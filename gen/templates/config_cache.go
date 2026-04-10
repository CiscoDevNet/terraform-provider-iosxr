//go:build ignore
// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/netascode/go-gnmi"
	"github.com/tidwall/gjson"
)

// Section below is generated&owned by "gen/generator.go". //template:begin allCachePaths

// allCachePaths contains one entry per unique module:/<first-path-segment> pair
// derived from every resource and data-source definition. Using only the first
// segment means a single gNMI Get per entry covers the entire subtree for that
// module. The list is filtered at runtime against gNMI Capabilities before the
// batch request is issued.
var allCachePaths = []string{
{{- range .}}
	"{{.}}",
{{- end}}
}

// End of section. //template:end allCachePaths

// ---------------------------------------------------------------------------
// DeviceCache
// ---------------------------------------------------------------------------

// DeviceCache holds cached gNMI Get responses keyed by "origin:/elem".
// Thread-safe via a single RWMutex.
type DeviceCache struct {
	mu        sync.RWMutex
	data      map[string][]byte
	fetchedAt map[string]int64
}

func NewDeviceCache() *DeviceCache {
	return &DeviceCache{
		data:      make(map[string][]byte),
		fetchedAt: make(map[string]int64),
	}
}

// Get returns cached bytes for path if still within TTL (0 = no expiry).
func (c *DeviceCache) Get(path string, ttlSeconds int64) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, ok := c.data[path]
	if !ok {
		return nil, false
	}
	if ttlSeconds > 0 && time.Now().Unix()-c.fetchedAt[path] >= ttlSeconds {
		return nil, false
	}
	return data, true
}

func (c *DeviceCache) Set(path string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[path] = data
	c.fetchedAt[path] = time.Now().Unix()
}

// Delete removes a path from the cache after a write operation.
func (c *DeviceCache) Delete(path string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, path)
	delete(c.fetchedAt, path)
}

// ---------------------------------------------------------------------------
// Capability-based path filtering
// ---------------------------------------------------------------------------

// FilterPathsByCapabilities calls gNMI Capabilities and returns the subset of
// allCachePaths whose origin module is advertised by the device.
func FilterPathsByCapabilities(ctx context.Context, client *gnmi.Client) []string {
	capRes, err := client.Capabilities(ctx)
	if err != nil {
		return allCachePaths
	}

	supported := make(map[string]struct{}, len(capRes.Models))
	for _, m := range capRes.Models {
		name := m.GetName()
		if strings.HasPrefix(name, "Cisco-IOS-XR-") {
			supported[name] = struct{}{}
		}
	}

	filtered := make([]string, 0, len(allCachePaths))
	for _, p := range allCachePaths {
		colonIdx := strings.Index(p, ":")
		if colonIdx < 0 {
			filtered = append(filtered, p)
			continue
		}
		if _, ok := supported[p[:colonIdx]]; ok {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// ---------------------------------------------------------------------------
// FetchAndCache — single batch gNMI Get
// ---------------------------------------------------------------------------

// FetchAndCache issues ONE batch gNMI Get with the supplied paths.
// IOS XR returns all data in a single notification with multiple Update entries.
// Each update's Path.Origin + Path.Elem[0] is used as the cache key.
func FetchAndCache(ctx context.Context, client *gnmi.Client, cache *DeviceCache, paths []string) error {
	if len(paths) == 0 {
		return nil
	}

	res, err := client.Get(ctx, paths)
	if err != nil {
		return err
	}

	for _, notif := range res.Notifications {
		for _, upd := range notif.Update {
			if upd.Path == nil || upd.Val == nil {
				continue
			}
			raw := upd.Val.GetJsonIetfVal()
			if len(raw) == 0 {
				continue
			}
			origin := upd.Path.GetOrigin()
			elems := upd.Path.GetElem()
			if origin == "" || len(elems) == 0 {
				continue
			}
			cache.Set(origin+":/"+elems[0].GetName(), raw)
		}
	}
	return nil
}

// WarmCache is a convenience wrapper: filters paths by capabilities then
// calls FetchAndCache in one shot.
func WarmCache(ctx context.Context, client *gnmi.Client, cache *DeviceCache, _ int64) error {
	paths := FilterPathsByCapabilities(ctx, client)
	if len(paths) == 0 {
		return nil
	}
	return FetchAndCache(ctx, client, cache, paths)
}

// ---------------------------------------------------------------------------
// GetFromCache — gjson-based cache lookup
// ---------------------------------------------------------------------------

// findRootPath returns the allCachePaths entry that best covers specificPath
// by longest-prefix match within the same module.
func findRootPath(specificPath string) string {
	colonIdx := strings.Index(specificPath, ":")
	if colonIdx < 0 {
		return specificPath
	}
	modulePrefix := specificPath[:colonIdx+1]
	pathPart := specificPath[colonIdx+1:]

	for _, r := range allCachePaths {
		if r == specificPath {
			return r
		}
	}

	best := ""
	for _, r := range allCachePaths {
		if !strings.HasPrefix(r, modulePrefix) {
			continue
		}
		rPart := r[colonIdx+1:]
		if strings.HasPrefix(pathPart, rPart) && len(r) > len(best) {
			best = r
		}
	}
	if best != "" {
		return best
	}

	re := regexp.MustCompile(`\[[^]]+]`)
	return modulePrefix + re.ReplaceAllString(pathPart, "")
}

// pathToGjsonQuery converts the relative path between specificPath and rootPath
// into a gjson query, handling both [key='value'] and [key=value] predicates.
func pathToGjsonQuery(specificPath, rootPath string) string {
	if specificPath == rootPath {
		return ""
	}
	colonIdx := strings.Index(specificPath, ":")
	if colonIdx < 0 {
		return ""
	}
	specPart := specificPath[colonIdx+1:]
	rootPart := rootPath[strings.Index(rootPath, ":")+1:]

	if !strings.HasPrefix(specPart, rootPart) {
		return ""
	}
	rel := strings.TrimPrefix(strings.TrimPrefix(specPart, rootPart), "/")
	if rel == "" {
		return ""
	}

	re := regexp.MustCompile(`\[([^=\]]+)=(?:'([^']*)'|([^\]']*))\]`)
	parts := strings.Split(rel, "/")
	var q strings.Builder
	first := true
	for _, part := range parts {
		if part == "" {
			continue
		}
		if !first {
			q.WriteByte('.')
		}
		first = false
		bracketIdx := strings.Index(part, "[")
		if bracketIdx >= 0 {
			matches := re.FindStringSubmatch(part[bracketIdx:])
			if len(matches) >= 3 {
				value := matches[2]
				if value == "" {
					value = matches[3]
				}
				q.WriteString(part[:bracketIdx])
				q.WriteString(".#(")
				q.WriteString(matches[1])
				q.WriteString("==\"")
				q.WriteString(value)
				q.WriteString("\")")
			} else {
				q.WriteString(part)
			}
		} else {
			q.WriteString(part)
		}
	}
	return q.String()
}

// GetFromCache retrieves data for specificPath from the cache using gjson
// filtering. Returns (nil, false) on cache miss.
func GetFromCache(ctx context.Context, cache *DeviceCache, specificPath string, ttlSeconds int64) ([]byte, bool) {
	rootPath := findRootPath(specificPath)

	cachedData, hit := cache.Get(rootPath, ttlSeconds)
	if !hit {
		return nil, false
	}

	if specificPath == rootPath {
		return cachedData, true
	}

	query := pathToGjsonQuery(specificPath, rootPath)
	if query == "" {
		return cachedData, true
	}

	result := gjson.GetBytes(cachedData, query)
	if !result.Exists() {
		return nil, false
	}
	return []byte(result.Raw), true
}

// ---------------------------------------------------------------------------
// ReadConfig — single entry point for all resource/data-source Read calls
// ---------------------------------------------------------------------------

// ReadConfig fetches configuration for resourcePath, using the cache when
// enabled. It replaces the duplicated fetch blocks in every generated Read.
//
//   - cacheEnabled=true: calls ensureWarmed (sync.Once, no-op after first call),
//     then returns cached JSON via gjson filtering.
//   - On cache miss or cacheEnabled=false: issues a single gNMI Get and stores
//     the result in cache for subsequent reads within the same refresh.
//   - notFound=true when the device returns "Requested element(s) not found".
//   - err!=nil for all other failures.
func ReadConfig(
	ctx context.Context,
	client *gnmi.Client,
	cache *DeviceCache,
	cacheEnabled bool,
	cacheTTL int64,
	ensureWarmed func(context.Context),
	resourcePath string,
) (body []byte, notFound bool, err error) {

	if cacheEnabled {
		ensureWarmed(ctx)
		if cached, hit := GetFromCache(ctx, cache, resourcePath, cacheTTL); hit {
			return cached, false, nil
		}
	}

	getResp, getErr := client.Get(ctx, []string{resourcePath})
	if getErr != nil {
		if strings.Contains(getErr.Error(), "Requested element(s) not found") {
			return nil, true, nil
		}
		return nil, false, fmt.Errorf("get: %w", getErr)
	}
	if len(getResp.Notifications) == 0 || len(getResp.Notifications[0].Update) == 0 {
		return nil, false, fmt.Errorf("gNMI response contains no data for %s", resourcePath)
	}

	raw := getResp.Notifications[0].Update[0].Val.GetJsonIetfVal()
	if cacheEnabled {
		cache.Set(resourcePath, raw)
	}
	return raw, false, nil
}


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

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	"Cisco-IOS-XR-um-aaa-cfg:/aaa",
	"Cisco-IOS-XR-um-route-policy-cfg:/routing-policy",
	"Cisco-IOS-XR-um-banner-cfg:/banners",
	"Cisco-IOS-XR-um-bfd-sbfd-cfg:/bfd",
	"Cisco-IOS-XR-um-router-bgp-cfg:/as-format",
	"Cisco-IOS-XR-um-router-bgp-cfg:/bmp",
	"Cisco-IOS-XR-um-call-home-cfg:/call-home",
	"Cisco-IOS-XR-um-cdp-cfg:/cdp",
	"Cisco-IOS-XR-um-cef-accounting-cfg:/cef",
	"Cisco-IOS-XR-8000-fib-platform-cfg:/cef",
	"Cisco-IOS-XR-um-policymap-classmap-cfg:/class-map",
	"Cisco-IOS-XR-um-cli-alias-cfg:/alias",
	"Cisco-IOS-XR-um-clock-exr-cfg:/clock",
	"Cisco-IOS-XR-um-control-plane-cfg:/control-plane",
	"Cisco-IOS-XR-ifmgr-cfg:/interface-configurations",
	"Cisco-IOS-XR-um-crypto-cfg:/crypto",
	"Cisco-IOS-XR-um-domain-cfg:/domain",
	"Cisco-IOS-XR-um-error-disable-cfg:/error-disable",
	"Cisco-IOS-XR-um-ethernet-cfm-cfg:/cfm",
	"Cisco-IOS-XR-um-ethernet-sla-cfg:/ethernet",
	"Cisco-IOS-XR-um-l2vpn-cfg:/evpn",
	"Cisco-IOS-XR-um-flow-cfg:/flow",
	"Cisco-IOS-XR-um-flow-cfg:/sampler-maps",
	"Cisco-IOS-XR-um-fpd-cfg:/fpd",
	"Cisco-IOS-XR-um-frequency-synchronization-cfg:/frequency",
	"Cisco-IOS-XR-um-ftp-tftp-cfg:/ftp",
	"Cisco-IOS-XR-um-l2vpn-cfg:/generic-interface-lists",
	"Cisco-IOS-XR-um-hostname-cfg:/hostname",
	"Cisco-IOS-XR-um-http-client-cfg:/http",
	"Cisco-IOS-XR-um-hw-module-profile-cfg:/hw-module",
	"Cisco-IOS-XR-um-8000-hw-module-profile-cfg:/hw-module",
	"Cisco-IOS-XR-um-hw-module-shut-cfg:/hw-module",
	"Cisco-IOS-XR-um-icmp-cfg:/icmp",
	"Cisco-IOS-XR-um-interface-cfg:/interfaces",
	"Cisco-IOS-XR-um-ipsla-cfg:/ipsla",
	"Cisco-IOS-XR-um-ipv4-access-list-cfg:/ipv4",
	"Cisco-IOS-XR-um-ipv4-prefix-list-cfg:/ipv4",
	"Cisco-IOS-XR-um-ipv6-cfg:/ipv6",
	"Cisco-IOS-XR-um-ipv6-access-list-cfg:/ipv6",
	"Cisco-IOS-XR-um-ipv6-prefix-list-cfg:/ipv6",
	"Cisco-IOS-XR-um-key-chain-cfg:/key",
	"Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn",
	"Cisco-IOS-XR-um-lacp-cfg:/lacp",
	"Cisco-IOS-XR-um-line-cfg:/line",
	"Cisco-IOS-XR-um-linux-networking-cfg:/linux",
	"Cisco-IOS-XR-um-lldp-cfg:/lldp",
	"Cisco-IOS-XR-um-logging-cfg:/logging",
	"Cisco-IOS-XR-um-interface-cfg:/logging",
	"Cisco-IOS-XR-um-lpts-punt-cfg:/lpts",
	"Cisco-IOS-XR-um-macsec-cfg:/macsec",
	"Cisco-IOS-XR-um-macsec-cfg:/macsec-policy",
	"Cisco-IOS-XR-um-monitor-session-cfg:/monitor-sessions",
	"Cisco-IOS-XR-um-mpls-ldp-cfg:/mpls",
	"Cisco-IOS-XR-um-mpls-oam-cfg:/mpls",
	"Cisco-IOS-XR-um-mpls-te-cfg:/mpls",
	"Cisco-IOS-XR-um-xml-agent-cfg:/netconf",
	"Cisco-IOS-XR-um-netconf-yang-cfg:/netconf-yang",
	"Cisco-IOS-XR-um-ntp-cfg:/ntp",
	"Cisco-IOS-XR-um-pce-cfg:/pce",
	"Cisco-IOS-XR-um-performance-measurement-cfg:/performance-measurement",
	"Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map",
	"Cisco-IOS-XR-um-ptp-cfg:/ptp",
	"Cisco-IOS-XR-um-router-bgp-cfg:/router",
	"Cisco-IOS-XR-um-router-hsrp-cfg:/router",
	"Cisco-IOS-XR-um-router-igmp-cfg:/router",
	"Cisco-IOS-XR-um-router-isis-cfg:/router",
	"Cisco-IOS-XR-um-router-mld-cfg:/router",
	"Cisco-IOS-XR-um-router-ospf-cfg:/router",
	"Cisco-IOS-XR-um-router-pim-cfg:/router",
	"Cisco-IOS-XR-um-router-static-cfg:/router",
	"Cisco-IOS-XR-um-router-vrrp-cfg:/router",
	"Cisco-IOS-XR-um-rsvp-cfg:/rsvp",
	"Cisco-IOS-XR-segment-routing-ms-cfg:/sr",
	"Cisco-IOS-XR-um-segment-routing-cfg:/segment-routing",
	"Cisco-IOS-XR-um-service-timestamps-cfg:/service",
	"Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server",
	"Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server-mibs",
	"Cisco-IOS-XR-um-vrf-cfg:/srlg",
	"Cisco-IOS-XR-um-ssh-cfg:/ssh",
	"Cisco-IOS-XR-um-tcp-cfg:/tcp",
	"Cisco-IOS-XR-um-telemetry-model-driven-cfg:/telemetry",
	"Cisco-IOS-XR-um-telnet-cfg:/telnet",
	"Cisco-IOS-XR-um-ftp-tftp-cfg:/tftp-fs",
	"Cisco-IOS-XR-um-ftp-tftp-cfg:/tftp",
	"Cisco-IOS-XR-um-tpa-cfg:/tpa",
	"Cisco-IOS-XR-um-track-cfg:/tracks",
	"Cisco-IOS-XR-um-vrf-cfg:/vrfs",
	"Cisco-IOS-XR-um-vty-pool-cfg:/vty-pool",
	"Cisco-IOS-XR-um-xml-agent-cfg:/xr-xml",
}

// End of section. //template:end allCachePaths

// ---------------------------------------------------------------------------
// DeviceCache
// ---------------------------------------------------------------------------

// DeviceCache holds cached gNMI Get responses keyed by the exact gNMI path.
// Each entry is the raw JSON-IETF bytes returned by the device for that path.
// Thread-safe via a single RWMutex.
//
// It also maintains a lazily-built, separate key->element index (listIndex)
// for fast point lookups into keyed-list blobs (e.g. "as-path-sets" with
// hundreds/thousands of "as-path-set" entries) -- see GetFromCache /
// singleTrailingFilter / getListIndex below for why this exists.
type DeviceCache struct {
	mu        sync.RWMutex
	data      map[string][]byte
	fetchedAt map[string]int64

	indexMu sync.RWMutex
	// listIndex caches, per "rootPath|containerQuery|keyField" combination,
	// a map from that keyed list's key VALUE to the matching element's raw
	// JSON. This turns N repeated per-resource GetFromCache lookups into
	// ONE array walk (O(n), done once) followed by N O(1) map lookups,
	// instead of re-running gjson's linear "#(key==value)" scan from
	// scratch on every single lookup (O(n) each -> O(n^2) total across n
	// sibling list resources, e.g. 1000 iosxr_as_path_set instances on one
	// device during a single `terraform plan`/`apply` refresh).
	listIndex map[string]map[string][]byte
}

// NewDeviceCache allocates an empty DeviceCache.
func NewDeviceCache() *DeviceCache {
	return &DeviceCache{
		data:      make(map[string][]byte),
		fetchedAt: make(map[string]int64),
	}
}

// Get returns the cached bytes for path and whether the entry is still within TTL.
// ttlSeconds == 0 means no expiry.
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

// Set stores bytes for path with the current timestamp. Invalidates the
// list-lookup index since any cached blob it was built from may have just
// changed (re-fetch after TTL expiry, cache miss fallback, etc.).
func (c *DeviceCache) Set(path string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[path] = data
	c.fetchedAt[path] = time.Now().Unix()
	c.invalidateListIndex()
}

// Delete removes a single path from the cache (used after write operations).
// Also invalidates the list-lookup index -- see Set.
//
// NOTE: c.data is keyed by "root" cache paths (one entry per
// module:/first-path-segment, e.g. "...-route-policy-cfg:/routing-policy"),
// as populated by FetchAndCache's batch warm -- NOT by the exact specific
// resource path (e.g. ".../as-path-set[set-name='X']") that
// Create/Update/Delete pass in here. ReadConfig's cache-miss fallback is the
// only path that ever stores under the exact specific path (its
// cache.Set(resourcePath, raw) call). So deleting only the literal `path`
// key is a no-op for the (common) case where the resource was actually
// served out of a root blob -- that stale blob would otherwise keep being
// returned as a cache HIT for this resource (and all its siblings) until
// the next full re-warm, contradicting the documented "cache is
// automatically invalidated after any write operation" behavior. Resolve
// and evict the root entry too so a write always forces a fresh fetch.
func (c *DeviceCache) Delete(path string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, path)
	delete(c.fetchedAt, path)
	if root := findRootPath(path); root != path {
		delete(c.data, root)
		delete(c.fetchedAt, root)
	}
	c.invalidateListIndex()
}

// invalidateListIndex drops the entire list-lookup index. Called whenever
// underlying cached data changes; the index is rebuilt lazily (and cheaply,
// one container at a time, only for containers actually queried) on the
// next GetFromCache call that needs it. Caller need not (but may) hold c.mu.
func (c *DeviceCache) invalidateListIndex() {
	c.indexMu.Lock()
	c.listIndex = nil
	c.indexMu.Unlock()
}

// getListIndex returns the value->element-JSON index for the keyed list at
// containerQuery (a dotted gjson path, e.g. "as-path-sets.as-path-set")
// within cachedData, building it on first use and reusing it for every
// subsequent lookup with the same indexKey until invalidated by a Set/Delete.
func (c *DeviceCache) getListIndex(indexKey string, cachedData []byte, containerQuery, keyField string) map[string][]byte {
	c.indexMu.RLock()
	if c.listIndex != nil {
		if idx, ok := c.listIndex[indexKey]; ok {
			c.indexMu.RUnlock()
			return idx
		}
	}
	c.indexMu.RUnlock()

	idx := make(map[string][]byte)
	arr := gjson.GetBytes(cachedData, containerQuery)
	if arr.IsArray() {
		arr.ForEach(func(_, el gjson.Result) bool {
			if keyVal := el.Get(keyField); keyVal.Exists() {
				idx[keyVal.String()] = []byte(el.Raw)
			}
			return true
		})
	}

	c.indexMu.Lock()
	if c.listIndex == nil {
		c.listIndex = make(map[string]map[string][]byte)
	}
	c.listIndex[indexKey] = idx
	c.indexMu.Unlock()

	return idx
}

// ---------------------------------------------------------------------------
// Capability-based path filtering
// ---------------------------------------------------------------------------

// FilterPathsByCapabilities calls gNMI Capabilities, builds the set of
// supported Cisco-IOS-XR-* module names, and returns the subset of
// allCachePaths whose origin module is advertised by the device.
//
// Example: "Cisco-IOS-XR-8000-fib-platform-cfg:/cef" is dropped when that
// module is not in the device capability list.
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
		origin := p[:colonIdx]
		if _, ok := supported[origin]; ok {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// FetchAndCache issues one batched gNMI Get across paths and stores each
// returned root subtree in cache. Retries a few times with backoff on
// transient errors (e.g. right after a large atomic batch commit, the
// device's datastore may still be converging and briefly reject or fail the
// very next Get) instead of giving up on the first attempt -- since this
// runs behind EnsureCacheWarmed's sync.Once, a single transient failure here
// would otherwise leave the cache permanently empty for the rest of this
// provider process, forcing every subsequent Read() to miss and fall back
// to individual live Gets (or, if those also race the same sync delay and
// are misread as "not found", a spurious full recreate -- see GetWithRetry).
func FetchAndCache(ctx context.Context, client *gnmi.Client, cache *DeviceCache, paths []string) error {
	if len(paths) == 0 {
		return nil
	}

	const maxRetries = 2
	baseDelay := 500 * time.Millisecond

	var res gnmi.GetRes
	var err error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		res, err = client.Get(ctx, paths)
		if err == nil {
			break
		}
		if attempt == maxRetries {
			tflog.Debug(ctx, fmt.Sprintf("device cache: batched Get for %d paths failed after %d attempt(s): %v", len(paths), attempt+1, err))
			return err
		}
		delay := baseDelay * time.Duration(1<<uint(attempt))
		tflog.Debug(ctx, fmt.Sprintf("device cache: batched Get for %d paths failed (attempt %d/%d): %v; retrying after %v in case device is still converging after a recent commit",
			len(paths), attempt+1, maxRetries+1, err, delay))
		time.Sleep(delay)
	}

	stored := 0
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
			stored++
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("device cache: 1 batched gNMI Get for %d paths stored %d cache entries", len(paths), stored))
	return nil
}

// ---------------------------------------------------------------------------
// GetFromCache — gjson-based cache lookup
// ---------------------------------------------------------------------------

// findRootPath returns the allCachePaths entry that best covers specificPath.
//
// It looks for the longest path in allCachePaths (same module prefix) that is a
// prefix of the specific path's path-part.
//
// Example:
//
//	specific = "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='Loopback100']"
//	→ root   = "Cisco-IOS-XR-um-interface-cfg:/interfaces"
func findRootPath(specificPath string) string {
	colonIdx := strings.Index(specificPath, ":")
	if colonIdx < 0 {
		return specificPath
	}
	modulePrefix := specificPath[:colonIdx+1]
	pathPart := specificPath[colonIdx+1:]

	// Exact match
	for _, r := range allCachePaths {
		if r == specificPath {
			return r
		}
	}

	// Longest-prefix match within the same module
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

	// Fallback: strip predicates and return
	re := regexp.MustCompile(`\[[^]]+]`)
	return modulePrefix + re.ReplaceAllString(pathPart, "")
}

// pathToGjsonQuery builds a gjson query that extracts the specific resource
// from the JSON blob stored under the root cache path.
//
// Examples:
//
//	specific = "…/interfaces/interface[interface-name='Loopback100']"
//	root     = "…/interfaces"
//	→ query  = "interface.#(interface-name==\"Loopback100\")"
//
//	specific = "…/ipv4/prefix-lists/prefix-list[prefix-list-name='PL1']"
//	root     = "…/ipv4"
//	→ query  = "prefix-lists.prefix-list.#(prefix-list-name==\"PL1\")"
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

	// Match both [key='value'] and [key=value] predicate forms
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
			container := part[:bracketIdx]
			matches := re.FindStringSubmatch(part[bracketIdx:])
			if len(matches) >= 3 {
				key := matches[1]
				// value is in group 2 (quoted) or group 3 (unquoted)
				value := matches[2]
				if value == "" {
					value = matches[3]
				}
				q.WriteString(container)
				q.WriteString(".#(")
				q.WriteString(key)
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

// trailingFilterRe matches a gjson query that ends in exactly one keyed-list
// filter, e.g. "as-path-sets.as-path-set.#(set-name==\"FOO\")" ->
// container="as-path-sets.as-path-set", key="set-name", value="FOO". Used by
// singleTrailingFilter to recognize the common flat-list case that the
// listIndex fast path (below) can serve in O(1) instead of O(n).
var trailingFilterRe = regexp.MustCompile(`^(.+)\.#\(([^=]+)==\"([^\"]*)\"\)$`)

// singleTrailingFilter reports whether query is EXACTLY ONE keyed-list
// filter at the end (no other "#(...)" anywhere in the query, i.e. not a
// nested keyed-list-within-keyed-list path). When true, it returns the
// container's dotted gjson path, the key field name, and the target value.
//
// This intentionally only optimizes the common single-level case (as used
// by as-path-sets, community-sets, prefix-lists, interfaces, etc.); nested
// multi-level filters fall back to the original per-call gjson scan below,
// which remains correct (just not accelerated) for those rarer paths.
func singleTrailingFilter(query string) (container, keyField, keyValue string, ok bool) {
	if strings.Count(query, "#(") != 1 {
		return "", "", "", false
	}
	m := trailingFilterRe.FindStringSubmatch(query)
	if m == nil {
		return "", "", "", false
	}
	return m[1], m[2], m[3], true
}

// GetFromCache retrieves data for a specific resource path from the cache.
//
// It finds the matching root path (e.g. "…/interfaces"), fetches the cached
// JSON blob for that root, then extracts the exact resource (e.g.
// interface[interface-name='Loopback100']) without any gNMI request.
//
// For the common case of a single trailing keyed-list filter, extraction
// uses a key->element index built once per (root, container) and reused for
// every sibling lookup (see DeviceCache.getListIndex) -- turning what would
// otherwise be N independent O(list length) gjson "#(...)" scans (O(n^2)
// total across n sibling resources, e.g. 1000 iosxr_as_path_set instances on
// one device) into one O(n) index build plus N O(1) map lookups. Other
// (nested/rarer) query shapes fall back to a direct gjson scan.
//
// Returns (nil, false) on cache miss.
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

	if container, keyField, keyValue, ok := singleTrailingFilter(query); ok {
		indexKey := rootPath + "|" + container + "|" + keyField
		idx := cache.getListIndex(indexKey, cachedData, container, keyField)
		raw, found := idx[keyValue]
		return raw, found
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

// ReadConfig fetches the configuration for resourcePath, using the cache when
// enabled. It replaces the duplicated if/else fetch blocks in every generated
// resource and data-source Read function.
//
// Behaviour:
//   - cacheEnabled=true: calls ensureWarmed (sync.Once, no-op after first call),
//     then returns cached JSON via gjson filtering.
//   - On cache miss or cacheEnabled=false: uses GetWithRetry (retry + exponential
//     backoff) to absorb post-Set device sync delays, exactly like the
//     pre-cache generated Read logic did, and stores the result in cache for
//     subsequent reads within the same refresh.
//   - notFound=true when the device returns "Requested element(s) not found"
//     (resource caller should remove itself from state).
//   - A successful but empty ({}) response (e.g. a keys-only list entry) is
//     NOT an error: body is returned empty ([]byte{}) with notFound=false so
//     callers can preserve state as-is (resources) or raise their own
//     "no data" error (data sources) — see IsEmptyRespBody.
//   - err!=nil for all other failures (caller should add a diagnostics error).
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
			tflog.Debug(ctx, fmt.Sprintf("device cache: HIT for %s (served from cache, no gNMI Get)", resourcePath))
			return cached, false, nil
		}
		tflog.Debug(ctx, fmt.Sprintf("device cache: MISS for %s (falling back to live gNMI Get)", resourcePath))
	}

	getResp, notFound, getErr := GetWithRetry(ctx, client, []string{resourcePath}, resourcePath)
	if getErr != nil {
		return nil, false, fmt.Errorf("get: %w", getErr)
	}
	if notFound {
		return nil, true, nil
	}

	if IsGnmiGetResponseEmpty(&getResp) {
		// A successful but empty ({}) response means the element exists but
		// the device returned no data (e.g. a keys-only list entry). Return
		// an empty body instead of erroring so callers can preserve state
		// as-is rather than treating this as a failure or deletion.
		if cacheEnabled {
			cache.Set(resourcePath, []byte{})
		}
		return []byte{}, false, nil
	}

	raw := getResp.Notifications[0].Update[0].Val.GetJsonIetfVal()

	if cacheEnabled {
		cache.Set(resourcePath, raw)
	}

	return raw, false, nil
}

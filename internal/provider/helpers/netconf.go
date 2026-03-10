// Copyright © 2025 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
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
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
)

// CloseNetconfConnection safely closes a NETCONF connection if reuse is disabled.
// When reuseConnection=true it is a no-op — the session stays open for subsequent operations.
//
// IMPORTANT: Must be called while the operation mutex is still held (deferred after
// AcquireNetconfLock) to prevent concurrent close attempts.
func CloseNetconfConnection(ctx context.Context, client *netconf.Client, reuseConnection bool) {
	if client == nil {
		return
	}
	if !reuseConnection {
		if err := client.Close(); err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to close NETCONF connection: %s", err))
		} else {
			tflog.Debug(ctx, "Successfully closed NETCONF connection")
		}
	}
}

// EnsureNetconfConnection checks if the NETCONF connection is open and reconnects if needed.
// When reuseConnection=true, it keeps the connection alive for better performance.
// When reuseConnection=false, it ensures a fresh connection for each operation.
func EnsureNetconfConnection(ctx context.Context, client *netconf.Client, reuseConnection bool, maxRetries int) error {
	if client == nil {
		return fmt.Errorf("NETCONF client is nil")
	}

	if maxRetries <= 0 {
		maxRetries = 3
	}

	if !reuseConnection {
		// No reuse: ensure connection is open before every operation
		return ensureOpenWithRetries(ctx, client, maxRetries)
	}

	// Fast path for reuse: if connection is already open, don't reconnect
	if !client.IsClosed() {
		return nil
	}

	// Connection is closed, need to reconnect
	tflog.Warn(ctx, "NETCONF session closed, reconnecting")
	return reconnectWithRetries(ctx, client, maxRetries)
}

// reconnectWithRetries attempts to reconnect a closed NETCONF session with exponential back-off
func reconnectWithRetries(ctx context.Context, client *netconf.Client, maxRetries int) error {
	retryDelay := 500 * time.Millisecond
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		_ = client.Close()
		InvalidateCapabilityCache(client)

		if attempt > 1 {
			delay := retryDelay * time.Duration(attempt)
			tflog.Debug(ctx, fmt.Sprintf("Reconnect attempt %d/%d, waiting %v", attempt, maxRetries, delay))
			time.Sleep(delay)
		}

		if err := openWithTimeout(ctx, client, openTimeout); err != nil {
			lastErr = err
			tflog.Warn(ctx, fmt.Sprintf("Reconnect attempt %d/%d failed: %s", attempt, maxRetries, err))
			continue
		}

		// Run health-check to confirm transport is working
		if err := netconfHealthCheck(ctx, client); err != nil {
			lastErr = err
			tflog.Warn(ctx, fmt.Sprintf("Reconnect attempt %d/%d health check failed: %s", attempt, maxRetries, err))
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("NETCONF connected on attempt %d", attempt))
		return nil
	}
	return fmt.Errorf("failed to connect NETCONF after %d attempts: %w", maxRetries, lastErr)
}

// AcquireNetconfLock acquires the appropriate lock for a NETCONF operation.
//
// Lock strategy based on reuseConnection and operation type:
// - When reuseConnection=false: Lock for ALL operations (serializes everything including Close)
// - When reuseConnection=true && isWrite=true: Lock for WRITE operations (Lock/EditConfig/Commit sequence)
// - When reuseConnection=true && isWrite=false: NO lock for READ operations (concurrent reads allowed)
//
// This prevents issues:
// - When reuse disabled: Prevents concurrent Close() attempts on same connection
// - When reuse enabled: Serializes write sequences but allows concurrent reads
//
// Returns true if lock was acquired, false if not acquired.
//
// Usage:
//
//	locked := helpers.AcquireNetconfLock(opMutex, reuseConnection, isWrite)
//	if locked {
//	    defer opMutex.Unlock()
//	}
//	defer helpers.CloseNetconfConnection(ctx, client, reuseConnection)
func AcquireNetconfLock(opMutex *sync.Mutex, reuseConnection bool, isWrite bool) bool {
	// Serialize all operations when reuse disabled (prevent concurrent Close)
	if !reuseConnection {
		opMutex.Lock()
		return true
	}

	// When reuse enabled, only serialize write operations
	// Read operations can run concurrently
	if isWrite {
		opMutex.Lock()
		return true
	}

	return false
}

// ensureOpenWithRetries opens a fresh connection (Close+Open) with retries.
// Used by the no-reuse path.
func ensureOpenWithRetries(ctx context.Context, client *netconf.Client, maxRetries int) error {
	retryDelay := 50 * time.Millisecond
	for attempt := 1; attempt <= maxRetries; attempt++ {
		_ = client.Close()
		InvalidateCapabilityCache(client)
		if attempt > 1 {
			time.Sleep(retryDelay * time.Duration(attempt))
		}
		if err := openWithTimeout(ctx, client, openTimeout); err != nil {
			if attempt < maxRetries {
				tflog.Warn(ctx, fmt.Sprintf("NETCONF Open attempt %d/%d failed: %s. Retrying...", attempt, maxRetries, err))
				continue
			}
			return fmt.Errorf("failed to open NETCONF connection after %d attempts: %w", maxRetries, err)
		}
		if attempt > 1 {
			tflog.Info(ctx, fmt.Sprintf("NETCONF connected on attempt %d", attempt))
		}
		return nil
	}
	return fmt.Errorf("failed to open NETCONF connection after %d attempts", maxRetries)
}

// openWithTimeout opens a NETCONF connection with a timeout.
// client.Open() can block indefinitely if the TCP connection is established
// but the SSH/NETCONF handshake never completes (e.g. slow or overloaded device).
func openWithTimeout(ctx context.Context, client *netconf.Client, timeout time.Duration) error {
	type result struct{ err error }
	ch := make(chan result, 1)
	go func() {
		ch <- result{err: client.Open()}
	}()
	tctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	select {
	case r := <-ch:
		return r.err
	case <-tctx.Done():
		// Best-effort close so the goroutine above can unblock.
		_ = client.Close()
		return fmt.Errorf("NETCONF Open timed out after %v", timeout)
	}
}

// openTimeout is the maximum time to wait for a NETCONF connection to open.
const openTimeout = 30 * time.Second

func netconfHealthCheck(ctx context.Context, client *netconf.Client) error {
	hctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	filter := GetSubtreeFilter("Cisco-IOS-XR-um-hostname-cfg:/hostname")
	_, err := client.GetConfig(hctx, "running", filter)
	if err == nil {
		return nil
	}

	errStr := strings.ToLower(err.Error())
	for _, token := range []string{"connection", "closed", "broken pipe", "eof", "timeout", "deadline exceeded"} {
		if strings.Contains(errStr, token) {
			return fmt.Errorf("transport error during health check: %w", err)
		}
	}
	// Non-transport NETCONF errors mean the session is alive.
	tflog.Debug(ctx, fmt.Sprintf("NETCONF health check got non-transport error (session OK): %s", err))
	return nil
}

// formatNetconfError extracts detailed error information from a NETCONF error.
func formatNetconfError(err error) string {
	if netconfErr, ok := err.(*netconf.NetconfError); ok {
		var b strings.Builder
		b.WriteString(netconfErr.Message)
		for i, e := range netconfErr.Errors {
			if i == 0 {
				b.WriteString("\n\nError Details:")
			}
			b.WriteString(fmt.Sprintf("\n  [%d] ", i+1))
			if e.ErrorMessage != "" {
				b.WriteString(e.ErrorMessage)
			}
			if e.ErrorPath != "" {
				b.WriteString(fmt.Sprintf(" (path: %s)", e.ErrorPath))
			}
			if e.ErrorType != "" || e.ErrorTag != "" {
				b.WriteString(fmt.Sprintf(" [type=%s, tag=%s]", e.ErrorType, e.ErrorTag))
			}
			if e.ErrorInfo != "" {
				b.WriteString(fmt.Sprintf("\n      Info: %s", e.ErrorInfo))
			}
		}
		return b.String()
	}
	return err.Error()
}

// candidateCapabilityURI is the IETF URN for the candidate datastore capability.
// IOS-XR always advertises this; it is the only supported write path.
const candidateCapabilityURI = "urn:ietf:params:netconf:capability:candidate:1.0"

// capabilityCache caches per-client capability check results.
// Key: "<pointer>:<uri>", Value: bool
var capabilityCache sync.Map

func cacheKey(client *netconf.Client, uri string) string {
	return fmt.Sprintf("%p:%s", client, uri)
}

func hasCapability(client *netconf.Client, uri string) bool {
	key := cacheKey(client, uri)
	if cached, ok := capabilityCache.Load(key); ok {
		return cached.(bool)
	}
	supported := client.ServerHasCapability(uri)
	capabilityCache.Store(key, supported)
	return supported
}

// InvalidateCapabilityCache removes all cached capability entries for the given client.
// Call this whenever the client reconnects so the next operation re-queries the server.
func InvalidateCapabilityCache(client *netconf.Client) {
	capabilityCache.Delete(cacheKey(client, candidateCapabilityURI))
}

// EditConfig edits the configuration on the device
// If the server supports the candidate capability, it will edit the configuration in the candidate datastore
// and commit it to the running datastore if commit is true.
// If the server does not support the candidate capability, it will edit the configuration in the running datastore.
//
// IMPORTANT: When connection reuse is enabled, callers MUST serialize calls to EditConfig using an
// application-level mutex that also covers EnsureNetconfConnectionDevice(). This prevents concurrent goroutines
// from attempting to acquire NETCONF datastore locks simultaneously on the same session.
//
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - body: string
//   - commit: bool
func EditConfig(ctx context.Context, client *netconf.Client, body string, commit bool) error {
	return EditConfigWithOptions(ctx, client, body, commit, false)
}

// EditConfigWithOptions edits the configuration on the device with additional options.
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - body: string
//   - commit: bool
//   - ignoreDataMissing: bool - if true, treats data-missing errors as success (useful for delete operations)
func EditConfigWithOptions(ctx context.Context, client *netconf.Client, body string, commit bool, ignoreDataMissing bool) error {
	if body == "" {
		tflog.Debug(ctx, "EditConfig called with empty body, skipping")
		return nil
	}

	// Ensure connection is open before checking capabilities
	// With lazy connections, Open() is idempotent and safe to call multiple times
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")

	if candidate {
		if commit {
			// Lock running datastore
			if _, err := client.Lock(ctx, "running"); err != nil {
				return fmt.Errorf("failed to lock running datastore: %s", formatNetconfError(err))
			}
			defer client.Unlock(ctx, "running")

			// Lock candidate datastore
			if _, err := client.Lock(ctx, "candidate"); err != nil {
				return fmt.Errorf("failed to lock candidate datastore: %s", formatNetconfError(err))
			}
			defer client.Unlock(ctx, "candidate")
		}

		tflog.Debug(ctx, "NETCONF edit-config body", map[string]interface{}{"xml": body})
		if _, err := client.EditConfig(ctx, "candidate", body); err != nil {
			// Check if this is a data-missing error and we should ignore it
			if ignoreDataMissing && isDataMissingError(err) {
				tflog.Debug(ctx, "Ignoring data-missing error during delete operation")
				return nil
			}
			return fmt.Errorf("failed to edit config: %s", formatNetconfError(err))
		}

		if commit {
			if _, err := client.Commit(ctx); err != nil {
				// Check if this is a "no changes to commit" error - not actually an error
				if isNoChangesToCommitError(err) {
					tflog.Debug(ctx, "No configuration changes to commit - continuing")
					return nil
				}
				// Check if this is a data-missing error and we should ignore it
				if ignoreDataMissing && isDataMissingError(err) {
					tflog.Debug(ctx, "Ignoring data-missing error during commit")
					return nil
				}
				return fmt.Errorf("failed to commit config: %s", formatNetconfError(err))
			}
		}
	} else {
		// Lock running datastore
		if _, err := client.Lock(ctx, "running"); err != nil {
			return fmt.Errorf("failed to lock running datastore: %s", formatNetconfError(err))
		}
		defer client.Unlock(ctx, "running")

		tflog.Debug(ctx, "NETCONF edit-config body", map[string]interface{}{"xml": body})
		if _, err := client.EditConfig(ctx, "running", body); err != nil {
			// Check if this is a data-missing error and we should ignore it
			if ignoreDataMissing && isDataMissingError(err) {
				tflog.Debug(ctx, "Ignoring data-missing error during delete operation")
				return nil
			}
			return fmt.Errorf("failed to edit config: %s", formatNetconfError(err))
		}
	}
	return nil
}

// isDataMissingError checks if a NETCONF error is a data-missing error
func isDataMissingError(err error) bool {
	if netconfErr, ok := err.(*netconf.NetconfError); ok {
		for _, e := range netconfErr.Errors {
			if e.ErrorTag == "data-missing" {
				return true
			}
		}
	}
	return false
}

// isNoChangesToCommitError checks if a NETCONF error is a "No configuration changes to commit" error
// This is not actually an error condition - it just means there were no changes to apply
func isNoChangesToCommitError(err error) bool {
	if netconfErr, ok := err.(*netconf.NetconfError); ok {
		for _, e := range netconfErr.Errors {
			if e.ErrorTag == "operation-failed" &&
				strings.Contains(e.ErrorMessage, "No configuration changes to commit") {
				return true
			}
		}
	}
	return false
}

// GetConfig retrieves configuration from the device.
// This is a wrapper around client.GetConfig for consistency with other helpers.
func GetConfig(ctx context.Context, client *netconf.Client, source string, filter netconf.Filter) (*netconf.Res, error) {
	res, err := client.GetConfig(ctx, source, filter)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetConfigWithTimeout is a backward compatibility wrapper for GetConfig
// Used by generated resources and data sources.
func GetConfigWithTimeout(ctx context.Context, client *netconf.Client, source string, filter netconf.Filter, timeout ...time.Duration) (netconf.Res, error) {
	// Default timeout of 30 seconds if not specified
	t := 30 * time.Second
	if len(timeout) > 0 {
		t = timeout[0]
	}

	// Create a context with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	res, err := GetConfig(timeoutCtx, client, source, filter)
	if err != nil {
		return netconf.Res{}, err
	}
	if res == nil {
		return netconf.Res{}, fmt.Errorf("nil response from GetConfig")
	}
	return *res, nil
}

// IsGetConfigResponseEmpty checks if a GetConfig response has an empty <data> element.
func IsGetConfigResponseEmpty(res *netconf.Res) bool {
	if res == nil {
		return true
	}
	rawXML := strings.TrimSpace(res.Res.Raw)
	if rawXML == "" {
		return true
	}
	if strings.Contains(rawXML, "<data/>") || strings.Contains(rawXML, "<data></data>") {
		return true
	}
	start := strings.Index(rawXML, "<data>")
	end := strings.Index(rawXML, "</data>")
	if start == -1 || end == -1 || start >= end {
		return true
	}
	content := strings.TrimSpace(rawXML[start+6 : end])
	return content == "" || !strings.Contains(content, "<")
}

// GetSubtreeFilter creates a NETCONF subtree filter from an XPath expression.
// IOS-XR devices do not support XPath filters, so we need to convert the XPath
// into a subtree filter.
func GetSubtreeFilter(xPath string) netconf.Filter {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	segments := splitXPathSegments(xPath)

	// Filter out empty segments (caused by // in XPath for augments)
	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	// Handle the case where the namespace is in a separate segment (e.g., "Cisco-IOS-XR-um-hostname-cfg:" followed by "hostname")
	var namespace string
	processedSegments := make([]string, 0, len(segments))
	for i, seg := range segments {
		if strings.HasSuffix(seg, ":") && i == 0 {
			namespace = strings.TrimSuffix(seg, ":")
		} else if seg != "" {
			processedSegments = append(processedSegments, seg)
		}
	}
	segments = processedSegments

	var content strings.Builder

	for i, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		var element string
		var segmentNamespace string
		if namespace == "" {
			if idx := strings.Index(elementName, ":"); idx != -1 {
				namespace = elementName[:idx]
				segmentNamespace = namespace
				element = elementName[idx+1:]
			} else {
				element = elementName
			}
		} else {
			// Check if this segment has its own namespace prefix
			if idx := strings.Index(elementName, ":"); idx != -1 {
				segmentNamespace = elementName[:idx]
				element = elementName[idx+1:]
			} else {
				element = removeNamespacePrefix(elementName)
			}
		}

		if element == "" {
			continue
		}

		content.WriteString("<")
		content.WriteString(element)

		// Add xmlns if this segment has a namespace (either from first segment or its own prefix)
		if i == 0 && namespace != "" {
			nsURL := getNamespaceURL(namespace)
			content.WriteString(fmt.Sprintf(" xmlns=\"%s\"", nsURL))
		} else if i > 0 && segmentNamespace != "" {
			nsURL := getNamespaceURL(segmentNamespace)
			content.WriteString(fmt.Sprintf(" xmlns=\"%s\"", nsURL))
		}

		if len(keys) > 0 {
			content.WriteString(">")
			for _, kv := range keys {
				keyName := removeNamespacePrefix(kv.Key)
				content.WriteString(fmt.Sprintf("<%s>%s</%s>", keyName, kv.Value, keyName))
			}
		} else {
			// Always use open tag (never self-closing). IOS-XR returns empty
			// <data></data> when a self-closing subtree filter element is used.
			content.WriteString(">")
		}
	}

	for i := len(segments) - 1; i >= 0; i-- {
		segment := segments[i]
		elementName, keys := parseXPathSegment(segment)
		element := removeNamespacePrefix(elementName)
		if element == "" {
			continue
		}
		if len(keys) > 0 || i < len(segments)-1 {
			content.WriteString("</")
			content.WriteString(element)
			content.WriteString(">")
		} else {
			// Close the last element (was previously self-closed with />)
			content.WriteString("</")
			content.WriteString(element)
			content.WriteString(">")
		}
	}

	return netconf.Filter{Type: "subtree", Content: content.String()}
}

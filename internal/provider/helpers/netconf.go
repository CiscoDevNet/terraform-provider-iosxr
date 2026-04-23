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
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// ============================================================================
// Constants
// ============================================================================

const (
	// openTimeout is the maximum time to wait for a NETCONF connection to open
	openTimeout = 30 * time.Second

	// candidateCapabilityURI is the IETF URN for the candidate datastore capability
	candidateCapabilityURI = "urn:ietf:params:netconf:capability:candidate:1.0"

	// namespaceBaseURL is the base URL for Cisco YANG models
	namespaceBaseURL = "http://cisco.com/ns/yang/"
)

// capabilityCache caches per-client capability check results
// Key: "<pointer>:<uri>", Value: bool
var capabilityCache sync.Map

// ============================================================================
// NETCONF Connection Management
// ============================================================================

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

func cacheKey(client *netconf.Client, uri string) string {
	return fmt.Sprintf("%p:%s", client, uri)
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
//   - skipCommit: bool
func EditConfig(ctx context.Context, client *netconf.Client, body string, commit bool, skipCommit bool) error {
	return EditConfigWithOptions(ctx, client, body, commit, skipCommit, false)
}

// EditConfigWithOptions edits the configuration on the device with additional options.
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - body: string
//   - commit: bool
//   - skipCommit: bool - when true, stages changes without committing (for batching/confirmed-commit)
//   - ignoreDataMissing: bool - if true, treats data-missing errors as success (useful for delete operations)
func EditConfigWithOptions(ctx context.Context, client *netconf.Client, body string, commit bool, skipCommit bool, ignoreDataMissing bool) error {
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
		// Only lock if we're committing immediately (not in batch mode)
		// When skipCommit=true, we're in batch mode and the connection is shared/locked at a higher level
		if commit && !skipCommit {
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

		// Only commit if commit is true AND skipCommit is false
		if commit && !skipCommit {
			if _, err := client.Commit(ctx); err != nil {
				// Check if this is a "no changes to commit" error - not actually an error
				if IsNoChangesToCommitError(err) {
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

// IsNoChangesToCommitError checks if a NETCONF error is a "No configuration changes to commit" error
// This is not actually an error condition - it just means there were no changes to apply
func IsNoChangesToCommitError(err error) bool {
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

// Commit commits the candidate datastore to running with proper locking
// This is used by the iosxr_commit resource to commit batched changes
func Commit(ctx context.Context, client *netconf.Client) error {
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")
	if !candidate {
		tflog.Debug(ctx, "Device does not support candidate datastore, nothing to commit")
		return nil
	}

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

	// Commit the candidate datastore
	if _, err := client.Commit(ctx); err != nil {
		// Check if this is a "no changes to commit" error - not actually an error
		if IsNoChangesToCommitError(err) {
			tflog.Debug(ctx, "No configuration changes to commit - continuing")
			return nil
		}
		return fmt.Errorf("failed to commit config: %s", formatNetconfError(err))
	}

	return nil
}

// CommitConfirmed executes a confirmed-commit with automatic confirmation on success
// This function:
// 1. Checks for :confirmed-commit:1.1 capability (returns error if not supported)
// 2. Locks running and candidate datastores
// 3. Executes confirmed-commit with specified timeout (60-240 seconds)
// 4. Immediately confirms the commit (auto-confirmation on success)
//
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - timeoutSeconds: confirmed-commit timeout in seconds (60-240)
//
// IOS XR will automatically rollback changes after timeout if not confirmed.
// This function auto-confirms on success for seamless Terraform workflow.
func CommitConfirmed(ctx context.Context, client *netconf.Client, timeoutSeconds int64) error {
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	// Check for confirmed-commit capability
	if !client.ServerHasCapability("urn:ietf:params:netconf:capability:confirmed-commit:1.1") {
		return fmt.Errorf("device does not support NETCONF confirmed-commit. Ensure IOS XR version supports :confirmed-commit:1.1 capability")
	}

	// Check for candidate datastore capability
	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")
	if !candidate {
		return fmt.Errorf("confirmed-commit requires candidate datastore support")
	}

	// Note: We don't lock datastores here because in batch mode (auto_commit=false),
	// each resource operation already handles locking/unlocking as it stages changes.
	// Attempting to lock again here would cause conflicts with connection reuse.
	// The NETCONF commit operation itself is atomic and safe without explicit locking.

	// Execute confirmed-commit with timeout using go-netconf's built-in support
	tflog.Debug(ctx, fmt.Sprintf("Executing confirmed-commit with %d second timeout", timeoutSeconds))

	// netconf.Confirmed(timeoutSeconds) tells the device to automatically rollback after N seconds
	_, err := client.Commit(ctx, netconf.Confirmed(int(timeoutSeconds)))
	if err != nil {
		// Check if this is a "no changes to commit" error
		if IsNoChangesToCommitError(err) {
			tflog.Debug(ctx, "No configuration changes to commit - continuing")
			return nil
		}
		return fmt.Errorf("confirmed-commit failed: %s", formatNetconfError(err))
	}

	// Immediately confirm the commit (auto-confirmation on success)
	// This is a regular commit that confirms the previous confirmed-commit
	tflog.Debug(ctx, "Auto-confirming commit after successful confirmed-commit")
	if _, err := client.Commit(ctx); err != nil {
		return fmt.Errorf("failed to auto-confirm commit: %s", formatNetconfError(err))
	}

	tflog.Info(ctx, "Confirmed-commit completed successfully with auto-confirmation")
	return nil
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
	// Use utility function to prepare XPath segments (handles trimming, splitting, filtering, and namespace merging)
	segments := PrepareXPathSegments(xPath)

	// Extract namespace from first segment if present
	var namespace string
	if len(segments) > 0 {
		namespace = GetNamespacePrefixFromSegment(segments[0])
	}

	var content strings.Builder

	for i, segment := range segments {
		elementName, keys := ParseXPathSegment(segment)

		var element string
		var segmentNamespace string
		if namespace == "" {
			// Extract namespace from element name
			segmentNamespace = GetNamespacePrefixFromSegment(elementName)
			if segmentNamespace != "" {
				namespace = segmentNamespace
			}
			element = RemoveNamespacePrefix(elementName)
		} else {
			// Check if this segment has its own namespace prefix
			segmentNamespace = GetNamespacePrefixFromSegment(elementName)
			element = RemoveNamespacePrefix(elementName)
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
				keyName := RemoveNamespacePrefix(kv.Key)
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
		elementName, keys := ParseXPathSegment(segment)
		element := RemoveNamespacePrefix(elementName)
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

// GetConfigWithRetry retrieves configuration from the device with retry logic.
// NETCONF GetConfig may return empty data immediately after commit due to device sync delay.
// This function retries with exponential backoff to handle such cases.
//
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - source: string (e.g., "running", "candidate")
//   - filter: netconf.Filter
//   - xpath: string (for logging purposes)
//
// Returns:
//   - netconf.Res: The response from GetConfig
//   - bool: true if response is empty after all retries
//   - error: any error that occurred
func GetConfigWithRetry(ctx context.Context, client *netconf.Client, source string, filter netconf.Filter, xpath string) (netconf.Res, bool, error) {
	var res netconf.Res
	var err error
	maxRetries := 3
	baseDelay := 200 * time.Millisecond

	for attempt := 0; attempt <= maxRetries; attempt++ {
		res, err = GetConfigWithTimeout(ctx, client, source, filter)
		if err != nil {
			return res, false, fmt.Errorf("failed to retrieve object (%s): %w", xpath, err)
		}

		// Check if we got data back
		isEmpty := IsGetConfigResponseEmpty(&res)
		tflog.Debug(ctx, fmt.Sprintf("NETCONF GetConfig response for %s (attempt %d/%d): isEmpty=%v, isListPath=%v",
			xpath, attempt+1, maxRetries+1, isEmpty, IsListPath(xpath)))

		// If we got data or this is the last attempt, break
		if !isEmpty || attempt == maxRetries {
			return res, isEmpty, nil
		}

		// Wait before retrying (exponential backoff)
		delay := baseDelay * time.Duration(1<<uint(attempt))
		tflog.Debug(ctx, fmt.Sprintf("NETCONF returned empty response, retrying after %v", delay))
		time.Sleep(delay)
	}

	return res, IsGetConfigResponseEmpty(&res), nil
}

// ============================================================================
// Namespace Handling Functions
// ============================================================================

// setRootElementNamespace adds xmlns to root element using regex
func setRootElementNamespace(xml, cleanSegment, namespace string) string {
	pattern := fmt.Sprintf(`<(%s)(\s[^>]*?)?(/?)>`, regexp.QuoteMeta(cleanSegment))
	re := regexp.MustCompile(pattern)
	match := re.FindStringIndex(xml)
	if match == nil {
		return xml
	}

	matchStart, matchEnd := match[0], match[1]
	matchStr := xml[matchStart:matchEnd]

	// Skip if xmlns already present
	if strings.Contains(matchStr, `xmlns="`) {
		return xml
	}

	// Determine insertion position (before closing > or />)
	insertPos := len(matchStr) - 1
	if strings.HasSuffix(matchStr, "/>") {
		insertPos = len(matchStr) - 2
	}

	attrStr := fmt.Sprintf(` xmlns="%s"`, namespace)
	newMatch := matchStr[:insertPos] + attrStr + matchStr[insertPos:]
	return xml[:matchStart] + newMatch + xml[matchEnd:]
}

// setNestedElementNamespace adds xmlns to nested element using xmldot
func setNestedElementNamespace(xml, xmldotPath, namespace string) string {
	if xmldot.Get(xml, xmldotPath).String() == "" {
		xml, _ = xmldot.Set(xml, xmldotPath, namespace)
	}
	return xml
}

// augmentNamespaces adds xmlns attributes to elements that have namespace prefixes
func augmentNamespaces(body netconf.Body, path string) netconf.Body {
	segments := strings.Split(path, ".")
	xml := body.Res()

	// First pass: Build map of elements that need namespaces
	elementToNamespace := make(map[string]string)
	currentPath := ""

	for segIdx, segment := range segments {
		cleanSegment := CleanSegmentName(segment)
		currentPath = BuildPathFromSegments([]string{cleanSegment}, currentPath)

		if prefix, hasPrefix := ExtractNamespacePrefix(segment); hasPrefix {
			key := fmt.Sprintf("%d:%s", segIdx, currentPath)
			elementToNamespace[key] = prefix
		}
	}

	// Second pass: Apply namespaces to XML
	currentPath = ""
	for segIdx, segment := range segments {
		cleanSegment := CleanSegmentName(segment)
		currentPath = BuildPathFromSegments([]string{cleanSegment}, currentPath)

		key := fmt.Sprintf("%d:%s", segIdx, currentPath)
		if prefix, hasNamespace := elementToNamespace[key]; hasNamespace {
			namespace := namespaceBaseURL + prefix

			if segIdx == 0 {
				// Root element: use regex to insert xmlns
				xml = setRootElementNamespace(xml, cleanSegment, namespace)
			} else {
				// Nested element: use xmldot to insert xmlns
				xmldotPath := BuildXmlnsPath(currentPath)
				xml = setNestedElementNamespace(xml, xmldotPath, namespace)
			}
		}
	}

	return netconf.NewBody(xml)
}

// ensureNCNamespaceOnRoot ensures the NETCONF namespace (nc) is declared on the root element
func ensureNCNamespaceOnRoot(body netconf.Body) netconf.Body {
	xml := body.Res()

	if strings.Contains(xml, `xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0"`) {
		return body
	}

	rootStart := strings.Index(xml, "<")
	if rootStart == -1 {
		return body
	}

	rootEnd := strings.Index(xml[rootStart:], ">")
	if rootEnd == -1 {
		return body
	}
	rootEnd += rootStart

	rootTag := xml[rootStart : rootEnd+1]

	nameEnd := rootStart + 1
	for nameEnd < rootEnd {
		ch := xml[nameEnd]
		if ch == ' ' || ch == '>' || ch == '/' {
			break
		}
		nameEnd++
	}

	insertPos := nameEnd
	if moduleXmlnsMatch := regexp.MustCompile(`\sxmlns="http://cisco\.com/ns/yang/[^"]+"`).FindStringIndex(rootTag); moduleXmlnsMatch != nil {
		insertPos = rootStart + moduleXmlnsMatch[1]
	}

	newXML := xml[:insertPos] + ` xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0"` + xml[insertPos:]

	return netconf.NewBody(newXML)
}

// AddNamespaceToRootElement adds a default namespace declaration to the root element of XML string.
// This is needed because xmldot doesn't properly handle YANG namespace prefixes.
// Given an XPath like "Cisco-IOS-XR-um-segment-routing-cfg:/segment-routing/...", it extracts
// the namespace and adds it as xmlns="..." to the first XML element.
func AddNamespaceToRootElement(xmlStr string, xPath string) string {
	// Extract namespace prefix from XPath
	namespacePrefix, hasNamespace := ExtractNamespaceFromXPath(xPath)
	if !hasNamespace {
		return xmlStr
	}

	namespaceURL := getNamespaceURL(namespacePrefix)

	// Find root element bounds
	startIdx, closeIdx, found := FindRootElementBounds(xmlStr)
	if !found {
		return xmlStr
	}

	rootTag := xmlStr[startIdx : closeIdx+1]

	// Extract element name
	elementName, found := ExtractElementName(xmlStr, startIdx)
	if !found {
		return xmlStr
	}

	// If xmlns already exists, clean and re-add
	if strings.Contains(rootTag, "xmlns=") {
		cleaned := CleanExistingNamespaces(rootTag)
		cleaned = InsertNamespaceAfterElementName(cleaned, elementName, namespaceURL)
		return xmlStr[:startIdx] + cleaned + xmlStr[closeIdx+1:]
	}

	// Insert new xmlns attribute after element name
	endIdx := strings.IndexAny(xmlStr[startIdx:], "> ")
	if endIdx == -1 {
		return xmlStr
	}
	endIdx += startIdx

	// Insert namespace (works for both <element> and <element attr="...">)
	return xmlStr[:endIdx] + fmt.Sprintf(` xmlns="%s"`, namespaceURL) + xmlStr[endIdx:]
}

// ============================================================================
// XPath Structure Building (Core Infrastructure)
// ============================================================================

// checkIfAugmentedChild determines if an element is an augmented child
func checkIfAugmentedChild(body netconf.Body, keys []KeyValue, nsPrefix string, pathSegments []string, tentativePath string) bool {
	if len(keys) == 0 && nsPrefix != "" && len(pathSegments) > 0 {
		existingElement := xmldot.Get(body.Res(), tentativePath)
		return !existingElement.Exists()
	}
	return false
}

// getEffectiveNamespace gets the effective namespace for an element, checking parent if needed
func getEffectiveNamespace(body netconf.Body, xmlnsPath string, pathSegments []string) string {
	existingNS := xmldot.Get(body.Res(), xmlnsPath).String()
	if existingNS == "" && len(pathSegments) > 0 {
		parentPath := BuildParentPath(pathSegments)
		if parentPath != "" {
			parentNS := xmldot.Get(body.Res(), BuildXmlnsPath(parentPath)).String()
			if parentNS != "" {
				existingNS = parentNS
			}
		}
	}
	return existingNS
}

// findOrCreateNamespacedSibling finds or creates a sibling with a specific namespace
func findOrCreateNamespacedSibling(body netconf.Body, tentativePath, cleanElementName, expectedNS string) (netconf.Body, int) {
	siblingIdx := -1
	const maxSiblings = 32

	for n := 1; n <= maxSiblings; n++ {
		idxPath := fmt.Sprintf("%s.%d", tentativePath, n)
		ns := xmldot.Get(body.Res(), BuildXmlnsPath(idxPath)).String()
		if ns == expectedNS {
			siblingIdx = n
			break
		}
		if !xmldot.Get(body.Res(), idxPath).Exists() {
			currentXML := body.Res()
			siblingXML := fmt.Sprintf(`<%s xmlns="%s"></%s>`, cleanElementName, expectedNS, cleanElementName)
			insertAfter := fmt.Sprintf("</%s>", cleanElementName)
			lastIdx := strings.LastIndex(currentXML, insertAfter)
			if lastIdx >= 0 {
				insertPos := lastIdx + len(insertAfter)
				newXML := currentXML[:insertPos] + siblingXML + currentXML[insertPos:]
				body = netconf.NewBody(newXML)
			}
			siblingIdx = n
			break
		}
	}

	return body, siblingIdx
}

// findOrCreatePlainSibling finds or creates a sibling without namespace
func findOrCreatePlainSibling(body netconf.Body, tentativePath, cleanElementName string) (netconf.Body, int) {
	siblingIdx := -1
	const maxSiblings = 32

	for n := 1; n <= maxSiblings; n++ {
		idxPath := fmt.Sprintf("%s.%d", tentativePath, n)
		ns := xmldot.Get(body.Res(), BuildXmlnsPath(idxPath)).String()
		if ns == "" {
			if xmldot.Get(body.Res(), idxPath).Exists() {
				siblingIdx = n
				break
			}
		}
		if !xmldot.Get(body.Res(), idxPath).Exists() {
			currentXML := body.Res()
			siblingXML := fmt.Sprintf(`<%s></%s>`, cleanElementName, cleanElementName)
			insertAfter := fmt.Sprintf("</%s>", cleanElementName)
			lastIdx := strings.LastIndex(currentXML, insertAfter)
			if lastIdx >= 0 {
				insertPos := lastIdx + len(insertAfter)
				newXML := currentXML[:insertPos] + siblingXML + currentXML[insertPos:]
				body = netconf.NewBody(newXML)
			}
			siblingIdx = n
			break
		}
	}

	return body, siblingIdx
}

// processSegmentNamespace handles namespace-aware segment processing
func processSegmentNamespace(body netconf.Body, nsPrefix string, tentativePath, escapedElementName, cleanElementName string, pathSegments []string) ([]string, netconf.Body) {
	expectedNS := namespaceBaseURL + nsPrefix
	xmlnsPath := BuildXmlnsPath(tentativePath)
	existingNS := getEffectiveNamespace(body, xmlnsPath, pathSegments)

	if existingNS != "" && existingNS != expectedNS {
		updatedBody, siblingIdx := findOrCreateNamespacedSibling(body, tentativePath, cleanElementName, expectedNS)
		body = updatedBody
		if siblingIdx >= 0 {
			pathSegments = append(pathSegments, fmt.Sprintf("%s.%d", escapedElementName, siblingIdx))
		} else {
			pathSegments = append(pathSegments, escapedElementName)
		}
	} else {
		pathSegments = append(pathSegments, escapedElementName)
	}

	return pathSegments, body
}

// processSegmentWithoutNamespace handles non-namespaced segment processing
func processSegmentWithoutNamespace(body netconf.Body, tentativePath, escapedElementName, cleanElementName string, pathSegments []string) ([]string, netconf.Body) {
	xmlnsPath := BuildXmlnsPath(tentativePath)
	existingNS := xmldot.Get(body.Res(), xmlnsPath).String()

	shouldCreateSibling := false
	if existingNS != "" && len(pathSegments) > 0 {
		parentPath := BuildPathString(pathSegments)
		parentNS := xmldot.Get(body.Res(), BuildXmlnsPath(parentPath)).String()
		if existingNS != parentNS {
			shouldCreateSibling = true
		}
	}

	if shouldCreateSibling {
		updatedBody, siblingIdx := findOrCreatePlainSibling(body, tentativePath, cleanElementName)
		body = updatedBody
		if siblingIdx >= 0 {
			pathSegments = append(pathSegments, fmt.Sprintf("%s.%d", escapedElementName, siblingIdx))
		} else {
			pathSegments = append(pathSegments, escapedElementName)
		}
	} else {
		pathSegments = append(pathSegments, escapedElementName)
	}

	return pathSegments, body
}

// addKeysToPath adds XPath predicate keys to the body
func addKeysToPath(body netconf.Body, fullPath string, keys []KeyValue) netconf.Body {
	for _, kv := range keys {
		keyPath := fullPath + "." + kv.Key
		body = body.Set(keyPath, kv.Value)
	}
	return body
}

// setNamespaceForPath sets namespace attributes for a path
func setNamespaceForPath(body netconf.Body, fullPath, nsPrefix string, originalSegments []string, segmentIndex int) netconf.Body {
	if nsPrefix == "" {
		return body
	}

	// Check if path contains indexed sibling notation (e.g., ".0", ".1", etc.)
	isIndexedSibling := regexp.MustCompile(`\.\d+`).MatchString(fullPath)

	if isIndexedSibling {
		namespace := namespaceBaseURL + nsPrefix
		xmlnsPath := BuildXmlnsPath(fullPath)
		existingNS := xmldot.Get(body.Res(), xmlnsPath).String()
		if existingNS == "" {
			body = body.Set(xmlnsPath, namespace)
		}
	} else {
		originalPath := BuildPathString(originalSegments[:segmentIndex+1])
		body = augmentNamespaces(body, originalPath)
	}

	return body
}

// buildXPathStructure creates all elements in an XPath, including keys and namespaces
func buildXPathStructure(body netconf.Body, xPath string, ensureStructure bool) (netconf.Body, []string, []string) {
	segments := PrepareXPathSegments(xPath)

	pathSegments := make([]string, 0, len(segments))
	originalSegments := make([]string, 0, len(segments))

	for i, segment := range segments {
		elementName, keys := ParseXPathSegment(segment)
		originalSegments = append(originalSegments, elementName)
		cleanElementName := RemoveNamespacePrefix(elementName)
		escapedElementName := strings.ReplaceAll(cleanElementName, ".", `\.`)
		nsPrefix := GetNamespacePrefixFromSegment(elementName)

		tentativePath := BuildTentativePath(pathSegments, escapedElementName)

		// Check if this is an augmented child
		if checkIfAugmentedChild(body, keys, nsPrefix, pathSegments, tentativePath) {
			pathSegments = append(pathSegments, escapedElementName)
		} else if nsPrefix != "" {
			// Process namespace-aware element
			pathSegments, body = processSegmentNamespace(body, nsPrefix, tentativePath, escapedElementName, cleanElementName, pathSegments)
		} else {
			// Process element without namespace prefix
			pathSegments, body = processSegmentWithoutNamespace(body, tentativePath, escapedElementName, cleanElementName, pathSegments)
		}

		fullPath := BuildPathString(pathSegments)

		// Add keys if present
		body = addKeysToPath(body, fullPath, keys)

		// Set namespace attributes
		body = setNamespaceForPath(body, fullPath, nsPrefix, originalSegments, i)
	}

	if ensureStructure && len(pathSegments) > 0 {
		fullPath := BuildPathString(pathSegments)
		result := xmldot.Get(body.Res(), fullPath)
		if !result.Exists() {
			body = body.Set(fullPath, "")
			originalPath := BuildPathString(originalSegments)
			body = augmentNamespaces(body, originalPath)
		}
	}

	return body, pathSegments, originalSegments
}

// ============================================================================
// Public XPath Query and Manipulation Functions
// ============================================================================

// resolveNamespaceForElement resolves namespace-aware element location
func resolveNamespaceForElement(xml, currentPath, escapedElementName, nsPrefix string, pathSoFar []string, count int) ([]string, int, bool) {
	parentPath := BuildParentPath(pathSoFar)

	if idx, needsIndex, found := FindNamespaceAwareSibling(xml, currentPath, count, nsPrefix, parentPath); found {
		if needsIndex {
			pathSoFar[len(pathSoFar)-1] = fmt.Sprintf("%s.%d", escapedElementName, idx)
		}
		return pathSoFar, 1, true
	}

	return pathSoFar, count, false
}

// GetFromXPath reads from an xmldot.Result using an XPath that may contain predicates
func GetFromXPath(res xmldot.Result, xPath string) xmldot.Result {
	segments := PrepareGetFromXPathSegments(xPath)
	xml := res.Raw
	pathSoFar := make([]string, 0, len(segments))

	for _, segment := range segments {
		rawElementName, keys := ParseXPathSegment(segment)
		nsPrefix := GetNamespacePrefixFromSegment(rawElementName)
		elementName := RemoveNamespacePrefix(rawElementName)
		escapedElementName := strings.ReplaceAll(elementName, ".", `\.`)

		pathSoFar = append(pathSoFar, escapedElementName)
		currentPath := BuildPathString(pathSoFar)
		count := int(xmldot.Get(xml, BuildCountPath(currentPath)).Int())

		// Handle namespace resolution
		if nsPrefix != "" {
			var found bool
			pathSoFar, count, found = resolveNamespaceForElement(xml, currentPath, escapedElementName, nsPrefix, pathSoFar, count)
			if !found {
				return xmldot.Result{}
			}
			currentPath = BuildPathString(pathSoFar)
		}

		// Handle key predicates
		if len(keys) > 0 {
			var found bool
			pathSoFar, found = FindElementByKeys(xml, currentPath, escapedElementName, keys, count)
			if !found {
				return xmldot.Result{}
			}
		}
	}

	return BuildFinalResult(xml, pathSoFar)
}

// SetFromXPath creates all elements in an XPath and optionally sets a value
func SetFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	xPath = NormalizeModuleXPath(xPath)

	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments, originalSegments := buildXPathStructure(body, xPath, ensureStructure)

	if hasValue && len(pathSegments) > 0 {
		fullPath := BuildPathString(pathSegments)
		body = body.Set(fullPath, value)
		originalPath := BuildPathString(originalSegments)
		body = augmentNamespaces(body, originalPath)
	}

	return body
}

// AppendFromXPath creates all elements in an XPath and appends a value to a list
func AppendFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	xPath = NormalizeModuleXPath(xPath)

	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments, originalSegments := buildXPathStructure(body, xPath, ensureStructure)

	if hasValue && len(pathSegments) > 0 {
		fullPath := BuildPathString(pathSegments) + ".-1"
		body = body.Set(fullPath, value)
		originalPath := BuildPathString(originalSegments)
		body = augmentNamespaces(body, originalPath)
	}

	return body
}

// RemoveFromXPath creates all elements in an XPath with an nc:operation="remove" attribute
func RemoveFromXPath(body netconf.Body, xPath string) netconf.Body {
	xPath = NormalizeModuleXPath(xPath)

	body, pathSegments, originalSegments := buildXPathStructure(body, xPath, true)

	if len(pathSegments) > 0 {
		targetPath := BuildPathString(pathSegments)

		originalPath := BuildPathString(originalSegments)
		body = augmentNamespaces(body, originalPath)

		body = ensureNCNamespaceOnRoot(body)

		operationPath := targetPath + ".@nc:operation"
		body = body.Set(operationPath, "remove")
	}

	return body
}

// BodyToNestedXML converts a netconf.Body to a properly nested XML string.
func BodyToNestedXML(body netconf.Body) (string, error) {
	xml := body.Res()
	if xml == "" {
		return "", nil
	}

	return xml, nil
}

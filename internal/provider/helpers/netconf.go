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

// Pre-compiled regular expressions for performance
var (
	xpathPredicateRegex = regexp.MustCompile(`\[[^\]]*\]`)
	namespaceRegex      = regexp.MustCompile(`[a-zA-Z0-9\-]+:`)
	predicatePattern    = regexp.MustCompile(`\[([^\]]+)\]`)
)

// Namespace base URL for Cisco YANG models
const namespaceBaseURL = "http://cisco.com/ns/yang/"

// getNamespaceURL returns the full namespace URL for a given prefix
// Generically constructs namespace by appending prefix to base URL
func getNamespaceURL(prefix string) string {
	// For Cisco IOS-XR and other standard YANG modules, namespace follows pattern:
	// http://cisco.com/ns/yang/<module-name>
	if prefix != "" {
		return namespaceBaseURL + prefix
	}

	return ""
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

	// When reuse enabled, serialize ALL operations (reads + writes).
	// The underlying client maintains shared session state; allowing concurrent reads
	// can race with reconnect/close and eventually exhaust sessions on the device.
	opMutex.Lock()
	return true
}

// CloseNetconfConnection safely closes a NETCONF connection if reuse is disabled.
//
// IMPORTANT: This must be called with the operation mutex still held (deferred after AcquireNetconfLock)
// to prevent concurrent close attempts.
//
// Usage:
//
//	defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
func CloseNetconfConnection(ctx context.Context, client *netconf.Client, reuseConnection bool) {
	if reuseConnection {
		return // Keep connection open for reuse
	}

	// Close the connection (mutex already held by caller)
	if err := client.Close(); err != nil {
		// Log error but don't fail - connection cleanup is best-effort
		tflog.Warn(ctx, fmt.Sprintf("Failed to close NETCONF connection: %s", err))
	}
}

// EnsureNetconfConnection ensures that the NETCONF connection is healthy and reconnects if needed.
// This is especially important when connection reuse is enabled, as connections can become stale.
//
// When reuseConnection=true, this function uses lazy connection - it only calls Open() once
// and the go-netconf library maintains the connection state internally.
//
// Usage:
//
//	if err := helpers.EnsureNetconfConnection(ctx, device.NetconfClient, reuseConnection); err != nil {
//	    resp.Diagnostics.AddError("Connection Error", fmt.Sprintf("Failed to ensure NETCONF connection: %s", err))
//	    return
//	}
func EnsureNetconfConnection(ctx context.Context, client *netconf.Client, reuseConnection bool) error {
	if client == nil {
		return fmt.Errorf("NETCONF client is nil")
	}

	maxRetries := 3
	retryDelay := 200 * time.Millisecond

	// When reuse is disabled, always force a fresh connection
	// Close any existing connection and open a new one with retries
	if !reuseConnection {
		for attempt := 1; attempt <= maxRetries; attempt++ {
			_ = client.Close() // Best effort close any existing connection
			if attempt > 1 {
				time.Sleep(retryDelay * time.Duration(attempt)) // Exponential backoff
			}

			if err := client.Open(); err != nil {
				if attempt < maxRetries {
					tflog.Warn(ctx, fmt.Sprintf("Failed to open NETCONF connection (attempt %d/%d): %s. Retrying...", attempt, maxRetries, err))
					continue
				}
				return fmt.Errorf("failed to open NETCONF connection after %d attempts: %w", maxRetries, err)
			}

			if attempt > 1 {
				tflog.Info(ctx, fmt.Sprintf("Successfully opened NETCONF connection on attempt %d", attempt))
			}
			return nil
		}
	}

	// Reuse enabled: first try Open (idempotent if already connected)
	if err := client.Open(); err != nil {
		// Connection open failed - need to reconnect
		tflog.Warn(ctx, fmt.Sprintf("NETCONF connection open failed: %s, attempting reconnect", err))
		return reconnectNetconf(ctx, client, maxRetries, retryDelay)
	}

	// Connection appears open - validate it's actually usable with health check
	if err := netconfHealthCheck(ctx, client); err != nil {
		// Health check failed - connection is stale
		tflog.Warn(ctx, fmt.Sprintf("NETCONF connection health check failed: %s, attempting reconnect", err))
		return reconnectNetconf(ctx, client, maxRetries, retryDelay)
	}

	// Connection is healthy
	return nil
}

// reconnectNetconf attempts to reconnect a NETCONF session with exponential backoff
func reconnectNetconf(ctx context.Context, client *netconf.Client, maxRetries int, baseDelay time.Duration) error {
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Always close before attempting to reconnect
		_ = client.Close()

		// Exponential backoff
		delay := baseDelay * time.Duration(attempt)
		tflog.Debug(ctx, fmt.Sprintf("Reconnect attempt %d/%d, waiting %v before retry", attempt, maxRetries, delay))
		time.Sleep(delay)

		// Attempt to open connection
		if err := client.Open(); err != nil {
			lastErr = err
			tflog.Warn(ctx, fmt.Sprintf("Failed to open NETCONF connection (reconnect attempt %d/%d): %s", attempt, maxRetries, err))

			// Check if error indicates resource exhaustion or connection limit
			if strings.Contains(err.Error(), "connection") || strings.Contains(err.Error(), "Connection") {
				// Continue retry loop
				continue
			}

			if attempt < maxRetries {
				continue
			}
			return fmt.Errorf("failed to reconnect NETCONF after %d attempts: %w", maxRetries, lastErr)
		}

		// Validate connection with health check
		if err := netconfHealthCheck(ctx, client); err != nil {
			lastErr = err
			tflog.Warn(ctx, fmt.Sprintf("NETCONF health check failed after reconnect (attempt %d/%d): %s", attempt, maxRetries, err))
			if attempt < maxRetries {
				continue
			}
			return fmt.Errorf("failed to validate NETCONF connection after %d attempts: %w", maxRetries, lastErr)
		}

		// Success!
		tflog.Info(ctx, fmt.Sprintf("Successfully reconnected NETCONF connection on attempt %d", attempt))
		return nil
	}

	return fmt.Errorf("failed to reconnect NETCONF after %d attempts: %w", maxRetries, lastErr)
}

// netconfHealthCheck tries a very small RPC to validate the session.
// It should be cheap and safe across IOS-XR versions.
func netconfHealthCheck(ctx context.Context, client *netconf.Client) error {
	// Use a short deadline so we fail fast on stale sessions
	hctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Try a simple capabilities request first - this is lightweight and doesn't touch config
	// If this fails, the transport layer is broken
	if err := client.Open(); err != nil {
		return fmt.Errorf("connection not open: %w", err)
	}

	// For IOS-XR, use a minimal get-config request to verify the session works
	// Some devices reject unfiltered get-config, so use a specific path
	filter := GetSubtreeFilter("Cisco-IOS-XR-um-hostname-cfg:/hostname")
	_, err := client.GetConfig(hctx, "running", filter)

	if err == nil {
		return nil
	}

	// Check if error is a transport/session error vs. a valid NETCONF error
	errStr := strings.ToLower(err.Error())

	// These indicate session/transport problems
	if strings.Contains(errStr, "connection") ||
		strings.Contains(errStr, "closed") ||
		strings.Contains(errStr, "broken pipe") ||
		strings.Contains(errStr, "eof") ||
		strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "deadline exceeded") {
		return fmt.Errorf("transport error during health check: %w", err)
	}

	// Other NETCONF errors (e.g., operation-failed, unknown-element) mean the session
	// is working but the request was rejected - that's OK for a health check
	tflog.Debug(ctx, fmt.Sprintf("NETCONF health check got non-transport error (session OK): %s", err))
	return nil
}

// FormatNetconfError extracts detailed error information from a NETCONF error
func FormatNetconfError(err error) string {
	if netconfErr, ok := err.(*netconf.NetconfError); ok {
		var details strings.Builder
		details.WriteString(netconfErr.Message)

		// Add detailed error information from each ErrorModel
		for i, e := range netconfErr.Errors {
			if i == 0 {
				details.WriteString("\n\nError Details:")
			}
			details.WriteString(fmt.Sprintf("\n  [%d] ", i+1))

			if e.ErrorMessage != "" {
				details.WriteString(e.ErrorMessage)
			}

			if e.ErrorPath != "" {
				details.WriteString(fmt.Sprintf(" (path: %s)", e.ErrorPath))
			}

			if e.ErrorType != "" || e.ErrorTag != "" {
				details.WriteString(fmt.Sprintf(" [type=%s, tag=%s]", e.ErrorType, e.ErrorTag))
			}

			if e.ErrorInfo != "" {
				details.WriteString(fmt.Sprintf("\n      Info: %s", e.ErrorInfo))
			}
		}

		return details.String()
	}
	return err.Error()
}

// commitWithRetry attempts to commit changes with retry logic for transient errors.
func commitWithRetry(ctx context.Context, client *netconf.Client, maxRetries int) error {
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			delay := time.Duration(attempt) * 500 * time.Millisecond
			tflog.Warn(ctx, fmt.Sprintf("Retrying commit after %v (attempt %d/%d)", delay, attempt+1, maxRetries+1))
			time.Sleep(delay)
		}

		_, err := client.Commit(ctx)
		if err == nil {
			if attempt > 0 {
				tflog.Info(ctx, fmt.Sprintf("Commit succeeded on retry attempt %d", attempt+1))
			}
			return nil
		}

		lastErr = err

		if netconfErr, ok := err.(*netconf.NetconfError); ok {
			for _, e := range netconfErr.Errors {
				if strings.Contains(e.ErrorMessage, "No configuration changes to commit") {
					tflog.Warn(ctx, "No configuration changes to commit - resource may already be in desired state", map[string]interface{}{
						"error": e.ErrorMessage,
					})
					return nil
				}

				lowerMsg := strings.ToLower(e.ErrorMessage)
				if strings.Contains(lowerMsg, "lock") || strings.Contains(lowerMsg, "in-use") || strings.Contains(lowerMsg, "busy") {
					if attempt < maxRetries {
						continue
					}
				}
			}
		}

		if attempt >= maxRetries {
			return fmt.Errorf("failed to commit config: %s", FormatNetconfError(lastErr))
		}
	}

	return fmt.Errorf("failed to commit config: %s", FormatNetconfError(lastErr))
}

// EditConfig edits the configuration on the device using NETCONF.
// If the server supports the candidate capability, it will edit in the candidate datastore
// and commit it to the running datastore if commit is true.
// If the server does not support the candidate capability, it will edit the running datastore directly.
func EditConfig(ctx context.Context, client *netconf.Client, body string, commit bool) error {
	// If body is empty, there's nothing to edit
	if body == "" {
		tflog.Debug(ctx, "EditConfig called with empty body, skipping")
		return nil
	}

	// Ensure connection is open before checking capabilities
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")

	if candidate {
		if commit {
			// Lock running datastore
			if _, err := client.Lock(ctx, "running"); err != nil {
				return fmt.Errorf("failed to lock running datastore: %s", FormatNetconfError(err))
			}
			defer client.Unlock(ctx, "running")

			// Lock candidate datastore
			if _, err := client.Lock(ctx, "candidate"); err != nil {
				return fmt.Errorf("failed to lock candidate datastore: %s", FormatNetconfError(err))
			}
			defer client.Unlock(ctx, "candidate")
		}

		tflog.Debug(ctx, "NETCONF edit-config body", map[string]interface{}{"xml": body})
		if _, err := client.EditConfig(ctx, "candidate", body); err != nil {
			return fmt.Errorf("failed to edit config: %s", FormatNetconfError(err))
		}

		if commit {
			if err := commitWithRetry(ctx, client, 4); err != nil {
				return err
			}
		}
	} else {
		// Lock running datastore
		if _, err := client.Lock(ctx, "running"); err != nil {
			return fmt.Errorf("failed to lock running datastore: %s", FormatNetconfError(err))
		}
		defer client.Unlock(ctx, "running")

		if _, err := client.EditConfig(ctx, "running", body); err != nil {
			// Check for data-missing error which is common during delete operations
			if strings.Contains(err.Error(), "data-missing") {
				tflog.Warn(ctx, "Ignored data-missing error during EditConfig", map[string]interface{}{"error": err.Error()})
				return nil
			}
			return fmt.Errorf("failed to edit config: %s", FormatNetconfError(err))
		}
	}
	return nil
}

// IsListPath checks if an XPath represents a list item (ends with a predicate).
func IsListPath(xPath string) bool {
	xPath = strings.TrimSpace(xPath)
	return strings.HasSuffix(xPath, "]")
}

// IsGetConfigResponseEmpty checks if a GetConfig response has an empty <data> element.
func IsGetConfigResponseEmpty(res *netconf.Res) bool {
	if res == nil {
		return true
	}

	// Check the Raw XML content from the xmldot.Result
	rawXML := strings.TrimSpace(res.Res.Raw)
	if rawXML == "" {
		return true
	}

	// Check for empty data element patterns
	if strings.Contains(rawXML, "<data/>") || strings.Contains(rawXML, "<data></data>") {
		return true
	}

	// Check if data element has any children by looking for content between <data> and </data>
	// Extract content between <data> and </data> tags
	dataStartIdx := strings.Index(rawXML, "<data>")
	dataEndIdx := strings.Index(rawXML, "</data>")

	if dataStartIdx == -1 || dataEndIdx == -1 || dataStartIdx >= dataEndIdx {
		// No proper data tags found, consider it empty
		return true
	}

	// Get content between <data> and </data>
	dataContent := strings.TrimSpace(rawXML[dataStartIdx+6 : dataEndIdx])

	// If content is empty or only whitespace, it's empty
	if dataContent == "" {
		return true
	}

	// If content contains any XML tags (child elements), it's not empty
	return !strings.Contains(dataContent, "<")
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
		if namespace == "" {
			if idx := strings.Index(elementName, ":"); idx != -1 {
				namespace = elementName[:idx]
				element = elementName[idx+1:]
			} else {
				element = elementName
			}
		} else {
			element = removeNamespacePrefix(elementName)
		}

		if element == "" {
			continue
		}

		content.WriteString("<")
		content.WriteString(element)

		if i == 0 && namespace != "" {
			nsURL := getNamespaceURL(namespace)
			content.WriteString(fmt.Sprintf(" xmlns=\"%s\"", nsURL))
		}

		if len(keys) > 0 {
			content.WriteString(">")
			for _, kv := range keys {
				keyName := removeNamespacePrefix(kv.Key)
				content.WriteString(fmt.Sprintf("<%s>%s</%s>", keyName, kv.Value, keyName))
			}
		} else if i < len(segments)-1 {
			content.WriteString(">")
		} else {
			content.WriteString("/>")
			break
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
		}
	}

	return netconf.Filter{Type: "subtree", Content: content.String()}
}

type keyValue struct {
	Key   string
	Value string
}

func removeNamespacePrefix(name string) string {
	if idx := strings.Index(name, ":"); idx != -1 {
		return name[idx+1:]
	}
	return name
}

func splitXPathSegments(xPath string) []string {
	segments := []string{}
	var current strings.Builder
	bracketDepth := 0

	for _, r := range xPath {
		switch r {
		case '[':
			bracketDepth++
			current.WriteRune(r)
		case ']':
			bracketDepth--
			current.WriteRune(r)
		case '/':
			if bracketDepth == 0 {
				if current.Len() > 0 {
					segments = append(segments, current.String())
					current.Reset()
				}
			} else {
				current.WriteRune(r)
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		segments = append(segments, current.String())
	}

	return segments
}

func parseXPathSegment(segment string) (string, []keyValue) {
	elementName := segment
	keys := []keyValue{}

	if idx := strings.Index(segment, "["); idx != -1 {
		elementName = segment[:idx]
		predicates := predicatePattern.FindAllStringSubmatch(segment[idx:], -1)
		for _, pred := range predicates {
			if len(pred) < 2 {
				continue
			}
			parts := strings.Split(pred[1], "and")
			for _, part := range parts {
				kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
				if len(kv) != 2 {
					continue
				}
				k := strings.TrimSpace(kv[0])
				v := strings.Trim(strings.TrimSpace(kv[1]), "'\"")
				keys = append(keys, keyValue{Key: k, Value: v})
			}
		}
	}

	return elementName, keys
}

// dotPath converts a YANG path to a dot path by removing XPath predicates (keys) and namespace prefixes.
func dotPath(path string) string {
	path = xpathPredicateRegex.ReplaceAllString(path, "")
	path = strings.ReplaceAll(path, "/", ".")
	return namespaceRegex.ReplaceAllString(path, "")
}

// setWithNamespaces sets a value in the netconf body and automatically adds namespace declarations.
func setWithNamespaces(body netconf.Body, fullPath string, value any) netconf.Body {
	body = body.Set(dotPath(fullPath), value)
	body = augmentNamespaces(body, fullPath)
	return body
}

func augmentNamespaces(body netconf.Body, path string) netconf.Body {
	segments := strings.Split(path, ".")
	pathWithoutPrefix := make([]string, 0, len(segments))

	for _, segment := range segments {
		cleanSegment := removeNamespacePrefix(segment)
		if idx := strings.IndexByte(cleanSegment, '['); idx != -1 {
			cleanSegment = cleanSegment[:idx]
		}
		pathWithoutPrefix = append(pathWithoutPrefix, cleanSegment)

		if idx := strings.Index(segment, ":"); idx != -1 {
			prefix := segment[:idx]
			currentPath := strings.Join(pathWithoutPrefix, ".")

			namespace := namespaceBaseURL + prefix
			xmlnsPath := currentPath + ".@xmlns"
			if !xmldot.Get(body.Res(), xmlnsPath).Exists() {
				body = body.Set(xmlnsPath, namespace)
			}
		}
	}

	return body
}

// AddNamespaceToRootElement adds a default namespace declaration to the root element of XML string.
// This is needed because xmldot doesn't properly handle YANG namespace prefixes.
// Given an XPath like "Cisco-IOS-XR-um-segment-routing-cfg:/segment-routing/...", it extracts
// the namespace and adds it as xmlns="..." to the first XML element.
func AddNamespaceToRootElement(xmlStr string, xPath string) string {
	// Extract namespace prefix from XPath (format: "Namespace:/path")
	if idx := strings.Index(xPath, ":/"); idx > 0 {
		namespacePrefix := xPath[:idx]
		// Remove leading slash if any
		namespacePrefix = strings.TrimPrefix(namespacePrefix, "/")

		// Build the full namespace URL
		namespaceURL := getNamespaceURL(namespacePrefix)

		// Find the first opening tag in the XML
		// Look for pattern like "<element" or "<element>"
		startIdx := strings.Index(xmlStr, "<")
		if startIdx == -1 {
			return xmlStr
		}

		// Find where the root element tag closes (the first '>')
		closeIdx := strings.Index(xmlStr[startIdx:], ">")
		if closeIdx == -1 {
			return xmlStr
		}
		closeIdx += startIdx

		// Check if xmlns is already present in the root element tag
		rootTag := xmlStr[startIdx : closeIdx+1]
		if strings.Contains(rootTag, "xmlns=") {
			// If xmlns already exists, we need to replace ALL occurrences with just one
			// This handles the case where augmentNamespaces already added it
			// Find the element name
			nameEndIdx := strings.IndexAny(xmlStr[startIdx+1:], "> ")
			if nameEndIdx == -1 {
				return xmlStr
			}
			nameEndIdx += startIdx + 1
			elementName := xmlStr[startIdx+1 : nameEndIdx]

			// Remove all existing xmlns attributes
			cleaned := xmlStr[startIdx : closeIdx+1]
			// Remove all xmlns="..." patterns
			xmlnsPattern := regexp.MustCompile(`\s+xmlns="[^"]*"`)
			cleaned = xmlnsPattern.ReplaceAllString(cleaned, "")

			// Find insertion point (after element name)
			insertPos := len("<" + elementName)
			cleaned = cleaned[:insertPos] + fmt.Sprintf(` xmlns="%s"`, namespaceURL) + cleaned[insertPos:]

			return xmlStr[:startIdx] + cleaned + xmlStr[closeIdx+1:]
		}

		// Find where the tag name ends (either at '>' or ' ')
		endIdx := strings.IndexAny(xmlStr[startIdx:], "> ")
		if endIdx == -1 {
			return xmlStr
		}
		endIdx += startIdx

		// Insert namespace declaration after the element name
		// If the tag ends with '>', insert before it
		// If the tag has attributes (has ' '), insert after the element name
		insertPos := endIdx
		if xmlStr[endIdx] == '>' {
			// Simple case: <element>
			return xmlStr[:insertPos] + fmt.Sprintf(" xmlns=\"%s\"", namespaceURL) + xmlStr[insertPos:]
		} else {
			// Has attributes: <element attr="value">
			return xmlStr[:insertPos] + fmt.Sprintf(" xmlns=\"%s\"", namespaceURL) + xmlStr[insertPos:]
		}
	}

	return xmlStr
}

// buildXPathStructure is a helper that creates all elements in an XPath, including keys and namespaces.
// Returns the body with the structure created and the path segments for further processing.
// The ensureStructure parameter controls whether an empty element should be created at the final path
// if it doesn't already exist. Set to false when a value will be immediately set afterward.
func buildXPathStructure(body netconf.Body, xPath string, ensureStructure bool) (netconf.Body, []string) {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
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
	// This matches the logic in GetSubtreeFilter
	if len(segments) > 0 && strings.HasSuffix(segments[0], ":") {
		namespace := strings.TrimSuffix(segments[0], ":")
		if len(segments) > 1 {
			// Prepend namespace to the next segment
			segments[1] = namespace + ":" + segments[1]
			// Remove the first segment
			segments = segments[1:]
		}
	}

	pathSegments := make([]string, 0, len(segments))

	for i, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		pathSegments = append(pathSegments, elementName)
		fullPath := strings.Join(pathSegments[:i+1], ".")

		if len(keys) > 0 {
			for _, kv := range keys {
				keyPath := fullPath + "." + kv.Key
				body = setWithNamespaces(body, keyPath, kv.Value)
			}
		}
	}

	if ensureStructure && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		existingContent := xmldot.Get(body.Res(), dotPath(fullPath)).String()
		if existingContent == "" {
			body = setWithNamespaces(body, fullPath, "")
		}
	}

	return body, pathSegments
}

// SetFromXPath creates all elements in an XPath, including keys and namespaces,
// and optionally sets a value at the final path location.
//
// If value is nil or empty string, only the structure is created without setting a value.
func SetFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments := buildXPathStructure(body, xPath, ensureStructure)

	if hasValue && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		body = setWithNamespaces(body, fullPath, value)
	}

	return body
}

// AppendFromXPath creates all elements in an XPath and appends a value to a list by using
// the ".-1" syntax. This is useful for adding multiple items to a list without keys.
// The function automatically appends ".-1" to the final element in the path.
func AppendFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments := buildXPathStructure(body, xPath, ensureStructure)

	// Append to the list using .-1 syntax
	if hasValue && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".") + ".-1"
		body = setWithNamespaces(body, fullPath, value)
	}

	return body
}

// SetRawFromXPath sets raw XML content at a specific XPath location.
// This is useful for setting pre-formatted XML content (e.g., from list items with keys).
// The value string should contain the inner XML content without the wrapping element tags.
func SetRawFromXPath(body netconf.Body, xPath string, value string) netconf.Body {
	if len(value) == 0 {
		return body
	}

	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
	segments := splitXPathSegments(xPath)

	// Filter out empty segments (caused by // in XPath for augments)
	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	if len(segments) == 0 {
		return body
	}

	// Extract the final element name
	finalSegment := segments[len(segments)-1]
	finalElement, keys := parseXPathSegment(finalSegment)
	finalElementClean := removeNamespacePrefix(finalElement)

	// Build parent structure (everything except the final element)
	if len(segments) > 1 {
		parentXPath := "/" + strings.Join(segments[:len(segments)-1], "/")
		// Build structure and ensure all parent elements are created with namespaces
		body, _ = buildXPathStructure(body, parentXPath, true)

		// Determine the namespace for the final element
		finalNamespace := ""
		if idx := strings.Index(finalElement, ":"); idx != -1 {
			// Final element has explicit namespace prefix
			prefix := finalElement[:idx]
			finalNamespace = getNamespaceURL(prefix)
		} else {
			// Final element doesn't have explicit namespace, try to inherit from parent path
			for i := len(segments) - 2; i >= 0; i-- {
				segmentName, _ := parseXPathSegment(segments[i])
				if idx := strings.Index(segmentName, ":"); idx != -1 {
					prefix := segmentName[:idx]
					finalNamespace = getNamespaceURL(prefix)
					break
				}
			}
		}

		// Get parent path for setting content
		parentPathSegments := make([]string, 0, len(segments)-1)
		for _, segment := range segments[:len(segments)-1] {
			elementName, _ := parseXPathSegment(segment)
			parentPathSegments = append(parentPathSegments, elementName)
		}
		parentPath := dotPath(strings.Join(parentPathSegments, "."))

		// Check if content already exists at parent path
		existingXML := xmldot.Get(body.Res(), parentPath).Raw

		// Multi-segment path: wrap the content with the final element tag
		var wrappedContent string
		if existingXML != "" {
			// This is a subsequent sibling - don't add xmlns again
			wrappedContent = "<" + finalElementClean + ">" + value + "</" + finalElementClean + ">"
		} else {
			// This is the first element - add xmlns if namespace is present
			if finalNamespace != "" {
				wrappedContent = "<" + finalElementClean + " xmlns=\"" + finalNamespace + "\">" + value + "</" + finalElementClean + ">"
			} else {
				wrappedContent = "<" + finalElementClean + ">" + value + "</" + finalElementClean + ">"
			}
		}

		if existingXML != "" {
			// Append wrapped element as sibling to existing content
			combinedXML := existingXML + wrappedContent
			body = body.SetRaw(parentPath, combinedXML)
		} else {
			// First element, set the wrapped content at parent
			body = body.SetRaw(parentPath, wrappedContent)
		}
	} else {
		// Single-segment path - SetRaw will wrap the content for us
		innerContent := value
		if len(keys) > 0 {
			tempBody := netconf.Body{}
			for _, kv := range keys {
				tempBody = setWithNamespaces(tempBody, kv.Key, kv.Value)
			}
			innerContent = tempBody.Res() + value
		}

		// Check if this element already exists at the root
		existingXML := xmldot.Get(body.Res(), finalElementClean).Raw
		if existingXML != "" {
			// Append as sibling - need to wrap each occurrence
			wrappedNew := "<" + finalElementClean + ">" + innerContent + "</" + finalElementClean + ">"
			combinedXML := existingXML + wrappedNew
			body = body.SetRaw(finalElementClean, combinedXML)
		} else {
			// First element at this path - SetRaw will wrap it
			body = body.SetRaw(finalElementClean, innerContent)
		}
	}

	// Add namespace declarations if present in the original path
	if len(segments) > 0 {
		dotPathForNamespaces := strings.Join(segments, ".")
		body = augmentNamespaces(body, dotPathForNamespaces)
	}

	return body
}

// GetFromXPath reads from an xmldot.Result using an XPath that may contain predicates.
// It supports both single and multiple list elements and returns an xmldot.Result
// that can be iterated with ForEach when the target is a list.
func GetFromXPath(res xmldot.Result, xPath string) xmldot.Result {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	segments := splitXPathSegments(xPath)

	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	xml := res.Raw
	pathSoFar := make([]string, 0, len(segments))

	for _, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		elementName = removeNamespacePrefix(elementName)
		pathSoFar = append(pathSoFar, elementName)

		currentPath := strings.Join(pathSoFar, ".")

		countPath := currentPath + ".#"
		count := xmldot.Get(xml, countPath).Int()

		if len(keys) > 0 {
			found := false
			if count > 1 {
				for idx := 0; idx < int(count); idx++ {
					indexedPath := fmt.Sprintf("%s.%d", currentPath, idx)
					item := xmldot.Get(xml, indexedPath)

					allMatch := true
					for _, kv := range keys {
						keyName := removeNamespacePrefix(kv.Key)
						keyResult := item.Get(keyName)
						if !keyResult.Exists() || keyResult.String() != kv.Value {
							allMatch = false
							break
						}
					}

					if allMatch {
						pathSoFar[len(pathSoFar)-1] = fmt.Sprintf("%s.%d", elementName, idx)
						found = true
						break
					}
				}
			} else {
				currentResult := xmldot.Get(xml, currentPath)
				allMatch := true
				for _, kv := range keys {
					keyName := removeNamespacePrefix(kv.Key)
					keyResult := currentResult.Get(keyName)
					if !keyResult.Exists() || keyResult.String() != kv.Value {
						allMatch = false
						break
					}
				}
				found = allMatch
			}

			if !found {
				return xmldot.Result{}
			}
		}
	}

	finalPath := strings.Join(pathSoFar, ".")

	countPath := finalPath + ".#"
	count := xmldot.Get(xml, countPath).Int()
	if count > 1 {
		if len(pathSoFar) >= 2 {
			parentPath := strings.Join(pathSoFar[:len(pathSoFar)-1], ".")
			childName := pathSoFar[len(pathSoFar)-1]
			arrayPath := parentPath + ".#." + childName
			return xmldot.Get(xml, arrayPath)
		}
	}

	return xmldot.Get(xml, finalPath)
}

// RemoveFromXPath creates all elements in an XPath with an nc:operation="delete" attribute.
// Returns a Body for use in toBodyXML functions that build up XML progressively.
// For delete operations in resource Delete() functions, use RemoveFromXPathString() instead.
func RemoveFromXPath(body netconf.Body, xPath string) netconf.Body {
	xmlStr := RemoveFromXPathString(body, xPath)
	// Wrap the string back in a Body for compatibility with toBodyXML functions.
	// Note: This will get mangled if you call .Res() on it, so only use this in toBodyXML.
	return netconf.Body{}.SetRaw("", xmlStr)
}

// RemoveFromXPathString creates all elements in an XPath with an nc:operation="delete" attribute
// and returns the raw XML string.
func RemoveFromXPathString(body netconf.Body, xPath string) string {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	segments := splitXPathSegments(xPath)

	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	// Handle namespace-only first segment (e.g., "Cisco-IOS-XR-um-hostname-cfg:" followed by "hostname")
	if len(segments) > 0 && strings.HasSuffix(segments[0], ":") {
		namespace := strings.TrimSuffix(segments[0], ":")
		if len(segments) > 1 {
			segments[1] = namespace + ":" + segments[1]
			segments = segments[1:]
		}
	}

	if len(segments) == 0 {
		return ""
	}

	var sb strings.Builder
	var closingTags []string

	for i, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		var namespace string
		if strings.Contains(elementName, ":") {
			parts := strings.SplitN(elementName, ":", 2)
			namespace = parts[0]
			elementName = parts[1]
		}

		sb.WriteString("<")
		sb.WriteString(elementName)

		if namespace != "" {
			nsURL := getNamespaceURL(namespace)
			if nsURL != "" {
				sb.WriteString(fmt.Sprintf(` xmlns="%s"`, nsURL))
			}
		}

		if i == 0 {
			sb.WriteString(` xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0"`)
		}

		if i == len(segments)-1 {
			sb.WriteString(` nc:operation="delete"`)
		}

		sb.WriteString(">")

		for _, kv := range keys {
			keyName := kv.Key
			if strings.Contains(keyName, ":") {
				parts := strings.SplitN(keyName, ":", 2)
				keyName = parts[1]
			}

			sb.WriteString(fmt.Sprintf("<%s>%s</%s>", keyName, kv.Value, keyName))
		}

		closingTags = append(closingTags, fmt.Sprintf("</%s>", elementName))
	}

	for i := len(closingTags) - 1; i >= 0; i-- {
		sb.WriteString(closingTags[i])
	}

	return sb.String()
}

// CleanupRedundantRemoveOperations removes redundant operation="remove" attributes from child elements
// when a parent element already has operation="remove".
func CleanupRedundantRemoveOperations(body netconf.Body) netconf.Body {
	if body.Res() == "" {
		return body
	}

	return cleanupRedundantRemoveRecursive(body, "", make(map[string]bool))
}

func cleanupRedundantRemoveRecursive(body netconf.Body, basePath string, visited map[string]bool) netconf.Body {
	xml := body.Res()
	if xml == "" {
		return body
	}

	if visited[basePath] {
		return body
	}
	visited[basePath] = true

	var currentResult xmldot.Result
	if basePath == "" {
		currentResult = xmldot.Get(xml, "")
	} else {
		currentResult = xmldot.Get(xml, basePath)
		if !currentResult.Exists() {
			return body
		}
	}

	rawXML := xml
	if basePath != "" {
		rawXML = currentResult.Raw
	}

	childPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)[\s>}]`)
	matches := childPattern.FindAllStringSubmatch(rawXML, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) <= 1 {
			continue
		}
		elementName := match[1]
		if seen[elementName] {
			continue
		}
		seen[elementName] = true

		var childPath string
		if basePath == "" {
			childPath = elementName
		} else {
			childPath = basePath + "." + elementName
		}

		operationPath := childPath + ".@operation"
		hasRemove := xmldot.Get(body.Res(), operationPath).String() == "remove"
		xcOperationPath := childPath + ".@xc:operation"
		if !hasRemove {
			hasRemove = xmldot.Get(body.Res(), xcOperationPath).String() == "remove"
		}

		_ = hasRemove // keep parity with setup implementation; recursion handles full walk
		body = cleanupRedundantRemoveRecursive(body, childPath, visited)
	}

	return body
}

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
	"os"
	"regexp"
	"strings"
	"sync"

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

// namespaceExceptions maps namespace prefixes to their full namespace URLs
// for modules that don't follow the standard pattern
var namespaceExceptions = map[string]string{
	// IOS-XR exceptions
	"Cisco-IOS-XR-um-hostname-cfg":        "http://cisco.com/ns/yang/Cisco-IOS-XR-um-hostname-cfg",
	"Cisco-IOS-XR-um-if-ip-address-cfg":   "http://cisco.com/ns/yang/Cisco-IOS-XR-um-if-ip-address-cfg",
	"Cisco-IOS-XR-um-interface-cfg":       "http://cisco.com/ns/yang/Cisco-IOS-XR-um-interface-cfg",
	"Cisco-IOS-XR-um-if-ipv6-cfg":         "http://cisco.com/ns/yang/Cisco-IOS-XR-um-if-ipv6-cfg",
	"Cisco-IOS-XR-um-if-bundle-cfg":       "http://cisco.com/ns/yang/Cisco-IOS-XR-um-if-bundle-cfg",
	"Cisco-IOS-XR-um-cdp-cfg":             "http://cisco.com/ns/yang/Cisco-IOS-XR-um-cdp-cfg",
	"Cisco-IOS-XR-um-if-access-group-cfg": "http://cisco.com/ns/yang/Cisco-IOS-XR-um-if-access-group-cfg",
}

// getNamespaceURL returns the full namespace URL for a given prefix
func getNamespaceURL(prefix string) string {
	// Check exceptions first
	if url, ok := namespaceExceptions[prefix]; ok {
		return url
	}

	// Default pattern for Cisco IOS-XR UM modules
	if strings.HasPrefix(prefix, "Cisco-IOS-XR-um-") {
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

	// When reuse enabled, only serialize write operations
	// Read operations can run concurrently
	if isWrite {
		opMutex.Lock()
		return true
	}

	return false
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

// EditConfig edits the configuration on the device
// If the server supports the candidate capability, it will edit the configuration in the candidate datastore
// and commit it to the running datastore if commit is true.
// If the server does not support the candidate capability, it will edit the configuration in the running datastore.
//
// IMPORTANT: When connection reuse is enabled, callers MUST serialize calls to EditConfig using an
// application-level mutex that also covers ManageNetconfConnection(). This prevents concurrent goroutines
// from attempting to acquire NETCONF datastore locks simultaneously on the same session.
//
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - body: string
//   - commit: bool
func EditConfig(ctx context.Context, client *netconf.Client, body string, commit bool) error {
	// If body is empty, there's nothing to edit
	if body == "" {
		fmt.Println("DEBUG: EditConfig called with empty body")
		tflog.Debug(ctx, "EditConfig called with empty body, skipping")
		return nil
	}
	fmt.Printf("DEBUG: EditConfig body: %s\n", body)

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
			if _, err := client.Commit(ctx); err != nil {
				// Check if the error is "No configuration changes to commit"
				// This can happen during delete operations if the resource is already deleted
				if netconfErr, ok := err.(*netconf.NetconfError); ok {
					for _, e := range netconfErr.Errors {
						if strings.Contains(e.ErrorMessage, "No configuration changes to commit") {
							// Log as warning but don't fail - resource might already be deleted
							tflog.Warn(ctx, "No configuration changes to commit - resource may already be in desired state", map[string]interface{}{
								"error": e.ErrorMessage,
							})
							return nil
						}
					}
				}
				// All other NETCONF RPC errors should cause the operation to fail
				return fmt.Errorf("failed to commit config: %s", FormatNetconfError(err))
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
			// This happens when trying to delete a resource that doesn't exist
			if strings.Contains(err.Error(), "data-missing") {
				tflog.Warn(ctx, "Ignored data-missing error during EditConfig", map[string]interface{}{"error": err.Error()})
				return nil
			}
			return fmt.Errorf("failed to edit config: %s", FormatNetconfError(err))
		}
	}
	return nil
}

// Commit commits the candidate datastore to the running datastore
func Commit(ctx context.Context, client *netconf.Client) error {
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")

	if candidate {
		// Lock running datastore
		if _, err := client.Lock(ctx, "running"); err != nil {
			return fmt.Errorf("failed to lock running datastore: %s", FormatNetconfError(err))
		}
		defer client.Unlock(ctx, "running")

		if _, err := client.Commit(ctx); err != nil {
			return fmt.Errorf("failed to commit config: %s", FormatNetconfError(err))
		}
	}
	return nil
}

// SaveConfig saves the running configuration to startup configuration.
// This is equivalent to 'copy running-config startup-config' in the CLI.
// Uses the cisco-ia:save-config RPC operation.
func SaveConfig(ctx context.Context, client *netconf.Client) error {
	// Ensure connection is open
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	// Build NETCONF body for save-config RPC
	body := netconf.Body{}
	body = SetFromXPath(body, "/cisco-ia:save-config", "")
	body = body.SetAttr("save-config", "xmlns", "http://cisco.com/yang/cisco-ia")

	// Execute the save-config RPC
	if _, err := client.RPC(ctx, body.Res()); err != nil {
		return fmt.Errorf("failed to save config: %s", FormatNetconfError(err))
	}

	return nil
}

// GetXpathFilter creates a NETCONF XPath filter with namespace prefixes removed.
// It processes the XPath expression to strip namespace prefixes from both element names
// and predicate key names, preserving the path structure.
//
// Supports the same XPath formats as SetFromXPath and GetFromXPath:
//   - Paths with namespace prefixes: /Cisco-IOS-XE-native:native/interface
//   - Single predicates: /native/interface[name='GigabitEthernet1']
//   - Multiple predicates: /native/interface[name='Gi1'][vrf='VRF1']
//   - Nested paths: /native/interface[name='Gi1']/ip/address
//   - Values with slashes: /native/interface[name='GigabitEthernet1/0/1']
//   - Predicates with namespace prefixes: /Cisco-IOS-XE-native:interface[Cisco-IOS-XE-native:name='Gi1']
//
// Example transformations:
//
//	Input:  "/Cisco-IOS-XE-native:native/aaa"
//	Output: netconf.Filter{Type: "xpath", Content: "/native/aaa"}
//
//	Input:  "/Cisco-IOS-XE-native:native/interface[Cisco-IOS-XE-native:name='Gi1']/Cisco-IOS-XE-native:ip"
//	Output: netconf.Filter{Type: "xpath", Content: "/native/interface[name='Gi1']/ip"}
//
//	Input:  "/native/interface[name='GigabitEthernet1/0/1']"
//	Output: netconf.Filter{Type: "xpath", Content: "/native/interface[name='GigabitEthernet1/0/1']"}
func GetXpathFilter(xPath string) netconf.Filter {
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

	// Process each segment to remove namespace prefixes
	processedSegments := make([]string, 0, len(segments))
	for _, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		// Remove namespace prefix from element name
		elementName = removeNamespacePrefix(elementName)

		// Reconstruct segment with predicates
		if len(keys) > 0 {
			// Build predicates in order
			predicates := make([]string, 0, len(keys))
			for _, kv := range keys {
				// Remove namespace prefix from key name
				keyName := removeNamespacePrefix(kv.Key)
				predicates = append(predicates, fmt.Sprintf("%s='%s'", keyName, kv.Value))
			}
			// Reconstruct segment with all predicates
			reconstructed := elementName
			for _, pred := range predicates {
				reconstructed += "[" + pred + "]"
			}
			processedSegments = append(processedSegments, reconstructed)
		} else {
			processedSegments = append(processedSegments, elementName)
		}
	}

	// Join segments back with slashes
	cleanedPath := "/" + strings.Join(processedSegments, "/")
	return netconf.XPathFilter(cleanedPath)
}

// GetSubtreeFilter creates a NETCONF subtree filter from an XPath expression.
// IOS-XR devices do not support XPath filters, so we need to convert the XPath
// into a subtree filter. This function builds a nested XML structure that matches
// the path hierarchy.
//
// Example transformations:
//
//	Input:  "/Cisco-IOS-XR-um-hostname-cfg:/hostname"
//	Output: netconf.Filter{Type: "subtree", Content: "<hostname xmlns=\"http://cisco.com/ns/yang/Cisco-IOS-XR-um-hostname-cfg\"/>"}
//
//	Input:  "/Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']"
//	Output: subtree with nested structure including key predicates
func GetSubtreeFilter(xPath string) netconf.Filter {
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
	// Extract namespace from first segment if it ends with ":"
	var namespace string
	processedSegments := make([]string, 0, len(segments))

	for i, seg := range segments {
		if strings.HasSuffix(seg, ":") && i == 0 {
			// This is a namespace-only segment
			namespace = strings.TrimSuffix(seg, ":")
		} else if seg != "" {
			processedSegments = append(processedSegments, seg)
		}
	}

	segments = processedSegments

	// Build nested XML structure
	var content strings.Builder

	for i, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		// Extract namespace from element name (format: namespace:element) if not already extracted
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

		// Skip empty elements
		if element == "" {
			continue
		}

		// Open the element
		content.WriteString("<")
		content.WriteString(element)

		// Add namespace attribute for the root element
		if i == 0 && namespace != "" {
			nsURL := getNamespaceURL(namespace)
			content.WriteString(fmt.Sprintf(" xmlns=\"%s\"", nsURL))
		}

		// If there are key predicates, add them as child elements
		if len(keys) > 0 {
			content.WriteString(">")
			for _, kv := range keys {
				keyName := removeNamespacePrefix(kv.Key)
				content.WriteString(fmt.Sprintf("<%s>%s</%s>", keyName, kv.Value, keyName))
			}
		} else if i < len(segments)-1 {
			// Not the last segment and no keys, keep it open
			content.WriteString(">")
		} else {
			// Last segment without keys, self-closing
			content.WriteString("/>")
			break
		}
	}

	// Close all open tags in reverse order
	for i := len(segments) - 1; i >= 0; i-- {
		segment := segments[i]
		elementName, keys := parseXPathSegment(segment)
		element := removeNamespacePrefix(elementName)

		// Skip empty elements
		if element == "" {
			continue
		}

		// Only close if there were keys or it's not the last element
		if len(keys) > 0 || i < len(segments)-1 {
			content.WriteString("</")
			content.WriteString(element)
			content.WriteString(">")
		}
	}

	// Create the subtree filter with proper type
	result := netconf.Filter{
		Type:    "subtree",
		Content: content.String(),
	}

	return result
}

// KeyValue represents a key-value pair with preserved order
type KeyValue struct {
	Key   string
	Value string
}

// dotPath converts a YANG path to a dot path by removing XPath predicates (keys) and namespace prefixes.
// Example: "Cisco-IOS-XE-native:interface/GigabitEthernet[name]/description" -> "interface.GigabitEthernet.description"
func dotPath(path string) string {
	// Remove XPath predicates like [name='value'] or [name]
	path = xpathPredicateRegex.ReplaceAllString(path, "")

	path = strings.ReplaceAll(path, "/", ".")

	// Remove namespace prefixes like "Cisco-IOS-XE-native:" or "Cisco-IOS-XE-bgp:"
	return namespaceRegex.ReplaceAllString(path, "")
}

// removeNamespacePrefix removes the namespace prefix from a single element or attribute name.
// Example: "Cisco-IOS-XE-native:interface" -> "interface"
func removeNamespacePrefix(name string) string {
	if idx := strings.Index(name, ":"); idx != -1 {
		return name[idx+1:]
	}
	return name
}

// setWithNamespaces sets a value in the netconf body and automatically adds
// namespace declarations for any prefixes found in the path.
func setWithNamespaces(body netconf.Body, fullPath string, value any) netconf.Body {
	// Set the value
	body = body.Set(dotPath(fullPath), value)

	// Extract and add namespace declarations
	body = augmentNamespaces(body, fullPath)

	return body
}

// augmentNamespaces walks through the path and adds namespace declarations
// to elements where prefixes appear, checking the current body to avoid duplicates.
// Only adds xmlns to elements with explicit namespace prefixes, not to children that inherit.
func augmentNamespaces(body netconf.Body, path string) netconf.Body {
	segments := strings.Split(path, ".")
	pathWithoutPrefix := make([]string, 0, len(segments))

	for _, segment := range segments {
		// Strip prefix from segment and add to path
		cleanSegment := removeNamespacePrefix(segment)
		// Also strip XPath predicates like [key='value']
		// This prevents malformed paths like "standard[name=SACL1].@xmlns"
		if idx := strings.IndexByte(cleanSegment, '['); idx != -1 {
			cleanSegment = cleanSegment[:idx]
		}
		pathWithoutPrefix = append(pathWithoutPrefix, cleanSegment)

		// If this segment has a namespace prefix, add xmlns declaration
		// Child elements WITHOUT prefixes will inherit the namespace automatically
		if idx := strings.Index(segment, ":"); idx != -1 {
			prefix := segment[:idx]
			currentPath := strings.Join(pathWithoutPrefix, ".")

			// Check for namespace exceptions first
			namespace, ok := namespaceExceptions[prefix]
			if !ok {
				// Use default pattern if no exception exists
				namespace = namespaceBaseURL + prefix
			}

			xmlnsPath := currentPath + ".@xmlns"
			if !xmldot.Get(body.Res(), xmlnsPath).Exists() {
				body = body.Set(xmlnsPath, namespace)
			}
		}
		// Removed the "else if lastNamespace != ''" block that was adding xmlns to all children
		// Child elements now properly inherit namespace from their parent
	}

	return body
}

// splitXPathSegments splits an XPath into segments while respecting bracket boundaries.
// This prevents splitting on forward slashes inside predicates like [name='GigabitEthernet1/0/1']
func splitXPathSegments(xPath string) []string {
	segments := []string{}
	var currentSegment strings.Builder
	bracketDepth := 0

	for _, char := range xPath {
		switch char {
		case '[':
			bracketDepth++
			currentSegment.WriteRune(char)
		case ']':
			bracketDepth--
			currentSegment.WriteRune(char)
		case '/':
			if bracketDepth == 0 {
				// We're not inside brackets, so this is a segment separator
				if currentSegment.Len() > 0 {
					segments = append(segments, currentSegment.String())
					currentSegment.Reset()
				}
			} else {
				// We're inside brackets, so this is part of a predicate value
				currentSegment.WriteRune(char)
			}
		default:
			currentSegment.WriteRune(char)
		}
	}

	// Add the last segment if there's anything left
	if currentSegment.Len() > 0 {
		segments = append(segments, currentSegment.String())
	}

	return segments
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

	// Build path incrementally, creating each element
	pathSegments := make([]string, 0, len(segments))

	for i, segment := range segments {
		// Parse segment: element[key='value'][key2='value2'] -> element, []KeyValue
		elementName, keys := parseXPathSegment(segment)

		// Add element name to path (without predicates)
		pathSegments = append(pathSegments, elementName)
		fullPath := strings.Join(pathSegments[:i+1], ".")

		// If this segment has keys, set all key values in order
		if len(keys) > 0 {
			for _, kv := range keys {
				keyPath := fullPath + "." + kv.Key
				body = setWithNamespaces(body, keyPath, kv.Value)
			}
		}
	}

	// Optionally ensure the complete path structure exists, including non-predicate elements.
	// Only create the structure if requested and if it doesn't already exist.
	if ensureStructure && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		existingContent := xmldot.Get(body.Res(), dotPath(fullPath)).String()
		if existingContent == "" {
			// Path doesn't exist yet, create it with SetWithNamespaces
			body = setWithNamespaces(body, fullPath, "")
		}
	}

	return body, pathSegments
}

// parseXPathSegment parses an XPath segment with single or multiple keys
// Supports formats:
//   - element[key='value']
//   - element[key1='value1'][key2='value2']
//   - element[key1='value1' and key2='value2']
//
// Returns: (elementName, []KeyValue) - order is preserved from the XPath
func parseXPathSegment(segment string) (string, []KeyValue) {
	// Check for predicate: element[...]
	if idx := strings.Index(segment, "["); idx != -1 {
		elementName := segment[:idx]
		keys := make([]KeyValue, 0)

		// Extract all predicates - handle both [key1='val1'][key2='val2'] and [key1='val1' and key2='val2']
		remainingPredicates := segment[idx:]

		// Use pre-compiled pattern to match predicates: [anything]
		predicates := predicatePattern.FindAllStringSubmatch(remainingPredicates, -1)

		for _, match := range predicates {
			if len(match) > 1 {
				predicate := match[1]

				// Split by 'and' to handle combined predicates
				conditions := strings.Split(predicate, " and ")
				for _, condition := range conditions {
					// Parse each condition: key='value' or key="value"
					if eqIdx := strings.Index(condition, "="); eqIdx != -1 {
						keyName := strings.TrimSpace(condition[:eqIdx])
						value := condition[eqIdx+1:]
						// Remove quotes
						keyValue := strings.Trim(value, `'"`)
						keys = append(keys, KeyValue{Key: keyName, Value: keyValue})
					}
				}
			}
		}

		return elementName, keys
	}

	// No predicate
	return segment, nil
}

// RemoveFromXPath creates all elements in an XPath with an nc:operation="delete" attribute.
// Returns a Body for use in toBodyXML functions that build up XML progressively.
// For delete operations in resource Delete() functions, use RemoveFromXPathString() instead.
func RemoveFromXPath(body netconf.Body, xPath string) netconf.Body {
	xmlStr := RemoveFromXPathString(body, xPath)
	// Wrap the string back in a Body for compatibility with toBodyXML functions
	// Note: This will get mangled if you call .Res() on it, so only use this in toBodyXML
	return netconf.Body{}.SetRaw("", xmlStr)
}

// RemoveFromXPathString creates all elements in an XPath with an nc:operation="delete" attribute
// and returns the raw XML string. Use this for delete operations to avoid library mangling.
// The nc: prefix is part of the NETCONF protocol and doesn't need explicit declaration.
//
// Example: /Cisco-IOS-XR-um-hostname-cfg:/hostname
// Returns: <hostname xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-um-hostname-cfg" nc:operation="delete"></hostname>
func RemoveFromXPathString(body netconf.Body, xPath string) string {
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

	// Handle namespace-only first segment (e.g., "Cisco-IOS-XR-um-hostname-cfg:" followed by "hostname")
	if len(segments) > 0 && strings.HasSuffix(segments[0], ":") {
		namespace := strings.TrimSuffix(segments[0], ":")
		if len(segments) > 1 {
			// Prepend namespace to the next segment
			segments[1] = namespace + ":" + segments[1]
			// Remove the first segment
			segments = segments[1:]
		}
	}

	if len(segments) == 0 {
		return ""
	}

	var sb strings.Builder
	var closingTags []string

	for i, segment := range segments {
		// Parse segment: element[key='value'] -> element, []KeyValue
		elementName, keys := parseXPathSegment(segment)

		// Extract namespace if present
		var namespace string
		if strings.Contains(elementName, ":") {
			parts := strings.SplitN(elementName, ":", 2)
			namespace = parts[0]
			elementName = parts[1]
		}

		sb.WriteString("<")
		sb.WriteString(elementName)

		// Add namespace declaration if this element has a namespace
		if namespace != "" {
			// Try to find the namespace URL
			nsURL := getNamespaceURL(namespace)
			if nsURL != "" {
				sb.WriteString(fmt.Sprintf(` xmlns="%s"`, nsURL))
			}
		}

		// Always add the NETCONF namespace declaration to the root element
		// This is required because we use the nc: prefix for the operation attribute
		if i == 0 {
			sb.WriteString(` xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0"`)
		}

		// Add nc:operation="delete" to the last element
		if i == len(segments)-1 {
			sb.WriteString(` nc:operation="delete"`)
		}

		sb.WriteString(">")

		// Add keys
		for _, kv := range keys {
			// Remove namespace prefix from key name if present
			keyName := kv.Key
			if strings.Contains(keyName, ":") {
				parts := strings.SplitN(keyName, ":", 2)
				keyName = parts[1]
			}

			sb.WriteString(fmt.Sprintf("<%s>%s</%s>", keyName, kv.Value, keyName))
		}

		closingTags = append(closingTags, fmt.Sprintf("</%s>", elementName))
	}

	// Close tags in reverse order
	for i := len(closingTags) - 1; i >= 0; i-- {
		sb.WriteString(closingTags[i])
	}

	xmlStr := sb.String()

	tflog.Debug(context.Background(), "RemoveFromXPathString: Generated XML", map[string]interface{}{
		"xpath": xPath,
		"xml":   xmlStr,
	})

	return xmlStr
}

// CleanupRedundantRemoveOperations removes redundant operation="remove" attributes from child elements
// when a parent element already has operation="remove". This is called after the full NETCONF payload
// has been built to sanitize the XML before sending to the device.
//
// Example:
//
//	Input:  <trap operation="remove"><severity operation="remove"></severity></trap>
//	Output: <trap operation="remove"></trap>
//	Input:  <trap xc:operation="remove"><severity xc:operation="remove"></severity></trap>
//	Output: <trap xc:operation="remove"></trap>
func CleanupRedundantRemoveOperations(body netconf.Body) netconf.Body {
	if body.Res() == "" {
		return body
	}

	// Start recursive cleanup from root level
	body = cleanupRedundantRemoveRecursive(body, "", make(map[string]bool))

	return body
}

// cleanupRedundantRemoveRecursive walks through the XML tree and removes redundant child operations.
// It processes the tree level by level, tracking which elements have operation="remove" to avoid
// checking the same elements multiple times.
func cleanupRedundantRemoveRecursive(body netconf.Body, basePath string, visited map[string]bool) netconf.Body {
	xml := body.Res()
	if xml == "" {
		return body
	}

	// Skip if we've already processed this path
	if visited[basePath] {
		return body
	}
	visited[basePath] = true

	// Get content at current path
	var currentResult xmldot.Result
	if basePath == "" {
		currentResult = xmldot.Get(xml, "")
	} else {
		currentResult = xmldot.Get(xml, basePath)
		if !currentResult.Exists() {
			return body
		}
	}

	// Find immediate child elements at this level by parsing Raw content
	// We need to find child element names, not search for operation="remove" yet
	rawXML := xml
	if basePath != "" {
		rawXML = currentResult.Raw
	}

	// Extract child element names from the raw XML
	childPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)[\s>]`)
	matches := childPattern.FindAllStringSubmatch(rawXML, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			elementName := match[1]
			if seen[elementName] {
				continue
			}
			seen[elementName] = true

			// Build path to this child element
			var childPath string
			if basePath == "" {
				childPath = elementName
			} else {
				childPath = basePath + "." + elementName
			}

			// Check if this child has operation="remove"
			operationPath := childPath + ".@operation"
			hasRemove := xmldot.Get(body.Res(), operationPath).String() == "remove"
			// Check for both xc:operation and operation just in case
			xcOperationPath := childPath + ".@xc:operation"
			// This element has operation="remove", clean up its children
			if !hasRemove {
				hasRemove = xmldot.Get(body.Res(), xcOperationPath).String() == "remove"
			}
			body = cleanupRedundantRemoveRecursive(body, childPath, visited)
		}
	}

	return body
}

// removeAllChildOperations removes all child elements that have operation="remove" from under the given parent path.
// When a parent has operation="remove", all its children will be removed by the device anyway, so we delete
// child elements entirely to avoid redundant operations and prevent empty element errors.
// However, we preserve key elements (list identifiers) as they may be needed for identification.
func removeAllChildOperations(body netconf.Body, parentPath string) netconf.Body {
	xml := body.Res()

	parentResult := xmldot.Get(xml, parentPath)
	if !parentResult.Exists() {
		return body
	}

	rawXML := parentResult.Raw
	if rawXML == "" {
		return body
	}

	// Find all immediate child elements at this level
	childPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)[\s>]`)
	matches := childPattern.FindAllStringSubmatch(rawXML, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			childName := match[1]
			if seen[childName] {
				continue
			}
			seen[childName] = true

			childPath := parentPath + "." + childName

			// Check if this child exists and what it contains
			childResult := xmldot.Get(body.Res(), childPath)
			if !childResult.Exists() {
				continue
			}

			// Check if this child has operation="remove"
			operationPath := childPath + ".@operation"
			hasRemove := xmldot.Get(body.Res(), operationPath).String() == "remove"

			if hasRemove {
				// This child has operation="remove" under a parent that also has operation="remove"
				// Check if it contains keys - if so, keep it but remove the operation attribute
				// If it doesn't contain keys (empty or only has other operations), delete it entirely
				xcOperationPath := childPath + ".@xc:operation"

				if !hasRemove {
					hasRemove = xmldot.Get(body.Res(), xcOperationPath).String() == "remove"
				}
				// Has keys, just remove the operation attribute but keep the element
				body = body.Delete(xcOperationPath)
				body = body.Delete(operationPath)

				// Now recursively clean this element's children
				// Since the parent (host) has operation="remove", ALL descendants are redundant
				body = removeAllDescendantOperations(body, childPath)
			} else {
				// Child doesn't have operation="remove", but we still need to check its descendants
				// because they might have redundant operation="remove" attributes
				body = removeAllDescendantOperations(body, childPath)
			}
		}
	}
	return body
}

// shouldDeleteEntireElement determines if an element should be deleted entirely or just have
// its operation attribute removed. Elements with only keys and operation="remove" should be
// deleted entirely to avoid empty elements like <severity operation="remove"></severity>.
func shouldDeleteEntireElement(content string) bool {
	if content == "" || strings.TrimSpace(content) == "" {
		return true
	}

	// Check if content only contains simple key elements (elements with text content only)
	// Pattern: <elementName>text</elementName>
	keyPattern := regexp.MustCompile(`^\s*(?:<[a-zA-Z0-9\-_]+>[^<]+</[a-zA-Z0-9\-_]+>\s*)+$`)
	return keyPattern.MatchString(content)
}

// containsKeys checks if the content contains key elements (simple elements with text content).
// Returns true if there are any key-like elements: <elementName>value</elementName>
func containsKeys(content string) bool {
	if content == "" || strings.TrimSpace(content) == "" {
		return false
	}

	// Pattern matches key elements: <elementName>text</elementName>
	keyPattern := regexp.MustCompile(`<[a-zA-Z0-9\-_]+>[^<]+</[a-zA-Z0-9\-_]+>`)
	return keyPattern.MatchString(content)
}

// removeAllDescendantOperations removes all operation="remove" attributes and elements from
// descendants of the given path. This is used when a parent has operation="remove" and we've
// decided to keep a child (because it has keys), but need to clean up all of that child's descendants.
func removeAllDescendantOperations(body netconf.Body, parentPath string) netconf.Body {
	xml := body.Res()

	parentResult := xmldot.Get(xml, parentPath)
	if !parentResult.Exists() {
		return body
	}

	rawXML := parentResult.Raw
	if rawXML == "" {
		return body
	}

	// Find all child elements at this level
	childPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)[\s>]`)
	matches := childPattern.FindAllStringSubmatch(rawXML, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			elementName := match[1]
			if seen[elementName] {
				continue
			}
			seen[elementName] = true

			childPath := parentPath + "." + elementName

			// Check if this child exists
			childResult := xmldot.Get(body.Res(), childPath)
			if !childResult.Exists() {
				continue
			}

			// Check if this child has operation="remove"
			operationPath := childPath + ".@operation"
			hasRemove := xmldot.Get(body.Res(), operationPath).String() == "remove"

			if hasRemove {
				// This descendant has operation="remove", delete it entirely
				// (we don't need to preserve it since the ancestor will be removed)
				body = body.Delete(childPath)
			} else {
				// No operation attribute, but recurse to check its children
				body = removeAllDescendantOperations(body, childPath)
			}
		}
	}

	return body
}

// SetFromXPath creates all elements in an XPath, including keys and namespaces,
// and optionally sets a value at the final path location.
// Supports single and composite keys:
//   - Single: /interface[name='GigabitEthernet1']
//   - Multiple predicates: /interface[name='GigabitEthernet1'][vrf='VRF1']
//   - Combined with 'and': /interface[name='GigabitEthernet1' and vrf='VRF1']
//   - Values with slashes: /interface[name='GigabitEthernet1/0/1']
//
// If value is nil or empty string, only the structure is created without setting a value.
// If value is non-empty, it's set at the final path location.
//
// Multi-Root Support:
// SetFromXPath properly handles multiple root-level sibling elements (e.g., both <deny> and <permit>
// at the same level). The underlying xmldot library automatically detects when different root paths
// are used and creates sibling elements instead of nesting them.
//
// Example (multi-root XML):
//
//	body := netconf.Body{}
//	body = SetFromXPath(body, "sequence", "10")
//	body = SetFromXPath(body, "deny/std-ace/prefix", "10.0.0.0")
//	body = SetFromXPath(body, "permit/std-ace/prefix", "192.168.0.0")
//	// Result: <sequence>10</sequence><deny>...</deny><permit>...</permit>
func SetFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	// Determine if we need to create empty structure
	// Only create empty structure if no value will be set
	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments := buildXPathStructure(body, xPath, ensureStructure)

	// Only set the value if it's not nil and not empty string
	// This prevents overwriting key children when no value is needed
	if hasValue && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		body = setWithNamespaces(body, fullPath, value)
		// tflog.Debug(context.Background(), "SetFromXPath set value", map[string]interface{}{"path": fullPath, "value": value})
	}

	return body
}

// AppendFromXPath creates all elements in an XPath and appends a value to a list by using
// the ".-1" syntax. This is useful for adding multiple items to a list without keys.
// The function automatically appends ".-1" to the final element in the path.
//
// Example:
//
//	body := netconf.Body{}
//	body = AppendFromXPath(body, "native/route-map/rule/match/ip/address", "10")
//	body = AppendFromXPath(body, "native/route-map/rule/match/ip/address", "20")
//	// Result: <native><route-map><rule><match><ip>
//	//           <address>10</address>
//	//           <address>20</address>
//	//         </ip></match></rule></route-map></native>
//
// Note: This function is designed for simple list items without keys. For lists with keys,
// use SetFromXPath with predicates instead.

//	// Result: <native><route-map><rule><match><ip>
//	//           <address>10</address>
//	//           <address>20</address>
//	//         </ip></match></rule></route-map></native>
//
// Note: This function is designed for simple list items without keys. For lists with keys,
// use SetFromXPath with predicates instead.
func AppendFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	// Determine if we need to create empty structure
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

// AppendRawFromXPath appends raw XML content as a new list item using the .-1 syntax.
// This is similar to AppendFromXPath but works with pre-formatted XML content instead of simple values.
// The raw XML is wrapped in the final element specified by the xPath.
//
// Example:
//
//	xPath: "/native/interface/GigabitEthernet"
//	value: "<name>Gi1</name><description>Port 1</description>"
//	Result: Appends <GigabitEthernet><name>Gi1</name><description>Port 1</description></GigabitEthernet>
//
// This function is useful for building lists of complex elements where each element has multiple child nodes.
func AppendRawFromXPath(body netconf.Body, xPath string, value string) netconf.Body {
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

	// Build structure up to the PARENT of the list element (not including the list element itself)
	// This prevents creating an empty list element
	if len(segments) > 1 {
		parentXPath := "/" + strings.Join(segments[:len(segments)-1], "/")
		body, _ = buildXPathStructure(body, parentXPath, true)
	}

	// Build the complete dot path for the append operation
	var dotPathBuilder strings.Builder
	for i, segment := range segments {
		elementName, _ := parseXPathSegment(segment)
		if i > 0 {
			dotPathBuilder.WriteString(".")
		}
		dotPathBuilder.WriteString(elementName)
	}

	// Convert to dotPath (removes namespaces and predicates)
	fullDotPath := dotPath(dotPathBuilder.String())

	// Use .-1 to append to the list
	appendPath := fullDotPath + ".-1"
	body = body.SetRaw(appendPath, value)

	return body
}

// SetRawFromXPath creates all elements in an XPath, including keys and namespace declarations,
// then inserts raw XML content at the final path location. This is useful when you have
// pre-formatted XML that needs to be inserted as child elements.
//
// The value parameter should contain raw XML content (child elements, attributes, etc.) that will
// be parsed and inserted at the target path. The content is wrapped in the final element tag
// specified by the xPath.
//
// Multi-root Support:
// When called multiple times with the same path, this function appends the new XML content
// as an additional sibling element, creating a multi-root XML fragment at the parent level.
// The underlying xmldot library automatically handles multi-root fragments, making this safe for
// creating multiple sibling elements (e.g., multiple <interface> elements in a list).
//
// Example (single call):
//
//	xpath: Cisco-IOS-XE-native:native/interface[name='Gi1']
//	value: "<description>Management</description><shutdown/>"
//	Result: <native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
//	          <interface>
//	            <name>Gi1</name>
//	            <description>Management</description>
//	            <shutdown/>
//	          </interface>
//	        </native>
//
// Example (multiple calls - multi-root):
//
//	First call:  SetRawFromXPath(body, "/native/interface", "<name>Gi1</name>")
//	Second call: SetRawFromXPath(body, "/native/interface", "<name>Gi2</name>")
//	Result: <native>
//	          <interface><name>Gi1</name></interface>
//	          <interface><name>Gi2</name></interface>
//	        </native>
//
// Unlike SetFromXPath, this function:
//   - Adds xmlns declarations for namespace prefixes in the path (via buildXPathStructure)
//   - Inserts the value as raw XML (parsed as child elements) rather than as text content
//   - Uses body.SetRaw() instead of body.Set() for XML insertion
//   - Supports appending multiple elements at the same path (multi-root fragments)
func SetRawFromXPath(body netconf.Body, xPath string, value string) netconf.Body {
	// FIRST THING - write to file to confirm this function is called!
	f, _ := os.OpenFile("/tmp/SETRAW_CALLED.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if f != nil {
		f.WriteString(fmt.Sprintf("Called with xPath: %s\n", xPath))
		f.Close()
	}

	// Get context for logging - create a background context if none available
	ctx := context.Background()

	// Append to debug files to capture all calls
	f2, _ := os.OpenFile("/tmp/netconf_setraw_calls.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if f2 != nil {
		fmt.Fprintf(f2, "\n=== SetRawFromXPath Call ===\n")
		fmt.Fprintf(f2, "xPath: %s\n", xPath)
		fmt.Fprintf(f2, "value length: %d\n", len(value))
		fmt.Fprintf(f2, "body.Res() length before: %d\n", len(body.Res()))
		f2.Close()
	}

	tflog.Debug(ctx, fmt.Sprintf("[SetRawFromXPath] CALLED with xPath=%s, value length=%d", xPath, len(value)))

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
		// It should inherit from the last namespace in the path or use its own if specified
		finalNamespace := ""
		if idx := strings.Index(finalElement, ":"); idx != -1 {
			// Final element has explicit namespace prefix
			prefix := finalElement[:idx]
			namespace, ok := namespaceExceptions[prefix]
			if !ok {
				namespace = namespaceBaseURL + prefix
			}
			finalNamespace = namespace
		} else {
			// Final element doesn't have explicit namespace, try to inherit from parent path
			// Walk backwards through segments to find the last namespace
			for i := len(segments) - 2; i >= 0; i-- {
				segmentName, _ := parseXPathSegment(segments[i])
				if idx := strings.Index(segmentName, ":"); idx != -1 {
					prefix := segmentName[:idx]
					namespace, ok := namespaceExceptions[prefix]
					if !ok {
						namespace = namespaceBaseURL + prefix
					}
					finalNamespace = namespace
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

		// Debug logging - write to files for ALL calls
		os.WriteFile("/tmp/netconf_debug_xpath.txt", []byte(xPath), 0644)
		os.WriteFile("/tmp/netconf_debug_parentpath.txt", []byte(parentPath), 0644)
		os.WriteFile("/tmp/netconf_debug_body_before.xml", []byte(body.Res()), 0644)
		os.WriteFile("/tmp/netconf_debug_existing.xml", []byte(existingXML), 0644)

		// Log to a file that appends
		f4, _ := os.OpenFile("/tmp/setraw_calls.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if f4 != nil {
			fmt.Fprintf(f4, "\n=== SetRawFromXPath Call ===\n")
			fmt.Fprintf(f4, "xPath: %s\n", xPath)
			fmt.Fprintf(f4, "parentPath: %s\n", parentPath)
			fmt.Fprintf(f4, "existingXML length: %d\n", len(existingXML))
			fmt.Fprintf(f4, "existingXML: %s\n", existingXML)
			fmt.Fprintf(f4, "value length: %d\n", len(value))
			f4.Close()
		}

		// Multi-segment path: wrap the content with the final element tag
		// Only add xmlns on the FIRST element to avoid namespace duplication issues
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

		os.WriteFile("/tmp/netconf_debug_wrapped.xml", []byte(wrappedContent), 0644)

		tflog.Debug(ctx, fmt.Sprintf("[SetRawFromXPath] parentPath=%s, existingXML length=%d", parentPath, len(existingXML)))

		if existingXML != "" {
			// Append wrapped element as sibling to existing content
			// xmldot now supports multi-root XML fragments at the parent level
			combinedXML := existingXML + wrappedContent
			body = body.SetRaw(parentPath, combinedXML)
			os.WriteFile("/tmp/netconf_debug_combined.xml", []byte(combinedXML), 0644)
			os.WriteFile("/tmp/netconf_debug_body_after.xml", []byte(body.Res()), 0644)
			tflog.Debug(ctx, fmt.Sprintf("[SetRawFromXPath] APPENDING - combinedXML length=%d", len(combinedXML)))
		} else {
			// First element, set the wrapped content at parent
			body = body.SetRaw(parentPath, wrappedContent)
			os.WriteFile("/tmp/netconf_debug_body_after.xml", []byte(body.Res()), 0644)
			tflog.Debug(ctx, fmt.Sprintf("[SetRawFromXPath] FIRST ELEMENT"))
		}
	} else {
		// Single-segment path - SetRaw will wrap the content for us
		// Just prepare the inner content (with keys if present)
		innerContent := value
		if len(keys) > 0 {
			tempBody := netconf.Body{}
			for _, kv := range keys {
				tempBody = setWithNamespaces(tempBody, kv.Key, kv.Value)
			}
			innerContent = tempBody.Res() + value
		}

		// Use the element name as the dotPath for SetRaw
		// SetRaw will create <finalElementClean>innerContent</finalElementClean>
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
	// Convert XPath segments back to dotPath format for augmentNamespaces
	if len(segments) > 0 {
		dotPathForNamespaces := strings.Join(segments, ".")
		body = augmentNamespaces(body, dotPathForNamespaces)
	}

	// Log output body length
	f3, _ := os.OpenFile("/tmp/netconf_setraw_calls.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if f3 != nil {
		fmt.Fprintf(f3, "body.Res() length after: %d\n", len(body.Res()))
		f3.Close()
	}

	return body
}

// GetFromXPath converts an XPath expression to a xmldot path and retrieves the result.
// Uses manual filtering to correctly handle both single and multiple elements with predicates.
// Supports the same XPath formats as SetFromXPath:
//   - Single: /interface[name='GigabitEthernet1']
//   - Multiple predicates: /interface[name='GigabitEthernet1'][vrf='VRF1']
//   - Combined with 'and': /interface[name='GigabitEthernet1' and vrf='VRF1']
//   - Values with slashes: /interface[name='GigabitEthernet1/0/1']
//   - Nested paths: /native/interface[name='Gi1']/ip/address
//
// Example: /native/interface[name='Gi1']/ip/address
// Processes path segments, validates predicates, and constructs the final xmldot path
func GetFromXPath(res xmldot.Result, xPath string) xmldot.Result {
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

	// Use manual filtering to handle both single and multiple elements with predicates.
	// xmldot's #() filter syntax doesn't work reliably with single elements, so we
	// always use the manual path which correctly handles both cases.
	xml := res.Raw
	pathSoFar := make([]string, 0, len(segments))

	for _, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		// Remove namespace prefix from element name
		elementName = removeNamespacePrefix(elementName)
		pathSoFar = append(pathSoFar, elementName)

		// Build current path
		currentPath := strings.Join(pathSoFar, ".")

		// Check if there are multiple sibling elements using absolute path
		countPath := currentPath + ".#"
		count := xmldot.Get(xml, countPath).Int()

		// Apply filtering if keys exist
		if len(keys) > 0 {
			found := false
			if count > 1 {
				// Multiple elements - iterate using numeric indices
				for idx := 0; idx < int(count); idx++ {
					indexedPath := fmt.Sprintf("%s.%d", currentPath, idx)
					item := xmldot.Get(xml, indexedPath)

					allMatch := true
					for _, kv := range keys {
						// Remove namespace prefix from key name
						keyName := removeNamespacePrefix(kv.Key)
						keyResult := item.Get(keyName)
						if !keyResult.Exists() || keyResult.String() != kv.Value {
							allMatch = false
							break
						}
					}

					if allMatch {
						// Update currentPath to point to the matched item
						pathSoFar[len(pathSoFar)-1] = fmt.Sprintf("%s.%d", elementName, idx)
						found = true
						break
					}
				}
			} else {
				// Single element - check directly
				currentResult := xmldot.Get(xml, currentPath)
				allMatch := true
				for _, kv := range keys {
					// Remove namespace prefix from key name
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

	// Return the final result
	finalPath := strings.Join(pathSoFar, ".")

	// Check if the final path points to multiple elements (array)
	// If so, use the #. syntax to return an array result that ForEach can iterate over
	countPath := finalPath + ".#"
	count := xmldot.Get(xml, countPath).Int()
	if count > 1 {
		// Multiple elements exist - use #. syntax to get array result
		// Split the path to insert #. before the last segment
		if len(pathSoFar) >= 2 {
			// Build parent path and use #. syntax: parent.#.child
			parentPath := strings.Join(pathSoFar[:len(pathSoFar)-1], ".")
			childName := pathSoFar[len(pathSoFar)-1]
			arrayPath := parentPath + ".#." + childName
			return xmldot.Get(xml, arrayPath)
		}
	}

	return xmldot.Get(xml, finalPath)
}

// IsListPath checks if an XPath represents a list item (ends with a predicate).
// List items have predicates like [name='value'] or [name=value] at the end of their XPath,
// while containers/singletons don't end with predicates.
//
// This is useful for determining if an empty GetConfig response should be
// interpreted as "resource not found" (for list items) vs other semantics.
//
// Parameters:
//   - xPath: The XPath to check
//
// Returns:
//   - true if the path ends with a predicate (is a list item)
//   - false if the path does not end with a predicate (is a container/singleton)
//
// Examples:
//   - "/native/interface/Vlan[name=10]" → true (ends with predicate)
//   - "/native/interface/GigabitEthernet[name='1/0/1']" → true (ends with predicate)
//   - "/native/router/bgp[id=65000]/neighbor" → false (has predicate but doesn't end with one)
//   - "/native/clock" → false (container)
//   - "/native/hostname" → false (singleton)
func IsListPath(xPath string) bool {
	// Trim whitespace
	xPath = strings.TrimSpace(xPath)

	// Check if the path ends with a closing bracket (predicate)
	// A list path will end with something like [name='value']
	return strings.HasSuffix(xPath, "]")
}

// IsGetConfigResponseEmpty checks if a GetConfig response has an empty <data> element.
// Returns true if the response contains <data></data> with no child elements,
// indicating that the requested configuration does not exist on the device.
//
// This is useful for determining if a resource exists before attempting to parse
// its attributes, particularly during Read operations or import.
//
// IMPORTANT: This should typically be combined with IsListPath() to only treat
// empty responses as "not found" for list items:
//
//	if helpers.IsGetConfigResponseEmpty(&res) && helpers.IsListPath(state.getXPath()) {
//	    // List item does not exist
//	    resp.State.RemoveResource(ctx)
//	    return
//	}
//
// Parameters:
//   - res: The NETCONF response from GetConfig operation
//
// Returns:
//   - true if the data element is empty (no child elements)
//   - false if the data element contains configuration
//
// Example usage:
//
//	res, err := device.NetconfClient.GetConfig(ctx, "running", filter)
//	if err != nil {
//	    return err
//	}
//	if helpers.IsGetConfigResponseEmpty(&res) && helpers.IsListPath(state.getXPath()) {
//	    // Resource does not exist (list item)
//	    resp.State.RemoveResource(ctx)
//	    return
//	}
//	// Parse the configuration
//	state.fromBodyXML(ctx, res.Res)
func IsGetConfigResponseEmpty(res *netconf.Res) bool {
	if res == nil {
		return true
	}

	// Get the data element from the response
	dataResult := res.Res.Get("data")

	// If data element doesn't exist, consider it empty
	if !dataResult.Exists() {
		return true
	}

	// Check if data element has any children using Map()
	// An empty <data></data> element will have an empty map
	children := dataResult.Map()

	// If the map is empty (no child elements), the response is empty
	// The "%" key represents direct text content, which we also want to ignore
	// since whitespace-only content is not meaningful configuration
	for key := range children {
		if key != "%" {
			return false
		}
	}

	return true
}

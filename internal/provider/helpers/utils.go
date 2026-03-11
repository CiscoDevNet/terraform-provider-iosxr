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
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/xmldot"
	"github.com/tidwall/gjson"
)

// ============================================================================
// Generic Utilities
// ============================================================================

// Contains checks if a string slice contains a specific string
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// RemoveEmptyStrings removes empty strings from a slice
func RemoveEmptyStrings(s []string) []string {
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}

// Must panics if error is not nil, otherwise returns the value
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Add adds two integers
func Add(a, b int) int {
	return a + b
}

// IsLast checks if index is last element of list
func IsLast(index int, len int) bool {
	return index+1 == len
}

// ============================================================================
// Terraform Type Conversion Utilities (JSON/gjson)
// ============================================================================

// GetValueSlice converts a slice of gjson.Result to a slice of Terraform attr.Value.
func GetValueSlice(result []gjson.Result) []attr.Value {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return v
}

// GetStringList converts gjson results to Terraform string list
func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

// GetInt64List converts gjson results to Terraform int64 list
func GetInt64List(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

// GetStringSlice converts gjson results to Terraform string slice
func GetStringSlice(result []gjson.Result) []types.String {
	v := make([]types.String, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return v
}

// GetStringSet converts gjson results to Terraform string set
func GetStringSet(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

// GetInt64Set converts gjson results to Terraform int64 set
func GetInt64Set(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

// ============================================================================
// Terraform Type Conversion Utilities (XML/xmldot)
// ============================================================================

// GetStringListXML converts xmldot results to Terraform string list
func GetStringListXML(result []xmldot.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

// GetInt64ListXML converts xmldot results to Terraform int64 list
func GetInt64ListXML(result []xmldot.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

// GetStringSetXML converts xmldot results to Terraform string set
func GetStringSetXML(result []xmldot.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

// GetInt64SetXML converts xmldot results to Terraform int64 set
func GetInt64SetXML(result []xmldot.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

// GetAllChildElements returns all child elements with the given name from a parent xmldot.Result.
// This is needed when multiple sibling elements have the same name in XML.
func GetAllChildElements(parent xmldot.Result, elementName string) []xmldot.Result {
	// Check if there are multiple elements
	countPath := elementName + ".#"
	count := parent.Get(countPath).Int()

	if count == 0 {
		// No elements found
		return []xmldot.Result{}
	}

	if count == 1 {
		// Single element - return it directly
		elem := parent.Get(elementName)
		if elem.Exists() {
			return []xmldot.Result{elem}
		}
		return []xmldot.Result{}
	}

	// Multiple elements - iterate using numeric indices
	results := make([]xmldot.Result, 0, count)
	for i := 0; i < int(count); i++ {
		indexPath := fmt.Sprintf("%s.%d", elementName, i)
		elem := parent.Get(indexPath)
		if elem.Exists() {
			results = append(results, elem)
		}
	}
	return results
}

// ============================================================================
// String Manipulation and Naming Conversion Utilities
// ============================================================================

// ToYangShortName gets short YANG name without prefix (xxx:abc -> abc)
func ToYangShortName(s string) string {
	elements := strings.Split(s, "/")
	for i := range elements {
		if strings.Contains(elements[i], ":") {
			elements[i] = strings.Split(elements[i], ":")[1]
		}
	}
	return strings.Join(elements, "/")
}

// ToGoName converts TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		if strings.Contains(value, ":") {
			value = strings.Split(value, ":")[1]
		}
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// CamelCase converts string to camel case
func CamelCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// SnakeCase converts string to snake case
func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// RemoveNamespacePrefix removes the namespace prefix from an element name
// Example: "Cisco-IOS-XR-um-logging-cfg:suppress" -> "suppress"
func RemoveNamespacePrefix(name string) string {
	if idx := strings.Index(name, ":"); idx != -1 {
		return name[idx+1:]
	}
	return name
}

// GetNamespacePrefixFromSegment extracts the namespace prefix from a segment name
// Example: "Cisco-IOS-XR-um-logging-correlator-cfg:suppress" -> "Cisco-IOS-XR-um-logging-correlator-cfg"
func GetNamespacePrefixFromSegment(elementName string) string {
	if idx := strings.Index(elementName, ":"); idx != -1 {
		return elementName[:idx]
	}
	return ""
}

// HasNamespacePrefix checks if a yang path contains any namespace prefix (e.g. "prefix:")
func HasNamespacePrefix(yangName string) bool {
	return strings.Contains(yangName, ":")
}

// TopElementName returns the first path element name stripped of any namespace
// prefix. e.g. "Cisco-IOS-XR-um-logging-correlator-cfg:suppress/rules/rule"
// → "suppress", and "suppress/duplicates" → "suppress".
func TopElementName(yangName string) string {
	first := yangName
	if idx := strings.Index(yangName, "/"); idx != -1 {
		first = yangName[:idx]
	}
	if idx := strings.Index(first, ":"); idx != -1 {
		return first[idx+1:]
	}
	return first
}

// XmlNamespacePrefixFromXPath extracts the YANG module prefix from the first
// segment of an XPath that carries one, e.g.
//
//	"Cisco-IOS-XR-um-logging-cfg:suppress/duplicates" -> "Cisco-IOS-XR-um-logging-cfg"
//	"suppress/duplicates"                              -> ""
func XmlNamespacePrefixFromXPath(xp string) string {
	first := xp
	if idx := strings.Index(xp, "/"); idx != -1 {
		first = xp[:idx]
	}
	if idx := strings.Index(first, ":"); idx != -1 {
		return first[:idx]
	}
	return ""
}

// ExtractNamespacePrefix extracts namespace prefix from segment if present
// Example: "Cisco-IOS-XR:interface" -> ("Cisco-IOS-XR", true)
func ExtractNamespacePrefix(segment string) (string, bool) {
	if idx := strings.Index(segment, ":"); idx != -1 {
		return segment[:idx], true
	}
	return "", false
}

// ExtractNamespaceFromXPath extracts namespace prefix from XPath
// Example: "Cisco-IOS-XR-um-hostname-cfg:/hostname" -> ("Cisco-IOS-XR-um-hostname-cfg", true)
func ExtractNamespaceFromXPath(xPath string) (string, bool) {
	if idx := strings.Index(xPath, ":/"); idx > 0 {
		namespacePrefix := strings.TrimPrefix(xPath[:idx], "/")
		return namespacePrefix, true
	}
	return "", false
}

// getNamespaceURL returns the full namespace URL for a given prefix
func getNamespaceURL(prefix string) string {
	if prefix != "" {
		return namespaceBaseURL + prefix
	}
	return ""
}

// CleanSegmentName removes namespace prefix and predicate from segment
// Example: "Cisco-IOS-XR:interface[name='test']" -> "interface"
func CleanSegmentName(segment string) string {
	cleaned := RemoveNamespacePrefix(segment)
	if idx := strings.IndexByte(cleaned, '['); idx != -1 {
		cleaned = cleaned[:idx]
	}
	return cleaned
}

// ============================================================================
// Path Building and Manipulation Utilities
// ============================================================================

// BuildPathString joins path segments with dots
func BuildPathString(segments []string) string {
	return strings.Join(segments, ".")
}

// BuildParentPath returns the parent path by removing the last segment
func BuildParentPath(pathSegments []string) string {
	if len(pathSegments) <= 1 {
		return ""
	}
	return strings.Join(pathSegments[:len(pathSegments)-1], ".")
}

// BuildXmlnsPath builds an xmldot path with .@xmlns suffix
func BuildXmlnsPath(path string) string {
	return path + ".@xmlns"
}

// BuildCountPath builds a count path with .# suffix
func BuildCountPath(path string) string {
	return path + ".#"
}

// BuildPathFromSegments builds dotted path from segments with optional current path
// Example: (["element"], "parent") -> "parent.element"
func BuildPathFromSegments(segments []string, currentPath string) string {
	if currentPath != "" {
		return currentPath + "." + segments[0]
	}
	return segments[0]
}

// BuildTentativePath creates a tentative path by appending a new element
// Example: (["parent", "child"], "element") -> "parent.child.element"
func BuildTentativePath(pathSegments []string, escapedElementName string) string {
	return strings.Join(append(pathSegments[:len(pathSegments):len(pathSegments)], escapedElementName), ".")
}

// ToDotPath converts path to dot notation
func ToDotPath(path string) string {
	// Remove leading slash
	path = strings.TrimPrefix(path, "/")

	// Split by /, escape dots in each segment (which are part of YANG element names),
	// then join with . (path separator)
	parts := strings.Split(path, "/")
	for i, part := range parts {
		// Escape dots in element names for sjson/gjson
		// In Go string literals, \\\\ becomes \\ in the string, which sjson interprets as \
		parts[i] = strings.ReplaceAll(part, ".", "\\\\.")
	}
	path = strings.Join(parts, ".")

	// Replace double slashes with single dot (if any remain)
	path = strings.ReplaceAll(path, "//", ".")

	return path
}

// RemoveLastPathElement removes last element of path
func RemoveLastPathElement(p string) string {
	if idx := strings.LastIndex(p, "/"); idx != -1 {
		return p[:idx]
	}
	return ""
}

// GetLastPathElement gets last element of path
func GetLastPathElement(path string) string {
	// Remove namespace prefix if present
	// e.g., "ipv4//Cisco-IOS-XR-um-if-ip-address-cfg:addresses/address/address" -> "address"
	// Split by / and get the last non-empty element
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] != "" {
			// Remove namespace prefix if present (e.g., "Cisco-IOS-XR-um:element" -> "element")
			element := parts[i]
			if idx := strings.LastIndex(element, ":"); idx >= 0 {
				element = element[idx+1:]
			}
			return element
		}
	}
	return ""
}

// LastElement returns the last element of a YANG path with its namespace prefix.
// Example: "Cisco-IOS-XE-native:native/interface/GigabitEthernet=1" -> "Cisco-IOS-XE-native:GigabitEthernet"
func LastElement(path string) string {
	pes := strings.Split(path, "/")
	var prefix, element string
	for _, pe := range pes {
		// remove key
		if strings.Contains(pe, "=") {
			pe = pe[:strings.Index(pe, "=")]
		}
		if strings.Contains(pe, ":") {
			prefix = strings.Split(pe, ":")[0]
			element = strings.Split(pe, ":")[1]
		} else {
			element = pe
		}
	}
	return prefix + ":" + element
}

// ============================================================================
// XPath Utilities
// ============================================================================

// IsListPath checks if an XPath represents a list item (ends with a predicate).
func IsListPath(xPath string) bool {
	return strings.HasSuffix(strings.TrimSpace(xPath), "]")
}

// NormalizeModuleXPath converts IOS-XR module-prefixed XPaths
// Example: `MODULE:/some/path` -> `MODULE:some/path`
func NormalizeModuleXPath(xPath string) string {
	idx := strings.Index(xPath, ":/")
	if idx == -1 {
		return xPath
	}
	if strings.HasPrefix(xPath, "/") {
		return xPath
	}
	module := xPath[:idx]
	rest := xPath[idx+2:]
	if module == "" || rest == "" {
		return xPath
	}
	return module + ":" + rest
}

// SplitXPathSegments splits an XPath into segments while respecting bracket boundaries
func SplitXPathSegments(xPath string) []string {
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

// PrepareXPathSegments normalizes and cleans XPath segments for building structures.
// This includes filtering empty segments and merging namespace prefixes.
func PrepareXPathSegments(xPath string) []string {
	xPath = NormalizeModuleXPath(xPath)
	xPath = strings.TrimPrefix(xPath, "/")
	segments := SplitXPathSegments(xPath)

	// Filter empty segments
	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	// Merge leading namespace prefix with next segment
	if len(segments) > 0 && strings.HasSuffix(segments[0], ":") {
		namespace := strings.TrimSuffix(segments[0], ":")
		if len(segments) > 1 {
			segments[1] = namespace + ":" + segments[1]
			segments = segments[1:]
		}
	}

	return segments
}

// PrepareGetFromXPathSegments normalizes and prepares segments for reading from XPath.
// This merges namespace segments (e.g., ["ns:", "element"] -> ["ns:element"]).
func PrepareGetFromXPathSegments(xPath string) []string {
	xPath = NormalizeModuleXPath(xPath)
	xPath = strings.TrimPrefix(xPath, "/")
	segments := SplitXPathSegments(xPath)

	// Filter empty segments
	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	// Merge namespace segments (ns: + element)
	mergedSegments := make([]string, 0, len(segments))
	for i := 0; i < len(segments); i++ {
		if strings.HasSuffix(segments[i], ":") && i+1 < len(segments) {
			mergedSegments = append(mergedSegments, segments[i]+segments[i+1])
			i++
		} else {
			mergedSegments = append(mergedSegments, segments[i])
		}
	}

	return mergedSegments
}

// KeyValue represents a key-value pair in XPath predicates
type KeyValue struct {
	Key   string
	Value string
}

// ParseXPathSegment parses an XPath segment to extract the element name and key-value pairs
// Example: "interface[name='GigabitEthernet0/0/0/0']" returns ("interface", [{"name", "GigabitEthernet0/0/0/0"}])
func ParseXPathSegment(segment string) (string, []KeyValue) {
	elementName := segment
	keys := []KeyValue{}

	if idx := strings.Index(segment, "["); idx != -1 {
		elementName = segment[:idx]
		// Match predicates: [key='value'] or [key="value"]
		predicatePattern := `\[([^]]+)\]`
		re := regexp.MustCompile(predicatePattern)
		predicates := re.FindAllStringSubmatch(segment[idx:], -1)
		for _, pred := range predicates {
			if len(pred) < 2 {
				continue
			}
			// Split by " and " to handle multiple predicates
			parts := strings.Split(pred[1], " and ")
			for _, part := range parts {
				kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
				if len(kv) != 2 {
					continue
				}
				k := strings.TrimSpace(kv[0])
				v := strings.Trim(strings.TrimSpace(kv[1]), "'\"")
				keys = append(keys, KeyValue{Key: k, Value: v})
			}
		}
	}

	return elementName, keys
}

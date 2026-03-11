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
	return types.ListValueMust(types.StringType, GetValueSlice(result))
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
	return types.SetValueMust(types.StringType, GetValueSlice(result))
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

// getStringAttrValuesXML is a helper that extracts string attr.Values from xmldot results
func getStringAttrValuesXML(result []xmldot.Result) []attr.Value {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return v
}

// getInt64AttrValuesXML is a helper that extracts int64 attr.Values from xmldot results
func getInt64AttrValuesXML(result []xmldot.Result) []attr.Value {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return v
}

// GetStringListXML converts xmldot results to Terraform string list
func GetStringListXML(result []xmldot.Result) types.List {
	return types.ListValueMust(types.StringType, getStringAttrValuesXML(result))
}

// GetInt64ListXML converts xmldot results to Terraform int64 list
func GetInt64ListXML(result []xmldot.Result) types.List {
	return types.ListValueMust(types.Int64Type, getInt64AttrValuesXML(result))
}

// GetStringSetXML converts xmldot results to Terraform string set
func GetStringSetXML(result []xmldot.Result) types.Set {
	return types.SetValueMust(types.StringType, getStringAttrValuesXML(result))
}

// GetInt64SetXML converts xmldot results to Terraform int64 set
func GetInt64SetXML(result []xmldot.Result) types.Set {
	return types.SetValueMust(types.Int64Type, getInt64AttrValuesXML(result))
}

// GetAllChildElements returns all child elements with the given name from a parent xmldot.Result.
// This is needed when multiple sibling elements have the same name in XML.
func GetAllChildElements(parent xmldot.Result, elementName string) []xmldot.Result {
	// Check if there are multiple elements using BuildCountPath
	count := parent.Get(BuildCountPath(elementName)).Int()

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
		elements[i] = RemoveNamespacePrefix(elements[i])
	}
	return strings.Join(elements, "/")
}

// ToGoName converts TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		value = RemoveNamespacePrefix(value)
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
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
	return RemoveNamespacePrefix(first)
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
	return GetNamespacePrefixFromSegment(first)
}

// ExtractNamespacePrefix extracts namespace prefix from segment if present
// Example: "Cisco-IOS-XR:interface" -> ("Cisco-IOS-XR", true)
func ExtractNamespacePrefix(segment string) (string, bool) {
	prefix := GetNamespacePrefixFromSegment(segment)
	return prefix, prefix != ""
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
	return BuildPathString(pathSegments[:len(pathSegments)-1])
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
	newSegments := append(pathSegments[:len(pathSegments):len(pathSegments)], escapedElementName)
	return BuildPathString(newSegments)
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
			return RemoveNamespacePrefix(parts[i])
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

// prepareXPathSegmentsBase is a helper that normalizes, splits, and filters XPath segments
func prepareXPathSegmentsBase(xPath string) []string {
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
	return filteredSegments
}

// PrepareXPathSegments normalizes and cleans XPath segments for building structures.
// This includes filtering empty segments and merging namespace prefixes.
func PrepareXPathSegments(xPath string) []string {
	segments := prepareXPathSegmentsBase(xPath)

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
	segments := prepareXPathSegmentsBase(xPath)

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

// ============================================================================
// XML String Manipulation Utilities
// ============================================================================

// FindRootElementBounds finds the start and end positions of root element tag
func FindRootElementBounds(xmlStr string) (startIdx, closeIdx int, found bool) {
	startIdx = strings.Index(xmlStr, "<")
	if startIdx == -1 {
		return 0, 0, false
	}

	closeIdx = strings.Index(xmlStr[startIdx:], ">")
	if closeIdx == -1 {
		return 0, 0, false
	}
	closeIdx += startIdx

	return startIdx, closeIdx, true
}

// ExtractElementName extracts the element name from XML
func ExtractElementName(xmlStr string, startIdx int) (string, bool) {
	nameEndIdx := strings.IndexAny(xmlStr[startIdx+1:], "> ")
	if nameEndIdx == -1 {
		return "", false
	}
	nameEndIdx += startIdx + 1
	return xmlStr[startIdx+1 : nameEndIdx], true
}

// CleanExistingNamespaces removes all xmlns declarations from a tag
func CleanExistingNamespaces(rootTag string) string {
	cleaned := rootTag
	// Remove standard xmlns declarations
	cleaned = regexp.MustCompile(`\s+xmlns="[^"]*"`).ReplaceAllString(cleaned, "")
	// Remove malformed namespace declarations
	cleaned = regexp.MustCompile(`\s+xmlns:_xmlns="[^"]*"`).ReplaceAllString(cleaned, "")
	cleaned = regexp.MustCompile(`\s+_xmlns:nc="[^"]*"`).ReplaceAllString(cleaned, "")
	return cleaned
}

// InsertNamespaceAfterElementName inserts xmlns attribute after element name
func InsertNamespaceAfterElementName(cleaned, elementName, namespaceURL string) string {
	insertPos := len("<" + elementName)
	return cleaned[:insertPos] + fmt.Sprintf(` xmlns="%s"`, namespaceURL) + cleaned[insertPos:]
}

// InjectXMLSibling injects extra XML fragment(s) as siblings just before the closing tag
// of the outermost element in xmlStr.
// Example: InjectXMLSibling("<logging xmlns="..."><x/></logging>", "<y/>")
// returns  "<logging xmlns="..."><x/><y/></logging>".
// If extraXML is empty or xmlStr has no closing root tag, xmlStr is returned unchanged.
func InjectXMLSibling(xmlStr string, extraXML string) string {
	extraXML = strings.TrimSpace(extraXML)
	if extraXML == "" {
		return xmlStr
	}
	// Find the last closing tag
	closeStart := strings.LastIndex(xmlStr, "</")
	if closeStart == -1 {
		return xmlStr
	}
	return xmlStr[:closeStart] + extraXML + xmlStr[closeStart:]
}

// ExtractInnerXML returns the XML content between the outermost element's opening and closing tags.
// Given "<logging xmlns="..."><suppress>...</suppress></logging>", it returns
// "<suppress>...</suppress>".
// If the input is empty or has no inner content, it returns "".
func ExtractInnerXML(xmlStr string) string {
	xmlStr = strings.TrimSpace(xmlStr)
	if xmlStr == "" {
		return ""
	}
	// Find end of opening tag (first '>')
	openEnd := strings.Index(xmlStr, ">")
	if openEnd == -1 {
		return ""
	}
	// Self-closing root element e.g. "<logging/>" has no inner content
	if strings.HasSuffix(xmlStr[:openEnd+1], "/>") {
		return ""
	}
	// Find the last closing tag (last '</')
	closeStart := strings.LastIndex(xmlStr, "</")
	if closeStart == -1 || closeStart <= openEnd {
		return ""
	}
	inner := strings.TrimSpace(xmlStr[openEnd+1 : closeStart])
	return inner
}

// CheckKeysMatch checks if an xmldot item matches all key-value pairs
func CheckKeysMatch(item xmldot.Result, keys []KeyValue) bool {
	for _, kv := range keys {
		keyName := RemoveNamespacePrefix(kv.Key)
		keyResult := item.Get(keyName)
		if !keyResult.Exists() || keyResult.String() != kv.Value {
			return false
		}
	}
	return true
}

// FindElementByKeys finds element matching all key predicates in XML
// Returns the path segments and whether the element was found
func FindElementByKeys(xml, currentPath, escapedElementName string, keys []KeyValue, count int) ([]string, bool) {
	pathSoFar := strings.Split(currentPath, ".")

	if count > 1 {
		// Multiple elements - search for matching keys
		for idx := 0; idx < count; idx++ {
			indexedPath := fmt.Sprintf("%s.%d", currentPath, idx)
			item := xmldot.Get(xml, indexedPath)
			if CheckKeysMatch(item, keys) {
				pathSoFar[len(pathSoFar)-1] = fmt.Sprintf("%s.%d", escapedElementName, idx)
				return pathSoFar, true
			}
		}
	} else {
		// Single element - check if it matches
		item := xmldot.Get(xml, currentPath)
		if CheckKeysMatch(item, keys) {
			return pathSoFar, true
		}
	}

	return pathSoFar, false
}

// BuildFinalResult builds final xmldot result from path, handling arrays
func BuildFinalResult(xml string, pathSoFar []string) xmldot.Result {
	finalPath := BuildPathString(pathSoFar)
	count := xmldot.Get(xml, BuildCountPath(finalPath)).Int()

	if count > 1 && len(pathSoFar) >= 2 {
		parentPath := BuildParentPath(pathSoFar)
		childName := pathSoFar[len(pathSoFar)-1]
		arrayPath := parentPath + ".#." + childName
		return xmldot.Get(xml, arrayPath)
	}

	return xmldot.Get(xml, finalPath)
}

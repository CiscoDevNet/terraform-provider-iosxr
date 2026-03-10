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
	"fmt"
	"regexp"
	"strings"

	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// ============================================================================
// Constants and Types
// ============================================================================

// Namespace base URL for Cisco YANG models
const namespaceBaseURL = "http://cisco.com/ns/yang/"

// Pre-compiled regular expressions for performance
var (
	predicatePattern = regexp.MustCompile(`\[([^]]+)\]`)
)

type keyValue struct {
	Key   string
	Value string
}

// ============================================================================
// String Utility Functions
// ============================================================================

// removeNamespacePrefix removes the namespace prefix from an element name
// Example: "Cisco-IOS-XR-um-logging-cfg:suppress" -> "suppress"
func removeNamespacePrefix(name string) string {
	if idx := strings.Index(name, ":"); idx != -1 {
		return name[idx+1:]
	}
	return name
}

// getNamespacePrefixFromSegment extracts the namespace prefix from a segment name
// Example: "Cisco-IOS-XR-um-logging-correlator-cfg:suppress" -> "Cisco-IOS-XR-um-logging-correlator-cfg"
func getNamespacePrefixFromSegment(elementName string) string {
	if idx := strings.Index(elementName, ":"); idx != -1 {
		return elementName[:idx]
	}
	return ""
}

// getNamespaceURL returns the full namespace URL for a given prefix
func getNamespaceURL(prefix string) string {
	if prefix != "" {
		return namespaceBaseURL + prefix
	}
	return ""
}

// normalizeModuleXPath converts IOS-XR module-prefixed XPaths
// Example: `MODULE:/some/path` -> `MODULE:some/path`
func normalizeModuleXPath(xPath string) string {
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

// ============================================================================
// XPath Parsing Functions
// ============================================================================

// splitXPathSegments splits an XPath into segments while respecting bracket boundaries
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

// parseXPathSegment parses an XPath segment to extract the element name and key-value pairs
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
			parts := strings.Split(pred[1], " and ")
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

// ============================================================================
// Namespace Handling Functions
// ============================================================================

// probeIndexCount returns the actual number of same-named sibling elements
func probeIndexCount(xml string, basePath string) int {
	const maxProbeCount = 100
	for n := 0; n < maxProbeCount; n++ {
		if !xmldot.Get(xml, fmt.Sprintf("%s.%d", basePath, n)).Exists() {
			return n
		}
	}
	return maxProbeCount
}

// findNamespaceAwareSibling selects the correct index among multiple same-named sibling elements
// by matching the expected xmlns attribute.
func findNamespaceAwareSibling(xml string, currentPath string, count int, nsPrefix string, parentPath string) (idx int, needsIndex bool, found bool) {
	if nsPrefix == "" {
		return 0, false, true
	}

	expectedNS := getNamespaceURL(nsPrefix)

	resolveEffectiveNS := func(itemPath string) string {
		ns := xmldot.Get(xml, itemPath+".@xmlns").String()
		if ns != "" {
			return ns
		}
		if parentPath != "" {
			parentNS := xmldot.Get(xml, parentPath+".@xmlns").String()
			if parentNS != "" {
				return parentNS
			}
		}
		return ""
	}

	if count <= 1 {
		probed := probeIndexCount(xml, currentPath)
		if probed > count {
			count = probed
		}
	}

	if count <= 1 {
		item := xmldot.Get(xml, currentPath)
		if !item.Exists() {
			return 0, false, false
		}
		effectiveNS := resolveEffectiveNS(currentPath)
		if effectiveNS != "" && effectiveNS != expectedNS {
			return 0, false, false
		}
		return 0, false, true
	}

	for i := 0; i < count; i++ {
		indexedPath := fmt.Sprintf("%s.%d", currentPath, i)
		effectiveNS := resolveEffectiveNS(indexedPath)
		if effectiveNS == expectedNS {
			return i, true, true
		}
	}
	return 0, false, false
}

// augmentNamespaces adds xmlns attributes to elements that have namespace prefixes
func augmentNamespaces(body netconf.Body, path string) netconf.Body {
	segments := strings.Split(path, ".")
	xml := body.Res()

	elementToNamespace := make(map[string]string)

	currentPath := ""
	for segIdx, segment := range segments {
		cleanSegment := removeNamespacePrefix(segment)
		if idx := strings.IndexByte(cleanSegment, '['); idx != -1 {
			cleanSegment = cleanSegment[:idx]
		}

		if currentPath != "" {
			currentPath += "."
		}
		currentPath += cleanSegment

		if idx := strings.Index(segment, ":"); idx != -1 {
			prefix := segment[:idx]
			key := fmt.Sprintf("%d:%s", segIdx, currentPath)
			elementToNamespace[key] = prefix
		}
	}

	currentPath = ""
	for segIdx, segment := range segments {
		cleanSegment := removeNamespacePrefix(segment)
		if idx := strings.IndexByte(cleanSegment, '['); idx != -1 {
			cleanSegment = cleanSegment[:idx]
		}

		if currentPath != "" {
			currentPath += "."
		}
		currentPath += cleanSegment

		key := fmt.Sprintf("%d:%s", segIdx, currentPath)
		if prefix, hasNamespace := elementToNamespace[key]; hasNamespace {
			namespace := namespaceBaseURL + prefix

			if segIdx == 0 {
				pattern := fmt.Sprintf(`<(%s)(\s[^>]*?)?(/?)>`, regexp.QuoteMeta(cleanSegment))
				re := regexp.MustCompile(pattern)
				match := re.FindStringIndex(xml)
				if match != nil {
					matchStart, matchEnd := match[0], match[1]
					matchStr := xml[matchStart:matchEnd]

					if !strings.Contains(matchStr, `xmlns="`) {
						var insertPos int
						if strings.HasSuffix(matchStr, "/>") {
							insertPos = len(matchStr) - 2
						} else {
							insertPos = len(matchStr) - 1
						}
						attrStr := fmt.Sprintf(` xmlns="%s"`, namespace)
						newMatch := matchStr[:insertPos] + attrStr + matchStr[insertPos:]
						xml = xml[:matchStart] + newMatch + xml[matchEnd:]
					}
				}
			} else {
				xmldotPath := currentPath + ".@xmlns"
				existingNS := xmldot.Get(xml, xmldotPath).String()

				if existingNS == "" {
					xml, _ = xmldot.Set(xml, xmldotPath, namespace)
				}
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

		// Get the root element tag
		rootTag := xmlStr[startIdx : closeIdx+1]

		// Find the element name
		nameEndIdx := strings.IndexAny(xmlStr[startIdx+1:], "> ")
		if nameEndIdx == -1 {
			return xmlStr
		}
		nameEndIdx += startIdx + 1
		elementName := xmlStr[startIdx+1 : nameEndIdx]

		// Check if xmlns= is already present in the root element tag
		if strings.Contains(rootTag, "xmlns=") {
			// Remove ALL xmlns="..." patterns to avoid duplicates
			cleaned := rootTag
			xmlnsPattern := regexp.MustCompile(`\s+xmlns="[^"]*"`)
			cleaned = xmlnsPattern.ReplaceAllString(cleaned, "")

			// Also remove any malformed namespace declarations
			cleaned = regexp.MustCompile(`\s+xmlns:_xmlns="[^"]*"`).ReplaceAllString(cleaned, "")
			cleaned = regexp.MustCompile(`\s+_xmlns:nc="[^"]*"`).ReplaceAllString(cleaned, "")

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

// ============================================================================
// XPath Structure Building (Core Infrastructure)
// ============================================================================

// buildXPathStructure creates all elements in an XPath, including keys and namespaces
func buildXPathStructure(body netconf.Body, xPath string, ensureStructure bool) (netconf.Body, []string, []string) {
	xPath = normalizeModuleXPath(xPath)
	xPath = strings.TrimPrefix(xPath, "/")

	segments := splitXPathSegments(xPath)

	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	if len(segments) > 0 && strings.HasSuffix(segments[0], ":") {
		namespace := strings.TrimSuffix(segments[0], ":")
		if len(segments) > 1 {
			segments[1] = namespace + ":" + segments[1]
			segments = segments[1:]
		}
	}

	pathSegments := make([]string, 0, len(segments))
	originalSegments := make([]string, 0, len(segments))

	for i, segment := range segments {
		elementName, keys := parseXPathSegment(segment)
		originalSegments = append(originalSegments, elementName)
		cleanElementName := removeNamespacePrefix(elementName)
		escapedElementName := strings.ReplaceAll(cleanElementName, ".", `\.`)

		nsPrefix := getNamespacePrefixFromSegment(elementName)

		isAugmentedChild := false
		if len(keys) == 0 && nsPrefix != "" && len(pathSegments) > 0 {
			tentativePath := strings.Join(append(pathSegments[:len(pathSegments):len(pathSegments)], escapedElementName), ".")
			existingElement := xmldot.Get(body.Res(), tentativePath)
			if !existingElement.Exists() {
				isAugmentedChild = true
			}
		}

		if isAugmentedChild {
			pathSegments = append(pathSegments, escapedElementName)
		} else if nsPrefix != "" {
			expectedNS := namespaceBaseURL + nsPrefix
			tentativePath := strings.Join(append(pathSegments[:len(pathSegments):len(pathSegments)], escapedElementName), ".")
			xmlnsPath := tentativePath + ".@xmlns"
			existingNS := xmldot.Get(body.Res(), xmlnsPath).String()

			if existingNS == "" && len(pathSegments) > 0 {
				parentPath := strings.Join(pathSegments, ".")
				parentNS := xmldot.Get(body.Res(), parentPath+".@xmlns").String()
				if parentNS != "" {
					existingNS = parentNS
				}
			}

			if existingNS != "" && existingNS != expectedNS {
				siblingIdx := -1
				const maxSiblings = 32
				for n := 1; n <= maxSiblings; n++ {
					idxPath := fmt.Sprintf("%s.%d", tentativePath, n)
					ns := xmldot.Get(body.Res(), idxPath+".@xmlns").String()
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
				if siblingIdx >= 0 {
					pathSegments = append(pathSegments, fmt.Sprintf("%s.%d", escapedElementName, siblingIdx))
				} else {
					pathSegments = append(pathSegments, escapedElementName)
				}
			} else {
				pathSegments = append(pathSegments, escapedElementName)
			}
		} else {
			tentativePath := strings.Join(append(pathSegments[:len(pathSegments):len(pathSegments)], escapedElementName), ".")
			xmlnsPath := tentativePath + ".@xmlns"
			existingNS := xmldot.Get(body.Res(), xmlnsPath).String()

			shouldCreateSibling := false
			if existingNS != "" && len(pathSegments) > 0 {
				parentPath := strings.Join(pathSegments, ".")
				parentNS := xmldot.Get(body.Res(), parentPath+".@xmlns").String()

				if existingNS != parentNS {
					shouldCreateSibling = true
				}
			}

			if shouldCreateSibling {
				siblingIdx := -1
				const maxSiblings = 32
				for n := 1; n <= maxSiblings; n++ {
					idxPath := fmt.Sprintf("%s.%d", tentativePath, n)
					ns := xmldot.Get(body.Res(), idxPath+".@xmlns").String()
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
				if siblingIdx >= 0 {
					pathSegments = append(pathSegments, fmt.Sprintf("%s.%d", escapedElementName, siblingIdx))
				} else {
					pathSegments = append(pathSegments, escapedElementName)
				}
			} else {
				pathSegments = append(pathSegments, escapedElementName)
			}
		}

		fullPath := strings.Join(pathSegments, ".")

		if len(keys) > 0 {
			for _, kv := range keys {
				keyPath := fullPath + "." + kv.Key
				body = body.Set(keyPath, kv.Value)
			}
		}

		if nsPrefix != "" {
			// Check if path contains indexed sibling notation (e.g., ".0", ".1", etc.)
			isIndexedSibling := regexp.MustCompile(`\.\d+`).MatchString(fullPath)

			if isIndexedSibling {
				namespace := namespaceBaseURL + nsPrefix
				xmlnsPath := fullPath + ".@xmlns"
				existingNS := xmldot.Get(body.Res(), xmlnsPath).String()
				if existingNS == "" {
					body = body.Set(xmlnsPath, namespace)
				}
			} else {
				originalPath := strings.Join(originalSegments[:i+1], ".")
				body = augmentNamespaces(body, originalPath)
			}
		}
	}

	if ensureStructure && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		result := xmldot.Get(body.Res(), fullPath)
		if !result.Exists() {
			body = body.Set(fullPath, "")
			originalPath := strings.Join(originalSegments, ".")
			body = augmentNamespaces(body, originalPath)
		}
	}

	return body, pathSegments, originalSegments
}

// ============================================================================
// Public XPath Query and Manipulation Functions
// ============================================================================

// GetFromXPath reads from an xmldot.Result using an XPath that may contain predicates
func GetFromXPath(res xmldot.Result, xPath string) xmldot.Result {
	xPath = normalizeModuleXPath(xPath)
	xPath = strings.TrimPrefix(xPath, "/")
	segments := splitXPathSegments(xPath)

	filteredSegments := make([]string, 0, len(segments))
	for _, seg := range segments {
		if seg != "" {
			filteredSegments = append(filteredSegments, seg)
		}
	}
	segments = filteredSegments

	mergedSegments := make([]string, 0, len(segments))
	for i := 0; i < len(segments); i++ {
		if strings.HasSuffix(segments[i], ":") && i+1 < len(segments) {
			mergedSegments = append(mergedSegments, segments[i]+segments[i+1])
			i++
		} else {
			mergedSegments = append(mergedSegments, segments[i])
		}
	}
	segments = mergedSegments

	xml := res.Raw
	pathSoFar := make([]string, 0, len(segments))

	for _, segment := range segments {
		rawElementName, keys := parseXPathSegment(segment)
		nsPrefix := getNamespacePrefixFromSegment(rawElementName)
		elementName := removeNamespacePrefix(rawElementName)
		escapedElementName := strings.ReplaceAll(elementName, ".", `\.`)
		pathSoFar = append(pathSoFar, escapedElementName)
		currentPath := strings.Join(pathSoFar, ".")
		countPath := currentPath + ".#"
		count := xmldot.Get(xml, countPath).Int()

		if nsPrefix != "" {
			parentPath := strings.Join(pathSoFar[:len(pathSoFar)-1], ".")
			if idx, needsIndex, found := findNamespaceAwareSibling(xml, currentPath, int(count), nsPrefix, parentPath); found {
				if needsIndex {
					pathSoFar[len(pathSoFar)-1] = fmt.Sprintf("%s.%d", escapedElementName, idx)
					currentPath = strings.Join(pathSoFar, ".")
				}
				count = 1
			} else {
				return xmldot.Result{}
			}
		}

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
						pathSoFar[len(pathSoFar)-1] = fmt.Sprintf("%s.%d", escapedElementName, idx)
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

// SetFromXPath creates all elements in an XPath and optionally sets a value
func SetFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	xPath = normalizeModuleXPath(xPath)

	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments, originalSegments := buildXPathStructure(body, xPath, ensureStructure)

	if hasValue && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		body = body.Set(fullPath, value)
		originalPath := strings.Join(originalSegments, ".")
		body = augmentNamespaces(body, originalPath)
	}

	return body
}

// AppendFromXPath creates all elements in an XPath and appends a value to a list
func AppendFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	xPath = normalizeModuleXPath(xPath)

	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments, originalSegments := buildXPathStructure(body, xPath, ensureStructure)

	if hasValue && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".") + ".-1"
		body = body.Set(fullPath, value)
		originalPath := strings.Join(originalSegments, ".")
		body = augmentNamespaces(body, originalPath)
	}

	return body
}

// RemoveFromXPath creates all elements in an XPath with an nc:operation="remove" attribute
func RemoveFromXPath(body netconf.Body, xPath string) netconf.Body {
	xPath = normalizeModuleXPath(xPath)

	body, pathSegments, originalSegments := buildXPathStructure(body, xPath, true)

	if len(pathSegments) > 0 {
		targetPath := strings.Join(pathSegments, ".")

		originalPath := strings.Join(originalSegments, ".")
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

	// Post-process XML to fix any malformed elements created by indexed siblings
	// Remove any empty/malformed tags like "</>"
	xml = strings.ReplaceAll(xml, "</>", "")
	xml = strings.ReplaceAll(xml, "< >", "")
	xml = strings.ReplaceAll(xml, "</ >", "")

	return xml, nil
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

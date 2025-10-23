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
	"encoding/json"
	"fmt"
	"html"
	"strings"
)

// GetNetconfXml converts XPath to NETCONF XML
func GetNetconfXml(xpath, operation, body string) (string, error) {
	if err := validateOperation(operation, body); err != nil {
		return "", err
	}

	rootNamespace, xpathParts, err := parseXPath(xpath)
	if err != nil {
		return "", err
	}

	var filter string
	if operation == "update" {
		predicateKeys := extractPredicateKeys(xpathParts)
		filter = buildUpdateXML(xpathParts, rootNamespace, body, predicateKeys)
	} else {
		filter = buildSubtreeXML(xpathParts, rootNamespace, operation, len(xpathParts))
	}

	// Wrap delete/update operations
	if operation == "delete" || operation == "update" {
		filter = fmt.Sprintf("<config>%s</config>", filter)
	}

	return filter, nil
}

func validateOperation(operation, body string) error {
	if operation != "get" && operation != "delete" && operation != "update" {
		return fmt.Errorf("operation must be 'get', 'delete', or 'update'")
	}
	if operation == "update" && body == "" {
		return fmt.Errorf("body is required for update operations")
	}
	return nil
}

func parseXPath(xpath string) (string, []string, error) {
	if xpath == "" {
		return "", nil, fmt.Errorf("xpath cannot be empty")
	}

	colonIdx := strings.Index(xpath, ":")
	if colonIdx <= 0 {
		return "", nil, fmt.Errorf("xpath must include Cisco-IOS-XR- namespace prefix")
	}

	module := xpath[:colonIdx]
	if !strings.HasPrefix(module, "Cisco-IOS-XR-") {
		return "", nil, fmt.Errorf("invalid namespace prefix: %s", module)
	}

	rootNamespace := fmt.Sprintf("http://cisco.com/ns/yang/%s", module)
	xpathWithoutNS := strings.TrimPrefix(xpath[colonIdx+1:], "/")
	parts := splitXPath(xpathWithoutNS)

	if len(parts) == 0 {
		return "", nil, fmt.Errorf("invalid xpath")
	}

	return rootNamespace, parts, nil
}

// buildUpdateXML builds XML with preserved JSON key order
func buildUpdateXML(parts []string, rootNamespace, body string, excludeKeys map[string]bool) string {
	return buildUpdateXMLRecursive(parts, rootNamespace, body, excludeKeys, len(parts), len(parts))
}

func buildUpdateXMLRecursive(parts []string, rootNamespace, body string, excludeKeys map[string]bool, totalDepth, currentDepth int) string {
	if len(parts) == 0 {
		return ""
	}

	tagName, predicate := parseElement(parts[0])
	if validateXMLName(tagName) != nil {
		return ""
	}

	elementNS := extractNamespaceFromElement(parts[0])
	isLeaf := len(parts) == 1
	isRoot := currentDepth == totalDepth

	var elem string
	if isLeaf {
		bodyXML := jsonToXML(body, excludeKeys)
		elem = formatElement(tagName, predicate+bodyXML)
	} else {
		children := buildUpdateXMLRecursive(parts[1:], "", body, excludeKeys, totalDepth, currentDepth-1)
		elem = formatElement(tagName, predicate+children)
	}

	return applyNamespace(elem, tagName, elementNS, rootNamespace, isRoot)
}

func parseElement(part string) (string, string) {
	tagName := removeNamespacePrefix(part)
	tagNameClean, predicate := parsePartWithPredicate(tagName)
	return tagNameClean, predicate
}

func formatElement(tagName, content string) string {
	return fmt.Sprintf("<%s>%s</%s>", tagName, content, tagName)
}

func applyNamespace(elem, tagName, elementNS, rootNS string, isRoot bool) string {
	if elementNS != "" {
		return addNamespaceToElement(elem, tagName, elementNS)
	}
	if isRoot && rootNS != "" {
		return addNamespaceToElement(elem, tagName, rootNS)
	}
	return elem
}

// jsonToXML converts JSON string to XML preserving key order
func jsonToXML(jsonStr string, excludeKeys map[string]bool) string {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return ""
	}

	keys := parseJSONKeys(jsonStr)
	var sb strings.Builder

	for _, key := range keys {
		if value, exists := data[key]; exists && !excludeKeys[key] {
			sb.WriteString(jsonValueToXML(key, value, excludeKeys))
		}
	}

	return sb.String()
}

// parseJSONKeys extracts top-level keys from JSON in order
func parseJSONKeys(jsonStr string) []string {
	var keys []string
	var key strings.Builder
	var inString, inKey, escape, afterColon bool
	depth := 0

	for i := 0; i < len(jsonStr); i++ {
		ch := jsonStr[i]

		if escape {
			escape = false
			if inKey {
				key.WriteByte(ch)
			}
			continue
		}

		if ch == '\\' {
			escape = true
			if inKey {
				key.WriteByte(ch)
			}
			continue
		}

		if ch == '"' {
			if inString {
				if inKey && depth == 1 && !afterColon {
					keys = append(keys, key.String())
					key.Reset()
					inKey = false
				}
				inString = false
			} else {
				inString = true
				if depth == 1 && !afterColon && isAfterKeyPosition(jsonStr, i) {
					inKey = true
				}
			}
		} else if inString && inKey {
			key.WriteByte(ch)
		} else if !inString {
			switch ch {
			case '{':
				depth++
				afterColon = false
			case '}':
				depth--
				afterColon = false
			case ':':
				if depth == 1 {
					afterColon = true
				}
			case ',':
				if depth == 1 {
					afterColon = false
				}
			}
		}
	}

	return keys
}

func isAfterKeyPosition(jsonStr string, pos int) bool {
	for i := pos - 1; i >= 0; i-- {
		ch := jsonStr[i]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			continue
		}
		return ch == '{' || ch == ','
	}
	return false
}

// jsonValueToXML converts a single JSON key-value pair to XML
func jsonValueToXML(key string, value interface{}, excludeKeys map[string]bool) string {
	if excludeKeys[key] {
		return ""
	}

	xmlKey, nsPrefix := splitNamespacedKey(key)
	if validateXMLName(xmlKey) != nil {
		return ""
	}

	switch v := value.(type) {
	case map[string]interface{}:
		return mapToXML(xmlKey, nsPrefix, v, excludeKeys)
	case []interface{}:
		return arrayToXML(xmlKey, nsPrefix, v, excludeKeys)
	case string:
		return scalarToXML(xmlKey, nsPrefix, html.EscapeString(v))
	case nil:
		return emptyElementXML(xmlKey, nsPrefix)
	default:
		return scalarToXML(xmlKey, nsPrefix, html.EscapeString(fmt.Sprintf("%v", v)))
	}
}

func splitNamespacedKey(key string) (string, string) {
	if !strings.Contains(key, ":") {
		return key, ""
	}
	parts := strings.SplitN(key, ":", 2)
	return parts[1], normalizeNamespace(parts[0])
}

func mapToXML(xmlKey, nsPrefix string, m map[string]interface{}, excludeKeys map[string]bool) string {
	if len(m) == 0 {
		return emptyElementXML(xmlKey, nsPrefix)
	}

	childJSON, _ := json.Marshal(m)
	childXML := jsonToXML(string(childJSON), excludeKeys)
	return wrapWithNamespace(xmlKey, nsPrefix, childXML)
}

func arrayToXML(xmlKey, nsPrefix string, arr []interface{}, excludeKeys map[string]bool) string {
	var sb strings.Builder
	for _, item := range arr {
		switch v := item.(type) {
		case map[string]interface{}:
			childJSON, _ := json.Marshal(v)
			childXML := jsonToXML(string(childJSON), excludeKeys)
			sb.WriteString(wrapWithNamespace(xmlKey, nsPrefix, childXML))
		case nil:
			sb.WriteString(emptyElementXML(xmlKey, nsPrefix))
		default:
			escapedValue := html.EscapeString(fmt.Sprintf("%v", v))
			sb.WriteString(wrapWithNamespace(xmlKey, nsPrefix, escapedValue))
		}
	}
	return sb.String()
}

func scalarToXML(xmlKey, nsPrefix, value string) string {
	return wrapWithNamespace(xmlKey, nsPrefix, value)
}

func emptyElementXML(xmlKey, nsPrefix string) string {
	if nsPrefix != "" {
		return fmt.Sprintf(`<%s xmlns="%s"/>`, xmlKey, nsPrefix)
	}
	return fmt.Sprintf("<%s/>", xmlKey)
}

func wrapWithNamespace(xmlKey, nsPrefix, content string) string {
	if nsPrefix != "" {
		return fmt.Sprintf(`<%s xmlns="%s">%s</%s>`, xmlKey, nsPrefix, content, xmlKey)
	}
	return fmt.Sprintf("<%s>%s</%s>", xmlKey, content, xmlKey)
}

func splitXPath(xpath string) []string {
	var parts []string
	var current strings.Builder
	inBracket := false

	for _, char := range xpath {
		switch char {
		case '[':
			inBracket = true
			current.WriteRune(char)
		case ']':
			inBracket = false
			current.WriteRune(char)
		case '/':
			if !inBracket && current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			} else if inBracket {
				current.WriteRune(char)
			}
		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}

func extractPredicateKeys(parts []string) map[string]bool {
	keys := make(map[string]bool)
	for _, part := range parts {
		part = removeNamespacePrefix(part)
		if idx := strings.Index(part, "["); idx >= 0 {
			predicateStr := part[idx+1 : len(part)-1]
			if strings.Contains(predicateStr, "=") {
				kvParts := strings.SplitN(predicateStr, "=", 2)
				keys[strings.TrimSpace(kvParts[0])] = true
			}
		}
	}
	return keys
}

func buildSubtreeXML(parts []string, rootNamespace, operation string, totalDepth int) string {
	return buildSubtreeXMLRecursive(parts, rootNamespace, operation, totalDepth, totalDepth)
}

func buildSubtreeXMLRecursive(parts []string, rootNamespace, operation string, totalDepth, currentDepth int) string {
	if len(parts) == 0 {
		return ""
	}

	tagName, predicate := parseElement(parts[0])
	if validateXMLName(tagName) != nil {
		return ""
	}

	elementNS := extractNamespaceFromElement(parts[0])
	isLeaf := len(parts) == 1
	isRoot := currentDepth == totalDepth

	var elem string
	if isLeaf {
		deleteAttr := ""
		if operation == "delete" {
			deleteAttr = ` xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" nc:operation="delete"`
		}

		if predicate != "" {
			elem = fmt.Sprintf("<%s%s>%s</%s>", tagName, deleteAttr, predicate, tagName)
		} else {
			elem = fmt.Sprintf("<%s%s/>", tagName, deleteAttr)
		}
	} else {
		children := buildSubtreeXMLRecursive(parts[1:], "", operation, totalDepth, currentDepth-1)
		elem = formatElement(tagName, predicate+children)
	}

	return applyNamespace(elem, tagName, elementNS, rootNamespace, isRoot)
}

func addNamespaceToElement(elem, tagName, namespace string) string {
	patterns := []struct {
		old string
		new string
	}{
		{fmt.Sprintf("<%s ", tagName), fmt.Sprintf(`<%s xmlns="%s" `, tagName, namespace)},
		{fmt.Sprintf("<%s>", tagName), fmt.Sprintf(`<%s xmlns="%s">`, tagName, namespace)},
		{fmt.Sprintf("<%s/>", tagName), fmt.Sprintf(`<%s xmlns="%s"/>`, tagName, namespace)},
	}

	for _, p := range patterns {
		if strings.HasPrefix(elem, p.old) {
			return strings.Replace(elem, p.old, p.new, 1)
		}
	}
	return elem
}

func parsePartWithPredicate(part string) (string, string) {
	if !strings.Contains(part, "[") {
		return part, ""
	}

	idx := strings.Index(part, "[")
	tagName := part[:idx]
	predicateStr := part[idx+1 : len(part)-1]

	if !strings.Contains(predicateStr, "=") {
		return tagName, ""
	}

	kvParts := strings.SplitN(predicateStr, "=", 2)
	key := strings.TrimSpace(kvParts[0])
	value := strings.Trim(strings.TrimSpace(kvParts[1]), "\"'")

	if validateXMLName(key) != nil {
		return tagName, ""
	}

	escapedValue := html.EscapeString(value)
	return tagName, fmt.Sprintf("<%s>%s</%s>", key, escapedValue, key)
}

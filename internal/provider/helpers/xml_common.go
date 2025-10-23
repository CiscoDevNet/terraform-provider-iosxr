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
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

var validXMLName = regexp.MustCompile(`^[a-zA-Z_][\w\.\-]*$`)

// validateXMLName validates that a name is a valid XML element name
func validateXMLName(name string) error {
	if !validXMLName.MatchString(name) {
		return fmt.Errorf("invalid XML element name: %s", name)
	}
	return nil
}

// extractNamespace extracts the xmlns attribute from XML attributes
func extractNamespace(attrs []xml.Attr) string {
	for _, attr := range attrs {
		if attr.Name.Local == "xmlns" {
			return attr.Value
		}
	}
	return ""
}

// extractNamespacePrefix extracts the last part of a namespace URL
func extractNamespacePrefix(namespace string) string {
	parts := strings.Split(namespace, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return namespace
}

// removeNamespacePrefix removes Cisco-IOS-XR namespace prefix from element name
func removeNamespacePrefix(elementName string) string {
	if idx := strings.Index(elementName, ":"); idx > 0 {
		if strings.HasPrefix(elementName[:idx], "Cisco-IOS-XR-") {
			return elementName[idx+1:]
		}
	}
	return elementName
}

// extractNamespaceFromElement extracts the full namespace URL from an element name
func extractNamespaceFromElement(elementName string) string {
	if idx := strings.Index(elementName, ":"); idx > 0 {
		module := elementName[:idx]
		if strings.HasPrefix(module, "Cisco-IOS-XR-") {
			return fmt.Sprintf("http://cisco.com/ns/yang/%s", module)
		}
	}
	return ""
}

// normalizeNamespace converts a short namespace to a full URL if needed
func normalizeNamespace(ns string) string {
	if strings.HasPrefix(ns, "http://") || strings.HasPrefix(ns, "https://") {
		return ns
	}
	return fmt.Sprintf("http://cisco.com/ns/yang/%s", ns)
}

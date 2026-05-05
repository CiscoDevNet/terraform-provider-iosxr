// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
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

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// ============================================================================
// Yang Resource Model - Serialization/Deserialization
// ============================================================================
//
// This file implements the `iosxr_yang` resource which provides generic YANG
// path-based configuration for IOS-XR devices.
//
// **Protocol-Specific Naming Convention:**
//
//   gNMI (JSON IETF encoding):
//     • toBody(ctx)       → Serialize to JSON for gNMI Set operations
//     • fromBody(ctx, res) → Deserialize from gNMI Get/Subscribe JSON responses
//
//   NETCONF (XML encoding):
//     • toBodyXML(ctx)       → Serialize to XML for NETCONF edit-config
//     • fromBodyXML(ctx, res) → Deserialize from NETCONF get-config XML responses
//
// Both protocols support:
//   - Attributes (key-value pairs, with special markers <NULL> and <EMPTY>)
//   - Lists (keyed collections with nested attributes)
//   - RPL normalization (routing policy language text formatting)
//
// ============================================================================

type Yang struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Delete     types.Bool   `tfsdk:"delete"`
	Attributes types.Map    `tfsdk:"attributes"`
	Lists      []YangList   `tfsdk:"lists"`
}

type YangList struct {
	Name types.String `tfsdk:"name"`
	Key  types.String `tfsdk:"key"`
	// Items is defined in schema as list(map(string)). Using types.List here ensures
	// Terraform Framework can reliably deserialize it for both plan and state.
	Items  types.List `tfsdk:"items"`
	Values types.List `tfsdk:"values"`
}

// extractKeysFromPath extracts key-value pairs from XPath predicates
// Example: "/path[key1='value1'][key2='value2']" -> map["key1"]="value1", map["key2"]="value2"
// Also handles: "/path[key1=value1][key2=value2]" (without quotes)
func extractKeysFromPath(path string) map[string]string {
	keys := make(map[string]string)

	// Match patterns like [key='value'], [key="value"], or [key=value] (no quotes)
	re := regexp.MustCompile(`\[([^=\]]+)='([^']+)'\]|\[([^=\]]+)="([^"]+)"\]|\[([^=\]]+)=([^\]]+)\]`)
	matches := re.FindAllStringSubmatch(path, -1)

	for _, match := range matches {
		var key, value string
		if match[1] != "" {
			// Single quote match
			key = strings.TrimSpace(match[1])
			value = match[2]
		} else if match[3] != "" {
			// Double quote match
			key = strings.TrimSpace(match[3])
			value = match[4]
		} else if match[5] != "" {
			// No quote match
			key = strings.TrimSpace(match[5])
			value = strings.TrimSpace(match[6])
		}
		if key != "" {
			keys[key] = value
		}
	}

	return keys
}

// extractTerminalSegmentKeysFromPath extracts key-value pairs ONLY from the
// terminal (last) path segment. This prevents parent-segment keys from being
// injected as child elements of a different (non-keyed) terminal node.
//
// Example:
//
//	"Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF2]/rd/two-byte-as" → {} (terminal has no keys)
//	"Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF2]"                → {"vrf-name": "VRF2"}
func extractTerminalSegmentKeysFromPath(path string) map[string]string {
	// Normalise: strip the "MODULE:/" prefix that IOS-XR paths carry
	// e.g. "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[…]" → "/vrfs/vrf[…]"
	if colonSlash := strings.Index(path, ":/"); colonSlash != -1 {
		path = path[colonSlash+1:]
	}

	// Split the path into segments, respecting brackets
	var segments []string
	var current strings.Builder
	bracketDepth := 0

	for _, r := range path {
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

	if len(segments) == 0 {
		return map[string]string{}
	}

	// Only look at the last segment
	lastSegment := segments[len(segments)-1]
	return extractKeysFromPath(lastSegment)
}

type YangDataSourceModel struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Attributes types.Map    `tfsdk:"attributes"`
	Lists      []YangList   `tfsdk:"lists"`
}

func (data Yang) getPath() string {
	return data.Path.ValueString()
}

func (data YangDataSourceModel) getPath() string {
	return data.Path.ValueString()
}

// if last path element has a key -> remove it
func (data Yang) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

// extractKeysFromPath extracts key names from a YANG path
// e.g., "path[key1=value1][key2=value2]" returns ["key1", "key2"]
func (data Yang) extractKeysFromPath() map[string]bool {
	keys := make(map[string]bool)
	path := data.getPath()

	// Match patterns like [key=value]
	re := regexp.MustCompile(`\[([^=\]]+)=`)
	matches := re.FindAllStringSubmatch(path, -1)

	for _, match := range matches {
		if len(match) > 1 {
			keys[match[1]] = true
		}
	}

	return keys
}

// ============================================================================
// gNMI Serialization (JSON IETF)
// ============================================================================

// toBody serializes the Yang resource to JSON (gNMI/JSON_IETF encoding).
func (data Yang) toBody(ctx context.Context) string {
	body := ""

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	// Extract keys from the path that should be filtered out
	keysInPath := data.extractKeysFromPath()

	// Filter attributes to exclude keys that are in the path
	filteredAttributes := make(map[string]string)
	for attr, value := range attributes {
		// Skip attributes that are keys in the path
		if !keysInPath[attr] {
			// Don't filter out <NULL> - we'll handle it specially in toBody
			// It's used for presence containers/leaves
			filteredAttributes[attr] = value
		}
	}

	// Check if we have content after filtering
	hasContent := len(filteredAttributes) > 0 || len(data.Lists) > 0

	// Special case for gNMI: if we have no content after filtering keys,
	// but we have keys in the path, include those keys in the body.
	// This allows creating presence containers (like interfaces) with just keys.
	if !hasContent && len(keysInPath) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("toBody: No attributes after filtering keys, including keys in body for presence container: %v", keysInPath))
		// Use the keys as the body content
		filteredAttributes = make(map[string]string)
		for key := range keysInPath {
			if val, exists := attributes[key]; exists {
				filteredAttributes[key] = val
			}
		}
		hasContent = len(filteredAttributes) > 0
	}

	if !hasContent {
		return ""
	}

	// For gNMI, don't wrap in path element - the path already specifies the container
	// Send only the leaf values directly
	body = "{}"
	for attr, value := range filteredAttributes {
		originalAttr := attr
		attr = strings.ReplaceAll(attr, "/", ".")
		if value == "<NULL>" {
			// For <NULL> marker, set as [null] array (presence leaf without value)
			body, _ = sjson.Set(body, attr, []interface{}{nil})
		} else if value == "<EMPTY>" {
			// For explicit <EMPTY> marker, set an empty object (presence container)
			body, _ = sjson.Set(body, attr, map[string]interface{}{})
		} else if value == "" {
			// For empty strings on container paths (containing /), skip them
			// as gNMI will reject "data is presented at none leaf node"
			// Only leaf nodes should have values set
			if strings.Contains(originalAttr, "/") {
				// This looks like a container path, skip it
				tflog.Debug(ctx, fmt.Sprintf("Skipping empty value for container path: %s", originalAttr))
				continue
			}
			// For leaf nodes with empty string, set empty object
			body, _ = sjson.Set(body, attr, map[string]interface{}{})
		} else {
			// Apply RPL normalization for attributes that contain RPL content
			if strings.Contains(attr, "rpl-") || strings.Contains(attr, "rpl_") {
				// Normalize RPL text: trim trailing spaces from each line, then ensure single trailing newline
				lines := strings.Split(value, "\n")
				for i := range lines {
					lines[i] = strings.TrimRight(lines[i], " \t\r")
				}
				value = strings.Join(lines, "\n")
				// Trim all trailing whitespace and add exactly one newline
				value = strings.TrimRight(value, " \t\n\r") + "\n"
			}
			body, _ = sjson.Set(body, attr, value)
		}
	}

	// Handle lists without path element prefix
	for i := range data.Lists {
		listName := strings.ReplaceAll(data.Lists[i].Name.ValueString(), "/", ".")
		if !data.Lists[i].Items.IsNull() && !data.Lists[i].Items.IsUnknown() && len(data.Lists[i].Items.Elements()) > 0 {
			body, _ = sjson.Set(body, listName, []interface{}{})
			for _, item := range data.Lists[i].Items.Elements() {
				listAttributes, _ := decodeListItem(ctx, item)
				itemBody := "{}"
				for attr, value := range listAttributes {
					attr = strings.ReplaceAll(attr, "/", ".")
					if value == "<EMPTY>" {
						itemBody, _ = sjson.Set(itemBody, attr, map[string]interface{}{})
					} else if value != "<NULL>" {
						if strings.Contains(attr, "rpl-") || strings.Contains(attr, "rpl_") {
							lines := strings.Split(value, "\n")
							for i := range lines {
								lines[i] = strings.TrimRight(lines[i], " \t\r")
							}
							value = strings.Join(lines, "\n")
							value = strings.TrimRight(value, " \t\n\r") + "\n"
						}
						itemBody, _ = sjson.Set(itemBody, attr, value)
					}
				}
				body, _ = sjson.SetRaw(body, listName+".-1", itemBody)
			}
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			var values []string
			data.Lists[i].Values.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, listName, values)
		}
	}

	return body
}

func (data Yang) toBodyXML(ctx context.Context) string {

	body := netconf.Body{}

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	// Extract keys from the path that should be filtered out
	keysInPath := data.extractKeysFromPath()

	tflog.Debug(ctx, fmt.Sprintf("toBodyXML: path=%s, attributes=%v, lists=%d, keysInPath=%v",
		data.Path.ValueString(), attributes, len(data.Lists), keysInPath))

	// Log detailed list information for debugging
	for i := range data.Lists {
		tflog.Debug(ctx, fmt.Sprintf("toBodyXML: List[%d] name=%s, Items.IsNull=%v, Items.IsUnknown=%v, Items.Elements=%d, Values.IsNull=%v, Values.IsUnknown=%v, Values.Elements=%d",
			i, data.Lists[i].Name.ValueString(),
			data.Lists[i].Items.IsNull(), data.Lists[i].Items.IsUnknown(), len(data.Lists[i].Items.Elements()),
			data.Lists[i].Values.IsNull(), data.Lists[i].Values.IsUnknown(), len(data.Lists[i].Values.Elements())))
	}

	// For NETCONF, key attributes must be included in the XML body as child elements
	// (unlike gNMI where they are encoded in the path). Inject path-key values first
	// for any keys the user has not explicitly provided in attributes.
	// Only inject keys from the terminal (last) path segment; parent-segment keys
	// must NOT be added as children of a different, deeper node.
	pathKeyValues := extractTerminalSegmentKeysFromPath(data.Path.ValueString())
	for keyAttr, keyValue := range pathKeyValues {
		if _, userProvided := attributes[keyAttr]; !userProvided {
			tflog.Debug(ctx, fmt.Sprintf("toBodyXML: Injecting path key attribute %s=%s into body", keyAttr, keyValue))
			body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+keyAttr, keyValue)
		}
	}

	for attr, value := range attributes {
		// For NETCONF, do NOT skip key attributes - they must be present in the XML body
		// so the device knows which list entry to create/modify.
		_ = keysInPath // retained for awareness; no longer used to filter
		// For <NULL>, create an empty element (presence leaf)
		if value == "<NULL>" {
			body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+attr, "")
			continue
		}
		// For <EMPTY>, create an empty container element
		if value == "<EMPTY>" {
			// Create empty element (just the structure, no value)
			body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+attr, "")
		} else if value == "" {
			// For empty strings on container paths (containing /), skip them
			// Only leaf nodes should have empty values set
			if strings.Contains(attr, "/") {
				// This looks like a container path, skip it
				tflog.Debug(ctx, fmt.Sprintf("Skipping empty value for container path in NETCONF: %s", attr))
				continue
			}
			// For leaf nodes with empty string, create empty element
			body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+attr, "")
		} else {
			// Apply RPL normalization for attributes that contain RPL content
			// This ensures consistency with fromBodyXML normalization
			if strings.Contains(attr, "rpl-") || strings.Contains(attr, "rpl_") {
				// Normalize RPL text: trim trailing spaces from each line, then ensure single trailing newline
				lines := strings.Split(value, "\n")
				for i := range lines {
					lines[i] = strings.TrimRight(lines[i], " \t\r")
				}
				value = strings.Join(lines, "\n")
				// Trim all trailing whitespace and add exactly one newline
				value = strings.TrimRight(value, " \t\n\r") + "\n"
			}
			body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+attr, value)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("toBodyXML: After attributes, body length=%d, processing %d lists", len(body.Res()), len(data.Lists)))

	// Handle lists - build them directly into the body structure
	for i := range data.Lists {
		listName := data.Lists[i].Name.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("toBodyXML: Processing list[%d]: name=%s", i, listName))

		// Check if this list has items (keyed list)
		if !data.Lists[i].Items.IsNull() && !data.Lists[i].Items.IsUnknown() && len(data.Lists[i].Items.Elements()) > 0 {
			tflog.Debug(ctx, fmt.Sprintf("toBodyXML: List '%s' has %d items", listName, len(data.Lists[i].Items.Elements())))

			for ii, item := range data.Lists[i].Items.Elements() {
				listAttributes, _ := decodeListItem(ctx, item)
				tflog.Debug(ctx, fmt.Sprintf("toBodyXML: Item[%d] has %d attributes", ii, len(listAttributes)))

				// Build the full XPath for this list item including the resource path
				listItemPath := data.Path.ValueString() + "/" + data.Lists[i].Name.ValueString()

				// Add key predicates if the list has keys
				if !data.Lists[i].Key.IsNull() && !data.Lists[i].Key.IsUnknown() && data.Lists[i].Key.ValueString() != "" {
					keys := strings.Split(data.Lists[i].Key.ValueString(), ",")
					var preds []string
					for _, k := range keys {
						k = strings.TrimSpace(k)
						if k == "" {
							continue
						}
						// Get key value from list item attributes
						if v, ok := listAttributes[k]; ok && v != "" {
							preds = append(preds, fmt.Sprintf("[%s='%s']", k, v))
						}
					}
					if len(preds) > 0 {
						listItemPath = listItemPath + strings.Join(preds, "")
					}
				}

				tflog.Debug(ctx, fmt.Sprintf("toBodyXML: Setting list item at path='%s'", listItemPath))

				// Set each attribute of the list item using SetFromXPath
				for attr, value := range listAttributes {
					if value == "<NULL>" {
						continue // Skip <NULL> values
					}
					attrPath := listItemPath + "/" + attr
					if value == "<EMPTY>" {
						body = helpers.SetFromXPath(body, attrPath, "")
					} else {
						// Apply RPL normalization
						if strings.Contains(attr, "rpl-") || strings.Contains(attr, "rpl_") {
							lines := strings.Split(value, "\n")
							for i := range lines {
								lines[i] = strings.TrimRight(lines[i], " \t\r")
							}
							value = strings.Join(lines, "\n")
							value = strings.TrimRight(value, " \t\n\r") + "\n"
						}
						body = helpers.SetFromXPath(body, attrPath, value)
					}
				}
			}
		} else if !data.Lists[i].Values.IsNull() && !data.Lists[i].Values.IsUnknown() && len(data.Lists[i].Values.Elements()) > 0 {
			// Check if this list has values (leaf-list)
			tflog.Debug(ctx, fmt.Sprintf("toBodyXML: List '%s' has %d leaf-list values", listName, len(data.Lists[i].Values.Elements())))
			var values []string
			data.Lists[i].Values.ElementsAs(ctx, &values, false)
			tflog.Debug(ctx, fmt.Sprintf("toBodyXML: Extracted %d values: %v", len(values), values))

			listPath := data.Path.ValueString() + "/" + data.Lists[i].Name.ValueString()
			tflog.Debug(ctx, fmt.Sprintf("toBodyXML: Appending leaf-list values at path='%s' with %d values", listPath, len(values)))

			// Use AppendFromXPath for each value individually (like generated resources do)
			for _, v := range values {
				body = helpers.AppendFromXPath(body, listPath, v)
			}
			tflog.Debug(ctx, fmt.Sprintf("toBodyXML: After appending leaf-list, body length=%d", len(body.Res())))
		} else {
			tflog.Debug(ctx, fmt.Sprintf("toBodyXML: List '%s' has no items or values - skipping", listName))
		}
	}

	bodyStr := body.Res()
	tflog.Debug(ctx, fmt.Sprintf("toBodyXML: Generated body length=%d", len(bodyStr)))
	// Add namespace declaration to root element
	bodyStr = helpers.AddNamespaceToRootElement(bodyStr, data.Path.ValueString())
	return bodyStr
}

// fromBody deserializes the JSON response from gNMI into the Yang struct.
func (data *Yang) fromBody(ctx context.Context, res []byte) {
	// For gNMI responses, the JSON structure directly contains the leaf values
	// without the path prefix, so we read attributes directly
	if !data.Attributes.IsNull() && !data.Attributes.IsUnknown() {
		attributes := make(map[string]string)
		data.Attributes.ElementsAs(ctx, &attributes, false)

		// Get keys that are in the path - these won't be in the response
		// but should be preserved in state
		keysInPath := data.extractKeysFromPath()

		tflog.Debug(ctx, fmt.Sprintf("fromBody: path=%s, keysInPath=%v, attributes=%v, response=%s",
			data.Path.ValueString(), keysInPath, attributes, string(res)))

		for attr := range attributes {
			// Skip attributes that are keys in the path - preserve their original values
			if keysInPath[attr] {
				tflog.Debug(ctx, fmt.Sprintf("fromBody: Skipping key attribute '%s' (preserving value '%s')", attr, attributes[attr]))
				continue
			}

			// If the original value was <NULL>, keep it as <NULL> (don't try to read from device)
			if attributes[attr] == "<NULL>" {
				tflog.Debug(ctx, fmt.Sprintf("fromBody: Preserving <NULL> value for attribute '%s'", attr))
				continue
			}

			attrPath := strings.ReplaceAll(attr, "/", ".")
			value := gjson.GetBytes(res, attrPath)

			if !value.Exists() ||
				value.Raw == "[null]" {
				// Value doesn't exist in device response
				// Preserve the planned value instead of setting to empty string
				// This handles optional/default attributes that device may not return
				tflog.Debug(ctx, fmt.Sprintf("fromBody: Attribute '%s' not in response, preserving planned value '%s'", attr, attributes[attr]))
				// Don't modify attributes[attr] - keep existing value
			} else if value.IsObject() && len(value.Map()) == 0 {
				// Empty container should be represented as <EMPTY>
				attributes[attr] = "<EMPTY>"
			} else {
				attrValue := value.String()
				// Apply RPL normalization for attributes that contain RPL content
				// (e.g., "rpl-route-policy", "rpl-community-set", etc.)
				if strings.Contains(attr, "rpl-") || strings.Contains(attr, "rpl_") {
					// Normalize RPL text: trim trailing spaces from each line, then ensure single trailing newline
					lines := strings.Split(attrValue, "\n")
					for i := range lines {
						lines[i] = strings.TrimRight(lines[i], " \t\r")
					}
					attrValue = strings.Join(lines, "\n")
					// Trim all trailing whitespace and add exactly one newline
					attrValue = strings.TrimRight(attrValue, " \t\n\r") + "\n"
				}
				attributes[attr] = attrValue
			}
		}

		tflog.Debug(ctx, fmt.Sprintf("fromBody: Final attributes after processing: %v", attributes))
		data.Attributes, _ = types.MapValueFrom(ctx, types.StringType, attributes)
	}

	for i := range data.Lists {
		keys := strings.Split(data.Lists[i].Key.ValueString(), ",")
		namePath := strings.ReplaceAll(data.Lists[i].Name.ValueString(), "/", ".")
		if !data.Lists[i].Items.IsNull() && !data.Lists[i].Items.IsUnknown() && len(data.Lists[i].Items.Elements()) > 0 {
			newElems := make([]attr.Value, 0, len(data.Lists[i].Items.Elements()))
			for _, item := range data.Lists[i].Items.Elements() {
				attributes, _ := decodeListItem(ctx, item)

				// Find item by key(s)
				var keyValues []string
				for _, key := range keys {
					keyValues = append(keyValues, attributes[key])
				}

				// find item by key(s)
				var r gjson.Result
				gjson.GetBytes(res, namePath).ForEach(
					func(_, v gjson.Result) bool {
						found := false
						for ik := range keys {
							keyPath := strings.ReplaceAll(keys[ik], "/", ".")
							if v.Get(keyPath).String() == keyValues[ik] {
								found = true
								continue
							}
							found = false
							break
						}
						if found {
							r = v
							return false
						}
						return true
					},
				)

				for attr := range attributes {
					attrPath := strings.ReplaceAll(attr, "/", ".")
					value := r.Get(attrPath)
					if !value.Exists() || value.Raw == "[null]" {
						attributes[attr] = ""
					} else if value.IsObject() && len(value.Map()) == 0 {
						attributes[attr] = "<EMPTY>"
					} else {
						attributes[attr] = value.String()
					}
				}

				mv, _ := types.MapValueFrom(ctx, types.StringType, attributes)
				newElems = append(newElems, mv)
			}
			data.Lists[i].Items = types.ListValueMust(types.MapType{ElemType: types.StringType}, newElems)
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			values := gjson.GetBytes(res, namePath)
			if values.IsArray() {
				data.Lists[i].Values = types.ListValueMust(data.Lists[i].Values.ElementType(ctx), helpers.GetValueSlice(values.Array()))
			}
		}
	}
}

// fromBodyXML deserializes the XML response from NETCONF into the Yang struct.
func (data *Yang) fromBodyXML(ctx context.Context, res xmldot.Result) {

	// Extract key-value pairs from path predicates
	// e.g., "/path[key1='value1'][key2='value2']" -> map["key1"]="value1", map["key2"]="value2"
	pathKeys := extractKeysFromPath(data.Path.ValueString())

	tflog.Debug(ctx, fmt.Sprintf("fromBodyXML: path=%s, extracted pathKeys=%v", data.Path.ValueString(), pathKeys))

	// Parse attributes
	attributes := data.Attributes.Elements()
	for attr := range attributes {
		// First check if this attribute is a key in the path
		if keyValue, isKey := pathKeys[attr]; isKey {
			tflog.Debug(ctx, fmt.Sprintf("fromBodyXML: Setting key attribute '%s' from path to '%s'", attr, keyValue))
			attributes[attr] = types.StringValue(keyValue)
		} else {
			// Not a key, read from XML response
			value := helpers.GetFromXPath(res, "data/"+data.Path.ValueString()+"/"+attr)
			if !value.Exists() || value.String() == "" {
				// Value doesn't exist in device response
				// Preserve the planned value instead of setting to empty string
				// This handles optional/default attributes that device may not return
				tflog.Debug(ctx, fmt.Sprintf("fromBodyXML: Attribute '%s' not in response, preserving planned value", attr))
				// Don't modify attributes[attr] - keep existing value
			} else {
				attrValue := value.String()
				// Apply RPL normalization for attributes that contain RPL content
				// (e.g., "rpl-route-policy", "rpl-community-set", etc.)
				if strings.Contains(attr, "rpl-") || strings.Contains(attr, "rpl_") {
					// Normalize RPL text: trim trailing spaces from each line, then ensure single trailing newline
					lines := strings.Split(attrValue, "\n")
					for i := range lines {
						lines[i] = strings.TrimRight(lines[i], " \t\r")
					}
					attrValue = strings.Join(lines, "\n")
					// Trim all trailing whitespace and add exactly one newline
					attrValue = strings.TrimRight(attrValue, " \t\n\r") + "\n"
				}
				attributes[attr] = types.StringValue(attrValue)
			}
		}
	}
	data.Attributes = types.MapValueMust(types.StringType, attributes)

	// Parse lists
	for i := range data.Lists {
		keys := strings.Split(data.Lists[i].Key.ValueString(), ",")
		listName := data.Lists[i].Name.ValueString()

		if !data.Lists[i].Items.IsNull() && !data.Lists[i].Items.IsUnknown() && len(data.Lists[i].Items.Elements()) > 0 {
			newElems := make([]attr.Value, 0, len(data.Lists[i].Items.Elements()))
			for _, item := range data.Lists[i].Items.Elements() {
				itemAttributes, _ := decodeListItem(ctx, item)

				// Build key values
				var keyValues []string
				for _, key := range keys {
					keyValues = append(keyValues, itemAttributes[key])
				}

				xpathPredicates := ""
				for ik, key := range keys {
					if ik > 0 {
						xpathPredicates += " and "
					}
					xpathPredicates += key + "='" + keyValues[ik] + "'"
				}
				itemXPath := listName + "[" + xpathPredicates + "]"
				itemResult := helpers.GetFromXPath(res, "data/"+data.Path.ValueString()+"/"+itemXPath)

				// Update attributes from response, preserving planned values where absent
				for attr := range itemAttributes {
					value := helpers.GetFromXPath(itemResult, attr)
					if !value.Exists() || value.String() == "" {
						// preserve
						continue
					}
					attrValue := value.String()
					if strings.Contains(attr, "rpl-") || strings.Contains(attr, "rpl_") {
						lines := strings.Split(attrValue, "\n")
						for i := range lines {
							lines[i] = strings.TrimRight(lines[i], " \t\r")
						}
						attrValue = strings.Join(lines, "\n")
						attrValue = strings.TrimRight(attrValue, " \t\n\r") + "\n"
					}
					itemAttributes[attr] = attrValue
				}

				mv, _ := types.MapValueFrom(ctx, types.StringType, itemAttributes)
				newElems = append(newElems, mv)
			}
			data.Lists[i].Items = types.ListValueMust(types.MapType{ElemType: types.StringType}, newElems)
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			// Simple leaf-list values
			listResult := helpers.GetFromXPath(res, "data/"+data.Path.ValueString()+"/"+listName)
			if listResult.IsArray() {
				values := make([]attr.Value, 0)
				for _, v := range listResult.Array() {
					values = append(values, types.StringValue(v.String()))
				}
				data.Lists[i].Values = types.ListValueMust(data.Lists[i].Values.ElementType(ctx), values)
			}
		}
	}
}

// decodeListItem attempts to decode a list item represented as either a types.Map
// (legacy) or an attr.Value (from elements of types.List) into map[string]string.
func decodeListItem(ctx context.Context, v attr.Value) (map[string]string, error) {
	out := make(map[string]string)
	if v == nil {
		return out, nil
	}
	if v.IsNull() || v.IsUnknown() {
		return out, nil
	}

	// Most common: item is a map(string)
	if m, ok := v.(types.Map); ok {
		diags := m.ElementsAs(ctx, &out, false)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to decode list item map")
		}
		return out, nil
	}

	// In case the value gets decoded as an object with string attrs.
	if o, ok := v.(types.Object); ok {
		attrs := o.Attributes()
		for k, av := range attrs {
			if sv, ok := av.(types.String); ok {
				if !sv.IsNull() && !sv.IsUnknown() {
					out[k] = sv.ValueString()
				}
			}
		}
		return out, nil
	}

	return out, nil
}

func (data *Yang) getDeletedItems(ctx context.Context, state Yang) []string {
	deletedItems := make([]string, 0)
	for l := range state.Lists {
		name := state.Lists[l].Name.ValueString()
		namePath := strings.ReplaceAll(name, "/", ".")
		keys := strings.Split(state.Lists[l].Key.ValueString(), ",")
		var dataList YangList
		for _, dl := range data.Lists {
			if dl.Name.ValueString() == name {
				dataList = dl
			}
		}

		stateItemElems := []attr.Value{}
		if !state.Lists[l].Items.IsNull() && !state.Lists[l].Items.IsUnknown() {
			stateItemElems = state.Lists[l].Items.Elements()
		}
		planItemElems := []attr.Value{}
		if !dataList.Items.IsNull() && !dataList.Items.IsUnknown() {
			planItemElems = dataList.Items.Elements()
		}

		if len(stateItemElems) > 0 {
			// check if state item is also included in plan, if not delete item
			for _, sItem := range stateItemElems {
				slia, _ := decodeListItem(ctx, sItem)

				// if state key values are empty move on to next item
				emptyKey := false
				for _, key := range keys {
					if slia[key] == "" {
						emptyKey = true
						break
					}
				}
				if emptyKey {
					continue
				}

				// find data (plan) item with matching key values
				found := false
				for _, pItem := range planItemElems {
					dlia, _ := decodeListItem(ctx, pItem)
					for _, key := range keys {
						if dlia[key] == slia[key] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						break
					}
				}

				// if no matching item in plan found -> delete
				if !found {
					keyValues := make([]string, len(keys))
					for k, key := range keys {
						keyValues[k] = slia[key]
					}
					deletedItems = append(deletedItems, state.getPath()+"/"+namePath+"="+strings.Join(keyValues, ","))
				}
			}
		} else if len(state.Lists[l].Values.Elements()) > 0 {
			var slv []string
			state.Lists[l].Values.ElementsAs(ctx, &slv, false)
			// check if state value is also included in plan, if not delete value from list
			for _, stateValue := range slv {
				found := false
				var dlv []string
				dataList.Values.ElementsAs(ctx, &dlv, false)
				for _, dataValue := range dlv {
					if stateValue == dataValue {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, state.getPath()+"/"+namePath+"="+stateValue)
				}
			}
		}
	}
	return deletedItems
}

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

type Yang struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Delete     types.Bool   `tfsdk:"delete"`
	Attributes types.Map    `tfsdk:"attributes"`
	Lists      []YangList   `tfsdk:"lists"`
}

type YangList struct {
	Name   types.String `tfsdk:"name"`
	Key    types.String `tfsdk:"key"`
	Items  []types.Map  `tfsdk:"items"`
	Values types.List   `tfsdk:"values"`
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

	// Only create the wrapper if we have actual content after filtering
	hasContent := len(filteredAttributes) > 0 || len(data.Lists) > 0
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
		if len(data.Lists[i].Items) > 0 {
			body, _ = sjson.Set(body, listName, []interface{}{})
			for ii := range data.Lists[i].Items {
				var listAttributes map[string]string
				data.Lists[i].Items[ii].ElementsAs(ctx, &listAttributes, false)
				itemBody := "{}"
				for attr, value := range listAttributes {
					attr = strings.ReplaceAll(attr, "/", ".")
					if value == "<EMPTY>" {
						itemBody, _ = sjson.Set(itemBody, attr, map[string]interface{}{})
					} else if value != "<NULL>" {
						// Apply RPL normalization for list item attributes that contain RPL content
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

	for attr, value := range attributes {
		// Skip attributes that are keys in the path
		if keysInPath[attr] {
			continue
		}
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
	for i := range data.Lists {
		if len(data.Lists[i].Items) > 0 {
			for ii := range data.Lists[i].Items {
				var listAttributes map[string]string
				data.Lists[i].Items[ii].ElementsAs(ctx, &listAttributes, false)
				attrs := netconf.Body{}
				for attr, value := range listAttributes {
					// Skip attributes with <NULL> value
					if value == "<NULL>" {
						continue
					}
					// For <EMPTY>, create an empty container element
					if value == "<EMPTY>" {
						attrs = helpers.SetFromXPath(attrs, attr, "")
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
						attrs = helpers.SetFromXPath(attrs, attr, value)
					}
				}
				body = helpers.SetRawFromXPath(body, data.Path.ValueString()+"/"+data.Lists[i].Name.ValueString(), attrs.Res())
			}
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			var values []string
			data.Lists[i].Values.ElementsAs(ctx, &values, false)
			body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+data.Lists[i].Name.ValueString(), values)
		}
	}

	bodyStr := body.Res()
	// Add namespace declaration to root element
	bodyStr = helpers.AddNamespaceToRootElement(bodyStr, data.Path.ValueString())
	return bodyStr
}

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
		if len(data.Lists[i].Items) > 0 {
			for ii := range data.Lists[i].Items {
				var keyValues []string
				for _, key := range keys {
					v, _ := data.Lists[i].Items[ii].Elements()[key].ToTerraformValue(ctx)
					var keyValue string
					v.As(&keyValue)
					keyValues = append(keyValues, keyValue)
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

				attributes := make(map[string]string)
				data.Lists[i].Items[ii].ElementsAs(ctx, &attributes, false)
				for attr := range attributes {
					attrPath := strings.ReplaceAll(attr, "/", ".")
					value := r.Get(attrPath)
					if !value.Exists() ||
						value.Raw == "[null]" {
						attributes[attr] = ""
					} else if value.IsObject() && len(value.Map()) == 0 {
						// Empty container should be represented as <EMPTY>
						attributes[attr] = "<EMPTY>"
					} else {
						attributes[attr] = value.String()
					}
				}
				data.Lists[i].Items[ii], _ = types.MapValueFrom(ctx, types.StringType, attributes)
			}
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			values := gjson.GetBytes(res, namePath)
			if values.IsArray() {
				data.Lists[i].Values = types.ListValueMust(data.Lists[i].Values.ElementType(ctx), helpers.GetValueSlice(values.Array()))
			}
		}
	}
}

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

		if len(data.Lists[i].Items) > 0 {
			// Complex list items with multiple attributes
			for ii := range data.Lists[i].Items {
				// Get key values from plan to find the matching item
				var keyValues []string
				for _, key := range keys {
					v, _ := data.Lists[i].Items[ii].Elements()[key].ToTerraformValue(ctx)
					var keyValue string
					v.As(&keyValue)
					keyValues = append(keyValues, keyValue)
				}

				// Build XPath to find the specific list item by key(s)
				xpathPredicates := ""
				for ik, key := range keys {
					if ik > 0 {
						xpathPredicates += " and "
					}
					xpathPredicates += key + "='" + keyValues[ik] + "'"
				}
				itemXPath := listName + "[" + xpathPredicates + "]"

				// Find the matching list item in XML response
				itemResult := helpers.GetFromXPath(res, "data/"+data.Path.ValueString()+"/"+itemXPath)

				tflog.Debug(ctx, fmt.Sprintf("fromBodyXML: Looking for list item with XPath='%s', found=%v",
					"data/"+data.Path.ValueString()+"/"+itemXPath, itemResult.Exists()))

				// Parse attributes from the matched item
				itemAttributes := data.Lists[i].Items[ii].Elements()
				for attr := range itemAttributes {
					value := helpers.GetFromXPath(itemResult, attr)
					tflog.Debug(ctx, fmt.Sprintf("fromBodyXML: List item attr='%s', valueExists=%v, value='%s'",
						attr, value.Exists(), value.String()))
					if !value.Exists() || value.String() == "" {
						// For keys, if not in XML, use the value from the plan
						// Check if this attr is a key
						isKey := false
						for _, key := range keys {
							if attr == key {
								isKey = true
								break
							}
						}
						if isKey {
							// Keep the existing value from plan (already in itemAttributes)
							continue
						}
						// For non-key attributes, if device doesn't return a value,
						// preserve the planned value (don't overwrite with empty)
						// This handles optional/default attributes
						tflog.Debug(ctx, fmt.Sprintf("fromBodyXML: List item attribute '%s' not in response, preserving planned value", attr))
						// Keep existing value - don't overwrite
						continue
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
						itemAttributes[attr] = types.StringValue(attrValue)
					}
				}
				data.Lists[i].Items[ii] = types.MapValueMust(types.StringType, itemAttributes)
			}
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
		if len(state.Lists[l].Items) > 0 {
			// check if state item is also included in plan, if not delete item
			for i := range state.Lists[l].Items {
				var slia map[string]string
				state.Lists[l].Items[i].ElementsAs(ctx, &slia, false)

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
				for dli := range dataList.Items {
					var dlia map[string]string
					dataList.Items[dli].ElementsAs(ctx, &dlia, false)
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

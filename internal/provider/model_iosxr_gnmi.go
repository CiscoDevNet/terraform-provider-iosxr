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

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	EMPTY_TAG string = "<EMPTY>"
)

type Gnmi struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Delete     types.Bool   `tfsdk:"delete"`
	Attributes types.Map    `tfsdk:"attributes"`
	Lists      []GnmiList   `tfsdk:"lists"`
}

type GnmiList struct {
	Name   types.String   `tfsdk:"name"`
	Key    types.String   `tfsdk:"key"`
	Items  []types.Map    `tfsdk:"items"`
	Values []types.String `tfsdk:"values"`
}

type GnmiData struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Attributes types.Map    `tfsdk:"attributes"`
}

func (data Gnmi) getPath() string {
	return data.Path.ValueString()
}

func (data Gnmi) toBody(ctx context.Context) string {
	body := "{}"

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	for attr, value := range attributes {
		attr = strings.ReplaceAll(attr, "/", ".")
		tflog.Debug(ctx, fmt.Sprintf("Setting attribute %s to %s", attr, value))
		if value == EMPTY_TAG {
			body, _ = sjson.Set(body, attr, map[string]interface{}{})
		} else {
			body, _ = sjson.Set(body, attr, value)
		}
	}

	for i := range data.Lists {
		listName := strings.ReplaceAll(data.Lists[i].Name.ValueString(), "/", ".")
		if len(data.Lists[i].Items) > 0 {
			body, _ = sjson.Set(body, listName, []interface{}{})
			for ii := range data.Lists[i].Items {
				var listAttributes map[string]string
				data.Lists[i].Items[ii].ElementsAs(ctx, &listAttributes, false)
				attrs := ""
				for attr, value := range listAttributes {
					attr = strings.ReplaceAll(attr, "/", ".")
					if value == EMPTY_TAG {
						attrs, _ = sjson.Set(attrs, attr, map[string]interface{}{})
					} else {
						attrs, _ = sjson.Set(attrs, attr, value)
					}
				}
				body, _ = sjson.SetRaw(body, listName+".-1", attrs)
			}
		} else if len(data.Lists[i].Values) > 0 {
			for _, value := range data.Lists[i].Values {
				body, _ = sjson.Set(body, listName+".-1", value.ValueString())
			}
		}
	}

	return body
}

func (data *Gnmi) fromBody(ctx context.Context, res []byte) diag.Diagnostics {
	var diags diag.Diagnostics
	// Extract a list of keys from the path
	keys := make([]string, 0)
	path := data.Path.ValueString()
	if strings.HasSuffix(path, "]") {
		keyValuePairs := strings.Split(path[strings.LastIndex(path, "[")+1:len(path)-1], ",")
		for _, v := range keyValuePairs {
			keys = append(keys, strings.Split(v, "=")[0])
		}
	}

	attributes := data.Attributes.Elements()
	for attr := range attributes {
		attrPath := strings.ReplaceAll(attr, "/", ".")
		value := gjson.GetBytes(res, attrPath)
		if value.IsObject() && len(value.Map()) == 0 {
			attributes[attr] = types.StringValue(EMPTY_TAG)
		} else if !value.Exists() || value.Raw == "[null]" {
			if !helpers.Contains(keys, attr) {
				attributes[attr] = types.StringValue("")
			}
		} else {
			attributes[attr] = types.StringValue(value.String())
		}
	}
	if len(attributes) > 0 {
		data.Attributes = types.MapValueMust(types.StringType, attributes)
	}

	for i := range data.Lists {
		keys := strings.Split(data.Lists[i].Key.ValueString(), ",")
		namePath := strings.ReplaceAll(data.Lists[i].Name.ValueString(), "/", ".")
		if len(data.Lists[i].Items) > 0 {
			for ii := range data.Lists[i].Items {
				var keyValues []string
				for _, key := range keys {
					elements := data.Lists[i].Items[ii].Elements()
					element, ok := elements[key]
					if !ok {
						diags.AddError("Missing key", fmt.Sprintf("Cannot locate key '%s' in list item: %v", key, elements))
						return diags
					}
					v, _ := element.ToTerraformValue(ctx)
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
							if v.Get(keys[ik]).String() == keyValues[ik] {
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

				listAttributes := data.Lists[i].Items[ii].Elements()
				for attr := range listAttributes {
					attrPath := strings.ReplaceAll(attr, "/", ".")
					value := r.Get(attrPath)
					if value.IsObject() && len(value.Map()) == 0 {
						listAttributes[attr] = types.StringValue(EMPTY_TAG)
					} else if !value.Exists() || value.Raw == "[null]" {
						listAttributes[attr] = types.StringValue("")
					} else {
						listAttributes[attr] = types.StringValue(value.String())
					}
				}
				if len(listAttributes) > 0 {
					data.Lists[i].Items[ii] = types.MapValueMust(types.StringType, listAttributes)
				}
			}
		} else if len(data.Lists[i].Values) > 0 {
			values := gjson.GetBytes(res, namePath)
			if values.IsArray() {
				data.Lists[i].Values = helpers.GetStringSlice(values.Array())
			}
		}
	}
	return diags
}

func (data *Gnmi) getDeletedItems(ctx context.Context, state Gnmi) []string {
	deletedItems := make([]string, 0)
	for l := range state.Lists {
		name := state.Lists[l].Name.ValueString()
		namePath := strings.ReplaceAll(name, "/", ".")
		keys := strings.Split(state.Lists[l].Key.ValueString(), ",")
		var dataList GnmiList
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
					keyString := ""
					for _, key := range keys {
						keyString += fmt.Sprintf("[%s=%s]", key, slia[key])
					}
					deletedItems = append(deletedItems, state.getPath()+"/"+namePath+keyString)
				}
			}
		}
	}
	return deletedItems
}

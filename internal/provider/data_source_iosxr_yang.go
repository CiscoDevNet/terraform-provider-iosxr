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
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &YangDataSource{}
	_ datasource.DataSourceWithConfigure = &YangDataSource{}
)

func NewYangDataSource() datasource.DataSource {
	return &YangDataSource{}
}

type YangDataSource struct {
	data *IosxrProviderData
}

func (d *YangDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_yang"
}

func (d *YangDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can retrieve one or more attributes via gNMI.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "A gNMI path, e.g. `openconfig-interfaces:/interfaces/interface`.",
				Required:            true,
			},
			"attributes": schema.MapAttribute{
				MarkdownDescription: "Map of key-value pairs which represents the attributes and its values.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"lists": schema.ListNestedAttribute{
				MarkdownDescription: "List of lists with items.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "List name.",
							Computed:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: "List key attribute name.",
							Computed:            true,
						},
						"values": schema.ListAttribute{
							MarkdownDescription: "List values.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"items": schema.ListAttribute{
							MarkdownDescription: "List items.",
							Computed:            true,
							ElementType:         types.MapType{ElemType: types.StringType},
						},
					},
				},
			},
		},
	}
}

func (d *YangDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxrProviderData)
}

func (d *YangDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state YangDataSourceModel

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, ok := d.data.Devices[config.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.ValueString()))

	attributes := make(map[string]attr.Value)

	if device.Managed {
		if device.Protocol == "gnmi" {
			// Ensure connection is healthy (reconnect if stale)
			if err := helpers.EnsureGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection); err != nil {
				resp.Diagnostics.AddError("gNMI Connection Error", fmt.Sprintf("Failed to ensure connection: %s", err))
				return
			}

			// Ensure connection is closed when function exits (if reuse disabled)
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)

			getResp, err := device.GnmiClient.Get(ctx, []string{config.Path.ValueString()})
			if err != nil {
				resp.Diagnostics.AddError("Unable to apply gNMI Get operation", err.Error())
				return
			}

			// Parse gNMI response - the JSON structure directly contains the leaf values
			respBody := getResp.Notifications[0].Update[0].Val.GetJsonIetfVal()
			tflog.Debug(ctx, fmt.Sprintf("gNMI data source response: %s", string(respBody)))

			// Parse response to extract attributes
			// The response might be nested under the last element of the path or at the root
			parsed := gjson.ParseBytes(respBody)

			// Try to find the data - check if it's nested under the last path element
			lastElement := helpers.LastElement(config.Path.ValueString())
			var dataToExtract gjson.Result

			if parsed.Get(lastElement).Exists() {
				// Data is nested under last path element
				dataToExtract = parsed.Get(lastElement)
				tflog.Debug(ctx, fmt.Sprintf("gNMI data source: data nested under '%s'", lastElement))
			} else {
				// Data is at root level
				dataToExtract = parsed
				tflog.Debug(ctx, "gNMI data source: data at root level")
			}

			// Extract attributes from the data
			for attrName, value := range dataToExtract.Map() {
				// Use attribute name as-is from response
				if !value.IsObject() && !value.IsArray() {
					attributes[attrName] = types.StringValue(value.String())
					tflog.Debug(ctx, fmt.Sprintf("gNMI data source: extracted attribute %s=%s", attrName, value.String()))
				} else if value.IsObject() && len(value.Map()) == 0 {
					// Empty container
					attributes[attrName] = types.StringValue("<EMPTY>")
				}
			}
		} else {
			// Serialize NETCONF operations (all ops when reuse disabled, reads concurrent when reuse enabled)
			locked := helpers.AcquireNetconfLock(&device.NetconfOpMutex, device.ReuseConnection, false)
			if locked {
				defer device.NetconfOpMutex.Unlock()
			}
			defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)

			// Ensure connection is healthy (reconnect if stale)
			if err := helpers.EnsureNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection); err != nil {
				resp.Diagnostics.AddError("NETCONF Connection Error", fmt.Sprintf("Failed to ensure connection: %s", err))
				return
			}

			filter := helpers.GetSubtreeFilter(config.Path.ValueString())
			res, err := device.NetconfClient.GetConfig(ctx, "running", filter)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object, got error: %s", err))
				return
			}

			// For NETCONF, parse XML response to extract all leaf elements
			// Extract key-value pairs from path predicates
			pathKeys := extractKeysFromPath(config.Path.ValueString())
			tflog.Debug(ctx, fmt.Sprintf("NETCONF data source: path=%s, extracted pathKeys=%v", config.Path.ValueString(), pathKeys))

			dataPath := "data/" + config.Path.ValueString()
			result := res.Res

			// Get the container element
			containerResult := helpers.GetFromXPath(result, dataPath)
			if containerResult.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("NETCONF data source: found container at path=%s, raw=%s", dataPath, containerResult.Raw))

				// containerResult.Raw contains XML - we need to extract all leaf elements
				// Use xmldot's Get method to access each child element by name
				// Parse the XML to discover all child element names
				rawXML := containerResult.Raw

				// Simple approach: parse XML to extract leaf element names and use xmldot.Get() for values
				// Look for patterns like <tagname>value</tagname> (leaf nodes without nested elements)
				// Go regex doesn't support backreferences, so extract tag names and validate separately
				leafPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)>([^<]+)</([a-zA-Z0-9\-_]+)>`)
				matches := leafPattern.FindAllStringSubmatch(rawXML, -1)

				for _, match := range matches {
					if len(match) >= 4 {
						openTag := match[1]
						closeTag := match[3]
						// Only process if open and close tags match
						if openTag == closeTag {
							// Skip namespace declarations and xmlns attributes
							if openTag != "" && !strings.HasPrefix(openTag, "xmlns") {
								// Use xmldot's Get to retrieve the actual value
								leafValue := containerResult.Get(openTag)
								if leafValue.Exists() {
									attributes[openTag] = types.StringValue(leafValue.String())
									tflog.Debug(ctx, fmt.Sprintf("NETCONF data source: extracted attribute %s=%s", openTag, leafValue.String()))
								}
							}
						}
					}
				}
			}

			// Add keys from path as attributes (these won't be in the XML response)
			for key, value := range pathKeys {
				if _, exists := attributes[key]; !exists {
					attributes[key] = types.StringValue(value)
				}
			}

			tflog.Debug(ctx, fmt.Sprintf("NETCONF data source read completed: path=%s, attributes=%v", dataPath, attributes))
		}
	}

	state.Path = types.StringValue(config.Path.ValueString())
	state.Id = types.StringValue(config.Path.ValueString())
	state.Attributes = types.MapValueMust(types.StringType, attributes)
	state.Lists = []YangList{}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

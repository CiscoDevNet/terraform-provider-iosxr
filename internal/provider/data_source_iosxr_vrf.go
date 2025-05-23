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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VRFDataSource{}
	_ datasource.DataSourceWithConfigure = &VRFDataSource{}
)

func NewVRFDataSource() datasource.DataSource {
	return &VRFDataSource{}
}

type VRFDataSource struct {
	client *client.Client
}

func (d *VRFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf"
}

func (d *VRFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the VRF configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: "VRF name",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "A description for the VRF",
				Computed:            true,
			},
			"vpn_id": schema.StringAttribute{
				MarkdownDescription: "VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part",
				Computed:            true,
			},
			"address_family_ipv4_unicast": schema.BoolAttribute{
				MarkdownDescription: "Unicast sub address family",
				Computed:            true,
			},
			"address_family_ipv4_unicast_import_route_policy": schema.StringAttribute{
				MarkdownDescription: "Use route-policy for import filtering",
				Computed:            true,
			},
			"address_family_ipv4_unicast_export_route_policy": schema.StringAttribute{
				MarkdownDescription: "Use route-policy for export",
				Computed:            true,
			},
			"address_family_ipv4_multicast": schema.BoolAttribute{
				MarkdownDescription: "Multicast topology",
				Computed:            true,
			},
			"address_family_ipv4_flowspec": schema.BoolAttribute{
				MarkdownDescription: "Flowspec sub address family",
				Computed:            true,
			},
			"address_family_ipv6_unicast": schema.BoolAttribute{
				MarkdownDescription: "Unicast sub address family",
				Computed:            true,
			},
			"address_family_ipv6_unicast_import_route_policy": schema.StringAttribute{
				MarkdownDescription: "Use route-policy for import filtering",
				Computed:            true,
			},
			"address_family_ipv6_unicast_export_route_policy": schema.StringAttribute{
				MarkdownDescription: "Use route-policy for export",
				Computed:            true,
			},
			"address_family_ipv6_multicast": schema.BoolAttribute{
				MarkdownDescription: "Multicast topology",
				Computed:            true,
			},
			"address_family_ipv6_flowspec": schema.BoolAttribute{
				MarkdownDescription: "Flowspec sub address family",
				Computed:            true,
			},
			"rd_two_byte_as_as_number": schema.StringAttribute{
				MarkdownDescription: "bgp as-number",
				Computed:            true,
			},
			"rd_two_byte_as_index": schema.Int64Attribute{
				MarkdownDescription: "ASN2:index (hex or decimal format)",
				Computed:            true,
			},
			"rd_four_byte_as_as_number": schema.StringAttribute{
				MarkdownDescription: "4-byte AS number",
				Computed:            true,
			},
			"rd_four_byte_as_index": schema.Int64Attribute{
				MarkdownDescription: "ASN2:index (hex or decimal format)",
				Computed:            true,
			},
			"rd_ip_address_ipv4_address": schema.StringAttribute{
				MarkdownDescription: "configure this node",
				Computed:            true,
			},
			"rd_ip_address_index": schema.Int64Attribute{
				MarkdownDescription: "IPv4Address:index (hex or decimal format)",
				Computed:            true,
			},
			"address_family_ipv4_unicast_import_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Two Byte AS Number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_import_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Four Byte AS number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_import_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "IPv4Address:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_export_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Two Byte AS Number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_export_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Four Byte AS number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_export_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "IPv4Address:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_import_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Two Byte AS Number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_import_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Four Byte AS number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_import_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "IPv4Address:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_export_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Two Byte AS Number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_export_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Four Byte AS number",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "ASN2:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_export_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: "IPv4Address:index (hex or decimal format)",
							Computed:            true,
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: "These are stitching RTs",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VRFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *VRFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VRFData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, err := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Unable to apply gNMI Get operation", err.Error())
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

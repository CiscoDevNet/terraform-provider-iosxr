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
	_ datasource.DataSource              = &RouterStaticVRFIPv6MulticastDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterStaticVRFIPv6MulticastDataSource{}
)

func NewRouterStaticVRFIPv6MulticastDataSource() datasource.DataSource {
	return &RouterStaticVRFIPv6MulticastDataSource{}
}

type RouterStaticVRFIPv6MulticastDataSource struct {
	client *client.Client
}

func (d *RouterStaticVRFIPv6MulticastDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_static_vrf_ipv6_multicast"
}

func (d *RouterStaticVRFIPv6MulticastDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router Static VRF IPv6 Multicast configuration.",

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
				MarkdownDescription: "VRF Static route configuration subcommands",
				Required:            true,
			},
			"prefix_address": schema.StringAttribute{
				MarkdownDescription: "Destination prefix",
				Required:            true,
			},
			"prefix_length": schema.Int64Attribute{
				MarkdownDescription: "Destination prefix length",
				Required:            true,
			},
			"nexthop_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Forwarding interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Forwarding interface",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "description of the static route",
							Computed:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: "Set tag for this route",
							Computed:            true,
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: "Distance metric for this route",
							Computed:            true,
						},
						"permanent": schema.BoolAttribute{
							MarkdownDescription: "Permanent route",
							Computed:            true,
						},
						"track": schema.StringAttribute{
							MarkdownDescription: "Enable object tracking for static route",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Set metric for this route",
							Computed:            true,
						},
					},
				},
			},
			"nexthop_interface_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Forwarding interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Forwarding interface",
							Computed:            true,
						},
						"address": schema.StringAttribute{
							MarkdownDescription: "Forwarding router's address",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "description of the static route",
							Computed:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: "Set tag for this route",
							Computed:            true,
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: "Distance metric for this route",
							Computed:            true,
						},
						"permanent": schema.BoolAttribute{
							MarkdownDescription: "Permanent route",
							Computed:            true,
						},
						"track": schema.StringAttribute{
							MarkdownDescription: "Enable object tracking for static route",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Set metric for this route",
							Computed:            true,
						},
					},
				},
			},
			"nexthop_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Forwarding router's address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "Forwarding router's address",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "description of the static route",
							Computed:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: "Set tag for this route",
							Computed:            true,
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: "Distance metric for this route",
							Computed:            true,
						},
						"permanent": schema.BoolAttribute{
							MarkdownDescription: "Permanent route",
							Computed:            true,
						},
						"track": schema.StringAttribute{
							MarkdownDescription: "Enable object tracking for static route",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Set metric for this route",
							Computed:            true,
						},
					},
				},
			},
			"vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "Destination VRF",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf_name": schema.StringAttribute{
							MarkdownDescription: "Destination VRF",
							Computed:            true,
						},
						"nexthop_interfaces": schema.ListNestedAttribute{
							MarkdownDescription: "Forwarding interface",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface_name": schema.StringAttribute{
										MarkdownDescription: "Forwarding interface",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "description of the static route",
										Computed:            true,
									},
									"tag": schema.Int64Attribute{
										MarkdownDescription: "Set tag for this route",
										Computed:            true,
									},
									"distance_metric": schema.Int64Attribute{
										MarkdownDescription: "Distance metric for this route",
										Computed:            true,
									},
									"permanent": schema.BoolAttribute{
										MarkdownDescription: "Permanent route",
										Computed:            true,
									},
									"track": schema.StringAttribute{
										MarkdownDescription: "Enable object tracking for static route",
										Computed:            true,
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: "Set metric for this route",
										Computed:            true,
									},
								},
							},
						},
						"nexthop_interface_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "Forwarding interface",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface_name": schema.StringAttribute{
										MarkdownDescription: "Forwarding interface",
										Computed:            true,
									},
									"address": schema.StringAttribute{
										MarkdownDescription: "Forwarding router's address",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "description of the static route",
										Computed:            true,
									},
									"tag": schema.Int64Attribute{
										MarkdownDescription: "Set tag for this route",
										Computed:            true,
									},
									"distance_metric": schema.Int64Attribute{
										MarkdownDescription: "Distance metric for this route",
										Computed:            true,
									},
									"permanent": schema.BoolAttribute{
										MarkdownDescription: "Permanent route",
										Computed:            true,
									},
									"track": schema.StringAttribute{
										MarkdownDescription: "Enable object tracking for static route",
										Computed:            true,
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: "Set metric for this route",
										Computed:            true,
									},
									"bfd_fast_detect_minimum_interval": schema.Int64Attribute{
										MarkdownDescription: "Hello interval",
										Computed:            true,
									},
									"bfd_fast_detect_multiplier": schema.Int64Attribute{
										MarkdownDescription: "Detect multiplier",
										Computed:            true,
									},
								},
							},
						},
						"nexthop_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "Forwarding router's address",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Forwarding router's address",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "description of the static route",
										Computed:            true,
									},
									"tag": schema.Int64Attribute{
										MarkdownDescription: "Set tag for this route",
										Computed:            true,
									},
									"distance_metric": schema.Int64Attribute{
										MarkdownDescription: "Distance metric for this route",
										Computed:            true,
									},
									"permanent": schema.BoolAttribute{
										MarkdownDescription: "Permanent route",
										Computed:            true,
									},
									"track": schema.StringAttribute{
										MarkdownDescription: "Enable object tracking for static route",
										Computed:            true,
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: "Set metric for this route",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *RouterStaticVRFIPv6MulticastDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterStaticVRFIPv6MulticastDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterStaticVRFIPv6MulticastData

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

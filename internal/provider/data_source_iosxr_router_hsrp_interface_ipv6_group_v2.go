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
	_ datasource.DataSource              = &RouterHSRPInterfaceIPv6GroupV2DataSource{}
	_ datasource.DataSourceWithConfigure = &RouterHSRPInterfaceIPv6GroupV2DataSource{}
)

func NewRouterHSRPInterfaceIPv6GroupV2DataSource() datasource.DataSource {
	return &RouterHSRPInterfaceIPv6GroupV2DataSource{}
}

type RouterHSRPInterfaceIPv6GroupV2DataSource struct {
	client *client.Client
}

func (d *RouterHSRPInterfaceIPv6GroupV2DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_hsrp_interface_ipv6_group_v2"
}

func (d *RouterHSRPInterfaceIPv6GroupV2DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router HSRP Interface IPv6 Group V2 configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "HSRP interface configuration subcommands",
				Required:            true,
			},
			"group_id": schema.Int64Attribute{
				MarkdownDescription: "group number version 2",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "MGO session name",
				Computed:            true,
			},
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "Use specified mac address for the virtual router",
				Computed:            true,
			},
			"timers_hold_time": schema.Int64Attribute{
				MarkdownDescription: "Hold time in seconds",
				Computed:            true,
			},
			"timers_hold_time2": schema.Int64Attribute{
				MarkdownDescription: "Hold time in seconds",
				Computed:            true,
			},
			"timers_msec": schema.Int64Attribute{
				MarkdownDescription: "Specify hellotime in milliseconds",
				Computed:            true,
			},
			"timers_msec2": schema.Int64Attribute{
				MarkdownDescription: "Specify hold time in milliseconds",
				Computed:            true,
			},
			"preempt_delay": schema.Int64Attribute{
				MarkdownDescription: "Wait before preempting",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Priority level",
				Computed:            true,
			},
			"bfd_fast_detect_peer_ipv6": schema.StringAttribute{
				MarkdownDescription: "BFD peer interface IPv6 address",
				Computed:            true,
			},
			"bfd_fast_detect_peer_interface": schema.StringAttribute{
				MarkdownDescription: "Select an interface over which to run BFD",
				Computed:            true,
			},
			"track_objects": schema.ListNestedAttribute{
				MarkdownDescription: "Object tracking",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object_name": schema.StringAttribute{
							MarkdownDescription: "Object tracking",
							Computed:            true,
						},
						"priority_decrement": schema.Int64Attribute{
							MarkdownDescription: "Priority decrement",
							Computed:            true,
						},
					},
				},
			},
			"track_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Configure tracking",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"track_name": schema.StringAttribute{
							MarkdownDescription: "Configure tracking",
							Computed:            true,
						},
						"priority_decrement": schema.Int64Attribute{
							MarkdownDescription: "Priority decrement",
							Computed:            true,
						},
					},
				},
			},
			"addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Global HSRP IPv6 address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "Set Global HSRP IPv6 address",
							Computed:            true,
						},
					},
				},
			},
			"address_link_local_autoconfig": schema.BoolAttribute{
				MarkdownDescription: "Autoconfigure the HSRP IPv6 linklocal address",
				Computed:            true,
			},
			"address_link_local_autoconfig_legacy_compatible": schema.BoolAttribute{
				MarkdownDescription: "Autoconfigure for Legacy compatibility (with IOS/NX-OS)",
				Computed:            true,
			},
			"address_link_local_ipv6_address": schema.StringAttribute{
				MarkdownDescription: "HSRP IPv6 linklocal address",
				Computed:            true,
			},
		},
	}
}

func (d *RouterHSRPInterfaceIPv6GroupV2DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterHSRPInterfaceIPv6GroupV2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterHSRPInterfaceIPv6GroupV2Data

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

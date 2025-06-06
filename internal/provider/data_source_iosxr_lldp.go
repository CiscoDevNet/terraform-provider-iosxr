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
	_ datasource.DataSource              = &LLDPDataSource{}
	_ datasource.DataSourceWithConfigure = &LLDPDataSource{}
)

func NewLLDPDataSource() datasource.DataSource {
	return &LLDPDataSource{}
}

type LLDPDataSource struct {
	client *client.Client
}

func (d *LLDPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lldp"
}

func (d *LLDPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the LLDP configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"holdtime": schema.Int64Attribute{
				MarkdownDescription: "Specify the holdtime (in sec) to be sent in packets",
				Computed:            true,
			},
			"timer": schema.Int64Attribute{
				MarkdownDescription: "Specify the rate at which LLDP packets are sent (in sec)",
				Computed:            true,
			},
			"reinit": schema.Int64Attribute{
				MarkdownDescription: "Delay (in sec) for LLDP initialization on any interface",
				Computed:            true,
			},
			"subinterfaces_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable LLDP over Sub-interfaces as well",
				Computed:            true,
			},
			"management_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable LLDP over Management interface as well",
				Computed:            true,
			},
			"priorityaddr_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable LLDP to use Management interface address first(if configured)",
				Computed:            true,
			},
			"extended_show_width_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable Extended Show LLDP Neighbor Width",
				Computed:            true,
			},
			"tlv_select_management_address_disable": schema.BoolAttribute{
				MarkdownDescription: "disable Management Address TLV",
				Computed:            true,
			},
			"tlv_select_port_description_disable": schema.BoolAttribute{
				MarkdownDescription: "disable Port Description TLV",
				Computed:            true,
			},
			"tlv_select_system_capabilities_disable": schema.BoolAttribute{
				MarkdownDescription: "disable System Capabilities TLV",
				Computed:            true,
			},
			"tlv_select_system_description_disable": schema.BoolAttribute{
				MarkdownDescription: "disable System Description TLV",
				Computed:            true,
			},
			"tlv_select_system_name_disable": schema.BoolAttribute{
				MarkdownDescription: "disable System Name TLV",
				Computed:            true,
			},
		},
	}
}

func (d *LLDPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *LLDPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LLDPData

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

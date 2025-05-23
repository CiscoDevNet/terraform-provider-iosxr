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
	_ datasource.DataSource              = &L2VPNPWClassDataSource{}
	_ datasource.DataSourceWithConfigure = &L2VPNPWClassDataSource{}
)

func NewL2VPNPWClassDataSource() datasource.DataSource {
	return &L2VPNPWClassDataSource{}
}

type L2VPNPWClassDataSource struct {
	client *client.Client
}

func (d *L2VPNPWClassDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2vpn_pw_class"
}

func (d *L2VPNPWClassDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the L2VPN PW Class configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Pseudowire class template",
				Required:            true,
			},
			"encapsulation_mpls": schema.BoolAttribute{
				MarkdownDescription: "Set pseudowire encapsulation to MPLS",
				Computed:            true,
			},
			"encapsulation_mpls_transport_mode_ethernet": schema.BoolAttribute{
				MarkdownDescription: "Ethernet port mode",
				Computed:            true,
			},
			"encapsulation_mpls_transport_mode_vlan": schema.BoolAttribute{
				MarkdownDescription: "Vlan tagged mode",
				Computed:            true,
			},
			"encapsulation_mpls_transport_mode_passthrough": schema.BoolAttribute{
				MarkdownDescription: "passthrough incoming tags",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_pw_label": schema.BoolAttribute{
				MarkdownDescription: "Enable PW VC label based load balancing",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_transmit": schema.BoolAttribute{
				MarkdownDescription: "Insert Flow label on transmit ",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_transmit_static": schema.BoolAttribute{
				MarkdownDescription: "Set Flow label parameters statically",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_receive": schema.BoolAttribute{
				MarkdownDescription: "Discard Flow label on receive",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_receive_static": schema.BoolAttribute{
				MarkdownDescription: "Set Flow label parameters statically",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_both": schema.BoolAttribute{
				MarkdownDescription: "Insert/Discard Flow label on transmit/recceive",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_both_static": schema.BoolAttribute{
				MarkdownDescription: "Set Flow label parameters statically",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_code_one7": schema.BoolAttribute{
				MarkdownDescription: "Legacy code value",
				Computed:            true,
			},
			"encapsulation_mpls_load_balancing_flow_label_code_one7_disable": schema.BoolAttribute{
				MarkdownDescription: "Disables sending code 17 TLV",
				Computed:            true,
			},
		},
	}
}

func (d *L2VPNPWClassDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *L2VPNPWClassDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config L2VPNPWClassData

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

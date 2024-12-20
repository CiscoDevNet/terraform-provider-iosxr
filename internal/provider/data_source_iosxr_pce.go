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
	_ datasource.DataSource              = &PCEDataSource{}
	_ datasource.DataSourceWithConfigure = &PCEDataSource{}
)

func NewPCEDataSource() datasource.DataSource {
	return &PCEDataSource{}
}

type PCEDataSource struct {
	client *client.Client
}

func (d *PCEDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pce"
}

func (d *PCEDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the PCE configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"address_ipv4": schema.StringAttribute{
				MarkdownDescription: "IPv4 address",
				Computed:            true,
			},
			"address_ipv6": schema.StringAttribute{
				MarkdownDescription: "IPv6 address",
				Computed:            true,
			},
			"state_sync_ipv4s": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv4 address",
							Computed:            true,
						},
					},
				},
			},
			"peer_filter_ipv4_access_list": schema.StringAttribute{
				MarkdownDescription: "Access-list for IPv4 peer filtering",
				Computed:            true,
			},
			"api_authentication_digest": schema.BoolAttribute{
				MarkdownDescription: "Use HTTP Digest authentication (MD5)",
				Computed:            true,
			},
			"api_sibling_ipv4": schema.StringAttribute{
				MarkdownDescription: "IPv4 address of the PCE sibling",
				Computed:            true,
			},
			"api_users": schema.ListNestedAttribute{
				MarkdownDescription: "Northbound API username",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"user_name": schema.StringAttribute{
							MarkdownDescription: "Northbound API username",
							Computed:            true,
						},
						"password_encrypted": schema.StringAttribute{
							MarkdownDescription: "Specify unencrypted password",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *PCEDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *PCEDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PCEData

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

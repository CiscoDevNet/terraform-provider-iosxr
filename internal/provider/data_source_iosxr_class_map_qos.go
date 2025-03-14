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
	_ datasource.DataSource              = &ClassMapQoSDataSource{}
	_ datasource.DataSourceWithConfigure = &ClassMapQoSDataSource{}
)

func NewClassMapQoSDataSource() datasource.DataSource {
	return &ClassMapQoSDataSource{}
}

type ClassMapQoSDataSource struct {
	client *client.Client
}

func (d *ClassMapQoSDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_class_map_qos"
}

func (d *ClassMapQoSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Class Map QoS configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"class_map_name": schema.StringAttribute{
				MarkdownDescription: "Name of the class-map",
				Required:            true,
			},
			"match_any": schema.BoolAttribute{
				MarkdownDescription: "Match any match criteria (default)",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Set description for this class-map",
				Computed:            true,
			},
			"match_dscp": schema.ListAttribute{
				MarkdownDescription: "DSCP value",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"match_mpls_experimental_topmost": schema.ListAttribute{
				MarkdownDescription: "MPLS experimental label",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"match_qos_group": schema.ListAttribute{
				MarkdownDescription: "QoS Group Id",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"match_traffic_class": schema.ListAttribute{
				MarkdownDescription: "Traffic Class Id",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *ClassMapQoSDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *ClassMapQoSDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ClassMapQoSData

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

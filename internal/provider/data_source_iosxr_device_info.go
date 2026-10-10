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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &DeviceInfoDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceInfoDataSource{}
)

func NewDeviceInfoDataSource() datasource.DataSource {
	return &DeviceInfoDataSource{}
}

type DeviceInfoDataSource struct {
	data *IosxrProviderData
}

type DeviceInfoData struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	Version types.String `tfsdk:"version"`
}

func (d *DeviceInfoDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_info"
}

func (d *DeviceInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Returns metadata about a managed device, including the auto-detected IOS-XR version. No additional gNMI connection is made — the version is read from the provider's internal cache.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The device name (same as `device`).",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "Auto-detected IOS-XR version in `major.minor` format (e.g. `24.4`, `25.4`). Empty string if version detection was skipped or unavailable.",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceInfoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxrProviderData)
}

func (d *DeviceInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceInfoData

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceName := config.Device.ValueString()
	device, ok := d.data.Devices[deviceName]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device",
			fmt.Sprintf("Device '%s' does not exist in provider configuration.", deviceName))
		return
	}

	// Use "default" as the id when no device name was specified (empty key = default device).
	id := deviceName
	if id == "" {
		id = "default"
	}

	tflog.Debug(ctx, fmt.Sprintf("iosxr_device_info: Reading device info for '%s'", id))

	state := DeviceInfoData{
		Device:  config.Device,
		Id:      types.StringValue(id),
		Version: types.StringValue(device.Version),
	}

	tflog.Debug(ctx, fmt.Sprintf("iosxr_device_info: device='%s' version='%s'", id, device.Version))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

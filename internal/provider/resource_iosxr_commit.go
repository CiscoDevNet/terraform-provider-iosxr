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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewCommitResource() resource.Resource {
	return &CommitResource{}
}

type CommitResource struct {
	data *IosxrProviderData
}

type Commit struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
	Commit types.Bool   `tfsdk:"commit"`
}

func (r *CommitResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_commit"
}

func (r *CommitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource commits all pending batch operations for a device. Use this with `depends_on` to control when batched operations are flushed to the device. Only needed when `auto_commit=false` (batch mode).",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Internal identifier.",
				Computed:            true,
			},
			"commit": schema.BoolAttribute{
				MarkdownDescription: "This attribute is only used internally to trigger commits on every apply.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

func (r *CommitResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.data = req.ProviderData.(*IosxrProviderData)
}

func (r *CommitResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Commit

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	deviceName := plan.Device.ValueString()
	if deviceName == "" {
		deviceName = "default"
	}

	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Beginning commit for device '%s'", deviceName))

	if device.Managed && !device.AutoCommit {
		ops := device.DrainCandidateOps()
		if len(ops) == 0 {
			tflog.Info(ctx, fmt.Sprintf("iosxr_commit: No pending operations for device '%s'", deviceName))
		} else {
			// Count operation types for logging
			updates, deletes, replaces := 0, 0, 0
			for _, op := range ops {
				switch op.OperationType {
				case "update":
					updates++
				case "delete":
					deletes++
				case "replace":
					replaces++
				}
			}

			tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Committing %d operation(s) for device '%s' (Updates: %d, Deletes: %d, Replaces: %d)",
				len(ops), deviceName, updates, deletes, replaces))

			if !r.data.ReuseConnection {
				defer device.Client.Disconnect()
			}
			_, err := device.Client.Set(ctx, ops)
			if err != nil {
				// Re-queue on failure
				device.AppendCandidateOps(ops)
				resp.Diagnostics.AddError("iosxr_commit: Failed to commit operations", err.Error())
				return
			}

			tflog.Info(ctx, fmt.Sprintf("iosxr_commit: SUCCESS - Committed %d operation(s) to device '%s'", len(ops), deviceName))
		}
	} else if device.AutoCommit {
		tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Auto-commit enabled for device '%s', nothing to commit", deviceName))
	}

	plan.Id = types.StringValue(fmt.Sprintf("commit-%s-%d", deviceName, time.Now().Unix()))

	tflog.Debug(ctx, fmt.Sprintf("iosxr_commit: Create finished successfully"))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Commit

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set commit to false - this causes terraform to see it as changed and trigger Update() on next apply
	// This is the same pattern used in iosxe_commit resource
	resp.State.SetAttribute(ctx, path.Root("commit"), false)
}

func (r *CommitResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan Commit

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	deviceName := plan.Device.ValueString()
	if deviceName == "" {
		deviceName = "default"
	}

	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Beginning commit for device '%s' (UPDATE)", deviceName))

	if device.Managed && !device.AutoCommit {
		ops := device.DrainCandidateOps()
		if len(ops) == 0 {
			tflog.Info(ctx, fmt.Sprintf("iosxr_commit: No pending operations for device '%s'", deviceName))
		} else {
			// Count operation types for logging
			updates, deletes, replaces := 0, 0, 0
			for _, op := range ops {
				switch op.OperationType {
				case "update":
					updates++
				case "delete":
					deletes++
				case "replace":
					replaces++
				}
			}

			tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Committing %d operation(s) for device '%s' (Updates: %d, Deletes: %d, Replaces: %d)",
				len(ops), deviceName, updates, deletes, replaces))

			if !r.data.ReuseConnection {
				defer device.Client.Disconnect()
			}
			_, err := device.Client.Set(ctx, ops)
			if err != nil {
				// Re-queue on failure
				device.AppendCandidateOps(ops)
				resp.Diagnostics.AddError("iosxr_commit: Failed to commit operations", err.Error())
				return
			}

			tflog.Info(ctx, fmt.Sprintf("iosxr_commit: SUCCESS - Committed %d operation(s) to device '%s'", len(ops), deviceName))
		}
	} else if device.AutoCommit {
		tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Auto-commit enabled for device '%s', nothing to commit", deviceName))
	}

	// Set ID with timestamp to track when commit was executed
	plan.Id = types.StringValue(fmt.Sprintf("commit-%s-%d", deviceName, time.Now().Unix()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Commit resource doesn't need to do anything on delete
	tflog.Debug(ctx, "iosxr_commit: Delete called (no-op)")
	resp.State.RemoveResource(ctx)
}

func (r *CommitResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError(
		"Import not supported",
		"The iosxr_commit resource does not support import. It's a trigger resource that commits pending batch operations.",
	)
}

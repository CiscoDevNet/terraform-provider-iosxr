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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
}

func (r *CommitResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_commit"
}

func (r *CommitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Commits all pending batch operations for a device in a single atomic gNMI Set call. " +
			"Use `depends_on` to ensure all resources are staged before committing. " +
			"Only needed when `auto_commit=false` (batch mode). " +
			"This resource always re-applies on every terraform apply (NETCONF PR #332 pattern). " +
			"On destroy, each resource commits its own delete immediately — this resource is a no-op on destroy.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of this resource.",
				Computed:            true,
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

// commitBatch drains the device CandidateStore and flushes all staged gNMI operations
// in a single atomic Set call. This is the gNMI equivalent of NETCONF's CommitBatch().
// Returns an error if the gNMI Set fails; ops are re-queued on failure.
func (r *CommitResource) commitBatch(ctx context.Context, device *IosxrProviderDataDevice, deviceName string) error {
	if device.AutoCommit {
		tflog.Info(ctx, fmt.Sprintf("iosxr_commit: auto_commit=true for device '%s', nothing to batch commit", deviceName))
		return nil
	}
	if !device.Managed {
		tflog.Info(ctx, fmt.Sprintf("iosxr_commit: device '%s' is not managed, skipping", deviceName))
		return nil
	}

	ops := device.DrainCandidateOps()
	if len(ops) == 0 {
		tflog.Info(ctx, fmt.Sprintf("iosxr_commit: No pending operations for device '%s'", deviceName))
		return nil
	}

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
	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Flushing %d staged operation(s) for device '%s' (updates=%d, deletes=%d, replaces=%d)",
		len(ops), deviceName, updates, deletes, replaces))

	if !r.data.ReuseConnection {
		defer func() { _ = device.Client.Disconnect() }()
	}

	_, err := device.Client.Set(ctx, ops)
	if err != nil {
		device.AppendCandidateOps(ops) // re-queue on failure so next apply can retry
		return fmt.Errorf("gNMI Set failed: %w", err)
	}

	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: SUCCESS - %d operation(s) committed to device '%s'", len(ops), deviceName))
	return nil
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
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device",
			fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	deviceName := plan.Device.ValueString()
	if deviceName == "" {
		deviceName = "default"
	}

	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Create - device '%s'", deviceName))

	if err := r.commitBatch(ctx, device, deviceName); err != nil {
		resp.Diagnostics.AddError("iosxr_commit: Batch commit failed", err.Error())
		return
	}

	plan.Id = types.StringValue(fmt.Sprintf("commit-%s", deviceName))
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Always remove this resource from state so Terraform plans a Create on every apply.
	// Create() calls commitBatch() which flushes all staged gNMI operations (no-op if
	// nothing is staged). This is the same mechanism as NETCONF PR #332: no lifecycle
	// blocks or extra trigger attributes needed.
	tflog.Debug(ctx, "iosxr_commit: Read - removing from state to force re-apply on next plan (PR #332 pattern)")
	resp.State.RemoveResource(ctx)
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
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device",
			fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	deviceName := plan.Device.ValueString()
	if deviceName == "" {
		deviceName = "default"
	}

	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Update - device '%s'", deviceName))

	if err := r.commitBatch(ctx, device, deviceName); err != nil {
		resp.Diagnostics.AddError("iosxr_commit: Batch commit failed", err.Error())
		return
	}

	plan.Id = types.StringValue(fmt.Sprintf("commit-%s", deviceName))
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No-op: during destroy, each resource commits its own delete operations immediately
	// (auto_commit setting is ignored during destroy — resources always self-commit on delete).
	// This matches the NETCONF PR #332 destroy behavior.
	tflog.Debug(ctx, "iosxr_commit: Delete (no-op - resources handle their own destroy commits)")
	resp.State.RemoveResource(ctx)
}

func (r *CommitResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("iosxr_commit: ImportState id='%s'", req.ID))
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

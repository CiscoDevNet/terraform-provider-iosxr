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

// commitBatch drains the device CandidateStore and flushes all staged gNMI
// operations in a single atomic Set call. Returns an error if the gNMI Set
// fails; ops are re-queued on failure so a subsequent apply can retry them.
//
// This is shared by the `iosxr_commit` action (action_iosxr_commit.go),
// which is the only thing that actually invokes it.
func commitBatch(ctx context.Context, data *IosxrProviderData, device *IosxrProviderDataDevice, deviceName string) error {
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

	if !data.ReuseConnection {
		defer func() { _ = device.GnmiClient.Disconnect() }()
	}

	_, err := device.GnmiClient.Set(ctx, ops)
	if err != nil {
		device.AppendCandidateOps(ops) // re-queue on failure so next apply can retry
		return fmt.Errorf("gNMI Set failed: %w", err)
	}

	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: SUCCESS - %d operation(s) committed to device '%s'", len(ops), deviceName))
	return nil
}

// NewCommitTriggerResource registers the `iosxr_commit_trigger` resource. It
// has no gNMI side effects of its own -- its only job is to plan an
// in-place update whenever its `triggers` attribute changes, which reliably
// fires a `lifecycle { action_trigger { events = [after_update], actions =
// [action.iosxr_commit...] } }` block attached to it, invoking the real
// `iosxr_commit` action (see action_iosxr_commit.go).
//
// Set `triggers` to an expression that changes whenever the batched
// resources' staged changes should be flushed, e.g.:
//
//	triggers = {
//	  hostname = jsonencode(iosxr_hostname.example)
//	  domain   = jsonencode(iosxr_domain.example)
//	}
//
// This is the same "triggers" idiom used by null_resource/terraform_data:
// Terraform's own plan diffing (comparing `triggers` against prior state)
// determines whether this resource -- and therefore the action -- needs to
// run, so a `terraform apply` with nothing staged shows NO changes at all
// (unlike an earlier design that used ModifyPlan to force an update on
// every single apply, which always showed a noisy in-place "update" even
// when nothing was staged).
//
// Use `depends_on` to list every resource whose staged operations should be
// included in that flush, and make sure this resource appears after them in
// the dependency graph so the action only fires once everything is staged
// (true batching), rather than attaching action_trigger to each staged
// resource individually (which could fire prematurely/multiple times if
// Terraform applies independent resources concurrently).
//
// On `terraform destroy` (or simply removing this block from configuration),
// Delete() is a no-op -- there is nothing to flush on the Create/Update
// path at that point. Use a second instance of this SAME resource type with
// `events = [before_destroy]` (or after_destroy) and the OPPOSITE
// `depends_on` direction to flush staged Delete operations; see the
// `iosxr_commit_trigger.on_destroy` example in main.tf.
func NewCommitTriggerResource() resource.Resource {
	return &CommitTriggerResource{}
}

type CommitTriggerResource struct{}

type CommitTrigger struct {
	Device   types.String  `tfsdk:"device"`
	Id       types.String  `tfsdk:"id"`
	Triggers types.Dynamic `tfsdk:"triggers"`
}

func (r *CommitTriggerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_commit_trigger"
}

func (r *CommitTriggerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Plans an in-place update whenever `triggers` changes, so it can be used to " +
			"reliably invoke the `iosxr_commit` action via a `lifecycle { action_trigger { ... } }` block only " +
			"when there is actually something to flush (unlike an always-updates-every-apply design, this " +
			"produces a clean 'no changes' plan when nothing is staged). Set `triggers` to a value derived " +
			"from every resource whose staged (`auto_commit=false`) Create/Update operations should be " +
			"flushed together in one atomic gNMI Set call, e.g. `jsonencode(iosxr_hostname.example)`. This " +
			"resource has no gNMI side effects of its own and is a no-op on destroy.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of this resource.",
				Computed:            true,
			},
			"triggers": schema.DynamicAttribute{
				MarkdownDescription: "An arbitrary value (string, map, object, etc.) that should change " +
					"whenever the batched resources' staged changes need to be flushed. A common pattern is " +
					"`jsonencode(...)` of each batched resource, e.g. `{ hostname = " +
					"jsonencode(iosxr_hostname.example) }`.",
				Optional: true,
			},
		},
	}
}

func (r *CommitTriggerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CommitTrigger

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceName := plan.Device.ValueString()
	if deviceName == "" {
		deviceName = "default"
	}

	plan.Id = types.StringValue(fmt.Sprintf("commit-trigger-%s", deviceName))

	tflog.Debug(ctx, fmt.Sprintf("iosxr_commit_trigger: Create - device '%s'", deviceName))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitTriggerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CommitTrigger

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitTriggerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CommitTrigger

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceName := plan.Device.ValueString()
	if deviceName == "" {
		deviceName = "default"
	}

	// id is Computed and deterministic (derived only from device), but must
	// still be explicitly (re-)set here: whenever "id" itself shows up as
	// unknown in the plan (e.g. after a schema change adds/removes other
	// attributes), Update() is responsible for resolving every unknown
	// computed value -- leaving it unset would make Terraform report
	// "Provider returned invalid result object after apply".
	plan.Id = types.StringValue(fmt.Sprintf("commit-trigger-%s", deviceName))

	tflog.Debug(ctx, fmt.Sprintf("iosxr_commit_trigger: Update - device '%s' (triggers changed, firing action_trigger)", deviceName))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitTriggerResource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// No-op: this resource has no gNMI side effects. On destroy, each
	// managed resource commits its own delete immediately (auto_commit is
	// ignored during destroy -- see gen/templates/resource.go's Delete
	// template), completely independent of this resource and the
	// `iosxr_commit` action.
	tflog.Debug(ctx, "iosxr_commit_trigger: Delete (no-op)")
}

func (r *CommitTriggerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("iosxr_commit_trigger: ImportState id='%s'", req.ID))
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

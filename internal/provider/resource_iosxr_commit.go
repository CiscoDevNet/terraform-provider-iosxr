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

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CommitResource{}
var _ resource.ResourceWithImportState = &CommitResource{}

func NewCommitResource() resource.Resource {
	return &CommitResource{}
}

// CommitResource defines the resource implementation.
type CommitResource struct {
	data *IosxrProviderData
}

type Commit struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
}

func (r *CommitResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_commit"
}

func (r *CommitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Commits all pending batch operations for a device. " +
			"Use `depends_on` to ensure all resources are staged before committing. " +
			"Only needed when `auto_commit=false` (batch mode). " +
			"This resource always re-applies on every terraform apply (PR #329 pattern). " +
			"The provider automatically sets `commit = true` — no user input required.",

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

func (r *CommitResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
		resp.Diagnostics.AddAttributeError(
			path.Root("device"),
			"Invalid device",
			fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()),
		)
		return
	}

	if device.Protocol == "netconf" {
		// In batch mode, check whether any dependent resource actually staged changes
		// by inspecting the batch locks. Locks are acquired by EditConfigBatch() when
		// resources stage NETCONF edits during apply. Because iosxr_commit has depends_on
		// on all those resources, by the time Create() runs here they have all finished
		// — so IsBatchLocked() is a reliable signal that there is something to commit.
		if !device.AutoCommit && !device.IsBatchLocked() {
			tflog.Info(ctx, "iosxr_commit: batch mode, no staged changes (batch locks not held) — skipping commit")
		} else {
			// Locks are held (staged changes exist) OR auto-commit mode: proceed with NETCONF commit
			locked := helpers.AcquireNetconfLock(device.GetOpMutex(), device.ReuseConnection, true)
			defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
			if locked {
				defer device.GetOpMutex().Unlock()
			}

			if err := helpers.EnsureNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection, device.MaxRetries); err != nil {
				resp.Diagnostics.AddError(
					"NETCONF Connection Error",
					fmt.Sprintf("Failed to ensure connection: %s", err),
				)
				return
			}

			if !device.AutoCommit {
				// Batch mode: commit staged edits and release locks
				tflog.Info(ctx, "iosxr_commit: batch locks held — committing staged changes")
				if err := helpers.CommitBatch(ctx, device.NetconfClient); err != nil {
					if helpers.IsNoChangesToCommitError(err) {
						tflog.Info(ctx, "iosxr_commit: no staged changes to commit (no-op)")
					} else {
						resp.Diagnostics.AddError(
							"Commit Failed",
							fmt.Sprintf("Failed to commit configuration in batch mode: %s", err),
						)
						return
					}
				}
				if err := helpers.ReleaseBatchLocks(ctx, device.NetconfClient, device); err != nil {
					tflog.Warn(ctx, fmt.Sprintf("Failed to release batch locks after commit: %s", err))
				}
				tflog.Info(ctx, "iosxr_commit: batch commit completed, locks released")
			} else {
				// Auto-commit mode: confirmed-commit or regular commit
				tflog.Info(ctx, fmt.Sprintf("Commit check: ConfirmedCommit=%v, AutoCommit=%v, Timeout=%d", device.ConfirmedCommit, device.AutoCommit, device.ConfirmedCommitTimeout))
				if device.ConfirmedCommit {
					tflog.Info(ctx, fmt.Sprintf("Executing confirmed-commit with %d second timeout", device.ConfirmedCommitTimeout))
					if err := helpers.CommitConfirmed(ctx, device.NetconfClient, device.ConfirmedCommitTimeout); err != nil {
						resp.Diagnostics.AddError(
							"Confirmed Commit Failed",
							fmt.Sprintf("Failed to execute confirmed-commit: %s", err),
						)
						return
					}
					tflog.Info(ctx, "Confirmed-commit completed successfully with auto-confirmation")
				} else {
					tflog.Info(ctx, "Executing regular NETCONF commit")
					if err := helpers.Commit(ctx, device.NetconfClient); err != nil {
						resp.Diagnostics.AddError(
							"Commit Failed",
							fmt.Sprintf("Failed to commit configuration: %s", err),
						)
						return
					}
					tflog.Info(ctx, "NETCONF commit completed successfully")
				}
			}
		}
	} else {
		// gNMI protocol — no-op
		tflog.Info(ctx, "gNMI protocol detected - commit operation skipped (gNMI is auto-commit)")
	}

	plan.Id = types.StringValue("commit")

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "iosxr_commit: Read - removing state to force re-create on next plan (PR #329 pattern)")
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
		resp.Diagnostics.AddAttributeError(
			path.Root("device"),
			"Invalid device",
			fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()),
		)
		return
	}

	if device.Protocol == "netconf" {
		// Lock for NETCONF operations
		locked := helpers.AcquireNetconfLock(device.GetOpMutex(), device.ReuseConnection, true)
		defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
		if locked {
			defer device.GetOpMutex().Unlock()
		}

		if err := helpers.EnsureNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection, device.MaxRetries); err != nil {
			resp.Diagnostics.AddError(
				"NETCONF Connection Error",
				fmt.Sprintf("Failed to ensure connection: %s", err),
			)
			return
		}

		// Execute confirmed-commit if enabled, otherwise regular commit
		tflog.Info(ctx, fmt.Sprintf("Commit check: ConfirmedCommit=%v, AutoCommit=%v, Timeout=%d", device.ConfirmedCommit, device.AutoCommit, device.ConfirmedCommitTimeout))
		if device.ConfirmedCommit {
			tflog.Info(ctx, fmt.Sprintf("Executing confirmed-commit with %d second timeout", device.ConfirmedCommitTimeout))
			if err := helpers.CommitConfirmed(ctx, device.NetconfClient, device.ConfirmedCommitTimeout); err != nil {
				resp.Diagnostics.AddError(
					"Confirmed Commit Failed",
					fmt.Sprintf("Failed to execute confirmed-commit: %s", err),
				)
				return
			}
			tflog.Info(ctx, "Confirmed-commit completed successfully with auto-confirmation")
		} else {
			// Regular commit operation
			tflog.Info(ctx, "Executing regular NETCONF commit")
			if err := helpers.Commit(ctx, device.NetconfClient); err != nil {
				resp.Diagnostics.AddError(
					"Commit Failed",
					fmt.Sprintf("Failed to commit configuration: %s", err),
				)
				return
			}
			tflog.Info(ctx, "NETCONF commit completed successfully")
		}
	} else {
		// gNMI protocol - no-op (log only)
		tflog.Info(ctx, "gNMI protocol detected - commit operation skipped (gNMI is auto-commit)")
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CommitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Commit

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(
			path.Root("device"),
			"Invalid device",
			fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()),
		)
		return
	}

	if device.Protocol == "netconf" {
		// Lock for NETCONF operations
		locked := helpers.AcquireNetconfLock(device.GetOpMutex(), device.ReuseConnection, true)
		defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
		if locked {
			defer device.GetOpMutex().Unlock()
		}

		if err := helpers.EnsureNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection, device.MaxRetries); err != nil {
			resp.Diagnostics.AddError(
				"NETCONF Connection Error",
				fmt.Sprintf("Failed to ensure connection: %s", err),
			)
			return
		}

		// During destroy, commit whatever is in the candidate datastore (batched deletes)
		// Don't use confirmed-commit during destroy - just regular commit
		tflog.Info(ctx, "Executing NETCONF commit during destroy (regular commit, no confirmed-commit)")
		if err := helpers.Commit(ctx, device.NetconfClient); err != nil {
			resp.Diagnostics.AddError(
				"Commit Failed During Destroy",
				fmt.Sprintf("Failed to commit configuration during destroy: %s", err),
			)
			return
		}
		tflog.Info(ctx, "NETCONF commit during destroy completed successfully")
	} else {
		// gNMI protocol - no-op (log only)
		tflog.Info(ctx, "gNMI protocol detected - commit operation skipped during destroy (gNMI is auto-commit)")
	}
}

func (r *CommitResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

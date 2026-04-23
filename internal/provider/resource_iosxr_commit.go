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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	Id              types.String `tfsdk:"id"`
	Device          types.String `tfsdk:"device"`
	CommitOnDestroy types.Bool   `tfsdk:"commit_on_destroy"`
}

func (r *CommitResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_commit"
}

func (r *CommitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `This resource executes a NETCONF commit operation to persist all pending configuration changes.

## Behavior
- **NETCONF Protocol**: Performs a commit operation to copy candidate datastore to running datastore
- **gNMI Protocol**: No operation (gNMI is auto-commit by design)
- **Confirmed Commit Mode**: When ` + "`confirmed_commit=true`" + ` in provider configuration:
  - Executes confirmed-commit with the configured timeout (60-240 seconds)
  - Auto-confirms on success for seamless Terraform workflow
  - IOS XR automatically rolls back if timeout expires without confirmation
  - Requires :confirmed-commit:1.1 capability on the device

## Use Cases
1. **Batch Operations**: Group multiple resource changes into one transaction
2. **Atomic Deployments**: Ensure all-or-nothing configuration changes
3. **Confirmed Commit**: Safe configuration deployment with automatic rollback capability
4. **Explicit Commit**: Control exactly when changes are committed to running datastore

## Example Usage

### Basic Commit (Auto-Commit)
` + "```hcl" + `
provider "iosxr" {
  auto_commit = true  # Individual resources auto-commit
}

resource "iosxr_hostname" "example" {
  system_network_name = "router1"
}
# No iosxr_commit needed - changes committed immediately
` + "```" + `

### Manual Batch Commit
` + "```hcl" + `
provider "iosxr" {
  auto_commit = false  # Stage changes without committing
}

resource "iosxr_hostname" "example" {
  system_network_name = "router1"
}

resource "iosxr_logging" "example" {
  ipv4_dscp = "af11"
}

# Commit all changes in one transaction
resource "iosxr_commit" "batch" {
  depends_on = [
    iosxr_hostname.example,
    iosxr_logging.example,
  ]

  lifecycle {
    replace_triggered_by = [
      iosxr_hostname.example,
      iosxr_logging.example,
    ]
  }
}
` + "```" + `

### Confirmed Commit with Automatic Rollback
` + "```hcl" + `
provider "iosxr" {
  protocol                   = "netconf"
  auto_commit                = true
  confirmed_commit           = true
  confirmed_commit_timeout   = 120  # 2 minutes timeout
}

resource "iosxr_hostname" "example" {
  system_network_name = "router1"
}

resource "iosxr_logging" "example" {
  ipv4_dscp = "af11"
}

# Executes confirmed-commit with automatic confirmation on success
resource "iosxr_commit" "safe_deploy" {
  depends_on = [
    iosxr_hostname.example,
    iosxr_logging.example,
  ]

  lifecycle {
    replace_triggered_by = [
      iosxr_hostname.example,
      iosxr_logging.example,
    ]
  }
}
# If commit succeeds, changes are auto-confirmed
# If timeout expires, IOS XR automatically rolls back all changes
` + "```" + `

## Notes
- This resource is primarily for NETCONF protocol users
- For gNMI protocol, this resource has no effect (gNMI is auto-commit)
- Use ` + "`depends_on`" + ` to ensure proper ordering of configuration changes
- Use ` + "`replace_triggered_by`" + ` lifecycle meta-argument to force re-execution when dependencies change
- The resource ID is a static timestamp of when the commit was executed
- Confirmed commit requires IOS XR support for :confirmed-commit:1.1 capability
`,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the commit operation (timestamp).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"device": schema.StringAttribute{
				MarkdownDescription: "Device name. This corresponds to the name of a device configured in the provider.",
				Optional:            true,
			},
			"commit_on_destroy": schema.BoolAttribute{
				MarkdownDescription: "Whether to commit pending changes when this resource is destroyed. Default is `true`. Set to `false` for commit resources that should only commit during create/update but not during destroy.",
				Optional:            true,
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

	plan.Id = types.StringValue("commit")

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

	// Commit resource is stateless - just preserve the existing state
	tflog.Debug(ctx, "Commit resource read - preserving existing state")

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
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

	// DESTROY: Do nothing - each resource commits individually during destroy
	// User requirement: "for destroy, don't batch the requests"
	// Resources ignore auto_commit during destroy and commit immediately
	tflog.Info(ctx, "Commit resource delete - skipping commit (resources commit individually during destroy)")
}

func (r *CommitResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

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
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NewCommitAction registers the `iosxr_commit` action, which flushes a
// device's staged (auto_commit=false) Create/Update operations in a single
// atomic gNMI Set call. This replaces the earlier `iosxr_commit` RESOURCE
// design, which relied on Read() always calling RemoveResource() to force
// Terraform to re-run Create() on every apply -- a trick that only works
// when Terraform is willing to plan a Create/Update, which never happens
// during a pure `terraform destroy` plan (delete-only, for every resource).
//
// Actions have no state and are invoked explicitly via a resource's
// `lifecycle { action_trigger { ... } }` block (see CommitTriggerResource in
// resource_iosxr_commit.go), so this logic is completely decoupled from any
// resource's CRUD/state semantics. Per-resource Delete() handlers already
// commit their own deletes immediately regardless of auto_commit, so a full
// destroy is entirely unaffected by (and does not depend on) this action.
func NewCommitAction() action.Action {
	return &CommitAction{}
}

type CommitAction struct {
	data *IosxrProviderData
}

type CommitActionModel struct {
	Device types.String `tfsdk:"device"`
}

func (a *CommitAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_commit"
}

func (a *CommitAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Commits pending staged (auto_commit=false) Create/Update operations in one " +
			"atomic gNMI Set call per device. Invoke via a `lifecycle { action_trigger { ... } }` block on a " +
			"resource (see `iosxr_commit_trigger`) so it runs after all the resources whose staged changes " +
			"should be flushed together. If `device` is omitted, every device in the provider configuration " +
			"is flushed independently (each gets its own Set call; a single shared `iosxr_commit_trigger` " +
			"pair can therefore be reused across all devices). If `device` is set, only that device is " +
			"flushed. Delete operations always commit immediately regardless of `auto_commit` and are " +
			"unaffected by this action, including during `terraform destroy`.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration. Omit to flush every " +
					"configured device independently in a single invocation.",
				Optional: true,
			},
		},
	}
}

func (a *CommitAction) Configure(_ context.Context, req action.ConfigureRequest, _ *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	a.data = req.ProviderData.(*IosxrProviderData)
}

func (a *CommitAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CommitActionModel

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// device omitted entirely (null): flush every configured device, each
	// with its own gNMI Set call, so one shared action/trigger pair can
	// batch-commit across all devices while each device's changes still
	// land in a separate, independent atomic Set.
	if config.Device.IsNull() {
		a.invokeAll(ctx, resp)
		return
	}

	deviceKey := config.Device.ValueString()
	device, ok := a.data.Devices[deviceKey]
	if !ok {
		resp.Diagnostics.AddError("Invalid device",
			fmt.Sprintf("Device '%s' does not exist in provider configuration.", deviceKey))
		return
	}

	a.invokeOne(ctx, resp, deviceKey, device)
}

// invokeAll flushes every configured device's candidate store, one gNMI Set
// call per device. A failure on one device does not prevent the others from
// being flushed; all errors are collected and reported together.
func (a *CommitAction) invokeAll(ctx context.Context, resp *action.InvokeResponse) {
	deviceKeys := make([]string, 0, len(a.data.Devices))
	for key := range a.data.Devices {
		deviceKeys = append(deviceKeys, key)
	}
	sort.Strings(deviceKeys) // deterministic order across invocations

	var failed []string
	for _, deviceKey := range deviceKeys {
		device := a.data.Devices[deviceKey]
		deviceName := deviceKey
		if deviceName == "" {
			deviceName = "default"
		}

		tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Invoke - device '%s'", deviceName))
		if resp.SendProgress != nil {
			resp.SendProgress(action.InvokeProgressEvent{
				Message: fmt.Sprintf("Flushing staged operations for device '%s'", deviceName),
			})
		}

		if err := commitBatch(ctx, a.data, device, deviceName); err != nil {
			failed = append(failed, fmt.Sprintf("%s: %s", deviceName, err.Error()))
			continue
		}

		if resp.SendProgress != nil {
			resp.SendProgress(action.InvokeProgressEvent{
				Message: fmt.Sprintf("Committed staged operations for device '%s'", deviceName),
			})
		}
	}

	if len(failed) > 0 {
		resp.Diagnostics.AddError("iosxr_commit: Batch commit failed for one or more devices",
			strings.Join(failed, "; "))
	}
}

func (a *CommitAction) invokeOne(ctx context.Context, resp *action.InvokeResponse, deviceKey string, device *IosxrProviderDataDevice) {
	deviceName := deviceKey
	if deviceName == "" {
		deviceName = "default"
	}

	tflog.Info(ctx, fmt.Sprintf("iosxr_commit: Invoke - device '%s'", deviceName))

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("Flushing staged operations for device '%s'", deviceName),
		})
	}

	if err := commitBatch(ctx, a.data, device, deviceName); err != nil {
		resp.Diagnostics.AddError("iosxr_commit: Batch commit failed", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("Committed staged operations for device '%s'", deviceName),
		})
	}
}

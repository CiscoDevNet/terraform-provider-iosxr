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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-gnmi"
	"github.com/netascode/go-netconf"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &YangResource{}
var _ resource.ResourceWithImportState = &YangResource{}

func NewYangResource() resource.Resource {
	return &YangResource{}
}

type YangResource struct {
	data *IosxrProviderData
}

func (r *YangResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_yang"
}

func (r *YangResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages IOS-XR objects via YANG paths. This resource can only manage a single object. It is able to read the state and therefore reconcile configuration drift.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "A YANG path, e.g. `openconfig-interfaces:interfaces`.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"delete": schema.BoolAttribute{
				MarkdownDescription: "Delete object during destroy operation. Default value is `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"attributes": schema.MapAttribute{
				MarkdownDescription: "Map of key-value pairs which represents the YANG leafs and its values.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
			},
			"lists": schema.ListNestedAttribute{
				MarkdownDescription: "YANG lists.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "YANG list name.",
							Required:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: "YANG list key attribute. In case of multiple keys, those should be separated by a comma (`,`).",
							Optional:            true,
						},
						"items": schema.ListAttribute{
							MarkdownDescription: "List of maps of key-value pairs which represents the YANG leafs and its values.",
							Optional:            true,
							ElementType:         types.MapType{ElemType: types.StringType},
						},
						"values": schema.ListAttribute{
							MarkdownDescription: "YANG leaf-list values.",
							Optional:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *YangResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Yang

	diag := req.Config.Get(ctx, &data)

	if diag.HasError() {
		return
	}

	for l := range data.Lists {
		listName := data.Lists[l].Name.ValueString()
		hasItems := len(data.Lists[l].Items) > 0
		hasValues := len(data.Lists[l].Values.Elements()) > 0
		hasKey := !data.Lists[l].Key.IsNull() && data.Lists[l].Key.ValueString() != ""

		// Validate that either items or values is present, but not both
		if !hasItems && !hasValues {
			resp.Diagnostics.AddAttributeError(
				path.Root("lists"),
				"Invalid List Configuration",
				fmt.Sprintf("List '%s' must contain either 'items' or 'values'.", listName),
			)
			continue
		}

		if hasItems && hasValues {
			resp.Diagnostics.AddAttributeError(
				path.Root("lists"),
				"Invalid List Configuration",
				fmt.Sprintf("List '%s' cannot contain both 'items' and 'values'.", listName),
			)
			continue
		}

		// If items is used, key must be present
		if hasItems && !hasKey {
			resp.Diagnostics.AddAttributeError(
				path.Root("lists"),
				"Invalid List Configuration",
				fmt.Sprintf("List '%s' with 'items' must have a 'key' attribute.", listName),
			)
			continue
		}

		// If items is used, validate that all keys are present in each item
		if hasItems && hasKey {
			keyString := data.Lists[l].Key.ValueString()
			// Split the key by comma to handle composite keys (e.g., "name,type")
			keys := strings.Split(keyString, ",")

			for i := range data.Lists[l].Items {
				var m map[string]string
				data.Lists[l].Items[i].ElementsAs(ctx, &m, false)

				// Check that all keys are present in the item
				for _, key := range keys {
					key = strings.TrimSpace(key) // Trim whitespace
					if _, ok := m[key]; !ok {
						resp.Diagnostics.AddAttributeError(
							path.Root("lists"),
							"Invalid List Configuration",
							fmt.Sprintf("Key '%s' is missing in list item of list '%s'.", key, listName),
						)
					}
				}
			}
		}
	}
}

func (r *YangResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxrProviderData)
}

func (r *YangResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Yang

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Path.ValueString()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		if device.Protocol == "gnmi" {
			// Ensure connection is healthy (reconnect if stale)
			if err := helpers.EnsureGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection); err != nil {
				resp.Diagnostics.AddError("gNMI Connection Error", fmt.Sprintf("Failed to ensure connection: %s", err))
				return
			}

			// Ensure connection is closed when function exits (if reuse disabled)
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)

			body := plan.toBody(ctx)
			tflog.Debug(ctx, fmt.Sprintf("Yang toBody result: %s", body))
			// Only send if we have actual content
			if body != "" && body != "{}" {
				tflog.Debug(ctx, fmt.Sprintf("Sending gNMI Update: path=%s, body=%s", plan.Path.ValueString(), body))
				_, err := device.GnmiClient.Set(ctx, []gnmi.SetOperation{gnmi.Update(plan.Path.ValueString(), body)})
				if err != nil {
					resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
					return
				}
			}
		} else {
			// Serialize NETCONF operations when reuse disabled, or writes when reuse enabled
			locked := helpers.AcquireNetconfLock(&device.NetconfOpMutex, device.ReuseConnection, true)
			defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
			if locked {
				defer device.NetconfOpMutex.Unlock()
			}

			bodyStr := plan.toBodyXML(ctx)
			tflog.Debug(ctx, fmt.Sprintf("NETCONF CREATE: body=%s", bodyStr))

			if err := helpers.EditConfig(ctx, device.NetconfClient, bodyStr, true); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
	}

	plan.Id = plan.Path

	if plan.Attributes.IsUnknown() {
		plan.Attributes = types.MapNull(types.StringType)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Path.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *YangResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Yang

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Path.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed {
		if device.Protocol == "gnmi" {
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)
			getResp, err := device.GnmiClient.Get(ctx, []string{state.Path.ValueString()})
			if err != nil {
				if strings.Contains(err.Error(), "Requested element(s) not found") {
					// Resource not found - this can happen for resources that only have key attributes
					// The device returns "not found" because there's no actual config, just the container
					tflog.Debug(ctx, fmt.Sprintf("%s: Resource not found, preserving key attributes", state.Path.ValueString()))

					// For resources with only key attributes, we need to ensure those keys stay in state
					// even though the device returns "not found"
					// This is correct behavior - the keys are in the path, not in the config

					// Don't modify state.Attributes or state.Lists
					// They should already have the correct values from Create
					// Just skip the read since there's nothing to read
				} else {
					resp.Diagnostics.AddError("Unable to apply gNMI Get operation", err.Error())
					return
				}
			} else {
				respBody := getResp.Notifications[0].Update[0].Val.GetJsonIetfVal()
				state.fromBody(ctx, respBody)
			}
		} else {
			// Serialize NETCONF operations (all ops when reuse disabled, reads concurrent when reuse enabled)
			locked := helpers.AcquireNetconfLock(&device.NetconfOpMutex, device.ReuseConnection, false)
			defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
			if locked {
				defer device.NetconfOpMutex.Unlock()
			}

			filter := helpers.GetSubtreeFilter(state.Path.ValueString())
			res, err := device.NetconfClient.GetConfig(ctx, "running", filter)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object, got error: %s", err))
				return
			}

			state.fromBodyXML(ctx, res.Res)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Path.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *YangResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Yang

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Path.ValueString()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		if device.Protocol == "gnmi" {
			// Ensure connection is healthy (reconnect if stale)
			if err := helpers.EnsureGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection); err != nil {
				resp.Diagnostics.AddError("gNMI Connection Error", fmt.Sprintf("Failed to ensure connection: %s", err))
				return
			}

			// Ensure connection is closed when function exits (if reuse disabled)
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)

			var ops []gnmi.SetOperation
			body := plan.toBody(ctx)
			if body != "" && body != "{}" {
				ops = append(ops, gnmi.Update(plan.Path.ValueString(), body))
			}

			deletedItems := plan.getDeletedItems(ctx, state)
			tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedItems))

			for _, i := range deletedItems {
				ops = append(ops, gnmi.Delete(i))
			}

			if len(ops) > 0 {
				_, err := device.GnmiClient.Set(ctx, ops)
				if err != nil {
					resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
					return
				}
			}
		} else {
			// Serialize NETCONF operations when reuse disabled, or writes when reuse enabled
			locked := helpers.AcquireNetconfLock(&device.NetconfOpMutex, device.ReuseConnection, true)
			defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
			if locked {
				defer device.NetconfOpMutex.Unlock()
			}

			bodyStr := plan.toBodyXML(ctx)
			tflog.Debug(ctx, fmt.Sprintf("NETCONF UPDATE: body=%s", bodyStr))

			if err := helpers.EditConfig(ctx, device.NetconfClient, bodyStr, true); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
	}

	if plan.Attributes.IsUnknown() {
		plan.Attributes = types.MapNull(types.StringType)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Path.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *YangResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Yang

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Path.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed && state.Delete.ValueBool() {
		if device.Protocol == "gnmi" {
			// Ensure connection is healthy (reconnect if stale)
			if err := helpers.EnsureGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection); err != nil {
				resp.Diagnostics.AddError("gNMI Connection Error", fmt.Sprintf("Failed to ensure connection: %s", err))
				return
			}

			// Ensure connection is closed when function exits (if reuse disabled)
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)

			_, err := device.GnmiClient.Set(ctx, []gnmi.SetOperation{gnmi.Delete(state.Path.ValueString())})
			if err != nil && !strings.Contains(err.Error(), "Requested element(s) not found") {
				resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
				return
			}
		} else {
			// Serialize NETCONF operations when reuse disabled, or writes when reuse enabled
			locked := helpers.AcquireNetconfLock(&device.NetconfOpMutex, device.ReuseConnection, true)
			defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
			if locked {
				defer device.NetconfOpMutex.Unlock()
			}

			body := netconf.Body{}
			body = helpers.RemoveFromXPath(body, state.Path.ValueString())
			bodyStr := body.Res()
			tflog.Debug(ctx, fmt.Sprintf("NETCONF DELETE: body=%s", bodyStr))

			if err := helpers.EditConfig(ctx, device.NetconfClient, bodyStr, true); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Path.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *YangResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", req.ID))

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("path"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", req.ID))
}

type resourceGnmiType struct{}

var _ resource.Resource = (*GnmiResource)(nil)

func NewGnmiResource() resource.Resource {
	return &GnmiResource{}
}

type GnmiResource struct {
	data *IosxrProviderData
}

func (r *GnmiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gnmi"
}

func (r *GnmiResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages IOS-XR objects via gNMI calls. This resource can only manage a single object. It is able to read the state and therefore reconcile configuration drift.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "A gNMI path, e.g. `openconfig-interfaces:/interfaces/interface`.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"delete": schema.BoolAttribute{
				MarkdownDescription: "Delete object during destroy operation. Default value is `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"attributes": schema.MapAttribute{
				MarkdownDescription: "Map of key-value pairs which represents the attributes and its values. To indicate an empty YANG container use `<EMPTY>` as the value. To omit an attribute entirely (null value) use `<NULL>` as the value.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"lists": schema.ListNestedAttribute{
				MarkdownDescription: "YANG lists.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "YANG list name.",
							Required:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: "YANG list key attribute(s). In case of multiple keys, those should be separated by a comma (`,`).",
							Optional:            true,
						},
						"items": schema.ListAttribute{
							MarkdownDescription: "List of maps of key-value pairs which represents the attributes and its values. To indicate an empty YANG container use `<EMPTY>` as the value. To omit an attribute entirely (null value) use `<NULL>` as the value.",
							Optional:            true,
							ElementType:         types.MapType{ElemType: types.StringType},
						},
						"values": schema.ListAttribute{
							MarkdownDescription: "YANG leaf-list values.",
							Optional:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *GnmiResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Yang

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	for _, l := range data.Lists {
		hasValues := len(l.Values.Elements()) > 0
		hasItems := len(l.Items) > 0

		if !hasValues && !hasItems {
			resp.Diagnostics.AddAttributeError(
				path.Root("lists"),
				"Invalid List Configuration",
				"List must contain either items or values.",
			)
		}
		if hasValues && hasItems {
			resp.Diagnostics.AddAttributeError(
				path.Root("lists"),
				"Invalid List Configuration",
				"List cannot contain both items and values.",
			)
		}
		if hasItems && l.Key.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("lists"),
				"Invalid List Configuration",
				"List must contain a key.",
			)
		}
	}
}

func (r *GnmiResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxrProviderData)
}

func (r *GnmiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Yang

	// Read plan
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	if device.Managed {
		body := plan.toBody(ctx)

		// Only send Set operation if we have actual data to configure
		// Skip if body is empty or contains only empty JSON object
		shouldSend := body != "" && body != "{}"

		if shouldSend {
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)
			_, err := device.GnmiClient.Set(ctx, []gnmi.SetOperation{gnmi.Update(plan.Path.ValueString(), body)})
			if err != nil {
				resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
				return
			}
		}
	}

	plan.Id = plan.Path

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *GnmiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Yang

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	// Only perform Get operation if we have attributes or lists to read
	if device.Managed && (!state.Attributes.IsNull() || len(state.Lists) > 0) {
		defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)
		getResp, err := device.GnmiClient.Get(ctx, []string{state.Path.ValueString()})
		if err != nil {
			if strings.Contains(err.Error(), "Requested element(s) not found") {
				resp.State.RemoveResource(ctx)
				return
			} else {
				resp.Diagnostics.AddError("Unable to apply gNMI Get operation", err.Error())
				return
			}
		}

		state.fromBody(ctx, getResp.Notifications[0].Update[0].Val.GetJsonIetfVal())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *GnmiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Yang

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	if device.Managed {
		var ops []gnmi.SetOperation

		if !plan.Attributes.IsUnknown() {
			body := plan.toBody(ctx)
			// Only add update operation if we have actual content (not empty or just {})
			if body != "" && body != "{}" {
				ops = append(ops, gnmi.Update(plan.Path.ValueString(), body))
			}
		}

		deletedListItems := plan.getDeletedItems(ctx, state)
		tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

		for _, i := range deletedListItems {
			ops = append(ops, gnmi.Delete(i))
		}

		// Only execute Set if we have operations to perform
		if len(ops) > 0 {
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)
			_, err := device.GnmiClient.Set(ctx, ops)
			if err != nil {
				resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *GnmiResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Yang

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	if device.Managed {
		if state.Delete.ValueBool() {
			defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)
			_, err := device.GnmiClient.Set(ctx, []gnmi.SetOperation{gnmi.Delete(state.Path.ValueString())})
			if err != nil {
				resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *GnmiResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")
	idParts = helpers.RemoveEmptyStrings(idParts)

	if len(idParts) != 1 && len(idParts) != 2 {
		expectedIdentifier := "Expected import identifier with format: '<path>'"
		expectedIdentifier += " or '<path>,<device>'"
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("%s. Got: %q", expectedIdentifier, req.ID),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", idParts[0]))

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("path"), idParts[0])...)
	if len(idParts) == 2 {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device"), idParts[1])...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", idParts[0]))
}

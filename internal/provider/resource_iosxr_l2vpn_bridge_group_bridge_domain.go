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
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewL2VPNBridgeGroupBridgeDomainResource() resource.Resource {
	return &L2VPNBridgeGroupBridgeDomainResource{}
}

type L2VPNBridgeGroupBridgeDomainResource struct {
	client *client.Client
}

func (r *L2VPNBridgeGroupBridgeDomainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2vpn_bridge_group_bridge_domain"
}

func (r *L2VPNBridgeGroupBridgeDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the L2VPN Bridge Group Bridge Domain configuration.",

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
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"bridge_group_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the group the bridge belongs to").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"bridge_domain_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure bridge domain").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 27),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"evis": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ethernet VPN identifier").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Ethernet VPN identifier").AddIntegerRangeDescription(1, 65534).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65534),
							},
						},
					},
				},
			},
			"vnis": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VxLAN VPN identifier").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vni_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VxLAN VPN identifier").AddIntegerRangeDescription(1, 16777215).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 16777215),
							},
						},
					},
				},
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum transmission unit (payload) for this Bridge Domain").AddIntegerRangeDescription(46, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(46, 65535),
				},
			},
			"storm_control_broadcast_pps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the storm control pps").AddIntegerRangeDescription(1, 160000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 160000),
				},
			},
			"storm_control_broadcast_kbps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the storm control kbps").AddIntegerRangeDescription(64, 1280000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(64, 1280000),
				},
			},
			"storm_control_multicast_pps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the storm control pps").AddIntegerRangeDescription(1, 160000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 160000),
				},
			},
			"storm_control_multicast_kbps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the storm control kbps").AddIntegerRangeDescription(64, 1280000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(64, 1280000),
				},
			},
			"storm_control_unknown_unicast_pps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the storm control pps").AddIntegerRangeDescription(1, 160000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 160000),
				},
			},
			"storm_control_unknown_unicast_kbps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the storm control kbps").AddIntegerRangeDescription(64, 1280000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(64, 1280000),
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify interface name").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify interface name").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"split_horizon_group": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure split-horizon group").String,
							Optional:            true,
						},
					},
				},
			},
			"segment_routing_srv6_evis": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ethernet VPN identifier for srv6").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Ethernet VPN identifier for srv6").AddIntegerRangeDescription(1, 65534).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65534),
							},
						},
					},
				},
			},
		},
	}
}

func (r *L2VPNBridgeGroupBridgeDomainResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *L2VPNBridgeGroupBridgeDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan L2VPNBridgeGroupBridgeDomain

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, err := r.client.Set(ctx, plan.Device.ValueString(), ops...)
	if err != nil {
		resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
		return
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *L2VPNBridgeGroupBridgeDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state L2VPNBridgeGroupBridgeDomain

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	import_ := false
	if state.Id.ValueString() == "" {
		import_ = true
		state.Id = types.StringValue(state.getPath())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, err := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "Requested element(s) not found") {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError("Unable to apply gNMI Get operation", err.Error())
			return
		}
	}

	respBody := getResp.Notification[0].Update[0].Val.GetJsonIetfVal()
	if import_ {
		state.fromBody(ctx, respBody)
	} else {
		state.updateFromBody(ctx, respBody)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *L2VPNBridgeGroupBridgeDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state L2VPNBridgeGroupBridgeDomain

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

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	deletedListItems := plan.getDeletedItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, err := r.client.Set(ctx, plan.Device.ValueString(), ops...)
	if err != nil {
		resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *L2VPNBridgeGroupBridgeDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state L2VPNBridgeGroupBridgeDomain

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var ops []client.SetOperation
	deleteMode := "all"
	if state.DeleteMode.ValueString() == "all" {
		deleteMode = "all"
	} else if state.DeleteMode.ValueString() == "attributes" {
		deleteMode = "attributes"
	}

	if deleteMode == "all" {
		ops = append(ops, client.SetOperation{Path: state.Id.ValueString(), Body: "", Operation: client.Delete})
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		for _, i := range deletePaths {
			ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
		}
	}

	_, err := r.client.Set(ctx, state.Device.ValueString(), ops...)
	if err != nil {
		resp.Diagnostics.AddError("Unable to apply gNMI Set operation", err.Error())
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *L2VPNBridgeGroupBridgeDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <bridge_group_name>,<bridge_domain_name>. Got: %q", req.ID),
		)
		return
	}
	value0 := idParts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("bridge_group_name"), value0)...)
	value1 := idParts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("bridge_domain_name"), value1)...)
}

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

func NewRouterBGPVRFNeighborAddressFamilyResource() resource.Resource {
	return &RouterBGPVRFNeighborAddressFamilyResource{}
}

type RouterBGPVRFNeighborAddressFamilyResource struct {
	client *client.Client
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_vrf_neighbor_address_family"
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router BGP VRF Neighbor Address Family configuration.",

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
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a vrf name").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"neighbor_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"af_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter Address Family command mode").AddStringEnumDescription("all-address-family", "ipv4-flowspec", "ipv4-labeled-unicast", "ipv4-mdt", "ipv4-multicast", "ipv4-mvpn", "ipv4-rt-filter", "ipv4-sr-policy", "ipv4-tunnel", "ipv4-unicast", "ipv6-flowspec", "ipv6-labeled-unicast", "ipv6-multicast", "ipv6-mvpn", "ipv6-sr-policy", "ipv6-unicast", "l2vpn-evpn", "l2vpn-mspw", "l2vpn-vpls-vpws", "link-state-link-state", "vpnv4-flowspec", "vpnv4-multicast", "vpnv4-unicast", "vpnv6-flowspec", "vpnv6-multicast", "vpnv6-unicast").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all-address-family", "ipv4-flowspec", "ipv4-labeled-unicast", "ipv4-mdt", "ipv4-multicast", "ipv4-mvpn", "ipv4-rt-filter", "ipv4-sr-policy", "ipv4-tunnel", "ipv4-unicast", "ipv6-flowspec", "ipv6-labeled-unicast", "ipv6-multicast", "ipv6-mvpn", "ipv6-sr-policy", "ipv6-unicast", "l2vpn-evpn", "l2vpn-mspw", "l2vpn-vpls-vpws", "link-state-link-state", "vpnv4-flowspec", "vpnv4-multicast", "vpnv4-unicast", "vpnv6-flowspec", "vpnv6-multicast", "vpnv6-unicast"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"route_policy_in": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply route policy to inbound routes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"route_policy_out": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply route policy to outbound routes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"default_originate_route_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Route policy to specify criteria to originate default").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"default_originate_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent default-originate being inherited from a parent group").String,
				Optional:            true,
			},
			"next_hop_self": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable the next hop calculation for this neighbor").String,
				Optional:            true,
			},
			"next_hop_self_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent next-hop-self from being inherited from the parent").String,
				Optional:            true,
			},
			"soft_reconfiguration_inbound_always": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Always use soft reconfig, even if route refresh is supported").String,
				Optional:            true,
			},
			"send_community_ebgp_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent send-community-ebgp from being inherited from the parent").String,
				Optional:            true,
			},
			"remove_private_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Remove private AS number from outbound updates").String,
				Optional:            true,
			},
			"remove_private_as_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent remove-private-AS from being inherited from the parent").String,
				Optional:            true,
			},
			"remove_private_as_entire_aspath": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("remove only if all ASes in the path are private").String,
				Optional:            true,
			},
			"remove_private_as_internal": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("remove only if all ASes in the path are private").String,
				Optional:            true,
			},
			"remove_private_as_internal_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent remove-private-AS from being inherited from the parent").String,
				Optional:            true,
			},
			"remove_private_as_inbound": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Remove private AS number from inbound updates").String,
				Optional:            true,
			},
			"remove_private_as_inbound_entire_aspath": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("remove only if all ASes in the path are private").String,
				Optional:            true,
			},
			"remove_private_as_inbound_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prevent remove-private-AS from being inherited from the parent").String,
				Optional:            true,
			},
		},
	}
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterBGPVRFNeighborAddressFamilyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterBGPVRFNeighborAddressFamily

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

func (r *RouterBGPVRFNeighborAddressFamilyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 4 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <as_number>,<vrf_name>,<neighbor_address>,<af_name>. Got: %q", req.ID),
		)
		return
	}
	value0 := idParts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("as_number"), value0)...)
	value1 := idParts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_name"), value1)...)
	value2 := idParts[2]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("neighbor_address"), value2)...)
	value3 := idParts[3]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("af_name"), value3)...)
}

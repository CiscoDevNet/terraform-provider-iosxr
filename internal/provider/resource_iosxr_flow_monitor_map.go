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

func NewFlowMonitorMapResource() resource.Resource {
	return &FlowMonitorMapResource{}
}

type FlowMonitorMapResource struct {
	client *client.Client
}

func (r *FlowMonitorMapResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flow_monitor_map"
}

func (r *FlowMonitorMapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Flow Monitor Map configuration.",

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Monitor map name - maximum 32 characters").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"exporters": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify flow exporter map name").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify flow exporter map name").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
					},
				},
			},
			"option_outphysint": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("export output interfaces as physical interfaces").String,
				Optional:            true,
			},
			"option_filtered": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable filtering of records").String,
				Optional:            true,
			},
			"option_bgpattr": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("export bgp attributes AS_PATH and STD_COMMUNITY").String,
				Optional:            true,
			},
			"option_outbundlemember": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("export output physical interfaces of bundle interface").String,
				Optional:            true,
			},
			"record_ipv4": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 raw record format").String,
				Optional:            true,
			},
			"record_ipv4_destination": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Destination Based NetFlow Accounting").String,
				Optional:            true,
			},
			"record_ipv4_destination_tos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Destination Based NetFlow Accounting TOS").String,
				Optional:            true,
			},
			"record_ipv4_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Autonomous System based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_protocol_port": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Protocol-Port based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_prefix": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix based agregation").String,
				Optional:            true,
			},
			"record_ipv4_source_prefix": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("source prefix based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_destination_prefix": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_as_tos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AS-TOS based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_protocol_port_tos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Protocol, port and tos based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_prefix_tos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix TOS based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_source_prefix_tos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source, Prefix and TOS based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_destination_prefix_tos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination, prefix and tos based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_prefix_port": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix port based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_bgp_nexthop_tos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP, nexthop and tos based aggregation").String,
				Optional:            true,
			},
			"record_ipv4_peer_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Record peer AS").String,
				Optional:            true,
			},
			"record_ipv4_gtp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPV4 gtp record format").String,
				Optional:            true,
			},
			"record_ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 raw record format").String,
				Optional:            true,
			},
			"record_ipv6_destination": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Destination Based NetFlow Accounting").String,
				Optional:            true,
			},
			"record_ipv6_peer_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Record peer AS").String,
				Optional:            true,
			},
			"record_ipv6_gtp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPV6 gtp record format").String,
				Optional:            true,
			},
			"record_mpls": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MPLS record format").String,
				Optional:            true,
			},
			"record_mpls_ipv4_fields": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MPLS with IPv4 fields format").String,
				Optional:            true,
			},
			"record_mpls_ipv6_fields": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MPLS with IPv6 fields format").String,
				Optional:            true,
			},
			"record_mpls_ipv4_ipv6_fields": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MPLS with IPv4 and IPv6 fields format").String,
				Optional:            true,
			},
			"record_mpls_labels": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Labels to be used for Hashing").AddIntegerRangeDescription(1, 6).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 6),
				},
			},
			"record_map_t": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("map-t translation based Netflow").String,
				Optional:            true,
			},
			"record_sflow": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("sFlow based flow").String,
				Optional:            true,
			},
			"record_datalink_record": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Layer2 traffic based flow").String,
				Optional:            true,
			},
			"record_default_rtp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default RTP record format").String,
				Optional:            true,
			},
			"record_default_mdi": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default MDI record format").String,
				Optional:            true,
			},
			"cache_entries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the number of entries in the flow cache").AddIntegerRangeDescription(4096, 1000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(4096, 1000000),
				},
			},
			"cache_timeout_active": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the active flow timeout").AddIntegerRangeDescription(1, 604800).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 604800),
				},
			},
			"cache_timeout_inactive": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the inactive flow timeout").AddIntegerRangeDescription(0, 604800).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 604800),
				},
			},
			"cache_timeout_update": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the update timeout").AddIntegerRangeDescription(1, 604800).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 604800),
				},
			},
			"cache_timeout_rate_limit": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum number of entries to age each second").AddIntegerRangeDescription(1, 1000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1000000),
				},
			},
			"cache_permanent": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable removal of entries from flow cache").String,
				Optional:            true,
			},
			"cache_immediate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Immediate removal of entries from flow cache").String,
				Optional:            true,
			},
			"hw_cache_timeout_inactive": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the inactive timeout").AddIntegerRangeDescription(50, 1800).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(50, 1800),
				},
			},
			"sflow_options": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("submode to configure sFlow related options").String,
				Optional:            true,
			},
			"sflow_options_extended_router": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable extended-router flow data type").String,
				Optional:            true,
			},
			"sflow_options_extended_gateway": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable extended-gateway flow data type").String,
				Optional:            true,
			},
			"sflow_options_extended_ipv4_tunnel_egress": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable extended-ipv4-tunnel-egress flow data type").String,
				Optional:            true,
			},
			"sflow_options_extended_ipv6_tunnel_egress": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable extended-ipv6-tunnel-egress flow data type").String,
				Optional:            true,
			},
			"sflow_options_if_counters_polling_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable if-counters counter sampling rate").AddIntegerRangeDescription(5, 1800).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 1800),
				},
			},
			"sflow_options_sample_header_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify maximum sample-header size to be exported").AddIntegerRangeDescription(128, 200).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(128, 200),
				},
			},
			"sflow_options_input_ifindex": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify ifindex related options").AddStringEnumDescription("physical").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("physical"),
				},
			},
			"sflow_options_output_ifindex": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify ifindex related options").AddStringEnumDescription("physical").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("physical"),
				},
			},
		},
	}
}

func (r *FlowMonitorMapResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *FlowMonitorMapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FlowMonitorMap

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

func (r *FlowMonitorMapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FlowMonitorMap

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

func (r *FlowMonitorMapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FlowMonitorMap

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

func (r *FlowMonitorMapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FlowMonitorMap

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var ops []client.SetOperation
	deleteMode := "all"

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

func (r *FlowMonitorMapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <name>. Got: %q", req.ID),
		)
		return
	}
	value0 := idParts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), value0)...)
}

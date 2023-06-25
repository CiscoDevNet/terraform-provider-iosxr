// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewEVPNEVIResource() resource.Resource {
	return &EVPNEVIResource{}
}

type EVPNEVIResource struct {
	client *client.Client
}

func (r *EVPNEVIResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_evpn_evi"
}

func (r *EVPNEVIResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the EVPN EVI configuration.",

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
			"vpn_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure EVPN Instance VPN ID").AddIntegerRangeDescription(1, 65534).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description for this EVPN Instance").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"load_balancing": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure EVPN Instance load-balancing").String,
				Optional:            true,
			},
			"load_balancing_flow_label_static": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static configuration of Flow Label").String,
				Optional:            true,
			},
			"bgp_rd_two_byte_as_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"bgp_rd_two_byte_as_assigned_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"bgp_rd_four_byte_as_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number").AddIntegerRangeDescription(65536, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(65536, 4294967295),
				},
			},
			"bgp_rd_four_byte_as_assigned_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"bgp_rd_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
			},
			"bgp_rd_ipv4_address_assigned_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP-address:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"bgp_route_target_import_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
					},
				},
			},
			"bgp_route_target_import_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number").AddIntegerRangeDescription(65536, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
					},
				},
			},
			"bgp_route_target_import_ipv4_address_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP-address:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
					},
				},
			},
			"bgp_route_target_export_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
					},
				},
			},
			"bgp_route_target_export_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number").AddIntegerRangeDescription(65536, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
					},
				},
			},
			"bgp_route_target_export_ipv4_address_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP-address:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
					},
				},
			},
			"bgp_route_policy_import": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Import route policy").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"bgp_route_policy_export": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Export route policy").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"advertise_mac": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure EVPN Instance MAC advertisement").String,
				Optional:            true,
			},
			"unknown_unicast_suppression": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enabling unknown unicast suppression").String,
				Optional:            true,
			},
			"control_word_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disabling control-word").String,
				Optional:            true,
			},
			"etree": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure EVPN Instance E-Tree").String,
				Optional:            true,
			},
			"etree_leaf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Designate EVPN Instance as EVPN E-Tree Leaf Site").String,
				Optional:            true,
			},
			"etree_rt_leaf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Designate EVPN Instance as EVPN E-Tree Route-Target Leaf Site").String,
				Optional:            true,
			},
		},
	}
}

func (r *EVPNEVIResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *EVPNEVIResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EVPNEVI

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

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *EVPNEVIResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EVPNEVI

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.updateFromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *EVPNEVIResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state EVPNEVI

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

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *EVPNEVIResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state EVPNEVI

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

	_, diags = r.client.Set(ctx, state.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *EVPNEVIResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

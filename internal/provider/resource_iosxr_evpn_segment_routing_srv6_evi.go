// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

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
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
)

var _ resource.Resource = (*EVPNSegmentRoutingSRv6EVIResource)(nil)

func NewEVPNSegmentRoutingSRv6EVIResource() resource.Resource {
	return &EVPNSegmentRoutingSRv6EVIResource{}
}

type EVPNSegmentRoutingSRv6EVIResource struct {
	client *client.Client
}

func (r *EVPNSegmentRoutingSRv6EVIResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_evpn_segment_routing_srv6_evi"
}

func (r *EVPNSegmentRoutingSRv6EVIResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the EVPN Segment Routing SRv6 EVI configuration.",

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
			"bgp_route_target_import_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
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
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
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
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP-address:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
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
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
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
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
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
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"assigned_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP-address:nn (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
					},
				},
			},
			"advertise_mac": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure EVPN Instance MAC advertisement").String,
				Optional:            true,
			},
			"locator": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EVI locator to use for EVPN SID allocation").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 58),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9A-Z][a-z0-9A-Z_.:]*`), ""),
				},
			},
		},
	}
}

func (r *EVPNSegmentRoutingSRv6EVIResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *EVPNSegmentRoutingSRv6EVIResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EVPNSegmentRoutingSRv6EVI

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

func (r *EVPNSegmentRoutingSRv6EVIResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EVPNSegmentRoutingSRv6EVI

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

func (r *EVPNSegmentRoutingSRv6EVIResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state EVPNSegmentRoutingSRv6EVI

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

func (r *EVPNSegmentRoutingSRv6EVIResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state EVPNSegmentRoutingSRv6EVI

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	_, diags = r.client.Set(ctx, state.Device.ValueString(), client.SetOperation{Path: state.getPath(), Body: "", Operation: client.Delete})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *EVPNSegmentRoutingSRv6EVIResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

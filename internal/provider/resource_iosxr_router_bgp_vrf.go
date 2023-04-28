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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
)

var _ resource.Resource = (*RouterBGPVRFResource)(nil)

func NewRouterBGPVRFResource() resource.Resource {
	return &RouterBGPVRFResource{}
}

type RouterBGPVRFResource struct {
	client *client.Client
}

func (r *RouterBGPVRFResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_vrf"
}

func (r *RouterBGPVRFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router BGP VRF configuration.",

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
			"rd_auto": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Automatic route distinguisher").String,
				Optional:            true,
			},
			"rd_two_byte_as_as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Required:            true,
			},
			"rd_two_byte_as_index": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"rd_four_byte_as_as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("4-byte AS number").String,
				Required:            true,
			},
			"rd_four_byte_as_index": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"rd_ip_address_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("configure this node").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
			},
			"rd_ip_address_index": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4Address:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Distribute a default route").String,
				Optional:            true,
			},
			"default_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("default redistributed metric").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"timers_bgp_keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP timers").AddIntegerRangeDescription(0, 65535).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"timers_bgp_holdtime": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Holdtime. Set 0 to disable keepalives/hold time.").String,
				Required:            true,
			},
			"bfd_minimum_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello interval").AddIntegerRangeDescription(3, 30000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 30000),
				},
			},
			"bfd_multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(2, 16).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 16),
				},
			},
			"neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"neighbor_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
							Optional:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Neighbor specific description").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
							},
						},
						"ignore_connected_check": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Bypass the directly connected nexthop check for single-hop eBGP peering").String,
							Optional:            true,
						},
						"ebgp_multihop_maximum_hop_count": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("maximum hop count").AddIntegerRangeDescription(1, 255).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"bfd_minimum_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hello interval").AddIntegerRangeDescription(3, 30000).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(3, 30000),
							},
						},
						"bfd_multiplier": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(2, 16).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(2, 16),
							},
						},
						"local_as": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
							Optional:            true,
						},
						"local_as_no_prepend": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not prepend local AS to announcements from this neighbor").String,
							Optional:            true,
						},
						"local_as_replace_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prepend only local AS to announcements to this neighbor").String,
							Optional:            true,
						},
						"local_as_dual_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Dual-AS mode").String,
							Optional:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specifies an ENCRYPTED password will follow").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(!.+)|([^!].+)`), ""),
							},
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administratively shut down this neighbor").String,
							Optional:            true,
						},
						"timers_keepalive_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("BGP timers").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"timers_holdtime": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Holdtime. Set 0 to disable keepalives/hold time.").String,
							Required:            true,
						},
						"update_source": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source of routing updates").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"ttl_security": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable EBGP TTL security").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *RouterBGPVRFResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterBGPVRFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterBGPVRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterBGPVRFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterBGPVRF

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

func (r *RouterBGPVRFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterBGPVRF

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterBGPVRFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterBGPVRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	_, diags = r.client.Set(ctx, state.Device.ValueString(), state.getPath(), "", client.Delete)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *RouterBGPVRFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

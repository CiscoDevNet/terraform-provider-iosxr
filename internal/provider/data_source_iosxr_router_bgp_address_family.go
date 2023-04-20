// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterBGPAddressFamilyDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterBGPAddressFamilyDataSource{}
)

func NewRouterBGPAddressFamilyDataSource() datasource.DataSource {
	return &RouterBGPAddressFamilyDataSource{}
}

type RouterBGPAddressFamilyDataSource struct {
	client *client.Client
}

func (d *RouterBGPAddressFamilyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_address_family"
}

func (d *RouterBGPAddressFamilyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router BGP Address Family configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "bgp as-number",
				Required:            true,
			},
			"af_name": schema.StringAttribute{
				MarkdownDescription: "Enter Address Family command mode",
				Required:            true,
			},
			"additional_paths_send": schema.BoolAttribute{
				MarkdownDescription: "Additional paths Send capability",
				Computed:            true,
			},
			"additional_paths_receive": schema.BoolAttribute{
				MarkdownDescription: "Additional paths Receive capability",
				Computed:            true,
			},
			"additional_paths_selection_route_policy": schema.StringAttribute{
				MarkdownDescription: "Route-policy for additional paths selection",
				Computed:            true,
			},
			"allocate_label_all_unlabeled_path": schema.BoolAttribute{
				MarkdownDescription: "Allocate label for unlabeled paths too",
				Computed:            true,
			},
			"advertise_best_external": schema.BoolAttribute{
				MarkdownDescription: "Advertise best-external path",
				Computed:            true,
			},
			"allocate_label_all": schema.BoolAttribute{
				MarkdownDescription: "Allocate labels for all prefixes",
				Computed:            true,
			},
			"maximum_paths_ebgp_multipath": schema.Int64Attribute{
				MarkdownDescription: "eBGP-multipath",
				Computed:            true,
			},
			"maximum_paths_eibgp_multipath": schema.Int64Attribute{
				MarkdownDescription: "eiBGP-multipath",
				Computed:            true,
			},
			"maximum_paths_ibgp_multipath": schema.Int64Attribute{
				MarkdownDescription: "iBGP-multipath",
				Computed:            true,
			},
			"label_mode_per_ce": schema.BoolAttribute{
				MarkdownDescription: "Set per CE label mode",
				Computed:            true,
			},
			"label_mode_per_vrf": schema.BoolAttribute{
				MarkdownDescription: "Set per VRF label mode",
				Computed:            true,
			},
			"redistribute_connected": schema.BoolAttribute{
				MarkdownDescription: "Connected routes",
				Computed:            true,
			},
			"redistribute_connected_metric": schema.Int64Attribute{
				MarkdownDescription: "Metric for redistributed routes",
				Computed:            true,
			},
			"redistribute_static": schema.BoolAttribute{
				MarkdownDescription: "Static routes",
				Computed:            true,
			},
			"redistribute_static_metric": schema.Int64Attribute{
				MarkdownDescription: "Metric for redistributed routes",
				Computed:            true,
			},
			"aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 Aggregate address and mask or masklength",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 Aggregate address and mask or masklength",
							Computed:            true,
						},
						"masklength": schema.Int64Attribute{
							MarkdownDescription: "Network in prefix/length format (prefix part)",
							Computed:            true,
						},
						"as_set": schema.BoolAttribute{
							MarkdownDescription: "Generate AS set path information",
							Computed:            true,
						},
						"as_confed_set": schema.BoolAttribute{
							MarkdownDescription: "Generate AS confed set path information",
							Computed:            true,
						},
						"summary_only": schema.BoolAttribute{
							MarkdownDescription: "Filter more specific routes from updates",
							Computed:            true,
						},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 network and mask or masklength",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 network and mask or masklength",
							Computed:            true,
						},
						"masklength": schema.Int64Attribute{
							MarkdownDescription: "Network in prefix/length format (prefix part)",
							Computed:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: "Route-policy to modify the attributes",
							Computed:            true,
						},
					},
				},
			},
			"redistribute_isis": schema.ListNestedAttribute{
				MarkdownDescription: "ISO IS-IS",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"instance_name": schema.StringAttribute{
							MarkdownDescription: "ISO IS-IS",
							Computed:            true,
						},
						"level_one": schema.BoolAttribute{
							MarkdownDescription: "Redistribute ISIS level 1 routes",
							Computed:            true,
						},
						"level_one_two": schema.BoolAttribute{
							MarkdownDescription: "Redistribute ISIS level 2 ISIS routes",
							Computed:            true,
						},
						"level_one_two_one_inter_area": schema.BoolAttribute{
							MarkdownDescription: "Redistribute ISIS level 1 inter-area routes",
							Computed:            true,
						},
						"level_one_one_inter_area": schema.BoolAttribute{
							MarkdownDescription: "Redistribute ISIS level 1 inter-area routes",
							Computed:            true,
						},
						"level_two": schema.BoolAttribute{
							MarkdownDescription: "Redistribute ISIS level 2 ISIS routes",
							Computed:            true,
						},
						"level_two_one_inter_area": schema.BoolAttribute{
							MarkdownDescription: "Redistribute ISIS level 1 inter-area routes",
							Computed:            true,
						},
						"level_one_inter_area": schema.BoolAttribute{
							MarkdownDescription: "Redistribute ISIS level 1 inter-area routes",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Metric for redistributed routes",
							Computed:            true,
						},
					},
				},
			},
			"redistribute_ospf": schema.ListNestedAttribute{
				MarkdownDescription: "Open Shortest Path First (OSPF or OSPFv3)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"router_tag": schema.StringAttribute{
							MarkdownDescription: "Open Shortest Path First (OSPF)",
							Computed:            true,
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: "Redistribute OSPF internal routes",
							Computed:            true,
						},
						"match_internal_external": schema.BoolAttribute{
							MarkdownDescription: "Redistribute OSPF external routes",
							Computed:            true,
						},
						"match_internal_nssa_external": schema.BoolAttribute{
							MarkdownDescription: "Redistribute OSPF NSSA external routes",
							Computed:            true,
						},
						"match_external": schema.BoolAttribute{
							MarkdownDescription: "Redistribute OSPF external routes",
							Computed:            true,
						},
						"match_external_nssa_external": schema.BoolAttribute{
							MarkdownDescription: "Redistribute OSPF NSSA external routes",
							Computed:            true,
						},
						"match_nssa_external": schema.BoolAttribute{
							MarkdownDescription: "Redistribute OSPF NSSA external routes",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Metric for redistributed routes",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *RouterBGPAddressFamilyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterBGPAddressFamilyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterBGPAddressFamily

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

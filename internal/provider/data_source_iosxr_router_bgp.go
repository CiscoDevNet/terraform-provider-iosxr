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
	_ datasource.DataSource              = &RouterBGPDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterBGPDataSource{}
)

func NewRouterBGPDataSource() datasource.DataSource {
	return &RouterBGPDataSource{}
}

type RouterBGPDataSource struct {
	client *client.Client
}

func (d *RouterBGPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp"
}

func (d *RouterBGPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router BGP configuration.",

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
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: "Distribute a default route",
				Computed:            true,
			},
			"default_metric": schema.Int64Attribute{
				MarkdownDescription: "default redistributed metric",
				Computed:            true,
			},
			"timers_bgp_keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: "BGP timers",
				Computed:            true,
			},
			"timers_bgp_holdtime": schema.StringAttribute{
				MarkdownDescription: "Holdtime. Set 0 to disable keepalives/hold time.",
				Computed:            true,
			},
			"bgp_router_id": schema.StringAttribute{
				MarkdownDescription: "Configure Router-id",
				Computed:            true,
			},
			"bgp_graceful_restart_graceful_reset": schema.BoolAttribute{
				MarkdownDescription: "Reset gracefully if configuration change forces a peer reset",
				Computed:            true,
			},
			"ibgp_policy_out_enforce_modifications": schema.BoolAttribute{
				MarkdownDescription: "Allow policy to modify all attributes",
				Computed:            true,
			},
			"bgp_log_neighbor_changes_detail": schema.BoolAttribute{
				MarkdownDescription: "Include extra detail in change messages",
				Computed:            true,
			},
			"bfd_minimum_interval": schema.Int64Attribute{
				MarkdownDescription: "Hello interval",
				Computed:            true,
			},
			"bfd_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Detect multiplier",
				Computed:            true,
			},
			"neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "Neighbor address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"neighbor_address": schema.StringAttribute{
							MarkdownDescription: "Neighbor address",
							Computed:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: "bgp as-number",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Neighbor specific description",
							Computed:            true,
						},
						"use_neighbor_group": schema.StringAttribute{
							MarkdownDescription: "Inherit configuration from a neighbor-group",
							Computed:            true,
						},
						"ignore_connected_check": schema.BoolAttribute{
							MarkdownDescription: "Bypass the directly connected nexthop check for single-hop eBGP peering",
							Computed:            true,
						},
						"ebgp_multihop_maximum_hop_count": schema.Int64Attribute{
							MarkdownDescription: "maximum hop count",
							Computed:            true,
						},
						"bfd_minimum_interval": schema.Int64Attribute{
							MarkdownDescription: "Hello interval",
							Computed:            true,
						},
						"bfd_multiplier": schema.Int64Attribute{
							MarkdownDescription: "Detect multiplier",
							Computed:            true,
						},
						"local_as": schema.StringAttribute{
							MarkdownDescription: "bgp as-number",
							Computed:            true,
						},
						"local_as_no_prepend": schema.BoolAttribute{
							MarkdownDescription: "Do not prepend local AS to announcements from this neighbor",
							Computed:            true,
						},
						"local_as_replace_as": schema.BoolAttribute{
							MarkdownDescription: "Prepend only local AS to announcements to this neighbor",
							Computed:            true,
						},
						"local_as_dual_as": schema.BoolAttribute{
							MarkdownDescription: "Dual-AS mode",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Specifies an ENCRYPTED password will follow",
							Computed:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: "Administratively shut down this neighbor",
							Computed:            true,
						},
						"timers_keepalive_interval": schema.Int64Attribute{
							MarkdownDescription: "BGP timers",
							Computed:            true,
						},
						"timers_holdtime": schema.StringAttribute{
							MarkdownDescription: "Holdtime. Set 0 to disable keepalives/hold time.",
							Computed:            true,
						},
						"update_source": schema.StringAttribute{
							MarkdownDescription: "Source of routing updates",
							Computed:            true,
						},
						"ttl_security": schema.BoolAttribute{
							MarkdownDescription: "Enable EBGP TTL security",
							Computed:            true,
						},
					},
				},
			},
			"neighbor_groups": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a Neighbor-group",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"neighbor_group_name": schema.StringAttribute{
							MarkdownDescription: "Specify a Neighbor-group",
							Computed:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: "bgp as-number",
							Computed:            true,
						},
						"update_source": schema.StringAttribute{
							MarkdownDescription: "Source of routing updates",
							Computed:            true,
						},
						"ao_key_chain_name": schema.StringAttribute{
							MarkdownDescription: "Name of the key chain - maximum 32 characters",
							Computed:            true,
						},
						"ao_include_tcp_options_enable": schema.BoolAttribute{
							MarkdownDescription: "Include other TCP options in the header",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *RouterBGPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterBGPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterBGP

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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterISISDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterISISDataSource{}
)

func NewRouterISISDataSource() datasource.DataSource {
	return &RouterISISDataSource{}
}

type RouterISISDataSource struct {
	client *client.Client
}

func (d *RouterISISDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_isis"
}

func (d *RouterISISDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router ISIS configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the retrieved object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"process_id": {
				MarkdownDescription: "Process ID",
				Type:                types.StringType,
				Required:            true,
			},
			"is_type": {
				MarkdownDescription: "Area type (level)",
				Type:                types.StringType,
				Computed:            true,
			},
			"nets": {
				MarkdownDescription: "A Network Entity Title (NET) for this process",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"net_id": {
						MarkdownDescription: "A Network Entity Title (NET) for this process",
						Type:                types.StringType,
						Computed:            true,
					},
				}),
			},
			"address_families": {
				MarkdownDescription: "IS-IS address family",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"af_name": {
						MarkdownDescription: "Address family name",
						Type:                types.StringType,
						Computed:            true,
					},
					"saf_name": {
						MarkdownDescription: "Sub address family name",
						Type:                types.StringType,
						Computed:            true,
					},
					"mpls_ldp_auto_config": {
						MarkdownDescription: "Enable LDP IGP interface auto-configuration",
						Type:                types.BoolType,
						Computed:            true,
					},
					"metric_style_narrow": {
						MarkdownDescription: "Use old style of TLVs with narrow metric",
						Type:                types.BoolType,
						Computed:            true,
					},
					"metric_style_wide": {
						MarkdownDescription: "Use new style of TLVs to carry wider metric",
						Type:                types.BoolType,
						Computed:            true,
					},
					"metric_style_transition": {
						MarkdownDescription: "Send and accept both styles of TLVs during transition",
						Type:                types.BoolType,
						Computed:            true,
					},
					"router_id_interface_name": {
						MarkdownDescription: "Router ID Interface",
						Type:                types.StringType,
						Computed:            true,
					},
					"router_id_ip_address": {
						MarkdownDescription: "Router ID address",
						Type:                types.StringType,
						Computed:            true,
					},
					"default_information_originate": {
						MarkdownDescription: "Distribute a default route",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"interfaces": {
				MarkdownDescription: "Enter the IS-IS interface configuration submode",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"interface_name": {
						MarkdownDescription: "Enter the IS-IS interface configuration submode",
						Type:                types.StringType,
						Computed:            true,
					},
					"circuit_type": {
						MarkdownDescription: "Configure circuit type for interface",
						Type:                types.StringType,
						Computed:            true,
					},
					"hello_padding_disable": {
						MarkdownDescription: "Disable hello-padding",
						Type:                types.BoolType,
						Computed:            true,
					},
					"hello_padding_sometimes": {
						MarkdownDescription: "Enable hello-padding during adjacency formation only",
						Type:                types.BoolType,
						Computed:            true,
					},
					"priority": {
						MarkdownDescription: "Set priority for Designated Router election",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"point_to_point": {
						MarkdownDescription: "Treat active LAN interface as point-to-point",
						Type:                types.BoolType,
						Computed:            true,
					},
					"passive": {
						MarkdownDescription: "Do not establish adjacencies over this interface",
						Type:                types.BoolType,
						Computed:            true,
					},
					"suppressed": {
						MarkdownDescription: "Do not advertise connected prefixes of this interface",
						Type:                types.BoolType,
						Computed:            true,
					},
					"shutdown": {
						MarkdownDescription: "Shutdown IS-IS on this interface",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (d *RouterISISDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterISISDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterISIS

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

	config.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	tflog.Debug(ctx, fmt.Sprintf("DSDEBUG: %v", config))
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

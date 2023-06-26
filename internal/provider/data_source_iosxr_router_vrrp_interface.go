// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterVRRPInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterVRRPInterfaceDataSource{}
)

func NewRouterVRRPInterfaceDataSource() datasource.DataSource {
	return &RouterVRRPInterfaceDataSource{}
}

type RouterVRRPInterfaceDataSource struct {
	client *client.Client
}

func (d *RouterVRRPInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_vrrp_interface"
}

func (d *RouterVRRPInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router VRRP Interface configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "VRRP interface configuration subcommands",
				Required:            true,
			},
			"mac_refresh": schema.Int64Attribute{
				MarkdownDescription: "Set the Subordinate MAC-refresh rate for this interface",
				Computed:            true,
			},
			"delay_minimum": schema.Int64Attribute{
				MarkdownDescription: "Set minimum delay on every interface up event",
				Computed:            true,
			},
			"delay_reload": schema.Int64Attribute{
				MarkdownDescription: "Set reload delay for first interface up event",
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
		},
	}
}

func (d *RouterVRRPInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterVRRPInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterVRRPInterface

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

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
	_ datasource.DataSource              = &RouterISISInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterISISInterfaceDataSource{}
)

func NewRouterISISInterfaceDataSource() datasource.DataSource {
	return &RouterISISInterfaceDataSource{}
}

type RouterISISInterfaceDataSource struct {
	client *client.Client
}

func (d *RouterISISInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_isis_interface"
}

func (d *RouterISISInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router ISIS Interface configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"process_id": schema.StringAttribute{
				MarkdownDescription: "Process ID",
				Required:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Enter the IS-IS interface configuration submode",
				Required:            true,
			},
			"circuit_type": schema.StringAttribute{
				MarkdownDescription: "Configure circuit type for interface",
				Computed:            true,
			},
			"hello_padding_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable hello-padding",
				Computed:            true,
			},
			"hello_padding_sometimes": schema.BoolAttribute{
				MarkdownDescription: "Enable hello-padding during adjacency formation only",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Set priority for Designated Router election",
				Computed:            true,
			},
			"point_to_point": schema.BoolAttribute{
				MarkdownDescription: "Treat active LAN interface as point-to-point",
				Computed:            true,
			},
			"passive": schema.BoolAttribute{
				MarkdownDescription: "Do not establish adjacencies over this interface",
				Computed:            true,
			},
			"suppressed": schema.BoolAttribute{
				MarkdownDescription: "Do not advertise connected prefixes of this interface",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown IS-IS on this interface",
				Computed:            true,
			},
			"hello_password_text": schema.StringAttribute{
				MarkdownDescription: "The encrypted LSP/SNP password",
				Computed:            true,
			},
			"hello_password_hmac_md5": schema.StringAttribute{
				MarkdownDescription: "The encrypted LSP/SNP password",
				Computed:            true,
			},
			"hello_password_keychain": schema.StringAttribute{
				MarkdownDescription: "Specifies a Key Chain name will follow",
				Computed:            true,
			},
		},
	}
}

func (d *RouterISISInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterISISInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterISISInterface

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

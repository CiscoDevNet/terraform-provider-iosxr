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
	_ datasource.DataSource              = &IPv4PrefixListDataSource{}
	_ datasource.DataSourceWithConfigure = &IPv4PrefixListDataSource{}
)

func NewIPv4PrefixListDataSource() datasource.DataSource {
	return &IPv4PrefixListDataSource{}
}

type IPv4PrefixListDataSource struct {
	client *client.Client
}

func (d *IPv4PrefixListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv4_prefix_list"
}

func (d *IPv4PrefixListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the IPv4 Prefix List configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"prefix_list_name": schema.StringAttribute{
				MarkdownDescription: "Name of a prefix list - maximum 32 characters",
				Required:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "Sequence number",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sequence_number": schema.Int64Attribute{
							MarkdownDescription: "Sequence number",
							Computed:            true,
						},
						"remark": schema.StringAttribute{
							MarkdownDescription: "Comments for prefix list",
							Computed:            true,
						},
						"permission": schema.StringAttribute{
							MarkdownDescription: "specify the type to be either deny (or) permit",
							Computed:            true,
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: "IPv4 address prefix",
							Computed:            true,
						},
						"mask": schema.StringAttribute{
							MarkdownDescription: "Mask length of IPv4 address",
							Computed:            true,
						},
						"match_prefix_length_eq": schema.Int64Attribute{
							MarkdownDescription: "Exact prefix length to be matched",
							Computed:            true,
						},
						"match_prefix_length_ge": schema.Int64Attribute{
							MarkdownDescription: "Minimum prefix length to be matched",
							Computed:            true,
						},
						"match_prefix_length_le": schema.Int64Attribute{
							MarkdownDescription: "Maximum prefix length to be matched",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *IPv4PrefixListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *IPv4PrefixListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IPv4PrefixList

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

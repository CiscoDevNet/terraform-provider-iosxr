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
	_ datasource.DataSource              = &RouterVRRPInterfaceAddressFamilyIPv6DataSource{}
	_ datasource.DataSourceWithConfigure = &RouterVRRPInterfaceAddressFamilyIPv6DataSource{}
)

func NewRouterVRRPInterfaceAddressFamilyIPv6DataSource() datasource.DataSource {
	return &RouterVRRPInterfaceAddressFamilyIPv6DataSource{}
}

type RouterVRRPInterfaceAddressFamilyIPv6DataSource struct {
	client *client.Client
}

func (d *RouterVRRPInterfaceAddressFamilyIPv6DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_vrrp_interface_address_family_ipv6"
}

func (d *RouterVRRPInterfaceAddressFamilyIPv6DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router VRRP Interface Address Family IPv6 configuration.",

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
			"vrrp_id": schema.Int64Attribute{
				MarkdownDescription: "VRRP configuration",
				Required:            true,
			},
			"global_address": schema.StringAttribute{
				MarkdownDescription: "Set Global VRRP IPv6 address",
				Computed:            true,
			},
			"address_linklocal_linklocal_address": schema.StringAttribute{
				MarkdownDescription: "VRRP IPv6 linklocal address",
				Computed:            true,
			},
			"address_linklocal_autoconfig": schema.BoolAttribute{
				MarkdownDescription: "Autoconfigure the VRRP IPv6 linklocal address",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Set priority level",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Configure VRRP Session name",
				Computed:            true,
			},
			"timer_advertisement_time_in_seconds": schema.Int64Attribute{
				MarkdownDescription: "Advertisement time in seconds",
				Computed:            true,
			},
			"timer_advertisement_time_in_milliseconds": schema.Int64Attribute{
				MarkdownDescription: "Configure in milliseconds",
				Computed:            true,
			},
			"timer_force": schema.BoolAttribute{
				MarkdownDescription: "Force the configured values to be used",
				Computed:            true,
			},
			"preempt_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable preemption",
				Computed:            true,
			},
			"preempt_delay": schema.Int64Attribute{
				MarkdownDescription: "Wait before preempting",
				Computed:            true,
			},
			"accept_mode_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable accept mode",
				Computed:            true,
			},
			"track_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Track an interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Track an interface",
							Computed:            true,
						},
						"priority_decrement": schema.Int64Attribute{
							MarkdownDescription: "Priority decrement",
							Computed:            true,
						},
					},
				},
			},
			"track_objects": schema.ListNestedAttribute{
				MarkdownDescription: "Object Tracking",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object_name": schema.StringAttribute{
							MarkdownDescription: "Object to be tracked",
							Computed:            true,
						},
						"priority_decrement": schema.Int64Attribute{
							MarkdownDescription: "Priority decrement",
							Computed:            true,
						},
					},
				},
			},
			"bfd_fast_detect_peer_ipv6": schema.StringAttribute{
				MarkdownDescription: "VRRP BFD remote interface IP address",
				Computed:            true,
			},
		},
	}
}

func (d *RouterVRRPInterfaceAddressFamilyIPv6DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterVRRPInterfaceAddressFamilyIPv6DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterVRRPInterfaceAddressFamilyIPv6

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

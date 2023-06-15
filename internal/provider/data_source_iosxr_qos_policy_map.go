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
	_ datasource.DataSource              = &QOSPolicyMapDataSource{}
	_ datasource.DataSourceWithConfigure = &QOSPolicyMapDataSource{}
)

func NewQOSPolicyMapDataSource() datasource.DataSource {
	return &QOSPolicyMapDataSource{}
}

type QOSPolicyMapDataSource struct {
	client *client.Client
}

func (d *QOSPolicyMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_policy_map"
}

func (d *QOSPolicyMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the QOS Policy Map configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"policy_map_name": schema.StringAttribute{
				MarkdownDescription: "Name of the policymap",
				Required:            true,
			},
			"class_name": schema.StringAttribute{
				MarkdownDescription: "Name of the class-map",
				Computed:            true,
			},
			"class_type": schema.StringAttribute{
				MarkdownDescription: "The type of class-map",
				Computed:            true,
			},
			"class_set_mpls_experimental_topmost": schema.Int64Attribute{
				MarkdownDescription: "Sets the experimental value of the MPLS packet top-most labels.",
				Computed:            true,
			},
			"class_set_dscp": schema.StringAttribute{
				MarkdownDescription: "Set IP DSCP (DiffServ CodePoint)",
				Computed:            true,
			},
			"class_priority_level": schema.Int64Attribute{
				MarkdownDescription: "Configure a priority level",
				Computed:            true,
			},
			"class_queue_limits": schema.ListNestedAttribute{
				MarkdownDescription: "Configure queue-limit (taildrop threshold) for this class",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value": schema.StringAttribute{
							MarkdownDescription: "queue-limit value",
							Computed:            true,
						},
						"unit": schema.StringAttribute{
							MarkdownDescription: "queue-limit unit",
							Computed:            true,
						},
					},
				},
			},
			"class_service_policy_name": schema.StringAttribute{
				MarkdownDescription: "Name of the child service policy",
				Computed:            true,
			},
			"class_police_rate_value": schema.StringAttribute{
				MarkdownDescription: "Committed Information Rate",
				Computed:            true,
			},
			"class_police_rate_unit": schema.StringAttribute{
				MarkdownDescription: "Rate unit",
				Computed:            true,
			},
			"class_shape_average_rate_value": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"class_shape_average_rate_unit": schema.StringAttribute{
				MarkdownDescription: "Shape rate unit",
				Computed:            true,
			},
			"class_bandwidth_remaining_unit": schema.StringAttribute{
				MarkdownDescription: "Bandwidth value unit",
				Computed:            true,
			},
			"class_bandwidth_remaining_value": schema.StringAttribute{
				MarkdownDescription: "Bandwidth value",
				Computed:            true,
			},
		},
	}
}

func (d *QOSPolicyMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *QOSPolicyMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config QOSPolicyMap

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

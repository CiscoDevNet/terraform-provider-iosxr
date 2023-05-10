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
	_ datasource.DataSource              = &SegmentRoutingCandidatePathsDataSource{}
	_ datasource.DataSourceWithConfigure = &SegmentRoutingCandidatePathsDataSource{}
)

func NewSegmentRoutingCandidatePathsDataSource() datasource.DataSource {
	return &SegmentRoutingCandidatePathsDataSource{}
}

type SegmentRoutingCandidatePathsDataSource struct {
	client *client.Client
}

func (d *SegmentRoutingCandidatePathsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_segment_routing_candidate_paths"
}

func (d *SegmentRoutingCandidatePathsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Segment Routing Candidate Paths configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"policy_name": schema.StringAttribute{
				MarkdownDescription: "Policy name",
				Required:            true,
			},
			"path_index": schema.Int64Attribute{
				MarkdownDescription: "Path-option preference",
				Required:            true,
			},
			"candidate_paths_type": schema.ListNestedAttribute{
				MarkdownDescription: "Policy configuration",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Path-option type",
							Computed:            true,
						},
						"pcep": schema.BoolAttribute{
							MarkdownDescription: "Path Computation Element Protocol",
							Computed:            true,
						},
						"metric_metric_type": schema.StringAttribute{
							MarkdownDescription: "Metric type",
							Computed:            true,
						},
						"hop_type": schema.StringAttribute{
							MarkdownDescription: "Type of dynamic path to be computed",
							Computed:            true,
						},
						"segment_list_name": schema.StringAttribute{
							MarkdownDescription: "Segment-list name",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SegmentRoutingCandidatePathsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *SegmentRoutingCandidatePathsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SegmentRoutingCandidatePaths

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

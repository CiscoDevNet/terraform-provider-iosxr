package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

type dataSourceGnmiType struct{}

func (t dataSourceGnmiType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can retrieve one or more attributes via gNMI.",

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
			"path": {
				MarkdownDescription: "A gNMI path, e.g. `openconfig-interfaces:/interfaces/interface`.",
				Type:                types.StringType,
				Required:            true,
			},
			"attributes": {
				MarkdownDescription: "Map of key-value pairs which represents the attributes and its values.",
				Type:                types.MapType{ElemType: types.StringType},
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceGnmiType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceGnmi{
		provider: provider,
	}, diags
}

type dataSourceGnmi struct {
	provider provider
}

func (d dataSourceGnmi) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state GnmiDataSource

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.Value))

	getResp, diags := d.provider.client.Get(ctx, config.Device.Value, config.Path.Value)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Path.Value = config.Path.Value
	state.Id.Value = config.Path.Value

	attributes := make(map[string]attr.Value)

	for attr, value := range gjson.ParseBytes(getResp.Notification[0].Update[0].Val.GetJsonIetfVal()).Map() {
		attributes[attr] = types.String{Value: value.String()}
	}
	state.Attributes.Elems = attributes
	state.Attributes.ElemType = types.StringType

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ASPathSet struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

type ASPathSetData struct {
	Device  types.String `tfsdk:"device"`
	Id      types.String `tfsdk:"id"`
	SetName types.String `tfsdk:"set_name"`
	Rpl     types.String `tfsdk:"rpl"`
}

func (data ASPathSet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/as-path-sets/as-path-set[set-name=%s]", data.SetName.ValueString())
}

func (data ASPathSetData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/as-path-sets/as-path-set[set-name=%s]", data.SetName.ValueString())
}

func (data ASPathSet) toBody(ctx context.Context) string {
	body := "{}"
	if !data.SetName.IsNull() && !data.SetName.IsUnknown() {
		body, _ = sjson.Set(body, "set-name", data.SetName.ValueString())
	}
	if !data.Rpl.IsNull() && !data.Rpl.IsUnknown() {
		body, _ = sjson.Set(body, "rplas-path-set", data.Rpl.ValueString())
	}
	return body
}

func (data *ASPathSet) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rplas-path-set"); value.Exists() && !data.Rpl.IsNull() {
		data.Rpl = types.StringValue(value.String())
	} else {
		data.Rpl = types.StringNull()
	}
}

func (data *ASPathSetData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rplas-path-set"); value.Exists() {
		data.Rpl = types.StringValue(value.String())
	}
}

func (data *ASPathSet) getDeletedListItems(ctx context.Context, state ASPathSet) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *ASPathSet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *ASPathSet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Rpl.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/rplas-path-set", data.getPath()))
	}
	return deletePaths
}

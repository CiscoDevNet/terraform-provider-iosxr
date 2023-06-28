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
	Device       types.String `tfsdk:"device"`
	Id           types.String `tfsdk:"id"`
	DeleteMode   types.String `tfsdk:"delete_mode"`
	SetName      types.String `tfsdk:"set_name"`
	RplasPathSet types.String `tfsdk:"rplas_path_set"`
}

type ASPathSetData struct {
	Device       types.String `tfsdk:"device"`
	Id           types.String `tfsdk:"id"`
	SetName      types.String `tfsdk:"set_name"`
	RplasPathSet types.String `tfsdk:"rplas_path_set"`
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
	if !data.RplasPathSet.IsNull() && !data.RplasPathSet.IsUnknown() {
		body, _ = sjson.Set(body, "rplas-path-set", data.RplasPathSet.ValueString())
	}
	return body
}

func (data *ASPathSet) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rplas-path-set"); value.Exists() && !data.RplasPathSet.IsNull() {
		data.RplasPathSet = types.StringValue(value.String())
	} else {
		data.RplasPathSet = types.StringNull()
	}
}

func (data *ASPathSetData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rplas-path-set"); value.Exists() {
		data.RplasPathSet = types.StringValue(value.String())
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
	if !data.RplasPathSet.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/rplas-path-set", data.getPath()))
	}
	return deletePaths
}

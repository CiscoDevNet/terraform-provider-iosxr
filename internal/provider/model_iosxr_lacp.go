// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type LACP struct {
	Device   types.String `tfsdk:"device"`
	Id       types.String `tfsdk:"id"`
	Mac      types.String `tfsdk:"mac"`
	Priority types.Int64  `tfsdk:"priority"`
}

type LACPData struct {
	Device   types.String `tfsdk:"device"`
	Id       types.String `tfsdk:"id"`
	Mac      types.String `tfsdk:"mac"`
	Priority types.Int64  `tfsdk:"priority"`
}

func (data LACP) getPath() string {
	return "Cisco-IOS-XR-um-lacp-cfg:/lacp/system"
}

func (data LACPData) getPath() string {
	return "Cisco-IOS-XR-um-lacp-cfg:/lacp/system"
}

func (data LACP) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Mac.IsNull() && !data.Mac.IsUnknown() {
		body, _ = sjson.Set(body, "mac", data.Mac.ValueString())
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		body, _ = sjson.Set(body, "priority", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	}
	return body
}

func (data *LACP) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "mac"); value.Exists() && !data.Mac.IsNull() {
		data.Mac = types.StringValue(value.String())
	} else {
		data.Mac = types.StringNull()
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
}

func (data *LACPData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "mac"); value.Exists() {
		data.Mac = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
}

func (data *LACP) getDeletedListItems(ctx context.Context, state LACP) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *LACP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *LACP) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Mac.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/mac", data.getPath()))
	}
	if !data.Priority.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/priority", data.getPath()))
	}
	return deletePaths
}

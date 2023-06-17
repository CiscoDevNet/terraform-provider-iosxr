// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SegmentRoutingTEPolicyCandidatePath struct {
	Device     types.String                                   `tfsdk:"device"`
	Id         types.String                                   `tfsdk:"id"`
	PolicyName types.String                                   `tfsdk:"policy_name"`
	PathIndex  types.Int64                                    `tfsdk:"path_index"`
	PathInfos  []SegmentRoutingTEPolicyCandidatePathPathInfos `tfsdk:"path_infos"`
}
type SegmentRoutingTEPolicyCandidatePathPathInfos struct {
	Type            types.String `tfsdk:"type"`
	Pcep            types.Bool   `tfsdk:"pcep"`
	MetricType      types.String `tfsdk:"metric_type"`
	HopType         types.String `tfsdk:"hop_type"`
	SegmentListName types.String `tfsdk:"segment_list_name"`
}

func (data SegmentRoutingTEPolicyCandidatePath) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy[policy-name=%s]/candidate-paths/preferences/preference[path-index=%d]", data.PolicyName.ValueString(), data.PathIndex.ValueInt64())
}

func (data SegmentRoutingTEPolicyCandidatePath) toBody(ctx context.Context) string {
	body := "{}"
	if !data.PathIndex.IsNull() && !data.PathIndex.IsUnknown() {
		body, _ = sjson.Set(body, "path-index", strconv.FormatInt(data.PathIndex.ValueInt64(), 10))
	}
	if len(data.PathInfos) > 0 {
		body, _ = sjson.Set(body, "path-infos.path-info", []interface{}{})
		for index, item := range data.PathInfos {
			if !item.Type.IsNull() && !item.Type.IsUnknown() {
				body, _ = sjson.Set(body, "path-infos.path-info"+"."+strconv.Itoa(index)+"."+"type", item.Type.ValueString())
			}
			if !item.Pcep.IsNull() && !item.Pcep.IsUnknown() {
				if item.Pcep.ValueBool() {
					body, _ = sjson.Set(body, "path-infos.path-info"+"."+strconv.Itoa(index)+"."+"pcep", map[string]string{})
				}
			}
			if !item.MetricType.IsNull() && !item.MetricType.IsUnknown() {
				body, _ = sjson.Set(body, "path-infos.path-info"+"."+strconv.Itoa(index)+"."+"metric.metric-type", item.MetricType.ValueString())
			}
			if !item.HopType.IsNull() && !item.HopType.IsUnknown() {
				body, _ = sjson.Set(body, "path-infos.path-info"+"."+strconv.Itoa(index)+"."+"hop-type", item.HopType.ValueString())
			}
			if !item.SegmentListName.IsNull() && !item.SegmentListName.IsUnknown() {
				body, _ = sjson.Set(body, "path-infos.path-info"+"."+strconv.Itoa(index)+"."+"segment-list-name", item.SegmentListName.ValueString())
			}
		}
	}
	return body
}

func (data *SegmentRoutingTEPolicyCandidatePath) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.PathInfos {
		keys := [...]string{"type", "hop-type", "segment-list-name"}
		keyValues := [...]string{data.PathInfos[i].Type.ValueString(), data.PathInfos[i].HopType.ValueString(), data.PathInfos[i].SegmentListName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "path-infos.path-info").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("type"); value.Exists() && !data.PathInfos[i].Type.IsNull() {
			data.PathInfos[i].Type = types.StringValue(value.String())
		} else {
			data.PathInfos[i].Type = types.StringNull()
		}
		if value := r.Get("pcep"); !data.PathInfos[i].Pcep.IsNull() {
			if value.Exists() {
				data.PathInfos[i].Pcep = types.BoolValue(true)
			} else {
				data.PathInfos[i].Pcep = types.BoolValue(false)
			}
		} else {
			data.PathInfos[i].Pcep = types.BoolNull()
		}
		if value := r.Get("metric.metric-type"); value.Exists() && !data.PathInfos[i].MetricType.IsNull() {
			data.PathInfos[i].MetricType = types.StringValue(value.String())
		} else {
			data.PathInfos[i].MetricType = types.StringNull()
		}
		if value := r.Get("hop-type"); value.Exists() && !data.PathInfos[i].HopType.IsNull() {
			data.PathInfos[i].HopType = types.StringValue(value.String())
		} else {
			data.PathInfos[i].HopType = types.StringNull()
		}
		if value := r.Get("segment-list-name"); value.Exists() && !data.PathInfos[i].SegmentListName.IsNull() {
			data.PathInfos[i].SegmentListName = types.StringValue(value.String())
		} else {
			data.PathInfos[i].SegmentListName = types.StringNull()
		}
	}
}

func (data *SegmentRoutingTEPolicyCandidatePath) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "path-infos.path-info"); value.Exists() {
		data.PathInfos = make([]SegmentRoutingTEPolicyCandidatePathPathInfos, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SegmentRoutingTEPolicyCandidatePathPathInfos{}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pcep"); cValue.Exists() {
				item.Pcep = types.BoolValue(true)
			} else {
				item.Pcep = types.BoolValue(false)
			}
			if cValue := v.Get("metric.metric-type"); cValue.Exists() {
				item.MetricType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("hop-type"); cValue.Exists() {
				item.HopType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("segment-list-name"); cValue.Exists() {
				item.SegmentListName = types.StringValue(cValue.String())
			}
			data.PathInfos = append(data.PathInfos, item)
			return true
		})
	}
}

func (data *SegmentRoutingTEPolicyCandidatePath) getDeletedListItems(ctx context.Context, state SegmentRoutingTEPolicyCandidatePath) []string {
	deletedListItems := make([]string, 0)
	for i := range state.PathInfos {
		keys := [...]string{"type", "hop-type", "segment-list-name"}
		stateKeyValues := [...]string{state.PathInfos[i].Type.ValueString(), state.PathInfos[i].HopType.ValueString(), state.PathInfos[i].SegmentListName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.PathInfos[i].Type.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.PathInfos[i].HopType.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.PathInfos[i].SegmentListName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.PathInfos {
			found = true
			if state.PathInfos[i].Type.ValueString() != data.PathInfos[j].Type.ValueString() {
				found = false
			}
			if state.PathInfos[i].HopType.ValueString() != data.PathInfos[j].HopType.ValueString() {
				found = false
			}
			if state.PathInfos[i].SegmentListName.ValueString() != data.PathInfos[j].SegmentListName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/path-infos/path-info%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *SegmentRoutingTEPolicyCandidatePath) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.PathInfos {
		keys := [...]string{"type", "hop-type", "segment-list-name"}
		keyValues := [...]string{data.PathInfos[i].Type.ValueString(), data.PathInfos[i].HopType.ValueString(), data.PathInfos[i].SegmentListName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.PathInfos[i].Pcep.IsNull() && !data.PathInfos[i].Pcep.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/path-infos/path-info%v/pcep", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

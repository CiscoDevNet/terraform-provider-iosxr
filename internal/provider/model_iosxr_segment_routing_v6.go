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

type SegmentRoutingV6 struct {
	Device                     types.String               `tfsdk:"device"`
	Id                         types.String               `tfsdk:"id"`
	Enable                     types.Bool                 `tfsdk:"enable"`
	EncapsulationSourceAddress types.String               `tfsdk:"encapsulation_source_address"`
	Locators                   []SegmentRoutingV6Locators `tfsdk:"locators"`
}
type SegmentRoutingV6Locators struct {
	LocatorEnable        types.Bool   `tfsdk:"locator_enable"`
	Name                 types.String `tfsdk:"name"`
	MicroSegmentBehavior types.String `tfsdk:"micro_segment_behavior"`
	Prefix               types.String `tfsdk:"prefix"`
	PrefixLength         types.Int64  `tfsdk:"prefix_length"`
}

func (data SegmentRoutingV6) getPath() string {
	return "Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6"
}

func (data SegmentRoutingV6) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Enable.IsNull() && !data.Enable.IsUnknown() {
		if data.Enable.ValueBool() {
			body, _ = sjson.Set(body, "enable", map[string]string{})
		}
	}
	if !data.EncapsulationSourceAddress.IsNull() && !data.EncapsulationSourceAddress.IsUnknown() {
		body, _ = sjson.Set(body, "encapsulation.source-address", data.EncapsulationSourceAddress.ValueString())
	}
	if len(data.Locators) > 0 {
		body, _ = sjson.Set(body, "locators.locators.locator", []interface{}{})
		for index, item := range data.Locators {
			if !item.LocatorEnable.IsNull() && !item.LocatorEnable.IsUnknown() {
				if item.LocatorEnable.ValueBool() {
					body, _ = sjson.Set(body, "locators.locators.locator"+"."+strconv.Itoa(index)+"."+"locator-enable", map[string]string{})
				}
			}
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "locators.locators.locator"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.MicroSegmentBehavior.IsNull() && !item.MicroSegmentBehavior.IsUnknown() {
				body, _ = sjson.Set(body, "locators.locators.locator"+"."+strconv.Itoa(index)+"."+"micro-segment.behavior", item.MicroSegmentBehavior.ValueString())
			}
			if !item.Prefix.IsNull() && !item.Prefix.IsUnknown() {
				body, _ = sjson.Set(body, "locators.locators.locator"+"."+strconv.Itoa(index)+"."+"prefix.prefix", item.Prefix.ValueString())
			}
			if !item.PrefixLength.IsNull() && !item.PrefixLength.IsUnknown() {
				body, _ = sjson.Set(body, "locators.locators.locator"+"."+strconv.Itoa(index)+"."+"prefix.prefix-length", strconv.FormatInt(item.PrefixLength.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *SegmentRoutingV6) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "enable"); !data.Enable.IsNull() {
		if value.Exists() {
			data.Enable = types.BoolValue(true)
		} else {
			data.Enable = types.BoolValue(false)
		}
	} else {
		data.Enable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.source-address"); value.Exists() && !data.EncapsulationSourceAddress.IsNull() {
		data.EncapsulationSourceAddress = types.StringValue(value.String())
	} else {
		data.EncapsulationSourceAddress = types.StringNull()
	}
	for i := range data.Locators {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Locators[i].Name.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "locators.locators.locator").ForEach(
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
		if value := r.Get("locator-enable"); !data.Locators[i].LocatorEnable.IsNull() {
			if value.Exists() {
				data.Locators[i].LocatorEnable = types.BoolValue(true)
			} else {
				data.Locators[i].LocatorEnable = types.BoolValue(false)
			}
		} else {
			data.Locators[i].LocatorEnable = types.BoolNull()
		}
		if value := r.Get("name"); value.Exists() && !data.Locators[i].Name.IsNull() {
			data.Locators[i].Name = types.StringValue(value.String())
		} else {
			data.Locators[i].Name = types.StringNull()
		}
		if value := r.Get("micro-segment.behavior"); value.Exists() && !data.Locators[i].MicroSegmentBehavior.IsNull() {
			data.Locators[i].MicroSegmentBehavior = types.StringValue(value.String())
		} else {
			data.Locators[i].MicroSegmentBehavior = types.StringNull()
		}
		if value := r.Get("prefix.prefix"); value.Exists() && !data.Locators[i].Prefix.IsNull() {
			data.Locators[i].Prefix = types.StringValue(value.String())
		} else {
			data.Locators[i].Prefix = types.StringNull()
		}
		if value := r.Get("prefix.prefix-length"); value.Exists() && !data.Locators[i].PrefixLength.IsNull() {
			data.Locators[i].PrefixLength = types.Int64Value(value.Int())
		} else {
			data.Locators[i].PrefixLength = types.Int64Null()
		}
	}
}

func (data *SegmentRoutingV6) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "enable"); value.Exists() {
		data.Enable = types.BoolValue(true)
	} else {
		data.Enable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.source-address"); value.Exists() {
		data.EncapsulationSourceAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "locators.locators.locator"); value.Exists() {
		data.Locators = make([]SegmentRoutingV6Locators, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SegmentRoutingV6Locators{}
			if cValue := v.Get("locator-enable"); cValue.Exists() {
				item.LocatorEnable = types.BoolValue(true)
			} else {
				item.LocatorEnable = types.BoolValue(false)
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("micro-segment.behavior"); cValue.Exists() {
				item.MicroSegmentBehavior = types.StringValue(cValue.String())
			}
			if cValue := v.Get("prefix.prefix"); cValue.Exists() {
				item.Prefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("prefix.prefix-length"); cValue.Exists() {
				item.PrefixLength = types.Int64Value(cValue.Int())
			}
			data.Locators = append(data.Locators, item)
			return true
		})
	}
}

func (data *SegmentRoutingV6) getDeletedListItems(ctx context.Context, state SegmentRoutingV6) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Locators {
		keys := [...]string{"name"}
		stateKeyValues := [...]string{state.Locators[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Locators[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Locators {
			found = true
			if state.Locators[i].Name.ValueString() != data.Locators[j].Name.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/locators/locators/locator%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *SegmentRoutingV6) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Enable.IsNull() && !data.Enable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/enable", data.getPath()))
	}

	for i := range data.Locators {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Locators[i].Name.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.Locators[i].LocatorEnable.IsNull() && !data.Locators[i].LocatorEnable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/locators/locators/locator%v/locator-enable", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

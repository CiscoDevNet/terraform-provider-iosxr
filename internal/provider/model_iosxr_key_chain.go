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

type KeyChain struct {
	Device types.String   `tfsdk:"device"`
	Id     types.String   `tfsdk:"id"`
	Name   types.String   `tfsdk:"name"`
	Keys   []KeyChainKeys `tfsdk:"keys"`
}
type KeyChainKeys struct {
	KeyName                           types.String `tfsdk:"key_name"`
	KeyStringPassword                 types.String `tfsdk:"key_string_password"`
	CryptographicAlgorithm            types.String `tfsdk:"cryptographic_algorithm"`
	AcceptLifetimeStartTimeHour       types.Int64  `tfsdk:"accept_lifetime_start_time_hour"`
	AcceptLifetimeStartTimeMinute     types.Int64  `tfsdk:"accept_lifetime_start_time_minute"`
	AcceptLifetimeStartTimeSecond     types.Int64  `tfsdk:"accept_lifetime_start_time_second"`
	AcceptLifetimeStartTimeDayOfMonth types.Int64  `tfsdk:"accept_lifetime_start_time_day_of_month"`
	AcceptLifetimeStartTimeMonth      types.String `tfsdk:"accept_lifetime_start_time_month"`
	AcceptLifetimeStartTimeYear       types.Int64  `tfsdk:"accept_lifetime_start_time_year"`
	AcceptLifetimeInfinite            types.Bool   `tfsdk:"accept_lifetime_infinite"`
	SendLifetimeStartTimeHour         types.Int64  `tfsdk:"send_lifetime_start_time_hour"`
	SendLifetimeStartTimeMinute       types.Int64  `tfsdk:"send_lifetime_start_time_minute"`
	SendLifetimeStartTimeSecond       types.Int64  `tfsdk:"send_lifetime_start_time_second"`
	SendLifetimeStartTimeDayOfMonth   types.Int64  `tfsdk:"send_lifetime_start_time_day_of_month"`
	SendLifetimeStartTimeMonth        types.String `tfsdk:"send_lifetime_start_time_month"`
	SendLifetimeStartTimeYear         types.Int64  `tfsdk:"send_lifetime_start_time_year"`
	SendLifetimeInfinite              types.Bool   `tfsdk:"send_lifetime_infinite"`
}

func (data KeyChain) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-key-chain-cfg:/key/chains/chain[key-chain-name=%s]", data.Name.ValueString())
}

func (data KeyChain) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "key-chain-name", data.Name.ValueString())
	}
	if len(data.Keys) > 0 {
		body, _ = sjson.Set(body, "keys.key", []interface{}{})
		for index, item := range data.Keys {
			if !item.KeyName.IsNull() && !item.KeyName.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"key-name", item.KeyName.ValueString())
			}
			if !item.KeyStringPassword.IsNull() && !item.KeyStringPassword.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"key-string.password", item.KeyStringPassword.ValueString())
			}
			if !item.CryptographicAlgorithm.IsNull() && !item.CryptographicAlgorithm.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"cryptographic-algorithm", item.CryptographicAlgorithm.ValueString())
			}
			if !item.AcceptLifetimeStartTimeHour.IsNull() && !item.AcceptLifetimeStartTimeHour.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"accept-lifetime.start-time.hour", strconv.FormatInt(item.AcceptLifetimeStartTimeHour.ValueInt64(), 10))
			}
			if !item.AcceptLifetimeStartTimeMinute.IsNull() && !item.AcceptLifetimeStartTimeMinute.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"accept-lifetime.start-time.minute", strconv.FormatInt(item.AcceptLifetimeStartTimeMinute.ValueInt64(), 10))
			}
			if !item.AcceptLifetimeStartTimeSecond.IsNull() && !item.AcceptLifetimeStartTimeSecond.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"accept-lifetime.start-time.second", strconv.FormatInt(item.AcceptLifetimeStartTimeSecond.ValueInt64(), 10))
			}
			if !item.AcceptLifetimeStartTimeDayOfMonth.IsNull() && !item.AcceptLifetimeStartTimeDayOfMonth.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"accept-lifetime.start-time.day-of-month", strconv.FormatInt(item.AcceptLifetimeStartTimeDayOfMonth.ValueInt64(), 10))
			}
			if !item.AcceptLifetimeStartTimeMonth.IsNull() && !item.AcceptLifetimeStartTimeMonth.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"accept-lifetime.start-time.month", item.AcceptLifetimeStartTimeMonth.ValueString())
			}
			if !item.AcceptLifetimeStartTimeYear.IsNull() && !item.AcceptLifetimeStartTimeYear.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"accept-lifetime.start-time.year", strconv.FormatInt(item.AcceptLifetimeStartTimeYear.ValueInt64(), 10))
			}
			if !item.AcceptLifetimeInfinite.IsNull() && !item.AcceptLifetimeInfinite.IsUnknown() {
				if item.AcceptLifetimeInfinite.ValueBool() {
					body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"accept-lifetime.infinite", map[string]string{})
				}
			}
			if !item.SendLifetimeStartTimeHour.IsNull() && !item.SendLifetimeStartTimeHour.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"send-lifetime.start-time.hour", strconv.FormatInt(item.SendLifetimeStartTimeHour.ValueInt64(), 10))
			}
			if !item.SendLifetimeStartTimeMinute.IsNull() && !item.SendLifetimeStartTimeMinute.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"send-lifetime.start-time.minute", strconv.FormatInt(item.SendLifetimeStartTimeMinute.ValueInt64(), 10))
			}
			if !item.SendLifetimeStartTimeSecond.IsNull() && !item.SendLifetimeStartTimeSecond.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"send-lifetime.start-time.second", strconv.FormatInt(item.SendLifetimeStartTimeSecond.ValueInt64(), 10))
			}
			if !item.SendLifetimeStartTimeDayOfMonth.IsNull() && !item.SendLifetimeStartTimeDayOfMonth.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"send-lifetime.start-time.day-of-month", strconv.FormatInt(item.SendLifetimeStartTimeDayOfMonth.ValueInt64(), 10))
			}
			if !item.SendLifetimeStartTimeMonth.IsNull() && !item.SendLifetimeStartTimeMonth.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"send-lifetime.start-time.month", item.SendLifetimeStartTimeMonth.ValueString())
			}
			if !item.SendLifetimeStartTimeYear.IsNull() && !item.SendLifetimeStartTimeYear.IsUnknown() {
				body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"send-lifetime.start-time.year", strconv.FormatInt(item.SendLifetimeStartTimeYear.ValueInt64(), 10))
			}
			if !item.SendLifetimeInfinite.IsNull() && !item.SendLifetimeInfinite.IsUnknown() {
				if item.SendLifetimeInfinite.ValueBool() {
					body, _ = sjson.Set(body, "keys.key"+"."+strconv.Itoa(index)+"."+"send-lifetime.infinite", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *KeyChain) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.Keys {
		keys := [...]string{"key-name"}
		keyValues := [...]string{data.Keys[i].KeyName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "keys.key").ForEach(
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
		if value := r.Get("key-name"); value.Exists() && !data.Keys[i].KeyName.IsNull() {
			data.Keys[i].KeyName = types.StringValue(value.String())
		} else {
			data.Keys[i].KeyName = types.StringNull()
		}
		if value := r.Get("key-string.password"); value.Exists() && !data.Keys[i].KeyStringPassword.IsNull() {
			data.Keys[i].KeyStringPassword = types.StringValue(value.String())
		} else {
			data.Keys[i].KeyStringPassword = types.StringNull()
		}
		if value := r.Get("cryptographic-algorithm"); value.Exists() && !data.Keys[i].CryptographicAlgorithm.IsNull() {
			data.Keys[i].CryptographicAlgorithm = types.StringValue(value.String())
		} else {
			data.Keys[i].CryptographicAlgorithm = types.StringNull()
		}
		if value := r.Get("accept-lifetime.start-time.hour"); value.Exists() && !data.Keys[i].AcceptLifetimeStartTimeHour.IsNull() {
			data.Keys[i].AcceptLifetimeStartTimeHour = types.Int64Value(value.Int())
		} else {
			data.Keys[i].AcceptLifetimeStartTimeHour = types.Int64Null()
		}
		if value := r.Get("accept-lifetime.start-time.minute"); value.Exists() && !data.Keys[i].AcceptLifetimeStartTimeMinute.IsNull() {
			data.Keys[i].AcceptLifetimeStartTimeMinute = types.Int64Value(value.Int())
		} else {
			data.Keys[i].AcceptLifetimeStartTimeMinute = types.Int64Null()
		}
		if value := r.Get("accept-lifetime.start-time.second"); value.Exists() && !data.Keys[i].AcceptLifetimeStartTimeSecond.IsNull() {
			data.Keys[i].AcceptLifetimeStartTimeSecond = types.Int64Value(value.Int())
		} else {
			data.Keys[i].AcceptLifetimeStartTimeSecond = types.Int64Null()
		}
		if value := r.Get("accept-lifetime.start-time.day-of-month"); value.Exists() && !data.Keys[i].AcceptLifetimeStartTimeDayOfMonth.IsNull() {
			data.Keys[i].AcceptLifetimeStartTimeDayOfMonth = types.Int64Value(value.Int())
		} else {
			data.Keys[i].AcceptLifetimeStartTimeDayOfMonth = types.Int64Null()
		}
		if value := r.Get("accept-lifetime.start-time.month"); value.Exists() && !data.Keys[i].AcceptLifetimeStartTimeMonth.IsNull() {
			data.Keys[i].AcceptLifetimeStartTimeMonth = types.StringValue(value.String())
		} else {
			data.Keys[i].AcceptLifetimeStartTimeMonth = types.StringNull()
		}
		if value := r.Get("accept-lifetime.start-time.year"); value.Exists() && !data.Keys[i].AcceptLifetimeStartTimeYear.IsNull() {
			data.Keys[i].AcceptLifetimeStartTimeYear = types.Int64Value(value.Int())
		} else {
			data.Keys[i].AcceptLifetimeStartTimeYear = types.Int64Null()
		}
		if value := r.Get("accept-lifetime.infinite"); !data.Keys[i].AcceptLifetimeInfinite.IsNull() {
			if value.Exists() {
				data.Keys[i].AcceptLifetimeInfinite = types.BoolValue(true)
			} else {
				data.Keys[i].AcceptLifetimeInfinite = types.BoolValue(false)
			}
		} else {
			data.Keys[i].AcceptLifetimeInfinite = types.BoolNull()
		}
		if value := r.Get("send-lifetime.start-time.hour"); value.Exists() && !data.Keys[i].SendLifetimeStartTimeHour.IsNull() {
			data.Keys[i].SendLifetimeStartTimeHour = types.Int64Value(value.Int())
		} else {
			data.Keys[i].SendLifetimeStartTimeHour = types.Int64Null()
		}
		if value := r.Get("send-lifetime.start-time.minute"); value.Exists() && !data.Keys[i].SendLifetimeStartTimeMinute.IsNull() {
			data.Keys[i].SendLifetimeStartTimeMinute = types.Int64Value(value.Int())
		} else {
			data.Keys[i].SendLifetimeStartTimeMinute = types.Int64Null()
		}
		if value := r.Get("send-lifetime.start-time.second"); value.Exists() && !data.Keys[i].SendLifetimeStartTimeSecond.IsNull() {
			data.Keys[i].SendLifetimeStartTimeSecond = types.Int64Value(value.Int())
		} else {
			data.Keys[i].SendLifetimeStartTimeSecond = types.Int64Null()
		}
		if value := r.Get("send-lifetime.start-time.day-of-month"); value.Exists() && !data.Keys[i].SendLifetimeStartTimeDayOfMonth.IsNull() {
			data.Keys[i].SendLifetimeStartTimeDayOfMonth = types.Int64Value(value.Int())
		} else {
			data.Keys[i].SendLifetimeStartTimeDayOfMonth = types.Int64Null()
		}
		if value := r.Get("send-lifetime.start-time.month"); value.Exists() && !data.Keys[i].SendLifetimeStartTimeMonth.IsNull() {
			data.Keys[i].SendLifetimeStartTimeMonth = types.StringValue(value.String())
		} else {
			data.Keys[i].SendLifetimeStartTimeMonth = types.StringNull()
		}
		if value := r.Get("send-lifetime.start-time.year"); value.Exists() && !data.Keys[i].SendLifetimeStartTimeYear.IsNull() {
			data.Keys[i].SendLifetimeStartTimeYear = types.Int64Value(value.Int())
		} else {
			data.Keys[i].SendLifetimeStartTimeYear = types.Int64Null()
		}
		if value := r.Get("send-lifetime.infinite"); !data.Keys[i].SendLifetimeInfinite.IsNull() {
			if value.Exists() {
				data.Keys[i].SendLifetimeInfinite = types.BoolValue(true)
			} else {
				data.Keys[i].SendLifetimeInfinite = types.BoolValue(false)
			}
		} else {
			data.Keys[i].SendLifetimeInfinite = types.BoolNull()
		}
	}
}

func (data *KeyChain) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "keys.key"); value.Exists() {
		data.Keys = make([]KeyChainKeys, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := KeyChainKeys{}
			if cValue := v.Get("key-name"); cValue.Exists() {
				item.KeyName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("key-string.password"); cValue.Exists() {
				item.KeyStringPassword = types.StringValue(cValue.String())
			}
			if cValue := v.Get("cryptographic-algorithm"); cValue.Exists() {
				item.CryptographicAlgorithm = types.StringValue(cValue.String())
			}
			if cValue := v.Get("accept-lifetime.start-time.hour"); cValue.Exists() {
				item.AcceptLifetimeStartTimeHour = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("accept-lifetime.start-time.minute"); cValue.Exists() {
				item.AcceptLifetimeStartTimeMinute = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("accept-lifetime.start-time.second"); cValue.Exists() {
				item.AcceptLifetimeStartTimeSecond = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("accept-lifetime.start-time.day-of-month"); cValue.Exists() {
				item.AcceptLifetimeStartTimeDayOfMonth = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("accept-lifetime.start-time.month"); cValue.Exists() {
				item.AcceptLifetimeStartTimeMonth = types.StringValue(cValue.String())
			}
			if cValue := v.Get("accept-lifetime.start-time.year"); cValue.Exists() {
				item.AcceptLifetimeStartTimeYear = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("accept-lifetime.infinite"); cValue.Exists() {
				item.AcceptLifetimeInfinite = types.BoolValue(true)
			} else {
				item.AcceptLifetimeInfinite = types.BoolValue(false)
			}
			if cValue := v.Get("send-lifetime.start-time.hour"); cValue.Exists() {
				item.SendLifetimeStartTimeHour = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("send-lifetime.start-time.minute"); cValue.Exists() {
				item.SendLifetimeStartTimeMinute = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("send-lifetime.start-time.second"); cValue.Exists() {
				item.SendLifetimeStartTimeSecond = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("send-lifetime.start-time.day-of-month"); cValue.Exists() {
				item.SendLifetimeStartTimeDayOfMonth = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("send-lifetime.start-time.month"); cValue.Exists() {
				item.SendLifetimeStartTimeMonth = types.StringValue(cValue.String())
			}
			if cValue := v.Get("send-lifetime.start-time.year"); cValue.Exists() {
				item.SendLifetimeStartTimeYear = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("send-lifetime.infinite"); cValue.Exists() {
				item.SendLifetimeInfinite = types.BoolValue(true)
			} else {
				item.SendLifetimeInfinite = types.BoolValue(false)
			}
			data.Keys = append(data.Keys, item)
			return true
		})
	}
}

func (data *KeyChain) fromPlan(ctx context.Context, plan KeyChain) {
	data.Device = plan.Device
	data.Name = types.StringValue(plan.Name.ValueString())
}

func (data *KeyChain) getDeletedListItems(ctx context.Context, state KeyChain) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Keys {
		keys := [...]string{"key-name"}
		stateKeyValues := [...]string{state.Keys[i].KeyName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Keys[i].KeyName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Keys {
			found = true
			if state.Keys[i].KeyName.ValueString() != data.Keys[j].KeyName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/keys/key%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *KeyChain) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}
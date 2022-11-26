// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouterBGPVRFAddressFamily struct {
	Device                      types.String                                  `tfsdk:"device"`
	Id                          types.String                                  `tfsdk:"id"`
	AsNumber                    types.String                                  `tfsdk:"as_number"`
	VrfName                     types.String                                  `tfsdk:"vrf_name"`
	AfName                      types.String                                  `tfsdk:"af_name"`
	MaximumPathsEbgpMultipath   types.Int64                                   `tfsdk:"maximum_paths_ebgp_multipath"`
	MaximumPathsEibgpMultipath  types.Int64                                   `tfsdk:"maximum_paths_eibgp_multipath"`
	MaximumPathsIbgpMultipath   types.Int64                                   `tfsdk:"maximum_paths_ibgp_multipath"`
	LabelModePerCe              types.Bool                                    `tfsdk:"label_mode_per_ce"`
	LabelModePerVrf             types.Bool                                    `tfsdk:"label_mode_per_vrf"`
	RedistributeConnected       types.Bool                                    `tfsdk:"redistribute_connected"`
	RedistributeConnectedMetric types.Int64                                   `tfsdk:"redistribute_connected_metric"`
	RedistributeStatic          types.Bool                                    `tfsdk:"redistribute_static"`
	RedistributeStaticMetric    types.Int64                                   `tfsdk:"redistribute_static_metric"`
	AggregateAddresses          []RouterBGPVRFAddressFamilyAggregateAddresses `tfsdk:"aggregate_addresses"`
	Networks                    []RouterBGPVRFAddressFamilyNetworks           `tfsdk:"networks"`
	RedistributeOspf            []RouterBGPVRFAddressFamilyRedistributeOspf   `tfsdk:"redistribute_ospf"`
}
type RouterBGPVRFAddressFamilyAggregateAddresses struct {
	Address     types.String `tfsdk:"address"`
	Masklength  types.Int64  `tfsdk:"masklength"`
	AsSet       types.Bool   `tfsdk:"as_set"`
	AsConfedSet types.Bool   `tfsdk:"as_confed_set"`
	SummaryOnly types.Bool   `tfsdk:"summary_only"`
}
type RouterBGPVRFAddressFamilyNetworks struct {
	Address    types.String `tfsdk:"address"`
	Masklength types.Int64  `tfsdk:"masklength"`
}
type RouterBGPVRFAddressFamilyRedistributeOspf struct {
	RouterTag                 types.String `tfsdk:"router_tag"`
	MatchInternal             types.Bool   `tfsdk:"match_internal"`
	MatchInternalExternal     types.Bool   `tfsdk:"match_internal_external"`
	MatchInternalNssaExternal types.Bool   `tfsdk:"match_internal_nssa_external"`
	MatchExternal             types.Bool   `tfsdk:"match_external"`
	MatchExternalNssaExternal types.Bool   `tfsdk:"match_external_nssa_external"`
	MatchNssaExternal         types.Bool   `tfsdk:"match_nssa_external"`
	Metric                    types.Int64  `tfsdk:"metric"`
}

func (data RouterBGPVRFAddressFamily) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/vrfs/vrf[vrf-name=%s]/address-families/address-family[af-name=%s]", data.AsNumber.ValueString(), data.VrfName.ValueString(), data.AfName.ValueString())
}

func (data RouterBGPVRFAddressFamily) toBody() string {
	body := "{}"
	if !data.MaximumPathsEbgpMultipath.IsNull() && !data.MaximumPathsEbgpMultipath.IsUnknown() {
		body, _ = sjson.Set(body, "maximum-paths.ebgp.multipath", strconv.FormatInt(data.MaximumPathsEbgpMultipath.ValueInt64(), 10))
	}
	if !data.MaximumPathsEibgpMultipath.IsNull() && !data.MaximumPathsEibgpMultipath.IsUnknown() {
		body, _ = sjson.Set(body, "maximum-paths.eibgp.multipath", strconv.FormatInt(data.MaximumPathsEibgpMultipath.ValueInt64(), 10))
	}
	if !data.MaximumPathsIbgpMultipath.IsNull() && !data.MaximumPathsIbgpMultipath.IsUnknown() {
		body, _ = sjson.Set(body, "maximum-paths.ibgp.multipath", strconv.FormatInt(data.MaximumPathsIbgpMultipath.ValueInt64(), 10))
	}
	if !data.LabelModePerCe.IsNull() && !data.LabelModePerCe.IsUnknown() {
		if data.LabelModePerCe.ValueBool() {
			body, _ = sjson.Set(body, "label.mode.per-ce", map[string]string{})
		}
	}
	if !data.LabelModePerVrf.IsNull() && !data.LabelModePerVrf.IsUnknown() {
		if data.LabelModePerVrf.ValueBool() {
			body, _ = sjson.Set(body, "label.mode.per-vrf", map[string]string{})
		}
	}
	if !data.RedistributeConnected.IsNull() && !data.RedistributeConnected.IsUnknown() {
		if data.RedistributeConnected.ValueBool() {
			body, _ = sjson.Set(body, "redistribute.connected", map[string]string{})
		}
	}
	if !data.RedistributeConnectedMetric.IsNull() && !data.RedistributeConnectedMetric.IsUnknown() {
		body, _ = sjson.Set(body, "redistribute.connected.metric", strconv.FormatInt(data.RedistributeConnectedMetric.ValueInt64(), 10))
	}
	if !data.RedistributeStatic.IsNull() && !data.RedistributeStatic.IsUnknown() {
		if data.RedistributeStatic.ValueBool() {
			body, _ = sjson.Set(body, "redistribute.static", map[string]string{})
		}
	}
	if !data.RedistributeStaticMetric.IsNull() && !data.RedistributeStaticMetric.IsUnknown() {
		body, _ = sjson.Set(body, "redistribute.static.metric", strconv.FormatInt(data.RedistributeStaticMetric.ValueInt64(), 10))
	}
	if len(data.AggregateAddresses) > 0 {
		body, _ = sjson.Set(body, "aggregate-addresses.aggregate-address", []interface{}{})
		for index, item := range data.AggregateAddresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "aggregate-addresses.aggregate-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Masklength.IsNull() && !item.Masklength.IsUnknown() {
				body, _ = sjson.Set(body, "aggregate-addresses.aggregate-address"+"."+strconv.Itoa(index)+"."+"masklength", strconv.FormatInt(item.Masklength.ValueInt64(), 10))
			}
			if !item.AsSet.IsNull() && !item.AsSet.IsUnknown() {
				if item.AsSet.ValueBool() {
					body, _ = sjson.Set(body, "aggregate-addresses.aggregate-address"+"."+strconv.Itoa(index)+"."+"as-set", map[string]string{})
				}
			}
			if !item.AsConfedSet.IsNull() && !item.AsConfedSet.IsUnknown() {
				if item.AsConfedSet.ValueBool() {
					body, _ = sjson.Set(body, "aggregate-addresses.aggregate-address"+"."+strconv.Itoa(index)+"."+"as-confed-set", map[string]string{})
				}
			}
			if !item.SummaryOnly.IsNull() && !item.SummaryOnly.IsUnknown() {
				if item.SummaryOnly.ValueBool() {
					body, _ = sjson.Set(body, "aggregate-addresses.aggregate-address"+"."+strconv.Itoa(index)+"."+"summary-only", map[string]string{})
				}
			}
		}
	}
	if len(data.Networks) > 0 {
		body, _ = sjson.Set(body, "networks.network", []interface{}{})
		for index, item := range data.Networks {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "networks.network"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Masklength.IsNull() && !item.Masklength.IsUnknown() {
				body, _ = sjson.Set(body, "networks.network"+"."+strconv.Itoa(index)+"."+"masklength", strconv.FormatInt(item.Masklength.ValueInt64(), 10))
			}
		}
	}
	if len(data.RedistributeOspf) > 0 {
		body, _ = sjson.Set(body, "redistribute.ospf", []interface{}{})
		for index, item := range data.RedistributeOspf {
			if !item.RouterTag.IsNull() && !item.RouterTag.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"router-tag", item.RouterTag.ValueString())
			}
			if !item.MatchInternal.IsNull() && !item.MatchInternal.IsUnknown() {
				if item.MatchInternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.internal", map[string]string{})
				}
			}
			if !item.MatchInternalExternal.IsNull() && !item.MatchInternalExternal.IsUnknown() {
				if item.MatchInternalExternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.internal.external", map[string]string{})
				}
			}
			if !item.MatchInternalNssaExternal.IsNull() && !item.MatchInternalNssaExternal.IsUnknown() {
				if item.MatchInternalNssaExternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.internal.nssa-external", map[string]string{})
				}
			}
			if !item.MatchExternal.IsNull() && !item.MatchExternal.IsUnknown() {
				if item.MatchExternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.external", map[string]string{})
				}
			}
			if !item.MatchExternalNssaExternal.IsNull() && !item.MatchExternalNssaExternal.IsUnknown() {
				if item.MatchExternalNssaExternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.external.nssa-external", map[string]string{})
				}
			}
			if !item.MatchNssaExternal.IsNull() && !item.MatchNssaExternal.IsUnknown() {
				if item.MatchNssaExternal.ValueBool() {
					body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"match.nssa-external", map[string]string{})
				}
			}
			if !item.Metric.IsNull() && !item.Metric.IsUnknown() {
				body, _ = sjson.Set(body, "redistribute.ospf"+"."+strconv.Itoa(index)+"."+"metric", strconv.FormatInt(item.Metric.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *RouterBGPVRFAddressFamily) updateFromBody(res []byte) {
	if value := gjson.GetBytes(res, "maximum-paths.ebgp.multipath"); value.Exists() {
		data.MaximumPathsEbgpMultipath = types.Int64Value(value.Int())
	} else {
		data.MaximumPathsEbgpMultipath = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "maximum-paths.eibgp.multipath"); value.Exists() {
		data.MaximumPathsEibgpMultipath = types.Int64Value(value.Int())
	} else {
		data.MaximumPathsEibgpMultipath = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "maximum-paths.ibgp.multipath"); value.Exists() {
		data.MaximumPathsIbgpMultipath = types.Int64Value(value.Int())
	} else {
		data.MaximumPathsIbgpMultipath = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "label.mode.per-ce"); value.Exists() {
		data.LabelModePerCe = types.BoolValue(true)
	} else {
		data.LabelModePerCe = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "label.mode.per-vrf"); value.Exists() {
		data.LabelModePerVrf = types.BoolValue(true)
	} else {
		data.LabelModePerVrf = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.connected"); value.Exists() {
		data.RedistributeConnected = types.BoolValue(true)
	} else {
		data.RedistributeConnected = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.connected.metric"); value.Exists() {
		data.RedistributeConnectedMetric = types.Int64Value(value.Int())
	} else {
		data.RedistributeConnectedMetric = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "redistribute.static"); value.Exists() {
		data.RedistributeStatic = types.BoolValue(true)
	} else {
		data.RedistributeStatic = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.static.metric"); value.Exists() {
		data.RedistributeStaticMetric = types.Int64Value(value.Int())
	} else {
		data.RedistributeStaticMetric = types.Int64Null()
	}
	for i := range data.AggregateAddresses {
		keys := [...]string{"address", "masklength"}
		keyValues := [...]string{data.AggregateAddresses[i].Address.ValueString(), strconv.FormatInt(data.AggregateAddresses[i].Masklength.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "aggregate-addresses.aggregate-address").ForEach(
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
		if value := r.Get("address"); value.Exists() {
			data.AggregateAddresses[i].Address = types.StringValue(value.String())
		} else {
			data.AggregateAddresses[i].Address = types.StringNull()
		}
		if value := r.Get("masklength"); value.Exists() {
			data.AggregateAddresses[i].Masklength = types.Int64Value(value.Int())
		} else {
			data.AggregateAddresses[i].Masklength = types.Int64Null()
		}
		if value := r.Get("as-set"); value.Exists() {
			data.AggregateAddresses[i].AsSet = types.BoolValue(true)
		} else {
			data.AggregateAddresses[i].AsSet = types.BoolValue(false)
		}
		if value := r.Get("as-confed-set"); value.Exists() {
			data.AggregateAddresses[i].AsConfedSet = types.BoolValue(true)
		} else {
			data.AggregateAddresses[i].AsConfedSet = types.BoolValue(false)
		}
		if value := r.Get("summary-only"); value.Exists() {
			data.AggregateAddresses[i].SummaryOnly = types.BoolValue(true)
		} else {
			data.AggregateAddresses[i].SummaryOnly = types.BoolValue(false)
		}
	}
	for i := range data.Networks {
		keys := [...]string{"address", "masklength"}
		keyValues := [...]string{data.Networks[i].Address.ValueString(), strconv.FormatInt(data.Networks[i].Masklength.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "networks.network").ForEach(
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
		if value := r.Get("address"); value.Exists() {
			data.Networks[i].Address = types.StringValue(value.String())
		} else {
			data.Networks[i].Address = types.StringNull()
		}
		if value := r.Get("masklength"); value.Exists() {
			data.Networks[i].Masklength = types.Int64Value(value.Int())
		} else {
			data.Networks[i].Masklength = types.Int64Null()
		}
	}
	for i := range data.RedistributeOspf {
		keys := [...]string{"router-tag"}
		keyValues := [...]string{data.RedistributeOspf[i].RouterTag.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "redistribute.ospf").ForEach(
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
		if value := r.Get("router-tag"); value.Exists() {
			data.RedistributeOspf[i].RouterTag = types.StringValue(value.String())
		} else {
			data.RedistributeOspf[i].RouterTag = types.StringNull()
		}
		if value := r.Get("match.internal"); value.Exists() {
			data.RedistributeOspf[i].MatchInternal = types.BoolValue(true)
		} else {
			data.RedistributeOspf[i].MatchInternal = types.BoolValue(false)
		}
		if value := r.Get("match.internal.external"); value.Exists() {
			data.RedistributeOspf[i].MatchInternalExternal = types.BoolValue(true)
		} else {
			data.RedistributeOspf[i].MatchInternalExternal = types.BoolValue(false)
		}
		if value := r.Get("match.internal.nssa-external"); value.Exists() {
			data.RedistributeOspf[i].MatchInternalNssaExternal = types.BoolValue(true)
		} else {
			data.RedistributeOspf[i].MatchInternalNssaExternal = types.BoolValue(false)
		}
		if value := r.Get("match.external"); value.Exists() {
			data.RedistributeOspf[i].MatchExternal = types.BoolValue(true)
		} else {
			data.RedistributeOspf[i].MatchExternal = types.BoolValue(false)
		}
		if value := r.Get("match.external.nssa-external"); value.Exists() {
			data.RedistributeOspf[i].MatchExternalNssaExternal = types.BoolValue(true)
		} else {
			data.RedistributeOspf[i].MatchExternalNssaExternal = types.BoolValue(false)
		}
		if value := r.Get("match.nssa-external"); value.Exists() {
			data.RedistributeOspf[i].MatchNssaExternal = types.BoolValue(true)
		} else {
			data.RedistributeOspf[i].MatchNssaExternal = types.BoolValue(false)
		}
		if value := r.Get("metric"); value.Exists() {
			data.RedistributeOspf[i].Metric = types.Int64Value(value.Int())
		} else {
			data.RedistributeOspf[i].Metric = types.Int64Null()
		}
	}
}

func (data *RouterBGPVRFAddressFamily) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "maximum-paths.ebgp.multipath"); value.Exists() {
		data.MaximumPathsEbgpMultipath = types.Int64Value(value.Int())
		data.MaximumPathsEbgpMultipath = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "maximum-paths.eibgp.multipath"); value.Exists() {
		data.MaximumPathsEibgpMultipath = types.Int64Value(value.Int())
		data.MaximumPathsEibgpMultipath = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "maximum-paths.ibgp.multipath"); value.Exists() {
		data.MaximumPathsIbgpMultipath = types.Int64Value(value.Int())
		data.MaximumPathsIbgpMultipath = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "label.mode.per-ce"); value.Exists() {
		data.LabelModePerCe = types.BoolValue(true)
	} else {
		data.LabelModePerCe = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "label.mode.per-vrf"); value.Exists() {
		data.LabelModePerVrf = types.BoolValue(true)
	} else {
		data.LabelModePerVrf = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.connected"); value.Exists() {
		data.RedistributeConnected = types.BoolValue(true)
	} else {
		data.RedistributeConnected = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.connected.metric"); value.Exists() {
		data.RedistributeConnectedMetric = types.Int64Value(value.Int())
		data.RedistributeConnectedMetric = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "redistribute.static"); value.Exists() {
		data.RedistributeStatic = types.BoolValue(true)
	} else {
		data.RedistributeStatic = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "redistribute.static.metric"); value.Exists() {
		data.RedistributeStaticMetric = types.Int64Value(value.Int())
		data.RedistributeStaticMetric = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "aggregate-addresses.aggregate-address"); value.Exists() {
		data.AggregateAddresses = make([]RouterBGPVRFAddressFamilyAggregateAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterBGPVRFAddressFamilyAggregateAddresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("masklength"); cValue.Exists() {
				item.Masklength = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("as-set"); cValue.Exists() {
				item.AsSet = types.BoolValue(true)
			}
			if cValue := v.Get("as-confed-set"); cValue.Exists() {
				item.AsConfedSet = types.BoolValue(true)
			}
			if cValue := v.Get("summary-only"); cValue.Exists() {
				item.SummaryOnly = types.BoolValue(true)
			}
			data.AggregateAddresses = append(data.AggregateAddresses, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "networks.network"); value.Exists() {
		data.Networks = make([]RouterBGPVRFAddressFamilyNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterBGPVRFAddressFamilyNetworks{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("masklength"); cValue.Exists() {
				item.Masklength = types.Int64Value(cValue.Int())
			}
			data.Networks = append(data.Networks, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "redistribute.ospf"); value.Exists() {
		data.RedistributeOspf = make([]RouterBGPVRFAddressFamilyRedistributeOspf, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterBGPVRFAddressFamilyRedistributeOspf{}
			if cValue := v.Get("router-tag"); cValue.Exists() {
				item.RouterTag = types.StringValue(cValue.String())
			}
			if cValue := v.Get("match.internal"); cValue.Exists() {
				item.MatchInternal = types.BoolValue(true)
			}
			if cValue := v.Get("match.internal.external"); cValue.Exists() {
				item.MatchInternalExternal = types.BoolValue(true)
			}
			if cValue := v.Get("match.internal.nssa-external"); cValue.Exists() {
				item.MatchInternalNssaExternal = types.BoolValue(true)
			}
			if cValue := v.Get("match.external"); cValue.Exists() {
				item.MatchExternal = types.BoolValue(true)
			}
			if cValue := v.Get("match.external.nssa-external"); cValue.Exists() {
				item.MatchExternalNssaExternal = types.BoolValue(true)
			}
			if cValue := v.Get("match.nssa-external"); cValue.Exists() {
				item.MatchNssaExternal = types.BoolValue(true)
			}
			if cValue := v.Get("metric"); cValue.Exists() {
				item.Metric = types.Int64Value(cValue.Int())
			}
			data.RedistributeOspf = append(data.RedistributeOspf, item)
			return true
		})
	}
}

func (data *RouterBGPVRFAddressFamily) fromPlan(plan RouterBGPVRFAddressFamily) {
	data.Device = plan.Device
	data.AsNumber = types.StringValue(plan.AsNumber.ValueString())
	data.VrfName = types.StringValue(plan.VrfName.ValueString())
	data.AfName = types.StringValue(plan.AfName.ValueString())
}

func (data *RouterBGPVRFAddressFamily) setUnknownValues() {
	if data.Device.IsUnknown() {
		data.Device = types.StringNull()
	}
	if data.Id.IsUnknown() {
		data.Id = types.StringNull()
	}
	if data.AsNumber.IsUnknown() {
		data.AsNumber = types.StringNull()
	}
	if data.VrfName.IsUnknown() {
		data.VrfName = types.StringNull()
	}
	if data.AfName.IsUnknown() {
		data.AfName = types.StringNull()
	}
	if data.MaximumPathsEbgpMultipath.IsUnknown() {
		data.MaximumPathsEbgpMultipath = types.Int64Null()
	}
	if data.MaximumPathsEibgpMultipath.IsUnknown() {
		data.MaximumPathsEibgpMultipath = types.Int64Null()
	}
	if data.MaximumPathsIbgpMultipath.IsUnknown() {
		data.MaximumPathsIbgpMultipath = types.Int64Null()
	}
	if data.LabelModePerCe.IsUnknown() {
		data.LabelModePerCe = types.BoolNull()
	}
	if data.LabelModePerVrf.IsUnknown() {
		data.LabelModePerVrf = types.BoolNull()
	}
	if data.RedistributeConnected.IsUnknown() {
		data.RedistributeConnected = types.BoolNull()
	}
	if data.RedistributeConnectedMetric.IsUnknown() {
		data.RedistributeConnectedMetric = types.Int64Null()
	}
	if data.RedistributeStatic.IsUnknown() {
		data.RedistributeStatic = types.BoolNull()
	}
	if data.RedistributeStaticMetric.IsUnknown() {
		data.RedistributeStaticMetric = types.Int64Null()
	}
	for i := range data.AggregateAddresses {
		if data.AggregateAddresses[i].Address.IsUnknown() {
			data.AggregateAddresses[i].Address = types.StringNull()
		}
		if data.AggregateAddresses[i].Masklength.IsUnknown() {
			data.AggregateAddresses[i].Masklength = types.Int64Null()
		}
		if data.AggregateAddresses[i].AsSet.IsUnknown() {
			data.AggregateAddresses[i].AsSet = types.BoolNull()
		}
		if data.AggregateAddresses[i].AsConfedSet.IsUnknown() {
			data.AggregateAddresses[i].AsConfedSet = types.BoolNull()
		}
		if data.AggregateAddresses[i].SummaryOnly.IsUnknown() {
			data.AggregateAddresses[i].SummaryOnly = types.BoolNull()
		}
	}
	for i := range data.Networks {
		if data.Networks[i].Address.IsUnknown() {
			data.Networks[i].Address = types.StringNull()
		}
		if data.Networks[i].Masklength.IsUnknown() {
			data.Networks[i].Masklength = types.Int64Null()
		}
	}
	for i := range data.RedistributeOspf {
		if data.RedistributeOspf[i].RouterTag.IsUnknown() {
			data.RedistributeOspf[i].RouterTag = types.StringNull()
		}
		if data.RedistributeOspf[i].MatchInternal.IsUnknown() {
			data.RedistributeOspf[i].MatchInternal = types.BoolNull()
		}
		if data.RedistributeOspf[i].MatchInternalExternal.IsUnknown() {
			data.RedistributeOspf[i].MatchInternalExternal = types.BoolNull()
		}
		if data.RedistributeOspf[i].MatchInternalNssaExternal.IsUnknown() {
			data.RedistributeOspf[i].MatchInternalNssaExternal = types.BoolNull()
		}
		if data.RedistributeOspf[i].MatchExternal.IsUnknown() {
			data.RedistributeOspf[i].MatchExternal = types.BoolNull()
		}
		if data.RedistributeOspf[i].MatchExternalNssaExternal.IsUnknown() {
			data.RedistributeOspf[i].MatchExternalNssaExternal = types.BoolNull()
		}
		if data.RedistributeOspf[i].MatchNssaExternal.IsUnknown() {
			data.RedistributeOspf[i].MatchNssaExternal = types.BoolNull()
		}
		if data.RedistributeOspf[i].Metric.IsUnknown() {
			data.RedistributeOspf[i].Metric = types.Int64Null()
		}
	}
}

func (data *RouterBGPVRFAddressFamily) getDeletedListItems(state RouterBGPVRFAddressFamily) []string {
	deletedListItems := make([]string, 0)
	for i := range state.AggregateAddresses {
		keys := [...]string{"address", "masklength"}
		stateKeyValues := [...]string{state.AggregateAddresses[i].Address.ValueString(), strconv.FormatInt(state.AggregateAddresses[i].Masklength.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.AggregateAddresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.AggregateAddresses[i].Masklength.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.AggregateAddresses {
			found = true
			if state.AggregateAddresses[i].Address.ValueString() != data.AggregateAddresses[j].Address.ValueString() {
				found = false
			}
			if state.AggregateAddresses[i].Masklength.ValueInt64() != data.AggregateAddresses[j].Masklength.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/aggregate-addresses/aggregate-address%v", state.getPath(), keyString))
		}
	}
	for i := range state.Networks {
		keys := [...]string{"address", "masklength"}
		stateKeyValues := [...]string{state.Networks[i].Address.ValueString(), strconv.FormatInt(state.Networks[i].Masklength.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Networks[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Networks[i].Masklength.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Networks {
			found = true
			if state.Networks[i].Address.ValueString() != data.Networks[j].Address.ValueString() {
				found = false
			}
			if state.Networks[i].Masklength.ValueInt64() != data.Networks[j].Masklength.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/networks/network%v", state.getPath(), keyString))
		}
	}
	for i := range state.RedistributeOspf {
		keys := [...]string{"router-tag"}
		stateKeyValues := [...]string{state.RedistributeOspf[i].RouterTag.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RedistributeOspf[i].RouterTag.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RedistributeOspf {
			found = true
			if state.RedistributeOspf[i].RouterTag.ValueString() != data.RedistributeOspf[j].RouterTag.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/redistribute/ospf%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterBGPVRFAddressFamily) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

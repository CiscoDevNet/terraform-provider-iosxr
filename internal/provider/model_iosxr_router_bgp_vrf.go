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

type RouterBGPVRF struct {
	Device                      types.String            `tfsdk:"device"`
	Id                          types.String            `tfsdk:"id"`
	AsNumber                    types.String            `tfsdk:"as_number"`
	VrfName                     types.String            `tfsdk:"vrf_name"`
	RdAuto                      types.Bool              `tfsdk:"rd_auto"`
	RdTwoByteAsAsNumber         types.String            `tfsdk:"rd_two_byte_as_as_number"`
	RdTwoByteAsIndex            types.Int64             `tfsdk:"rd_two_byte_as_index"`
	RdFourByteAsAsNumber        types.String            `tfsdk:"rd_four_byte_as_as_number"`
	RdFourByteAsIndex           types.Int64             `tfsdk:"rd_four_byte_as_index"`
	RdIpAddressIpv4Address      types.String            `tfsdk:"rd_ip_address_ipv4_address"`
	RdIpAddressIndex            types.Int64             `tfsdk:"rd_ip_address_index"`
	DefaultInformationOriginate types.Bool              `tfsdk:"default_information_originate"`
	DefaultMetric               types.Int64             `tfsdk:"default_metric"`
	TimersBgpKeepaliveInterval  types.Int64             `tfsdk:"timers_bgp_keepalive_interval"`
	TimersBgpHoldtime           types.String            `tfsdk:"timers_bgp_holdtime"`
	BfdMinimumInterval          types.Int64             `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier               types.Int64             `tfsdk:"bfd_multiplier"`
	Neighbors                   []RouterBGPVRFNeighbors `tfsdk:"neighbors"`
}
type RouterBGPVRFNeighbors struct {
	NeighborAddress             types.String `tfsdk:"neighbor_address"`
	RemoteAs                    types.String `tfsdk:"remote_as"`
	Description                 types.String `tfsdk:"description"`
	IgnoreConnectedCheck        types.Bool   `tfsdk:"ignore_connected_check"`
	EbgpMultihopMaximumHopCount types.Int64  `tfsdk:"ebgp_multihop_maximum_hop_count"`
	BfdMinimumInterval          types.Int64  `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier               types.Int64  `tfsdk:"bfd_multiplier"`
	LocalAs                     types.String `tfsdk:"local_as"`
	LocalAsNoPrepend            types.Bool   `tfsdk:"local_as_no_prepend"`
	LocalAsReplaceAs            types.Bool   `tfsdk:"local_as_replace_as"`
	LocalAsDualAs               types.Bool   `tfsdk:"local_as_dual_as"`
	Password                    types.String `tfsdk:"password"`
	Shutdown                    types.Bool   `tfsdk:"shutdown"`
	TimersKeepaliveInterval     types.Int64  `tfsdk:"timers_keepalive_interval"`
	TimersHoldtime              types.String `tfsdk:"timers_holdtime"`
	UpdateSource                types.String `tfsdk:"update_source"`
	TtlSecurity                 types.Bool   `tfsdk:"ttl_security"`
}

func (data RouterBGPVRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/vrfs/vrf[vrf-name=%s]", data.AsNumber.ValueString(), data.VrfName.ValueString())
}

func (data RouterBGPVRF) toBody(ctx context.Context) string {
	body := "{}"
	if !data.VrfName.IsNull() && !data.VrfName.IsUnknown() {
		body, _ = sjson.Set(body, "vrf-name", data.VrfName.ValueString())
	}
	if !data.RdAuto.IsNull() && !data.RdAuto.IsUnknown() {
		if data.RdAuto.ValueBool() {
			body, _ = sjson.Set(body, "rd.auto", map[string]string{})
		}
	}
	if !data.RdTwoByteAsAsNumber.IsNull() && !data.RdTwoByteAsAsNumber.IsUnknown() {
		body, _ = sjson.Set(body, "rd.two-byte-as.as-number", data.RdTwoByteAsAsNumber.ValueString())
	}
	if !data.RdTwoByteAsIndex.IsNull() && !data.RdTwoByteAsIndex.IsUnknown() {
		body, _ = sjson.Set(body, "rd.two-byte-as.index", strconv.FormatInt(data.RdTwoByteAsIndex.ValueInt64(), 10))
	}
	if !data.RdFourByteAsAsNumber.IsNull() && !data.RdFourByteAsAsNumber.IsUnknown() {
		body, _ = sjson.Set(body, "rd.four-byte-as.as-number", data.RdFourByteAsAsNumber.ValueString())
	}
	if !data.RdFourByteAsIndex.IsNull() && !data.RdFourByteAsIndex.IsUnknown() {
		body, _ = sjson.Set(body, "rd.four-byte-as.index", strconv.FormatInt(data.RdFourByteAsIndex.ValueInt64(), 10))
	}
	if !data.RdIpAddressIpv4Address.IsNull() && !data.RdIpAddressIpv4Address.IsUnknown() {
		body, _ = sjson.Set(body, "rd.ip-address.ipv4-address", data.RdIpAddressIpv4Address.ValueString())
	}
	if !data.RdIpAddressIndex.IsNull() && !data.RdIpAddressIndex.IsUnknown() {
		body, _ = sjson.Set(body, "rd.ip-address.index", strconv.FormatInt(data.RdIpAddressIndex.ValueInt64(), 10))
	}
	if !data.DefaultInformationOriginate.IsNull() && !data.DefaultInformationOriginate.IsUnknown() {
		if data.DefaultInformationOriginate.ValueBool() {
			body, _ = sjson.Set(body, "default-information.originate", map[string]string{})
		}
	}
	if !data.DefaultMetric.IsNull() && !data.DefaultMetric.IsUnknown() {
		body, _ = sjson.Set(body, "default-metric", strconv.FormatInt(data.DefaultMetric.ValueInt64(), 10))
	}
	if !data.TimersBgpKeepaliveInterval.IsNull() && !data.TimersBgpKeepaliveInterval.IsUnknown() {
		body, _ = sjson.Set(body, "timers.bgp.keepalive-interval", strconv.FormatInt(data.TimersBgpKeepaliveInterval.ValueInt64(), 10))
	}
	if !data.TimersBgpHoldtime.IsNull() && !data.TimersBgpHoldtime.IsUnknown() {
		body, _ = sjson.Set(body, "timers.bgp.holdtime", data.TimersBgpHoldtime.ValueString())
	}
	if !data.BfdMinimumInterval.IsNull() && !data.BfdMinimumInterval.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.ValueInt64(), 10))
	}
	if !data.BfdMultiplier.IsNull() && !data.BfdMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.ValueInt64(), 10))
	}
	if len(data.Neighbors) > 0 {
		body, _ = sjson.Set(body, "neighbors.neighbor", []interface{}{})
		for index, item := range data.Neighbors {
			if !item.NeighborAddress.IsNull() && !item.NeighborAddress.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"neighbor-address", item.NeighborAddress.ValueString())
			}
			if !item.RemoteAs.IsNull() && !item.RemoteAs.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"remote-as", item.RemoteAs.ValueString())
			}
			if !item.Description.IsNull() && !item.Description.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"description", item.Description.ValueString())
			}
			if !item.IgnoreConnectedCheck.IsNull() && !item.IgnoreConnectedCheck.IsUnknown() {
				if item.IgnoreConnectedCheck.ValueBool() {
					body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"ignore-connected-check", map[string]string{})
				}
			}
			if !item.EbgpMultihopMaximumHopCount.IsNull() && !item.EbgpMultihopMaximumHopCount.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"ebgp-multihop.maximum-hop-count", strconv.FormatInt(item.EbgpMultihopMaximumHopCount.ValueInt64(), 10))
			}
			if !item.BfdMinimumInterval.IsNull() && !item.BfdMinimumInterval.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"bfd.minimum-interval", strconv.FormatInt(item.BfdMinimumInterval.ValueInt64(), 10))
			}
			if !item.BfdMultiplier.IsNull() && !item.BfdMultiplier.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"bfd.multiplier", strconv.FormatInt(item.BfdMultiplier.ValueInt64(), 10))
			}
			if !item.LocalAs.IsNull() && !item.LocalAs.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"local-as.as-number", item.LocalAs.ValueString())
			}
			if !item.LocalAsNoPrepend.IsNull() && !item.LocalAsNoPrepend.IsUnknown() {
				if item.LocalAsNoPrepend.ValueBool() {
					body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"local-as.no-prepend", map[string]string{})
				}
			}
			if !item.LocalAsReplaceAs.IsNull() && !item.LocalAsReplaceAs.IsUnknown() {
				if item.LocalAsReplaceAs.ValueBool() {
					body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"local-as.no-prepend.replace-as", map[string]string{})
				}
			}
			if !item.LocalAsDualAs.IsNull() && !item.LocalAsDualAs.IsUnknown() {
				if item.LocalAsDualAs.ValueBool() {
					body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"local-as.no-prepend.replace-as.dual-as", map[string]string{})
				}
			}
			if !item.Password.IsNull() && !item.Password.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"password.encrypted", item.Password.ValueString())
			}
			if !item.Shutdown.IsNull() && !item.Shutdown.IsUnknown() {
				if item.Shutdown.ValueBool() {
					body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"shutdown", map[string]string{})
				}
			}
			if !item.TimersKeepaliveInterval.IsNull() && !item.TimersKeepaliveInterval.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"timers.keepalive-interval", strconv.FormatInt(item.TimersKeepaliveInterval.ValueInt64(), 10))
			}
			if !item.TimersHoldtime.IsNull() && !item.TimersHoldtime.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"timers.holdtime", item.TimersHoldtime.ValueString())
			}
			if !item.UpdateSource.IsNull() && !item.UpdateSource.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"update-source", item.UpdateSource.ValueString())
			}
			if !item.TtlSecurity.IsNull() && !item.TtlSecurity.IsUnknown() {
				if item.TtlSecurity.ValueBool() {
					body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"ttl-security", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *RouterBGPVRF) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rd.auto"); !data.RdAuto.IsNull() {
		if value.Exists() {
			data.RdAuto = types.BoolValue(true)
		} else {
			data.RdAuto = types.BoolValue(false)
		}
	} else {
		data.RdAuto = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "rd.two-byte-as.as-number"); value.Exists() && !data.RdTwoByteAsAsNumber.IsNull() {
		data.RdTwoByteAsAsNumber = types.StringValue(value.String())
	} else {
		data.RdTwoByteAsAsNumber = types.StringNull()
	}
	if value := gjson.GetBytes(res, "rd.two-byte-as.index"); value.Exists() && !data.RdTwoByteAsIndex.IsNull() {
		data.RdTwoByteAsIndex = types.Int64Value(value.Int())
	} else {
		data.RdTwoByteAsIndex = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "rd.four-byte-as.as-number"); value.Exists() && !data.RdFourByteAsAsNumber.IsNull() {
		data.RdFourByteAsAsNumber = types.StringValue(value.String())
	} else {
		data.RdFourByteAsAsNumber = types.StringNull()
	}
	if value := gjson.GetBytes(res, "rd.four-byte-as.index"); value.Exists() && !data.RdFourByteAsIndex.IsNull() {
		data.RdFourByteAsIndex = types.Int64Value(value.Int())
	} else {
		data.RdFourByteAsIndex = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "rd.ip-address.ipv4-address"); value.Exists() && !data.RdIpAddressIpv4Address.IsNull() {
		data.RdIpAddressIpv4Address = types.StringValue(value.String())
	} else {
		data.RdIpAddressIpv4Address = types.StringNull()
	}
	if value := gjson.GetBytes(res, "rd.ip-address.index"); value.Exists() && !data.RdIpAddressIndex.IsNull() {
		data.RdIpAddressIndex = types.Int64Value(value.Int())
	} else {
		data.RdIpAddressIndex = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "default-information.originate"); !data.DefaultInformationOriginate.IsNull() {
		if value.Exists() {
			data.DefaultInformationOriginate = types.BoolValue(true)
		} else {
			data.DefaultInformationOriginate = types.BoolValue(false)
		}
	} else {
		data.DefaultInformationOriginate = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "default-metric"); value.Exists() && !data.DefaultMetric.IsNull() {
		data.DefaultMetric = types.Int64Value(value.Int())
	} else {
		data.DefaultMetric = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.bgp.keepalive-interval"); value.Exists() && !data.TimersBgpKeepaliveInterval.IsNull() {
		data.TimersBgpKeepaliveInterval = types.Int64Value(value.Int())
	} else {
		data.TimersBgpKeepaliveInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.bgp.holdtime"); value.Exists() && !data.TimersBgpHoldtime.IsNull() {
		data.TimersBgpHoldtime = types.StringValue(value.String())
	} else {
		data.TimersBgpHoldtime = types.StringNull()
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() && !data.BfdMinimumInterval.IsNull() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	} else {
		data.BfdMinimumInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() && !data.BfdMultiplier.IsNull() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	} else {
		data.BfdMultiplier = types.Int64Null()
	}
	for i := range data.Neighbors {
		keys := [...]string{"neighbor-address"}
		keyValues := [...]string{data.Neighbors[i].NeighborAddress.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "neighbors.neighbor").ForEach(
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
		if value := r.Get("neighbor-address"); value.Exists() && !data.Neighbors[i].NeighborAddress.IsNull() {
			data.Neighbors[i].NeighborAddress = types.StringValue(value.String())
		} else {
			data.Neighbors[i].NeighborAddress = types.StringNull()
		}
		if value := r.Get("remote-as"); value.Exists() && !data.Neighbors[i].RemoteAs.IsNull() {
			data.Neighbors[i].RemoteAs = types.StringValue(value.String())
		} else {
			data.Neighbors[i].RemoteAs = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.Neighbors[i].Description.IsNull() {
			data.Neighbors[i].Description = types.StringValue(value.String())
		} else {
			data.Neighbors[i].Description = types.StringNull()
		}
		if value := r.Get("ignore-connected-check"); !data.Neighbors[i].IgnoreConnectedCheck.IsNull() {
			if value.Exists() {
				data.Neighbors[i].IgnoreConnectedCheck = types.BoolValue(true)
			} else {
				data.Neighbors[i].IgnoreConnectedCheck = types.BoolValue(false)
			}
		} else {
			data.Neighbors[i].IgnoreConnectedCheck = types.BoolNull()
		}
		if value := r.Get("ebgp-multihop.maximum-hop-count"); value.Exists() && !data.Neighbors[i].EbgpMultihopMaximumHopCount.IsNull() {
			data.Neighbors[i].EbgpMultihopMaximumHopCount = types.Int64Value(value.Int())
		} else {
			data.Neighbors[i].EbgpMultihopMaximumHopCount = types.Int64Null()
		}
		if value := r.Get("bfd.minimum-interval"); value.Exists() && !data.Neighbors[i].BfdMinimumInterval.IsNull() {
			data.Neighbors[i].BfdMinimumInterval = types.Int64Value(value.Int())
		} else {
			data.Neighbors[i].BfdMinimumInterval = types.Int64Null()
		}
		if value := r.Get("bfd.multiplier"); value.Exists() && !data.Neighbors[i].BfdMultiplier.IsNull() {
			data.Neighbors[i].BfdMultiplier = types.Int64Value(value.Int())
		} else {
			data.Neighbors[i].BfdMultiplier = types.Int64Null()
		}
		if value := r.Get("local-as.as-number"); value.Exists() && !data.Neighbors[i].LocalAs.IsNull() {
			data.Neighbors[i].LocalAs = types.StringValue(value.String())
		} else {
			data.Neighbors[i].LocalAs = types.StringNull()
		}
		if value := r.Get("local-as.no-prepend"); !data.Neighbors[i].LocalAsNoPrepend.IsNull() {
			if value.Exists() {
				data.Neighbors[i].LocalAsNoPrepend = types.BoolValue(true)
			} else {
				data.Neighbors[i].LocalAsNoPrepend = types.BoolValue(false)
			}
		} else {
			data.Neighbors[i].LocalAsNoPrepend = types.BoolNull()
		}
		if value := r.Get("local-as.no-prepend.replace-as"); !data.Neighbors[i].LocalAsReplaceAs.IsNull() {
			if value.Exists() {
				data.Neighbors[i].LocalAsReplaceAs = types.BoolValue(true)
			} else {
				data.Neighbors[i].LocalAsReplaceAs = types.BoolValue(false)
			}
		} else {
			data.Neighbors[i].LocalAsReplaceAs = types.BoolNull()
		}
		if value := r.Get("local-as.no-prepend.replace-as.dual-as"); !data.Neighbors[i].LocalAsDualAs.IsNull() {
			if value.Exists() {
				data.Neighbors[i].LocalAsDualAs = types.BoolValue(true)
			} else {
				data.Neighbors[i].LocalAsDualAs = types.BoolValue(false)
			}
		} else {
			data.Neighbors[i].LocalAsDualAs = types.BoolNull()
		}
		if value := r.Get("password.encrypted"); value.Exists() && !data.Neighbors[i].Password.IsNull() {
			data.Neighbors[i].Password = types.StringValue(value.String())
		} else {
			data.Neighbors[i].Password = types.StringNull()
		}
		if value := r.Get("shutdown"); !data.Neighbors[i].Shutdown.IsNull() {
			if value.Exists() {
				data.Neighbors[i].Shutdown = types.BoolValue(true)
			} else {
				data.Neighbors[i].Shutdown = types.BoolValue(false)
			}
		} else {
			data.Neighbors[i].Shutdown = types.BoolNull()
		}
		if value := r.Get("timers.keepalive-interval"); value.Exists() && !data.Neighbors[i].TimersKeepaliveInterval.IsNull() {
			data.Neighbors[i].TimersKeepaliveInterval = types.Int64Value(value.Int())
		} else {
			data.Neighbors[i].TimersKeepaliveInterval = types.Int64Null()
		}
		if value := r.Get("timers.holdtime"); value.Exists() && !data.Neighbors[i].TimersHoldtime.IsNull() {
			data.Neighbors[i].TimersHoldtime = types.StringValue(value.String())
		} else {
			data.Neighbors[i].TimersHoldtime = types.StringNull()
		}
		if value := r.Get("update-source"); value.Exists() && !data.Neighbors[i].UpdateSource.IsNull() {
			data.Neighbors[i].UpdateSource = types.StringValue(value.String())
		} else {
			data.Neighbors[i].UpdateSource = types.StringNull()
		}
		if value := r.Get("ttl-security"); !data.Neighbors[i].TtlSecurity.IsNull() {
			if value.Exists() {
				data.Neighbors[i].TtlSecurity = types.BoolValue(true)
			} else {
				data.Neighbors[i].TtlSecurity = types.BoolValue(false)
			}
		} else {
			data.Neighbors[i].TtlSecurity = types.BoolNull()
		}
	}
}

func (data *RouterBGPVRF) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rd.auto"); value.Exists() {
		data.RdAuto = types.BoolValue(true)
	} else {
		data.RdAuto = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "rd.two-byte-as.as-number"); value.Exists() {
		data.RdTwoByteAsAsNumber = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "rd.two-byte-as.index"); value.Exists() {
		data.RdTwoByteAsIndex = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "rd.four-byte-as.as-number"); value.Exists() {
		data.RdFourByteAsAsNumber = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "rd.four-byte-as.index"); value.Exists() {
		data.RdFourByteAsIndex = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "rd.ip-address.ipv4-address"); value.Exists() {
		data.RdIpAddressIpv4Address = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "rd.ip-address.index"); value.Exists() {
		data.RdIpAddressIndex = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate = types.BoolValue(true)
	} else {
		data.DefaultInformationOriginate = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "default-metric"); value.Exists() {
		data.DefaultMetric = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.bgp.keepalive-interval"); value.Exists() {
		data.TimersBgpKeepaliveInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.bgp.holdtime"); value.Exists() {
		data.TimersBgpHoldtime = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "neighbors.neighbor"); value.Exists() {
		data.Neighbors = make([]RouterBGPVRFNeighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterBGPVRFNeighbors{}
			if cValue := v.Get("neighbor-address"); cValue.Exists() {
				item.NeighborAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("remote-as"); cValue.Exists() {
				item.RemoteAs = types.StringValue(cValue.String())
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ignore-connected-check"); cValue.Exists() {
				item.IgnoreConnectedCheck = types.BoolValue(true)
			} else {
				item.IgnoreConnectedCheck = types.BoolValue(false)
			}
			if cValue := v.Get("ebgp-multihop.maximum-hop-count"); cValue.Exists() {
				item.EbgpMultihopMaximumHopCount = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("bfd.minimum-interval"); cValue.Exists() {
				item.BfdMinimumInterval = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("bfd.multiplier"); cValue.Exists() {
				item.BfdMultiplier = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("local-as.as-number"); cValue.Exists() {
				item.LocalAs = types.StringValue(cValue.String())
			}
			if cValue := v.Get("local-as.no-prepend"); cValue.Exists() {
				item.LocalAsNoPrepend = types.BoolValue(true)
			} else {
				item.LocalAsNoPrepend = types.BoolValue(false)
			}
			if cValue := v.Get("local-as.no-prepend.replace-as"); cValue.Exists() {
				item.LocalAsReplaceAs = types.BoolValue(true)
			} else {
				item.LocalAsReplaceAs = types.BoolValue(false)
			}
			if cValue := v.Get("local-as.no-prepend.replace-as.dual-as"); cValue.Exists() {
				item.LocalAsDualAs = types.BoolValue(true)
			} else {
				item.LocalAsDualAs = types.BoolValue(false)
			}
			if cValue := v.Get("password.encrypted"); cValue.Exists() {
				item.Password = types.StringValue(cValue.String())
			}
			if cValue := v.Get("shutdown"); cValue.Exists() {
				item.Shutdown = types.BoolValue(true)
			} else {
				item.Shutdown = types.BoolValue(false)
			}
			if cValue := v.Get("timers.keepalive-interval"); cValue.Exists() {
				item.TimersKeepaliveInterval = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("timers.holdtime"); cValue.Exists() {
				item.TimersHoldtime = types.StringValue(cValue.String())
			}
			if cValue := v.Get("update-source"); cValue.Exists() {
				item.UpdateSource = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ttl-security"); cValue.Exists() {
				item.TtlSecurity = types.BoolValue(true)
			} else {
				item.TtlSecurity = types.BoolValue(false)
			}
			data.Neighbors = append(data.Neighbors, item)
			return true
		})
	}
}

func (data *RouterBGPVRF) getDeletedListItems(ctx context.Context, state RouterBGPVRF) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Neighbors {
		keys := [...]string{"neighbor-address"}
		stateKeyValues := [...]string{state.Neighbors[i].NeighborAddress.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Neighbors[i].NeighborAddress.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Neighbors {
			found = true
			if state.Neighbors[i].NeighborAddress.ValueString() != data.Neighbors[j].NeighborAddress.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbors/neighbor%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterBGPVRF) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.RdAuto.IsNull() && !data.RdAuto.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/rd/auto", data.getPath()))
	}
	if !data.DefaultInformationOriginate.IsNull() && !data.DefaultInformationOriginate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/default-information/originate", data.getPath()))
	}
	for i := range data.Neighbors {
		keys := [...]string{"neighbor-address"}
		keyValues := [...]string{data.Neighbors[i].NeighborAddress.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.Neighbors[i].IgnoreConnectedCheck.IsNull() && !data.Neighbors[i].IgnoreConnectedCheck.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/neighbors/neighbor%v/ignore-connected-check", data.getPath(), keyString))
		}
		if !data.Neighbors[i].LocalAsNoPrepend.IsNull() && !data.Neighbors[i].LocalAsNoPrepend.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/neighbors/neighbor%v/local-as/no-prepend", data.getPath(), keyString))
		}
		if !data.Neighbors[i].LocalAsReplaceAs.IsNull() && !data.Neighbors[i].LocalAsReplaceAs.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/neighbors/neighbor%v/local-as/no-prepend/replace-as", data.getPath(), keyString))
		}
		if !data.Neighbors[i].LocalAsDualAs.IsNull() && !data.Neighbors[i].LocalAsDualAs.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/neighbors/neighbor%v/local-as/no-prepend/replace-as/dual-as", data.getPath(), keyString))
		}
		if !data.Neighbors[i].Shutdown.IsNull() && !data.Neighbors[i].Shutdown.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/neighbors/neighbor%v/shutdown", data.getPath(), keyString))
		}
		if !data.Neighbors[i].TtlSecurity.IsNull() && !data.Neighbors[i].TtlSecurity.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/neighbors/neighbor%v/ttl-security", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

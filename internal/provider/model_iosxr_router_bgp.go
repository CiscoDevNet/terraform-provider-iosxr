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

type RouterBGP struct {
	Device                                types.String              `tfsdk:"device"`
	Id                                    types.String              `tfsdk:"id"`
	AsNumber                              types.String              `tfsdk:"as_number"`
	DefaultInformationOriginate           types.Bool                `tfsdk:"default_information_originate"`
	DefaultMetric                         types.Int64               `tfsdk:"default_metric"`
	NsrDisable                            types.Bool                `tfsdk:"nsr_disable"`
	BgpRedistributeInternal               types.Bool                `tfsdk:"bgp_redistribute_internal"`
	SegmentRoutingSrv6Locator             types.String              `tfsdk:"segment_routing_srv6_locator"`
	TimersBgpKeepaliveInterval            types.Int64               `tfsdk:"timers_bgp_keepalive_interval"`
	TimersBgpHoldtime                     types.String              `tfsdk:"timers_bgp_holdtime"`
	BgpRouterId                           types.String              `tfsdk:"bgp_router_id"`
	BgpGracefulRestartGracefulReset       types.Bool                `tfsdk:"bgp_graceful_restart_graceful_reset"`
	IbgpPolicyOutEnforceModifications     types.Bool                `tfsdk:"ibgp_policy_out_enforce_modifications"`
	BgpLogNeighborChangesDetail           types.Bool                `tfsdk:"bgp_log_neighbor_changes_detail"`
	BfdMinimumInterval                    types.Int64               `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier                         types.Int64               `tfsdk:"bfd_multiplier"`
	NexthopValidationColorExtcommSrPolicy types.Bool                `tfsdk:"nexthop_validation_color_extcomm_sr_policy"`
	NexthopValidationColorExtcommDisable  types.Bool                `tfsdk:"nexthop_validation_color_extcomm_disable"`
	Neighbors                             []RouterBGPNeighbors      `tfsdk:"neighbors"`
	NeighborGroups                        []RouterBGPNeighborGroups `tfsdk:"neighbor_groups"`
}
type RouterBGPNeighbors struct {
	NeighborAddress             types.String `tfsdk:"neighbor_address"`
	RemoteAs                    types.String `tfsdk:"remote_as"`
	Description                 types.String `tfsdk:"description"`
	UseNeighborGroup            types.String `tfsdk:"use_neighbor_group"`
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
type RouterBGPNeighborGroups struct {
	Name                      types.String `tfsdk:"name"`
	RemoteAs                  types.String `tfsdk:"remote_as"`
	UpdateSource              types.String `tfsdk:"update_source"`
	AoKeyChainName            types.String `tfsdk:"ao_key_chain_name"`
	AoIncludeTcpOptionsEnable types.Bool   `tfsdk:"ao_include_tcp_options_enable"`
	BfdMinimumInterval        types.Int64  `tfsdk:"bfd_minimum_interval"`
}

func (data RouterBGP) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]", data.AsNumber.ValueString())
}

func (data RouterBGP) toBody(ctx context.Context) string {
	body := "{}"
	if !data.AsNumber.IsNull() && !data.AsNumber.IsUnknown() {
		body, _ = sjson.Set(body, "as-number", data.AsNumber.ValueString())
	}
	if !data.DefaultInformationOriginate.IsNull() && !data.DefaultInformationOriginate.IsUnknown() {
		if data.DefaultInformationOriginate.ValueBool() {
			body, _ = sjson.Set(body, "default-information.originate", map[string]string{})
		}
	}
	if !data.DefaultMetric.IsNull() && !data.DefaultMetric.IsUnknown() {
		body, _ = sjson.Set(body, "default-metric", strconv.FormatInt(data.DefaultMetric.ValueInt64(), 10))
	}
	if !data.NsrDisable.IsNull() && !data.NsrDisable.IsUnknown() {
		if data.NsrDisable.ValueBool() {
			body, _ = sjson.Set(body, "nsr.disable", map[string]string{})
		}
	}
	if !data.BgpRedistributeInternal.IsNull() && !data.BgpRedistributeInternal.IsUnknown() {
		if data.BgpRedistributeInternal.ValueBool() {
			body, _ = sjson.Set(body, "bgp.redistribute-internal", map[string]string{})
		}
	}
	if !data.SegmentRoutingSrv6Locator.IsNull() && !data.SegmentRoutingSrv6Locator.IsUnknown() {
		body, _ = sjson.Set(body, "segment-routing.srv6.locator", data.SegmentRoutingSrv6Locator.ValueString())
	}
	if !data.TimersBgpKeepaliveInterval.IsNull() && !data.TimersBgpKeepaliveInterval.IsUnknown() {
		body, _ = sjson.Set(body, "timers.bgp.keepalive-interval", strconv.FormatInt(data.TimersBgpKeepaliveInterval.ValueInt64(), 10))
	}
	if !data.TimersBgpHoldtime.IsNull() && !data.TimersBgpHoldtime.IsUnknown() {
		body, _ = sjson.Set(body, "timers.bgp.holdtime", data.TimersBgpHoldtime.ValueString())
	}
	if !data.BgpRouterId.IsNull() && !data.BgpRouterId.IsUnknown() {
		body, _ = sjson.Set(body, "bgp.router-id", data.BgpRouterId.ValueString())
	}
	if !data.BgpGracefulRestartGracefulReset.IsNull() && !data.BgpGracefulRestartGracefulReset.IsUnknown() {
		if data.BgpGracefulRestartGracefulReset.ValueBool() {
			body, _ = sjson.Set(body, "bgp.graceful-restart.graceful-reset", map[string]string{})
		}
	}
	if !data.IbgpPolicyOutEnforceModifications.IsNull() && !data.IbgpPolicyOutEnforceModifications.IsUnknown() {
		if data.IbgpPolicyOutEnforceModifications.ValueBool() {
			body, _ = sjson.Set(body, "ibgp.policy.out.enforce-modifications", map[string]string{})
		}
	}
	if !data.BgpLogNeighborChangesDetail.IsNull() && !data.BgpLogNeighborChangesDetail.IsUnknown() {
		if data.BgpLogNeighborChangesDetail.ValueBool() {
			body, _ = sjson.Set(body, "bgp.log.neighbor.changes.detail", map[string]string{})
		}
	}
	if !data.BfdMinimumInterval.IsNull() && !data.BfdMinimumInterval.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.ValueInt64(), 10))
	}
	if !data.BfdMultiplier.IsNull() && !data.BfdMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.ValueInt64(), 10))
	}
	if !data.NexthopValidationColorExtcommSrPolicy.IsNull() && !data.NexthopValidationColorExtcommSrPolicy.IsUnknown() {
		if data.NexthopValidationColorExtcommSrPolicy.ValueBool() {
			body, _ = sjson.Set(body, "nexthop.validation.color-extcomm.sr-policy", map[string]string{})
		}
	}
	if !data.NexthopValidationColorExtcommDisable.IsNull() && !data.NexthopValidationColorExtcommDisable.IsUnknown() {
		if data.NexthopValidationColorExtcommDisable.ValueBool() {
			body, _ = sjson.Set(body, "nexthop.validation.color-extcomm.disable", map[string]string{})
		}
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
			if !item.UseNeighborGroup.IsNull() && !item.UseNeighborGroup.IsUnknown() {
				body, _ = sjson.Set(body, "neighbors.neighbor"+"."+strconv.Itoa(index)+"."+"use.neighbor-group", item.UseNeighborGroup.ValueString())
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
	if len(data.NeighborGroups) > 0 {
		body, _ = sjson.Set(body, "neighbor-groups.neighbor-group", []interface{}{})
		for index, item := range data.NeighborGroups {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor-groups.neighbor-group"+"."+strconv.Itoa(index)+"."+"neighbor-group-name", item.Name.ValueString())
			}
			if !item.RemoteAs.IsNull() && !item.RemoteAs.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor-groups.neighbor-group"+"."+strconv.Itoa(index)+"."+"remote-as", item.RemoteAs.ValueString())
			}
			if !item.UpdateSource.IsNull() && !item.UpdateSource.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor-groups.neighbor-group"+"."+strconv.Itoa(index)+"."+"update-source", item.UpdateSource.ValueString())
			}
			if !item.AoKeyChainName.IsNull() && !item.AoKeyChainName.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor-groups.neighbor-group"+"."+strconv.Itoa(index)+"."+"ao.key-chain-name", item.AoKeyChainName.ValueString())
			}
			if !item.AoIncludeTcpOptionsEnable.IsNull() && !item.AoIncludeTcpOptionsEnable.IsUnknown() {
				if item.AoIncludeTcpOptionsEnable.ValueBool() {
					body, _ = sjson.Set(body, "neighbor-groups.neighbor-group"+"."+strconv.Itoa(index)+"."+"ao.include-tcp-options.enable", map[string]string{})
				}
			}
			if !item.BfdMinimumInterval.IsNull() && !item.BfdMinimumInterval.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor-groups.neighbor-group"+"."+strconv.Itoa(index)+"."+"bfd.minimum-interval", strconv.FormatInt(item.BfdMinimumInterval.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *RouterBGP) updateFromBody(ctx context.Context, res []byte) {
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
	if value := gjson.GetBytes(res, "nsr.disable"); !data.NsrDisable.IsNull() {
		if value.Exists() {
			data.NsrDisable = types.BoolValue(true)
		} else {
			data.NsrDisable = types.BoolValue(false)
		}
	} else {
		data.NsrDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "bgp.redistribute-internal"); !data.BgpRedistributeInternal.IsNull() {
		if value.Exists() {
			data.BgpRedistributeInternal = types.BoolValue(true)
		} else {
			data.BgpRedistributeInternal = types.BoolValue(false)
		}
	} else {
		data.BgpRedistributeInternal = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "segment-routing.srv6.locator"); value.Exists() && !data.SegmentRoutingSrv6Locator.IsNull() {
		data.SegmentRoutingSrv6Locator = types.StringValue(value.String())
	} else {
		data.SegmentRoutingSrv6Locator = types.StringNull()
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
	if value := gjson.GetBytes(res, "bgp.router-id"); value.Exists() && !data.BgpRouterId.IsNull() {
		data.BgpRouterId = types.StringValue(value.String())
	} else {
		data.BgpRouterId = types.StringNull()
	}
	if value := gjson.GetBytes(res, "bgp.graceful-restart.graceful-reset"); !data.BgpGracefulRestartGracefulReset.IsNull() {
		if value.Exists() {
			data.BgpGracefulRestartGracefulReset = types.BoolValue(true)
		} else {
			data.BgpGracefulRestartGracefulReset = types.BoolValue(false)
		}
	} else {
		data.BgpGracefulRestartGracefulReset = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "ibgp.policy.out.enforce-modifications"); !data.IbgpPolicyOutEnforceModifications.IsNull() {
		if value.Exists() {
			data.IbgpPolicyOutEnforceModifications = types.BoolValue(true)
		} else {
			data.IbgpPolicyOutEnforceModifications = types.BoolValue(false)
		}
	} else {
		data.IbgpPolicyOutEnforceModifications = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "bgp.log.neighbor.changes.detail"); !data.BgpLogNeighborChangesDetail.IsNull() {
		if value.Exists() {
			data.BgpLogNeighborChangesDetail = types.BoolValue(true)
		} else {
			data.BgpLogNeighborChangesDetail = types.BoolValue(false)
		}
	} else {
		data.BgpLogNeighborChangesDetail = types.BoolNull()
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
	if value := gjson.GetBytes(res, "nexthop.validation.color-extcomm.sr-policy"); !data.NexthopValidationColorExtcommSrPolicy.IsNull() {
		if value.Exists() {
			data.NexthopValidationColorExtcommSrPolicy = types.BoolValue(true)
		} else {
			data.NexthopValidationColorExtcommSrPolicy = types.BoolValue(false)
		}
	} else {
		data.NexthopValidationColorExtcommSrPolicy = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "nexthop.validation.color-extcomm.disable"); !data.NexthopValidationColorExtcommDisable.IsNull() {
		if value.Exists() {
			data.NexthopValidationColorExtcommDisable = types.BoolValue(true)
		} else {
			data.NexthopValidationColorExtcommDisable = types.BoolValue(false)
		}
	} else {
		data.NexthopValidationColorExtcommDisable = types.BoolNull()
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
		if value := r.Get("use.neighbor-group"); value.Exists() && !data.Neighbors[i].UseNeighborGroup.IsNull() {
			data.Neighbors[i].UseNeighborGroup = types.StringValue(value.String())
		} else {
			data.Neighbors[i].UseNeighborGroup = types.StringNull()
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
	for i := range data.NeighborGroups {
		keys := [...]string{"neighbor-group-name"}
		keyValues := [...]string{data.NeighborGroups[i].Name.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "neighbor-groups.neighbor-group").ForEach(
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
		if value := r.Get("neighbor-group-name"); value.Exists() && !data.NeighborGroups[i].Name.IsNull() {
			data.NeighborGroups[i].Name = types.StringValue(value.String())
		} else {
			data.NeighborGroups[i].Name = types.StringNull()
		}
		if value := r.Get("remote-as"); value.Exists() && !data.NeighborGroups[i].RemoteAs.IsNull() {
			data.NeighborGroups[i].RemoteAs = types.StringValue(value.String())
		} else {
			data.NeighborGroups[i].RemoteAs = types.StringNull()
		}
		if value := r.Get("update-source"); value.Exists() && !data.NeighborGroups[i].UpdateSource.IsNull() {
			data.NeighborGroups[i].UpdateSource = types.StringValue(value.String())
		} else {
			data.NeighborGroups[i].UpdateSource = types.StringNull()
		}
		if value := r.Get("ao.key-chain-name"); value.Exists() && !data.NeighborGroups[i].AoKeyChainName.IsNull() {
			data.NeighborGroups[i].AoKeyChainName = types.StringValue(value.String())
		} else {
			data.NeighborGroups[i].AoKeyChainName = types.StringNull()
		}
		if value := r.Get("ao.include-tcp-options.enable"); !data.NeighborGroups[i].AoIncludeTcpOptionsEnable.IsNull() {
			if value.Exists() {
				data.NeighborGroups[i].AoIncludeTcpOptionsEnable = types.BoolValue(true)
			} else {
				data.NeighborGroups[i].AoIncludeTcpOptionsEnable = types.BoolValue(false)
			}
		} else {
			data.NeighborGroups[i].AoIncludeTcpOptionsEnable = types.BoolNull()
		}
		if value := r.Get("bfd.minimum-interval"); value.Exists() && !data.NeighborGroups[i].BfdMinimumInterval.IsNull() {
			data.NeighborGroups[i].BfdMinimumInterval = types.Int64Value(value.Int())
		} else {
			data.NeighborGroups[i].BfdMinimumInterval = types.Int64Null()
		}
	}
}

func (data *RouterBGP) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate = types.BoolValue(true)
	} else {
		data.DefaultInformationOriginate = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "default-metric"); value.Exists() {
		data.DefaultMetric = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "nsr.disable"); value.Exists() {
		data.NsrDisable = types.BoolValue(true)
	} else {
		data.NsrDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "bgp.redistribute-internal"); value.Exists() {
		data.BgpRedistributeInternal = types.BoolValue(true)
	} else {
		data.BgpRedistributeInternal = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "segment-routing.srv6.locator"); value.Exists() {
		data.SegmentRoutingSrv6Locator = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "timers.bgp.keepalive-interval"); value.Exists() {
		data.TimersBgpKeepaliveInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.bgp.holdtime"); value.Exists() {
		data.TimersBgpHoldtime = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bgp.router-id"); value.Exists() {
		data.BgpRouterId = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bgp.graceful-restart.graceful-reset"); value.Exists() {
		data.BgpGracefulRestartGracefulReset = types.BoolValue(true)
	} else {
		data.BgpGracefulRestartGracefulReset = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "ibgp.policy.out.enforce-modifications"); value.Exists() {
		data.IbgpPolicyOutEnforceModifications = types.BoolValue(true)
	} else {
		data.IbgpPolicyOutEnforceModifications = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "bgp.log.neighbor.changes.detail"); value.Exists() {
		data.BgpLogNeighborChangesDetail = types.BoolValue(true)
	} else {
		data.BgpLogNeighborChangesDetail = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "nexthop.validation.color-extcomm.sr-policy"); value.Exists() {
		data.NexthopValidationColorExtcommSrPolicy = types.BoolValue(true)
	} else {
		data.NexthopValidationColorExtcommSrPolicy = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "nexthop.validation.color-extcomm.disable"); value.Exists() {
		data.NexthopValidationColorExtcommDisable = types.BoolValue(true)
	} else {
		data.NexthopValidationColorExtcommDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "neighbors.neighbor"); value.Exists() {
		data.Neighbors = make([]RouterBGPNeighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterBGPNeighbors{}
			if cValue := v.Get("neighbor-address"); cValue.Exists() {
				item.NeighborAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("remote-as"); cValue.Exists() {
				item.RemoteAs = types.StringValue(cValue.String())
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			}
			if cValue := v.Get("use.neighbor-group"); cValue.Exists() {
				item.UseNeighborGroup = types.StringValue(cValue.String())
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
	if value := gjson.GetBytes(res, "neighbor-groups.neighbor-group"); value.Exists() {
		data.NeighborGroups = make([]RouterBGPNeighborGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterBGPNeighborGroups{}
			if cValue := v.Get("neighbor-group-name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("remote-as"); cValue.Exists() {
				item.RemoteAs = types.StringValue(cValue.String())
			}
			if cValue := v.Get("update-source"); cValue.Exists() {
				item.UpdateSource = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ao.key-chain-name"); cValue.Exists() {
				item.AoKeyChainName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ao.include-tcp-options.enable"); cValue.Exists() {
				item.AoIncludeTcpOptionsEnable = types.BoolValue(true)
			} else {
				item.AoIncludeTcpOptionsEnable = types.BoolValue(false)
			}
			if cValue := v.Get("bfd.minimum-interval"); cValue.Exists() {
				item.BfdMinimumInterval = types.Int64Value(cValue.Int())
			}
			data.NeighborGroups = append(data.NeighborGroups, item)
			return true
		})
	}
}

func (data *RouterBGP) getDeletedListItems(ctx context.Context, state RouterBGP) []string {
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
	for i := range state.NeighborGroups {
		keys := [...]string{"neighbor-group-name"}
		stateKeyValues := [...]string{state.NeighborGroups[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.NeighborGroups[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.NeighborGroups {
			found = true
			if state.NeighborGroups[i].Name.ValueString() != data.NeighborGroups[j].Name.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbor-groups/neighbor-group%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterBGP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.DefaultInformationOriginate.IsNull() && !data.DefaultInformationOriginate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/default-information/originate", data.getPath()))
	}
	if !data.NsrDisable.IsNull() && !data.NsrDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/nsr/disable", data.getPath()))
	}
	if !data.BgpRedistributeInternal.IsNull() && !data.BgpRedistributeInternal.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/bgp/redistribute-internal", data.getPath()))
	}
	if !data.BgpGracefulRestartGracefulReset.IsNull() && !data.BgpGracefulRestartGracefulReset.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/bgp/graceful-restart/graceful-reset", data.getPath()))
	}
	if !data.IbgpPolicyOutEnforceModifications.IsNull() && !data.IbgpPolicyOutEnforceModifications.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ibgp/policy/out/enforce-modifications", data.getPath()))
	}
	if !data.BgpLogNeighborChangesDetail.IsNull() && !data.BgpLogNeighborChangesDetail.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/bgp/log/neighbor/changes/detail", data.getPath()))
	}
	if !data.NexthopValidationColorExtcommSrPolicy.IsNull() && !data.NexthopValidationColorExtcommSrPolicy.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/nexthop/validation/color-extcomm/sr-policy", data.getPath()))
	}
	if !data.NexthopValidationColorExtcommDisable.IsNull() && !data.NexthopValidationColorExtcommDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/nexthop/validation/color-extcomm/disable", data.getPath()))
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
	for i := range data.NeighborGroups {
		keys := [...]string{"neighbor-group-name"}
		keyValues := [...]string{data.NeighborGroups[i].Name.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.NeighborGroups[i].AoIncludeTcpOptionsEnable.IsNull() && !data.NeighborGroups[i].AoIncludeTcpOptionsEnable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/neighbor-groups/neighbor-group%v/ao/include-tcp-options/enable", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

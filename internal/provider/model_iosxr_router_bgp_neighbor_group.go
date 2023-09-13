// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

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

type RouterBGPNeighborGroup struct {
	Device                            types.String                            `tfsdk:"device"`
	Id                                types.String                            `tfsdk:"id"`
	DeleteMode                        types.String                            `tfsdk:"delete_mode"`
	AsNumber                          types.String                            `tfsdk:"as_number"`
	Name                              types.String                            `tfsdk:"name"`
	RemoteAs                          types.String                            `tfsdk:"remote_as"`
	UpdateSource                      types.String                            `tfsdk:"update_source"`
	AdvertisementIntervalSeconds      types.Int64                             `tfsdk:"advertisement_interval_seconds"`
	AdvertisementIntervalMilliseconds types.Int64                             `tfsdk:"advertisement_interval_milliseconds"`
	AoKeyChainName                    types.String                            `tfsdk:"ao_key_chain_name"`
	AoIncludeTcpOptionsEnable         types.Bool                              `tfsdk:"ao_include_tcp_options_enable"`
	BfdMinimumInterval                types.Int64                             `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier                     types.Int64                             `tfsdk:"bfd_multiplier"`
	BfdFastDetect                     types.Bool                              `tfsdk:"bfd_fast_detect"`
	BfdFastDetectStrictMode           types.Bool                              `tfsdk:"bfd_fast_detect_strict_mode"`
	BfdFastDetectInheritanceDisable   types.Bool                              `tfsdk:"bfd_fast_detect_inheritance_disable"`
	AddressFamilies                   []RouterBGPNeighborGroupAddressFamilies `tfsdk:"address_families"`
}

type RouterBGPNeighborGroupData struct {
	Device                            types.String                            `tfsdk:"device"`
	Id                                types.String                            `tfsdk:"id"`
	AsNumber                          types.String                            `tfsdk:"as_number"`
	Name                              types.String                            `tfsdk:"name"`
	RemoteAs                          types.String                            `tfsdk:"remote_as"`
	UpdateSource                      types.String                            `tfsdk:"update_source"`
	AdvertisementIntervalSeconds      types.Int64                             `tfsdk:"advertisement_interval_seconds"`
	AdvertisementIntervalMilliseconds types.Int64                             `tfsdk:"advertisement_interval_milliseconds"`
	AoKeyChainName                    types.String                            `tfsdk:"ao_key_chain_name"`
	AoIncludeTcpOptionsEnable         types.Bool                              `tfsdk:"ao_include_tcp_options_enable"`
	BfdMinimumInterval                types.Int64                             `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier                     types.Int64                             `tfsdk:"bfd_multiplier"`
	BfdFastDetect                     types.Bool                              `tfsdk:"bfd_fast_detect"`
	BfdFastDetectStrictMode           types.Bool                              `tfsdk:"bfd_fast_detect_strict_mode"`
	BfdFastDetectInheritanceDisable   types.Bool                              `tfsdk:"bfd_fast_detect_inheritance_disable"`
	AddressFamilies                   []RouterBGPNeighborGroupAddressFamilies `tfsdk:"address_families"`
}
type RouterBGPNeighborGroupAddressFamilies struct {
	AfName                                 types.String `tfsdk:"af_name"`
	SoftReconfigurationInboundAlways       types.Bool   `tfsdk:"soft_reconfiguration_inbound_always"`
	NextHopSelfInheritanceDisable          types.Bool   `tfsdk:"next_hop_self_inheritance_disable"`
	RouteReflectorClient                   types.Bool   `tfsdk:"route_reflector_client"`
	RouteReflectorClientInheritanceDisable types.Bool   `tfsdk:"route_reflector_client_inheritance_disable"`
}

func (data RouterBGPNeighborGroup) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/neighbor-groups/neighbor-group[neighbor-group-name=%s]", data.AsNumber.ValueString(), data.Name.ValueString())
}

func (data RouterBGPNeighborGroupData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/neighbor-groups/neighbor-group[neighbor-group-name=%s]", data.AsNumber.ValueString(), data.Name.ValueString())
}

func (data RouterBGPNeighborGroup) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "neighbor-group-name", data.Name.ValueString())
	}
	if !data.RemoteAs.IsNull() && !data.RemoteAs.IsUnknown() {
		body, _ = sjson.Set(body, "remote-as", data.RemoteAs.ValueString())
	}
	if !data.UpdateSource.IsNull() && !data.UpdateSource.IsUnknown() {
		body, _ = sjson.Set(body, "update-source", data.UpdateSource.ValueString())
	}
	if !data.AdvertisementIntervalSeconds.IsNull() && !data.AdvertisementIntervalSeconds.IsUnknown() {
		body, _ = sjson.Set(body, "advertisement-interval.time-in-seconds", strconv.FormatInt(data.AdvertisementIntervalSeconds.ValueInt64(), 10))
	}
	if !data.AdvertisementIntervalMilliseconds.IsNull() && !data.AdvertisementIntervalMilliseconds.IsUnknown() {
		body, _ = sjson.Set(body, "advertisement-interval.time-in-milliseconds", strconv.FormatInt(data.AdvertisementIntervalMilliseconds.ValueInt64(), 10))
	}
	if !data.AoKeyChainName.IsNull() && !data.AoKeyChainName.IsUnknown() {
		body, _ = sjson.Set(body, "ao.key-chain-name", data.AoKeyChainName.ValueString())
	}
	if !data.AoIncludeTcpOptionsEnable.IsNull() && !data.AoIncludeTcpOptionsEnable.IsUnknown() {
		if data.AoIncludeTcpOptionsEnable.ValueBool() {
			body, _ = sjson.Set(body, "ao.include-tcp-options.enable", map[string]string{})
		}
	}
	if !data.BfdMinimumInterval.IsNull() && !data.BfdMinimumInterval.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.ValueInt64(), 10))
	}
	if !data.BfdMultiplier.IsNull() && !data.BfdMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.ValueInt64(), 10))
	}
	if !data.BfdFastDetect.IsNull() && !data.BfdFastDetect.IsUnknown() {
		if data.BfdFastDetect.ValueBool() {
			body, _ = sjson.Set(body, "bfd.fast-detect", map[string]string{})
		}
	}
	if !data.BfdFastDetectStrictMode.IsNull() && !data.BfdFastDetectStrictMode.IsUnknown() {
		if data.BfdFastDetectStrictMode.ValueBool() {
			body, _ = sjson.Set(body, "bfd.fast-detect.strict-mode", map[string]string{})
		}
	}
	if !data.BfdFastDetectInheritanceDisable.IsNull() && !data.BfdFastDetectInheritanceDisable.IsUnknown() {
		if data.BfdFastDetectInheritanceDisable.ValueBool() {
			body, _ = sjson.Set(body, "bfd.fast-detect.inheritance-disable", map[string]string{})
		}
	}
	if len(data.AddressFamilies) > 0 {
		body, _ = sjson.Set(body, "address-families.address-family", []interface{}{})
		for index, item := range data.AddressFamilies {
			if !item.AfName.IsNull() && !item.AfName.IsUnknown() {
				body, _ = sjson.Set(body, "address-families.address-family"+"."+strconv.Itoa(index)+"."+"af-name", item.AfName.ValueString())
			}
			if !item.SoftReconfigurationInboundAlways.IsNull() && !item.SoftReconfigurationInboundAlways.IsUnknown() {
				if item.SoftReconfigurationInboundAlways.ValueBool() {
					body, _ = sjson.Set(body, "address-families.address-family"+"."+strconv.Itoa(index)+"."+"soft-reconfiguration.inbound.always", map[string]string{})
				}
			}
			if !item.NextHopSelfInheritanceDisable.IsNull() && !item.NextHopSelfInheritanceDisable.IsUnknown() {
				if item.NextHopSelfInheritanceDisable.ValueBool() {
					body, _ = sjson.Set(body, "address-families.address-family"+"."+strconv.Itoa(index)+"."+"next-hop-self.inheritance-disable", map[string]string{})
				}
			}
			if !item.RouteReflectorClient.IsNull() && !item.RouteReflectorClient.IsUnknown() {
				if item.RouteReflectorClient.ValueBool() {
					body, _ = sjson.Set(body, "address-families.address-family"+"."+strconv.Itoa(index)+"."+"route-reflector-client", map[string]string{})
				}
			}
			if !item.RouteReflectorClientInheritanceDisable.IsNull() && !item.RouteReflectorClientInheritanceDisable.IsUnknown() {
				if item.RouteReflectorClientInheritanceDisable.ValueBool() {
					body, _ = sjson.Set(body, "address-families.address-family"+"."+strconv.Itoa(index)+"."+"route-reflector-client.inheritance-disable", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *RouterBGPNeighborGroup) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "remote-as"); value.Exists() && !data.RemoteAs.IsNull() {
		data.RemoteAs = types.StringValue(value.String())
	} else {
		data.RemoteAs = types.StringNull()
	}
	if value := gjson.GetBytes(res, "update-source"); value.Exists() && !data.UpdateSource.IsNull() {
		data.UpdateSource = types.StringValue(value.String())
	} else {
		data.UpdateSource = types.StringNull()
	}
	if value := gjson.GetBytes(res, "advertisement-interval.time-in-seconds"); value.Exists() && !data.AdvertisementIntervalSeconds.IsNull() {
		data.AdvertisementIntervalSeconds = types.Int64Value(value.Int())
	} else {
		data.AdvertisementIntervalSeconds = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "advertisement-interval.time-in-milliseconds"); value.Exists() && !data.AdvertisementIntervalMilliseconds.IsNull() {
		data.AdvertisementIntervalMilliseconds = types.Int64Value(value.Int())
	} else {
		data.AdvertisementIntervalMilliseconds = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "ao.key-chain-name"); value.Exists() && !data.AoKeyChainName.IsNull() {
		data.AoKeyChainName = types.StringValue(value.String())
	} else {
		data.AoKeyChainName = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ao.include-tcp-options.enable"); !data.AoIncludeTcpOptionsEnable.IsNull() {
		if value.Exists() {
			data.AoIncludeTcpOptionsEnable = types.BoolValue(true)
		} else {
			data.AoIncludeTcpOptionsEnable = types.BoolValue(false)
		}
	} else {
		data.AoIncludeTcpOptionsEnable = types.BoolNull()
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
	if value := gjson.GetBytes(res, "bfd.fast-detect"); !data.BfdFastDetect.IsNull() {
		if value.Exists() {
			data.BfdFastDetect = types.BoolValue(true)
		} else {
			data.BfdFastDetect = types.BoolValue(false)
		}
	} else {
		data.BfdFastDetect = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.strict-mode"); !data.BfdFastDetectStrictMode.IsNull() {
		if value.Exists() {
			data.BfdFastDetectStrictMode = types.BoolValue(true)
		} else {
			data.BfdFastDetectStrictMode = types.BoolValue(false)
		}
	} else {
		data.BfdFastDetectStrictMode = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.inheritance-disable"); !data.BfdFastDetectInheritanceDisable.IsNull() {
		if value.Exists() {
			data.BfdFastDetectInheritanceDisable = types.BoolValue(true)
		} else {
			data.BfdFastDetectInheritanceDisable = types.BoolValue(false)
		}
	} else {
		data.BfdFastDetectInheritanceDisable = types.BoolNull()
	}
	for i := range data.AddressFamilies {
		keys := [...]string{"af-name"}
		keyValues := [...]string{data.AddressFamilies[i].AfName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "address-families.address-family").ForEach(
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
		if value := r.Get("af-name"); value.Exists() && !data.AddressFamilies[i].AfName.IsNull() {
			data.AddressFamilies[i].AfName = types.StringValue(value.String())
		} else {
			data.AddressFamilies[i].AfName = types.StringNull()
		}
		if value := r.Get("soft-reconfiguration.inbound.always"); !data.AddressFamilies[i].SoftReconfigurationInboundAlways.IsNull() {
			if value.Exists() {
				data.AddressFamilies[i].SoftReconfigurationInboundAlways = types.BoolValue(true)
			} else {
				data.AddressFamilies[i].SoftReconfigurationInboundAlways = types.BoolValue(false)
			}
		} else {
			data.AddressFamilies[i].SoftReconfigurationInboundAlways = types.BoolNull()
		}
		if value := r.Get("next-hop-self.inheritance-disable"); !data.AddressFamilies[i].NextHopSelfInheritanceDisable.IsNull() {
			if value.Exists() {
				data.AddressFamilies[i].NextHopSelfInheritanceDisable = types.BoolValue(true)
			} else {
				data.AddressFamilies[i].NextHopSelfInheritanceDisable = types.BoolValue(false)
			}
		} else {
			data.AddressFamilies[i].NextHopSelfInheritanceDisable = types.BoolNull()
		}
		if value := r.Get("route-reflector-client"); !data.AddressFamilies[i].RouteReflectorClient.IsNull() {
			if value.Exists() {
				data.AddressFamilies[i].RouteReflectorClient = types.BoolValue(true)
			} else {
				data.AddressFamilies[i].RouteReflectorClient = types.BoolValue(false)
			}
		} else {
			data.AddressFamilies[i].RouteReflectorClient = types.BoolNull()
		}
		if value := r.Get("route-reflector-client.inheritance-disable"); !data.AddressFamilies[i].RouteReflectorClientInheritanceDisable.IsNull() {
			if value.Exists() {
				data.AddressFamilies[i].RouteReflectorClientInheritanceDisable = types.BoolValue(true)
			} else {
				data.AddressFamilies[i].RouteReflectorClientInheritanceDisable = types.BoolValue(false)
			}
		} else {
			data.AddressFamilies[i].RouteReflectorClientInheritanceDisable = types.BoolNull()
		}
	}
}

func (data *RouterBGPNeighborGroupData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "remote-as"); value.Exists() {
		data.RemoteAs = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "update-source"); value.Exists() {
		data.UpdateSource = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "advertisement-interval.time-in-seconds"); value.Exists() {
		data.AdvertisementIntervalSeconds = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "advertisement-interval.time-in-milliseconds"); value.Exists() {
		data.AdvertisementIntervalMilliseconds = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "ao.key-chain-name"); value.Exists() {
		data.AoKeyChainName = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ao.include-tcp-options.enable"); value.Exists() {
		data.AoIncludeTcpOptionsEnable = types.BoolValue(true)
	} else {
		data.AoIncludeTcpOptionsEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect"); value.Exists() {
		data.BfdFastDetect = types.BoolValue(true)
	} else {
		data.BfdFastDetect = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.strict-mode"); value.Exists() {
		data.BfdFastDetectStrictMode = types.BoolValue(true)
	} else {
		data.BfdFastDetectStrictMode = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.inheritance-disable"); value.Exists() {
		data.BfdFastDetectInheritanceDisable = types.BoolValue(true)
	} else {
		data.BfdFastDetectInheritanceDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "address-families.address-family"); value.Exists() {
		data.AddressFamilies = make([]RouterBGPNeighborGroupAddressFamilies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterBGPNeighborGroupAddressFamilies{}
			if cValue := v.Get("af-name"); cValue.Exists() {
				item.AfName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("soft-reconfiguration.inbound.always"); cValue.Exists() {
				item.SoftReconfigurationInboundAlways = types.BoolValue(true)
			} else {
				item.SoftReconfigurationInboundAlways = types.BoolValue(false)
			}
			if cValue := v.Get("next-hop-self.inheritance-disable"); cValue.Exists() {
				item.NextHopSelfInheritanceDisable = types.BoolValue(true)
			} else {
				item.NextHopSelfInheritanceDisable = types.BoolValue(false)
			}
			if cValue := v.Get("route-reflector-client"); cValue.Exists() {
				item.RouteReflectorClient = types.BoolValue(true)
			} else {
				item.RouteReflectorClient = types.BoolValue(false)
			}
			if cValue := v.Get("route-reflector-client.inheritance-disable"); cValue.Exists() {
				item.RouteReflectorClientInheritanceDisable = types.BoolValue(true)
			} else {
				item.RouteReflectorClientInheritanceDisable = types.BoolValue(false)
			}
			data.AddressFamilies = append(data.AddressFamilies, item)
			return true
		})
	}
}

func (data *RouterBGPNeighborGroup) getDeletedListItems(ctx context.Context, state RouterBGPNeighborGroup) []string {
	deletedListItems := make([]string, 0)
	for i := range state.AddressFamilies {
		keys := [...]string{"af-name"}
		stateKeyValues := [...]string{state.AddressFamilies[i].AfName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.AddressFamilies[i].AfName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.AddressFamilies {
			found = true
			if state.AddressFamilies[i].AfName.ValueString() != data.AddressFamilies[j].AfName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/address-families/address-family%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterBGPNeighborGroup) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.AoIncludeTcpOptionsEnable.IsNull() && !data.AoIncludeTcpOptionsEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ao/include-tcp-options/enable", data.getPath()))
	}
	if !data.BfdFastDetect.IsNull() && !data.BfdFastDetect.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/bfd/fast-detect", data.getPath()))
	}
	if !data.BfdFastDetectStrictMode.IsNull() && !data.BfdFastDetectStrictMode.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/bfd/fast-detect/strict-mode", data.getPath()))
	}
	if !data.BfdFastDetectInheritanceDisable.IsNull() && !data.BfdFastDetectInheritanceDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/bfd/fast-detect/inheritance-disable", data.getPath()))
	}
	for i := range data.AddressFamilies {
		keys := [...]string{"af-name"}
		keyValues := [...]string{data.AddressFamilies[i].AfName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.AddressFamilies[i].SoftReconfigurationInboundAlways.IsNull() && !data.AddressFamilies[i].SoftReconfigurationInboundAlways.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address-families/address-family%v/soft-reconfiguration/inbound/always", data.getPath(), keyString))
		}
		if !data.AddressFamilies[i].NextHopSelfInheritanceDisable.IsNull() && !data.AddressFamilies[i].NextHopSelfInheritanceDisable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address-families/address-family%v/next-hop-self/inheritance-disable", data.getPath(), keyString))
		}
		if !data.AddressFamilies[i].RouteReflectorClient.IsNull() && !data.AddressFamilies[i].RouteReflectorClient.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address-families/address-family%v/route-reflector-client", data.getPath(), keyString))
		}
		if !data.AddressFamilies[i].RouteReflectorClientInheritanceDisable.IsNull() && !data.AddressFamilies[i].RouteReflectorClientInheritanceDisable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address-families/address-family%v/route-reflector-client/inheritance-disable", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

func (data *RouterBGPNeighborGroup) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.RemoteAs.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/remote-as", data.getPath()))
	}
	if !data.UpdateSource.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/update-source", data.getPath()))
	}
	if !data.AdvertisementIntervalSeconds.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/advertisement-interval", data.getPath()))
	}
	if !data.AdvertisementIntervalMilliseconds.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/advertisement-interval", data.getPath()))
	}
	if !data.AoKeyChainName.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ao/key-chain-name", data.getPath()))
	}
	if !data.AoIncludeTcpOptionsEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ao/include-tcp-options/enable", data.getPath()))
	}
	if !data.BfdMinimumInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/minimum-interval", data.getPath()))
	}
	if !data.BfdMultiplier.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/multiplier", data.getPath()))
	}
	if !data.BfdFastDetect.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect", data.getPath()))
	}
	if !data.BfdFastDetectStrictMode.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/strict-mode", data.getPath()))
	}
	if !data.BfdFastDetectInheritanceDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/inheritance-disable", data.getPath()))
	}
	for i := range data.AddressFamilies {
		keys := [...]string{"af-name"}
		keyValues := [...]string{data.AddressFamilies[i].AfName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address-families/address-family%v", data.getPath(), keyString))
	}
	return deletePaths
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Domain struct {
	Device                types.String        `tfsdk:"device"`
	Id                    types.String        `tfsdk:"id"`
	DeleteMode            types.String        `tfsdk:"delete_mode"`
	Domains               []DomainDomains     `tfsdk:"domains"`
	LookupDisable         types.Bool          `tfsdk:"lookup_disable"`
	LookupSourceInterface types.String        `tfsdk:"lookup_source_interface"`
	Name                  types.String        `tfsdk:"name"`
	Ipv4Hosts             []DomainIpv4Hosts   `tfsdk:"ipv4_hosts"`
	NameServers           []DomainNameServers `tfsdk:"name_servers"`
	Ipv6Hosts             []DomainIpv6Hosts   `tfsdk:"ipv6_hosts"`
	Multicast             types.String        `tfsdk:"multicast"`
	DefaultFlowsDisable   types.Bool          `tfsdk:"default_flows_disable"`
}

type DomainData struct {
	Device                types.String        `tfsdk:"device"`
	Id                    types.String        `tfsdk:"id"`
	Domains               []DomainDomains     `tfsdk:"domains"`
	LookupDisable         types.Bool          `tfsdk:"lookup_disable"`
	LookupSourceInterface types.String        `tfsdk:"lookup_source_interface"`
	Name                  types.String        `tfsdk:"name"`
	Ipv4Hosts             []DomainIpv4Hosts   `tfsdk:"ipv4_hosts"`
	NameServers           []DomainNameServers `tfsdk:"name_servers"`
	Ipv6Hosts             []DomainIpv6Hosts   `tfsdk:"ipv6_hosts"`
	Multicast             types.String        `tfsdk:"multicast"`
	DefaultFlowsDisable   types.Bool          `tfsdk:"default_flows_disable"`
}
type DomainDomains struct {
	DomainName types.String `tfsdk:"domain_name"`
	Order      types.Int64  `tfsdk:"order"`
}
type DomainIpv4Hosts struct {
	HostName  types.String `tfsdk:"host_name"`
	IpAddress types.List   `tfsdk:"ip_address"`
}
type DomainNameServers struct {
	Address types.String `tfsdk:"address"`
	Order   types.Int64  `tfsdk:"order"`
}
type DomainIpv6Hosts struct {
	HostName    types.String `tfsdk:"host_name"`
	Ipv6Address types.List   `tfsdk:"ipv6_address"`
}

func (data Domain) getPath() string {
	return "Cisco-IOS-XR-um-domain-cfg:/domain"
}

func (data DomainData) getPath() string {
	return "Cisco-IOS-XR-um-domain-cfg:/domain"
}

func (data Domain) toBody(ctx context.Context) string {
	body := "{}"
	if !data.LookupDisable.IsNull() && !data.LookupDisable.IsUnknown() {
		if data.LookupDisable.ValueBool() {
			body, _ = sjson.Set(body, "lookup.disable", map[string]string{})
		}
	}
	if !data.LookupSourceInterface.IsNull() && !data.LookupSourceInterface.IsUnknown() {
		body, _ = sjson.Set(body, "lookup.source-interface", data.LookupSourceInterface.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Multicast.IsNull() && !data.Multicast.IsUnknown() {
		body, _ = sjson.Set(body, "multicast", data.Multicast.ValueString())
	}
	if !data.DefaultFlowsDisable.IsNull() && !data.DefaultFlowsDisable.IsUnknown() {
		if data.DefaultFlowsDisable.ValueBool() {
			body, _ = sjson.Set(body, "default-flows.disable", map[string]string{})
		}
	}
	if len(data.Domains) > 0 {
		body, _ = sjson.Set(body, "list.domain", []interface{}{})
		for index, item := range data.Domains {
			if !item.DomainName.IsNull() && !item.DomainName.IsUnknown() {
				body, _ = sjson.Set(body, "list.domain"+"."+strconv.Itoa(index)+"."+"domain-name", item.DomainName.ValueString())
			}
			if !item.Order.IsNull() && !item.Order.IsUnknown() {
				body, _ = sjson.Set(body, "list.domain"+"."+strconv.Itoa(index)+"."+"order", strconv.FormatInt(item.Order.ValueInt64(), 10))
			}
		}
	}
	if len(data.Ipv4Hosts) > 0 {
		body, _ = sjson.Set(body, "ipv4.hosts.host", []interface{}{})
		for index, item := range data.Ipv4Hosts {
			if !item.HostName.IsNull() && !item.HostName.IsUnknown() {
				body, _ = sjson.Set(body, "ipv4.hosts.host"+"."+strconv.Itoa(index)+"."+"host-name", item.HostName.ValueString())
			}
			if !item.IpAddress.IsNull() && !item.IpAddress.IsUnknown() {
				var values []string
				item.IpAddress.ElementsAs(ctx, &values, false)
				body, _ = sjson.Set(body, "ipv4.hosts.host"+"."+strconv.Itoa(index)+"."+"ip-address", values)
			}
		}
	}
	if len(data.NameServers) > 0 {
		body, _ = sjson.Set(body, "name-servers.name-server", []interface{}{})
		for index, item := range data.NameServers {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "name-servers.name-server"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Order.IsNull() && !item.Order.IsUnknown() {
				body, _ = sjson.Set(body, "name-servers.name-server"+"."+strconv.Itoa(index)+"."+"order", strconv.FormatInt(item.Order.ValueInt64(), 10))
			}
		}
	}
	if len(data.Ipv6Hosts) > 0 {
		body, _ = sjson.Set(body, "ipv6.host.host", []interface{}{})
		for index, item := range data.Ipv6Hosts {
			if !item.HostName.IsNull() && !item.HostName.IsUnknown() {
				body, _ = sjson.Set(body, "ipv6.host.host"+"."+strconv.Itoa(index)+"."+"host-name", item.HostName.ValueString())
			}
			if !item.Ipv6Address.IsNull() && !item.Ipv6Address.IsUnknown() {
				var values []string
				item.Ipv6Address.ElementsAs(ctx, &values, false)
				body, _ = sjson.Set(body, "ipv6.host.host"+"."+strconv.Itoa(index)+"."+"ipv6-address", values)
			}
		}
	}
	return body
}

func (data *Domain) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.Domains {
		keys := [...]string{"domain-name", "order"}
		keyValues := [...]string{data.Domains[i].DomainName.ValueString(), strconv.FormatInt(data.Domains[i].Order.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "list.domain").ForEach(
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
		if value := r.Get("domain-name"); value.Exists() && !data.Domains[i].DomainName.IsNull() {
			data.Domains[i].DomainName = types.StringValue(value.String())
		} else {
			data.Domains[i].DomainName = types.StringNull()
		}
		if value := r.Get("order"); value.Exists() && !data.Domains[i].Order.IsNull() {
			data.Domains[i].Order = types.Int64Value(value.Int())
		} else {
			data.Domains[i].Order = types.Int64Null()
		}
	}
	if value := gjson.GetBytes(res, "lookup.disable"); !data.LookupDisable.IsNull() {
		if value.Exists() {
			data.LookupDisable = types.BoolValue(true)
		} else {
			data.LookupDisable = types.BoolValue(false)
		}
	} else {
		data.LookupDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "lookup.source-interface"); value.Exists() && !data.LookupSourceInterface.IsNull() {
		data.LookupSourceInterface = types.StringValue(value.String())
	} else {
		data.LookupSourceInterface = types.StringNull()
	}
	if value := gjson.GetBytes(res, "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := range data.Ipv4Hosts {
		keys := [...]string{"host-name"}
		keyValues := [...]string{data.Ipv4Hosts[i].HostName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "ipv4.hosts.host").ForEach(
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
		if value := r.Get("host-name"); value.Exists() && !data.Ipv4Hosts[i].HostName.IsNull() {
			data.Ipv4Hosts[i].HostName = types.StringValue(value.String())
		} else {
			data.Ipv4Hosts[i].HostName = types.StringNull()
		}
		if value := r.Get("ip-address"); value.Exists() && !data.Ipv4Hosts[i].IpAddress.IsNull() {
			data.Ipv4Hosts[i].IpAddress = helpers.GetStringList(value.Array())
		} else {
			data.Ipv4Hosts[i].IpAddress = types.ListNull(types.StringType)
		}
	}
	for i := range data.NameServers {
		keys := [...]string{"address", "order"}
		keyValues := [...]string{data.NameServers[i].Address.ValueString(), strconv.FormatInt(data.NameServers[i].Order.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "name-servers.name-server").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.NameServers[i].Address.IsNull() {
			data.NameServers[i].Address = types.StringValue(value.String())
		} else {
			data.NameServers[i].Address = types.StringNull()
		}
		if value := r.Get("order"); value.Exists() && !data.NameServers[i].Order.IsNull() {
			data.NameServers[i].Order = types.Int64Value(value.Int())
		} else {
			data.NameServers[i].Order = types.Int64Null()
		}
	}
	for i := range data.Ipv6Hosts {
		keys := [...]string{"host-name"}
		keyValues := [...]string{data.Ipv6Hosts[i].HostName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "ipv6.host.host").ForEach(
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
		if value := r.Get("host-name"); value.Exists() && !data.Ipv6Hosts[i].HostName.IsNull() {
			data.Ipv6Hosts[i].HostName = types.StringValue(value.String())
		} else {
			data.Ipv6Hosts[i].HostName = types.StringNull()
		}
		if value := r.Get("ipv6-address"); value.Exists() && !data.Ipv6Hosts[i].Ipv6Address.IsNull() {
			data.Ipv6Hosts[i].Ipv6Address = helpers.GetStringList(value.Array())
		} else {
			data.Ipv6Hosts[i].Ipv6Address = types.ListNull(types.StringType)
		}
	}
	if value := gjson.GetBytes(res, "multicast"); value.Exists() && !data.Multicast.IsNull() {
		data.Multicast = types.StringValue(value.String())
	} else {
		data.Multicast = types.StringNull()
	}
	if value := gjson.GetBytes(res, "default-flows.disable"); !data.DefaultFlowsDisable.IsNull() {
		if value.Exists() {
			data.DefaultFlowsDisable = types.BoolValue(true)
		} else {
			data.DefaultFlowsDisable = types.BoolValue(false)
		}
	} else {
		data.DefaultFlowsDisable = types.BoolNull()
	}
}

func (data *DomainData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "list.domain"); value.Exists() {
		data.Domains = make([]DomainDomains, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DomainDomains{}
			if cValue := v.Get("domain-name"); cValue.Exists() {
				item.DomainName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("order"); cValue.Exists() {
				item.Order = types.Int64Value(cValue.Int())
			}
			data.Domains = append(data.Domains, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "lookup.disable"); value.Exists() {
		data.LookupDisable = types.BoolValue(true)
	} else {
		data.LookupDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "lookup.source-interface"); value.Exists() {
		data.LookupSourceInterface = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv4.hosts.host"); value.Exists() {
		data.Ipv4Hosts = make([]DomainIpv4Hosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DomainIpv4Hosts{}
			if cValue := v.Get("host-name"); cValue.Exists() {
				item.HostName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ip-address"); cValue.Exists() {
				item.IpAddress = helpers.GetStringList(cValue.Array())
			} else {
				item.IpAddress = types.ListNull(types.StringType)
			}
			data.Ipv4Hosts = append(data.Ipv4Hosts, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "name-servers.name-server"); value.Exists() {
		data.NameServers = make([]DomainNameServers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DomainNameServers{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("order"); cValue.Exists() {
				item.Order = types.Int64Value(cValue.Int())
			}
			data.NameServers = append(data.NameServers, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "ipv6.host.host"); value.Exists() {
		data.Ipv6Hosts = make([]DomainIpv6Hosts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DomainIpv6Hosts{}
			if cValue := v.Get("host-name"); cValue.Exists() {
				item.HostName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ipv6-address"); cValue.Exists() {
				item.Ipv6Address = helpers.GetStringList(cValue.Array())
			} else {
				item.Ipv6Address = types.ListNull(types.StringType)
			}
			data.Ipv6Hosts = append(data.Ipv6Hosts, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "multicast"); value.Exists() {
		data.Multicast = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "default-flows.disable"); value.Exists() {
		data.DefaultFlowsDisable = types.BoolValue(true)
	} else {
		data.DefaultFlowsDisable = types.BoolValue(false)
	}
}

func (data *Domain) getDeletedListItems(ctx context.Context, state Domain) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Domains {
		keys := [...]string{"domain-name", "order"}
		stateKeyValues := [...]string{state.Domains[i].DomainName.ValueString(), strconv.FormatInt(state.Domains[i].Order.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Domains[i].DomainName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Domains[i].Order.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Domains {
			found = true
			if state.Domains[i].DomainName.ValueString() != data.Domains[j].DomainName.ValueString() {
				found = false
			}
			if state.Domains[i].Order.ValueInt64() != data.Domains[j].Order.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/list/domain%v", state.getPath(), keyString))
		}
	}
	for i := range state.Ipv4Hosts {
		keys := [...]string{"host-name"}
		stateKeyValues := [...]string{state.Ipv4Hosts[i].HostName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv4Hosts[i].HostName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv4Hosts {
			found = true
			if state.Ipv4Hosts[i].HostName.ValueString() != data.Ipv4Hosts[j].HostName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ipv4/hosts/host%v", state.getPath(), keyString))
		}
	}
	for i := range state.NameServers {
		keys := [...]string{"address", "order"}
		stateKeyValues := [...]string{state.NameServers[i].Address.ValueString(), strconv.FormatInt(state.NameServers[i].Order.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.NameServers[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.NameServers[i].Order.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.NameServers {
			found = true
			if state.NameServers[i].Address.ValueString() != data.NameServers[j].Address.ValueString() {
				found = false
			}
			if state.NameServers[i].Order.ValueInt64() != data.NameServers[j].Order.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/name-servers/name-server%v", state.getPath(), keyString))
		}
	}
	for i := range state.Ipv6Hosts {
		keys := [...]string{"host-name"}
		stateKeyValues := [...]string{state.Ipv6Hosts[i].HostName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6Hosts[i].HostName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6Hosts {
			found = true
			if state.Ipv6Hosts[i].HostName.ValueString() != data.Ipv6Hosts[j].HostName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ipv6/host/host%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *Domain) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.Domains {
		keys := [...]string{"domain-name", "order"}
		keyValues := [...]string{data.Domains[i].DomainName.ValueString(), strconv.FormatInt(data.Domains[i].Order.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	if !data.LookupDisable.IsNull() && !data.LookupDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/lookup/disable", data.getPath()))
	}
	for i := range data.Ipv4Hosts {
		keys := [...]string{"host-name"}
		keyValues := [...]string{data.Ipv4Hosts[i].HostName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.NameServers {
		keys := [...]string{"address", "order"}
		keyValues := [...]string{data.NameServers[i].Address.ValueString(), strconv.FormatInt(data.NameServers[i].Order.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.Ipv6Hosts {
		keys := [...]string{"host-name"}
		keyValues := [...]string{data.Ipv6Hosts[i].HostName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	if !data.DefaultFlowsDisable.IsNull() && !data.DefaultFlowsDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/default-flows/disable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *Domain) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.Domains {
		keys := [...]string{"domain-name", "order"}
		keyValues := [...]string{data.Domains[i].DomainName.ValueString(), strconv.FormatInt(data.Domains[i].Order.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/list/domain%v", data.getPath(), keyString))
	}
	if !data.LookupDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/lookup/disable", data.getPath()))
	}
	if !data.LookupSourceInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/lookup/source-interface", data.getPath()))
	}
	if !data.Name.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/name", data.getPath()))
	}
	for i := range data.Ipv4Hosts {
		keys := [...]string{"host-name"}
		keyValues := [...]string{data.Ipv4Hosts[i].HostName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv4/hosts/host%v", data.getPath(), keyString))
	}
	for i := range data.NameServers {
		keys := [...]string{"address", "order"}
		keyValues := [...]string{data.NameServers[i].Address.ValueString(), strconv.FormatInt(data.NameServers[i].Order.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/name-servers/name-server%v", data.getPath(), keyString))
	}
	for i := range data.Ipv6Hosts {
		keys := [...]string{"host-name"}
		keyValues := [...]string{data.Ipv6Hosts[i].HostName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/host/host%v", data.getPath(), keyString))
	}
	if !data.Multicast.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/multicast", data.getPath()))
	}
	if !data.DefaultFlowsDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/default-flows/disable", data.getPath()))
	}
	return deletePaths
}

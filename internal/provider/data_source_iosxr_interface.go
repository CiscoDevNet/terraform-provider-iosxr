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

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &InterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &InterfaceDataSource{}
)

func NewInterfaceDataSource() datasource.DataSource {
	return &InterfaceDataSource{}
}

type InterfaceDataSource struct {
	client *client.Client
}

func (d *InterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface"
}

func (d *InterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Interface configuration subcommands",
				Required:            true,
			},
			"l2transport": schema.BoolAttribute{
				MarkdownDescription: "l2transport sub-interface",
				Computed:            true,
			},
			"point_to_point": schema.BoolAttribute{
				MarkdownDescription: "point-to-point sub-interface",
				Computed:            true,
			},
			"multipoint": schema.BoolAttribute{
				MarkdownDescription: "multipoint sub-interface",
				Computed:            true,
			},
			"dampening_decay_half_life_value": schema.Int64Attribute{
				MarkdownDescription: "Decay half life (in minutes)",
				Computed:            true,
			},
			"ipv4_point_to_point": schema.BoolAttribute{
				MarkdownDescription: "Enable point-to-point handling for this interface.",
				Computed:            true,
			},
			"service_policy_input": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a policy in the input direction",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the service policy. Set 'input' for 'service-ipsec and 'service-gre' interfaces",
							Computed:            true,
						},
					},
				},
			},
			"service_policy_output": schema.ListNestedAttribute{
				MarkdownDescription: "direction of service policy application",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the service policy. Set 'output' for 'service-ipsec and 'service-gre' interfaces",
							Computed:            true,
						},
					},
				},
			},
			"bfd_mode_ietf": schema.BoolAttribute{
				MarkdownDescription: "Use IETF standard for BoB",
				Computed:            true,
			},
			"encapsulation_dot1q_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "Configure first (outer) VLAN ID on the subinterface",
				Computed:            true,
			},
			"l2transport_encapsulation_dot1q_vlan_id": schema.StringAttribute{
				MarkdownDescription: "Single VLAN id or start of VLAN range",
				Computed:            true,
			},
			"l2transport_encapsulation_dot1q_second_dot1q": schema.StringAttribute{
				MarkdownDescription: "End of VLAN range",
				Computed:            true,
			},
			"rewrite_ingress_tag_pop_one": schema.BoolAttribute{
				MarkdownDescription: "Remove outer tag only",
				Computed:            true,
			},
			"rewrite_ingress_tag_pop_two": schema.BoolAttribute{
				MarkdownDescription: "Remove two outermost tags",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "shutdown the given interface",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Set the MTU on an interface",
				Computed:            true,
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: "Set the bandwidth of an interface",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Set description for this interface",
				Computed:            true,
			},
			"load_interval": schema.Int64Attribute{
				MarkdownDescription: "Specify interval for load calculation for an interface",
				Computed:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "Set VRF in which the interface operates",
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
			},
			"ipv4_netmask": schema.StringAttribute{
				MarkdownDescription: "IP subnet mask",
				Computed:            true,
			},
			"unnumbered": schema.StringAttribute{
				MarkdownDescription: "Enable IP processing without an explicit address",
				Computed:            true,
			},
			"ipv4_verify_unicast_source_reachable_via_type": schema.StringAttribute{
				MarkdownDescription: "Source reachable type",
				Computed:            true,
			},
			"ipv4_verify_unicast_source_reachable_via_allow_self_ping": schema.BoolAttribute{
				MarkdownDescription: "Allow router to ping itself (opens vulnerability in verification)",
				Computed:            true,
			},
			"ipv4_verify_unicast_source_reachable_via_allow_default": schema.BoolAttribute{
				MarkdownDescription: "Allow default route to match when checking source address",
				Computed:            true,
			},
			"ipv4_access_group_ingress_acl1": schema.StringAttribute{
				MarkdownDescription: "Access-list name",
				Computed:            true,
			},
			"ipv4_access_group_ingress_hardware_count": schema.BoolAttribute{
				MarkdownDescription: "Count packets in hardware",
				Computed:            true,
			},
			"ipv4_access_group_ingress_interface_statistics": schema.BoolAttribute{
				MarkdownDescription: "Per interface statistics in hardware",
				Computed:            true,
			},
			"ipv4_access_group_ingress_compress_level": schema.Int64Attribute{
				MarkdownDescription: "Specify ACL compression in hardware",
				Computed:            true,
			},
			"ipv4_access_group_egress_acl": schema.StringAttribute{
				MarkdownDescription: "Access-list name",
				Computed:            true,
			},
			"ipv4_access_group_egress_hardware_count": schema.BoolAttribute{
				MarkdownDescription: "Count packets in hardware",
				Computed:            true,
			},
			"ipv4_access_group_egress_interface_statistics": schema.BoolAttribute{
				MarkdownDescription: "Per interface statistics in hardware",
				Computed:            true,
			},
			"ipv4_access_group_egress_compress_level": schema.Int64Attribute{
				MarkdownDescription: "Specify ACL compression in hardware",
				Computed:            true,
			},
			"ipv6_verify_unicast_source_reachable_via_type": schema.StringAttribute{
				MarkdownDescription: "Source reachable type",
				Computed:            true,
			},
			"ipv6_verify_unicast_source_reachable_via_allow_self_ping": schema.BoolAttribute{
				MarkdownDescription: "Allow router to ping itself (opens vulnerability in verification)",
				Computed:            true,
			},
			"ipv6_verify_unicast_source_reachable_via_allow_default": schema.BoolAttribute{
				MarkdownDescription: "Allow default route to match when checking source address",
				Computed:            true,
			},
			"ipv6_access_group_ingress_acl1": schema.StringAttribute{
				MarkdownDescription: "Access-list name",
				Computed:            true,
			},
			"ipv6_access_group_ingress_interface_statistics": schema.BoolAttribute{
				MarkdownDescription: "Per interface statistics in hardware",
				Computed:            true,
			},
			"ipv6_access_group_ingress_compress_level": schema.Int64Attribute{
				MarkdownDescription: "Specify ACL compression in hardware",
				Computed:            true,
			},
			"ipv6_access_group_egress_acl1": schema.StringAttribute{
				MarkdownDescription: "Access-list name",
				Computed:            true,
			},
			"ipv6_access_group_egress_interface_statistics": schema.BoolAttribute{
				MarkdownDescription: "Per interface statistics in hardware",
				Computed:            true,
			},
			"ipv6_access_group_egress_compress_level": schema.Int64Attribute{
				MarkdownDescription: "Specify ACL compression in hardware",
				Computed:            true,
			},
			"ipv6_link_local_address": schema.StringAttribute{
				MarkdownDescription: "IPv6 address",
				Computed:            true,
			},
			"ipv6_link_local_zone": schema.StringAttribute{
				MarkdownDescription: "IPv6 address zone",
				Computed:            true,
			},
			"ipv6_autoconfig": schema.BoolAttribute{
				MarkdownDescription: "Enable slaac on Mgmt interface",
				Computed:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable IPv6 on interface",
				Computed:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 name or address",
							Computed:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: "Prefix length in bits",
							Computed:            true,
						},
						"zone": schema.StringAttribute{
							MarkdownDescription: "IPv6 address zone",
							Computed:            true,
						},
					},
				},
			},
			"bundle_minimum_active_links": schema.Int64Attribute{
				MarkdownDescription: "Set the number of active links needed to bring up this bundle",
				Computed:            true,
			},
			"bundle_maximum_active_links": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of active links in this bundle",
				Computed:            true,
			},
			"cdp": schema.BoolAttribute{
				MarkdownDescription: "Enable CDP on an interface",
				Computed:            true,
			},
			"bundle_shutdown": schema.BoolAttribute{
				MarkdownDescription: "Bring all links in the bundle down to Standby state",
				Computed:            true,
			},
			"bundle_load_balancing_hash_src_ip": schema.BoolAttribute{
				MarkdownDescription: "Use the source IP as the hash function",
				Computed:            true,
			},
			"bundle_load_balancing_hash_dst_ip": schema.BoolAttribute{
				MarkdownDescription: "Use the destination IP as the hash function",
				Computed:            true,
			},
			"bundle_id": schema.Int64Attribute{
				MarkdownDescription: "Add the port to an aggregated interface.",
				Computed:            true,
			},
			"bundle_id_mode": schema.StringAttribute{
				MarkdownDescription: "Specify the mode of operation.",
				Computed:            true,
			},
			"bundle_port_priority": schema.Int64Attribute{
				MarkdownDescription: "Priority for this port. Lower value is higher priority.",
				Computed:            true,
			},
			"flow_ipv4_ingress_monitors": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor for packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
					},
				},
			},
			"flow_ipv4_ingress_monitor_samplers": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor and sampler for incoming packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
						"sampler_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a sampler for packets",
							Computed:            true,
						},
					},
				},
			},
			"flow_ipv4_egress_monitors": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor for packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
					},
				},
			},
			"flow_ipv4_egress_monitor_samplers": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor and sampler for outgoing packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
						"sampler_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a sampler for packets",
							Computed:            true,
						},
					},
				},
			},
			"flow_ipv6_ingress_monitors": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor for packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
					},
				},
			},
			"flow_ipv6_ingress_monitor_samplers": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor and sampler for incoming packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
						"sampler_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a sampler for packets",
							Computed:            true,
						},
					},
				},
			},
			"flow_ipv6_egress_monitors": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor for packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
					},
				},
			},
			"flow_ipv6_egress_monitor_samplers": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a flow monitor and sampler for outgoing packets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a flow monitor for packets",
							Computed:            true,
						},
						"sampler_map_name": schema.StringAttribute{
							MarkdownDescription: "Specify a sampler for packets",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *InterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *InterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config InterfaceData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, err := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Unable to apply gNMI Get operation", err.Error())
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrFlowMonitorMap(t *testing.T) {
	if os.Getenv("FLOW") == "" {
		t.Skip("skipping test, set environment variable FLOW")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "exporters.0.name", "exporter1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "option_outphysint", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "option_filtered", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "option_bgpattr", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "option_outbundlemember", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_destination", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_destination_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_as", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_protocol_port", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_prefix", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_source_prefix", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_destination_prefix", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_as_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_protocol_port_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_prefix_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_source_prefix_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_destination_prefix_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_prefix_port", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_bgp_nexthop_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_peer_as", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv4_gtp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv6_destination", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv6_peer_as", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_ipv6_gtp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_mpls_ipv4_fields", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_mpls_ipv6_fields", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_mpls_ipv4_ipv6_fields", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_mpls_labels", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_map_t", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_sflow", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_datalink_record", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_default_rtp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "record_default_mdi", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "cache_entries", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "cache_timeout_active", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "cache_timeout_inactive", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "cache_timeout_update", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "cache_timeout_rate_limit", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "cache_permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "cache_immediate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "hw_cache_timeout_inactive", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_extended_router", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_extended_gateway", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_extended_ipv4_tunnel_egress", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_extended_ipv6_tunnel_egress", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_if_counters_polling_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_sample_header_size", "128"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_input_ifindex", "physical"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_monitor_map.test", "sflow_options_output_ifindex", "physical"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrFlowMonitorMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrFlowMonitorMapConfig() string {
	config := `resource "iosxr_flow_monitor_map" "test" {` + "\n"
	config += `	name = "monitor_map1"` + "\n"
	config += `	exporters = [{` + "\n"
	config += `		name = "exporter1"` + "\n"
	config += `	}]` + "\n"
	config += `	option_outphysint = true` + "\n"
	config += `	option_filtered = true` + "\n"
	config += `	option_bgpattr = true` + "\n"
	config += `	option_outbundlemember = true` + "\n"
	config += `	record_ipv4_destination = true` + "\n"
	config += `	record_ipv4_destination_tos = true` + "\n"
	config += `	record_ipv4_as = true` + "\n"
	config += `	record_ipv4_protocol_port = true` + "\n"
	config += `	record_ipv4_prefix = true` + "\n"
	config += `	record_ipv4_source_prefix = true` + "\n"
	config += `	record_ipv4_destination_prefix = true` + "\n"
	config += `	record_ipv4_as_tos = true` + "\n"
	config += `	record_ipv4_protocol_port_tos = true` + "\n"
	config += `	record_ipv4_prefix_tos = true` + "\n"
	config += `	record_ipv4_source_prefix_tos = true` + "\n"
	config += `	record_ipv4_destination_prefix_tos = true` + "\n"
	config += `	record_ipv4_prefix_port = true` + "\n"
	config += `	record_ipv4_bgp_nexthop_tos = true` + "\n"
	config += `	record_ipv4_peer_as = true` + "\n"
	config += `	record_ipv4_gtp = true` + "\n"
	config += `	record_ipv6_destination = true` + "\n"
	config += `	record_ipv6_peer_as = true` + "\n"
	config += `	record_ipv6_gtp = true` + "\n"
	config += `	record_mpls_ipv4_fields = true` + "\n"
	config += `	record_mpls_ipv6_fields = true` + "\n"
	config += `	record_mpls_ipv4_ipv6_fields = true` + "\n"
	config += `	record_mpls_labels = 2` + "\n"
	config += `	record_map_t = true` + "\n"
	config += `	record_sflow = true` + "\n"
	config += `	record_datalink_record = true` + "\n"
	config += `	record_default_rtp = true` + "\n"
	config += `	record_default_mdi = true` + "\n"
	config += `	cache_entries = 5000` + "\n"
	config += `	cache_timeout_active = 1` + "\n"
	config += `	cache_timeout_inactive = 0` + "\n"
	config += `	cache_timeout_update = 1` + "\n"
	config += `	cache_timeout_rate_limit = 5000` + "\n"
	config += `	cache_permanent = true` + "\n"
	config += `	cache_immediate = true` + "\n"
	config += `	hw_cache_timeout_inactive = 50` + "\n"
	config += `	sflow_options_extended_router = true` + "\n"
	config += `	sflow_options_extended_gateway = true` + "\n"
	config += `	sflow_options_extended_ipv4_tunnel_egress = true` + "\n"
	config += `	sflow_options_extended_ipv6_tunnel_egress = true` + "\n"
	config += `	sflow_options_if_counters_polling_interval = 5` + "\n"
	config += `	sflow_options_sample_header_size = 128` + "\n"
	config += `	sflow_options_input_ifindex = "physical"` + "\n"
	config += `	sflow_options_output_ifindex = "physical"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_flow_monitor_map" "test" {
			name = "monitor_map1"
			depends_on = [iosxr_flow_monitor_map.test]
		}
	`
	return config
}

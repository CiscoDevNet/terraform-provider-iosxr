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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrFlowExporterMap(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.exporter_map_name", "TEST"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.destination_ipv4_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.destination_ipv6_address", "1::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.destination_vrf", "28"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.source", "10.1.1.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.dscp", "62"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.packet_length", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.transport_udp", "1033"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.dfbit_set", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_export_format", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_template_data_timeout", "1024"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_template_options_timeout", "3033"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_template_timeout", "2222"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_options_interface_table_timeout", "6048"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_options_sampler_table_timeout", "4096"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_options_class_table_timeout", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_exporter_map.test", "exporter_maps.0.version_options_vrf_table_timeout", "122"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrFlowExporterMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrFlowExporterMapConfig() string {
	config := `resource "iosxr_flow_exporter_map" "test" {` + "\n"
	config += `	exporter_maps = [{` + "\n"
	config += `		exporter_map_name = "TEST"` + "\n"
	config += `		destination_ipv4_address = "10.1.1.1"` + "\n"
	config += `		destination_ipv6_address = "1::1"` + "\n"
	config += `		destination_vrf = "28"` + "\n"
	config += `		source = "10.1.1.4"` + "\n"
	config += `		dscp = 62` + "\n"
	config += `		packet_length = 512` + "\n"
	config += `		transport_udp = 1033` + "\n"
	config += `		dfbit_set = true` + "\n"
	config += `		version_export_format = "true"` + "\n"
	config += `		version_template_data_timeout = 1024` + "\n"
	config += `		version_template_options_timeout = 3033` + "\n"
	config += `		version_template_timeout = 2222` + "\n"
	config += `		version_options_interface_table_timeout = 6048` + "\n"
	config += `		version_options_sampler_table_timeout = 4096` + "\n"
	config += `		version_options_class_table_timeout = 255` + "\n"
	config += `		version_options_vrf_table_timeout = 122` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_flow_exporter_map" "test" {
			depends_on = [iosxr_flow_exporter_map.test]
		}
	`
	return config
}

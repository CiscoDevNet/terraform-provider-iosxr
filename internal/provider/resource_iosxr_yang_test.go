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

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrYang(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrYangConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_yang.test", "id", "Cisco-IOS-XR-um-hostname-cfg:/hostname"),
				),
			},
			{
				Config: testAccIosxrYangConfig_hostname("tf-router-1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_yang.test", "id", "Cisco-IOS-XR-um-hostname-cfg:/hostname"),
					resource.TestCheckResourceAttr("iosxr_yang.test", "attributes.system-network-name", "tf-router-1"),
				),
			},
			{
				ResourceName:  "iosxr_yang.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-hostname-cfg:/hostname",
			},
			{
				Config: testAccIosxrYangConfig_hostname("router-1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_yang.test", "attributes.system-network-name", "router-1"),
				),
			},
			{
				Config: testAccIosxrYangConfig_list(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_yang.test", "id", "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]"),
					resource.TestCheckResourceAttr("iosxr_yang.test", "lists.0.items.0.ipv4-address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_yang.test", "lists.0.items.0.ipv4-address-index", "1"),
					resource.TestCheckResourceAttr("iosxr_yang.test", "lists.0.items.0.stitching", "enable"),
				),
			},
			{
				Config: testAccIosxrYangConfig_leafList(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_yang.test", "id", "Cisco-IOS-XR-um-domain-cfg:/domain/ipv4/hosts/host[host-name=abc.cisco.com]"),
					resource.TestCheckResourceAttr("iosxr_yang.test", "attributes.host-name", "abc.cisco.com"),
					resource.TestCheckResourceAttr("iosxr_yang.test", "lists.0.values.0", "1.2.3.4"),
				),
			},
			{
				Config: testAccIosxrYangConfig_yangEmpty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_yang.test", "id", "Cisco-IOS-XR-um-logging-cfg:/logging"),
					resource.TestCheckResourceAttr("iosxr_yang.test", "attributes.suppress/duplicates", "<EMPTY>"),
				),
			},
			{
				Config: testAccIosxrYangConfig_yangEmptyNull(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_yang.test", "attributes.system-network-name", "router-null-test"),
				),
			},
		},
	})
}

func testAccIosxrYangConfig_empty() string {
	return `
	resource "iosxr_yang" "test" {
		path = "Cisco-IOS-XR-um-hostname-cfg:/hostname"
	}
	`
}

func testAccIosxrYangConfig_hostname(name string) string {
	return fmt.Sprintf(`
	resource "iosxr_yang" "test" {
		path = "Cisco-IOS-XR-um-hostname-cfg:/hostname"
		attributes = {
			"system-network-name" = "%s"
		}
	}
	`, name)
}

func testAccIosxrYangConfig_list() string {
	return `
	resource "iosxr_yang" "test" {
		path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]"
		attributes = {
			"vrf-name" = "VRF1"
			"description" = "Test VRF for list configuration"
		}
		lists = [
			{
				name = "address-family/ipv4/unicast/Cisco-IOS-XR-um-router-bgp-cfg:import/route-target/ipv4-address-route-targets/ipv4-address-route-target"
				key = "ipv4-address,ipv4-address-index"
				items = [
					{
						ipv4-address       = "1.1.1.1"
						ipv4-address-index = "1"
						stitching          = "enable"
					}
				]
			}
		]
	}
	`
}

func testAccIosxrYangConfig_leafList() string {
	return `
	resource "iosxr_yang" "test" {
		path = "Cisco-IOS-XR-um-domain-cfg:/domain/ipv4/hosts/host[host-name=abc.cisco.com]"
		attributes = {
			"host-name" = "abc.cisco.com"
		}
		lists = [
			{
				name = "ip-address"
				values = ["1.2.3.4"]
			}
		]
	}
	`
}

func testAccIosxrYangConfig_yangEmpty() string {
	return `
	resource "iosxr_yang" "test" {
		path = "Cisco-IOS-XR-um-logging-cfg:/logging"
		attributes = {
			"hostnameprefix" = "TEST-EMPTY"
			"suppress/duplicates" = "<EMPTY>"
		}
	}
	`
}

func testAccIosxrYangConfig_yangEmptyNull() string {
	return `
	resource "iosxr_yang" "test" {
		path = "Cisco-IOS-XR-um-hostname-cfg:/hostname"
		attributes = {
			"system-network-name" = "router-null-test"
		}
	}
	`
}

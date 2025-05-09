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

func TestAccIosxrRouterBGPVRFNeighborAddressFamily(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "af_name", "ipv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "route_policy_in", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "route_policy_out", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "default_originate_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "next_hop_self", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "next_hop_self_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "soft_reconfiguration_inbound_always", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "send_community_ebgp_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "remove_private_as", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "remove_private_as_entire_aspath", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_neighbor_address_family.test", "remove_private_as_inbound_inheritance_disable", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterBGPVRFNeighborAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPVRFNeighborAddressFamilyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterBGPVRFNeighborAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPVRFNeighborAddressFamilyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_bgp_vrf_neighbor_address_family.test",
		ImportState:   true,
		ImportStateId: "65001,VRF1,10.1.1.2,ipv4-unicast",
		Check:         resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

const testAccIosxrRouterBGPVRFNeighborAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]/Cisco-IOS-XR-um-router-bgp-cfg:rd/Cisco-IOS-XR-um-router-bgp-cfg:two-byte-as"
	attributes = {
		"as-number" = "1"
		"index" = "1"
	}
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
	lists = [
		{
			name = "address-families/address-family"
			key = "af-name"
			items = [
				{
					"af-name" = "vpnv4-unicast"
				},
			]
		},
	]
}

resource "iosxr_gnmi" "PreReq2" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/vrfs/vrf[vrf-name=VRF1]"
	delete = false
	attributes = {
		"vrf-name" = "VRF1"
	}
	lists = [
		{
			name = "address-families/address-family"
			key = "af-name"
			items = [
				{
					"af-name" = "ipv4-unicast"
				},
			]
		},
		{
			name = "neighbors/neighbor"
			key = "neighbor-address"
			items = [
				{
					"neighbor-address" = "10.1.1.2"
					"remote-as" = "65002"
				},
			]
		},
	]
	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]
}

resource "iosxr_gnmi" "PreReq3" {
	path = "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
	attributes = {
		"route-policy-name" = "ROUTE_POLICY_1"
		"rpl-route-policy" = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
	}
}

`

func testAccIosxrRouterBGPVRFNeighborAddressFamilyConfig_minimum() string {
	config := `resource "iosxr_router_bgp_vrf_neighbor_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	neighbor_address = "10.1.1.2"` + "\n"
	config += `	af_name = "ipv4-unicast"` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, iosxr_gnmi.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterBGPVRFNeighborAddressFamilyConfig_all() string {
	config := `resource "iosxr_router_bgp_vrf_neighbor_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	neighbor_address = "10.1.1.2"` + "\n"
	config += `	af_name = "ipv4-unicast"` + "\n"
	config += `	route_policy_in = "ROUTE_POLICY_1"` + "\n"
	config += `	route_policy_out = "ROUTE_POLICY_1"` + "\n"
	config += `	default_originate_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	next_hop_self = true` + "\n"
	config += `	next_hop_self_inheritance_disable = true` + "\n"
	config += `	soft_reconfiguration_inbound_always = true` + "\n"
	config += `	send_community_ebgp_inheritance_disable = true` + "\n"
	config += `	remove_private_as = true` + "\n"
	config += `	remove_private_as_entire_aspath = true` + "\n"
	config += `	remove_private_as_inbound_inheritance_disable = true` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, iosxr_gnmi.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}

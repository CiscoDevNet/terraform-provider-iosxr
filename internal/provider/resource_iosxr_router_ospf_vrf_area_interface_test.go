// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrRouterOSPFVRFAreaInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterOSPFVRFAreaInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "network_broadcast", "false"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "network_non_broadcast", "false"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "network_point_to_point", "true"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "network_point_to_multipoint", "false"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "cost", "20"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "priority", "100"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "passive_enable", "false"),
					resource.TestCheckResourceAttr("iosxr_router_ospf_vrf_area_interface.test", "passive_disable", "true"),
				),
			},
			{
				ResourceName:  "iosxr_router_ospf_vrf_area_interface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=OSPF1]/vrfs/vrf[vrf-name=VRF1]/areas/area[area-id=0]/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]",
			},
		},
	})
}

func testAccIosxrRouterOSPFVRFAreaInterfaceConfig_minimum() string {
	return `
	resource "iosxr_router_ospf_vrf_area_interface" "test" {
		process_name = "OSPF1"
		vrf_name = "VRF1"
		area_id = "0"
		interface_name = "GigabitEthernet0/0/0/1"
	}
	`
}

func testAccIosxrRouterOSPFVRFAreaInterfaceConfig_all() string {
	return `
	resource "iosxr_router_ospf_vrf_area_interface" "test" {
		process_name = "OSPF1"
		vrf_name = "VRF1"
		area_id = "0"
		interface_name = "GigabitEthernet0/0/0/1"
		network_broadcast = false
		network_non_broadcast = false
		network_point_to_point = true
		network_point_to_multipoint = false
		cost = 20
		priority = 100
		passive_enable = false
		passive_disable = true
	}
	`
}

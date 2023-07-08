// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterOSPFAreaInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_non_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_point_to_point", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_point_to_multipoint", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "cost", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "passive_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "passive_disable", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterOSPFAreaInterfaceConfig_minimum(),
			},
			{
				Config: testAccIosxrRouterOSPFAreaInterfaceConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_router_ospf_area_interface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=OSPF1]/areas/area[area-id=0]/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]",
			},
		},
	})
}

func testAccIosxrRouterOSPFAreaInterfaceConfig_minimum() string {
	config := `resource "iosxr_router_ospf_area_interface" "test" {` + "\n"
	config += `	process_name = "OSPF1"` + "\n"
	config += `	area_id = "0"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterOSPFAreaInterfaceConfig_all() string {
	config := `resource "iosxr_router_ospf_area_interface" "test" {` + "\n"
	config += `	process_name = "OSPF1"` + "\n"
	config += `	area_id = "0"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	network_broadcast = false` + "\n"
	config += `	network_non_broadcast = false` + "\n"
	config += `	network_point_to_point = true` + "\n"
	config += `	network_point_to_multipoint = false` + "\n"
	config += `	cost = 20` + "\n"
	config += `	priority = 100` + "\n"
	config += `	passive_enable = false` + "\n"
	config += `	passive_disable = true` + "\n"
	config += `}` + "\n"
	return config
}

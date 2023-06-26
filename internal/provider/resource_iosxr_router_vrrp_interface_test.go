// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrRouterVRRPInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterVRRPInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "mac_refresh", "14"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "delay_minimum", "1234"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "delay_reload", "4321"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "bfd_minimum_interval", "255"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface.test", "bfd_multiplier", "33"),
				),
			},
			{
				ResourceName:  "iosxr_router_vrrp_interface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]",
			},
		},
	})
}

func testAccIosxrRouterVRRPInterfaceConfig_minimum() string {
	return `
	resource "iosxr_router_vrrp_interface" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
	}
	`
}

func testAccIosxrRouterVRRPInterfaceConfig_all() string {
	return `
	resource "iosxr_router_vrrp_interface" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
		mac_refresh = 14
		delay_minimum = 1234
		delay_reload = 4321
		bfd_minimum_interval = 255
		bfd_multiplier = 33
	}
	`
}

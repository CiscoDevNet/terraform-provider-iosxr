// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterVRRPInterfaceAddressFamilyIPv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterVRRPInterfaceAddressFamilyIPv4Config_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "vrrp_id", "123"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "version", "2"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "priority", "250"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "name", "TEST"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "text_authentication", "7"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "timer_advertisement_time_in_seconds", "123"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "timer_force", "false"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "preempt_disable", "false"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "preempt_delay", "255"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "accept_mode_disable", "false"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "track_interfaces.0.interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "track_interfaces.0.priority_decrement", "12"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "track_objects.0.object_name", "OBJECT"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "track_objects.0.priority_decrement", "22"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv4.test", "bfd_fast_detect_peer_ipv4", "33.33.33.3"),
				),
			},
			{
				ResourceName:  "iosxr_router_vrrp_interface_address_family_ipv4.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]/address-family/ipv4/vrrps/vrrp[vrrp-id=%!d(string=123)][version=%!d(string=2)]",
			},
		},
	})
}

func testAccIosxrRouterVRRPInterfaceAddressFamilyIPv4Config_minimum() string {
	return `
	resource "iosxr_router_vrrp_interface_address_family_ipv4" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
		vrrp_id = 123
		version = 2
	}
	`
}

func testAccIosxrRouterVRRPInterfaceAddressFamilyIPv4Config_all() string {
	return `
	resource "iosxr_router_vrrp_interface_address_family_ipv4" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
		vrrp_id = 123
		version = 2
		address = "1.1.1.1"
		priority = 250
		name = "TEST"
		text_authentication = "7"
		timer_advertisement_time_in_seconds = 123
		timer_force = false
		preempt_disable = false
		preempt_delay = 255
		accept_mode_disable = false
		track_interfaces = [{
			interface_name = "GigabitEthernet0/0/0/1"
			priority_decrement = 12
		}]
		track_objects = [{
			object_name = "OBJECT"
			priority_decrement = 22
		}]
		bfd_fast_detect_peer_ipv4 = "33.33.33.3"
	}
	`
}

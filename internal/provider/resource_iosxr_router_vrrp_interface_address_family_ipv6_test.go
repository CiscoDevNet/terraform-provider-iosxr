// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterVRRPInterfaceAddressFamilyIPv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterVRRPInterfaceAddressFamilyIPv6Config_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "vrrp_id", "123"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "address_linklocal_autoconfig", "true"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "priority", "250"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "name", "TEST"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "timer_advertisement_seconds", "10"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "timer_force", "true"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "preempt_disable", "false"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "preempt_delay", "255"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "accept_mode_disable", "true"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_interfaces.0.interface_name", "GigabitEthernet0/0/0/5"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_interfaces.0.priority_decrement", "12"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_objects.0.object_name", "OBJECT"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "track_objects.0.priority_decrement", "22"),
					resource.TestCheckResourceAttr("iosxr_router_vrrp_interface_address_family_ipv6.test", "bfd_fast_detect_peer_ipv6", "3::3"),
				),
			},
			{
				ResourceName:  "iosxr_router_vrrp_interface_address_family_ipv6.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]/address-family/ipv6/vrrps/vrrp[vrrp-id=%!d(string=123)]",
			},
		},
	})
}

func testAccIosxrRouterVRRPInterfaceAddressFamilyIPv6Config_minimum() string {
	return `
	resource "iosxr_router_vrrp_interface_address_family_ipv6" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
		vrrp_id = 123
	}
	`
}

func testAccIosxrRouterVRRPInterfaceAddressFamilyIPv6Config_all() string {
	return `
	resource "iosxr_router_vrrp_interface_address_family_ipv6" "test" {
		interface_name = "GigabitEthernet0/0/0/1"
		vrrp_id = 123
		address_linklocal_autoconfig = true
		priority = 250
		name = "TEST"
		timer_advertisement_seconds = 10
		timer_force = true
		preempt_disable = false
		preempt_delay = 255
		accept_mode_disable = true
		track_interfaces = [{
			interface_name = "GigabitEthernet0/0/0/5"
			priority_decrement = 12
		}]
		track_objects = [{
			object_name = "OBJECT"
			priority_decrement = 22
		}]
		bfd_fast_detect_peer_ipv6 = "3::3"
	}
	`
}

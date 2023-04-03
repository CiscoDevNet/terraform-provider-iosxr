// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrEVPNEVI(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrEVPNEVIPrerequisitesConfig + testAccIosxrEVPNEVIConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "vpn_id", "1234"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "description", "My Description"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "load_balancing", "true"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "load_balancing_flow_label_static", "true"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_rd_two_byte_as_number", "1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_rd_two_byte_as_assigned_number", "1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_route_target_import_two_byte_as_format.0.as_number", "1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_route_target_import_two_byte_as_format.0.assigned_number", "1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_route_target_export_ipv4_address_format.0.ipv4_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_route_target_export_ipv4_address_format.0.assigned_number", "1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_route_policy_import", "ROUTE_POLICY_1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "bgp_route_policy_export", "ROUTE_POLICY_1"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "advertise_mac", "true"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "unknown_unicast_suppression", "true"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "control_word_disable", "true"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "etree", "true"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "etree_leaf", "false"),
					resource.TestCheckResourceAttr("iosxr_evpn_evi.test", "etree_rt_leaf", "true"),
				),
			},
			{
				ResourceName:  "iosxr_evpn_evi.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/evpn/evis/evi[vpn-id=1234]",
			},
		},
	})
}

const testAccIosxrEVPNEVIPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
  path = "Cisco-IOS-XR-um-route-policy-cfg:routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
  attributes = {
      route-policy-name = "ROUTE_POLICY_1"
      rpl-route-policy = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
  }
}

`

func testAccIosxrEVPNEVIConfig_minimum() string {
	return `
	resource "iosxr_evpn_evi" "test" {
		vpn_id = 1234
  		depends_on = [iosxr_gnmi.PreReq0, ]
	}
	`
}

func testAccIosxrEVPNEVIConfig_all() string {
	return `
	resource "iosxr_evpn_evi" "test" {
		vpn_id = 1234
		description = "My Description"
		load_balancing = true
		load_balancing_flow_label_static = true
		bgp_rd_two_byte_as_number = 1
		bgp_rd_two_byte_as_assigned_number = 1
		bgp_route_target_import_two_byte_as_format = [{
			as_number = 1
			assigned_number = 1
		}]
		bgp_route_target_export_ipv4_address_format = [{
			ipv4_address = "1.1.1.1"
			assigned_number = 1
		}]
		bgp_route_policy_import = "ROUTE_POLICY_1"
		bgp_route_policy_export = "ROUTE_POLICY_1"
		advertise_mac = true
		unknown_unicast_suppression = true
		control_word_disable = true
		etree = true
		etree_leaf = false
		etree_rt_leaf = true
  		depends_on = [iosxr_gnmi.PreReq0, ]
	}
	`
}

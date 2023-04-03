// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrVRFPrerequisitesConfig + testAccIosxrVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_vrf.test", "vrf_name", "VRF1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "description", "My VRF Description"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "vpn_id", "1000:1000"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_policy", "ROUTE_POLICY_1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_policy", "ROUTE_POLICY_1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_multicast", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_flowspec", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_policy", "ROUTE_POLICY_1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_policy", "ROUTE_POLICY_1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_multicast", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_flowspec", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "rd_two_byte_as_as_number", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "rd_two_byte_as_index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_two_byte_as_format.0.as_number", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_two_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_two_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_four_byte_as_format.0.as_number", "100000"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_four_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_four_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_ip_address_format.0.ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_ip_address_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_ip_address_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_two_byte_as_format.0.as_number", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_two_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_two_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_four_byte_as_format.0.as_number", "100000"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_four_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_four_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_ip_address_format.0.ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_ip_address_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_ip_address_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_two_byte_as_format.0.as_number", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_two_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_two_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_four_byte_as_format.0.as_number", "100000"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_four_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_four_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_ip_address_format.0.ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_ip_address_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_ip_address_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_two_byte_as_format.0.as_number", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_two_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_two_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_four_byte_as_format.0.as_number", "100000"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_four_byte_as_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_four_byte_as_format.0.stitching", "true"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_ip_address_format.0.ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_ip_address_format.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_ip_address_format.0.stitching", "true"),
				),
			},
			{
				ResourceName:  "iosxr_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]",
			},
		},
	})
}

const testAccIosxrVRFPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
  path = "Cisco-IOS-XR-um-route-policy-cfg:routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
  attributes = {
      route-policy-name = "ROUTE_POLICY_1"
      rpl-route-policy = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
  }
}

`

func testAccIosxrVRFConfig_minimum() string {
	return `
	resource "iosxr_vrf" "test" {
		vrf_name = "VRF1"
  		depends_on = [iosxr_gnmi.PreReq0, ]
	}
	`
}

func testAccIosxrVRFConfig_all() string {
	return `
	resource "iosxr_vrf" "test" {
		vrf_name = "VRF1"
		description = "My VRF Description"
		vpn_id = "1000:1000"
		address_family_ipv4_unicast = true
		address_family_ipv4_unicast_import_route_policy = "ROUTE_POLICY_1"
		address_family_ipv4_unicast_export_route_policy = "ROUTE_POLICY_1"
		address_family_ipv4_multicast = true
		address_family_ipv4_flowspec = true
		address_family_ipv6_unicast = true
		address_family_ipv6_unicast_import_route_policy = "ROUTE_POLICY_1"
		address_family_ipv6_unicast_export_route_policy = "ROUTE_POLICY_1"
		address_family_ipv6_multicast = true
		address_family_ipv6_flowspec = true
		rd_two_byte_as_as_number = "1"
		rd_two_byte_as_index = 1
		address_family_ipv4_unicast_import_route_target_two_byte_as_format = [{
			as_number = 1
			index = 1
			stitching = true
		}]
		address_family_ipv4_unicast_import_route_target_four_byte_as_format = [{
			as_number = 100000
			index = 1
			stitching = true
		}]
		address_family_ipv4_unicast_import_route_target_ip_address_format = [{
			ip_address = "1.1.1.1"
			index = 1
			stitching = true
		}]
		address_family_ipv4_unicast_export_route_target_two_byte_as_format = [{
			as_number = 1
			index = 1
			stitching = true
		}]
		address_family_ipv4_unicast_export_route_target_four_byte_as_format = [{
			as_number = 100000
			index = 1
			stitching = true
		}]
		address_family_ipv4_unicast_export_route_target_ip_address_format = [{
			ip_address = "1.1.1.1"
			index = 1
			stitching = true
		}]
		address_family_ipv6_unicast_import_route_target_two_byte_as_format = [{
			as_number = 1
			index = 1
			stitching = true
		}]
		address_family_ipv6_unicast_import_route_target_four_byte_as_format = [{
			as_number = 100000
			index = 1
			stitching = true
		}]
		address_family_ipv6_unicast_import_route_target_ip_address_format = [{
			ip_address = "1.1.1.1"
			index = 1
			stitching = true
		}]
		address_family_ipv6_unicast_export_route_target_two_byte_as_format = [{
			as_number = 1
			index = 1
			stitching = true
		}]
		address_family_ipv6_unicast_export_route_target_four_byte_as_format = [{
			as_number = 100000
			index = 1
			stitching = true
		}]
		address_family_ipv6_unicast_export_route_target_ip_address_format = [{
			ip_address = "1.1.1.1"
			index = 1
			stitching = true
		}]
  		depends_on = [iosxr_gnmi.PreReq0, ]
	}
	`
}

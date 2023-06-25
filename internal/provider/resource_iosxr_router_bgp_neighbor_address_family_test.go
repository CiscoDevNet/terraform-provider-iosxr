// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterBGPNeighborAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPNeighborAddressFamilyConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "af_name", "vpnv4-unicast"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "import_stitching_rt_re_originate_stitching_rt", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "route_reflector_client_inheritance_disable", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "advertise_vpnv4_unicast_enable_re_originated_stitching_rt", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "next_hop_self_inheritance_disable", "true"),
					resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "encapsulation_type_srv6", "true"),
				),
			},
			{
				ResourceName:  "iosxr_router_bgp_neighbor_address_family.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/neighbors/neighbor[neighbor-address=10.1.1.2]/address-families/address-family[af-name=vpnv4-unicast]",
			},
		},
	})
}

const testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/address-families/address-family[af-name=vpnv4-unicast]"
	delete = false
	attributes = {
		"af-name" = "vpnv4-unicast"
	}
	depends_on = [iosxr_gnmi.PreReq0, ]
}

resource "iosxr_gnmi" "PreReq2" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/neighbors/neighbor[neighbor-address=10.1.1.2]"
	delete = false
	attributes = {
		"neighbor-address" = "10.1.1.2"
		"remote-as" = "65002"
	}
	depends_on = [iosxr_gnmi.PreReq0, ]
}

`

func testAccIosxrRouterBGPNeighborAddressFamilyConfig_minimum() string {
	return `
	resource "iosxr_router_bgp_neighbor_address_family" "test" {
		as_number = "65001"
		neighbor_address = "10.1.1.2"
		af_name = "vpnv4-unicast"
		depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, ]
	}
	`
}

func testAccIosxrRouterBGPNeighborAddressFamilyConfig_all() string {
	return `
	resource "iosxr_router_bgp_neighbor_address_family" "test" {
		as_number = "65001"
		neighbor_address = "10.1.1.2"
		af_name = "vpnv4-unicast"
		import_stitching_rt_re_originate_stitching_rt = true
		route_reflector_client_inheritance_disable = true
		advertise_vpnv4_unicast_enable_re_originated_stitching_rt = true
		next_hop_self_inheritance_disable = true
		encapsulation_type_srv6 = true
  		depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, ]
	}
	`
}

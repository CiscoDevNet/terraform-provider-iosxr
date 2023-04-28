// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterISISInterfaceAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterISISInterfaceAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface_address_family.test", "fast_reroute_per_prefix_levels.0.level_id", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface_address_family.test", "fast_reroute_per_prefix_levels.0.ti_lfa", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface_address_family.test", "tag", "100"),
					resource.TestCheckResourceAttr("data.iosxr_router_isis_interface_address_family.test", "advertise_prefix_route_policy", "ROUTE_POLICY_1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterISISInterfaceAddressFamilyConfig = `

resource "iosxr_router_isis_interface_address_family" "test" {
	process_id = "P1"
	interface_name = "GigabitEthernet0/0/0/1"
	af_name = "ipv4"
	saf_name = "unicast"
	fast_reroute_per_prefix_levels = [{
		level_id = 1
		ti_lfa = true
	}]
	tag = 100
	advertise_prefix_route_policy = "ROUTE_POLICY_1"
}

data "iosxr_router_isis_interface_address_family" "test" {
	process_id = "P1"
	interface_name = "GigabitEthernet0/0/0/1"
	af_name = "ipv4"
	saf_name = "unicast"
	depends_on = [iosxr_router_isis_interface_address_family.test]
}
`

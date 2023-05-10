// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrSegmentRoutingISIS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSegmentRoutingISISConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_segment_routing_isis.test", "metric_style_wide", "true"),
					resource.TestCheckResourceAttr("data.iosxr_segment_routing_isis.test", "microloop_avoidance_segment_routing", "true"),
					resource.TestCheckResourceAttr("data.iosxr_segment_routing_isis.test", "router_id_interface_name", "Loopback0"),
					resource.TestCheckResourceAttr("data.iosxr_segment_routing_isis.test", "locators.0.locator_name", "AlgoLocator"),
					resource.TestCheckResourceAttr("data.iosxr_segment_routing_isis.test", "locators.0.level", "1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrSegmentRoutingISISConfig = `

resource "iosxr_segment_routing_isis" "test" {
	process_id = "P1"
	af_name = "ipv6"
	saf_name = "unicast"
	metric_style_wide = true
	microloop_avoidance_segment_routing = true
	router_id_interface_name = "Loopback0"
	locators = [{
		locator_name = "AlgoLocator"
		level = 1
	}]
}

data "iosxr_segment_routing_isis" "test" {
	process_id = "P1"
	af_name = "ipv6"
	saf_name = "unicast"
	depends_on = [iosxr_segment_routing_isis.test]
}
`

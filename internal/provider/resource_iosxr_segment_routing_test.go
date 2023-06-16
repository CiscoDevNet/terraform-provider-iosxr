// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSegmentRouting(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSegmentRoutingConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_segment_routing.test", "global_block_lower_bound", "16000"),
					resource.TestCheckResourceAttr("iosxr_segment_routing.test", "global_block_upper_bound", "29999"),
					resource.TestCheckResourceAttr("iosxr_segment_routing.test", "local_block_lower_bound", "15000"),
					resource.TestCheckResourceAttr("iosxr_segment_routing.test", "local_block_upper_bound", "15999"),
				),
			},
			{
				ResourceName:  "iosxr_segment_routing.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-segment-routing-ms-cfg:/sr",
			},
		},
	})
}

func testAccIosxrSegmentRoutingConfig_minimum() string {
	return `
	resource "iosxr_segment_routing" "test" {
		global_block_lower_bound = 16000
		global_block_upper_bound = 29999
		local_block_lower_bound = 15000
		local_block_upper_bound = 15999
	}
	`
}

func testAccIosxrSegmentRoutingConfig_all() string {
	return `
	resource "iosxr_segment_routing" "test" {
		global_block_lower_bound = 16000
		global_block_upper_bound = 29999
		local_block_lower_bound = 15000
		local_block_upper_bound = 15999
	}
	`
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrL2VPN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrL2VPNConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_l2vpn.test", "description", "My L2VPN Description"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn.test", "router_id", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn.test", "xconnect_groups.0.group_name", "P2P"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrL2VPNConfig = `

resource "iosxr_l2vpn" "test" {
	description = "My L2VPN Description"
	router_id = "1.2.3.4"
	xconnect_groups = [{
		group_name = "P2P"
	}]
}

data "iosxr_l2vpn" "test" {
	depends_on = [iosxr_l2vpn.test]
}
`

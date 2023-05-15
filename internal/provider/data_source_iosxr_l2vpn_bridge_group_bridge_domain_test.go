// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrL2VPNBridgeGroupBridgeDomain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrL2VPNBridgeGroupBridgeDomainConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_bridge_group_bridge_domain.test", "evis.0.vpn_id", "1234"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_bridge_group_bridge_domain.test", "vnis.0.vni_id", "1234"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_bridge_group_bridge_domain.test", "segment_routing_srv6_evis_evi_vpn_id", "32"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrL2VPNBridgeGroupBridgeDomainConfig = `

resource "iosxr_l2vpn_bridge_group_bridge_domain" "test" {
	bridge_group_name = "BG123"
	bridge_domain_name = "BD123"
	evis = [{
		vpn_id = 1234
	}]
	vnis = [{
		vni_id = 1234
	}]
	segment_routing_srv6_evis_evi_vpn_id = 32
}

data "iosxr_l2vpn_bridge_group_bridge_domain" "test" {
	bridge_group_name = "BG123"
	bridge_domain_name = "BD123"
	depends_on = [iosxr_l2vpn_bridge_group_bridge_domain.test]
}
`

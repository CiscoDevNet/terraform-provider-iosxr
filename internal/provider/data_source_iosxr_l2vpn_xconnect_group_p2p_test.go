// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrL2VPNXconnectGroupP2P(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrL2VPNXconnectGroupP2PConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_xconnect_group_p2p.test", "description", "My P2P Description"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_xconnect_group_p2p.test", "interfaces.0.interface_name", "GigabitEthernet0/0/0/2"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_xconnect_group_p2p.test", "ipv4_neighbors.0.address", "2.3.4.5"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_xconnect_group_p2p.test", "ipv4_neighbors.0.pw_id", "1"),
					resource.TestCheckResourceAttr("data.iosxr_l2vpn_xconnect_group_p2p.test", "ipv4_neighbors.0.pw_class", "PW_CLASS_1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrL2VPNXconnectGroupP2PConfig = `

resource "iosxr_l2vpn_xconnect_group_p2p" "test" {
	group_name = "P2P"
	p2p_xconnect_name = "XC"
	description = "My P2P Description"
	interfaces = [{
		interface_name = "GigabitEthernet0/0/0/2"
	}]
	ipv4_neighbors = [{
		address = "2.3.4.5"
		pw_id = 1
		pw_class = "PW_CLASS_1"
	}]
}

data "iosxr_l2vpn_xconnect_group_p2p" "test" {
	group_name = "P2P"
	p2p_xconnect_name = "XC"
	depends_on = [iosxr_l2vpn_xconnect_group_p2p.test]
}
`

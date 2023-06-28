// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrLLDP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrLLDPConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "holdtime", "50"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "timer", "6"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "reinit", "3"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "subinterfaces_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "priorityaddr_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "extended_show_width_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "tlv_select_management_address_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "tlv_select_port_description_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "tlv_select_system_capabilities_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "tlv_select_system_description_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_lldp.test", "tlv_select_system_name_disable", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrLLDPConfig = `

resource "iosxr_lldp" "test" {
	delete_mode = "attributes"
	holdtime = 50
	timer = 6
	reinit = 3
	subinterfaces_enable = true
	priorityaddr_enable = true
	extended_show_width_enable = true
	tlv_select_management_address_disable = true
	tlv_select_port_description_disable = true
	tlv_select_system_capabilities_disable = true
	tlv_select_system_description_disable = true
	tlv_select_system_name_disable = true
}

data "iosxr_lldp" "test" {
	depends_on = [iosxr_lldp.test]
}
`

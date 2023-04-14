// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrSNMPMIB(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSNMPMIBConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_snmp_mib.test", "ifmib_ifalias_long", "true"),
					resource.TestCheckResourceAttr("data.iosxr_snmp_mib.test", "ifindex_persist", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrSNMPMIBConfig = `

resource "iosxr_snmp_mib" "test" {
	ifmib_ifalias_long = true
	ifindex_persist = true
}

data "iosxr_snmp_mib" "test" {
	depends_on = [iosxr_snmp_mib.test]
}
`

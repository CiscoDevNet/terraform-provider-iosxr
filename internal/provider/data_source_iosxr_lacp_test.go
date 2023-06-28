// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrLACP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrLACPConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_lacp.test", "mac", "00:11:00:11:00:11"),
					resource.TestCheckResourceAttr("data.iosxr_lacp.test", "priority", "1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrLACPConfig = `

resource "iosxr_lacp" "test" {
	delete_mode = "attributes"
	mac = "00:11:00:11:00:11"
	priority = 1
}

data "iosxr_lacp" "test" {
	depends_on = [iosxr_lacp.test]
}
`

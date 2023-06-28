// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrESISet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrESISetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_esi_set.test", "rpl", "esi-set POLICYSET\n  1234.1234.1234.1234.1234\nend-set\n"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrESISetConfig = `

resource "iosxr_esi_set" "test" {
	set_name = "POLICYSET"
	rpl = "esi-set POLICYSET\n  1234.1234.1234.1234.1234\nend-set\n"
}

data "iosxr_esi_set" "test" {
	set_name = "POLICYSET"
	depends_on = [iosxr_esi_set.test]
}
`

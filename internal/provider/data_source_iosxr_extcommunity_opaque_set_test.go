// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrExtcommunityOpaqueSet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrExtcommunityOpaqueSetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_extcommunity_opaque_set.test", "rpl", "extcommunity-set opaque BLUE\n  100\nend-set\n"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrExtcommunityOpaqueSetConfig = `

resource "iosxr_extcommunity_opaque_set" "test" {
	set_name = "BLUE"
	rpl = "extcommunity-set opaque BLUE\n  100\nend-set\n"
}

data "iosxr_extcommunity_opaque_set" "test" {
	set_name = "BLUE"
	depends_on = [iosxr_extcommunity_opaque_set.test]
}
`
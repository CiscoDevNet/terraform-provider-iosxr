package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrGnmi(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrGnmiConfigInterface,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_gnmi.test", "id", "openconfig-system:/system/config"),
					resource.TestCheckResourceAttr("data.iosxr_gnmi.test", "attributes.hostname", "TF-ROUTER-1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrGnmiConfigInterface = `
resource "iosxr_gnmi" "test" {
	path = "openconfig-system:/system/config"
	attributes = {
		hostname = "TF-ROUTER-1"
	}
}

data "iosxr_gnmi" "test" {
	path = "openconfig-system:/system/config"
	depends_on = [iosxr_gnmi.test]
}
`

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrGnmi(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrGnmiConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "id", "openconfig-system:/system/config"),
				),
			},
			{
				Config: testAccIosxrGnmiConfig_interface("TF-ROUTER-1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "id", "openconfig-system:/system/config"),
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "attributes.hostname", "TF-ROUTER-1"),
				),
			},
			{
				ResourceName:  "iosxr_gnmi.test",
				ImportState:   true,
				ImportStateId: "openconfig-system:/system/config",
			},
			{
				Config: testAccIosxrGnmiConfig_interface("TF-ROUTER-1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "attributes.hostname", "TF-ROUTER-1"),
				),
			},
		},
	})
}

func testAccIosxrGnmiConfig_empty() string {
	return `
	resource "iosxr_gnmi" "test" {
		path = "openconfig-system:/system/config"
	}
	`
}

func testAccIosxrGnmiConfig_interface(name string) string {
	return fmt.Sprintf(`
	resource "iosxr_gnmi" "test" {
		path = "openconfig-system:/system/config"
		attributes = {
			hostname = "%s"
		}
	}
	`, name)
}

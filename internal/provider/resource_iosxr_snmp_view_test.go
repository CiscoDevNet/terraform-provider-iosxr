// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSNMPView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSNMPViewConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_snmp_view.test", "mib_view_family_name", "iso"),
					resource.TestCheckResourceAttr("iosxr_snmp_view.test", "included", "true"),
				),
			},
			{
				ResourceName:  "iosxr_snmp_view.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-snmp-server-cfg:snmp-server/views/view[view-name-VIEW2]/mib-view-families/mib-view-family[mib-view-family-name=iso]",
			},
		},
	})
}

func testAccIosxrSNMPViewConfig_minimum() string {
	return `
	resource "iosxr_snmp_view" "test" {
		view_name = "VIEW2"
		mib_view_family_name = "iso"
	}
	`
}

func testAccIosxrSNMPViewConfig_all() string {
	return `
	resource "iosxr_snmp_view" "test" {
		view_name = "VIEW2"
		mib_view_family_name = "iso"
		included = true
	}
	`
}
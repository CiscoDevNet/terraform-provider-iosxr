// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrSNMPServerView(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_snmp_server_view.test", "view_name", "VIEW12"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_snmp_server_view.test", "mib_view_families.0.name", "iso"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_snmp_server_view.test", "mib_view_families.0.included", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSNMPServerViewConfig_minimum(),
			},
			{
				Config: testAccIosxrSNMPServerViewConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_snmp_server_view.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server/views/view[view-name=VIEW12]",
			},
		},
	})
}

func testAccIosxrSNMPServerViewConfig_minimum() string {
	config := `resource "iosxr_snmp_server_view" "test" {` + "\n"
	config += `	view_name = "VIEW12"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrSNMPServerViewConfig_all() string {
	config := `resource "iosxr_snmp_server_view" "test" {` + "\n"
	config += `	view_name = "VIEW12"` + "\n"
	config += `	mib_view_families = [{` + "\n"
	config += `		name = "iso"` + "\n"
	config += `		included = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

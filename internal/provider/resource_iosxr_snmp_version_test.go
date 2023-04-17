// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSNMPVersion(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSNMPVersionConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_snmp_version.test", "version_v3_security_level", "true"),
				),
			},
			{
				ResourceName:  "iosxr_snmp_version.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-snmp-server-cfg:snmp-server/vrfs/vrf[vrf-name-%!s(MISSING)]/hosts/host[address=%!s(MISSING)]/traps/unencrypted/unencrypted-string[community-string=%!s(MISSING)]",
			},
		},
	})
}

func testAccIosxrSNMPVersionConfig_minimum() string {
	return `
	resource "iosxr_snmp_version" "test" {
		version_v3_security_level = "true"
	}
	`
}

func testAccIosxrSNMPVersionConfig_all() string {
	return `
	resource "iosxr_snmp_version" "test" {
		version_v3_security_level = "true"
	}
	`
}
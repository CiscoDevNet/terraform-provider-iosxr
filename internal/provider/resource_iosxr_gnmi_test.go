package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
				Config: testAccIosxrGnmiConfig_hostname("TF-ROUTER-1"),
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
				Config: testAccIosxrGnmiConfig_hostname("ROUTER-1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "attributes.hostname", "ROUTER-1"),
				),
			},
			{
				Config: testAccIosxrGnmiConfig_list(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "id", "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]"),
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "lists.0.items.0.ip-address", "1.1.1.1"),
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "lists.0.items.0.index", "1"),
					resource.TestCheckResourceAttr("iosxr_gnmi.test", "lists.0.items.0.stitching", "true"),
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

func testAccIosxrGnmiConfig_hostname(name string) string {
	return fmt.Sprintf(`
	resource "iosxr_gnmi" "test" {
		path = "openconfig-system:/system/config"
		attributes = {
			"hostname" = "%s"
		}
	}
	`, name)
}

func testAccIosxrGnmiConfig_list() string {
	return `
	resource "iosxr_gnmi" "test" {
		path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]"
		attributes = {
			"vrf-name" = "VRF1"
		}
		lists = [
			{
				name = "address-family/ipv4/unicast/Cisco-IOS-XR-um-router-bgp-cfg:import/route-target/ip-addresse-rts/ip-address-rt"
				key = "ip-address,index,stitching"
				items = [
					{
						ip-address = "1.1.1.1"
						index      = "1"
						stitching  = "true"
					}
				]
			}
		]
	}
	`
}

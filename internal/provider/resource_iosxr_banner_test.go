// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrBanner(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_banner.test", "banner_type", "login"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_banner.test", "line", " Hello user !"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrBannerConfig_minimum(),
			},
			{
				Config: testAccIosxrBannerConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_banner.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-banner-cfg:/banners/banner[banner-type=login]",
			},
		},
	})
}

func testAccIosxrBannerConfig_minimum() string {
	config := `resource "iosxr_banner" "test" {` + "\n"
	config += `	banner_type = "login"` + "\n"
	config += `	line = " Hello user !"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrBannerConfig_all() string {
	config := `resource "iosxr_banner" "test" {` + "\n"
	config += `	banner_type = "login"` + "\n"
	config += `	line = " Hello user !"` + "\n"
	config += `}` + "\n"
	return config
}

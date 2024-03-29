// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrFPD(t *testing.T) {
	if os.Getenv("FPD") == "" {
		t.Skip("skipping test, set environment variable FPD")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_fpd.test", "auto_upgrade_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_fpd.test", "auto_upgrade_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_fpd.test", "auto_reload_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_fpd.test", "auto_reload_disable", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrFPDConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrFPDConfig() string {
	config := `resource "iosxr_fpd" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	auto_upgrade_enable = false` + "\n"
	config += `	auto_upgrade_disable = false` + "\n"
	config += `	auto_reload_enable = false` + "\n"
	config += `	auto_reload_disable = false` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_fpd" "test" {
			depends_on = [iosxr_fpd.test]
		}
	`
	return config
}

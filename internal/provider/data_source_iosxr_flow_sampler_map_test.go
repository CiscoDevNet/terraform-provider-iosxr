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

func TestAccDataSourceIosxrFlowSamplerMap(t *testing.T) {
	if os.Getenv("FLOW") == "" {
		t.Skip("skipping test, set environment variable FLOW")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_sampler_map.test", "random", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_flow_sampler_map.test", "out_of", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrFlowSamplerMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrFlowSamplerMapConfig() string {
	config := `resource "iosxr_flow_sampler_map" "test" {` + "\n"
	config += `	name = "sampler_map1"` + "\n"
	config += `	random = 1` + "\n"
	config += `	out_of = 1` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_flow_sampler_map" "test" {
			name = "sampler_map1"
			depends_on = [iosxr_flow_sampler_map.test]
		}
	`
	return config
}

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

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

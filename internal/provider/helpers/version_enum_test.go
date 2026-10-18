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

package helpers_test

import (
	"testing"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// bgpASFormatModel mirrors the generated BGPASFormat struct for test purposes.
type bgpASFormatModel struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	DeleteMode types.String `tfsdk:"delete_mode"`
	AsFormat   types.String `tfsdk:"as_format"`
}

// bgpASFormatConstraints mirrors the GetEnumConstraints() output for bgp_as_format
// after the 25.4 test override: 24.4={asdot,asplain}, 25.4={asdot,asdot+}.
var bgpASFormatConstraints = []helpers.FieldEnumConstraint{
	{
		FieldPath: "as_format",
		VersionEnums: map[string][]string{
			"24.4": {"asdot", "asplain"},
			"25.4": {"asdot", "asdot+"},
		},
	},
}

func TestValidateVersionEnums(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		asFormat  *string // nil means types.StringNull()
		wantError bool
	}{
		{name: "24.4 allows asdot", version: "24.4", asFormat: strPtr("asdot"), wantError: false},
		{name: "24.4 allows asplain", version: "24.4", asFormat: strPtr("asplain"), wantError: false},
		{name: "24.4 rejects asdot+ (not in 24.4 set)", version: "24.4", asFormat: strPtr("asdot+"), wantError: true},
		{name: "25.4 allows asdot", version: "25.4", asFormat: strPtr("asdot"), wantError: false},
		{name: "25.4 allows asdot+", version: "25.4", asFormat: strPtr("asdot+"), wantError: false},
		{name: "25.4 rejects asplain (removed in 25.4)", version: "25.4", asFormat: strPtr("asplain"), wantError: true},
		{name: "null field skipped", version: "25.4", asFormat: nil, wantError: false},
		{name: "empty version skipped", version: "", asFormat: strPtr("asplain"), wantError: false},
		// 24.3 is below both 24.4 and 25.4 thresholds — no restriction applies.
		{name: "version below all thresholds skipped", version: "24.3", asFormat: strPtr("asplain"), wantError: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			asFormat := types.StringNull()
			if tc.asFormat != nil {
				asFormat = types.StringValue(*tc.asFormat)
			}
			plan := bgpASFormatModel{AsFormat: asFormat}
			var diags diag.Diagnostics
			helpers.ValidateVersionEnums(tc.version, plan, bgpASFormatConstraints, &diags)
			if diags.HasError() != tc.wantError {
				t.Errorf("ValidateVersionEnums(%q, AsFormat=%v): HasError() = %v, want %v (diags: %v)",
					tc.version, tc.asFormat, diags.HasError(), tc.wantError, diags)
			}
		})
	}
}

func strPtr(s string) *string { return &s }

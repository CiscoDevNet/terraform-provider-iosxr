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

package helpers

import (
	"encoding/base64"
	"testing"
)

func TestNormalizeVersion(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   string
		wantOk bool
	}{
		{name: "three_dotted_decimal", input: "24.4.2", want: "24.4", wantOk: true},
		{name: "two_dotted_decimal", input: "25.4", want: "25.4", wantOk: true},
		{name: "three_dotted_two_digit_patch", input: "25.4.10", want: "25.4", wantOk: true},
		{name: "three_dotted_two_digit_patch_26_2", input: "26.2.20", want: "26.2", wantOk: true},
		{name: "three_dotted_two_digit_minor", input: "25.10.2", want: "25.10", wantOk: true},
		{name: "two_dotted_two_digit_minor", input: "25.10", want: "25.10", wantOk: true},
		{name: "compact_format_rejected", input: "2442", want: "", wantOk: false},
		{name: "empty_string", input: "", want: "", wantOk: false},
		{name: "malformed_word", input: "garbage", want: "", wantOk: false},
		{name: "malformed_dotted_alpha", input: "abc.def", want: "", wantOk: false},
		{name: "gnmi_version", input: "0.10.0", want: "0.10", wantOk: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := NormalizeVersion(tc.input)
			if got != tc.want || ok != tc.wantOk {
				t.Errorf("NormalizeVersion(%q) = %q, %v; want %q, %v",
					tc.input, got, ok, tc.want, tc.wantOk)
			}
		})
	}
}

func TestValidateSupportedVersion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "three_dotted_decimal", input: "24.4.2", want: true},
		{name: "two_dotted_decimal", input: "25.4", want: true},
		{name: "compact_format_rejected", input: "2442", want: false},
		{name: "empty_string", input: "", want: false},
		{name: "malformed_word", input: "garbage", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateSupportedVersion(tc.input)
			if got != tc.want {
				t.Errorf("ValidateSupportedVersion(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestExtractVersionFromDevice(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "c8000_25_4_2", input: `{"label": "25.4.2", "hardware-info": "8000", "chassis-pid": "8201-24H8FH"}`, want: "25.4.2"},
		{name: "c8000_24_4_1", input: `{"label": "24.4.1", "hardware-info": "8000", "chassis-pid": "8201-24H8FH"}`, want: "24.4.1"},
		{name: "ncs5500_24_4_2", input: `{"label": "24.4.2", "hardware-info": "cisco NCS-5500 () processor"}`, want: "24.4.2"},
		{name: "xrv9k_25_4_2", input: `{"label": "25.4.2", "hardware-info": "cisco IOS-XRv 9000 () processor"}`, want: "25.4.2"},
		{name: "no_label", input: `{"hardware-info": "8000"}`, want: ""},
		{name: "empty_string", input: "", want: ""},
		{name: "invalid_json", input: "not json", want: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := extractVersionFromDevice(tc.input)
			if got != tc.want {
				t.Errorf("extractVersionFromDevice(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestDecodeBase64Payload(t *testing.T) {
	plaintext := `{"label": "24.4.2"}`
	encoded := base64.StdEncoding.EncodeToString([]byte(plaintext))

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "valid_base64", input: encoded, want: plaintext},
		{name: "invalid_base64", input: "not-valid-base64!!!", want: ""},
		{name: "empty_string", input: "", want: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := decodeBase64Payload(tc.input)
			if got != tc.want {
				t.Errorf("decodeBase64Payload(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

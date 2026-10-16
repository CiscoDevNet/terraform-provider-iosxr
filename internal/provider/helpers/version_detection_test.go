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
	"context"
	"encoding/base64"
	"testing"

	"github.com/netascode/go-gnmi"
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

func TestExtractHardwareFromDevice(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "c8000", input: `{"label": "25.4.2", "hardware-info": "8000", "chassis-pid": "8201-24H8FH"}`, want: "8000"},
		{name: "ncs5500", input: `{"label": "24.4.2", "hardware-info": "cisco NCS-5500 () processor"}`, want: "cisco NCS-5500 () processor"},
		{name: "xrv9k", input: `{"label": "25.4.2", "hardware-info": "cisco IOS-XRv 9000 () processor"}`, want: "cisco IOS-XRv 9000 () processor"},
		{name: "whitespace_trimmed", input: `{"label": "25.4.2", "hardware-info": "  8000 \n"}`, want: "8000"},
		{name: "no_hardware_info", input: `{"label": "24.4.2"}`, want: ""},
		{name: "empty_string", input: "", want: ""},
		{name: "invalid_json", input: "not json", want: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := extractHardwareFromDevice(tc.input)
			if got != tc.want {
				t.Errorf("extractHardwareFromDevice(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

// installVersionResponse builds the gNMI response shape the provider receives for install/version
func installVersionResponse(payload string) map[string]interface{} {
	return map[string]interface{}{
		"Notifications": []interface{}{
			map[string]interface{}{
				"update": []interface{}{
					map[string]interface{}{
						"val": map[string]interface{}{
							"Value": map[string]interface{}{
								"JsonIetfVal": base64.StdEncoding.EncodeToString([]byte(payload)),
							},
						},
					},
				},
			},
		},
	}
}

func TestExtractInstallInfoFromResponse(t *testing.T) {
	tests := []struct {
		name         string
		response     interface{}
		wantVersion  string
		wantHardware string
		wantErr      bool
	}{
		{name: "label_and_hardware", response: installVersionResponse(`{"label": "25.4.2", "hardware-info": "8000"}`), wantVersion: "25.4.2", wantHardware: "8000"},
		{name: "label_only", response: installVersionResponse(`{"label": "24.4.2"}`), wantVersion: "24.4.2", wantHardware: ""},
		{name: "hardware_only", response: installVersionResponse(`{"hardware-info": "8000"}`), wantErr: true},
		{name: "no_notifications", response: map[string]interface{}{}, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			version, hardware, err := extractInstallInfoFromResponse(context.Background(), tc.response)
			if (err != nil) != tc.wantErr {
				t.Fatalf("extractInstallInfoFromResponse() error = %v; wantErr %v", err, tc.wantErr)
			}
			if version != tc.wantVersion || hardware != tc.wantHardware {
				t.Errorf("extractInstallInfoFromResponse() = (%q, %q); want (%q, %q)", version, hardware, tc.wantVersion, tc.wantHardware)
			}
		})
	}
}

func TestDetectIosxrHardwareCache(t *testing.T) {
	ClearVersionCache("")
	t.Cleanup(func() { ClearVersionCache("") })

	// A cache hit does not use the client, so a zero value client is enough
	client := &gnmi.Client{}

	hardwareCache.Store("c8000", "8000")
	got, err := DetectIosxrHardware(context.Background(), client, "c8000")
	if err != nil || got != "8000" {
		t.Errorf("DetectIosxrHardware(cached 8000) = (%q, %v); want (\"8000\", nil)", got, err)
	}

	// An empty cached value means the device did not report hardware-info
	hardwareCache.Store("nohw", "")
	if got, err := DetectIosxrHardware(context.Background(), client, "nohw"); err == nil {
		t.Errorf("DetectIosxrHardware(cached empty) = (%q, nil); want an error", got)
	}

	if _, err := DetectIosxrHardware(context.Background(), nil, "c8000"); err == nil {
		t.Error("DetectIosxrHardware(nil client) returned no error")
	}
}

func TestClearVersionCacheClearsHardware(t *testing.T) {
	ClearVersionCache("")
	t.Cleanup(func() { ClearVersionCache("") })

	versionCache.Store("a", "24.4")
	hardwareCache.Store("a", "8000")
	versionCache.Store("b", "25.4")
	hardwareCache.Store("b", "cisco NCS-5500 () processor")

	ClearVersionCache("a")
	if _, ok := versionCache.Load("a"); ok {
		t.Error("versionCache still has device a after ClearVersionCache(a)")
	}
	if _, ok := hardwareCache.Load("a"); ok {
		t.Error("hardwareCache still has device a after ClearVersionCache(a)")
	}
	if _, ok := hardwareCache.Load("b"); !ok {
		t.Error("ClearVersionCache(a) removed the hardware of device b")
	}

	ClearVersionCache("")
	if _, ok := hardwareCache.Load("b"); ok {
		t.Error("hardwareCache still has device b after ClearVersionCache(\"\")")
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

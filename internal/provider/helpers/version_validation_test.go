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
	"sort"
	"testing"
)

func TestVersionAtLeast(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		required string
		want     bool
	}{
		{name: "equal", current: "25.4", required: "25.4", want: true},
		{name: "newer_major", current: "26.2", required: "25.4", want: true},
		{name: "older_major", current: "24.4", required: "25.4", want: false},
		{name: "newer_minor", current: "25.4", required: "25.2", want: true},
		{name: "older_minor", current: "25.2", required: "25.4", want: false},
		{name: "two_digit_minor_above_one_digit", current: "25.10", required: "25.4", want: true},
		{name: "one_digit_minor_below_two_digit", current: "25.4", required: "25.10", want: false},
		{name: "two_digit_minor_below_next_major", current: "25.10", required: "26.0", want: false},
		{name: "next_major_above_two_digit_minor", current: "26.0", required: "25.20", want: true},
		{name: "patch_ignored_current", current: "25.20.3", required: "25.20", want: true},
		{name: "patch_ignored_required", current: "25.20", required: "25.20.9", want: true},
		{name: "two_digit_patch_ignored", current: "26.2.20", required: "26.2.9", want: true},
		{name: "empty_current_allows", current: "", required: "26.2", want: true},
		{name: "empty_required_allows", current: "24.4", required: "", want: true},
		{name: "unparseable_current_allows", current: "garbage", required: "26.2", want: true},
		{name: "unparseable_required_allows", current: "24.4", required: "garbage", want: true},
		{name: "single_part_current_allows", current: "25", required: "26.2", want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := VersionAtLeast(tc.current, tc.required)
			if got != tc.want {
				t.Errorf("VersionAtLeast(%q, %q) = %v; want %v", tc.current, tc.required, got, tc.want)
			}
		})
	}
}

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{name: "equal", a: "25.4", b: "25.4", want: 0},
		{name: "less_by_major", a: "24.4", b: "25.4", want: -1},
		{name: "greater_by_major", a: "26.2", b: "25.4", want: 1},
		{name: "less_by_minor", a: "25.2", b: "25.4", want: -1},
		{name: "greater_by_two_digit_minor", a: "25.10", b: "25.4", want: 1},
		{name: "less_by_one_digit_minor", a: "25.4", b: "25.10", want: -1},
		{name: "two_digit_minor_below_next_major", a: "25.20", b: "26.2", want: -1},
		{name: "patch_ignored", a: "25.4.10", b: "25.4.2", want: 0},
		{name: "unparseable_a", a: "garbage", b: "25.4", want: 0},
		{name: "unparseable_b", a: "25.4", b: "", want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := compareVersions(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("compareVersions(%q, %q) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestCompareVersionsSort(t *testing.T) {
	versions := []string{"25.10", "25.4", "24.4"}
	sort.Slice(versions, func(i, j int) bool { return compareVersions(versions[i], versions[j]) < 0 })

	want := []string{"24.4", "25.4", "25.10"}
	for i := range want {
		if versions[i] != want[i] {
			t.Fatalf("sorted versions = %v; want %v", versions, want)
		}
	}
}

func TestFormatVersion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "dotted", input: "25.4", want: "25.4"},
		{name: "dotted_two_digit_minor", input: "25.10", want: "25.10"},
		{name: "empty", input: "", want: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatVersion(tc.input)
			if got != tc.want {
				t.Errorf("FormatVersion(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

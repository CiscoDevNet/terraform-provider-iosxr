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

import "testing"

func TestAddVersionRangeDescription(t *testing.T) {
	tests := []struct {
		name   string
		ranges map[string]struct{ Min, Max int64 }
		want   string
	}{
		{
			name:   "empty",
			ranges: map[string]struct{ Min, Max int64 }{},
			want:   "base",
		},
		{
			name: "one_digit_minors",
			ranges: map[string]struct{ Min, Max int64 }{
				"25.4": {Min: 1, Max: 20},
				"24.4": {Min: 1, Max: 10},
			},
			want: "base\n  - Range: `1`-`10` (v24.4), `1`-`20` (v25.4)",
		},
		{
			name: "two_digit_minor_sorts_after_one_digit",
			ranges: map[string]struct{ Min, Max int64 }{
				"25.10": {Min: 1, Max: 30},
				"25.4":  {Min: 1, Max: 20},
				"24.4":  {Min: 1, Max: 10},
			},
			want: "base\n  - Range: `1`-`10` (v24.4), `1`-`20` (v25.4), `1`-`30` (v25.10)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewAttributeDescription("base").AddVersionRangeDescription(tc.ranges).String
			if got != tc.want {
				t.Errorf("AddVersionRangeDescription() = %q; want %q", got, tc.want)
			}
		})
	}
}

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

func TestSelectYangPath(t *testing.T) {
	single := map[string]string{"24.4": "old/path", "25.4": "new/path"}
	three := map[string]string{
		"24.4": "monitor",
		"25.4": "monitor.monitor-level",
		"26.2": "monitor.monitor-level.mode",
	}

	tests := []struct {
		name         string
		version      string
		versionPaths map[string]string
		defaultPath  string
		want         string
	}{
		// single-hop (24.4 → 25.4)
		{"exact match on moved version returns new path", "25.4", single, "old/path", "new/path"},
		{"version below threshold returns old path", "24.4", single, "old/path", "old/path"},
		{"version above threshold returns new path", "26.1", single, "old/path", "new/path"},
		{"empty version returns defaultPath (safe fallback)", "", single, "old/path", "old/path"},
		{"patch component ignored: 24.4.2 treated as 24.4", "24.4.2", single, "old/path", "old/path"},
		{"patch component ignored: 25.4.2 treated as 25.4", "25.4.2", single, "old/path", "new/path"},
		// three-hop (24.4 → 25.4 → 26.2)
		{"three-hop: 24.4 returns base path", "24.4", three, "monitor", "monitor"},
		{"three-hop: 25.1 between 24.4 and 25.4 returns 24.4 path", "25.1", three, "monitor", "monitor"},
		{"three-hop: 25.4 returns second path", "25.4", three, "monitor", "monitor.monitor-level"},
		{"three-hop: 26.2 returns third path", "26.2", three, "monitor", "monitor.monitor-level.mode"},
		{"three-hop: 27.0 above all thresholds inherits 26.2", "27.0", three, "monitor", "monitor.monitor-level.mode"},
		{"three-hop: empty version returns defaultPath", "", three, "monitor", "monitor"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SelectYangPath(tc.version, tc.versionPaths, tc.defaultPath)
			if got != tc.want {
				t.Errorf("SelectYangPath(%q, %v, %q) = %q, want %q",
					tc.version, tc.versionPaths, tc.defaultPath, got, tc.want)
			}
		})
	}
}

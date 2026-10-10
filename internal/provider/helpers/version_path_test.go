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

func TestGetPathVersion(t *testing.T) {
	const (
		defPath   = "Cisco-IOS-XR-um-logging-cfg:/logging"
		newPath   = "Cisco-IOS-XR-um-logging-ng-cfg:/logging"
		newerPath = "Cisco-IOS-XR-um-logging-ng2-cfg:/logging"
	)

	byVersion := map[string]string{
		"25.4": newPath,
	}

	twoThresholds := map[string]string{
		"25.2": newPath,
		"25.4": newerPath,
	}

	tests := []struct {
		name      string
		version   string
		defPath   string
		byVersion map[string]string
		wantPath  string
	}{
		{
			name:      "empty version returns defaultPath",
			version:   "",
			defPath:   defPath,
			byVersion: byVersion,
			wantPath:  defPath,
		},
		{
			name:      "empty pathByVersion returns defaultPath",
			version:   "25.4",
			defPath:   defPath,
			byVersion: map[string]string{},
			wantPath:  defPath,
		},
		{
			name:      "version below all thresholds returns defaultPath",
			version:   "24.4",
			defPath:   defPath,
			byVersion: byVersion,
			wantPath:  defPath,
		},
		{
			name:      "exact match on threshold returns that path",
			version:   "25.4",
			defPath:   defPath,
			byVersion: byVersion,
			wantPath:  newPath,
		},
		{
			name:      "version above threshold returns that path",
			version:   "25.6",
			defPath:   defPath,
			byVersion: byVersion,
			wantPath:  newPath,
		},
		{
			name:      "version between two thresholds returns lower threshold path",
			version:   "25.3",
			defPath:   defPath,
			byVersion: twoThresholds,
			wantPath:  newPath,
		},
		{
			name:      "exact match on upper threshold returns upper path",
			version:   "25.4",
			defPath:   defPath,
			byVersion: twoThresholds,
			wantPath:  newerPath,
		},
		{
			name:      "version above all thresholds returns highest threshold path",
			version:   "26.1",
			defPath:   defPath,
			byVersion: twoThresholds,
			wantPath:  newerPath,
		},
		{
			name:      "patch component ignored: 24.4.2 treated as 24.4",
			version:   "24.4.2",
			defPath:   defPath,
			byVersion: byVersion,
			wantPath:  defPath,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetPathVersion(tc.version, tc.defPath, tc.byVersion)
			if got != tc.wantPath {
				t.Errorf("GetPathVersion(%q, %q, ...) = %q, want %q",
					tc.version, tc.defPath, got, tc.wantPath)
			}
		})
	}
}

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

// GetPathVersion returns the gNMI module path appropriate for currentVersion.
//
// pathByVersion maps a minimum version threshold to the path that should be
// used when the device is running that version or later. The highest threshold
// that is still ≤ currentVersion wins.
//
// If currentVersion is empty, or no threshold in pathByVersion is satisfied,
// defaultPath is returned — preserving backward compatibility with devices that
// do not report a version.
func GetPathVersion(currentVersion, defaultPath string, pathByVersion map[string]string) string {
	if currentVersion == "" || len(pathByVersion) == 0 {
		return defaultPath
	}

	bestVersion := ""
	bestPath := defaultPath

	for threshold, path := range pathByVersion {
		if VersionAtLeast(currentVersion, threshold) {
			// Pick the highest threshold satisfied.
			if bestVersion == "" || VersionAtLeast(threshold, bestVersion) {
				bestVersion = threshold
				bestPath = path
			}
		}
	}

	return bestPath
}

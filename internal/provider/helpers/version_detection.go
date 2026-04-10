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
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-gnmi"
)

// versionCache stores detected IOS-XR versions to avoid redundant queries
// Key: deviceName, Value: normalized version string (e.g., "2442")
var versionCache sync.Map

// SupportedVersions is the single source of truth for all supported IOS-XR versions.
// Key: dotted format (user-facing), Value: internal compact format used by validation logic.
var SupportedVersions = map[string]string{
	"24.4.2": "2442",
	"25.2.2": "2522",
}

// ParseVersion accepts either dotted format ("24.4.2") or compact format ("2442")
// and returns the internal compact format. Returns ("", false) if unsupported.
func ParseVersion(version string) (string, bool) {
	// Already compact format (e.g., "2442")
	for _, compact := range SupportedVersions {
		if version == compact {
			return compact, true
		}
	}
	// Dotted format (e.g., "24.4.2")
	if compact, ok := SupportedVersions[version]; ok {
		return compact, true
	}
	return "", false
}

// ValidateSupportedVersion checks if a version (compact or dotted) is supported
func ValidateSupportedVersion(version string) bool {
	_, ok := ParseVersion(version)
	return ok
}

// SupportedVersionList returns a sorted comma-separated list of supported versions (dotted format)
func SupportedVersionList() string {
	versions := make([]string, 0, len(SupportedVersions))
	for dotted := range SupportedVersions {
		versions = append(versions, dotted)
	}
	// Simple sort for consistent output
	for i := 0; i < len(versions); i++ {
		for j := i + 1; j < len(versions); j++ {
			if versions[i] > versions[j] {
				versions[i], versions[j] = versions[j], versions[i]
			}
		}
	}
	result := ""
	for i, v := range versions {
		if i > 0 {
			result += ", "
		}
		result += v
	}
	return result
}

// DetectIosxrVersion queries a device via gNMI to detect its IOS-XR version
// and returns the normalized version string (e.g., "2442" for 24.4.2)
// Uses Cisco IOS-XR unified YANG models (not OpenConfig)
// Results are cached per device to avoid redundant queries
func DetectIosxrVersion(ctx context.Context, client *gnmi.Client, deviceName string) (string, error) {
	if client == nil {
		return "", fmt.Errorf("gNMI client is nil")
	}

	// Check cache first
	if cached, ok := versionCache.Load(deviceName); ok {
		if version, ok := cached.(string); ok {
			tflog.Debug(ctx, fmt.Sprintf("Using cached IOS-XR version for device '%s': %s", deviceName, version))
			return version, nil
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Attempting to auto-detect IOS-XR version for device '%s'", deviceName))

	// First try: gNMI Capabilities Version field
	version, err := detectFromCapabilities(ctx, client)
	if err == nil && version != "" {
		if compact, ok := ParseVersion(version); ok {
			tflog.Info(ctx, fmt.Sprintf("Auto-detected IOS-XR version for device '%s' from gNMI capabilities: %s", deviceName, version))
			versionCache.Store(deviceName, compact)
			return compact, nil
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Failed to detect version from capabilities: %v", err))

	// Second try: CLI configuration path - most reliable
	result, err := client.Get(ctx, []string{"/Cisco-IOS-XR-cli-cfg:cli"})
	if err != nil {
		return "", fmt.Errorf("unable to auto-detect IOS-XR version from device: %w", err)
	}

	version, err = extractVersionFromResponse(ctx, result)
	if err != nil {
		return "", fmt.Errorf("unable to auto-detect IOS-XR version from device: %w", err)
	}

	compact, ok := ParseVersion(version)
	if !ok {
		return "", fmt.Errorf("detected unsupported IOS-XR version '%s'. Supported versions: %s", version, SupportedVersionList())
	}

	tflog.Info(ctx, fmt.Sprintf("Auto-detected IOS-XR version for device '%s': %s", deviceName, version))
	versionCache.Store(deviceName, compact)
	return compact, nil
}

// detectFromCapabilities tries to extract version from gNMI Capabilities response
// Note: This rarely succeeds - the CLI config path is the reliable method
func detectFromCapabilities(ctx context.Context, client *gnmi.Client) (string, error) {
	tflog.Debug(ctx, "Trying to detect version from gNMI Capabilities")

	// Get capabilities
	caps, err := client.Capabilities(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get capabilities: %w", err)
	}

	// Try to extract version from the Version field
	if caps.Version != "" {
		if v := extractVersionString(caps.Version); v != "" {
			tflog.Info(ctx, fmt.Sprintf("Found version in gNMI capabilities Version field: %s", v))
			return v, nil
		}
	}

	return "", fmt.Errorf("no version information found in capabilities")
}

// extractVersionFromResponse attempts to extract version information from gNMI response
// Handles Cisco IOS-XR CLI config which contains version in the header
func extractVersionFromResponse(ctx context.Context, response interface{}) (string, error) {
	// Convert response to JSON for easier parsing
	jsonData, err := json.Marshal(response)
	if err != nil {
		return "", fmt.Errorf("failed to marshal response: %w", err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Response JSON (first 1000 chars): %s", truncate(string(jsonData), 1000)))

	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Extract version from CLI configuration
	version := tryExtractFromCLI(data)
	if version != "" {
		tflog.Debug(ctx, fmt.Sprintf("Found version from CLI path: %s", version))
		return version, nil
	}

	return "", fmt.Errorf("no version information found in response")
}

// tryExtractFromCLI tries to extract version from CLI config/output
func tryExtractFromCLI(data map[string]interface{}) string {
	// The CLI data is in the gNMI response under "Notifications"
	// Navigate to the actual CLI content
	if notifications, ok := data["Notifications"].([]interface{}); ok && len(notifications) > 0 {
		if notif, ok := notifications[0].(map[string]interface{}); ok {
			if updates, ok := notif["update"].([]interface{}); ok && len(updates) > 0 {
				if update, ok := updates[0].(map[string]interface{}); ok {
					if val, ok := update["val"].(map[string]interface{}); ok {
						if value, ok := val["Value"].(map[string]interface{}); ok {
							if jsonIetfVal, ok := value["JsonIetfVal"].(string); ok {
								// Decode base64-encoded CLI output
								if decoded := decodeBase64CLI(jsonIetfVal); decoded != "" {
									if v := extractVersionFromCLIString(decoded); v != "" {
										return v
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return ""
}

// extractVersionFromCLIString parses CLI output text to find version information
// Example CLI output: "!! IOS XR Configuration 24.4.2"
func extractVersionFromCLIString(cliOutput string) string {
	// Split by newlines and search each line
	lines := strings.Split(cliOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Look for IOS XR Configuration line which reliably contains version
		if strings.Contains(strings.ToLower(line), "ios xr configuration") {
			if v := extractVersionString(line); v != "" {
				return v
			}
		}
	}

	return ""
}

// decodeBase64CLI decodes base64-encoded CLI output
func decodeBase64CLI(encoded string) string {
	// Try to decode as base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		// Not base64 or invalid, return empty
		return ""
	}
	return string(decoded)
}

// truncate helper to limit string length for logging
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// extractVersionString extracts version string from a text
func extractVersionString(text string) string {
	// Pattern to match version formats like:
	// - 7.5.2.28I
	// - 24.4.2
	// - 25.2.2
	patterns := []string{
		`(\d+\.\d+\.\d+\.\d+[A-Za-z]*)`, // 7.5.2.28I
		`(\d+\.\d+\.\d+)`,               // 24.4.2 or 25.2.2
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if matches := re.FindStringSubmatch(text); len(matches) > 1 {
			return matches[1]
		}
	}

	return ""
}

// FormatVersionError creates a user-friendly error message for version detection failures
func FormatVersionError(deviceName, host string, err error) string {
	return fmt.Sprintf(`Unable to auto-detect IOS-XR version for device '%s' at %s.

Error details: %v

Please explicitly specify the 'iosxr_version' attribute in your provider configuration:

  provider "iosxr" {
    iosxr_version = "24.4.2"  # optional, auto-detected if not set
    devices = [...]
  }

Supported versions: %s`, deviceName, host, err, SupportedVersionList())
}

// ClearVersionCache clears the version cache for a specific device or all devices
// If deviceName is empty, clears the entire cache
func ClearVersionCache(deviceName string) {
	if deviceName == "" {
		// Clear entire cache
		versionCache = sync.Map{}
	} else {
		// Clear specific device
		versionCache.Delete(deviceName)
	}
}

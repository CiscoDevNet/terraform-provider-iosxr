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
	"strconv"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-gnmi"
)

// versionCache stores detected IOS-XR versions to avoid redundant queries
// Key: deviceName, Value: normalized version string (e.g., "2442")
var versionCache sync.Map

// NormalizeVersion converts any user-facing version string to the canonical
// internal "MM.mm" (major.minor) dotted format. The patch component is always
// stripped so that "25.2.1", "25.2.2", and "25.2" all map to the same key "25.2".
//
// Accepted inputs:
//
//	3-part dotted   "24.4.2" → "24.4"   (patch stripped)
//	2-part dotted   "25.2"   → "25.2"   (already major.minor)
//	4-digit compact "2442"   → "24.4"   (legacy format, backward-compat)
//
// Returns ("", false) when the input cannot be parsed.
func NormalizeVersion(version string) (string, bool) {
	version = strings.TrimSpace(version)
	if version == "" {
		return "", false
	}

	// Legacy 4-digit compact format (e.g. "2442", "2512") → convert to dotted major.minor.
	// Format: first 2 chars = major, 3rd char = minor (patch digit is dropped).
	if len(version) == 4 && !strings.Contains(version, ".") {
		if _, err := strconv.Atoi(version); err == nil {
			return version[0:2] + "." + string(version[2]), true
		}
	}

	parts := strings.Split(version, ".")
	switch len(parts) {
	case 2:
		// Already "MM.mm" — validate both components are integers.
		if _, err := strconv.Atoi(parts[0]); err != nil {
			return "", false
		}
		if _, err := strconv.Atoi(parts[1]); err != nil {
			return "", false
		}
		return version, true
	case 3:
		// "MM.mm.pp" → strip patch, return "MM.mm".
		// All three parts must be valid integers (e.g. "24.4.3" → "24.4").
		if _, err := strconv.Atoi(parts[0]); err != nil {
			return "", false
		}
		if _, err := strconv.Atoi(parts[1]); err != nil {
			return "", false
		}
		if _, err := strconv.Atoi(parts[2]); err != nil {
			return "", false
		}
		return parts[0] + "." + parts[1], true
	default:
		return "", false
	}
}

// ParseVersion accepts any supported version format (dotted or compact) and returns
// the 4-digit internal compact format.  Returns ("", false) if parsing fails.
// Deprecated: prefer NormalizeVersion directly; ParseVersion is kept for compatibility.
func ParseVersion(version string) (string, bool) {
	return NormalizeVersion(version)
}

// ValidateSupportedVersion checks whether a version string can be successfully normalized.
// Any well-formed version (e.g., "25.2", "25.2.1", "2521") is accepted.
func ValidateSupportedVersion(version string) bool {
	_, ok := NormalizeVersion(version)
	return ok
}

// SupportedVersionList returns a human-readable description of accepted version formats.
func SupportedVersionList() string {
	return "MM.mm.pp (e.g. 25.2.1) or MM.mm (e.g. 25.2) — patch component is ignored, only major.minor matters"
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

	compact, ok := NormalizeVersion(version)
	if !ok {
		return "", fmt.Errorf("detected IOS-XR version '%s' could not be parsed. Expected format: %s", version, SupportedVersionList())
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

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
	"strconv"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-gnmi"
	"github.com/tidwall/gjson"
)

// versionCache stores detected IOS-XR versions to avoid redundant queries
// Key: deviceName, Value: normalized version string (e.g., "24.4")
var versionCache sync.Map

// hardwareCache stores the hardware identifier each device reports in the hardware-info field of
// install/version, to avoid redundant queries
// Key: deviceName, Value: trimmed hardware-info string (e.g., "8000"), empty if the device did not report it
var hardwareCache sync.Map

// NormalizeVersion converts any user-facing version string to the canonical
// internal "MM.mm" (major.minor) dotted format. The patch component is always
// stripped so that "25.2.1", "25.2.2", and "25.2" all map to the same key "25.2".
//
// Accepted inputs:
//
//	3-part dotted   "24.4.2" → "24.4"   (patch stripped)
//	2-part dotted   "25.2"   → "25.2"   (already major.minor)
//
// Returns ("", false) when the input cannot be parsed.
func NormalizeVersion(version string) (string, bool) {
	version = strings.TrimSpace(version)
	if version == "" {
		return "", false
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

// ParseVersion accepts any supported dotted version format and returns the
// normalized major.minor string.  Returns ("", false) if parsing fails.
// Deprecated: prefer NormalizeVersion directly; ParseVersion is kept for compatibility.
func ParseVersion(version string) (string, bool) {
	return NormalizeVersion(version)
}

// ValidateSupportedVersion checks whether a version string can be successfully normalized.
// Any well-formed dotted version (e.g., "25.2", "25.2.1") is accepted.
func ValidateSupportedVersion(version string) bool {
	_, ok := NormalizeVersion(version)
	return ok
}

// SupportedVersionList returns a human-readable description of accepted version formats.
func SupportedVersionList() string {
	return fmt.Sprintf(
		"Major.Minor (e.g. 25.4) or Major.Minor.Patch (e.g. 25.4.2) — Note: patch version is ignored\nSupported IOS-XR versions: %s",
		strings.Join(DefinitionVersions, ", "),
	)
}

// DetectIosxrVersion queries a device via gNMI to detect its IOS-XR version
// and returns the normalized version string (e.g., "24.4" for 24.4.2)
// Reads the label from Cisco-IOS-XR-install-oper:install/version
// Results are cached per device to avoid redundant queries
// The hardware-info field of the same response is cached for DetectIosxrHardware
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

	// Note: gNMI Capabilities Version field is not used — IOS-XR returns the gNMI
	// library version there, not the OS version.

	version, hardware, err := queryInstallVersion(ctx, client)
	if err != nil {
		return "", fmt.Errorf("unable to auto-detect IOS-XR version from device: %w", err)
	}

	normalized, ok := NormalizeVersion(version)
	if !ok {
		return "", fmt.Errorf("detected IOS-XR version '%s' could not be parsed. Expected format: %s", version, SupportedVersionList())
	}

	tflog.Info(ctx, fmt.Sprintf("Auto-detected IOS-XR version for device '%s': %s", deviceName, version))
	versionCache.Store(deviceName, normalized)
	hardwareCache.Store(deviceName, hardware)
	return normalized, nil
}

// DetectIosxrHardware returns the hardware identifier a device reports in the hardware-info field of
// Cisco-IOS-XR-install-oper:install/version (e.g., "8000" for a Cisco 8000)
// It is the counterpart of DetectIosxrVersion and shares its cache: after DetectIosxrVersion has run for the
// device the value is already cached and no second request is made
// A device that does not report hardware-info is an error, like a missing version label
func DetectIosxrHardware(ctx context.Context, client *gnmi.Client, deviceName string) (string, error) {
	if client == nil {
		return "", fmt.Errorf("gNMI client is nil")
	}

	// Check cache first. An empty cached value means the device did not report hardware-info
	if cached, ok := hardwareCache.Load(deviceName); ok {
		if hardware, ok := cached.(string); ok {
			if hardware == "" {
				return "", fmt.Errorf("unable to auto-detect IOS-XR hardware from device: no hardware-info in install/version response")
			}
			tflog.Debug(ctx, fmt.Sprintf("Using cached IOS-XR hardware for device '%s': %s", deviceName, hardware))
			return hardware, nil
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Attempting to auto-detect IOS-XR hardware for device '%s'", deviceName))

	version, hardware, err := queryInstallVersion(ctx, client)
	if err != nil {
		return "", fmt.Errorf("unable to auto-detect IOS-XR hardware from device: %w", err)
	}

	// The same response carries the version, so cache it as well
	if normalized, ok := NormalizeVersion(version); ok {
		versionCache.Store(deviceName, normalized)
	}
	hardwareCache.Store(deviceName, hardware)

	if hardware == "" {
		return "", fmt.Errorf("unable to auto-detect IOS-XR hardware from device: no hardware-info in install/version response")
	}

	tflog.Info(ctx, fmt.Sprintf("Auto-detected IOS-XR hardware for device '%s': %s", deviceName, hardware))
	return hardware, nil
}

// queryInstallVersion reads Cisco-IOS-XR-install-oper:install/version with one gNMI request and returns the
// release label and the hardware-info field. The hardware is empty if the device did not report it; only a
// missing label is an error. Callers wrap the error with their own context.
func queryInstallVersion(ctx context.Context, client *gnmi.Client) (string, string, error) {
	// install-oper version - small payload, label holds the OS release
	result, err := client.Get(ctx, []string{"/Cisco-IOS-XR-install-oper:install/version"})
	if err != nil {
		return "", "", err
	}

	return extractInstallInfoFromResponse(ctx, result)
}

// extractInstallInfoFromResponse attempts to extract version and hardware information from gNMI response
// Handles the Cisco IOS-XR install-oper version payload, which carries the release in its label field
// and the hardware identifier in its hardware-info field
func extractInstallInfoFromResponse(ctx context.Context, response interface{}) (string, string, error) {
	// Convert response to JSON for easier parsing
	jsonData, err := json.Marshal(response)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal response: %w", err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Response JSON (first 1000 chars): %s", truncate(string(jsonData), 1000)))

	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Extract version from the install-oper label
	payload := decodeInstallPayload(data)
	version := extractVersionFromDevice(payload)
	if version == "" {
		return "", "", fmt.Errorf("no label in install/version response")
	}
	tflog.Debug(ctx, fmt.Sprintf("Found version from install/version label: %s", version))

	// hardware-info is optional here; DetectIosxrHardware decides if its absence is an error
	hardware := extractHardwareFromDevice(payload)
	tflog.Debug(ctx, fmt.Sprintf("Found hardware from install/version hardware-info: %q", hardware))

	return version, hardware, nil
}

// decodeInstallPayload returns the decoded install/version JSON payload from the gNMI response, or "" if it cannot be found
func decodeInstallPayload(data map[string]interface{}) string {
	// The payload is in the gNMI response under "Notifications"
	// Navigate to the actual install/version content
	if notifications, ok := data["Notifications"].([]interface{}); ok && len(notifications) > 0 {
		if notif, ok := notifications[0].(map[string]interface{}); ok {
			if updates, ok := notif["update"].([]interface{}); ok && len(updates) > 0 {
				if update, ok := updates[0].(map[string]interface{}); ok {
					if val, ok := update["val"].(map[string]interface{}); ok {
						if value, ok := val["Value"].(map[string]interface{}); ok {
							if jsonIetfVal, ok := value["JsonIetfVal"].(string); ok {
								// Decode base64-encoded JSON payload
								return decodeBase64Payload(jsonIetfVal)
							}
						}
					}
				}
			}
		}
	}

	return ""
}

// extractVersionFromDevice reads the release label from an install/version JSON payload
// Example payload: {"label": "24.4.2", "hardware-info": "8000", ...}
func extractVersionFromDevice(payload string) string {
	return gjson.Get(payload, "label").String()
}

// extractHardwareFromDevice reads the hardware-info field from an install/version JSON payload
// Example payload: {"label": "24.4.2", "hardware-info": "8000", ...}
func extractHardwareFromDevice(payload string) string {
	return strings.TrimSpace(gjson.Get(payload, "hardware-info").String())
}

// decodeBase64Payload decodes a base64-encoded JSON payload
func decodeBase64Payload(encoded string) string {
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

// FormatVersionError creates a user-friendly error message for version detection failures
func FormatVersionError(deviceName, host string, err error) string {
	return fmt.Sprintf(`Unable to auto-detect IOS-XR version for device '%s' at %s.

Error details: %v

Check that the device is reachable and that the username and password are correct.
If the device should not be managed, set 'managed = false' for it.

Supported versions: %s`, deviceName, host, err, SupportedVersionList())
}

// FormatHardwareError creates a user-friendly error message for hardware detection failures
func FormatHardwareError(deviceName, host string, err error) string {
	return fmt.Sprintf(`Unable to auto-detect IOS-XR hardware for device '%s' at %s.

Error details: %v

Check that the device is reachable and that the username and password are correct.
If the device should not be managed, set 'managed = false' for it.`, deviceName, host, err)
}

// ClearVersionCache clears the version and hardware caches for a specific device or all devices
// If deviceName is empty, clears the entire cache
func ClearVersionCache(deviceName string) {
	if deviceName == "" {
		// Clear entire cache
		versionCache = sync.Map{}
		hardwareCache = sync.Map{}
	} else {
		// Clear specific device
		versionCache.Delete(deviceName)
		hardwareCache.Delete(deviceName)
	}
}

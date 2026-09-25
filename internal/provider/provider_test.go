// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"sync"
	"testing"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// Singleton provider instance to enable connection reuse across test steps
	testProviderInstance     provider.Provider
	testProviderInstanceOnce sync.Once

	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"iosxr": func() (tfprotov6.ProviderServer, error) {
			// Use a singleton provider instance to enable connection reuse across test steps
			testProviderInstanceOnce.Do(func() {
				testProviderInstance = New()
			})
			return providerserver.NewProtocol6(testProviderInstance)(), nil
		},
	}
)

// iosxrVersionAtLeast returns true when currentVersion meets minVersion, or when
// currentVersion is empty (permissive default so tests aren't skipped without a version set).
func iosxrVersionAtLeast(currentVersion, minVersion string) bool {
	return helpers.VersionAtLeast(currentVersion, minVersion)
}

// selectVersionExample returns the example value appropriate for the IOSXR_VERSION in the
// environment. Keys in byVersion are version thresholds (e.g. "25.4"); the highest threshold
// satisfied by IOSXR_VERSION wins. Falls back to baseExample when no threshold matches or
// IOSXR_VERSION is unset.
func selectVersionExample(byVersion map[string]string, baseExample string) string {
	ver := os.Getenv("IOSXR_VERSION")
	if ver == "" || len(byVersion) == 0 {
		return baseExample
	}
	bestVer := ""
	for v := range byVersion {
		if iosxrVersionAtLeast(ver, v) && (bestVer == "" || iosxrVersionAtLeast(v, bestVer)) {
			bestVer = v
		}
	}
	if bestVer == "" {
		return baseExample
	}
	return byVersion[bestVer]
}

// selectVersionTestTags returns the test tag set appropriate for the IOSXR_VERSION in the
// environment. Keys in byVersion are version thresholds; the highest threshold satisfied by
// IOSXR_VERSION wins. Falls back to baseTags when no threshold matches or IOSXR_VERSION is unset.
func selectVersionTestTags(byVersion map[string][]string, baseTags []string) []string {
	ver := os.Getenv("IOSXR_VERSION")
	if ver == "" || len(byVersion) == 0 {
		return baseTags
	}
	bestVer := ""
	for v := range byVersion {
		if iosxrVersionAtLeast(ver, v) && (bestVer == "" || iosxrVersionAtLeast(v, bestVer)) {
			bestVer = v
		}
	}
	if bestVer == "" {
		return baseTags
	}
	return byVersion[bestVer]
}

// selectVersionPrerequisitesConfig returns the prerequisite HCL config string for the
// exact IOS-XR version in the environment -- unlike selectVersionTestTags/
// selectVersionExample, there is no "highest version <= current wins" fallback.
// test_prerequisites deliberately does not inherit across versions: a version not
// present in configByVersion (including when IOSXR_VERSION is unset) has no
// prerequisites at all.
func selectVersionPrerequisitesConfig(configByVersion map[string]string) string {
	return configByVersion[os.Getenv("IOSXR_VERSION")]
}

// selectVersionDependsOn returns the version-appropriate depends_on line (with leading
// tab, e.g. "\tdepends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]"), or an empty
// string when the current version has no prerequisites.
func selectVersionDependsOn(dependsByVersion map[string]string) string {
	deps := selectVersionPrerequisitesConfig(dependsByVersion)
	if deps == "" {
		return ""
	}
	return "\tdepends_on = " + deps
}

// TestSelectVersionPrerequisitesConfig covers the exact-match-only behavior that
// distinguishes this selector from selectVersionTestTags/selectVersionExample: no
// "highest version <= current wins" fallback, and no default argument.
func TestSelectVersionPrerequisitesConfig(t *testing.T) {
	byVersion := map[string]string{
		"24.4": "base config",
		"25.4": "delta config",
	}

	t.Run("empty map returns empty string", func(t *testing.T) {
		t.Setenv("IOSXR_VERSION", "24.4")
		if got := selectVersionPrerequisitesConfig(nil); got != "" {
			t.Errorf("got %q, want empty string", got)
		}
	})

	t.Run("exact match on lower version", func(t *testing.T) {
		t.Setenv("IOSXR_VERSION", "24.4")
		if got := selectVersionPrerequisitesConfig(byVersion); got != "base config" {
			t.Errorf("got %q, want %q", got, "base config")
		}
	})

	t.Run("exact match on higher version", func(t *testing.T) {
		t.Setenv("IOSXR_VERSION", "25.4")
		if got := selectVersionPrerequisitesConfig(byVersion); got != "delta config" {
			t.Errorf("got %q, want %q", got, "delta config")
		}
	})

	t.Run("version not in map returns empty string, not a fallback", func(t *testing.T) {
		t.Setenv("IOSXR_VERSION", "26.2")
		if got := selectVersionPrerequisitesConfig(byVersion); got != "" {
			t.Errorf("got %q, want empty string (no fallback to a neighboring version)", got)
		}
	})

	t.Run("version below all keys returns empty string, not a fallback", func(t *testing.T) {
		t.Setenv("IOSXR_VERSION", "23.1")
		if got := selectVersionPrerequisitesConfig(byVersion); got != "" {
			t.Errorf("got %q, want empty string (no fallback to a neighboring version)", got)
		}
	})

	t.Run("IOSXR_VERSION unset returns empty string", func(t *testing.T) {
		os.Unsetenv("IOSXR_VERSION")
		if got := selectVersionPrerequisitesConfig(byVersion); got != "" {
			t.Errorf("got %q, want empty string", got)
		}
	})
}

// TestSelectVersionDependsOn covers the same exact-match semantics as
// TestSelectVersionPrerequisitesConfig, plus the "\tdepends_on = " prefix and the
// no-match/empty-string case, which must never leave a dangling "\tdepends_on = "
// with nothing after the "=" (invalid HCL).
func TestSelectVersionDependsOn(t *testing.T) {
	byVersion := map[string]string{
		"24.4": `[iosxr_gnmi.PreReq0, ]`,
	}

	t.Run("match returns prefixed depends_on line", func(t *testing.T) {
		t.Setenv("IOSXR_VERSION", "24.4")
		want := "\tdepends_on = [iosxr_gnmi.PreReq0, ]"
		if got := selectVersionDependsOn(byVersion); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("no match returns empty string, not a dangling prefix", func(t *testing.T) {
		t.Setenv("IOSXR_VERSION", "25.4")
		if got := selectVersionDependsOn(byVersion); got != "" {
			t.Errorf("got %q, want empty string", got)
		}
	})

	t.Run("IOSXR_VERSION unset returns empty string", func(t *testing.T) {
		os.Unsetenv("IOSXR_VERSION")
		if got := selectVersionDependsOn(byVersion); got != "" {
			t.Errorf("got %q, want empty string", got)
		}
	})
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	if v := os.Getenv("IOSXR_USERNAME"); v == "" {
		t.Fatal("IOSXR_USERNAME env variable must be set for acceptance tests")
	}
	if v := os.Getenv("IOSXR_PASSWORD"); v == "" {
		t.Fatal("IOSXR_PASSWORD env variable must be set for acceptance tests")
	}
	if v := os.Getenv("IOSXR_HOST"); v == "" {
		t.Fatal("IOSXR_HOST env variable must be set for acceptance tests")
	}
}

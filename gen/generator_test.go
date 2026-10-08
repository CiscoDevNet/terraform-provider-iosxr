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

// Tests for the generator merge logic.
//
// generator.go uses //go:build ignore (it is a go run tool, not a library),
// so these tests must be run with explicit file arguments:
//
//	go test -v generator.go generator_test.go
//
// Running "go test ./gen/..." will not find them.

//go:build ignore

package main

import (
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// mergeConfigs tests
// ---------------------------------------------------------------------------

func TestMergeConfigs(t *testing.T) {
	tests := []struct {
		name     string
		base     YamlConfig
		override YamlConfig
		check    func(t *testing.T, got YamlConfig)
	}{
		{
			name: "basic field override",
			base: YamlConfig{
				Name: "Logging",
				Path: "Cisco-IOS-XR-um-logging-cfg:/logging",
			},
			override: YamlConfig{
				Version:        "25.4",
				ResDescription: "Updated description",
				DocCategory:    "Logging",
			},
			check: func(t *testing.T, got YamlConfig) {
				if got.Name != "Logging" {
					t.Errorf("Name: got %q, want %q", got.Name, "Logging")
				}
				if got.Path != "Cisco-IOS-XR-um-logging-cfg:/logging" {
					t.Errorf("Path: got %q, should be unchanged", got.Path)
				}
				if got.ResDescription != "Updated description" {
					t.Errorf("ResDescription: got %q, want %q", got.ResDescription, "Updated description")
				}
				if got.DocCategory != "Logging" {
					t.Errorf("DocCategory: got %q, want %q", got.DocCategory, "Logging")
				}
			},
		},
		{
			name:     "name override",
			base:     YamlConfig{Name: "Service Timestamps Old", Path: "old-module:/service/timestamps"},
			override: YamlConfig{Version: "25.4", Name: "Service Timestamps", Path: "new-module:/service/timestamps"},
			check: func(t *testing.T, got YamlConfig) {
				if got.Name != "Service Timestamps" {
					t.Errorf("Name: got %q, want %q", got.Name, "Service Timestamps")
				}
				if got.Path != "new-module:/service/timestamps" {
					t.Errorf("Path: got %q, want %q", got.Path, "new-module:/service/timestamps")
				}
			},
		},
		{
			name:     "path unchanged",
			base:     YamlConfig{Name: "Logging", Path: "Cisco-IOS-XR-um-logging-cfg:/logging"},
			override: YamlConfig{Version: "25.4", ResDescription: "New desc"},
			check: func(t *testing.T, got YamlConfig) {
				if got.Path != "Cisco-IOS-XR-um-logging-cfg:/logging" {
					t.Errorf("Path: got %q, want %q (unchanged)", got.Path, "Cisco-IOS-XR-um-logging-cfg:/logging")
				}
			},
		},
		{
			name: "legacy resource",
			base: YamlConfig{
				Name: "Old Feature",
				Path: "old-module:/feature",
				Attributes: []YamlConfigAttribute{
					{YangName: "attr1", TfName: "attr1", Type: "String"},
				},
			},
			override: YamlConfig{Version: "25.4", Legacy: true},
			check: func(t *testing.T, got YamlConfig) {
				if got.RemovedInVersion != "25.4" {
					t.Errorf("RemovedInVersion: got %q, want %q", got.RemovedInVersion, "25.4")
				}
				if !got.Legacy {
					t.Error("Legacy: got false, want true")
				}
				// Attributes should be unchanged — early return preserves base state
				if len(got.Attributes) != 1 {
					t.Errorf("Attributes: got %d, want 1 (preserved from base)", len(got.Attributes))
				}
			},
		},
		{
			name:     "no_delete propagates",
			base:     YamlConfig{Name: "Feature", NoDelete: false},
			override: YamlConfig{Version: "25.4", NoDelete: true},
			check: func(t *testing.T, got YamlConfig) {
				if !got.NoDelete {
					t.Error("NoDelete: got false, want true")
				}
			},
		},
		{
			// Matches real usage (gen/definitions/25.4/service_timestamps.yaml): a single
			// delta authors both the old and new version's module path directly, rather
			// than the generator inferring divergence across merge steps.
			name: "path_version: single delta declares both keys",
			base: YamlConfig{Name: "Service Timestamps", Path: "old-module:/service/timestamps"},
			override: YamlConfig{
				Version: "25.4",
				PathVersion: map[string]string{
					"24.4": "old-module:/service/timestamps",
					"25.4": "new-module:/service/timestamps",
				},
			},
			check: func(t *testing.T, got YamlConfig) {
				if len(got.PathVersion) != 2 {
					t.Fatalf("PathVersion: got %d entries, want 2 (full map: %v)", len(got.PathVersion), got.PathVersion)
				}
				if got.PathVersion["24.4"] != "old-module:/service/timestamps" {
					t.Errorf("PathVersion[24.4]: got %q, want %q", got.PathVersion["24.4"], "old-module:/service/timestamps")
				}
				if got.PathVersion["25.4"] != "new-module:/service/timestamps" {
					t.Errorf("PathVersion[25.4]: got %q, want %q", got.PathVersion["25.4"], "new-module:/service/timestamps")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := mergeConfigs(tc.base, tc.override)
			tc.check(t, got)
		})
	}
}

// TestMergeConfigs_PathVersionThreeVersionChain covers the case a single mergeConfigs call
// can't exercise: path_version keys contributed by separate deltas across 3+ versions must
// all survive the cumulative fold, not just the keys from the single delta that happens to
// run last. Unlike VersionRanges/VersionEnums, path_version has no "did this change since
// the last merge" comparison to get wrong — mergeConfigs's union (gen/generator.go:2112-2120)
// just adds whatever keys each delta contributes — so this is a lower-risk field, but the
// accumulation itself was previously untested at the generator level for any version count.
func TestMergeConfigs_PathVersionThreeVersionChain(t *testing.T) {
	t.Run("keys contributed by separate deltas", func(t *testing.T) {
		base := YamlConfig{Name: "Feature", Path: "module-a:/feature"}
		after25 := mergeConfigs(base, YamlConfig{
			Version:     "25.4",
			PathVersion: map[string]string{"25.4": "module-b:/feature"},
		})
		after26 := mergeConfigs(after25, YamlConfig{
			Version:     "26.2",
			PathVersion: map[string]string{"26.2": "module-c:/feature"},
		})

		if len(after26.PathVersion) != 2 {
			t.Fatalf("PathVersion: got %d entries, want 2 (full map: %v)", len(after26.PathVersion), after26.PathVersion)
		}
		if after26.PathVersion["25.4"] != "module-b:/feature" {
			t.Errorf("PathVersion[25.4]: got %q, want %q (25.4's contribution must survive the 26.2 merge)",
				after26.PathVersion["25.4"], "module-b:/feature")
		}
		if after26.PathVersion["26.2"] != "module-c:/feature" {
			t.Errorf("PathVersion[26.2]: got %q, want %q", after26.PathVersion["26.2"], "module-c:/feature")
		}
	})

	t.Run("mixed multi-key delta then single-key delta", func(t *testing.T) {
		base := YamlConfig{Name: "Feature", Path: "module-a:/feature"}
		after25 := mergeConfigs(base, YamlConfig{
			Version: "25.4",
			PathVersion: map[string]string{
				"24.4": "module-a:/feature",
				"25.4": "module-b:/feature",
			},
		})
		after26 := mergeConfigs(after25, YamlConfig{
			Version:     "26.2",
			PathVersion: map[string]string{"26.2": "module-c:/feature"},
		})

		want := map[string]string{
			"24.4": "module-a:/feature",
			"25.4": "module-b:/feature",
			"26.2": "module-c:/feature",
		}
		if len(after26.PathVersion) != len(want) {
			t.Fatalf("PathVersion: got %d entries, want %d (full map: %v)", len(after26.PathVersion), len(want), after26.PathVersion)
		}
		for k, v := range want {
			if got := after26.PathVersion[k]; got != v {
				t.Errorf("PathVersion[%q]: got %q, want %q", k, got, v)
			}
		}
	})
}

// TestMergeConfigs_TestTagsThreeVersionChain covers F24 (BUG-3): resource-level test_tags
// must seed its version map with the "_base" sentinel, not the mutable base.Version field,
// or a divergence first occurring at fold 2+ mislabels the base entry under the wrong version.
func TestMergeConfigs_TestTagsThreeVersionChain(t *testing.T) {
	base := YamlConfig{Name: "Feature", TestTags: []string{"A"}}
	after25 := mergeConfigs(base, YamlConfig{Version: "25.4", TestTags: []string{"A"}})
	after26 := mergeConfigs(after25, YamlConfig{Version: "26.2", TestTags: []string{"B"}})

	if !stringSlicesEqual(after26.VersionTestTags["_base"], []string{"A"}) {
		t.Errorf("VersionTestTags[_base]: got %v, want %v (must not be mislabeled under an intermediate version string)",
			after26.VersionTestTags["_base"], []string{"A"})
	}
	if !stringSlicesEqual(after26.VersionTestTags["26.2"], []string{"B"}) {
		t.Errorf("VersionTestTags[26.2]: got %v, want %v", after26.VersionTestTags["26.2"], []string{"B"})
	}

	fixBaseVersionInRanges(&after26, "24.4")
	if !stringSlicesEqual(after26.VersionTestTags["24.4"], []string{"A"}) {
		t.Errorf("VersionTestTags[24.4] after fixup: got %v, want %v", after26.VersionTestTags["24.4"], []string{"A"})
	}
	if _, hasBase := after26.VersionTestTags["_base"]; hasBase {
		t.Error("VersionTestTags[_base]: still present after fixBaseVersionInRanges, want removed")
	}
}

// TestMergeConfigs_TestPrerequisitesRecordsOnlyDeclaredVersions: the merge records a
// version's test_prerequisites only when that version declares them. Base declares one
// prerequisite; 25.4 overrides with a two-prerequisite list; 26.2 declares none and gets
// no entry. Inheritance by higher versions happens at test time in
// selectVersionPrerequisitesConfig, not here.
func TestMergeConfigs_TestPrerequisitesRecordsOnlyDeclaredVersions(t *testing.T) {
	prereq24 := []YamlTest{{Path: "module-a:/prereq"}}
	prereq25 := []YamlTest{{Path: "module-a:/prereq"}, {Path: "module-b:/prereq"}}

	base := YamlConfig{Name: "Feature", TestPrerequisites: prereq24}
	after25 := mergeConfigs(base, YamlConfig{Version: "25.4", TestPrerequisites: prereq25})
	after26 := mergeConfigs(after25, YamlConfig{Version: "26.2"})

	if _, has24 := after26.VersionTestPrerequisites["24.4"]; has24 {
		t.Error("VersionTestPrerequisites[24.4]: present before fixBaseVersionInRanges, want absent (base-seeding hasn't run yet)")
	}
	if len(after26.VersionTestPrerequisites["25.4"]) != 2 {
		t.Errorf("VersionTestPrerequisites[25.4]: got %d entries, want 2", len(after26.VersionTestPrerequisites["25.4"]))
	}
	if _, has26 := after26.VersionTestPrerequisites["26.2"]; has26 {
		t.Error("VersionTestPrerequisites[26.2]: present, want absent (26.2 never declared its own test_prerequisites; inheritance is resolved by the test selector, not recorded by the merge)")
	}
	if _, hasBase := after26.VersionTestPrerequisites["_base"]; hasBase {
		t.Error("VersionTestPrerequisites[_base]: present, want absent (test_prerequisites never seeds a \"_base\" sentinel)")
	}

	// fixBaseVersionInRanges' base-seeding case requires SupportedVersions to know
	// whether this resource has more than one version -- mergeConfigs itself never
	// sets that field, only the real generation pipeline's outer loop does.
	after26.SupportedVersions = []string{"24.4", "25.4", "26.2"}
	fixBaseVersionInRanges(&after26, "24.4")

	if len(after26.VersionTestPrerequisites["24.4"]) != 1 {
		t.Errorf("VersionTestPrerequisites[24.4] after fixup: got %d entries, want 1", len(after26.VersionTestPrerequisites["24.4"]))
	}
	if _, has26 := after26.VersionTestPrerequisites["26.2"]; has26 {
		t.Error("VersionTestPrerequisites[26.2]: present after fixBaseVersionInRanges, want absent")
	}
	if _, hasBase := after26.VersionTestPrerequisites["_base"]; hasBase {
		t.Error("VersionTestPrerequisites[_base]: present after fixBaseVersionInRanges, want absent")
	}
}

// TestMergeConfigs_TestPrerequisitesSingleVersion_NoVersionMap covers the guard on
// fixBaseVersionInRanges' base-seeding case: a resource with only one supported version
// has no neighboring version to isolate from, so it must keep the plain single-constant
// template path instead of being forced into VersionTestPrerequisites.
func TestMergeConfigs_TestPrerequisitesSingleVersion_NoVersionMap(t *testing.T) {
	config := YamlConfig{Name: "Feature", TestPrerequisites: []YamlTest{{Path: "module-a:/prereq"}}}
	config.SupportedVersions = []string{"24.4"}

	fixBaseVersionInRanges(&config, "24.4")

	if config.VersionTestPrerequisites != nil {
		t.Errorf("VersionTestPrerequisites: got %v, want nil (single-version resource must not be forced into the versioned path)",
			config.VersionTestPrerequisites)
	}
	if len(config.TestPrerequisites) != 1 {
		t.Errorf("TestPrerequisites (plain scalar): got %d entries, want 1 (unchanged)", len(config.TestPrerequisites))
	}
}

// ---------------------------------------------------------------------------
// mergeAttributes tests
// ---------------------------------------------------------------------------

func TestMergeAttributes(t *testing.T) {
	tests := []struct {
		name     string
		base     []YamlConfigAttribute
		override []YamlConfigAttribute
		check    func(t *testing.T, got []YamlConfigAttribute)
	}{
		{
			name: "new attr gets AddedInVersion",
			base: []YamlConfigAttribute{
				{YangName: "console", TfName: "console", Type: "String"},
			},
			override: []YamlConfigAttribute{
				{YangName: "new-leaf", TfName: "new_leaf", Type: "String"},
			},
			check: func(t *testing.T, got []YamlConfigAttribute) {
				if len(got) != 2 {
					t.Fatalf("len: got %d, want 2", len(got))
				}
				newAttr := got[1]
				if newAttr.YangName != "new-leaf" {
					t.Errorf("YangName: got %q, want %q", newAttr.YangName, "new-leaf")
				}
				if newAttr.AddedInVersion != "25.4" {
					t.Errorf("AddedInVersion: got %q, want %q", newAttr.AddedInVersion, "25.4")
				}
			},
		},
		{
			name: "legacy attr gets RemovedInVersion",
			base: []YamlConfigAttribute{
				{YangName: "source-interface-name", TfName: "name", Type: "String", Id: true},
			},
			override: []YamlConfigAttribute{
				{YangName: "source-interface-name", Legacy: true},
			},
			check: func(t *testing.T, got []YamlConfigAttribute) {
				if len(got) != 1 {
					t.Fatalf("len: got %d, want 1 (legacy keeps the attr)", len(got))
				}
				if got[0].RemovedInVersion != "25.4" {
					t.Errorf("RemovedInVersion: got %q, want %q", got[0].RemovedInVersion, "25.4")
				}
				if !got[0].Legacy {
					t.Error("Legacy: got false, want true")
				}
				// TfName should be unchanged when override doesn't specify one
				if got[0].TfName != "name" {
					t.Errorf("TfName: got %q, want %q (unchanged)", got[0].TfName, "name")
				}
			},
		},
		{
			// F7 fix: legacy block with explicit tf_name renames the base attribute
			// so the natural name is freed for a replacement attribute.
			name: "legacy attr renames tf_name",
			base: []YamlConfigAttribute{
				{YangName: "console", TfName: "console", Type: "String"},
			},
			override: []YamlConfigAttribute{
				{YangName: "console", TfName: "console_legacy", Legacy: true},
			},
			check: func(t *testing.T, got []YamlConfigAttribute) {
				if len(got) != 1 {
					t.Fatalf("len: got %d, want 1", len(got))
				}
				if got[0].TfName != "console_legacy" {
					t.Errorf("TfName: got %q, want %q", got[0].TfName, "console_legacy")
				}
				if got[0].RemovedInVersion != "25.4" {
					t.Errorf("RemovedInVersion: got %q, want %q", got[0].RemovedInVersion, "25.4")
				}
			},
		},
		{
			// A legacy entry in the override that has no matching base attr is a no-op.
			name: "legacy attr not in base is skipped",
			base: []YamlConfigAttribute{
				{YangName: "existing", TfName: "existing", Type: "String"},
			},
			override: []YamlConfigAttribute{
				{YangName: "ghost", Legacy: true},
			},
			check: func(t *testing.T, got []YamlConfigAttribute) {
				if len(got) != 1 {
					t.Errorf("len: got %d, want 1 (ghost attr should be skipped)", len(got))
				}
			},
		},
		{
			name: "existing attr fields updated",
			base: []YamlConfigAttribute{
				{YangName: "severity", TfName: "severity", Type: "String", Description: "old desc"},
			},
			override: []YamlConfigAttribute{
				{YangName: "severity", Description: "new desc"},
			},
			check: func(t *testing.T, got []YamlConfigAttribute) {
				if len(got) != 1 {
					t.Fatalf("len: got %d, want 1", len(got))
				}
				if got[0].Description != "new desc" {
					t.Errorf("Description: got %q, want %q", got[0].Description, "new desc")
				}
				// Version not stamped on existing attrs
				if got[0].AddedInVersion != "" {
					t.Errorf("AddedInVersion: got %q, want empty (existing attr)", got[0].AddedInVersion)
				}
			},
		},
		{
			// When yang_names differ but tf_names match, merge uses tf_name match.
			name: "match by tf_name",
			base: []YamlConfigAttribute{
				{YangName: "old-yang", TfName: "shared_name", Type: "String"},
			},
			override: []YamlConfigAttribute{
				{YangName: "new-yang", TfName: "shared_name", Type: "Int64"},
			},
			check: func(t *testing.T, got []YamlConfigAttribute) {
				if len(got) != 1 {
					t.Fatalf("len: got %d, want 1 (matched by tf_name)", len(got))
				}
				if got[0].Type != "Int64" {
					t.Errorf("Type: got %q, want %q", got[0].Type, "Int64")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := mergeAttributes(tc.base, tc.override, "25.4")
			tc.check(t, got)
		})
	}
}

func TestMergeAttributes_NestedAttrsInheritAddedInVersion(t *testing.T) {
	// A completely new list attr in the override: its nested attrs should be
	// stamped with AddedInVersion too.
	base := []YamlConfigAttribute{}
	override := []YamlConfigAttribute{
		{
			YangName: "new-list",
			TfName:   "new_list",
			Type:     "List",
			Attributes: []YamlConfigAttribute{
				{YangName: "key", TfName: "key", Type: "String", Id: true},
				{YangName: "value", TfName: "value", Type: "String"},
			},
		},
	}

	got := mergeAttributes(base, override, "25.4")

	if len(got) != 1 {
		t.Fatalf("len: got %d, want 1", len(got))
	}
	if got[0].AddedInVersion != "25.4" {
		t.Errorf("outer AddedInVersion: got %q, want %q", got[0].AddedInVersion, "25.4")
	}
	for _, child := range got[0].Attributes {
		if child.AddedInVersion != "25.4" {
			t.Errorf("child %q AddedInVersion: got %q, want %q",
				child.YangName, child.AddedInVersion, "25.4")
		}
	}
}

func TestMergeAttributes_ExcludeTestFromDeltaVersion(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "a", TfName: "a", Type: "String"},
		{YangName: "b", TfName: "b", Type: "String", ExcludeTest: true},
	}
	override := []YamlConfigAttribute{
		{YangName: "a", ExcludeTest: true},
		{YangName: "b", ExcludeTest: true},
		{YangName: "c", TfName: "c", Type: "String", ExcludeTest: true},
	}

	got := mergeAttributes(base, override, "26.2")

	if got[0].ExcludeTest || got[0].ExcludeTestFrom != "26.2" {
		t.Errorf("a: got ExcludeTest=%v ExcludeTestFrom=%q, want false and 26.2", got[0].ExcludeTest, got[0].ExcludeTestFrom)
	}
	if !got[1].ExcludeTest || got[1].ExcludeTestFrom != "" {
		t.Errorf("b: base exclusion must stay global, got ExcludeTest=%v ExcludeTestFrom=%q", got[1].ExcludeTest, got[1].ExcludeTestFrom)
	}
	if !got[2].ExcludeTest || got[2].ExcludeTestFrom != "" {
		t.Errorf("c: new attribute keeps ExcludeTest, got ExcludeTest=%v ExcludeTestFrom=%q", got[2].ExcludeTest, got[2].ExcludeTestFrom)
	}
}

func TestMergeAttributes_ExcludeTestFromChain(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "a", TfName: "a", Type: "String"},
		{YangName: "b", TfName: "b", Type: "String"},
	}
	d254 := []YamlConfigAttribute{{YangName: "a", ExcludeTest: true}}
	d262 := []YamlConfigAttribute{
		{YangName: "a", ExcludeTest: true},
		{YangName: "b", ExcludeTest: true},
	}

	got := mergeAttributes(mergeAttributes(base, d254, "25.4"), d262, "26.2")

	if got[0].ExcludeTestFrom != "25.4" || got[0].ExcludeTest {
		t.Errorf("a: lowest version must win, got ExcludeTest=%v ExcludeTestFrom=%q, want false and 25.4", got[0].ExcludeTest, got[0].ExcludeTestFrom)
	}
	if got[1].ExcludeTestFrom != "26.2" || got[1].ExcludeTest {
		t.Errorf("b: got ExcludeTest=%v ExcludeTestFrom=%q, want false and 26.2", got[1].ExcludeTest, got[1].ExcludeTestFrom)
	}
}

func TestTestVersionGuardExpr(t *testing.T) {
	const added = `iosxrVersionAtLeast(os.Getenv("IOSXR_VERSION"), "25.4")`
	const removed = `!iosxrVersionAtLeast(os.Getenv("IOSXR_VERSION"), "26.2")`
	cases := []struct {
		name string
		attr YamlConfigAttribute
		want string
	}{
		{"none", YamlConfigAttribute{}, ""},
		{"added", YamlConfigAttribute{AddedInVersion: "25.4"}, added},
		{"added and removed", YamlConfigAttribute{AddedInVersion: "25.4", RemovedInVersion: "26.2"}, added + " && " + removed},
		{"exclude from", YamlConfigAttribute{ExcludeTestFrom: "26.2"}, removed},
	}
	for _, c := range cases {
		if got := TestVersionGuardExpr(c.attr); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}

func TestMergeAttributes_CompositeKeyPromotion(t *testing.T) {
	// Reproduces the logging source_interfaces scenario:
	// 24.4 key: source-interface-name (id:true)
	// 25.4: retire source-interface-name, add interface-name + vrf-name (both id:true)
	base := []YamlConfigAttribute{
		{YangName: "source-interface-name", TfName: "name", Type: "String", Id: true},
	}
	override := []YamlConfigAttribute{
		{YangName: "source-interface-name", Legacy: true},
		{YangName: "interface-name", TfName: "interface_name", Type: "String", Id: true},
		{YangName: "vrf-name", TfName: "vrf_name", Type: "String", Id: true},
	}

	got := mergeAttributes(base, override, "25.4")

	// Expect 3 attrs: retired name, new interface_name, new vrf_name
	if len(got) != 3 {
		t.Fatalf("len: got %d, want 3", len(got))
	}

	// Index 0: source-interface-name retired
	if got[0].RemovedInVersion != "25.4" {
		t.Errorf("name RemovedInVersion: got %q, want %q", got[0].RemovedInVersion, "25.4")
	}
	if !got[0].Id {
		t.Error("name Id: got false, want true (should retain Id)")
	}

	// Index 1: interface-name added
	if got[1].YangName != "interface-name" {
		t.Errorf("got[1] YangName: got %q, want %q", got[1].YangName, "interface-name")
	}
	if got[1].AddedInVersion != "25.4" {
		t.Errorf("interface-name AddedInVersion: got %q, want %q", got[1].AddedInVersion, "25.4")
	}
	if !got[1].Id {
		t.Error("interface-name Id: got false, want true")
	}

	// Index 2: vrf-name added
	if got[2].YangName != "vrf-name" {
		t.Errorf("got[2] YangName: got %q, want %q", got[2].YangName, "vrf-name")
	}
	if got[2].AddedInVersion != "25.4" {
		t.Errorf("vrf-name AddedInVersion: got %q, want %q", got[2].AddedInVersion, "25.4")
	}
}

func TestMergeAttributes_ReplacesYangName_OnKeyAttr(t *testing.T) {
	// Verifies that replaces_yang_name on an id:true attribute is fully processed:
	// Phase 1 (mergeAttributes): VersionYangNames populated with "_base" placeholder.
	// Phase 2 (fixAttributeBaseVersion): "_base" replaced with real base version,
	// MovedInVersion derived.
	base := []YamlConfigAttribute{
		{YangName: "name", TfName: "name", Type: "String", Id: true},
	}
	override := []YamlConfigAttribute{
		{YangName: "host", TfName: "name", Type: "String", Id: true, ReplacesYangName: "name"},
	}

	// Phase 1
	got := mergeAttributes(base, override, "25.4")
	if len(got) != 1 {
		t.Fatalf("len: got %d, want 1", len(got))
	}
	if got[0].YangName != "host" {
		t.Errorf("YangName: got %q, want %q", got[0].YangName, "host")
	}
	if !got[0].Id {
		t.Error("Id: got false, want true")
	}
	if got[0].TfName != "name" {
		t.Errorf("TfName: got %q, want %q", got[0].TfName, "name")
	}
	if got[0].AddedInVersion != "" {
		t.Errorf("AddedInVersion: got %q, want empty (not treated as new attr)", got[0].AddedInVersion)
	}
	if got[0].RemovedInVersion != "" {
		t.Errorf("RemovedInVersion: got %q, want empty", got[0].RemovedInVersion)
	}
	if got[0].VersionYangNames["_base"] != "name" {
		t.Errorf("VersionYangNames[_base]: got %q, want %q", got[0].VersionYangNames["_base"], "name")
	}
	if got[0].VersionYangNames["25.4"] != "host" {
		t.Errorf("VersionYangNames[25.4]: got %q, want %q", got[0].VersionYangNames["25.4"], "host")
	}
	if got[0].MovedInVersion != "" {
		t.Errorf("MovedInVersion: got %q, want empty (not set until fixAttributeBaseVersion)", got[0].MovedInVersion)
	}

	// Phase 2
	fixAttributeBaseVersion(&got[0], "24.4")
	if got[0].VersionYangNames["24.4"] != "name" {
		t.Errorf("VersionYangNames[24.4]: got %q, want %q", got[0].VersionYangNames["24.4"], "name")
	}
	if _, hasBase := got[0].VersionYangNames["_base"]; hasBase {
		t.Error("VersionYangNames[_base]: still present after fixAttributeBaseVersion, want removed")
	}
	if got[0].MovedInVersion != "25.4" {
		t.Errorf("MovedInVersion: got %q, want %q", got[0].MovedInVersion, "25.4")
	}
}

func TestMergeAttributes_ReplacesYangName_ThreeVersionChain(t *testing.T) {
	// Simulates an attribute that moves in 25.4 and again in 26.2.
	// Each delta points replaces_yang_name at the previous canonical yang_name.
	base := []YamlConfigAttribute{
		{YangName: "monitor", TfName: "monitor", Type: "String"},
	}
	// First move: 24.4 "monitor" → 25.4 "monitor/monitor-level"
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "monitor/monitor-level", TfName: "monitor", Type: "String", ReplacesYangName: "monitor"},
	}, "25.4")

	// Second move: current canonical "monitor/monitor-level" → 26.2 "monitor/monitor-level/mode"
	after26 := mergeAttributes(after25, []YamlConfigAttribute{
		{YangName: "monitor/monitor-level/mode", TfName: "monitor", Type: "String", ReplacesYangName: "monitor/monitor-level"},
	}, "26.2")

	if len(after26) != 1 {
		t.Fatalf("len: got %d, want 1", len(after26))
	}
	attr := after26[0]
	if attr.YangName != "monitor/monitor-level/mode" {
		t.Errorf("YangName: got %q, want %q", attr.YangName, "monitor/monitor-level/mode")
	}
	if attr.VersionYangNames["_base"] != "monitor" {
		t.Errorf("VersionYangNames[_base]: got %q, want %q", attr.VersionYangNames["_base"], "monitor")
	}
	if attr.VersionYangNames["25.4"] != "monitor/monitor-level" {
		t.Errorf("VersionYangNames[25.4]: got %q, want %q", attr.VersionYangNames["25.4"], "monitor/monitor-level")
	}
	if attr.VersionYangNames["26.2"] != "monitor/monitor-level/mode" {
		t.Errorf("VersionYangNames[26.2]: got %q, want %q", attr.VersionYangNames["26.2"], "monitor/monitor-level/mode")
	}

	// After fixAttributeBaseVersion, "_base" → "24.4" and MovedInVersion is the earliest move.
	fixAttributeBaseVersion(&attr, "24.4")
	if attr.VersionYangNames["24.4"] != "monitor" {
		t.Errorf("VersionYangNames[24.4]: got %q, want %q", attr.VersionYangNames["24.4"], "monitor")
	}
	if _, hasBase := attr.VersionYangNames["_base"]; hasBase {
		t.Error("VersionYangNames[_base]: still present after fixAttributeBaseVersion, want removed")
	}
	// MovedInVersion is the earliest non-base version key ("25.4" < "26.2").
	if attr.MovedInVersion != "25.4" {
		t.Errorf("MovedInVersion: got %q, want %q", attr.MovedInVersion, "25.4")
	}
}

// ---------------------------------------------------------------------------
// Stale-comparison-baseline regression tests (F18/F19/F21/F22).
//
// Each reproduces a 3-version chain where the 3rd version reverts to a value
// already seen in an earlier version. Before the fix, each of these silently
// dropped the 3rd version's VersionXxx entry because the merge compared the
// incoming delta against something other than the true immediately-preceding
// version's own value.
// ---------------------------------------------------------------------------

func TestMergeAttributes_RangeBaselineFreeze_ThreeVersionChain(t *testing.T) {
	// base(24.4)={1,10} -> 25.4={1,20} -> 26.2 reverts to {1,10}.
	base := []YamlConfigAttribute{
		{YangName: "mtu", TfName: "mtu", Type: "Int64", MinInt: 1, MaxInt: 10},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "mtu", MinInt: 1, MaxInt: 20},
	}, "25.4")
	after26 := mergeAttributes(after25, []YamlConfigAttribute{
		{YangName: "mtu", MinInt: 1, MaxInt: 10},
	}, "26.2")

	if len(after26) != 1 {
		t.Fatalf("len: got %d, want 1", len(after26))
	}
	attr := after26[0]
	if got := attr.VersionRanges["26.2"]; got != (RangeConstraint{Min: 1, Max: 10}) {
		t.Errorf("VersionRanges[26.2]: got %+v, want %+v", got, RangeConstraint{Min: 1, Max: 10})
	}
	if got := attr.VersionRanges["25.4"]; got != (RangeConstraint{Min: 1, Max: 20}) {
		t.Errorf("VersionRanges[25.4]: got %+v, want %+v", got, RangeConstraint{Min: 1, Max: 20})
	}
	// The schema-facing scalar stays frozen at the original base value -- deliberately
	// unaffected by this fix (see F18's Fix section: no template reads it once
	// VersionRanges is set).
	if attr.MinInt != 1 || attr.MaxInt != 10 {
		t.Errorf("MinInt/MaxInt: got {%d,%d}, want {1,10} (frozen, unaffected by the fix)", attr.MinInt, attr.MaxInt)
	}

	fixAttributeBaseVersion(&attr, "24.4")
	if got := attr.VersionRanges["24.4"]; got != (RangeConstraint{Min: 1, Max: 10}) {
		t.Errorf("VersionRanges[24.4] after fixup: got %+v, want %+v", got, RangeConstraint{Min: 1, Max: 10})
	}
	if _, hasBase := attr.VersionRanges["_base"]; hasBase {
		t.Error("VersionRanges[_base]: still present after fixAttributeBaseVersion, want removed")
	}
}

func TestMergeAttributes_EnumUnionStaleness_ThreeVersionChain(t *testing.T) {
	// base=[a,b,c] -> 25.4 removes c ([a,b]) -> 26.2 re-adds c ([a,b,c]).
	base := []YamlConfigAttribute{
		{YangName: "mode", TfName: "mode", Type: "String", EnumValues: []string{"a", "b", "c"}},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "mode", EnumValues: []string{"a", "b"}},
	}, "25.4")
	after26 := mergeAttributes(after25, []YamlConfigAttribute{
		{YangName: "mode", EnumValues: []string{"a", "b", "c"}},
	}, "26.2")

	if len(after26) != 1 {
		t.Fatalf("len: got %d, want 1", len(after26))
	}
	attr := after26[0]
	if !stringSlicesEqual(attr.VersionEnums["26.2"], []string{"a", "b", "c"}) {
		t.Errorf("VersionEnums[26.2]: got %v, want %v", attr.VersionEnums["26.2"], []string{"a", "b", "c"})
	}
	if !stringSlicesEqual(attr.VersionEnums["25.4"], []string{"a", "b"}) {
		t.Errorf("VersionEnums[25.4]: got %v, want %v", attr.VersionEnums["25.4"], []string{"a", "b"})
	}
	// The schema-facing union is unaffected by this fix -- still the correct superset.
	if !stringSlicesEqual(attr.EnumValues, []string{"a", "b", "c"}) {
		t.Errorf("EnumValues (union): got %v, want %v", attr.EnumValues, []string{"a", "b", "c"})
	}

	fixAttributeBaseVersion(&attr, "24.4")
	if !stringSlicesEqual(attr.VersionEnums["24.4"], []string{"a", "b", "c"}) {
		t.Errorf("VersionEnums[24.4] after fixup: got %v, want %v", attr.VersionEnums["24.4"], []string{"a", "b", "c"})
	}
	if _, hasBase := attr.VersionEnums["_base"]; hasBase {
		t.Error("VersionEnums[_base]: still present after fixAttributeBaseVersion, want removed")
	}
}

func TestMergeAttributes_StringLengthEnvelopeStaleness_ThreeVersionChain(t *testing.T) {
	// base min=5 -> 25.4 sets min=10 (less permissive) -> 26.2 reverts to min=5.
	base := []YamlConfigAttribute{
		{YangName: "name", TfName: "name", Type: "String", StringMinLength: 5},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "name", StringMinLength: 10},
	}, "25.4")
	after26 := mergeAttributes(after25, []YamlConfigAttribute{
		{YangName: "name", StringMinLength: 5},
	}, "26.2")

	if len(after26) != 1 {
		t.Fatalf("len: got %d, want 1", len(after26))
	}
	attr := after26[0]
	if got := attr.VersionStringLengths["26.2"]; got != (StringLengthConstraint{Min: 5}) {
		t.Errorf("VersionStringLengths[26.2]: got %+v, want %+v", got, StringLengthConstraint{Min: 5})
	}
	if got := attr.VersionStringLengths["25.4"]; got != (StringLengthConstraint{Min: 10}) {
		t.Errorf("VersionStringLengths[25.4]: got %+v, want %+v", got, StringLengthConstraint{Min: 10})
	}
	// The widened schema-facing scalar is unaffected by this fix -- still load-bearing
	// (gen/templates/resource.go reads it even when VersionStringLengths is set).
	if attr.StringMinLength != 5 {
		t.Errorf("StringMinLength (widened): got %d, want 5", attr.StringMinLength)
	}

	fixAttributeBaseVersion(&attr, "24.4")
	if got := attr.VersionStringLengths["24.4"]; got != (StringLengthConstraint{Min: 5}) {
		t.Errorf("VersionStringLengths[24.4] after fixup: got %+v, want %+v", got, StringLengthConstraint{Min: 5})
	}
	if _, hasBase := attr.VersionStringLengths["_base"]; hasBase {
		t.Error("VersionStringLengths[_base]: still present after fixAttributeBaseVersion, want removed")
	}
}

func TestMergeAttributes_TypeYangBoolFreezeReuse_ThreeVersionChain(t *testing.T) {
	// base 24.4="presence" -> 25.4="empty" -> 26.2 reverts to "presence".
	base := []YamlConfigAttribute{
		{YangName: "monitor-receiver", TfName: "monitor_receiver", Type: "Bool", TypeYangBool: "presence"},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "monitor-receiver", TypeYangBool: "empty"},
	}, "25.4")
	after26 := mergeAttributes(after25, []YamlConfigAttribute{
		{YangName: "monitor-receiver", TypeYangBool: "presence"},
	}, "26.2")

	if len(after26) != 1 {
		t.Fatalf("len: got %d, want 1", len(after26))
	}
	attr := after26[0]
	if got := attr.VersionTypeYangBool["26.2"]; got != "presence" {
		t.Errorf("VersionTypeYangBool[26.2]: got %q, want %q", got, "presence")
	}
	if got := attr.VersionTypeYangBool["25.4"]; got != "empty" {
		t.Errorf("VersionTypeYangBool[25.4]: got %q, want %q", got, "empty")
	}
	// The deliberately-frozen default is unaffected by this fix -- still the base value,
	// since it's read as GetPathVersion's defaultValue argument by TypeYangBoolExpr.
	if attr.TypeYangBool != "presence" {
		t.Errorf("TypeYangBool (frozen default): got %q, want %q", attr.TypeYangBool, "presence")
	}
	// VersionTypeYangBool never seeds a "_base" key -- confirmed by construction, unlike
	// every other VersionXxx map in this file.
	if _, hasBase := attr.VersionTypeYangBool["_base"]; hasBase {
		t.Error("VersionTypeYangBool[_base]: present, want absent (TypeYangBoolExpr never needs one)")
	}
}

func TestMergeAttributes_XPathVersionBleed_ThreeVersionChain(t *testing.T) {
	// A bare xpath-only change with no rename: base xpath="a/b" -> 25.4 xpath="c/d" ->
	// 26.2 reverts to "a/b".
	base := []YamlConfigAttribute{
		{YangName: "leaf", TfName: "leaf", Type: "String", XPath: "a/b"},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "leaf", XPath: "c/d"},
	}, "25.4")
	after26 := mergeAttributes(after25, []YamlConfigAttribute{
		{YangName: "leaf", XPath: "a/b"},
	}, "26.2")

	if len(after26) != 1 {
		t.Fatalf("len: got %d, want 1", len(after26))
	}
	attr := after26[0]
	if got := attr.VersionXPath["26.2"]; got != "a/b" {
		t.Errorf("VersionXPath[26.2]: got %q, want %q", got, "a/b")
	}
	if got := attr.VersionXPath["25.4"]; got != "c/d" {
		t.Errorf("VersionXPath[25.4]: got %q, want %q", got, "c/d")
	}
	// .XPath keeps its existing last-version-wins role, used as JsonPathExpr/KeyPathExpr's
	// per-version default -- unaffected by this fix.
	if attr.XPath != "a/b" {
		t.Errorf("XPath (last-version-wins): got %q, want %q", attr.XPath, "a/b")
	}

	fixAttributeBaseVersion(&attr, "24.4")
	if got := attr.VersionXPath["24.4"]; got != "a/b" {
		t.Errorf("VersionXPath[24.4] after fixup: got %q, want %q", got, "a/b")
	}
	if _, hasBase := attr.VersionXPath["_base"]; hasBase {
		t.Error("VersionXPath[_base]: still present after fixAttributeBaseVersion, want removed")
	}
}

// ---------------------------------------------------------------------------
// Regression tests: a delta that merely restates the base's own value on the very
// FIRST fold must not be treated as a divergence. Found via real generated output
// (go generate against actual definitions) after the fixes above shipped: each new
// unexported "last*" tracking field starts at its zero value, so comparing a restated
// base value against that zero value looked like a change even when nothing differs.
// ---------------------------------------------------------------------------

func TestMergeAttributes_RangeBaselineFreeze_RestatedBaseValueIsNotADivergence(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "mtu", TfName: "mtu", Type: "Int64", MinInt: 1, MaxInt: 10},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "mtu", MinInt: 1, MaxInt: 10}, // restates the base's own value
	}, "25.4")

	attr := after25[0]
	if attr.VersionRanges != nil {
		t.Errorf("VersionRanges: got %v, want nil (25.4 restated the base's own {1,10}, not a real change)", attr.VersionRanges)
	}
}

func TestMergeAttributes_EnumUnionStaleness_RestatedBaseValueIsNotADivergence(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "mode", TfName: "mode", Type: "String", EnumValues: []string{"a", "b", "c"}},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "mode", EnumValues: []string{"a", "b", "c"}}, // restates the base's own list
	}, "25.4")

	attr := after25[0]
	if attr.VersionEnums != nil {
		t.Errorf("VersionEnums: got %v, want nil (25.4 restated the base's own list, not a real change)", attr.VersionEnums)
	}
}

func TestMergeAttributes_StringLengthEnvelopeStaleness_RestatedBaseValueIsNotADivergence(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "name", TfName: "name", Type: "String", StringMinLength: 5},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "name", StringMinLength: 5}, // restates the base's own value
	}, "25.4")

	attr := after25[0]
	if attr.VersionStringLengths != nil {
		t.Errorf("VersionStringLengths: got %v, want nil (25.4 restated the base's own min=5, not a real change)", attr.VersionStringLengths)
	}
}

func TestMergeAttributes_TypeYangBoolFreezeReuse_RestatedBaseValueIsNotADivergence(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "monitor-receiver", TfName: "monitor_receiver", Type: "Bool", TypeYangBool: "presence"},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "monitor-receiver", TypeYangBool: "presence"}, // restates the base's own value
	}, "25.4")

	attr := after25[0]
	if attr.VersionTypeYangBool != nil {
		t.Errorf("VersionTypeYangBool: got %v, want nil (25.4 restated the base's own \"presence\", not a real change)", attr.VersionTypeYangBool)
	}
}

// ---------------------------------------------------------------------------
// VersionDeleteMode merge scenarios
// ---------------------------------------------------------------------------

func TestVersionDeleteMode(t *testing.T) {
	tests := []struct {
		name                   string
		baseDeleteParent       bool
		baseDeleteGrandparent  bool
		deltaReplacesYangName  string
		deltaDeleteParent      bool
		deltaDeleteGrandparent bool
		wantVersionDeleteMode  map[string]string // nil means expect nil (inherited/unchanged)
		checkFlags             bool              // also assert the OR-merged DeleteParent/DeleteGrandparent scalars
		wantDeleteParent       bool
		wantDeleteGrandparent  bool
	}{
		{
			// base has delete_parent, 25.4 has same yang_name (no replaces_yang_name) and no delete flag.
			// Expected: 25.4 inherits delete_parent — no VersionDeleteMode created.
			name:             "inherit delete_parent (no rename, no flag)",
			baseDeleteParent: true,
			checkFlags:       true,
			wantDeleteParent: true,
		},
		{
			// base has delete_grandparent, 25.4 has same yang_name (no replaces_yang_name) and no delete flag.
			// Expected: 25.4 inherits delete_grandparent — no VersionDeleteMode created.
			name:                  "inherit delete_grandparent (no rename, no flag)",
			baseDeleteGrandparent: true,
			checkFlags:            true,
			wantDeleteGrandparent: true,
		},
		{
			// base has delete_parent, 25.4 has replaces_yang_name and NO delete flag.
			// Expected: version-specific map — 24.4 deletes parent, 25.4 deletes leaf directly.
			name:                  "replacement drops delete_parent (rename, no flag)",
			baseDeleteParent:      true,
			deltaReplacesYangName: "maxfilesize",
			wantVersionDeleteMode: map[string]string{"_base": "parent", "25.4": ""},
		},
		{
			// base has delete_parent, 25.4 has replaces_yang_name WITH delete_parent.
			// Expected: both versions delete parent — no VersionDeleteMode, static path.
			name:                  "replacement keeps delete_parent (rename, with flag)",
			baseDeleteParent:      true,
			deltaReplacesYangName: "maxfilesize",
			deltaDeleteParent:     true,
			checkFlags:            true,
			wantDeleteParent:      true,
		},
		{
			// Same as above, delete_grandparent variant for base.
			name:                  "replacement drops delete_grandparent (rename, no flag)",
			baseDeleteGrandparent: true,
			deltaReplacesYangName: "maxfilesize",
			wantVersionDeleteMode: map[string]string{"_base": "grandparent", "25.4": ""},
		},
		{
			// base has delete_grandparent, 25.4 has replaces_yang_name WITH delete_grandparent.
			// Expected: no VersionDeleteMode, same mode throughout.
			name:                   "replacement keeps delete_grandparent (rename, with flag)",
			baseDeleteGrandparent:  true,
			deltaReplacesYangName:  "maxfilesize",
			deltaDeleteGrandparent: true,
			checkFlags:             true,
			wantDeleteGrandparent:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			base := []YamlConfigAttribute{{
				YangName:          "maxfilesize",
				TfName:            "maxfilesize",
				DeleteParent:      tc.baseDeleteParent,
				DeleteGrandparent: tc.baseDeleteGrandparent,
			}}
			yangName := "maxfilesize"
			if tc.deltaReplacesYangName != "" {
				yangName = "path/maxfilesize"
			}
			delta := []YamlConfigAttribute{{
				YangName:          yangName,
				TfName:            "maxfilesize",
				ReplacesYangName:  tc.deltaReplacesYangName,
				DeleteParent:      tc.deltaDeleteParent,
				DeleteGrandparent: tc.deltaDeleteGrandparent,
			}}
			got := mergeAttributes(base, delta, "25.4")
			attr := got[0]

			if tc.wantVersionDeleteMode == nil {
				if attr.VersionDeleteMode != nil {
					t.Errorf("VersionDeleteMode should be nil, got %v", attr.VersionDeleteMode)
				}
			} else {
				if attr.VersionDeleteMode == nil {
					t.Fatal("VersionDeleteMode should be set")
				}
				for k, want := range tc.wantVersionDeleteMode {
					if got := attr.VersionDeleteMode[k]; got != want {
						t.Errorf("VersionDeleteMode[%q] = %q, want %q (full map: %v)", k, got, want, attr.VersionDeleteMode)
					}
				}
			}

			if tc.checkFlags {
				if attr.DeleteParent != tc.wantDeleteParent {
					t.Errorf("DeleteParent = %v, want %v", attr.DeleteParent, tc.wantDeleteParent)
				}
				if attr.DeleteGrandparent != tc.wantDeleteGrandparent {
					t.Errorf("DeleteGrandparent = %v, want %v", attr.DeleteGrandparent, tc.wantDeleteGrandparent)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// getVersionValue tests
// ---------------------------------------------------------------------------

func TestGetVersionValue(t *testing.T) {
	const (
		defVal   = "default"
		newVal   = "new"
		newerVal = "newer"
	)

	byVersion := map[string]string{
		"25.4": newVal,
	}

	twoThresholds := map[string]string{
		"25.2": newVal,
		"25.4": newerVal,
	}

	tests := []struct {
		name      string
		target    string
		byVersion map[string]string
		defVal    string
		want      string
	}{
		{
			name:      "empty map returns defaultValue",
			target:    "25.4",
			byVersion: map[string]string{},
			defVal:    defVal,
			want:      defVal,
		},
		{
			name:      "target below all thresholds returns defaultValue",
			target:    "24.4",
			byVersion: byVersion,
			defVal:    defVal,
			want:      defVal,
		},
		{
			name:      "exact match on threshold",
			target:    "25.4",
			byVersion: byVersion,
			defVal:    defVal,
			want:      newVal,
		},
		{
			name:      "target above threshold",
			target:    "25.6",
			byVersion: byVersion,
			defVal:    defVal,
			want:      newVal,
		},
		{
			name:      "target between two thresholds returns lower threshold value",
			target:    "25.3",
			byVersion: twoThresholds,
			defVal:    defVal,
			want:      newVal,
		},
		{
			name:      "exact match on upper threshold",
			target:    "25.4",
			byVersion: twoThresholds,
			defVal:    defVal,
			want:      newerVal,
		},
		{
			name:      "target above all thresholds returns highest threshold value",
			target:    "26.1",
			byVersion: twoThresholds,
			defVal:    defVal,
			want:      newerVal,
		},
		{
			name:      "patch component ignored: 24.4.2 treated as 24.4",
			target:    "24.4.2",
			byVersion: byVersion,
			defVal:    defVal,
			want:      defVal,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getVersionValue(tc.target, tc.byVersion, tc.defVal)
			if got != tc.want {
				t.Errorf("getVersionValue(%q, %v, %q) = %q, want %q",
					tc.target, tc.byVersion, tc.defVal, got, tc.want)
			}
		})
	}
}

// TestGetDeletePathExpr_RenameWithoutModeChange_ThreeVersionChain reproduces the read-side
// stale-fallback bug fixed by getVersionValue: a mode change recorded at 25.4, followed by a
// 26.2 rename that restates the same mode (correctly producing no new VersionDeleteMode
// entry), used to bake the wrong, stale oldest mode into 26.2's generated path.
func TestGetDeletePathExpr_RenameWithoutModeChange_ThreeVersionChain(t *testing.T) {
	base := []YamlConfigAttribute{{
		YangName:     "o1/m1/a",
		TfName:       "x",
		DeleteParent: true,
	}}
	delta25 := []YamlConfigAttribute{{
		YangName:          "o2/m2/b",
		TfName:            "x",
		ReplacesYangName:  "o1/m1/a",
		DeleteGrandparent: true,
	}}
	delta26 := []YamlConfigAttribute{{
		YangName:          "o3/m3/c",
		TfName:            "x",
		ReplacesYangName:  "o2/m2/b",
		DeleteGrandparent: true, // restated, same mode as 25.4
	}}

	merged := mergeAttributes(base, delta25, "25.4")
	merged = mergeAttributes(merged, delta26, "26.2")
	attr := merged[0]

	if attr.VersionDeleteMode == nil {
		t.Fatal("VersionDeleteMode should be set after the 25.4 delta")
	}
	if _, has26 := attr.VersionDeleteMode["26.2"]; has26 {
		t.Errorf("VersionDeleteMode should have no 26.2 entry (mode unchanged from 25.4), got %v", attr.VersionDeleteMode)
	}

	fixAttributeBaseVersion(&attr, "24.4")

	got := GetDeletePathExpr(attr, "version")
	want := `helpers.SelectYangPath(version, map[string]string{"24.4": "o1/m1", "25.4": "o2", "26.2": "o3"}, "o1/m1")`
	if got != want {
		t.Errorf("GetDeletePathExpr =\n  %s\nwant\n  %s", got, want)
	}
}

// TestJsonPathExpr_XPathOnlyChange_NoRename covers F20: a bare xpath-only change with
// no rename at all, across a 3-version chain that reverts back to the original value.
func TestJsonPathExpr_XPathOnlyChange_NoRename(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "leaf", TfName: "leaf", Type: "String", XPath: "a/b"},
	}
	merged := mergeAttributes(base, []YamlConfigAttribute{{YangName: "leaf", XPath: "c/d"}}, "25.4")
	merged = mergeAttributes(merged, []YamlConfigAttribute{{YangName: "leaf", XPath: "a/b"}}, "26.2")
	attr := merged[0]
	fixAttributeBaseVersion(&attr, "24.4")

	got := JsonPathExpr(attr, "version")
	want := `helpers.SelectYangPath(version, map[string]string{"24.4": "a.b", "25.4": "c.d", "26.2": "a.b"}, "")`
	if got != want {
		t.Errorf("JsonPathExpr =\n  %s\nwant\n  %s", got, want)
	}

	gotKey := KeyPathExpr(attr, "version")
	wantKey := `helpers.SelectYangPath(version, map[string]string{"24.4": "a/b", "25.4": "c/d", "26.2": "a/b"}, "")`
	if gotKey != wantKey {
		t.Errorf("KeyPathExpr =\n  %s\nwant\n  %s", gotKey, wantKey)
	}
}

// TestJsonPathExpr_XPathVersionBleed_CombinedWithRename_ThreeVersionChain covers F20's
// composition with a rename: 25.4 renames the attribute AND independently changes its
// xpath in the same delta; 26.2 renames again without restating the xpath, which must
// still carry forward via VersionXPath's own cascade (getVersionValue), independent of
// the rename-driven default.
func TestJsonPathExpr_XPathVersionBleed_CombinedWithRename_ThreeVersionChain(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "o1/m1", TfName: "x", XPath: "orig/path"},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "o2", TfName: "x", ReplacesYangName: "o1/m1", XPath: "x/y"},
	}, "25.4")
	after26 := mergeAttributes(after25, []YamlConfigAttribute{
		{YangName: "o3", TfName: "x", ReplacesYangName: "o2"},
	}, "26.2")
	attr := after26[0]
	fixAttributeBaseVersion(&attr, "24.4")

	got := JsonPathExpr(attr, "version")
	want := `helpers.SelectYangPath(version, map[string]string{"24.4": "orig.path", "25.4": "x.y", "26.2": "x.y"}, "orig.path")`
	if got != want {
		t.Errorf("JsonPathExpr =\n  %s\nwant\n  %s", got, want)
	}
}

// F27/BUG-10: GetWidestRange had no test coverage despite being registered in the template
// FuncMap -- it was never called from any template until F27 wired it into the schema's Range
// validator.
func TestGetWidestRange_ThreeVersionChain(t *testing.T) {
	versionRanges := map[string]RangeConstraint{
		"24.4": {Min: 1, Max: 10},
		"25.4": {Min: 1, Max: 20},
		"26.2": {Min: 5, Max: 15},
	}
	got := GetWidestRange(versionRanges)
	want := []int64{1, 20}
	if len(got) != 2 || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetWidestRange = %v, want %v", got, want)
	}
}

func TestGetWidestRange_Empty(t *testing.T) {
	got := GetWidestRange(nil)
	want := []int64{0, 0}
	if len(got) != 2 || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("GetWidestRange(nil) = %v, want %v", got, want)
	}
}

// F28/BUG-11: hasAttributeVersionDifferences only checked AddedInVersion/RemovedInVersion/
// VersionRanges/VersionYangNames, silently omitting HasVersionDifferences from generated
// Create/Update/Delete functions for a resource whose only divergence was VersionEnums,
// VersionStringLengths, or VersionPatterns -- meaning ValidateVersionEnums/StringLengths/Patterns
// would never run for that resource at all.
func TestHasAttributeVersionDifferences_VersionEnumsOnly_True(t *testing.T) {
	attrs := []YamlConfigAttribute{
		{YangName: "severity", VersionEnums: map[string][]string{"25.4": {"alerts", "critical"}}},
	}
	if !hasAttributeVersionDifferences(attrs) {
		t.Error("hasAttributeVersionDifferences = false, want true for VersionEnums-only divergence")
	}
}

func TestHasAttributeVersionDifferences_VersionStringLengthsOnly_True(t *testing.T) {
	attrs := []YamlConfigAttribute{
		{YangName: "name", VersionStringLengths: map[string]StringLengthConstraint{"25.4": {Min: 1, Max: 32}}},
	}
	if !hasAttributeVersionDifferences(attrs) {
		t.Error("hasAttributeVersionDifferences = false, want true for VersionStringLengths-only divergence")
	}
}

func TestHasAttributeVersionDifferences_VersionPatternsOnly_True(t *testing.T) {
	attrs := []YamlConfigAttribute{
		{YangName: "hostname", VersionPatterns: map[string][]string{"25.4": {`^[a-z]+$`}}},
	}
	if !hasAttributeVersionDifferences(attrs) {
		t.Error("hasAttributeVersionDifferences = false, want true for VersionPatterns-only divergence")
	}
}

// VersionDeleteMode is a defensive-only addition -- it can never independently trip true in
// practice (it always co-occurs with VersionYangNames, already checked), but this confirms the
// check works in isolation regardless.
func TestHasAttributeVersionDifferences_VersionDeleteModeOnly_True(t *testing.T) {
	attrs := []YamlConfigAttribute{
		{YangName: "path", VersionDeleteMode: map[string]string{"25.4": "parent"}},
	}
	if !hasAttributeVersionDifferences(attrs) {
		t.Error("hasAttributeVersionDifferences = false, want true for VersionDeleteMode-only divergence")
	}
}

func TestHasAttributeVersionDifferences_NestedListAttribute(t *testing.T) {
	attrs := []YamlConfigAttribute{
		{
			YangName: "file", TfName: "file", Type: "List",
			Attributes: []YamlConfigAttribute{
				{YangName: "severity", VersionEnums: map[string][]string{"25.4": {"alerts"}}},
			},
		},
	}
	if !hasAttributeVersionDifferences(attrs) {
		t.Error("hasAttributeVersionDifferences = false, want true for VersionEnums nested inside a List attribute")
	}
}

// F26/BUG-5: findNoAugmentConfigViolations/validateNoAugmentConfigCarryover -- an attribute (or
// whole resource) with no_augment_config: true in an earlier version must have that flag
// restated in any later delta that re-lists it, or real YANG augmentation would silently
// overwrite hand-authored fields. See provider-yaml-authoring-rules.md Rule 4.

func TestFindNoAugmentConfigViolations_MissingRestatement_Fails(t *testing.T) {
	acc := YamlConfig{
		Name: "Logging",
		Attributes: []YamlConfigAttribute{
			{YangName: "archive/threshold", NoAugmentConfig: true, Type: "Int64"},
		},
	}
	raw := YamlConfig{
		Name: "Logging",
		Attributes: []YamlConfigAttribute{
			{YangName: "archive/threshold", Example: "80"}, // re-listed for an unrelated reason, flag omitted
		},
	}
	violations := findNoAugmentConfigViolations(acc, raw, "25.4")
	if len(violations) != 1 {
		t.Fatalf("violations = %v, want exactly 1", violations)
	}
	if !strings.Contains(violations[0], `"archive/threshold"`) || !strings.Contains(violations[0], "25.4") {
		t.Errorf("violation message = %q, want it to name the attribute and version", violations[0])
	}
}

func TestFindNoAugmentConfigViolations_CorrectRestatement_Passes(t *testing.T) {
	acc := YamlConfig{
		Name: "Logging",
		Attributes: []YamlConfigAttribute{
			{YangName: "archive/threshold", NoAugmentConfig: true, Type: "Int64"},
		},
	}
	raw := YamlConfig{
		Name: "Logging",
		Attributes: []YamlConfigAttribute{
			{YangName: "archive/threshold", NoAugmentConfig: true, Example: "80"},
		},
	}
	if violations := findNoAugmentConfigViolations(acc, raw, "25.4"); len(violations) != 0 {
		t.Errorf("violations = %v, want none (flag correctly restated)", violations)
	}
}

func TestFindNoAugmentConfigViolations_ThreeVersionChain(t *testing.T) {
	base := []YamlConfigAttribute{
		{YangName: "archive/threshold", NoAugmentConfig: true, Type: "Int64"},
	}
	after25 := mergeAttributes(base, []YamlConfigAttribute{
		{YangName: "archive/threshold", NoAugmentConfig: true, Example: "80"},
	}, "25.4")
	acc26 := YamlConfig{Name: "Logging", Attributes: after25}
	raw26 := YamlConfig{
		Name: "Logging",
		Attributes: []YamlConfigAttribute{
			{YangName: "archive/threshold", Example: "90"}, // v3 forgets the flag
		},
	}
	violations := findNoAugmentConfigViolations(acc26, raw26, "26.2")
	if len(violations) != 1 {
		t.Fatalf("violations = %v, want exactly 1 (attributed to 26.2)", violations)
	}
	if !strings.Contains(violations[0], "26.2") {
		t.Errorf("violation message = %q, want it attributed to 26.2, not an earlier version", violations[0])
	}
}

func TestFindNoAugmentConfigViolations_NestedListAttribute(t *testing.T) {
	acc := YamlConfig{
		Name: "Logging",
		Attributes: []YamlConfigAttribute{
			{
				YangName: "files/file", TfName: "file", Type: "List",
				Attributes: []YamlConfigAttribute{
					{YangName: "maxfilesize", NoAugmentConfig: true, Type: "Int64"},
				},
			},
		},
	}
	raw := YamlConfig{
		Name: "Logging",
		Attributes: []YamlConfigAttribute{
			{
				YangName: "files/file", TfName: "file", Type: "List",
				Attributes: []YamlConfigAttribute{
					{YangName: "maxfilesize", Example: "1000"}, // nested re-list, flag omitted
				},
			},
		},
	}
	violations := findNoAugmentConfigViolations(acc, raw, "25.4")
	if len(violations) != 1 {
		t.Fatalf("violations = %v, want exactly 1 from the nested list attribute", violations)
	}
	if !strings.Contains(violations[0], `"maxfilesize"`) {
		t.Errorf("violation message = %q, want it to name the nested attribute", violations[0])
	}
}

func TestFindNoAugmentConfigViolations_ResourceLevel_MissingRestatement_Fails(t *testing.T) {
	acc := YamlConfig{Name: "SomeResource", NoAugmentConfig: true}
	raw := YamlConfig{Name: "SomeResource"}
	violations := findNoAugmentConfigViolations(acc, raw, "25.4")
	if len(violations) != 1 {
		t.Fatalf("violations = %v, want exactly 1 (resource-level)", violations)
	}
	if !strings.Contains(violations[0], "SomeResource") || !strings.Contains(violations[0], "resource-level") {
		t.Errorf("violation message = %q, want it to identify the resource and call out resource-level", violations[0])
	}
}

func TestFindNoAugmentConfigViolations_ResourceLevel_CorrectRestatement_Passes(t *testing.T) {
	acc := YamlConfig{Name: "SomeResource", NoAugmentConfig: true}
	raw := YamlConfig{Name: "SomeResource", NoAugmentConfig: true}
	if violations := findNoAugmentConfigViolations(acc, raw, "25.4"); len(violations) != 0 {
		t.Errorf("violations = %v, want none (resource-level flag correctly restated)", violations)
	}
}

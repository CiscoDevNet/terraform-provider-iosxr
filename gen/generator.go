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

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/openconfig/goyang/pkg/yang"
	"gopkg.in/yaml.v3"
)

const (
	definitionsPath   = "./gen/definitions/"
	modelsPath        = "./gen/models/"
	providerTemplate  = "./gen/templates/provider.go"
	providerLocation  = "./internal/provider/provider.go"
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogLocation = "./templates/guides/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
)

type t struct {
	path          string
	prefix        string
	suffix        string
	versionSuffix bool // indicates if version should be added to filename
}

var templates = []t{
	{
		path:          "./gen/templates/model.go",
		prefix:        "./internal/provider/model_iosxr_",
		suffix:        ".go",
		versionSuffix: false, // Unified files, no version suffix
	},
	{
		path:          "./gen/templates/data_source.go",
		prefix:        "./internal/provider/data_source_iosxr_",
		suffix:        ".go",
		versionSuffix: false, // Unified files, no version suffix
	},
	{
		path:          "./gen/templates/data_source_test.go",
		prefix:        "./internal/provider/data_source_iosxr_",
		suffix:        "_test.go",
		versionSuffix: false, // Unified files, no version suffix
	},
	{
		path:          "./gen/templates/resource.go",
		prefix:        "./internal/provider/resource_iosxr_",
		suffix:        ".go",
		versionSuffix: false, // Unified files, no version suffix
	},
	{
		path:          "./gen/templates/resource_test.go",
		prefix:        "./internal/provider/resource_iosxr_",
		suffix:        "_test.go",
		versionSuffix: false, // Unified files, no version suffix
	},
	{
		path:          "./gen/templates/data-source.tf",
		prefix:        "./examples/data-sources/iosxr_",
		suffix:        "/data-source.tf",
		versionSuffix: false,
	},
	{
		path:          "./gen/templates/resource.tf",
		prefix:        "./examples/resources/iosxr_",
		suffix:        "/resource.tf",
		versionSuffix: false,
	},
	{
		path:          "./gen/templates/import.sh",
		prefix:        "./examples/resources/iosxr_",
		suffix:        "/import.sh",
		versionSuffix: false,
	},
}

type YamlConfig struct {
	Name                     string                `yaml:"name"`
	Version                  string                `yaml:"version"` // drives type-name suffix; empty for unified files
	SupportedVersions        []string              // All versions this resource supports (for unified files)
	BaseVersion              string                // The minimum/base version (first version in SupportedVersions)
	IntroducedInVersion      string                // Set when this resource only exists from a higher version (not the global min)
	HasVersionDifferences    bool                  // True if there are version-specific changes (added/removed fields or definitions)
	Path                     string                `yaml:"path"`
	AugmentPath              string                `yaml:"augment_path"`
	NoDelete                 bool                  `yaml:"no_delete"`
	NoDeleteAttributes       bool                  `yaml:"no_delete_attributes"`
	DefaultDeleteAttributes  bool                  `yaml:"default_delete_attributes"`
	TestTags                 []string              `yaml:"test_tags"`
	VersionTestTags          map[string][]string   // computed during merge: version → tag set (nil if same across all versions)
	SkipMinimumTest          bool                  `yaml:"skip_minimum_test"`
	NoAugmentConfig          bool                  `yaml:"no_augment_config"`
	DsDescription            string                `yaml:"ds_description"`
	ResDescription           string                `yaml:"res_description"`
	DocCategory              string                `yaml:"doc_category"`
	Legacy                   bool                  `yaml:"legacy"` // If true, entire resource is removed/not available in this version
	RemovedInVersion         string                // Version where entire resource was removed (set when Legacy is true)
	PathVersion              map[string]string     `yaml:"path_version"` // version threshold → gNMI module path override
	HasPathVersion           bool                  // True when PathVersion is non-empty (computed after merge)
	Attributes               []YamlConfigAttribute `yaml:"attributes"`
	TestPrerequisites        []YamlTest            `yaml:"test_prerequisites"`
	VersionTestPrerequisites map[string][]YamlTest // computed during merge: version → prerequisite set (nil if same across all versions)
}

type YamlConfigAttribute struct {
	YangName                 string                            `yaml:"yang_name"`
	YangScope                string                            `yaml:"yang_scope"`
	TfName                   string                            `yaml:"tf_name"`
	XPath                    string                            `yaml:"xpath"`
	VersionXPath             map[string]string                 // Version-specific xpath overrides (nil if same across all versions); independent of VersionYangNames
	Type                     string                            `yaml:"type"`
	ReadRaw                  bool                              `yaml:"read_raw"`
	TypeYangBool             string                            `yaml:"type_yang_bool"`
	VersionTypeYangBool      map[string]string                 // internal — computed by mergeAttributes; no yaml tag
	lastTypeYangBool         string                            // unexported: last delta's own explicit value, decoupled from the deliberately-frozen TypeYangBool default above. Never read by templates.
	Id                       bool                              `yaml:"id"`
	Reference                bool                              `yaml:"reference"`
	Mandatory                bool                              `yaml:"mandatory"`
	Optional                 bool                              `yaml:"optional"`
	WriteOnly                bool                              `yaml:"write_only"`
	Sensitive                bool                              `yaml:"sensitive"`
	ExcludeTest              bool                              `yaml:"exclude_test"`
	ExcludeTestFrom          string                            // computed during merge: delta version whose exclude_test applies from that version up (empty when not version-scoped)
	ExcludeExample           bool                              `yaml:"exclude_example"`
	IncludeExample           bool                              `yaml:"include_example"`
	Description              string                            `yaml:"description"`
	Example                  string                            `yaml:"example"`
	VersionExamples          map[string]string                 // computed during merge: version → example value (nil if same across all versions)
	EnumValues               []string                          `yaml:"enum_values"`
	lastEnumValues           []string                          // unexported: last delta's own explicit list, decoupled from the ever-growing union above. Never read by templates.
	MinInt                   int64                             `yaml:"min_int"`
	MaxInt                   int64                             `yaml:"max_int"`
	lastRangeMin             int64                             // unexported: last delta's own explicit value, decoupled from the frozen schema scalar above. Never read by templates, never part of any "_base" map — pure merge-time bookkeeping.
	lastRangeMax             int64                             // same, for the max side.
	StringPatterns           []string                          `yaml:"string_patterns"`
	StringMinLength          int64                             `yaml:"string_min_length"`
	StringMaxLength          int64                             `yaml:"string_max_length"`
	lastStringMinLength      int64                             // unexported: last delta's own explicit value, decoupled from the widened schema scalar above. Never read by templates, never part of any "_base" map — pure merge-time bookkeeping.
	lastStringMaxLength      int64                             // same, for the max side.
	DefaultValue             string                            `yaml:"default_value"`
	RequiresReplace          bool                              `yaml:"requires_replace"`
	NoAugmentConfig          bool                              `yaml:"no_augment_config"`
	DeleteParent             bool                              `yaml:"delete_parent"`
	DeleteGrandparent        bool                              `yaml:"delete_grandparent"`
	NoDelete                 bool                              `yaml:"no_delete"`
	TestTags                 []string                          `yaml:"test_tags"`
	VersionTestTags          map[string][]string               // computed during merge: version → tag set (nil if same across all versions)
	MinimumTestValue         string                            `yaml:"minimum_test_value"`
	VersionMinimumTestValues map[string]string                 // computed during merge: version → minimum test value (nil if same across all versions)
	AddedInVersion           string                            // Which version introduced this attribute (e.g., "25.4", "26.2") — dot-separated major.minor, matches gen/definitions/ subdirectory names
	RemovedInVersion         string                            // Which version removed this attribute (populated when legacy: true)
	Legacy                   bool                              `yaml:"legacy"`           // If true, this attribute is removed/dropped in this version
	YangTypeChange           bool                              `yaml:"yang_type_change"` // If true, treat as a new attribute addition despite matching yang_name (incompatible type change)
	VersionRanges            map[string]RangeConstraint        // Version-specific ranges for Int64 fields (nil if same across all versions)
	VersionEnums             map[string][]string               // Version-specific enum sets for String fields (nil if same across all versions)
	VersionStringLengths     map[string]StringLengthConstraint // Version-specific string length constraints (nil if same across all versions)
	VersionPatterns          map[string][]string               // Version-specific string patterns (nil if same across all versions)
	ReplacesYangName         string                            `yaml:"replaces_yang_name"`
	ReplacesXPath            string                            // preserved from base XPath before it is cleared
	VersionYangNames         map[string]string                 // computed during merge: version → yang_name
	MovedInVersion           string                            // earliest version with new path (derived in fixAttributeBaseVersion)
	VersionDeleteMode        map[string]string                 // Version-specific delete mode: "" (direct), "parent", "grandparent". Nil if mode is uniform across all versions.
	Attributes               []YamlConfigAttribute             `yaml:"attributes"`
}

// RangeConstraint represents min/max constraints for a version
type RangeConstraint struct {
	Min int64
	Max int64
}

// StringLengthConstraint represents string length min/max for a version
type StringLengthConstraint struct {
	Min int64
	Max int64
}

type YamlTest struct {
	Path         string              `yaml:"path"`
	NoDelete     bool                `yaml:"no_delete"`
	Attributes   []YamlTestAttribute `yaml:"attributes"`
	Lists        []YamlTestList      `yaml:"lists"`
	Dependencies []string            `yaml:"dependencies"`
}

type YamlTestAttribute struct {
	Name      string `yaml:"name"`
	Value     string `yaml:"value"`
	Reference string `yaml:"reference"`
}

type YamlTestList struct {
	Name   string             `yaml:"name"`
	Key    string             `yaml:"key"`
	Items  []YamlTestListItem `yaml:"items"`
	Values []string           `yaml:"values"`
}

type YamlTestListItem struct {
	Attributes []YamlTestAttribute `yaml:"attributes"`
}

// Templating helper function to get short YANG name without prefix (xxx:abc -> abc)
func ToYangShortName(s string) string {
	elements := strings.Split(s, "/")
	for i := range elements {
		if strings.Contains(elements[i], ":") {
			elements[i] = strings.Split(elements[i], ":")[1]
		}
	}
	return strings.Join(elements, "/")
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		if strings.Contains(value, ":") {
			value = strings.Split(value, ":")[1]
		}
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert YANG name to GO name
func ToJsonPath(yangPath, xPath string) string {
	path := yangPath
	if xPath != "" {
		path = xPath
	}

	// Split by /, escape dots in each segment, then join with .
	parts := strings.Split(path, "/")
	for i, part := range parts {
		parts[i] = strings.ReplaceAll(part, ".", "\\.")
	}
	return strings.Join(parts, ".")
}

// sortedVersionKeys returns the keys of m sorted in ascending version order.
// Used to produce deterministic map literals in generated code.
func sortedVersionKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return versionCompare(keys[i], keys[j]) < 0
	})
	return keys
}

// getVersionValue is the generator's build-time counterpart to helpers.GetPathVersion
// (internal/provider/helpers/version_path.go): given a sparse map of version thresholds,
// it returns the value for the highest threshold ≤ target, or defaultValue if none
// qualify. Needed wherever a build-time loop must resolve a value for a version key that
// isn't necessarily an explicit key in that map itself.
func getVersionValue(target string, byVersion map[string]string, defaultValue string) string {
	best := ""
	result := defaultValue
	for threshold, val := range byVersion {
		if versionCompare(target, threshold) >= 0 {
			if best == "" || versionCompare(threshold, best) > 0 {
				best = threshold
				result = val
			}
		}
	}
	return result
}

// JsonPathExpr returns a Go expression string for the gNMI JSON path of an attribute.
// For static attributes it returns a quoted string literal (e.g. "files.file").
// For attributes with a renamed YANG path it returns a helpers.SelectYangPath(...) call
// with a map[string]string of version thresholds so any number of moves is handled.
func JsonPathExpr(attr YamlConfigAttribute, versionVar string) string {
	path := ToJsonPath(attr.YangName, attr.XPath)
	if len(attr.VersionYangNames) == 0 && len(attr.VersionXPath) == 0 {
		return fmt.Sprintf("%q", path)
	}
	defaultPath := ToJsonPath(attr.ReplacesYangName, attr.ReplacesXPath)
	// allVersions is the union of both independent per-version axes: a rename
	// (VersionYangNames) and an xpath-only change (VersionXPath) -- mirrors
	// GetDeletePathExpr's identical composition of VersionYangNames + VersionDeleteMode.
	allVersions := make(map[string]string)
	for v := range attr.VersionYangNames {
		allVersions[v] = ""
	}
	for v := range attr.VersionXPath {
		allVersions[v] = ""
	}
	var entries []string
	for _, v := range sortedVersionKeys(allVersions) {
		yangName := attr.YangName
		xpath := attr.XPath
		if name, ok := attr.VersionYangNames[v]; ok {
			yangName = name
			if yangName == attr.ReplacesYangName {
				xpath = attr.ReplacesXPath
			} else if yangName != attr.YangName {
				xpath = ""
			}
		}
		// The rename-driven xpath (or lack thereof) is the default; an explicit
		// VersionXPath entry for this version, or the nearest lower threshold, refines it.
		xpath = getVersionValue(v, attr.VersionXPath, xpath)
		entries = append(entries, fmt.Sprintf("%q: %q", v, ToJsonPath(yangName, xpath)))
	}
	return fmt.Sprintf("helpers.SelectYangPath(%s, map[string]string{%s}, %q)",
		versionVar, strings.Join(entries, ", "), defaultPath)
}

// KeyPathExpr returns a Go expression string for a list key's YANG path,
// version-conditional when the key has moved (replaces_yang_name set on an id:true attr).
// Uses XPath form (slash-separated, gNMI predicate-compatible) not dot-notation.
func KeyPathExpr(attr YamlConfigAttribute, versionVar string) string {
	path := GetXPath(attr.YangName, attr.XPath)
	if len(attr.VersionYangNames) == 0 && len(attr.VersionXPath) == 0 {
		return fmt.Sprintf("%q", path)
	}
	defaultPath := GetXPath(attr.ReplacesYangName, attr.ReplacesXPath)
	allVersions := make(map[string]string)
	for v := range attr.VersionYangNames {
		allVersions[v] = ""
	}
	for v := range attr.VersionXPath {
		allVersions[v] = ""
	}
	var entries []string
	for _, v := range sortedVersionKeys(allVersions) {
		yangName := attr.YangName
		xpath := attr.XPath
		if name, ok := attr.VersionYangNames[v]; ok {
			yangName = name
			if yangName == attr.ReplacesYangName {
				xpath = attr.ReplacesXPath
			} else if yangName != attr.YangName {
				xpath = ""
			}
		}
		xpath = getVersionValue(v, attr.VersionXPath, xpath)
		entries = append(entries, fmt.Sprintf("%q: %q", v, GetXPath(yangName, xpath)))
	}
	return fmt.Sprintf("helpers.SelectYangPath(%s, map[string]string{%s}, %q)",
		versionVar, strings.Join(entries, ", "), defaultPath)
}

// TypeYangBoolExpr returns a Go expression that evaluates to the correct TypeYangBool
// string ("empty", "presence", or "boolean") for a given providerVersion at runtime.
// When VersionTypeYangBool is nil, returns a quoted constant (the static TypeYangBool value).
// When populated, returns a helpers.GetPathVersion(...) call with sorted version keys,
// matching the deterministic-ordering pattern used by JsonPathExpr/KeyPathExpr.
func TypeYangBoolExpr(attr YamlConfigAttribute, versionVar string) string {
	if len(attr.VersionTypeYangBool) == 0 {
		return fmt.Sprintf("%q", attr.TypeYangBool)
	}
	var entries []string
	for _, v := range sortedVersionKeys(attr.VersionTypeYangBool) {
		entries = append(entries, fmt.Sprintf("%q: %q", v, attr.VersionTypeYangBool[v]))
	}
	return fmt.Sprintf("helpers.GetPathVersion(%s, %q, map[string]string{%s})",
		versionVar, attr.TypeYangBool, strings.Join(entries, ", "))
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id || attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to return true if reference included in attributes
func HasReference(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to return number of import parts
func ImportParts(attributes []YamlConfigAttribute) int {
	parts := 0
	for _, attr := range attributes {
		if attr.Reference {
			parts += 1
		} else if attr.Id {
			parts += 1
		}
	}
	return parts
}

// Templating helper function to return import attributes
func ImportAttributes(config YamlConfig) []YamlConfigAttribute {
	attributes := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.Reference || attr.Id {
			attributes = append(attributes, attr)
		}
	}
	return attributes
}

// Templating helper function to get ID attributes from a list
func GetIdAttributes(attributes []YamlConfigAttribute) []YamlConfigAttribute {
	idAttrs := []YamlConfigAttribute{}
	for _, attr := range attributes {
		if attr.Id {
			idAttrs = append(idAttrs, attr)
		}
	}
	return idAttrs
}

// Templating helper function to get xpath if available
func GetXPath(yangPath, xPath string) string {
	if xPath != "" {
		return xPath
	}
	return yangPath
}

func GetDeletePath(attribute YamlConfigAttribute) string {
	path := attribute.XPath
	if attribute.DeleteGrandparent {
		// Remove two levels: grandparent
		return helpers.RemoveLastPathElement(helpers.RemoveLastPathElement(path))
	}
	if attribute.DeleteParent {
		return helpers.RemoveLastPathElement(path)
	}
	return path
}

// GetDeletePathExpr returns a Go expression string for the delete path of attr,
// accounting for both VersionYangNames (renamed paths) and VersionDeleteMode
// (changed delete modes). Models KeyPathExpr at generator.go:303.
//
// Returns a quoted static path for the fast path, or
// helpers.SelectYangPath(versionVar, map[string]string{...}, "default") for
// attributes whose effective delete path differs across versions.
func GetDeletePathExpr(attr YamlConfigAttribute, versionVar string) string {
	// Fast path: no per-version variation on path or mode.
	if len(attr.VersionYangNames) == 0 && len(attr.VersionDeleteMode) == 0 {
		return fmt.Sprintf("%q", GetDeletePath(attr))
	}
	// VersionYangNames is set: the YANG path itself differs between versions, so we
	// must always emit a SelectYangPath call — fall through to the slow path below.

	// Slow path: compute the effective delete path per version.
	allVersions := make(map[string]string)
	for v := range attr.VersionYangNames {
		allVersions[v] = ""
	}
	for v := range attr.VersionDeleteMode {
		allVersions[v] = ""
	}

	baseYangName := attr.ReplacesYangName
	if baseYangName == "" {
		baseYangName = attr.YangName
	}
	// baseMode: derive from OLDEST key in VersionDeleteMode (after fixAttributeBaseVersion
	// replaced "_base"). DO NOT use attr.DeleteGrandparent/attr.DeleteParent — OR-merged.
	var baseMode string
	if len(attr.VersionDeleteMode) > 0 {
		oldestKey := sortedVersionKeys(attr.VersionDeleteMode)[0]
		baseMode = attr.VersionDeleteMode[oldestKey]
	} else if attr.DeleteGrandparent {
		baseMode = "grandparent"
	} else if attr.DeleteParent {
		baseMode = "parent"
	}
	defaultPath := applyDeleteMode(GetXPath(baseYangName, attr.ReplacesXPath), baseMode)

	var entries []string
	for _, v := range sortedVersionKeys(allVersions) {
		yn := attr.YangName
		xp := attr.XPath
		if name, ok := attr.VersionYangNames[v]; ok {
			yn = name
			// Mirror KeyPathExpr: thread correct XPath for choice/case attributes.
			if yn == attr.ReplacesYangName {
				xp = attr.ReplacesXPath
			} else if yn != attr.YangName {
				xp = ""
			}
		}
		mode := getVersionValue(v, attr.VersionDeleteMode, baseMode)
		entries = append(entries, fmt.Sprintf("%q: %q", v, applyDeleteMode(GetXPath(yn, xp), mode)))
	}
	return fmt.Sprintf("helpers.SelectYangPath(%s, map[string]string{%s}, %q)",
		versionVar, strings.Join(entries, ", "), defaultPath)
}

func applyDeleteMode(p, mode string) string {
	switch mode {
	case "grandparent":
		return helpers.RemoveLastPathElement(helpers.RemoveLastPathElement(p))
	case "parent":
		return helpers.RemoveLastPathElement(p)
	default:
		return p
	}
}

// HasVersionDeleteMode returns true if any attribute (recursively) triggers the slow
// path of GetDeletePathExpr — has a non-nil VersionDeleteMode OR has VersionYangNames
// with delete_parent/grandparent active.
func HasVersionDeleteMode(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionDeleteMode) > 0 {
			return true
		}
		if (attr.DeleteParent || attr.DeleteGrandparent) && len(attr.VersionYangNames) > 0 {
			return true
		}
		if HasVersionDeleteMode(attr.Attributes) {
			return true
		}
	}
	return false
}

// Templating helper function to add two integers
func Add(a, b int) int {
	return a + b
}

// Templating helper function to get example dn
func GetExamplePath(path string, attributes []YamlConfigAttribute) string {
	a := make([]interface{}, 0, len(attributes))
	for _, attr := range attributes {
		if attr.Id || attr.Reference {
			a = append(a, attr.Example)
		}
	}
	return fmt.Sprintf(path, a...)
}

// ReverseAttributes reverses a slice of YamlConfigAttribute
func ReverseAttributes(attributes []YamlConfigAttribute) []YamlConfigAttribute {
	reversed := make([]YamlConfigAttribute, len(attributes))
	for i, v := range attributes {
		reversed[len(attributes)-1-i] = v
	}
	return reversed
}

// IsXmlNamespaceSibling reports whether attr must be emitted as a separate
// XML sibling body to prevent xmldot from merging same-named elements that
// belong to different YANG namespaces.
//
// An attribute needs sibling treatment when:
//  1. Its yang_name carries an explicit namespace prefix anywhere in the path
//     (e.g. "traps/Cisco-IOS-XR-um-mpls-l3vpn-cfg:mpls/l3vpn/all"), AND
//  2. At least one other attribute in the same sibling list shares the same root element.
//
// The explicit yaml flag xml_namespace_sibling:true is still honoured for
// backward compatibility, but is no longer required.
func IsXmlNamespaceSibling(attr YamlConfigAttribute, allAttrs []YamlConfigAttribute) bool {
	yangNames := make([]string, len(allAttrs))
	for i, a := range allAttrs {
		yangNames[i] = a.YangName
	}
	return helpers.IsXmlNamespaceSiblingStr(attr.YangName, yangNames)
}

// HasXmlNamespaceSiblings returns true when any attribute in the list needs
// separate XML sibling treatment (either via explicit flag or auto-detection).
func HasXmlNamespaceSiblings(attributes []YamlConfigAttribute) bool {
	yangNames := make([]string, len(attributes))
	for i, a := range attributes {
		yangNames[i] = a.YangName
	}
	return helpers.HasXmlNamespaceSiblingsStr(yangNames)
}

// SiblingGroup holds a set of attributes that share the same top-level XML
// namespace element (e.g. all "Cisco-IOS-XR-um-mpls-te-cfg:fast-reroute/…"
// attributes belong to the same group).  They must be emitted into a single
// nsBody so they produce one merged XML element rather than duplicate siblings.
type SiblingGroup struct {
	// Key is "<namespace-prefix>:<top-element-name>", e.g.
	// "Cisco-IOS-XR-um-mpls-te-cfg:fast-reroute"
	Key        string
	Attributes []YamlConfigAttribute
}

// GroupXmlSiblings returns the xml-namespace-sibling attributes from allAttrs
// collected into groups that share the same top-level element.  The order of
// groups follows the first occurrence of each key in allAttrs.
func GroupXmlSiblings(allAttrs []YamlConfigAttribute) []SiblingGroup {
	helperGroups := helpers.GroupXmlSiblingsStr(allAttrs, func(attr YamlConfigAttribute) string {
		return attr.YangName
	})

	// Convert helper.SiblingGroup[YamlConfigAttribute] to generator.SiblingGroup
	groups := make([]SiblingGroup, len(helperGroups))
	for i, hg := range helperGroups {
		groups[i] = SiblingGroup{
			Key:        hg.Key,
			Attributes: hg.Items,
		}
	}
	return groups
}

// Templating helper function to generate version suffix for type names.
// Dotted format "24.4" → "V24_4" (dot replaced with underscore for valid Go identifier).
func VersionSuffix(version string) string {
	if version == "" {
		return ""
	}
	return "V" + strings.ReplaceAll(version, ".", "_")
}

// CollectVersionConstraints recursively collects all attributes that have version constraints
// It filters out attributes where AddedInVersion equals the baseVersion since those are the default
func CollectVersionConstraints(attributes []YamlConfigAttribute, prefix string, baseVersion string) []YamlConfigAttribute {
	var result []YamlConfigAttribute
	for _, attr := range attributes {
		// Build the field path
		fieldPath := attr.TfName
		if prefix != "" {
			fieldPath = prefix + "." + attr.TfName
		}

		// Only include if version constraint exists AND it's not the base version
		shouldInclude := false
		if attr.AddedInVersion != "" && attr.AddedInVersion != baseVersion {
			shouldInclude = true
		}
		if attr.RemovedInVersion != "" {
			shouldInclude = true
		}

		if shouldInclude {
			constraintAttr := attr
			constraintAttr.TfName = fieldPath // Store the full path
			result = append(result, constraintAttr)
		}

		// Recursively check nested attributes
		if len(attr.Attributes) > 0 {
			nested := CollectVersionConstraints(attr.Attributes, fieldPath, baseVersion)
			result = append(result, nested...)
		}
	}
	return result
}

// HasVersionConstraints returns true if any attribute has version constraints
func HasVersionConstraints(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.AddedInVersion != "" || attr.RemovedInVersion != "" {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionConstraints(attr.Attributes) {
			return true
		}
	}
	return false
}

// hasVersionDifferences checks if a config has any version-specific differences
func hasVersionDifferences(config YamlConfig) bool {
	// Check if entire resource has version constraints
	if config.RemovedInVersion != "" {
		return true
	}

	// A resource that only exists from a higher version is blocked on older devices
	if config.IntroducedInVersion != "" {
		return true
	}

	// Check if any attribute has version constraints or range differences
	return hasAttributeVersionDifferences(config.Attributes)
}

// hasAttributeVersionDifferences recursively checks if any attribute has version-specific changes
func hasAttributeVersionDifferences(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.AddedInVersion != "" || attr.RemovedInVersion != "" {
			return true
		}
		if attr.VersionRanges != nil && len(attr.VersionRanges) > 0 {
			return true
		}
		if len(attr.VersionYangNames) > 0 {
			return true
		}
		// These three feed GetEnumConstraints/GetStringLengthConstraints/GetPatternConstraints
		// exactly like VersionRanges feeds GetRangeConstraints -- omitting them here was BUG-11.
		if len(attr.VersionEnums) > 0 {
			return true
		}
		if len(attr.VersionStringLengths) > 0 {
			return true
		}
		if len(attr.VersionPatterns) > 0 {
			return true
		}
		// Defensive only, not currently load-bearing: VersionDeleteMode can only ever be
		// populated together with VersionYangNames (both require replaces_yang_name on the same
		// attribute, see mergeAttributes), so this never independently trips true today. Kept
		// explicit rather than relying on that coupling implicitly.
		if len(attr.VersionDeleteMode) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && hasAttributeVersionDifferences(attr.Attributes) {
			return true
		}
	}
	return false
}

// VersionChangesData is the JSON structure written to gen/version_changes_data.json.
// It maps snake_case resource name → its added/removed attribute lists.
type VersionChangesData map[string]ResourceVersionChanges

type ResourceVersionChanges struct {
	Removed []VersionedAttrRow `json:"removed"`
}

func writeVersionChangesData(configs []YamlConfig) {
	data := make(VersionChangesData)
	for _, cfg := range configs {
		removed := CollectRemovedAttrs(cfg.Attributes, "")
		if len(removed) > 0 {
			data[SnakeCase(cfg.Name)] = ResourceVersionChanges{Removed: removed}
		}
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling version changes data: %v", err)
	}
	if err := os.WriteFile("./gen/version_changes_data.json", b, 0644); err != nil {
		log.Fatalf("Error writing version_changes_data.json: %v", err)
	}
}

// VersionedAttrRow holds one row of the version compatibility table in generated docs.
type VersionedAttrRow struct {
	TfName    string `json:"tf_name"`
	RemovedIn string `json:"removed_in,omitempty"`
}

// CollectRemovedAttrs recursively collects attributes that have RemovedInVersion set, sorted by name.
func CollectRemovedAttrs(attrs []YamlConfigAttribute, prefix string) []VersionedAttrRow {
	var rows []VersionedAttrRow
	for _, attr := range attrs {
		name := attr.TfName
		if prefix != "" {
			name = prefix + "." + attr.TfName
		}
		if attr.RemovedInVersion != "" {
			rows = append(rows, VersionedAttrRow{TfName: name, RemovedIn: attr.RemovedInVersion})
		}
		if len(attr.Attributes) > 0 {
			rows = append(rows, CollectRemovedAttrs(attr.Attributes, name)...)
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].TfName < rows[j].TfName })
	return rows
}

// FormatVersionDisplay formats a version string for display.
// New dotted format ("25.2", "24.4") is returned as-is.
// Legacy 4-digit compact format ("2512") is converted to "MM.mm" (patch dropped).
func FormatVersionDisplay(version string) string {
	if strings.Contains(version, ".") {
		return version // already dotted major.minor
	}
	if len(version) == 4 {
		if _, err := strconv.Atoi(version); err == nil {
			return version[0:2] + "." + string(version[2])
		}
	}
	return version
}

// FormatVersionRanges formats version-specific ranges for markdown description
func FormatVersionRanges(versionRanges map[string]RangeConstraint) string {
	if len(versionRanges) == 0 {
		return ""
	}

	// Sort versions for consistent output
	versions := make([]string, 0, len(versionRanges))
	for v := range versionRanges {
		versions = append(versions, v)
	}
	sort.Strings(versions)

	parts := make([]string, 0, len(versions))
	for _, v := range versions {
		r := versionRanges[v]
		parts = append(parts, fmt.Sprintf("`%d`-`%d` (v%s)", r.Min, r.Max, FormatVersionDisplay(v)))
	}

	return strings.Join(parts, ", ")
}

// HasVersionRanges returns true if any attribute has version-specific ranges
func HasVersionRanges(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionRanges) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionRanges(attr.Attributes) {
			return true
		}
	}
	return false
}

// RangeConstraintInfo represents a field with version-specific range constraints
type RangeConstraintInfo struct {
	FieldPath     string
	VersionRanges map[string]RangeConstraint
}

// CollectVersionRangeConstraints recursively collects all attributes that have version-specific ranges
func CollectVersionRangeConstraints(attributes []YamlConfigAttribute, prefix string, baseVersion string) []RangeConstraintInfo {
	var result []RangeConstraintInfo
	for _, attr := range attributes {
		// Build the field path
		fieldPath := attr.TfName
		if prefix != "" {
			fieldPath = prefix + "." + attr.TfName
		}

		// Only include if version-specific ranges exist
		if len(attr.VersionRanges) > 0 {
			result = append(result, RangeConstraintInfo{
				FieldPath:     fieldPath,
				VersionRanges: attr.VersionRanges,
			})
		}

		// Recursively check nested attributes
		if len(attr.Attributes) > 0 {
			nested := CollectVersionRangeConstraints(attr.Attributes, fieldPath, baseVersion)
			result = append(result, nested...)
		}
	}
	return result
}

// stringSlicesEqual returns true if both slices contain the same elements in the same order.
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// unionStringSlices returns a deduplicated union of a and b, preserving order (a first, then b extras).
func unionStringSlices(a, b []string) []string {
	seen := make(map[string]bool, len(a))
	result := make([]string, 0, len(a)+len(b))
	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	for _, v := range b {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// FormatVersionEnums formats per-version enum sets for markdown description.
func FormatVersionEnums(versionEnums map[string][]string) string {
	if len(versionEnums) == 0 {
		return ""
	}
	versions := make([]string, 0, len(versionEnums))
	for v := range versionEnums {
		versions = append(versions, v)
	}
	sort.Strings(versions)
	parts := make([]string, 0, len(versions))
	for _, v := range versions {
		vals := versionEnums[v]
		quoted := make([]string, len(vals))
		for i, val := range vals {
			quoted[i] = fmt.Sprintf("`%s`", val)
		}
		parts = append(parts, fmt.Sprintf("%s (v%s)", strings.Join(quoted, ", "), FormatVersionDisplay(v)))
	}
	return strings.Join(parts, ", ")
}

// HasVersionEnums returns true if any attribute in the slice has version-specific enum sets.
func HasVersionEnums(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionEnums) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionEnums(attr.Attributes) {
			return true
		}
	}
	return false
}

// FormatVersionExamples returns sorted "ver": "ex", entries for inline map literals in generated test code.
// The trailing comma is required when the closing "}" is on the next line (Go syntax rule).
// Registered under both "formatVersionExamples" and "formatVersionMinimumTestValues" — identical signature.
func FormatVersionExamples(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%q: %q", k, m[k]))
	}
	return strings.Join(parts, ", ") + ","
}

// TestVersionGuardExpr returns the Go condition that gates an attribute's test checks and config
// on IOSXR_VERSION: added in, removed in, and excluded from versions. Empty when none apply.
func TestVersionGuardExpr(attr YamlConfigAttribute) string {
	var conds []string
	if attr.AddedInVersion != "" {
		conds = append(conds, fmt.Sprintf(`iosxrVersionAtLeast(os.Getenv("IOSXR_VERSION"), %q)`, attr.AddedInVersion))
	}
	for _, v := range []string{attr.RemovedInVersion, attr.ExcludeTestFrom} {
		if v != "" {
			conds = append(conds, fmt.Sprintf(`!iosxrVersionAtLeast(os.Getenv("IOSXR_VERSION"), %q)`, v))
		}
	}
	return strings.Join(conds, " && ")
}

// HasVersionExamples returns true if any attribute (recursively) has version-specific examples.
func HasVersionExamples(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionExamples) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionExamples(attr.Attributes) {
			return true
		}
	}
	return false
}

// HasVersionMinimumTestValues returns true if any attribute (recursively) has version-specific minimum test values.
func HasVersionMinimumTestValues(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionMinimumTestValues) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionMinimumTestValues(attr.Attributes) {
			return true
		}
	}
	return false
}

// FormatVersionTestTags returns sorted "ver": []string{...}, entries for inline map literals in generated test code.
func FormatVersionTestTags(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		quotedTags := make([]string, len(m[k]))
		for i, t := range m[k] {
			quotedTags[i] = fmt.Sprintf("%q", t)
		}
		parts = append(parts, fmt.Sprintf("%q: []string{%s}", k, strings.Join(quotedTags, ", ")))
	}
	return strings.Join(parts, ", ") + ","
}

// FormatVersionTestPrerequisites returns sorted "ver": constName, entries for use in the
// selectVersionPrerequisitesConfig map literal in generated test code.
// prefix is the constant name prefix ("testAccIosxr" for resource tests,
// "testAccDataSourceIosxr" for data source tests); camelName is the CamelCase resource
// name (e.g. "Logging"), used to build constant names.
func FormatVersionTestPrerequisites(m map[string][]YamlTest, prefix, camelName string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, len(keys))
	for i, k := range keys {
		parts[i] = fmt.Sprintf("%q: %s%sPrerequisitesConfig_%s", k, prefix, camelName, VersionSuffix(k))
	}
	return strings.Join(parts, ",\n\t\t\t") + ","
}

// HasVersionTestTags returns true if any attribute (recursively) has version-specific test tags.
func HasVersionTestTags(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionTestTags) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionTestTags(attr.Attributes) {
			return true
		}
	}
	return false
}

// EnumConstraintInfo represents a field with version-specific enum constraints.
type EnumConstraintInfo struct {
	FieldPath    string
	VersionEnums map[string][]string
}

// CollectVersionEnumConstraints recursively collects all attributes that have version-specific enum sets.
func CollectVersionEnumConstraints(attributes []YamlConfigAttribute, prefix string) []EnumConstraintInfo {
	var result []EnumConstraintInfo
	for _, attr := range attributes {
		fieldPath := attr.TfName
		if prefix != "" {
			fieldPath = prefix + "." + attr.TfName
		}
		if len(attr.VersionEnums) > 0 {
			result = append(result, EnumConstraintInfo{
				FieldPath:    fieldPath,
				VersionEnums: attr.VersionEnums,
			})
		}
		if len(attr.Attributes) > 0 {
			nested := CollectVersionEnumConstraints(attr.Attributes, fieldPath)
			result = append(result, nested...)
		}
	}
	return result
}

// GetWidestRange calculates the widest range (min of all mins, max of all maxs) from version ranges
// Returns a slice [min, max] for compatibility with Go templates
func GetWidestRange(versionRanges map[string]RangeConstraint) []int64 {
	if len(versionRanges) == 0 {
		return []int64{0, 0}
	}

	var minRange, maxRange int64
	first := true

	for _, r := range versionRanges {
		if first {
			minRange = r.Min
			maxRange = r.Max
			first = false
		} else {
			if r.Min < minRange {
				minRange = r.Min
			}
			if r.Max > maxRange {
				maxRange = r.Max
			}
		}
	}

	return []int64{minRange, maxRange}
}

// FormatVersionStringLengths formats version-specific string lengths for markdown description.
func FormatVersionStringLengths(versionStringLengths map[string]StringLengthConstraint) string {
	if len(versionStringLengths) == 0 {
		return ""
	}
	versions := make([]string, 0, len(versionStringLengths))
	for v := range versionStringLengths {
		versions = append(versions, v)
	}
	sort.Strings(versions)
	parts := make([]string, 0, len(versions))
	for _, v := range versions {
		c := versionStringLengths[v]
		parts = append(parts, fmt.Sprintf("`%d`-`%d` (v%s)", c.Min, c.Max, FormatVersionDisplay(v)))
	}
	return strings.Join(parts, ", ")
}

// HasVersionStringLengths returns true if any attribute has version-specific string length constraints.
func HasVersionStringLengths(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionStringLengths) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionStringLengths(attr.Attributes) {
			return true
		}
	}
	return false
}

// StringLengthConstraintInfo represents a field with version-specific string length constraints.
type StringLengthConstraintInfo struct {
	FieldPath            string
	VersionStringLengths map[string]StringLengthConstraint
}

// CollectVersionStringLengthConstraints recursively collects attributes with version-specific string lengths.
func CollectVersionStringLengthConstraints(attributes []YamlConfigAttribute, prefix string) []StringLengthConstraintInfo {
	var result []StringLengthConstraintInfo
	for _, attr := range attributes {
		fieldPath := attr.TfName
		if prefix != "" {
			fieldPath = prefix + "." + attr.TfName
		}
		if len(attr.VersionStringLengths) > 0 {
			result = append(result, StringLengthConstraintInfo{
				FieldPath:            fieldPath,
				VersionStringLengths: attr.VersionStringLengths,
			})
		}
		if len(attr.Attributes) > 0 {
			nested := CollectVersionStringLengthConstraints(attr.Attributes, fieldPath)
			result = append(result, nested...)
		}
	}
	return result
}

// HasVersionPatterns returns true if any attribute has version-specific string patterns.
func HasVersionPatterns(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if len(attr.VersionPatterns) > 0 {
			return true
		}
		if len(attr.Attributes) > 0 && HasVersionPatterns(attr.Attributes) {
			return true
		}
	}
	return false
}

// PatternConstraintInfo represents a field with version-specific string patterns.
type PatternConstraintInfo struct {
	FieldPath       string
	VersionPatterns map[string][]string
}

// CollectVersionPatternConstraints recursively collects attributes with version-specific patterns.
func CollectVersionPatternConstraints(attributes []YamlConfigAttribute, prefix string) []PatternConstraintInfo {
	var result []PatternConstraintInfo
	for _, attr := range attributes {
		fieldPath := attr.TfName
		if prefix != "" {
			fieldPath = prefix + "." + attr.TfName
		}
		if len(attr.VersionPatterns) > 0 {
			result = append(result, PatternConstraintInfo{
				FieldPath:       fieldPath,
				VersionPatterns: attr.VersionPatterns,
			})
		}
		if len(attr.Attributes) > 0 {
			nested := CollectVersionPatternConstraints(attr.Attributes, fieldPath)
			result = append(result, nested...)
		}
	}
	return result
}

// findNoAugmentConfigViolations returns a log.Fatalf-ready message for every case where acc
// (the config merged through the immediately-preceding version) has NoAugmentConfig=true and
// raw (this version's own, not-yet-merged delta) re-lists the same resource or attribute
// without also restating NoAugmentConfig=true. Per provider-yaml-authoring-rules.md Rule 4,
// the merged NoAugmentConfig flag is sticky (mergeConfigs/mergeAttributes only ever set it to
// true), but the gate that actually skips YANG augmentation (augmentForVersion) reads the raw,
// pre-merge delta for the version being processed -- so an author re-touching an already
// hand-modeled resource or attribute for any unrelated reason, without repeating
// no_augment_config: true, would silently get real YANG data written back into it.
//
// Pure and testable; validateNoAugmentConfigCarryover below is the log.Fatalf wrapper called
// from main().
func findNoAugmentConfigViolations(acc, raw YamlConfig, version string) []string {
	var violations []string
	if acc.NoAugmentConfig && !raw.NoAugmentConfig {
		violations = append(violations, fmt.Sprintf(
			"%s: has no_augment_config: true (resource-level) in an earlier version but the %s "+
				"delta does not restate no_augment_config: true -- this would run full YANG "+
				"augmentation against every attribute in this delta.",
			acc.Name, version,
		))
	}
	violations = append(violations, findAttributeNoAugmentConfigViolations(acc.Attributes, raw.Attributes, acc.Name, version)...)
	return violations
}

// findAttributeNoAugmentConfigViolations walks acc and raw in lockstep, matching attributes the
// same way mergeAttributes does (yang_name / tf_name / replaces_yang_name), and recurses into
// nested List attributes exactly the way mergeAttributes' own nested merge does -- a flat,
// global lookup would incorrectly conflate same-named attributes nested under different parents.
func findAttributeNoAugmentConfigViolations(acc, raw []YamlConfigAttribute, resourceName, version string) []string {
	var violations []string
	for _, rawAttr := range raw {
		for _, accAttr := range acc {
			matched := accAttr.YangName == rawAttr.YangName ||
				(accAttr.TfName != "" && rawAttr.TfName != "" && accAttr.TfName == rawAttr.TfName) ||
				(rawAttr.ReplacesYangName != "" && accAttr.YangName == rawAttr.ReplacesYangName)
			if !matched {
				continue
			}
			if accAttr.NoAugmentConfig && !rawAttr.NoAugmentConfig {
				violations = append(violations, fmt.Sprintf(
					"attribute %q (%s): has no_augment_config: true in an earlier version but "+
						"the %s delta re-lists it without restating no_augment_config: true -- "+
						"this would silently overwrite hand-authored type/range/enum values with "+
						"real YANG data. Add no_augment_config: true to this attribute in the %s delta.",
					rawAttr.YangName, resourceName, version, version,
				))
			}
			if len(rawAttr.Attributes) > 0 && len(accAttr.Attributes) > 0 {
				violations = append(violations, findAttributeNoAugmentConfigViolations(accAttr.Attributes, rawAttr.Attributes, resourceName, version)...)
			}
			break
		}
	}
	return violations
}

// validateNoAugmentConfigCarryover fails the build on the first no_augment_config carryover
// violation found. Must be called with acc/raw BEFORE augmentForVersion runs on raw for this
// version -- augmentForVersion's YANG augmentation mutates raw's attributes in place (Go slice
// aliasing with the versionConfigs map entry), so by the time it returns, the hand-authored data
// this check exists to protect may already be gone.
func validateNoAugmentConfigCarryover(acc, raw YamlConfig, version string) {
	for _, msg := range findNoAugmentConfigViolations(acc, raw, version) {
		log.Fatalf("%s", msg)
	}
}

// hclReserved contains HCL2 keywords that cause parse errors when unquoted
// as the first key in an object literal (e.g. "for" triggers a for-expression).
var hclReserved = map[string]bool{"for": true, "if": true, "in": true}

// SortedAttrs returns a copy of attrs sorted by TfName. Used in .tf templates only
// so that example HCL is alphabetized without affecting Go schema or test generation.
// HCL reserved keywords are sorted to the end to avoid parse errors when they appear
// as the first key in an object literal (e.g. "for = ..." triggers a for-expression).
func SortedAttrs(attrs []YamlConfigAttribute) []YamlConfigAttribute {
	sorted := make([]YamlConfigAttribute, len(attrs))
	copy(sorted, attrs)
	sort.Slice(sorted, func(i, j int) bool {
		iReserved := hclReserved[sorted[i].TfName]
		jReserved := hclReserved[sorted[j].TfName]
		if iReserved != jReserved {
			return !iReserved // reserved sorts after non-reserved
		}
		return sorted[i].TfName < sorted[j].TfName
	})
	return sorted
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":                              ToGoName,
	"camelCase":                             CamelCase,
	"snakeCase":                             SnakeCase,
	"add":                                   Add,
	"isLast":                                helpers.IsLast,
	"sprintf":                               fmt.Sprintf,
	"removeLastPathElement":                 helpers.RemoveLastPathElement,
	"getXPath":                              GetXPath,
	"getDeletePath":                         GetDeletePath,
	"reverseAttributes":                     ReverseAttributes,
	"hasId":                                 HasId,
	"hasReference":                          HasReference,
	"importParts":                           ImportParts,
	"importAttributes":                      ImportAttributes,
	"getExamplePath":                        GetExamplePath,
	"toJsonPath":                            ToJsonPath,
	"jsonPathExpr":                          JsonPathExpr,
	"keyPathExpr":                           KeyPathExpr,
	"typeYangBoolExpr":                      TypeYangBoolExpr,
	"versionSuffix":                         VersionSuffix,
	"getDeletePathExpr":                     GetDeletePathExpr,
	"collectVersionConstraints":             CollectVersionConstraints,
	"hasVersionConstraints":                 HasVersionConstraints,
	"formatVersionRanges":                   FormatVersionRanges,
	"formatVersionDisplay":                  FormatVersionDisplay,
	"hasVersionRanges":                      HasVersionRanges,
	"collectVersionRangeConstraints":        CollectVersionRangeConstraints,
	"getWidestRange":                        GetWidestRange,
	"formatVersionEnums":                    FormatVersionEnums,
	"hasVersionEnums":                       HasVersionEnums,
	"formatVersionExamples":                 FormatVersionExamples,
	"testVersionGuardExpr":                  TestVersionGuardExpr,
	"formatVersionMinimumTestValues":        FormatVersionExamples,
	"hasVersionExamples":                    HasVersionExamples,
	"hasVersionMinimumTestValues":           HasVersionMinimumTestValues,
	"formatVersionTestTags":                 FormatVersionTestTags,
	"formatVersionTestPrerequisites":        FormatVersionTestPrerequisites,
	"hasVersionTestTags":                    HasVersionTestTags,
	"collectVersionEnumConstraints":         CollectVersionEnumConstraints,
	"formatVersionStringLengths":            FormatVersionStringLengths,
	"hasVersionStringLengths":               HasVersionStringLengths,
	"collectVersionStringLengthConstraints": CollectVersionStringLengthConstraints,
	"hasVersionPatterns":                    HasVersionPatterns,
	"collectVersionPatternConstraints":      CollectVersionPatternConstraints,
	"collectRemovedAttrs":                   CollectRemovedAttrs,
	"sortedAttrs":                           SortedAttrs,
	"getIdAttributes":                       GetIdAttributes,
	"getLastPathElement":                    helpers.GetLastPathElement,
	"toDotPath":                             helpers.ToDotPath,
	"hasPrefix":                             strings.HasPrefix,
	"hasXmlNamespaceSiblings":               HasXmlNamespaceSiblings,
	"isXmlNamespaceSibling":                 IsXmlNamespaceSibling,
	"xmlNamespacePrefixFromXPath":           helpers.XmlNamespacePrefixFromXPath,
	"groupXmlSiblings":                      GroupXmlSiblings,
}

func resolvePath(e *yang.Entry, path string) *yang.Entry {
	pathElements := strings.Split(path, "/")

	for _, pathElement := range pathElements {
		if len(pathElement) > 0 {
			// remove XPath predicate (e.g., [name=value] or [name=%v])
			if strings.Contains(pathElement, "[") {
				pathElement = pathElement[:strings.Index(pathElement, "[")]
			}
			// remove namespace prefix (e.g., Cisco-IOS-XE-bgp:bgp -> bgp)
			if strings.Contains(pathElement, ":") {
				pathElement = pathElement[strings.Index(pathElement, ":")+1:]
			}
			if _, ok := e.Dir[pathElement]; !ok {
				panic(fmt.Sprintf("Failed to resolve YANG path: %s, element: %s", path, pathElement))
			}
			e = e.Dir[pathElement]
		}
	}

	return e
}

// safeResolvePath resolves a YANG path without panicking.
// Returns (entry, true) on success, or (nil, false) if any path element is missing.
// Used for attributes that may not exist in the target YANG version (e.g. removed nodes).
func safeResolvePath(e *yang.Entry, path string) (*yang.Entry, bool) {
	for _, pathElement := range strings.Split(path, "/") {
		if pathElement == "" {
			continue
		}
		if strings.Contains(pathElement, "[") {
			pathElement = pathElement[:strings.Index(pathElement, "[")]
		}
		if strings.Contains(pathElement, ":") {
			pathElement = pathElement[strings.Index(pathElement, ":")+1:]
		}
		next, ok := e.Dir[pathElement]
		if !ok {
			return nil, false
		}
		e = next
	}
	return e, true
}

func addKeys(e *yang.Entry, config *YamlConfig) {
	first := true
	for {
		if e.Key != "" {
			keys := strings.Split(e.Key, " ")
			for _, key := range keys {
				var keyAttr *YamlConfigAttribute
				// check if key attribute already in config
				for i := range config.Attributes {
					if config.Attributes[i].YangScope != "" && config.Attributes[i].YangScope != e.Name {
						continue
					}
					if config.Attributes[i].YangName == key {
						keyAttr = &config.Attributes[i]
						break
					}
				}
				if keyAttr == nil {
					continue
				}
				if first {
					keyAttr.Id = true
					keyAttr.Reference = false
				} else {
					keyAttr.Id = false
					keyAttr.Reference = true
				}
				parseAttribute(e, keyAttr)
			}
		}
		first = false
		if e.Parent != nil {
			e = e.Parent
			continue
		}
		break
	}
}

// rangeBounds returns the outer bounds of a possibly multi-part YANG range. YANG requires ascending,
// non-overlapping segments, so the first segment holds the minimum and the last holds the maximum.
func rangeBounds(r yang.YangRange) (lo, hi int64) {
	first, last := r[0], r[len(r)-1]
	lo = int64(first.Min.Value)
	if first.Min.Negative {
		lo = -lo
	}
	maxV := last.Max.Value
	// hack to not introduce unsigned types
	if maxV > math.MaxInt64 {
		maxV = math.MaxInt64
	}
	return lo, int64(maxV)
}

func parseAttribute(e *yang.Entry, attr *YamlConfigAttribute) {
	leaf := resolvePath(e, attr.YangName)
	//fmt.Printf("%s, Entry: %+v\n\n", attr.YangName, e)
	//fmt.Printf("%s, Kind: %+v, Type: %+v\n\n", leaf.Name, leaf.Kind, leaf.Type)
	if leaf.Kind.String() == "Leaf" {
		if leaf.ListAttr != nil {
			if helpers.Contains([]string{"string", "union", "leafref"}, leaf.Type.Kind.String()) {
				attr.Type = "StringList"
			} else if helpers.Contains([]string{"int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
				attr.Type = "Int64List"
			} else {
				panic(fmt.Sprintf("Unknown leaf-list type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
			}
			// TODO parse union type
		} else if helpers.Contains([]string{"string", "union", "leafref"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			if leaf.Type.Length != nil {
				if attr.StringMinLength == 0 {
					attr.StringMinLength = int64(leaf.Type.Length[0].Min.Value)
				}
				if attr.StringMaxLength == 0 {
					max := leaf.Type.Length[0].Max.Value
					// hack to not introduce unsigned types
					if max > math.MaxInt64 {
						max = math.MaxInt64
					}
					attr.StringMaxLength = int64(max)
				}
			}
			if len(leaf.Type.Pattern) > 0 && len(attr.StringPatterns) == 0 {
				attr.StringPatterns = leaf.Type.Pattern
			}
		} else if helpers.Contains([]string{"int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
			attr.Type = "Int64"
			if leaf.Type.Range != nil {
				minV, maxV := rangeBounds(leaf.Type.Range)
				if attr.MinInt == 0 {
					attr.MinInt = minV
				}
				if attr.MaxInt == 0 {
					attr.MaxInt = maxV
				}
			}
		} else if helpers.Contains([]string{"boolean", "empty"}, leaf.Type.Kind.String()) {
			if attr.TypeYangBool == "" {
				if leaf.Type.Kind.String() == "boolean" {
					attr.TypeYangBool = "boolean"
				} else if leaf.Type.Kind.String() == "empty" {
					attr.TypeYangBool = "empty"
				}
			}
			attr.Type = "Bool"
		} else if helpers.Contains([]string{"enumeration"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			attr.EnumValues = leaf.Type.Enum.Names()
		} else {
			panic(fmt.Sprintf("Unknown leaf type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
		}
	}
	if _, ok := leaf.Extra["presence"]; ok {
		if attr.TypeYangBool == "" {
			attr.TypeYangBool = "presence"
		}
		attr.Type = "Bool"
	}
	if attr.XPath == "" {
		attr.XPath = attr.YangName
	}
	if attr.TfName == "" {
		tfName := strings.ReplaceAll(helpers.ToYangShortName(attr.XPath), "-", "_")
		tfName = strings.ReplaceAll(tfName, "/", "_")
		// Trim leading underscores to comply with tfsdk naming rules (must start with letter)
		tfName = strings.TrimLeft(tfName, "_")
		attr.TfName = tfName
	}
	if attr.Description == "" {
		attr.Description = strings.ReplaceAll(leaf.Description, "\n", " ")
	}
	if !attr.Mandatory && attr.DefaultValue == "" && !attr.Optional {
		foundChoice := false
		parent := leaf.Parent
		for parent != nil {
			if parent.IsChoice() {
				foundChoice = true
				break
			}
			parent = parent.Parent
		}
		if !foundChoice {
			attr.Mandatory = leaf.Mandatory.Value()
		}
	}
}

// resolveYangPathForVersion returns the path_version entry in effect for version, taking the
// delta's own entries over those accumulated from earlier versions. Empty if none applies.
func resolveYangPathForVersion(version string, accumulated, own map[string]string) string {
	merged := make(map[string]string, len(accumulated)+len(own))
	for k, v := range accumulated {
		merged[k] = v
	}
	for k, v := range own {
		merged[k] = v
	}
	return getVersionValue(version, merged, "")
}

func augmentConfig(config *YamlConfig, modelPaths []string) {
	path := ""
	if config.AugmentPath != "" {
		path = config.AugmentPath
	} else {
		path = config.Path
	}
	path = strings.TrimPrefix(path, "/")
	module := strings.Split(path, ":")[0]
	e, errors := yang.GetModule(module, modelPaths...)
	if e == nil {
		log.Fatalf("definition %q @ %s: YANG module %q not found (path %q): %+v\n"+
			"Add the module to gen/load_models.go, or set legacy / no_augment_config.",
			config.Name, config.Version, module, path, errors)
	}
	if len(errors) > 0 {
		fmt.Printf("YANG parser error(s): %+v\n\n", errors)
	}

	// Print definition/model info
	fmt.Printf("Processing definition: %s\n", config.Name)
	//fmt.Printf("Resolving yang model: %s ==> Resolved: %s\n", module, e.Name)

	p := path[len(module)+1:]
	e = resolvePath(e, p)

	addKeys(e, config)

	for ia := range config.Attributes {
		// Default XPath from YangName if not explicitly set (do this first for all attributes)
		if config.Attributes[ia].XPath == "" {
			config.Attributes[ia].XPath = config.Attributes[ia].YangName
		}

		// For Lists with NoAugmentConfig, still process child attributes to set their XPath
		if config.Attributes[ia].Type == "List" && config.Attributes[ia].NoAugmentConfig {
			for iaa := range config.Attributes[ia].Attributes {
				// Default XPath from YangName if not explicitly set
				if config.Attributes[ia].Attributes[iaa].XPath == "" {
					config.Attributes[ia].Attributes[iaa].XPath = config.Attributes[ia].Attributes[iaa].YangName
				}
				// If parent list has no_augment_config and child is an id attribute, inherit the flag
				if config.Attributes[ia].Attributes[iaa].Id && !config.Attributes[ia].Attributes[iaa].NoAugmentConfig {
					config.Attributes[ia].Attributes[iaa].NoAugmentConfig = true
				}
				// For nested lists, also set XPath for their children
				if config.Attributes[ia].Attributes[iaa].Type == "List" && config.Attributes[ia].Attributes[iaa].NoAugmentConfig {
					for iaaa := range config.Attributes[ia].Attributes[iaa].Attributes {
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].XPath == "" {
							config.Attributes[ia].Attributes[iaa].Attributes[iaaa].XPath = config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName
						}
						// If parent list has no_augment_config and child is an id attribute, inherit the flag
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Id && !config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig {
							config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig = true
						}
					}
				}
			}
		}

		if config.Attributes[ia].Id || config.Attributes[ia].Reference || config.Attributes[ia].NoAugmentConfig {
			continue
		}
		// Use safeResolvePath to skip attributes whose YANG path no longer exists in
		// this version's model (e.g. removed/legacy nodes in a newer version diff).
		if _, ok := safeResolvePath(e, config.Attributes[ia].YangName); !ok {
			continue
		}
		parseAttribute(e, &config.Attributes[ia])
		if config.Attributes[ia].Type == "List" {
			el := resolvePath(e, config.Attributes[ia].YangName)
			for iaa := range config.Attributes[ia].Attributes {
				// Default XPath from YangName if not explicitly set (do this first for all attributes)
				if config.Attributes[ia].Attributes[iaa].XPath == "" {
					config.Attributes[ia].Attributes[iaa].XPath = config.Attributes[ia].Attributes[iaa].YangName
				}
				if config.Attributes[ia].Attributes[iaa].NoAugmentConfig {
					continue
				}
				if _, ok := safeResolvePath(el, config.Attributes[ia].Attributes[iaa].YangName); !ok {
					continue
				}
				parseAttribute(el, &config.Attributes[ia].Attributes[iaa])
				if config.Attributes[ia].Attributes[iaa].Type == "List" {
					ell := resolvePath(el, config.Attributes[ia].Attributes[iaa].YangName)
					for iaaa := range config.Attributes[ia].Attributes[iaa].Attributes {
						// Default XPath from YangName if not explicitly set (do this first for all attributes)
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].XPath == "" {
							config.Attributes[ia].Attributes[iaa].Attributes[iaaa].XPath = config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName
						}
						if _, ok := safeResolvePath(ell, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName); !ok {
							continue
						}
						// If parent list has no_augment_config and child is an id attribute, inherit the flag
						if config.Attributes[ia].Attributes[iaa].NoAugmentConfig && config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Id && !config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig {
							config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig = true
						}
						// Only skip parseAttribute if no_augment_config, but still process children
						if !config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig {
							parseAttribute(ell, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa])
						}
						// Process children if this is a List (check Type from YAML, or if children exist)
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Type == "List" || config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Type == "Set" || len(config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes) > 0 {
							elll := resolvePath(ell, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName)
							for iaaaa := range config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes {
								if _, ok := safeResolvePath(elll, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].YangName); !ok {
									continue
								}
								// Default XPath from YangName if not explicitly set (do this first for all attributes)
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].XPath == "" {
									config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].XPath = config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].YangName
								}
								// If parent list has no_augment_config and child is an id attribute, inherit the flag
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig && config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Id && !config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].NoAugmentConfig {
									config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].NoAugmentConfig = true
								}
								// Only skip parseAttribute if no_augment_config, but still process children
								if !config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].NoAugmentConfig {
									parseAttribute(elll, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa])
								}
								// Process level 5 children if this is a List
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Type == "List" || config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Type == "Set" || len(config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes) > 0 {
									ellll := resolvePath(elll, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].YangName)
									for iaaaaa := range config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes {
										if _, ok := safeResolvePath(ellll, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].YangName); !ok {
											continue
										}
										// Default XPath from YangName if not explicitly set
										if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].XPath == "" {
											config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].XPath = config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].YangName
										}
										// If parent list has no_augment_config and child is an id attribute, inherit the flag
										if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].NoAugmentConfig && config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].Id && !config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].NoAugmentConfig {
											config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].NoAugmentConfig = true
										}
										if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa].NoAugmentConfig {
											continue
										}
										parseAttribute(ellll, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Attributes[iaaaaa])
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s configuration.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage the %s configuration.", config.Name)
	}
}

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	// Increase buffer to 1MB to handle long lines in large templates/generated files.
	// The default 64KB limit silently truncates the scan, producing a misleading
	// "unexpected EOF" parse error at whatever template line was last read.
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading template %s: %v", templatePath, err)
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template for %s: %v", outputPath, err)
	}

	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		defer existingFile.Close()
		existingScanner := bufio.NewScanner(existingFile)
		existingScanner.Buffer(make([]byte, 1024*1024), 1024*1024)
		var newContent string
		currentSectionName := ""
		processedSections := make(map[string]bool)
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
					processedSections[currentSectionName] = true
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}

		output = bytes.NewBufferString(newContent)
	} else {
		existingFile.Close()
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer f.Close()
	f.Write(output.Bytes())
}

// versionCompare compares two version strings. Returns -1 if v1 < v2, 0 if equal, 1 if v1 > v2
func versionCompare(v1, v2 string) int {
	// Handle simple numeric versions like "2442"
	// Also handle dot versions like "25.2.2"
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	// Convert to comparable format
	for i := 0; i < len(parts1) || i < len(parts2); i++ {
		var p1, p2 int
		if i < len(parts1) {
			p1, _ = strconv.Atoi(parts1[i])
		}
		if i < len(parts2) {
			p2, _ = strconv.Atoi(parts2[i])
		}
		if p1 < p2 {
			return -1
		}
		if p1 > p2 {
			return 1
		}
	}
	return 0
}

// mergeAttributes merges two attribute slices, with newer attributes overriding older ones
// overrideVersion indicates which version the override attributes come from
func mergeAttributes(base, override []YamlConfigAttribute, overrideVersion string) []YamlConfigAttribute {
	result := make([]YamlConfigAttribute, len(base))
	copy(result, base)

	for _, newAttr := range override {
		found := false
		for i := range result {
			// yang_type_change guard: missing tf_name is a misconfiguration
			if result[i].YangName == newAttr.YangName && newAttr.YangTypeChange && newAttr.TfName == "" {
				log.Fatalf("attribute %q: yang_type_change: true requires a non-empty tf_name", newAttr.YangName)
			}
			// yang_type_change guard: non-unique tf_name would silently merge instead of appending — misconfiguration
			if result[i].YangName == newAttr.YangName && newAttr.YangTypeChange && newAttr.TfName != "" && result[i].TfName == newAttr.TfName {
				log.Fatalf("attribute %q: yang_type_change: true requires a unique tf_name (conflicts with existing %q)", newAttr.YangName, result[i].TfName)
			}
			// yang_type_change bypass: same YANG path but incompatible type → treat as new addition, not a merge
			if result[i].YangName == newAttr.YangName &&
				newAttr.YangTypeChange && newAttr.TfName != "" &&
				result[i].TfName != newAttr.TfName && !newAttr.Legacy {
				continue
			}
			// Match by yang_name, tf_name, or replaces_yang_name (for YANG path renames)
			if result[i].YangName == newAttr.YangName ||
				(result[i].TfName != "" && newAttr.TfName != "" && result[i].TfName == newAttr.TfName) ||
				(newAttr.ReplacesYangName != "" && result[i].YangName == newAttr.ReplacesYangName) {
				if newAttr.ReplacesYangName != "" && newAttr.Legacy {
					log.Fatalf("attribute %q: replaces_yang_name and legacy cannot both be set", newAttr.TfName)
				}
				// Legacy: true means this attribute is removed in the override version
				// Instead of dropping it, mark it with RemovedInVersion for validation
				if newAttr.Legacy {
					result[i].RemovedInVersion = overrideVersion
					result[i].Legacy = true
					if newAttr.TfName != "" {
						result[i].TfName = newAttr.TfName
					}
					found = true
					break
				}

				// Recursively merge nested attributes if both have them
				if len(result[i].Attributes) > 0 && len(newAttr.Attributes) > 0 {
					result[i].Attributes = mergeAttributes(result[i].Attributes, newAttr.Attributes, overrideVersion)
				} else if len(newAttr.Attributes) > 0 {
					result[i].Attributes = newAttr.Attributes
					// Mark nested attributes with version
					for j := range result[i].Attributes {
						if result[i].Attributes[j].AddedInVersion == "" {
							result[i].Attributes[j].AddedInVersion = overrideVersion
						}
					}
				}

				// Snapshot the original XPath before any overrides, for use as ReplacesXPath
				originalXPath := result[i].XPath

				// Override all other fields from new attribute
				if newAttr.YangScope != "" {
					result[i].YangScope = newAttr.YangScope
				}
				if newAttr.TfName != "" {
					result[i].TfName = newAttr.TfName
				}
				if newAttr.XPath != "" {
					// originalXPath (captured above, before this override) is always the true
					// immediately-preceding version's own value -- .XPath is a plain
					// last-version-wins scalar with no freezing/unioning, so unlike
					// min_int/enum_values/string_min_length it needs no separate tracking
					// field to avoid staleness. Record a version-gated entry whenever this
					// delta's xpath actually differs from that value, independent of whether
					// replaces_yang_name is also set on the same delta.
					if originalXPath != "" && newAttr.XPath != originalXPath {
						if result[i].VersionXPath == nil {
							result[i].VersionXPath = make(map[string]string)
							result[i].VersionXPath["_base"] = originalXPath
						}
						result[i].VersionXPath[overrideVersion] = newAttr.XPath
					}
					result[i].XPath = newAttr.XPath
				}
				if newAttr.Type != "" {
					result[i].Type = newAttr.Type
				}
				if newAttr.ReadRaw {
					result[i].ReadRaw = newAttr.ReadRaw
				}
				typeYangBoolBaseline := result[i].lastTypeYangBool
				if typeYangBoolBaseline == "" {
					// Never tracked yet -- fall back to the real base value (still accurate
					// at this point, before this fold's own assignment below runs).
					typeYangBoolBaseline = result[i].TypeYangBool
				}
				if newAttr.TypeYangBool != "" && newAttr.TypeYangBool != typeYangBoolBaseline {
					// Bool encoding changed in this version — version-gate it, keep base as default.
					// Compared against typeYangBoolBaseline (the true immediately-preceding
					// delta's own value), not result[i].TypeYangBool, which is deliberately
					// frozen below.
					if result[i].VersionTypeYangBool == nil {
						result[i].VersionTypeYangBool = make(map[string]string)
					}
					result[i].VersionTypeYangBool[overrideVersion] = newAttr.TypeYangBool
					// Do NOT overwrite result[i].TypeYangBool — base encoding remains the default
				} else if newAttr.TypeYangBool != "" {
					// Same value — no version map needed
					result[i].TypeYangBool = newAttr.TypeYangBool
				}
				if newAttr.TypeYangBool != "" {
					// Track this delta's own explicit value unconditionally, decoupled from the
					// deliberately-frozen default above, so the next fold's comparison is always
					// against the true immediately-preceding version.
					result[i].lastTypeYangBool = newAttr.TypeYangBool
				}
				if newAttr.Id {
					result[i].Id = newAttr.Id
				}
				if newAttr.Reference {
					result[i].Reference = newAttr.Reference
				}
				if newAttr.Mandatory {
					result[i].Mandatory = newAttr.Mandatory
				}
				if newAttr.Optional {
					result[i].Optional = newAttr.Optional
				}
				if newAttr.WriteOnly {
					result[i].WriteOnly = newAttr.WriteOnly
				}
				if newAttr.Sensitive {
					result[i].Sensitive = newAttr.Sensitive
				}
				// exclude_test in a delta applies from that version up; a base-level exclusion stays global.
				if newAttr.ExcludeTest && !result[i].ExcludeTest && result[i].ExcludeTestFrom == "" {
					result[i].ExcludeTestFrom = overrideVersion
				}
				if newAttr.ExcludeExample {
					result[i].ExcludeExample = newAttr.ExcludeExample
				}
				if newAttr.IncludeExample {
					result[i].IncludeExample = newAttr.IncludeExample
				}
				if newAttr.Description != "" {
					result[i].Description = newAttr.Description
				}
				if newAttr.Example != "" {
					if result[i].Example != "" && result[i].Example != newAttr.Example {
						if result[i].VersionExamples == nil {
							result[i].VersionExamples = make(map[string]string)
							result[i].VersionExamples["_base"] = result[i].Example
						}
						result[i].VersionExamples[overrideVersion] = newAttr.Example
					}
					result[i].Example = newAttr.Example
				}
				lastEnumBaseline := result[i].lastEnumValues
				if lastEnumBaseline == nil {
					// Never tracked yet -- result[i].EnumValues is still the untouched base
					// list at this point (the union only ever grows inside this same block,
					// which also always updates lastEnumValues, so nil here means "no delta
					// has ever set EnumValues before this one").
					lastEnumBaseline = result[i].EnumValues
				}
				if len(newAttr.EnumValues) > 0 && !stringSlicesEqual(lastEnumBaseline, newAttr.EnumValues) {
					// Enum sets differ between versions — record per-version sets and use the union at schema level.
					// Compared against lastEnumBaseline (the true immediately-preceding delta's
					// own list), not result[i].EnumValues, which is a monotonically-growing union.
					if result[i].VersionEnums == nil {
						result[i].VersionEnums = make(map[string][]string)
						if len(result[i].EnumValues) > 0 {
							result[i].VersionEnums["_base"] = result[i].EnumValues
						}
					}
					result[i].VersionEnums[overrideVersion] = newAttr.EnumValues
					result[i].EnumValues = unionStringSlices(result[i].EnumValues, newAttr.EnumValues)
				} else if len(newAttr.EnumValues) > 0 {
					result[i].EnumValues = newAttr.EnumValues
				}
				if len(newAttr.EnumValues) > 0 {
					// Track this delta's own explicit list unconditionally, decoupled from the
					// ever-growing union above, so the next fold's comparison is always against
					// the true immediately-preceding version.
					result[i].lastEnumValues = newAttr.EnumValues
				}

				// Handle range merging - detect when ranges differ across versions
				if newAttr.MinInt != 0 || newAttr.MaxInt != 0 {
					// hasBase reflects whether a real prior range exists at all (unaffected by
					// the fix — read from the true schema scalar, which is only frozen, never
					// zeroed, once a range exists). The divergence comparison itself is against
					// lastRangeMin/Max — the true immediately-preceding delta's own value — not
					// result[i].MinInt/MaxInt, which freezes at the original base value once
					// VersionRanges exists.
					hasBase := result[i].MinInt != 0 || result[i].MaxInt != 0
					// A tracked range always has a non-zero max (goyang gives every integer a full-width range),
					// so lastRangeMax != 0 means "a delta has stated a range" and lastRangeMin may legitimately be 0.
					baseMin, baseMax := result[i].MinInt, result[i].MaxInt
					if result[i].lastRangeMax != 0 {
						baseMin, baseMax = result[i].lastRangeMin, result[i].lastRangeMax
					}
					overrideMin := newAttr.MinInt
					overrideMax := newAttr.MaxInt

					if hasBase && (overrideMin != baseMin || overrideMax != baseMax) {
						// Ranges differ - initialize version ranges map if needed
						if result[i].VersionRanges == nil {
							result[i].VersionRanges = make(map[string]RangeConstraint)
							// Add base version range if we have one
							result[i].VersionRanges["_base"] = RangeConstraint{Min: result[i].MinInt, Max: result[i].MaxInt}
						}
						// Add override version range
						result[i].VersionRanges[overrideVersion] = RangeConstraint{Min: overrideMin, Max: overrideMax}
					} else {
						// Ranges are the same or one is not set - use single range
						if newAttr.MinInt != 0 {
							result[i].MinInt = newAttr.MinInt
						}
						if newAttr.MaxInt != 0 {
							result[i].MaxInt = newAttr.MaxInt
						}
					}
					// Track this delta's own explicit range unconditionally, decoupled from the
					// frozen schema scalar above, so the next fold's comparison is always against
					// the true immediately-preceding version.
					result[i].lastRangeMin = newAttr.MinInt
					if newAttr.MaxInt != 0 {
						result[i].lastRangeMax = newAttr.MaxInt
					}
				}
				if len(newAttr.StringPatterns) > 0 && !stringSlicesEqual(result[i].StringPatterns, newAttr.StringPatterns) {
					if result[i].VersionPatterns == nil {
						result[i].VersionPatterns = make(map[string][]string)
						if len(result[i].StringPatterns) > 0 {
							result[i].VersionPatterns["_base"] = result[i].StringPatterns
						}
					}
					result[i].VersionPatterns[overrideVersion] = newAttr.StringPatterns
					result[i].StringPatterns = newAttr.StringPatterns
				} else if len(newAttr.StringPatterns) > 0 {
					result[i].StringPatterns = newAttr.StringPatterns
				}
				// Compared against lastStringMinLength/MaxLength (the true immediately-preceding
				// delta's own value), not result[i].StringMinLength/MaxLength, which only ever
				// widens and so can go stale relative to the true previous version.
				baseMin := result[i].lastStringMinLength
				if baseMin == 0 {
					// Never tracked yet -- result[i].StringMinLength is still the untouched
					// base value at this point (widening only ever runs inside this same
					// block, which also always updates lastStringMinLength).
					baseMin = result[i].StringMinLength
				}
				baseMax := result[i].lastStringMaxLength
				if baseMax == 0 {
					baseMax = result[i].StringMaxLength
				}
				newMin := newAttr.StringMinLength
				newMax := newAttr.StringMaxLength
				minChanged := newMin != 0 && newMin != baseMin
				maxChanged := newMax != 0 && newMax != baseMax
				if minChanged || maxChanged {
					if result[i].VersionStringLengths == nil {
						result[i].VersionStringLengths = make(map[string]StringLengthConstraint)
						// Seed "_base" from the real schema scalar (the true base value), not
						// from lastStringMinLength/MaxLength, which is still 0 on first divergence.
						if result[i].StringMinLength != 0 || result[i].StringMaxLength != 0 {
							result[i].VersionStringLengths["_base"] = StringLengthConstraint{Min: result[i].StringMinLength, Max: result[i].StringMaxLength}
						}
					}
					result[i].VersionStringLengths[overrideVersion] = StringLengthConstraint{Min: newMin, Max: newMax}
					// Widen to most permissive span across all versions. Unchanged from before --
					// this part was already correct.
					if newMin != 0 && (result[i].StringMinLength == 0 || newMin < result[i].StringMinLength) {
						result[i].StringMinLength = newMin
					}
					if newMax != 0 && newMax > result[i].StringMaxLength {
						result[i].StringMaxLength = newMax
					}
				} else {
					if newMin != 0 {
						result[i].StringMinLength = newMin
					}
					if newMax != 0 {
						result[i].StringMaxLength = newMax
					}
				}
				// Track this delta's own explicit value unconditionally, decoupled from the
				// widened scalar above, so the next fold's comparison is always against the
				// true immediately-preceding version.
				if newMin != 0 {
					result[i].lastStringMinLength = newMin
				}
				if newMax != 0 {
					result[i].lastStringMaxLength = newMax
				}
				if newAttr.DefaultValue != "" {
					result[i].DefaultValue = newAttr.DefaultValue
				}
				if newAttr.RequiresReplace {
					result[i].RequiresReplace = newAttr.RequiresReplace
				}
				if newAttr.NoAugmentConfig {
					result[i].NoAugmentConfig = newAttr.NoAugmentConfig
				}
				// Determine delete modes for version-diff logic.
				overrideMode := ""
				if newAttr.DeleteGrandparent {
					overrideMode = "grandparent"
				} else if newAttr.DeleteParent {
					overrideMode = "parent"
				}
				baseMode := ""
				if result[i].DeleteGrandparent {
					baseMode = "grandparent"
				} else if result[i].DeleteParent {
					baseMode = "parent"
				}
				if overrideMode != baseMode && newAttr.ReplacesYangName != "" {
					// Replacement with a different delete mode → per-version map.
					// Only replacements (replaces_yang_name set) can intentionally change
					// the delete mode; a non-replacing override without a delete flag
					// inherits the base mode via OR merge below.
					if result[i].VersionDeleteMode == nil {
						result[i].VersionDeleteMode = make(map[string]string)
						result[i].VersionDeleteMode["_base"] = baseMode
					}
					result[i].VersionDeleteMode[overrideVersion] = overrideMode
					if newAttr.DeleteParent {
						result[i].DeleteParent = true
					}
					if newAttr.DeleteGrandparent {
						result[i].DeleteGrandparent = true
					}
				} else {
					// Same mode, or no replaces_yang_name — OR-only merge.
					// Non-replacing overrides inherit the base delete mode implicitly.
					if newAttr.DeleteParent {
						result[i].DeleteParent = newAttr.DeleteParent
					}
					if newAttr.DeleteGrandparent {
						result[i].DeleteGrandparent = newAttr.DeleteGrandparent
					}
				}
				if newAttr.NoDelete {
					result[i].NoDelete = newAttr.NoDelete
				}
				if len(newAttr.TestTags) > 0 {
					if !stringSlicesEqual(result[i].TestTags, newAttr.TestTags) {
						if result[i].VersionTestTags == nil {
							result[i].VersionTestTags = make(map[string][]string)
							result[i].VersionTestTags["_base"] = result[i].TestTags
						}
						result[i].VersionTestTags[overrideVersion] = newAttr.TestTags
					}
					result[i].TestTags = newAttr.TestTags
				}
				if newAttr.MinimumTestValue != "" {
					if result[i].MinimumTestValue != "" && result[i].MinimumTestValue != newAttr.MinimumTestValue {
						if result[i].VersionMinimumTestValues == nil {
							result[i].VersionMinimumTestValues = make(map[string]string)
							result[i].VersionMinimumTestValues["_base"] = result[i].MinimumTestValue
						}
						result[i].VersionMinimumTestValues[overrideVersion] = newAttr.MinimumTestValue
					}
					result[i].MinimumTestValue = newAttr.MinimumTestValue
				}

				if newAttr.ReplacesYangName != "" {
					if result[i].Reference {
						log.Printf("WARNING: replaces_yang_name on Reference attribute %q has no effect (Reference attrs excluded from body)", newAttr.TfName)
					}
					if result[i].Id && (result[i].AddedInVersion != "" || result[i].RemovedInVersion != "") {
						log.Fatalf("attribute %q: replaces_yang_name on a key attribute (id: true) cannot be combined with AddedInVersion or RemovedInVersion — these code paths are not reconciled in the template", result[i].TfName)
					}
					if result[i].VersionYangNames == nil {
						result[i].VersionYangNames = make(map[string]string)
						result[i].VersionYangNames["_base"] = result[i].YangName
						result[i].ReplacesYangName = newAttr.ReplacesYangName
						result[i].ReplacesXPath = originalXPath
					}
					result[i].VersionYangNames[overrideVersion] = newAttr.YangName
					result[i].YangName = newAttr.YangName
					// XPath is already updated by the override block above (lines that handle newAttr.XPath != "")
				}

				found = true
				break
			}
		}
		if !found {
			// Skip legacy attributes that don't exist in the base — nothing to remove
			if newAttr.Legacy {
				continue
			}
			if newAttr.ReplacesYangName != "" {
				log.Fatalf("replaces_yang_name %q on attribute %q not found in base definition",
					newAttr.ReplacesYangName, newAttr.TfName)
			}
			// Add new attribute and mark with version
			newAttr.AddedInVersion = overrideVersion
			// Mark all nested attributes too
			markAttributesWithVersion(&newAttr, overrideVersion)
			result = append(result, newAttr)
		}
	}

	return result
}

// markAttributesWithVersion recursively marks attributes and their children with a version
func markAttributesWithVersion(attr *YamlConfigAttribute, version string) {
	if attr.AddedInVersion == "" {
		attr.AddedInVersion = version
	}
	for i := range attr.Attributes {
		markAttributesWithVersion(&attr.Attributes[i], version)
	}
}

// fixBaseVersionInRanges replaces "_base" placeholder with actual base version
func fixBaseVersionInRanges(config *YamlConfig, baseVersion string) {
	if config.VersionTestTags != nil {
		if base, exists := config.VersionTestTags["_base"]; exists {
			delete(config.VersionTestTags, "_base")
			config.VersionTestTags[baseVersion] = base
		}
	}
	// test_prerequisites is recorded per declaring version only (see mergeConfigs), so
	// there's no "_base" sentinel to resolve here. But if the base version declared its own
	// test_prerequisites and this resource has other versions too, the base's entry must
	// exist in the map under its own real key -- otherwise the single-constant template
	// branch would apply the base's prerequisites to every version's test, ignoring any
	// higher version's override.
	if len(config.SupportedVersions) > 1 && len(config.TestPrerequisites) > 0 {
		if config.VersionTestPrerequisites == nil {
			config.VersionTestPrerequisites = make(map[string][]YamlTest)
		}
		if _, exists := config.VersionTestPrerequisites[baseVersion]; !exists {
			config.VersionTestPrerequisites[baseVersion] = config.TestPrerequisites
		}
	}
	for i := range config.Attributes {
		fixAttributeBaseVersion(&config.Attributes[i], baseVersion)
	}
}

// fixAttributeBaseVersion recursively resolves the "_base" sentinel in attribute and its children.
// baseVersion is the label inherited from the parent (the resource base version at the top level);
// an attribute added in a later version labels its seed with its own AddedInVersion instead.
func fixAttributeBaseVersion(attr *YamlConfigAttribute, baseVersion string) {
	if attr.AddedInVersion != "" {
		baseVersion = attr.AddedInVersion
	}
	if attr.VersionRanges != nil {
		if baseRange, exists := attr.VersionRanges["_base"]; exists {
			delete(attr.VersionRanges, "_base")
			attr.VersionRanges[baseVersion] = baseRange
		}
	}
	if attr.VersionEnums != nil {
		if baseEnums, exists := attr.VersionEnums["_base"]; exists {
			delete(attr.VersionEnums, "_base")
			attr.VersionEnums[baseVersion] = baseEnums
		}
	}
	if attr.VersionStringLengths != nil {
		if base, exists := attr.VersionStringLengths["_base"]; exists {
			delete(attr.VersionStringLengths, "_base")
			attr.VersionStringLengths[baseVersion] = base
		}
	}
	if attr.VersionPatterns != nil {
		if base, exists := attr.VersionPatterns["_base"]; exists {
			delete(attr.VersionPatterns, "_base")
			attr.VersionPatterns[baseVersion] = base
		}
	}
	if attr.VersionXPath != nil {
		if base, exists := attr.VersionXPath["_base"]; exists {
			delete(attr.VersionXPath, "_base")
			attr.VersionXPath[baseVersion] = base
		}
	}
	if attr.VersionDeleteMode != nil {
		if baseMode, exists := attr.VersionDeleteMode["_base"]; exists {
			delete(attr.VersionDeleteMode, "_base")
			attr.VersionDeleteMode[baseVersion] = baseMode
		}
	}
	if attr.VersionYangNames != nil {
		if base, ok := attr.VersionYangNames["_base"]; ok {
			delete(attr.VersionYangNames, "_base")
			attr.VersionYangNames[baseVersion] = base
			for v := range attr.VersionYangNames {
				if v != baseVersion && (attr.MovedInVersion == "" || v < attr.MovedInVersion) {
					attr.MovedInVersion = v
				}
			}
		}
	}
	if attr.VersionExamples != nil {
		if baseEx, exists := attr.VersionExamples["_base"]; exists {
			delete(attr.VersionExamples, "_base")
			attr.VersionExamples[baseVersion] = baseEx
		}
	}
	if attr.VersionMinimumTestValues != nil {
		if baseVal, exists := attr.VersionMinimumTestValues["_base"]; exists {
			delete(attr.VersionMinimumTestValues, "_base")
			attr.VersionMinimumTestValues[baseVersion] = baseVal
		}
	}
	if attr.VersionTestTags != nil {
		if baseTags, exists := attr.VersionTestTags["_base"]; exists {
			delete(attr.VersionTestTags, "_base")
			attr.VersionTestTags[baseVersion] = baseTags
		}
	}
	for i := range attr.Attributes {
		fixAttributeBaseVersion(&attr.Attributes[i], baseVersion)
	}
}

// mergeConfigs merges two configurations, with the newer version overriding the older
func mergeConfigs(base, override YamlConfig) YamlConfig {
	merged := base

	// Check if entire resource is marked as legacy in this version
	if override.Legacy {
		// Mark entire resource as removed in this version
		merged.RemovedInVersion = override.Version
		merged.Legacy = true
		log.Printf("  Resource '%s' marked as legacy/removed in version %s", merged.Name, override.Version)
		return merged
	}

	// Override basic fields if set in override
	if override.Name != "" {
		merged.Name = override.Name
	}
	if override.Version != "" {
		merged.Version = override.Version
	}
	if override.Path != "" {
		merged.Path = override.Path
	}
	if override.AugmentPath != "" {
		merged.AugmentPath = override.AugmentPath
	}
	if override.NoDelete {
		merged.NoDelete = override.NoDelete
	}
	if override.NoDeleteAttributes {
		merged.NoDeleteAttributes = override.NoDeleteAttributes
	}
	if override.DefaultDeleteAttributes {
		merged.DefaultDeleteAttributes = override.DefaultDeleteAttributes
	}
	if len(override.TestTags) > 0 {
		if !stringSlicesEqual(merged.TestTags, override.TestTags) {
			if merged.VersionTestTags == nil {
				merged.VersionTestTags = make(map[string][]string)
				// "_base" sentinel, not base.Version -- mergeConfigs overwrites
				// merged.Version on every call, so by fold 2+ it no longer holds the
				// true original base version. Resolved to the real base version by
				// fixBaseVersionInRanges after all folds complete.
				merged.VersionTestTags["_base"] = merged.TestTags
			}
			merged.VersionTestTags[override.Version] = override.TestTags
		}
		merged.TestTags = override.TestTags
	}
	if override.SkipMinimumTest {
		merged.SkipMinimumTest = override.SkipMinimumTest
	}
	if override.NoAugmentConfig {
		merged.NoAugmentConfig = override.NoAugmentConfig
	}
	if override.DsDescription != "" {
		merged.DsDescription = override.DsDescription
	}
	if override.ResDescription != "" {
		merged.ResDescription = override.ResDescription
	}
	if override.DocCategory != "" {
		merged.DocCategory = override.DocCategory
	}
	// test_prerequisites is recorded only for the version that declares it, under its own
	// real version key, with no "_base" sentinel and no merged.TestPrerequisites
	// carry-forward. Inheritance by higher versions happens at test time, where
	// selectVersionPrerequisitesConfig picks the highest declared version at or below
	// IOSXR_VERSION.
	if len(override.TestPrerequisites) > 0 {
		if merged.VersionTestPrerequisites == nil {
			merged.VersionTestPrerequisites = make(map[string][]YamlTest)
		}
		merged.VersionTestPrerequisites[override.Version] = override.TestPrerequisites
	}

	// Accumulate path_version entries from the override.
	if len(override.PathVersion) > 0 {
		if merged.PathVersion == nil {
			merged.PathVersion = make(map[string]string)
		}
		for ver, path := range override.PathVersion {
			merged.PathVersion[ver] = path
		}
	}

	// Merge attributes - pass override version so new attributes can be marked
	merged.Attributes = mergeAttributes(base.Attributes, override.Attributes, override.Version)

	return merged
}

// globToRegex converts a glob-style pattern to a regular expression.
// A bare '*' becomes '.*', '?' becomes '.', and everything else is
// treated as a literal regex character so that full regex syntax also works.
// Examples:
//
//	"router_bgp*"          → case-insensitive prefix match
//	"router_(bgp|isis).*"  → full regex (mixed mode)
func globToRegex(pattern string) string {
	var b strings.Builder
	// If the pattern already looks like a "real" regex (contains anchors, groups,
	// character classes, alternation, or quantifiers other than *), pass it through
	// as-is. We detect this by checking for characters that have no glob meaning.
	isRegex := strings.ContainsAny(pattern, "()[]{}^$|+")
	if isRegex {
		return pattern
	}
	// Glob mode: escape everything except * and ?
	for _, ch := range pattern {
		switch ch {
		case '*':
			b.WriteString(".*")
		case '?':
			b.WriteString(".")
		default:
			b.WriteString(regexp.QuoteMeta(string(ch)))
		}
	}
	return b.String()
}

func main() {
	filterFlag := flag.String("filter", "", `Regex/glob filter for definition names (e.g. --filter="router_bgp*" or --filter="router_(bgp|isis).*")`)
	flag.Parse()

	// Build the compiled filter regex (nil = no filter = generate all).
	var filterRegex *regexp.Regexp
	if *filterFlag != "" {
		pattern := globToRegex(*filterFlag)
		var err error
		filterRegex, err = regexp.Compile("(?i)" + pattern)
		if err != nil {
			log.Fatalf("Invalid --filter pattern %q: %v", *filterFlag, err)
		}
		log.Printf("Filter active: %q  (compiled regex: %s)", *filterFlag, filterRegex)
	} else if flag.NArg() == 1 {
		// Legacy: positional argument treated as an exact (case-insensitive) name match.
		filterRegex = regexp.MustCompile("(?i)^" + regexp.QuoteMeta(flag.Arg(0)) + "$")
		log.Printf("Filter (legacy positional): exact match for %q", flag.Arg(0))
	}

	// Get all version directories and sort them (lowest to highest)
	versionDirs, err := os.ReadDir(definitionsPath)
	if err != nil {
		log.Fatalf("Error reading definitions directory: %v", err)
	}

	versions := make([]string, 0)
	for _, versionDir := range versionDirs {
		if versionDir.IsDir() {
			versions = append(versions, versionDir.Name())
		}
	}
	sort.Slice(versions, func(i, j int) bool {
		return versionCompare(versions[i], versions[j]) < 0
	})

	log.Printf("Found versions: %v", versions)

	// First pass: load all definition files per version
	// definitionsByNameAndVersion: resource name -> version -> raw (un-augmented) config
	definitionsByNameAndVersion := make(map[string]map[string]YamlConfig)

	for _, version := range versions {
		versionDefPath := filepath.Join(definitionsPath, version)
		items, err := os.ReadDir(versionDefPath)
		if err != nil {
			log.Printf("Warning: Could not read definitions for version %s: %v", version, err)
			continue
		}

		for _, filename := range items {
			if filepath.Ext(filename.Name()) != ".yaml" {
				continue
			}
			yamlFile, err := os.ReadFile(filepath.Join(versionDefPath, filename.Name()))
			if err != nil {
				log.Fatalf("Error reading file %s: %v", filename.Name(), err)
			}
			config := YamlConfig{}
			err = yaml.Unmarshal(yamlFile, &config)
			if err != nil {
				log.Fatalf("Error parsing yaml %s: %v", filename.Name(), err)
			}
			if definitionsByNameAndVersion[config.Name] == nil {
				definitionsByNameAndVersion[config.Name] = make(map[string]YamlConfig)
			}
			config.Version = version
			definitionsByNameAndVersion[config.Name][version] = config
		}
	}

	allConfigs := make([]YamlConfig, 0)

	// augmentedCache caches augmented configs per (name, version) so we don't re-parse YANG twice
	augmentedCache := make(map[string]YamlConfig) // key = "name@version"

	// Helper: augment a config from YANG models for a given version
	augmentForVersion := func(cfg YamlConfig, version string, accPathVersion map[string]string) YamlConfig {
		cacheKey := cfg.Name + "@" + version
		if cached, ok := augmentedCache[cacheKey]; ok {
			return cached
		}
		if !cfg.NoAugmentConfig && !cfg.Legacy {
			vModelPath := filepath.Join(modelsPath, version)
			var vModelPaths []string
			if _, err := os.Stat(vModelPath); err == nil {
				modelItems, _ := os.ReadDir(vModelPath)
				for _, item := range modelItems {
					if filepath.Ext(item.Name()) == ".yang" {
						vModelPaths = append(vModelPaths, filepath.Join(vModelPath, item.Name()))
					}
				}
			}
			if len(vModelPaths) == 0 {
				log.Fatalf("definition %q @ %s: no YANG models found in %s", cfg.Name, version, vModelPath)
			}
			log.Printf("  Augmenting '%s' version %s from YANG models", cfg.Name, version)
			// Without augment_path, use the path_version entry in effect for this version.
			origAugmentPath := cfg.AugmentPath
			if cfg.AugmentPath == "" {
				cfg.AugmentPath = resolveYangPathForVersion(version, accPathVersion, cfg.PathVersion)
			}
			augmentConfig(&cfg, vModelPaths)
			cfg.AugmentPath = origAugmentPath
		}
		augmentedCache[cacheKey] = cfg
		return cfg
	}

	// NEW UNIFIED APPROACH: Generate ONE file per resource with version metadata
	// For each resource name, merge ALL versions into a single unified config
	for defName, versionConfigs := range definitionsByNameAndVersion {
		if filterRegex != nil {
			// Match against both the raw name ("Router BGP AF Group") and the
			// snake_case name ("router_bgp_af_group") so that either form works
			// in the --filter flag.
			if !filterRegex.MatchString(defName) && !filterRegex.MatchString(SnakeCase(defName)) {
				continue
			}
		}

		log.Printf("Generating unified '%s' (merging all versions)", defName)

		// Collect all versions that have this resource, sorted
		resourceVersions := make([]string, 0)
		for _, v := range versions {
			if _, exists := versionConfigs[v]; exists {
				resourceVersions = append(resourceVersions, v)
			}
		}

		if len(resourceVersions) == 0 {
			log.Printf("  WARNING: No versions found for %s, skipping", defName)
			continue
		}

		log.Printf("  Found in versions: %v", resourceVersions)

		// Build cumulative merge across ALL versions
		var unifiedConfig YamlConfig
		for i, v := range resourceVersions {
			if i > 0 {
				validateNoAugmentConfigCarryover(unifiedConfig, versionConfigs[v], v)
			}
			augmented := augmentForVersion(versionConfigs[v], v, unifiedConfig.PathVersion)

			if i == 0 {
				// First version becomes the base - DO NOT mark attributes with version
				// because base version is the default and doesn't need tracking
				unifiedConfig = augmented
				log.Printf("  Base version: %s", v)
			} else {
				// Merge subsequent versions
				log.Printf("  Merging version: %s", v)
				unifiedConfig = mergeConfigs(unifiedConfig, augmented)
			}
		}

		// Set unified file properties
		unifiedConfig.Version = "" // No version suffix in type names
		unifiedConfig.SupportedVersions = resourceVersions
		unifiedConfig.BaseVersion = resourceVersions[0] // First version is the base

		// If this resource was introduced in a version above the global minimum,
		// record it so templates can emit "# Supported from version X".
		if len(versions) > 0 && resourceVersions[0] != versions[0] {
			unifiedConfig.IntroducedInVersion = resourceVersions[0]
		}

		// Fix base version placeholders in VersionRanges
		fixBaseVersionInRanges(&unifiedConfig, resourceVersions[0])

		// Detect if this resource has version-specific differences
		unifiedConfig.HasVersionDifferences = hasVersionDifferences(unifiedConfig)
		unifiedConfig.HasPathVersion = len(unifiedConfig.PathVersion) > 0

		// Generate unified files (no version suffix due to versionSuffix: false)
		for _, tmpl := range templates {
			outputFileName := tmpl.prefix + SnakeCase(unifiedConfig.Name)
			outputFileName += tmpl.suffix
			log.Printf("  Generating: %s", filepath.Base(outputFileName))
			renderTemplate(tmpl.path, outputFileName, unifiedConfig)
		}

		allConfigs = append(allConfigs, unifiedConfig)
	}

	// Group configs by name for provider template (not strictly needed now, kept for changelog etc.)
	type ProviderTemplateData struct {
		ResourcesByName       map[string][]YamlConfig
		DataSourcesByName     map[string][]YamlConfig
		UniqueResourceNames   []string
		UniqueDataSourceNames []string
	}

	templateData := ProviderTemplateData{
		ResourcesByName:   make(map[string][]YamlConfig),
		DataSourcesByName: make(map[string][]YamlConfig),
	}

	uniqueNames := make(map[string]bool)
	for _, config := range allConfigs {
		templateData.ResourcesByName[config.Name] = append(templateData.ResourcesByName[config.Name], config)
		templateData.DataSourcesByName[config.Name] = append(templateData.DataSourcesByName[config.Name], config)
		uniqueNames[config.Name] = true
	}
	for name := range uniqueNames {
		templateData.UniqueResourceNames = append(templateData.UniqueResourceNames, name)
		templateData.UniqueDataSourceNames = append(templateData.UniqueDataSourceNames, name)
	}
	sort.Strings(templateData.UniqueResourceNames)
	sort.Strings(templateData.UniqueDataSourceNames)

	log.Println("Generating provider.go")
	renderTemplate(providerTemplate, providerLocation, templateData)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))

	// Write supported_versions_gen.go so helpers.DefinitionVersions stays in sync
	// with the directories present under gen/definitions/.
	var versionLiterals []string
	for _, v := range versions {
		versionLiterals = append(versionLiterals, fmt.Sprintf("%q", v))
	}
	supportedVersionsContent := fmt.Sprintf(
		"// Code generated by \"gen/generator.go\"; DO NOT EDIT.\n\npackage helpers\n\n"+
			"// DefinitionVersions lists the IOS-XR versions that have a definitions\n"+
			"// directory under gen/definitions/. Updated automatically on make gen.\n"+
			"var DefinitionVersions = []string{%s}\n",
		strings.Join(versionLiterals, ", "),
	)
	if err := os.WriteFile("./internal/provider/helpers/supported_versions_gen.go",
		[]byte(supportedVersionsContent), 0644); err != nil {
		log.Fatalf("Error writing supported_versions_gen.go: %v", err)
	}
	log.Println("Generated internal/provider/helpers/supported_versions_gen.go")

	// Write version changes data for doc_version_changes.go to consume
	writeVersionChangesData(allConfigs)

	log.Printf("\nGeneration complete! Processed %d resource(s) across all versions.", len(allConfigs))
}

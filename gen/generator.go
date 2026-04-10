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
	Name                    string                `yaml:"name"`
	Version                 string                `yaml:"version"` // drives type-name suffix; empty for unified files
	SupportedVersions       []string              // All versions this resource supports (for unified files)
	BaseVersion             string                // The minimum/base version (first version in SupportedVersions)
	HasVersionDifferences   bool                  // True if there are version-specific changes (added/removed fields or definitions)
	Path                    string                `yaml:"path"`
	AugmentPath             string                `yaml:"augment_path"`
	NoDelete                bool                  `yaml:"no_delete"`
	NoDeleteAttributes      bool                  `yaml:"no_delete_attributes"`
	DefaultDeleteAttributes bool                  `yaml:"default_delete_attributes"`
	TestTags                []string              `yaml:"test_tags"`
	SkipMinimumTest         bool                  `yaml:"skip_minimum_test"`
	NoAugmentConfig         bool                  `yaml:"no_augment_config"`
	DsDescription           string                `yaml:"ds_description"`
	ResDescription          string                `yaml:"res_description"`
	DocCategory             string                `yaml:"doc_category"`
	Legacy                  bool                  `yaml:"legacy"` // If true, entire resource is removed/not available in this version
	RemovedInVersion        string                // Version where entire resource was removed (set when Legacy is true)
	Attributes              []YamlConfigAttribute `yaml:"attributes"`
	TestPrerequisites       []YamlTest            `yaml:"test_prerequisites"`
}

type YamlConfigAttribute struct {
	YangName          string                     `yaml:"yang_name"`
	YangScope         string                     `yaml:"yang_scope"`
	TfName            string                     `yaml:"tf_name"`
	XPath             string                     `yaml:"xpath"`
	Type              string                     `yaml:"type"`
	ReadRaw           bool                       `yaml:"read_raw"`
	TypeYangBool      string                     `yaml:"type_yang_bool"`
	Id                bool                       `yaml:"id"`
	Reference         bool                       `yaml:"reference"`
	Mandatory         bool                       `yaml:"mandatory"`
	Optional          bool                       `yaml:"optional"`
	WriteOnly         bool                       `yaml:"write_only"`
	Sensitive         bool                       `yaml:"sensitive"`
	ExcludeTest       bool                       `yaml:"exclude_test"`
	ExcludeExample    bool                       `yaml:"exclude_example"`
	IncludeExample    bool                       `yaml:"include_example"`
	Description       string                     `yaml:"description"`
	Example           string                     `yaml:"example"`
	EnumValues        []string                   `yaml:"enum_values"`
	MinInt            int64                      `yaml:"min_int"`
	MaxInt            int64                      `yaml:"max_int"`
	StringPatterns    []string                   `yaml:"string_patterns"`
	StringMinLength   int64                      `yaml:"string_min_length"`
	StringMaxLength   int64                      `yaml:"string_max_length"`
	DefaultValue      string                     `yaml:"default_value"`
	RequiresReplace   bool                       `yaml:"requires_replace"`
	NoAugmentConfig   bool                       `yaml:"no_augment_config"`
	DeleteParent      bool                       `yaml:"delete_parent"`
	DeleteGrandparent bool                       `yaml:"delete_grandparent"`
	NoDelete          bool                       `yaml:"no_delete"`
	TestTags          []string                   `yaml:"test_tags"`
	MinimumTestValue  string                     `yaml:"minimum_test_value"`
	AddedInVersion    string                     // Which version introduced this attribute (e.g., "2442", "2522")
	RemovedInVersion  string                     // Which version removed this attribute (populated when legacy: true)
	Legacy            bool                       `yaml:"legacy"` // If true, this attribute is removed/dropped in this version
	VersionRanges     map[string]RangeConstraint // Version-specific ranges for Int64 fields (nil if same across all versions)
	Attributes        []YamlConfigAttribute      `yaml:"attributes"`
}

// RangeConstraint represents min/max constraints for a version
type RangeConstraint struct {
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
		parts[i] = strings.ReplaceAll(part, ".", "\\\\.")
	}
	return strings.Join(parts, ".")
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

// Templating helper function to get xpath if available
func GetXPath(yangPath, xPath string) string {
	if xPath != "" {
		return xPath
	}
	return yangPath
}

func GetDeletePath(attribute YamlConfigAttribute) string {
	path := GetXPath(attribute.YangName, attribute.XPath)
	if attribute.DeleteGrandparent {
		// Remove two levels: grandparent
		return RemoveLastPathElement(RemoveLastPathElement(path))
	}
	if attribute.DeleteParent {
		return RemoveLastPathElement(path)
	}
	return path
}

func ReverseAttributes(attributes []YamlConfigAttribute) []YamlConfigAttribute {
	reversed := make([]YamlConfigAttribute, len(attributes))
	for i, v := range attributes {
		reversed[len(attributes)-1-i] = v
	}
	return reversed
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

// Templating helper function to identify last element of list
func IsLast(index int, len int) bool {
	return index+1 == len
}

// Templating helper function to remove last element of path
func RemoveLastPathElement(p string) string {
	return path.Dir(p)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Templating helper function to generate version suffix for type names (e.g., "2522" -> "V2522")
func VersionSuffix(version string) string {
	if version == "" {
		return ""
	}
	return "V" + version
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
		if len(attr.Attributes) > 0 && hasAttributeVersionDifferences(attr.Attributes) {
			return true
		}
	}
	return false
}

// FormatVersionDisplay formats a version string for display (e.g., "2442" -> "24.4.2")
func FormatVersionDisplay(version string) string {
	if len(version) == 4 {
		return fmt.Sprintf("%s.%s.%s", version[0:2], version[2:3], version[3:4])
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

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":                       ToGoName,
	"toJsonPath":                     ToJsonPath,
	"camelCase":                      CamelCase,
	"snakeCase":                      SnakeCase,
	"versionSuffix":                  VersionSuffix,
	"hasId":                          HasId,
	"hasReference":                   HasReference,
	"importParts":                    ImportParts,
	"importAttributes":               ImportAttributes,
	"add":                            Add,
	"getExamplePath":                 GetExamplePath,
	"isLast":                         IsLast,
	"sprintf":                        fmt.Sprintf,
	"removeLastPathElement":          RemoveLastPathElement,
	"getXPath":                       GetXPath,
	"getDeletePath":                  GetDeletePath,
	"reverseAttributes":              ReverseAttributes,
	"collectVersionConstraints":      CollectVersionConstraints,
	"hasVersionConstraints":          HasVersionConstraints,
	"formatVersionRanges":            FormatVersionRanges,
	"formatVersionDisplay":           FormatVersionDisplay,
	"hasVersionRanges":               HasVersionRanges,
	"collectVersionRangeConstraints": CollectVersionRangeConstraints,
	"getWidestRange":                 GetWidestRange,
}

func resolvePath(e *yang.Entry, path string) *yang.Entry {
	pathElements := strings.Split(path, "/")

	for _, pathElement := range pathElements {
		if len(pathElement) > 0 {
			// remove key
			if strings.Contains(pathElement, "[") {
				pathElement = pathElement[:strings.Index(pathElement, "[")]
			}
			// remove reference
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

func parseAttribute(e *yang.Entry, attr *YamlConfigAttribute) {
	leaf := resolvePath(e, attr.YangName)
	//fmt.Printf("%s, Entry: %+v\n\n", attr.YangName, e)
	//fmt.Printf("%s, Kind: %+v, Type: %+v\n\n", leaf.Name, leaf.Kind, leaf.Type)
	if leaf.Kind.String() == "Leaf" {
		if leaf.ListAttr != nil {
			if contains([]string{"string", "union", "leafref"}, leaf.Type.Kind.String()) {
				attr.Type = "StringList"
			} else if contains([]string{"int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
				attr.Type = "Int64List"
			} else {
				panic(fmt.Sprintf("Unknown leaf-list type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
			}
			// TODO parse union type
		} else if contains([]string{"string", "union", "leafref"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			if leaf.Type.Length != nil {
				attr.StringMinLength = int64(leaf.Type.Length[0].Min.Value)
				max := leaf.Type.Length[0].Max.Value
				// hack to not introduce unsigned types
				if max > math.MaxInt64 {
					max = math.MaxInt64
				}
				attr.StringMaxLength = int64(max)
			}
			if len(leaf.Type.Pattern) > 0 {
				attr.StringPatterns = leaf.Type.Pattern
			}
		} else if contains([]string{"int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
			attr.Type = "Int64"
			if leaf.Type.Range != nil {
				if attr.MinInt == 0 {
					attr.MinInt = int64(leaf.Type.Range[0].Min.Value)
					if leaf.Type.Range[0].Min.Negative {
						attr.MinInt = -attr.MinInt
					}
				}
				max := leaf.Type.Range[0].Max.Value
				// hack to not introduce unsigned types
				if max > math.MaxInt64 {
					max = math.MaxInt64
				}
				if attr.MaxInt == 0 {
					attr.MaxInt = int64(max)
				}
			}
		} else if contains([]string{"boolean", "empty"}, leaf.Type.Kind.String()) {
			if leaf.Type.Kind.String() == "boolean" {
				attr.TypeYangBool = "boolean"
			} else if leaf.Type.Kind.String() == "empty" {
				attr.TypeYangBool = "empty"
			}
			attr.Type = "Bool"
		} else if contains([]string{"enumeration"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			attr.EnumValues = leaf.Type.Enum.Names()
		} else {
			panic(fmt.Sprintf("Unknown leaf type, attribute: %s, type: %s", attr.YangName, leaf.Type.Kind.String()))
		}
	}
	if _, ok := leaf.Extra["presence"]; ok {
		attr.TypeYangBool = "presence"
		attr.Type = "Bool"
	}
	if attr.XPath == "" {
		attr.XPath = attr.YangName
	}
	if attr.TfName == "" {
		tfName := strings.ReplaceAll(ToYangShortName(attr.XPath), "-", "_")
		tfName = strings.ReplaceAll(tfName, "/", "_")
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

func augmentConfig(config *YamlConfig, modelPaths []string) {
	path := ""
	if config.AugmentPath != "" {
		path = config.AugmentPath
	} else {
		path = config.Path
	}

	module := strings.Split(path, ":")[0]
	e, errors := yang.GetModule(module, modelPaths...)
	if len(errors) > 0 {
		fmt.Printf("YANG parser error(s): %+v\n\n", errors)
		return
	}

	// Print definition/model info
	fmt.Printf("Processing definition: %s\n", config.Name)
	//fmt.Printf("Resolving yang model: %s ==> Resolved: %s\n", module, e.Name)

	p := path[len(module)+1:]
	e = resolvePath(e, p)

	addKeys(e, config)

	for ia := range config.Attributes {
		if config.Attributes[ia].Id || config.Attributes[ia].Reference || config.Attributes[ia].NoAugmentConfig {
			continue
		}
		parseAttribute(e, &config.Attributes[ia])
		if config.Attributes[ia].Type == "List" {
			el := resolvePath(e, config.Attributes[ia].YangName)
			for iaa := range config.Attributes[ia].Attributes {
				if config.Attributes[ia].Attributes[iaa].NoAugmentConfig || config.Attributes[ia].Attributes[iaa].Legacy {
					continue
				}
				parseAttribute(el, &config.Attributes[ia].Attributes[iaa])
				if config.Attributes[ia].Attributes[iaa].Type == "List" {
					ell := resolvePath(el, config.Attributes[ia].Attributes[iaa].YangName)
					for iaaa := range config.Attributes[ia].Attributes[iaa].Attributes {
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig || config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Legacy {
							continue
						}
						parseAttribute(ell, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa])
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Type == "List" {
							elll := resolvePath(ell, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName)
							for iaaaa := range config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes {
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].NoAugmentConfig || config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].Legacy {
									continue
								}
								parseAttribute(elll, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa])
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
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
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
		var newContent string
		currentSectionName := ""
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
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
			// Match by yang_name or tf_name
			if result[i].YangName == newAttr.YangName ||
				(result[i].TfName != "" && newAttr.TfName != "" && result[i].TfName == newAttr.TfName) {
				// Legacy: true means this attribute is removed in the override version
				// Instead of dropping it, mark it with RemovedInVersion for validation
				if newAttr.Legacy {
					result[i].RemovedInVersion = overrideVersion
					result[i].Legacy = true
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

				// Override all other fields from new attribute
				if newAttr.YangScope != "" {
					result[i].YangScope = newAttr.YangScope
				}
				if newAttr.TfName != "" {
					result[i].TfName = newAttr.TfName
				}
				if newAttr.XPath != "" {
					result[i].XPath = newAttr.XPath
				}
				if newAttr.Type != "" {
					result[i].Type = newAttr.Type
				}
				if newAttr.ReadRaw {
					result[i].ReadRaw = newAttr.ReadRaw
				}
				if newAttr.TypeYangBool != "" {
					result[i].TypeYangBool = newAttr.TypeYangBool
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
				if newAttr.ExcludeTest {
					result[i].ExcludeTest = newAttr.ExcludeTest
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
					result[i].Example = newAttr.Example
				}
				if len(newAttr.EnumValues) > 0 {
					result[i].EnumValues = newAttr.EnumValues
				}

				// Handle range merging - detect when ranges differ across versions
				if newAttr.MinInt != 0 || newAttr.MaxInt != 0 {
					// Check if ranges differ
					baseMin := result[i].MinInt
					baseMax := result[i].MaxInt
					overrideMin := newAttr.MinInt
					overrideMax := newAttr.MaxInt

					if (baseMin != 0 || baseMax != 0) && (overrideMin != baseMin || overrideMax != baseMax) {
						// Ranges differ - initialize version ranges map if needed
						if result[i].VersionRanges == nil {
							result[i].VersionRanges = make(map[string]RangeConstraint)
							// Add base version range if we have one
							if baseMin != 0 || baseMax != 0 {
								// Find base version - it's the first non-override version
								// We'll set this later when we know all versions
								result[i].VersionRanges["_base"] = RangeConstraint{Min: baseMin, Max: baseMax}
							}
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
				}
				if len(newAttr.StringPatterns) > 0 {
					result[i].StringPatterns = newAttr.StringPatterns
				}
				if newAttr.StringMinLength != 0 {
					result[i].StringMinLength = newAttr.StringMinLength
				}
				if newAttr.StringMaxLength != 0 {
					result[i].StringMaxLength = newAttr.StringMaxLength
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
				if newAttr.DeleteParent {
					result[i].DeleteParent = newAttr.DeleteParent
				}
				if newAttr.DeleteGrandparent {
					result[i].DeleteGrandparent = newAttr.DeleteGrandparent
				}
				if newAttr.NoDelete {
					result[i].NoDelete = newAttr.NoDelete
				}
				if len(newAttr.TestTags) > 0 {
					result[i].TestTags = newAttr.TestTags
				}
				if newAttr.MinimumTestValue != "" {
					result[i].MinimumTestValue = newAttr.MinimumTestValue
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
	for i := range config.Attributes {
		fixAttributeBaseVersion(&config.Attributes[i], baseVersion)
	}
}

// fixAttributeBaseVersion recursively fixes base version in attribute and its children
func fixAttributeBaseVersion(attr *YamlConfigAttribute, baseVersion string) {
	if attr.VersionRanges != nil {
		if baseRange, exists := attr.VersionRanges["_base"]; exists {
			delete(attr.VersionRanges, "_base")
			attr.VersionRanges[baseVersion] = baseRange
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
	if len(override.TestPrerequisites) > 0 {
		merged.TestPrerequisites = override.TestPrerequisites
	}

	// Merge attributes - pass override version so new attributes can be marked

	merged.Attributes = mergeAttributes(base.Attributes, override.Attributes, override.Version)

	return merged
}

func main() {
	resourceName := ""

	if len(os.Args) == 2 {
		resourceName = os.Args[1]
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
	augmentForVersion := func(cfg YamlConfig, version string) YamlConfig {
		cacheKey := cfg.Name + "@" + version
		if cached, ok := augmentedCache[cacheKey]; ok {
			return cached
		}
		if !cfg.NoAugmentConfig {
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
			if len(vModelPaths) > 0 {
				log.Printf("  Augmenting '%s' version %s from YANG models", cfg.Name, version)
				augmentConfig(&cfg, vModelPaths)
			}
		}
		augmentedCache[cacheKey] = cfg
		return cfg
	}

	// NEW UNIFIED APPROACH: Generate ONE file per resource with version metadata
	// For each resource name, merge ALL versions into a single unified config
	for defName, versionConfigs := range definitionsByNameAndVersion {
		if resourceName != "" && !strings.EqualFold(defName, resourceName) {
			continue
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
			augmented := augmentForVersion(versionConfigs[v], v)

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

		// Fix base version placeholders in VersionRanges
		fixBaseVersionInRanges(&unifiedConfig, resourceVersions[0])

		// Detect if this resource has version-specific differences
		unifiedConfig.HasVersionDifferences = hasVersionDifferences(unifiedConfig)

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

	log.Printf("\nGeneration complete! Processed %d resource(s) across all versions.", len(allConfigs))
}

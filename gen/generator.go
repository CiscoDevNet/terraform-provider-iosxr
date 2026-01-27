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
	path   string
	prefix string
	suffix string
}

var templates = []t{
	{
		path:   "./gen/templates/model.go",
		prefix: "./internal/provider/model_iosxr_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go",
		prefix: "./internal/provider/data_source_iosxr_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go",
		prefix: "./internal/provider/data_source_iosxr_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go",
		prefix: "./internal/provider/resource_iosxr_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go",
		prefix: "./internal/provider/resource_iosxr_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/data-source.tf",
		prefix: "./examples/data-sources/iosxr_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/resource.tf",
		prefix: "./examples/resources/iosxr_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/import.sh",
		prefix: "./examples/resources/iosxr_",
		suffix: "/import.sh",
	},
}

type YamlConfig struct {
	Name                    string                `yaml:"name"`
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
	Attributes              []YamlConfigAttribute `yaml:"attributes"`
	TestPrerequisites       []YamlTest            `yaml:"test_prerequisites"`
}

type YamlConfigAttribute struct {
	YangName          string                `yaml:"yang_name"`
	YangScope         string                `yaml:"yang_scope"`
	TfName            string                `yaml:"tf_name"`
	XPath             string                `yaml:"xpath"`
	Type              string                `yaml:"type"`
	ReadRaw           bool                  `yaml:"read_raw"`
	TypeYangBool      string                `yaml:"type_yang_bool"`
	Id                bool                  `yaml:"id"`
	Reference         bool                  `yaml:"reference"`
	Mandatory         bool                  `yaml:"mandatory"`
	Optional          bool                  `yaml:"optional"`
	WriteOnly         bool                  `yaml:"write_only"`
	Sensitive         bool                  `yaml:"sensitive"`
	ExcludeTest       bool                  `yaml:"exclude_test"`
	ExcludeExample    bool                  `yaml:"exclude_example"`
	IncludeExample    bool                  `yaml:"include_example"`
	Description       string                `yaml:"description"`
	Example           string                `yaml:"example"`
	EnumValues        []string              `yaml:"enum_values"`
	MinInt            int64                 `yaml:"min_int"`
	MaxInt            int64                 `yaml:"max_int"`
	StringPatterns    []string              `yaml:"string_patterns"`
	StringMinLength   int64                 `yaml:"string_min_length"`
	StringMaxLength   int64                 `yaml:"string_max_length"`
	DefaultValue      string                `yaml:"default_value"`
	RequiresReplace   bool                  `yaml:"requires_replace"`
	NoAugmentConfig   bool                  `yaml:"no_augment_config"`
	DeleteParent      bool                  `yaml:"delete_parent"`
	DeleteGrandparent bool                  `yaml:"delete_grandparent"`
	NoDelete          bool                  `yaml:"no_delete"`
	TestTags          []string              `yaml:"test_tags"`
	MinimumTestValue  string                `yaml:"minimum_test_value"`
	Attributes        []YamlConfigAttribute `yaml:"attributes"`
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
	path := attribute.XPath
	if attribute.DeleteGrandparent {
		// Remove two levels: grandparent
		return RemoveLastPathElement(RemoveLastPathElement(path))
	}
	if attribute.DeleteParent {
		return RemoveLastPathElement(path)
	}
	return path
}

func GetLastPathElement(path string) string {
	// Remove namespace prefix if present
	// e.g., "ipv4//Cisco-IOS-XR-um-if-ip-address-cfg:addresses/address/address" -> "address"
	// Split by / and get the last non-empty element
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] != "" {
			// Remove namespace prefix if present (e.g., "Cisco-IOS-XR-um:element" -> "element")
			element := parts[i]
			if idx := strings.LastIndex(element, ":"); idx >= 0 {
				element = element[idx+1:]
			}
			return element
		}
	}
	return ""
}

func ToDotPath(path string) string {
	// Remove leading slash
	path = strings.TrimPrefix(path, "/")
	// Replace double slashes with single dot
	path = strings.ReplaceAll(path, "//", ".")
	// Replace single slashes with dots
	path = strings.ReplaceAll(path, "/", ".")
	return path
}

// ToGnmiPath converts a path to GNMI format
// For iosxr (which uses NETCONF/gNMI), we just return the path as-is
func ToGnmiPath(path string) string {
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

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":              ToGoName,
	"camelCase":             CamelCase,
	"snakeCase":             SnakeCase,
	"hasId":                 HasId,
	"hasReference":          HasReference,
	"importParts":           ImportParts,
	"importAttributes":      ImportAttributes,
	"add":                   Add,
	"sprintf":               fmt.Sprintf,
	"removeLastPathElement": RemoveLastPathElement,
	"getDeletePath":         GetDeletePath,
	"getLastPathElement":    GetLastPathElement,
	"reverseAttributes":     ReverseAttributes,
	"toDotPath":             ToDotPath,
	"toGnmiPath":            ToGnmiPath,
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
		} else if helpers.Contains([]string{"int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
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
		} else if helpers.Contains([]string{"boolean", "empty"}, leaf.Type.Kind.String()) {
			if leaf.Type.Kind.String() == "boolean" {
				attr.TypeYangBool = "boolean"
			} else if leaf.Type.Kind.String() == "empty" {
				attr.TypeYangBool = "empty"
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
		attr.TypeYangBool = "presence"
		attr.Type = "Bool"
	}
	if attr.XPath == "" {
		attr.XPath = attr.YangName
	}
	if attr.TfName == "" {
		tfName := strings.ReplaceAll(ToYangShortName(attr.XPath), "-", "_")
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
		// Default XPath from YangName if not explicitly set (do this first for all attributes)
		if config.Attributes[ia].XPath == "" {
			config.Attributes[ia].XPath = config.Attributes[ia].YangName
		}
		if config.Attributes[ia].Id || config.Attributes[ia].Reference || config.Attributes[ia].NoAugmentConfig {
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
				parseAttribute(el, &config.Attributes[ia].Attributes[iaa])
				if config.Attributes[ia].Attributes[iaa].Type == "List" {
					ell := resolvePath(el, config.Attributes[ia].Attributes[iaa].YangName)
					for iaaa := range config.Attributes[ia].Attributes[iaa].Attributes {
						// Default XPath from YangName if not explicitly set (do this first for all attributes)
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].XPath == "" {
							config.Attributes[ia].Attributes[iaa].Attributes[iaaa].XPath = config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName
						}
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].NoAugmentConfig {
							continue
						}
						parseAttribute(ell, &config.Attributes[ia].Attributes[iaa].Attributes[iaaa])
						if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Type == "List" {
							elll := resolvePath(ell, config.Attributes[ia].Attributes[iaa].Attributes[iaaa].YangName)
							for iaaaa := range config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes {
								// Default XPath from YangName if not explicitly set (do this first for all attributes)
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].XPath == "" {
									config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].XPath = config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].YangName
								}
								if config.Attributes[ia].Attributes[iaa].Attributes[iaaa].Attributes[iaaaa].NoAugmentConfig {
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
		existingScanner := bufio.NewScanner(existingFile)
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
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	fmt.Println("=== GENERATOR STARTING ===")
	resourceName := ""
	if len(os.Args) == 2 {
		resourceName = os.Args[1]
		fmt.Printf("Filtering for resource: %s\n", resourceName)
	}

	items, _ := os.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))

	// Load configs
	for i, filename := range items {
		fmt.Printf("Processing: %s\n", filename.Name())
		yamlFile, err := os.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file '%s': %v", filename.Name(), err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml file '%s': %v", filename.Name(), err)
		}
		configs[i] = config
	}

	items, _ = os.ReadDir(modelsPath)
	modelPaths := make([]string, 0)

	// Iterate over yang models
	for _, item := range items {
		if filepath.Ext(item.Name()) == ".yang" {
			modelPaths = append(modelPaths, filepath.Join(modelsPath, item.Name()))
		}
	}

	for i := range configs {
		if resourceName != "" && configs[i].Name != resourceName {
			continue
		}
		// Augment config by yang models
		if !configs[i].NoAugmentConfig {
			augmentConfig(&configs[i], modelPaths)
		}

		// Iterate over templates and render files
		for _, t := range templates {
			renderTemplate(t.path, t.prefix+SnakeCase(configs[i].Name)+t.suffix, configs[i])
		}
	}

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, configs)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}

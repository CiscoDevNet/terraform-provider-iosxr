//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/openconfig/goyang/pkg/yang"
	"gopkg.in/yaml.v3"
)

const (
	definitionsPath  = "./gen/definitions/"
	modelsPath       = "./gen/models/"
	providerTemplate = "./gen/templates/provider.go"
	providerLocation = "./internal/provider/provider.go"
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
	Name              string                `yaml:"name"`
	Path              string                `yaml:"path"`
	NoDelete          bool                  `yaml:"no_delete"`
	ExcludeTest       bool                  `yaml:"exclude_test"`
	NoAugmentConfig   bool                  `yaml:"no_augment_config"`
	DsDescription     string                `yaml:"ds_description"`
	ResDescription    string                `yaml:"res_description"`
	DocCategory       string                `yaml:"doc_category"`
	Attributes        []YamlConfigAttribute `yaml:"attributes"`
	TestPrerequisites []YamlTest            `yaml:"test_prerequisites"`
}

type YamlConfigAttribute struct {
	YangName        string   `yaml:"yang_name"`
	TfName          string   `yaml:"tf_name"`
	Type            string   `yaml:"type"`
	Id              bool     `yaml:"id"`
	Reference       bool     `yaml:"reference"`
	Mandatory       bool     `yaml:"mandatory"`
	Optional        bool     `yaml:"optional"`
	WriteOnly       bool     `yaml:"write_only"`
	ExcludeTest     bool     `yaml:"exclude_test"`
	Description     string   `yaml:"description"`
	Example         string   `yaml:"example"`
	EnumValues      []string `yaml:"enum_values"`
	MinInt          int64    `yaml:"min_int"`
	MaxInt          int64    `yaml:"max_int"`
	StringPatterns  []string `yaml:"string_patterns"`
	StringMinLength int64    `yaml:"string_min_length"`
	StringMaxLength int64    `yaml:"string_max_length"`
	DefaultValue    string   `yaml:"default_value"`
}

type YamlTest struct {
	Path         string              `yaml:"path"`
	Attributes   []YamlTestAttribute `yaml:"attributes"`
	Dependencies []string            `yaml:"dependencies"`
}

type YamlTestAttribute struct {
	Name      string `yaml:"name"`
	Value     string `yaml:"value"`
	Reference string `yaml:"reference"`
}

// Templating helper function to get short YAMG name without prefix (xxx:abc -> abc)
func ToYangShortName(s string) string {
	if strings.Contains(s, ":") {
		s = strings.Split(s, ":")[1]
	}
	return s
}

// Templating helper function to convert YANG name to GO name
func ToGoName(s string) string {
	s = ToYangShortName(s)
	var g []string

	p := strings.Split(s, "-")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	s = strings.ReplaceAll(s, "/", "")
	if s == "Id" {
		s = "YangId"
	}
	return s
}

// Templating helper function to convert YANG name to GO name
func ToJsonPath(yangPath string) string {
	return strings.ReplaceAll(yangPath, "/", ".")
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":       ToGoName,
	"toJsonPath":     ToJsonPath,
	"camelCase":      CamelCase,
	"snakeCase":      SnakeCase,
	"hasId":          HasId,
	"getExamplePath": GetExamplePath,
	"isLast":         IsLast,
	"sprintf":        fmt.Sprintf,
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
					if config.Attributes[i].YangName == key {
						keyAttr = &config.Attributes[i]
					}
				}
				if keyAttr == nil {
					panic(fmt.Sprintf("Cannot find id/reference attribute in config: %s", key))
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
	// fmt.Printf("%s, Kind: %+v, Type: %+v\n\n", leaf.Name, leaf.Kind, leaf.Type)
	if leaf.Kind.String() == "Leaf" {
		// TODO parse union type
		if contains([]string{"string", "union"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			if leaf.Type.Length != nil {
				attr.StringMinLength = int64(leaf.Type.Length[0].Min.Value)
				attr.StringMaxLength = int64(leaf.Type.Length[0].Max.Value)
			}
			if len(leaf.Type.Pattern) > 0 {
				attr.StringPatterns = leaf.Type.Pattern
			}
		} else if contains([]string{"uint8", "uint16", "uint32", "uint64"}, leaf.Type.Kind.String()) {
			attr.Type = "Int64"
			if leaf.Type.Range != nil {
				attr.MinInt = int64(leaf.Type.Range[0].Min.Value)
				max := leaf.Type.Range[0].Max.Value
				// hack to not introduce unsigned types
				if max > math.MaxInt64 {
					max = math.MaxInt64
				}
				attr.MaxInt = int64(max)
			}
		} else if contains([]string{"boolean"}, leaf.Type.Kind.String()) {
			attr.Type = "Bool"
		} else if contains([]string{"enumeration"}, leaf.Type.Kind.String()) {
			attr.Type = "String"
			attr.EnumValues = leaf.Type.Enum.Names()
		}
	}
	if _, ok := leaf.Extra["presence"]; ok {
		attr.Type = "Bool"
	}
	if attr.TfName == "" {
		tfName := strings.ReplaceAll(ToYangShortName(attr.YangName), "-", "_")
		tfName = strings.ReplaceAll(tfName, "/", "_")
		attr.TfName = tfName
	}
	if attr.Description == "" {
		attr.Description = strings.ReplaceAll(leaf.Description, "\n", " ")
	}
	if !attr.Mandatory && attr.DefaultValue == "" && !attr.Optional {
		attr.Mandatory = leaf.Mandatory.Value()
	}
}

func augmentConfig(config *YamlConfig, modelPaths []string) {

	module := strings.Split(config.Path, ":")[0]
	e, errors := yang.GetModule(module, modelPaths...)
	if len(errors) > 0 {
		fmt.Printf("YANG parser error(s): %+v\n\n", errors)
		return
	}

	p := config.Path[len(module)+1:]
	e = resolvePath(e, p)

	addKeys(e, config)

	for ia := range config.Attributes {
		if config.Attributes[ia].Id || config.Attributes[ia].Reference {
			continue
		}
		parseAttribute(e, &config.Attributes[ia])
	}

	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s configuration.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage the %s configuration.", config.Name)
	}
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

	// create output file
	outputFile := filepath.Join(outputPath)
	os.MkdirAll(filepath.Dir(outputFile), 0755)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// format go code
	if strings.HasSuffix(templatePath, ".go") {
		fOutput, err := format.Source(output.Bytes())
		if err != nil {
			log.Fatalf("Error formatting go in %s: %v", templatePath, err)
		}
		f.Write(fOutput)
	} else {
		f.Write(output.Bytes())
	}
}

func main() {
	items, _ := ioutil.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))

	// Load configs
	for i, filename := range items {
		yamlFile, err := ioutil.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs[i] = config
	}

	items, _ = ioutil.ReadDir(modelsPath)
	modelPaths := make([]string, 0)

	// Iterate over yang models
	for _, item := range items {
		if filepath.Ext(item.Name()) == ".yang" {
			modelPaths = append(modelPaths, filepath.Join(modelsPath, item.Name()))
		}
	}

	for i := range configs {
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
}

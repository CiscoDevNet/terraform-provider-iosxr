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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var generalResources = []string{"gnmi", "cli"}

const (
	definitionsPath = "./gen/definitions/"
)

type YamlConfig struct {
	Name        string `yaml:"name"`
	DocCategory string `yaml:"doc_category"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}
var generalDocPaths = []string{"./docs/data-sources/", "./docs/resources/"}

func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

func main() {
	// Find the base version directory (lowest version = base)
	versionDirs, err := os.ReadDir(definitionsPath)
	if err != nil {
		log.Fatalf("Error reading definitions directory: %v", err)
	}

	// Find first version directory (base version)
	baseVersionPath := ""
	for _, d := range versionDirs {
		if d.IsDir() {
			baseVersionPath = filepath.Join(definitionsPath, d.Name())
			break
		}
	}
	if baseVersionPath == "" {
		log.Fatalf("No version directories found in %s", definitionsPath)
	}

	items, _ := os.ReadDir(baseVersionPath)
	configs := make([]YamlConfig, 0)

	// Load configs from base version only
	for _, filename := range items {
		if filepath.Ext(filename.Name()) != ".yaml" {
			continue
		}
		yamlFile, err := os.ReadFile(filepath.Join(baseVersionPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs = append(configs, config)
	}

	for i := range configs {
		for _, path := range docPaths {
			filename := path + SnakeCase(configs[i].Name) + ".md"
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+configs[i].DocCategory+`"`)

			os.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Update general resources with "General" subcategory
	for _, resource := range generalResources {
		for _, path := range generalDocPaths {
			filename := fmt.Sprintf("%s%s.md", path, resource)
			content, err := os.ReadFile(filename)
			if err != nil {
				// Skip if file doesn't exist (e.g., data source may not exist for all resources)
				if os.IsNotExist(err) {
					continue
				}
				log.Fatalf("Error opening documentation: %v", err)
			}
			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "General"`)
			os.WriteFile(filename, []byte(s), 0644)
		}
	}
}

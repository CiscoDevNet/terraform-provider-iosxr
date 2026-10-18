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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	versionChangesDataPath = "./gen/version_changes_data.json"
	resourceDocsPath       = "./docs/resources/"
	dataSourceDocsPath     = "./docs/data-sources/"
)

type VersionedAttrRow struct {
	TfName    string `json:"tf_name"`
	RemovedIn string `json:"removed_in,omitempty"`
}

type ResourceVersionChanges struct {
	Removed []VersionedAttrRow `json:"removed"`
}

type VersionChangesData map[string]ResourceVersionChanges

func buildVersionCompatSection(changes ResourceVersionChanges) string {
	var sb strings.Builder
	sb.WriteString("## Version Compatibility\n")
	sb.WriteString("\n### Removed from version\n\n")
	sb.WriteString("| Attribute | Version |\n")
	sb.WriteString("|-----------|:-------:|\n")
	for _, row := range changes.Removed {
		sb.WriteString(fmt.Sprintf("| `%s` | `%s` |\n", row.TfName, row.RemovedIn))
	}
	sb.WriteString("\n")
	return sb.String()
}

func main() {
	raw, err := os.ReadFile(versionChangesDataPath)
	if err != nil {
		log.Fatalf("Error reading %s: %v", versionChangesDataPath, err)
	}

	var data VersionChangesData
	if err := json.Unmarshal(raw, &data); err != nil {
		log.Fatalf("Error parsing %s: %v", versionChangesDataPath, err)
	}

	fmt.Println("rendering multi-version removed from docs")

	injectVersionCompat(data, resourceDocsPath, "resources")
	injectVersionCompat(data, dataSourceDocsPath, "data-sources")
}

func injectVersionCompat(data VersionChangesData, docsPath string, label string) {
	for resourceName, changes := range data {
		if len(changes.Removed) == 0 {
			continue
		}

		docFile := docsPath + resourceName + ".md"
		content, err := os.ReadFile(docFile)
		if err != nil {
			// Not all resources/data-sources have a doc file yet; skip silently
			continue
		}

		s := string(content)

		// Insert the Version Compatibility section before ## Example Usage.
		// Fall back to before ## Schema if Example Usage is absent.
		insertBefore := "## Example Usage"
		if !strings.Contains(s, insertBefore) {
			insertBefore = "## Schema"
		}
		if !strings.Contains(s, insertBefore) {
			fmt.Printf("skipping \"docs/%s/%s.md\": no insertion point found\n", label, resourceName)
			continue
		}

		section := buildVersionCompatSection(changes)
		s = strings.Replace(s, insertBefore, section+insertBefore, 1)

		if err := os.WriteFile(docFile, []byte(s), 0644); err != nil {
			log.Fatalf("Error writing %s: %v", docFile, err)
		}
		fmt.Printf("rendering \"docs/%s/%s.md\"\n", label, resourceName)
	}
}

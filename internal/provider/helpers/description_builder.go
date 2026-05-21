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
	"fmt"
	"sort"
	"strings"
)

type AttributeDescription struct {
	String string
}

func NewAttributeDescription(s string) *AttributeDescription {
	return &AttributeDescription{s}
}

func (d *AttributeDescription) AddDefaultValueDescription(defaultValue string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Default value: `%s`", d.String, defaultValue)
	return d
}

func (d *AttributeDescription) AddStringEnumDescription(values ...string) *AttributeDescription {
	v := make([]string, len(values))
	for i, value := range values {
		v[i] = fmt.Sprintf("`%s`", value)
	}
	d.String = fmt.Sprintf("%s\n  - Choices: %s", d.String, strings.Join(v, ", "))
	return d
}

func (d *AttributeDescription) AddIntegerRangeDescription(min, max int64) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Range: `%v`-`%v`", d.String, min, max)
	return d
}

func (d *AttributeDescription) AddVersionRangeDescription(versionRanges map[string]struct{ Min, Max int64 }) *AttributeDescription {
	if len(versionRanges) == 0 {
		return d
	}

	// Sort versions for consistent output
	versions := make([]string, 0, len(versionRanges))
	for v := range versionRanges {
		versions = append(versions, v)
	}
	sort.Strings(versions)

	d.String = fmt.Sprintf("%s\n  - Range:", d.String)
	for _, v := range versions {
		r := versionRanges[v]
		formattedVersion := FormatVersion(v)
		d.String = fmt.Sprintf("%s `%v`-`%v` (v%s),", d.String, r.Min, r.Max, formattedVersion)
	}
	// Remove trailing comma
	d.String = strings.TrimSuffix(d.String, ",")
	return d
}

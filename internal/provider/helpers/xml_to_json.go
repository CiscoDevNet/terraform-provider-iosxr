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
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// elementInfo tracks XML element state during parsing
type elementInfo struct {
	name           string
	namespace      string
	jsonPath       string
	hasText        bool
	childCount     int
	isListItem     bool
	parentName     string
	hasLeafChild   bool
	uniqueChildren map[string]int
}

func newElementInfo(name, namespace, jsonPath, parentName string, isListItem bool) elementInfo {
	return elementInfo{
		name:           name,
		namespace:      namespace,
		jsonPath:       jsonPath,
		parentName:     parentName,
		isListItem:     isListItem,
		uniqueChildren: make(map[string]int),
	}
}

// XMLToJSON converts XML input to JSON format with namespace prefixes
// Automatically skips container-only levels (no leaf values)
func XMLToJSON(xmlInput string) (string, error) {
	decoder := xml.NewDecoder(strings.NewReader(xmlInput))
	parser := &xmlParser{
		jsonStr: "{}",
		decoder: decoder,
	}
	return parser.parse()
}

type xmlParser struct {
	jsonStr      string
	decoder      *xml.Decoder
	stack        []elementInfo
	inData       bool
	dataDepth    int
	currentDepth int
}

func (p *xmlParser) parse() (string, error) {
	for {
		token, err := p.decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to parse XML: %w", err)
		}

		switch t := token.(type) {
		case xml.StartElement:
			if err := p.handleStartElement(t); err != nil {
				return "", err
			}
		case xml.EndElement:
			p.handleEndElement(t)
		case xml.CharData:
			p.handleCharData(t)
		}
	}
	return p.jsonStr, nil
}

func (p *xmlParser) handleStartElement(t xml.StartElement) error {
	p.currentDepth++

	if t.Name.Local == "data" {
		p.inData = true
		p.dataDepth = p.currentDepth
		return nil
	}

	if !p.inData || p.currentDepth <= p.dataDepth {
		return nil
	}

	// Skip first element after data (root container)
	if p.currentDepth == p.dataDepth+1 {
		p.stack = append(p.stack, newElementInfo(t.Name.Local, "", "", "", false))
		return nil
	}

	namespace := extractNamespace(t.Attr)
	key := buildKey(t.Name.Local, namespace)

	var parentName string
	var isListItem bool
	if len(p.stack) > 0 {
		parentName = p.stack[len(p.stack)-1].name
		isListItem = shouldBeListItem(t.Name.Local, parentName)
		p.stack[len(p.stack)-1].childCount++
		p.stack[len(p.stack)-1].uniqueChildren[t.Name.Local]++
	}

	jsonPath := p.buildJSONPath(key)
	p.stack = append(p.stack, newElementInfo(t.Name.Local, namespace, jsonPath, parentName, isListItem))

	return nil
}

func (p *xmlParser) buildJSONPath(key string) string {
	if len(p.stack) == 0 {
		return key
	}

	parent := p.stack[len(p.stack)-1]
	if parent.jsonPath == "" || (len(p.stack) == 1 && !parent.hasLeafChild) {
		return key
	}

	return parent.jsonPath + "." + key
}

func (p *xmlParser) handleEndElement(t xml.EndElement) {
	p.currentDepth--

	if t.Name.Local == "data" {
		p.inData = false
		return
	}

	if !p.inData || len(p.stack) == 0 {
		return
	}

	elem := p.stack[len(p.stack)-1]
	p.stack = p.stack[:len(p.stack)-1]

	// Skip container-only levels
	if p.shouldSkipContainer(elem) {
		p.flattenContainer(elem)
		return
	}

	// Mark parent as having leaf child
	if (elem.childCount == 0 || elem.hasText) && len(p.stack) > 0 {
		p.stack[len(p.stack)-1].hasLeafChild = true
	}

	if elem.jsonPath == "" {
		return
	}

	p.setJSONValue(elem)
}

func (p *xmlParser) shouldSkipContainer(elem elementInfo) bool {
	return len(p.stack) == 1 &&
		p.stack[0].jsonPath == "" &&
		elem.childCount > 0 &&
		!elem.hasText &&
		!elem.hasLeafChild &&
		elem.jsonPath != ""
}

func (p *xmlParser) flattenContainer(elem elementInfo) {
	val := gjson.Get(p.jsonStr, elem.jsonPath)
	if !val.Exists() {
		return
	}

	if m, ok := val.Value().(map[string]interface{}); ok {
		p.jsonStr = "{}"
		for k, v := range m {
			p.jsonStr, _ = sjson.Set(p.jsonStr, k, v)
		}
	}
}

func (p *xmlParser) setJSONValue(elem elementInfo) {
	if elem.hasText {
		if elem.isListItem {
			currentValue := gjson.Get(p.jsonStr, elem.jsonPath)
			if currentValue.Exists() {
				p.jsonStr, _ = sjson.Set(p.jsonStr, elem.jsonPath, []interface{}{currentValue.Value()})
			}
		}
	} else if elem.childCount == 0 {
		// Empty element
		value := map[string]interface{}{}
		if elem.isListItem {
			p.jsonStr, _ = sjson.Set(p.jsonStr, elem.jsonPath, []interface{}{value})
		} else {
			p.jsonStr, _ = sjson.Set(p.jsonStr, elem.jsonPath, value)
		}
	} else {
		// Has children
		if elem.isListItem {
			currentValue := gjson.Get(p.jsonStr, elem.jsonPath)
			if currentValue.Exists() && currentValue.IsObject() {
				p.jsonStr, _ = sjson.Set(p.jsonStr, elem.jsonPath, []interface{}{currentValue.Value()})
			}
		} else {
			currentValue := gjson.Get(p.jsonStr, elem.jsonPath)
			if !currentValue.Exists() {
				p.jsonStr, _ = sjson.Set(p.jsonStr, elem.jsonPath, map[string]interface{}{})
			}
		}
	}
}

func (p *xmlParser) handleCharData(t xml.CharData) {
	if !p.inData || len(p.stack) == 0 {
		return
	}

	text := strings.TrimSpace(string(t))
	if text == "" {
		return
	}

	elem := &p.stack[len(p.stack)-1]
	elem.hasText = true
	elem.hasLeafChild = true

	// Mark all parents as having leaf children
	for i := range p.stack {
		p.stack[i].hasLeafChild = true
	}

	if elem.jsonPath != "" {
		p.jsonStr, _ = sjson.Set(p.jsonStr, elem.jsonPath, text)
	}
}

func buildKey(localName, namespace string) string {
	if namespace != "" {
		nsPrefix := extractNamespacePrefix(namespace)
		return nsPrefix + ":" + localName
	}
	return localName
}

func shouldBeListItem(childName, parentName string) bool {
	if parentName == "" {
		return false
	}

	// Check if parent is plural of child
	if strings.HasSuffix(parentName, "s") && childName == strings.TrimSuffix(parentName, "s") {
		return true
	}

	// Check common list suffixes
	listSuffixes := []string{"-string", "-list", "-entry", "-item"}
	for _, suffix := range listSuffixes {
		if childName == parentName+suffix {
			return true
		}
	}

	return false
}

package helpers

import (
	"testing"
)

// ============================================================================
// IsListPath Tests
// ============================================================================

func TestIsListPath(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		expected bool
	}{
		{
			name:     "list with simple predicate",
			xPath:    "/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']",
			expected: true,
		},
		{
			name:     "list with quoted predicate",
			xPath:    `/interfaces/interface[interface-name="GigabitEthernet0/0/0/0"]`,
			expected: true,
		},
		{
			name:     "list with multiple predicates using and",
			xPath:    "/policy-map/class[name='CM-HIGH' and type='qos']",
			expected: true,
		},
		{
			name:     "list with multiple separate predicates",
			xPath:    "/policy-map/class[name='CM-HIGH'][type='qos']",
			expected: true,
		},
		{
			name:     "container without predicate",
			xPath:    "/interfaces/interface/ipv4",
			expected: false,
		},
		{
			name:     "simple single segment",
			xPath:    "/hostname",
			expected: false,
		},
		{
			name:     "predicate in middle segment but not at end",
			xPath:    "/interfaces/interface[interface-name='Gi0/0/0/0']/ipv4/address",
			expected: false,
		},
		{
			name:     "empty path",
			xPath:    "",
			expected: false,
		},
		{
			name:     "list path with trailing whitespace",
			xPath:    "/interfaces/interface[interface-name='Gi0/0/0/0']   ",
			expected: true,
		},
		{
			name:     "container path with trailing whitespace",
			xPath:    "/interfaces/interface   ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsListPath(tt.xPath)
			if result != tt.expected {
				t.Errorf("IsListPath(%q) = %v, expected %v", tt.xPath, result, tt.expected)
			}
		})
	}
}

// ============================================================================
// ParseXPathSegment Tests
// ============================================================================

func TestParseXPathSegment(t *testing.T) {
	tests := []struct {
		name         string
		segment      string
		wantElement  string
		wantKeyCount int
		wantKeys     []KeyValue
	}{
		{
			name:         "simple element without predicates",
			segment:      "interface",
			wantElement:  "interface",
			wantKeyCount: 0,
		},
		{
			name:         "element with single predicate single quotes",
			segment:      "interface[interface-name='GigabitEthernet0/0/0/0']",
			wantElement:  "interface",
			wantKeyCount: 1,
			wantKeys:     []KeyValue{{Key: "interface-name", Value: "GigabitEthernet0/0/0/0"}},
		},
		{
			name:         "element with single predicate double quotes",
			segment:      `interface[interface-name="GigabitEthernet0/0/0/0"]`,
			wantElement:  "interface",
			wantKeyCount: 1,
			wantKeys:     []KeyValue{{Key: "interface-name", Value: "GigabitEthernet0/0/0/0"}},
		},
		{
			name:         "element with multiple predicates using and",
			segment:      "class[name='CM-HIGH' and type='qos']",
			wantElement:  "class",
			wantKeyCount: 2,
			wantKeys: []KeyValue{
				{Key: "name", Value: "CM-HIGH"},
				{Key: "type", Value: "qos"},
			},
		},
		{
			name:         "element with multiple separate predicates",
			segment:      "class[name='CM-HIGH'][type='qos']",
			wantElement:  "class",
			wantKeyCount: 2,
			wantKeys: []KeyValue{
				{Key: "name", Value: "CM-HIGH"},
				{Key: "type", Value: "qos"},
			},
		},
		{
			name:         "element with three predicates using and",
			segment:      "route[vrf='default' and dst='10.0.0.0' and prefix='24']",
			wantElement:  "route",
			wantKeyCount: 3,
			wantKeys: []KeyValue{
				{Key: "vrf", Value: "default"},
				{Key: "dst", Value: "10.0.0.0"},
				{Key: "prefix", Value: "24"},
			},
		},
		{
			name:         "element with numeric key value",
			segment:      "vlan[id='100']",
			wantElement:  "vlan",
			wantKeyCount: 1,
			wantKeys:     []KeyValue{{Key: "id", Value: "100"}},
		},
		{
			name:         "element with dot in predicate value",
			segment:      "interface[interface-name='Bundle-Ether10.666']",
			wantElement:  "interface",
			wantKeyCount: 1,
			wantKeys:     []KeyValue{{Key: "interface-name", Value: "Bundle-Ether10.666"}},
		},
		{
			name:         "element name with hyphen",
			segment:      "interface-configuration[active='act'][interface-name='Gi0/0/0/0']",
			wantElement:  "interface-configuration",
			wantKeyCount: 2,
		},
		{
			name:         "element with namespace prefix",
			segment:      "Cisco-IOS-XR-um-hostname-cfg:hostname",
			wantElement:  "Cisco-IOS-XR-um-hostname-cfg:hostname",
			wantKeyCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element, keys := ParseXPathSegment(tt.segment)

			if element != tt.wantElement {
				t.Errorf("ParseXPathSegment(%q) element = %q, expected %q", tt.segment, element, tt.wantElement)
			}
			if len(keys) != tt.wantKeyCount {
				t.Errorf("ParseXPathSegment(%q) key count = %d, expected %d (keys: %v)",
					tt.segment, len(keys), tt.wantKeyCount, keys)
			}
			for i, kv := range tt.wantKeys {
				if i >= len(keys) {
					break
				}
				if keys[i].Key != kv.Key || keys[i].Value != kv.Value {
					t.Errorf("ParseXPathSegment(%q) key[%d] = {%q, %q}, expected {%q, %q}",
						tt.segment, i, keys[i].Key, keys[i].Value, kv.Key, kv.Value)
				}
			}
		})
	}
}

// ============================================================================
// SplitXPathSegments Tests
// ============================================================================

func TestSplitXPathSegments(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		expected []string
	}{
		{
			name:     "simple path",
			xPath:    "interfaces/interface/ipv4",
			expected: []string{"interfaces", "interface", "ipv4"},
		},
		{
			name:     "path with single predicate",
			xPath:    "interfaces/interface[interface-name='Gi0/0/0/0']/description",
			expected: []string{"interfaces", "interface[interface-name='Gi0/0/0/0']", "description"},
		},
		{
			name:     "path with slash in predicate value",
			xPath:    "interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/mtu",
			expected: []string{"interfaces", "interface[interface-name='GigabitEthernet0/0/0/0']", "mtu"},
		},
		{
			name:     "path with composite key using and",
			xPath:    "policy-map/class[name='CM-HIGH' and type='qos']/priority/level",
			expected: []string{"policy-map", "class[name='CM-HIGH' and type='qos']", "priority", "level"},
		},
		{
			name:     "path with multiple separate predicates",
			xPath:    "policy-map/class[name='CM-HIGH'][type='qos']/priority",
			expected: []string{"policy-map", "class[name='CM-HIGH'][type='qos']", "priority"},
		},
		{
			name:     "single segment",
			xPath:    "hostname",
			expected: []string{"hostname"},
		},
		{
			name:     "empty path",
			xPath:    "",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitXPathSegments(tt.xPath)
			if len(result) != len(tt.expected) {
				t.Errorf("SplitXPathSegments(%q) returned %d segments, expected %d\ngot: %v\nwant: %v",
					tt.xPath, len(result), len(tt.expected), result, tt.expected)
				return
			}
			for i, seg := range tt.expected {
				if result[i] != seg {
					t.Errorf("SplitXPathSegments(%q) segment[%d] = %q, expected %q",
						tt.xPath, i, result[i], seg)
				}
			}
		})
	}
}

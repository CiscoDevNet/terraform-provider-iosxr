package helpers

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// ============================================================================
// NETCONF Connection Management Tests
// ============================================================================

func TestCloseNetconfConnection(t *testing.T) {
	ctx := context.Background()

	t.Run("nil client with reuse enabled", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CloseNetconfConnection panicked with nil client and reuse enabled: %v", r)
			}
		}()
		CloseNetconfConnection(ctx, nil, true)
	})

	t.Run("nil client with reuse disabled", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CloseNetconfConnection panicked with nil client and reuse disabled: %v", r)
			}
		}()
		CloseNetconfConnection(ctx, nil, false)
	})
}

func TestAcquireNetconfLock(t *testing.T) {
	tests := []struct {
		name            string
		reuseConnection bool
		isWrite         bool
		expectLock      bool
	}{
		{
			name:            "no reuse locks regardless of operation type",
			reuseConnection: false,
			isWrite:         false,
			expectLock:      true,
		},
		{
			name:            "reuse write locks",
			reuseConnection: true,
			isWrite:         true,
			expectLock:      true,
		},
		{
			name:            "reuse read does not lock",
			reuseConnection: true,
			isWrite:         false,
			expectLock:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mutex sync.Mutex
			acquired := AcquireNetconfLock(&mutex, tt.reuseConnection, tt.isWrite)

			if acquired != tt.expectLock {
				t.Errorf("AcquireNetconfLock() = %v, expected %v", acquired, tt.expectLock)
			}

			if acquired {
				mutex.Unlock()
			}
		})
	}

	t.Run("no reuse blocks concurrent caller", func(t *testing.T) {
		var mutex sync.Mutex
		acquired := AcquireNetconfLock(&mutex, false, false)
		if !acquired {
			t.Fatal("first AcquireNetconfLock() should have acquired the lock")
		}

		// A second goroutine must not be able to acquire the lock while the first holds it.
		done := make(chan bool, 1)
		go func() {
			// TryLock returns false when the mutex is already held.
			done <- !mutex.TryLock()
		}()
		blocked := <-done

		mutex.Unlock()

		if !blocked {
			t.Error("second caller should have been blocked while first held the lock")
		}
	})
}

func TestEnsureNetconfConnection(t *testing.T) {
	ctx := context.Background()

	t.Run("nil client", func(t *testing.T) {
		err := EnsureNetconfConnection(ctx, nil, false, 3)
		if err == nil {
			t.Error("EnsureNetconfConnection() with nil client should return error")
		}
		if !strings.Contains(err.Error(), "client is nil") {
			t.Errorf("EnsureNetconfConnection() error = %v, should mention nil client", err)
		}
	})

	t.Run("default maxRetries", func(t *testing.T) {
		err := EnsureNetconfConnection(ctx, nil, false, 0)
		if err == nil {
			t.Error("EnsureNetconfConnection() with nil client should return error")
		}
	})
}

// ============================================================================
// NETCONF Filter Tests
// ============================================================================

func TestGetSubtreeFilter(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		contains []string
	}{
		{
			name:  "simple xpath",
			xPath: "Cisco-IOS-XR-um-hostname-cfg:/hostname",
			contains: []string{
				"<hostname",
				"xmlns=",
				"http://cisco.com/ns/yang/Cisco-IOS-XR-um-hostname-cfg",
				"</hostname>",
			},
		},
		{
			name:  "xpath with list predicate",
			xPath: "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']",
			contains: []string{
				"<interfaces",
				"<interface",
				"<interface-name>GigabitEthernet0/0/0/0</interface-name>",
				"</interface>",
				"</interfaces>",
			},
		},
		{
			name:  "nested path",
			xPath: "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface/ipv4/address",
			contains: []string{
				"xmlns=",
				"http://cisco.com/ns/yang/Cisco-IOS-XR-um-interface-cfg",
				"<interfaces",
				"<interface",
				"<ipv4",
				"<address",
				"</address>",
				"</ipv4>",
				"</interface>",
				"</interfaces>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := GetSubtreeFilter(tt.xPath)

			if filter.Type != "subtree" {
				t.Errorf("GetSubtreeFilter() type = %q, expected \"subtree\"", filter.Type)
			}

			for _, substr := range tt.contains {
				if !strings.Contains(filter.Content, substr) {
					t.Errorf("GetSubtreeFilter() result missing substring: %q\nResult: %s", substr, filter.Content)
				}
			}
		})
	}
}

// ============================================================================
// NETCONF Response Helper Tests
// ============================================================================

func TestIsGetConfigResponseEmpty(t *testing.T) {
	tests := []struct {
		name     string
		xmlStr   string
		expected bool
	}{
		{
			name:     "nil response",
			xmlStr:   "",
			expected: true,
		},
		{
			name:     "empty data element self-closing",
			xmlStr:   `<rpc-reply><data/></rpc-reply>`,
			expected: true,
		},
		{
			name:     "empty data element with closing tag",
			xmlStr:   `<rpc-reply><data></data></rpc-reply>`,
			expected: true,
		},
		{
			name:     "data element with whitespace only",
			xmlStr:   `<rpc-reply><data>   </data></rpc-reply>`,
			expected: true,
		},
		{
			name:     "data element with content",
			xmlStr:   `<rpc-reply><data><hostname>test</hostname></data></rpc-reply>`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res *netconf.Res
			if tt.xmlStr != "" {
				res = &netconf.Res{}
				res.Res = xmldot.Get(tt.xmlStr, "rpc-reply")
			}

			result := IsGetConfigResponseEmpty(res)
			if result != tt.expected {
				t.Errorf("IsGetConfigResponseEmpty() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// ============================================================================
// NETCONF Body Manipulation Tests
// ============================================================================

// xmlGet is a test helper that queries a path in an XML string using xmldot.
func xmlGet(xml, path string) string {
	return xmldot.Get(xml, path).String()
}

func TestSetFromXPath(t *testing.T) {
	tests := []struct {
		name        string
		xPath       string
		value       interface{}
		checkPath   string // xmldot path to verify the value was placed correctly
		checkVal    string
		extraChecks map[string]string // additional path→value assertions
	}{
		{
			name:      "set value on leaf element",
			xPath:     "Cisco-IOS-XR-um-hostname-cfg:/hostname/host-name",
			value:     "test-router",
			checkPath: "hostname.host-name",
			checkVal:  "test-router",
		},
		{
			name:      "nested path with single key",
			xPath:     "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/description",
			value:     "uplink",
			checkPath: "interfaces.interface.description",
			checkVal:  "uplink",
			extraChecks: map[string]string{
				"interfaces.interface.interface-name": "GigabitEthernet0/0/0/0",
			},
		},
		{
			name:      "element with multiple predicates using and",
			xPath:     "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name='PM-QOS']/class[name='CM-HIGH' and type='qos']/priority/level",
			value:     "1",
			checkPath: "policy-map.type.qos.class.priority.level",
			checkVal:  "1",
			extraChecks: map[string]string{
				"policy-map.type.qos.class.name": "CM-HIGH",
				"policy-map.type.qos.class.type": "qos",
			},
		},
		{
			name:      "element with multiple separate predicates",
			xPath:     "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name='PM-QOS']/class[name='CM-HIGH'][type='qos']/priority/level",
			value:     "1",
			checkPath: "policy-map.type.qos.class.priority.level",
			checkVal:  "1",
			extraChecks: map[string]string{
				"policy-map.type.qos.class.name": "CM-HIGH",
				"policy-map.type.qos.class.type": "qos",
			},
		},
		{
			name:      "path with slash in predicate value",
			xPath:     "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/mtu",
			value:     "9000",
			checkPath: "interfaces.interface.mtu",
			checkVal:  "9000",
			extraChecks: map[string]string{
				"interfaces.interface.interface-name": "GigabitEthernet0/0/0/0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.NewBody("")
			result := SetFromXPath(body, tt.xPath, tt.value)
			resultXML := result.Res()

			got := xmlGet(resultXML, tt.checkPath)
			if got != tt.checkVal {
				t.Errorf("SetFromXPath() value at %q = %q, expected %q\nXML: %s",
					tt.checkPath, got, tt.checkVal, resultXML)
			}

			for path, want := range tt.extraChecks {
				if v := xmlGet(resultXML, path); v != want {
					t.Errorf("SetFromXPath() extra check at %q = %q, expected %q\nXML: %s",
						path, v, want, resultXML)
				}
			}
		})
	}
}

// TestSetFromXPathPresenceContainer verifies that an empty value creates the element
// rather than silently producing an empty body.
func TestSetFromXPathPresenceContainer(t *testing.T) {
	body := netconf.NewBody("")
	result := SetFromXPath(body, "Cisco-IOS-XR-um-hostname-cfg:/hostname", "")
	resultXML := result.Res()

	if !xmldot.Get(resultXML, "hostname").Exists() {
		t.Errorf("SetFromXPath() with empty value should create <hostname> element, got empty body\nXML: %q", resultXML)
	}
}

// TestSetFromXPathNoDuplicateElements verifies that setting a value does not produce
// both an empty placeholder element and a valued element at the same path.
func TestSetFromXPathNoDuplicateElements(t *testing.T) {
	body := netconf.NewBody("")
	body = SetFromXPath(body, "Cisco-IOS-XR-um-hostname-cfg:/hostname/host-name", "router1")
	xml := body.Res()

	// Count occurrences of <host-name — there must be exactly one
	count := strings.Count(xml, "<host-name")
	if count != 1 {
		t.Errorf("SetFromXPath() produced %d <host-name elements, expected exactly 1\nXML: %s", count, xml)
	}
}

// TestSetFromXPathMultipleListEntries verifies that different key predicates on the
// same list element produce separate sibling entries, not nested or overwritten ones.
func TestSetFromXPathMultipleListEntries(t *testing.T) {
	body := netconf.NewBody("")

	base := "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name='PM-QOS-POLICY']"
	bp1 := base + "/class[name='CM-HIGH-PRIORITY' and type='qos']"
	body = SetFromXPath(body, bp1+"/name", "CM-HIGH-PRIORITY")
	body = SetFromXPath(body, bp1+"/type", "qos")
	body = SetFromXPath(body, bp1+"/priority/level", "2")

	bp2 := base + "/class[name='CM-REAL-TIME' and type='qos']"
	body = SetFromXPath(body, bp2+"/name", "CM-REAL-TIME")
	body = SetFromXPath(body, bp2+"/type", "qos")
	body = SetFromXPath(body, bp2+"/priority/level", "3")

	xml := body.Res()

	// Find which array index each class landed in, then verify its priority is correct.
	// This catches bugs where values from one class entry bleed into the other.
	var highIdx, realIdx int = -1, -1
	for i := 0; i < 4; i++ {
		name := xmldot.Get(xml, fmt.Sprintf("policy-map.type.qos.class.%d.name", i)).String()
		switch name {
		case "CM-HIGH-PRIORITY":
			highIdx = i
		case "CM-REAL-TIME":
			realIdx = i
		}
	}

	if highIdx == -1 {
		t.Fatalf("CM-HIGH-PRIORITY class entry not found in XML: %s", xml)
	}
	if realIdx == -1 {
		t.Fatalf("CM-REAL-TIME class entry not found in XML: %s", xml)
	}

	highLevel := xmldot.Get(xml, fmt.Sprintf("policy-map.type.qos.class.%d.priority.level", highIdx)).String()
	realLevel := xmldot.Get(xml, fmt.Sprintf("policy-map.type.qos.class.%d.priority.level", realIdx)).String()

	if highLevel != "2" {
		t.Errorf("CM-HIGH-PRIORITY priority level = %q, expected \"2\"\nXML: %s", highLevel, xml)
	}
	if realLevel != "3" {
		t.Errorf("CM-REAL-TIME priority level = %q, expected \"3\"\nXML: %s", realLevel, xml)
	}
}

// TestSetFromXPathThenAppendSiblings verifies that AppendFromXPath after SetFromXPath
// on the same parent path produces sibling elements, not nested ones.
func TestSetFromXPathThenAppendSiblings(t *testing.T) {
	body := netconf.NewBody("")
	base := "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='Bundle-Ether1']"

	body = SetFromXPath(body, base+"/bundle/id", "1")
	body = AppendFromXPath(body, base+"/bundle/load-balancing/hash", "src-ip")

	xml := body.Res()

	// id and hash must both be present; id must not be nested inside hash or vice versa
	id := xmlGet(xml, "interfaces.interface.bundle.id")
	hash := xmlGet(xml, "interfaces.interface.bundle.load-balancing.hash")

	if id != "1" {
		t.Errorf("bundle/id = %q, expected \"1\"\nXML: %s", id, xml)
	}
	if hash != "src-ip" {
		t.Errorf("bundle/load-balancing/hash = %q, expected \"src-ip\"\nXML: %s", hash, xml)
	}
}

func TestRemoveFromXPath(t *testing.T) {
	tests := []struct {
		name          string
		xPath         string
		operationPath string // xmldot path where @nc:operation must equal "remove"
	}{
		{
			name:          "path without predicates",
			xPath:         "Cisco-IOS-XR-um-hostname-cfg:/hostname",
			operationPath: "hostname.@nc:operation",
		},
		{
			name:          "simple single element with key",
			xPath:         "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']",
			operationPath: "interfaces.interface.@nc:operation",
		},
		{
			name:          "nested path with single key",
			xPath:         "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/description",
			operationPath: "interfaces.interface.description.@nc:operation",
		},
		{
			name:          "interface with slashes in name",
			xPath:         "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/mtu",
			operationPath: "interfaces.interface.mtu.@nc:operation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.NewBody("")
			result := RemoveFromXPath(body, tt.xPath)
			resultXML := result.Res()

			// nc namespace must be declared
			if !strings.Contains(resultXML, "xmlns:nc=") {
				t.Errorf("RemoveFromXPath() missing xmlns:nc declaration\nXML: %s", resultXML)
			}

			// operation="remove" must be on the exact target element
			op := xmlGet(resultXML, tt.operationPath)
			if op != "remove" {
				t.Errorf("RemoveFromXPath() @nc:operation at %q = %q, expected \"remove\"\nXML: %s",
					tt.operationPath, op, resultXML)
			}
		})
	}
}

// ============================================================================
// AppendFromXPath Tests
// ============================================================================

func TestAppendFromXPath(t *testing.T) {
	t.Run("sequential appends produce sibling elements in order", func(t *testing.T) {
		body := netconf.NewBody("")
		base := "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name='RP-IN']/rpl-route-policy"
		body = AppendFromXPath(body, base, "value-a")
		body = AppendFromXPath(body, base, "value-b")
		body = AppendFromXPath(body, base, "value-c")
		xml := body.Res()

		want := []string{"value-a", "value-b", "value-c"}
		for i, w := range want {
			got := xmldot.Get(xml, fmt.Sprintf("routing-policy.route-policies.route-policy.rpl-route-policy.%d", i)).String()
			if got != w {
				t.Errorf("AppendFromXPath() entry[%d] = %q, expected %q\nXML: %s", i, got, w, xml)
			}
		}
	})

	t.Run("empty value creates presence container", func(t *testing.T) {
		body := netconf.NewBody("")
		body = AppendFromXPath(body, "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='Loopback0']/ipv4/addresses/address", "")
		xml := body.Res()

		if !strings.Contains(xml, "<address") {
			t.Errorf("AppendFromXPath() with empty value should create <address> element\nXML: %s", xml)
		}
	})

	t.Run("namespace declared on root element when module prefix present", func(t *testing.T) {
		// augmentNamespaces runs only when value is non-empty; verify xmlns is added.
		body := netconf.NewBody("")
		body = AppendFromXPath(body,
			"Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name='RP-TEST']/rpl-route-policy",
			"permit 10\n")
		xml := body.Res()

		if !strings.Contains(xml, "xmlns") {
			t.Errorf("AppendFromXPath() should declare xmlns on namespaced root element\nXML: %s", xml)
		}
		if !strings.Contains(xml, "Cisco-IOS-XR-um-route-policy-cfg") {
			t.Errorf("AppendFromXPath() xmlns should reference the module namespace\nXML: %s", xml)
		}
	})
}

// ============================================================================
// GetFromXPath Tests
// ============================================================================

func TestGetFromXPath(t *testing.T) {
	// Build a representative IOS-XR XML response to query against
	const sampleXML = `<rpc-reply>
		<data>
			<interfaces xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-um-interface-cfg">
				<interface>
					<interface-name>GigabitEthernet0/0/0/0</interface-name>
					<description>uplink-primary</description>
					<mtu>9000</mtu>
				</interface>
				<interface>
					<interface-name>GigabitEthernet0/0/0/1</interface-name>
					<description>uplink-secondary</description>
					<mtu>1500</mtu>
				</interface>
			</interfaces>
			<hostname xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-um-hostname-cfg">
				<host-name>router1</host-name>
			</hostname>
		</data>
	</rpc-reply>`

	parsed := xmldot.Get(sampleXML, "rpc-reply.data")

	tests := []struct {
		name      string
		xPath     string
		wantVal   string
		wantEmpty bool
	}{
		{
			name:    "simple path without predicates",
			xPath:   "Cisco-IOS-XR-um-hostname-cfg:/hostname/host-name",
			wantVal: "router1",
		},
		{
			name:    "filter by key returns first match",
			xPath:   "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/description",
			wantVal: "uplink-primary",
		},
		{
			name:    "filter correctly returns second match",
			xPath:   "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/1']/description",
			wantVal: "uplink-secondary",
		},
		{
			name:    "path with value containing slashes",
			xPath:   "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/mtu",
			wantVal: "9000",
		},
		{
			name:      "non-existent path",
			xPath:     "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet0/0/0/0']/ipv4",
			wantEmpty: true,
		},
		{
			name:      "filter with no matching key",
			xPath:     "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name='GigabitEthernet9/9/9/9']/description",
			wantEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFromXPath(parsed, tt.xPath)

			if tt.wantEmpty {
				if result.Exists() {
					t.Errorf("GetFromXPath(%q) expected empty result, got: %q", tt.xPath, result.String())
				}
				return
			}

			if !result.Exists() {
				t.Errorf("GetFromXPath(%q) returned no result, expected %q", tt.xPath, tt.wantVal)
				return
			}
			if got := result.String(); got != tt.wantVal {
				t.Errorf("GetFromXPath(%q) = %q, expected %q", tt.xPath, got, tt.wantVal)
			}
		})
	}
}

func TestGetConfigWithRetry(t *testing.T) {
	ctx := context.Background()

	t.Run("nil client", func(t *testing.T) {
		// This will panic because the client library dereferences nil
		defer func() {
			if r := recover(); r != nil {
				// Expected panic from client library with nil client
				t.Logf("Expected panic from client library with nil client: %v", r)
			}
		}()
		_, _, _ = GetConfigWithRetry(ctx, nil, "running", netconf.Filter{}, "/test")
	})
}

package helpers

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/netascode/go-netconf"
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
			name:            "no reuse - should lock all operations",
			reuseConnection: false,
			isWrite:         false,
			expectLock:      true,
		},
		{
			name:            "no reuse - write operation - should lock",
			reuseConnection: false,
			isWrite:         true,
			expectLock:      true,
		},
		{
			name:            "reuse - write operation - should lock",
			reuseConnection: true,
			isWrite:         true,
			expectLock:      true,
		},
		{
			name:            "reuse - read operation - should NOT lock",
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

			// If lock was acquired, unlock it
			if acquired {
				mutex.Unlock()
			}
		})
	}
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
				"<interfaces",
				"<interface",
				"<ipv4",
				"<address",
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
			name:     "empty data element - self-closing",
			xmlStr:   `<rpc-reply><data/></rpc-reply>`,
			expected: true,
		},
		{
			name:     "empty data element - with closing tag",
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
				res.Res.Raw = tt.xmlStr
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

func TestSetFromXPath(t *testing.T) {
	tests := []struct {
		name       string
		xPath      string
		value      interface{}
		wantString string
	}{
		{
			name:       "simple xpath with value",
			xPath:      "Cisco-IOS-XR-um-hostname-cfg:/hostname/host-name",
			value:      "test-router",
			wantString: "test-router",
		},
		{
			name:       "xpath with namespace prefix",
			xPath:      "Cisco-IOS-XR-um-hostname-cfg:/hostname",
			value:      "",
			wantString: "<hostname",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.NewBody("")
			result := SetFromXPath(body, tt.xPath, tt.value)
			resultXML := result.Res()

			if !strings.Contains(resultXML, tt.wantString) {
				t.Errorf("SetFromXPath() should contain %q, got: %s", tt.wantString, resultXML)
			}
		})
	}
}

func TestRemoveFromXPath(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		contains []string
	}{
		{
			name:  "remove operation with namespace",
			xPath: "Cisco-IOS-XR-um-hostname-cfg:/hostname",
			contains: []string{
				"nc:operation=\"remove\"",
				"xmlns:nc=",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.NewBody("")
			result := RemoveFromXPath(body, tt.xPath)
			resultXML := result.Res()

			for _, substr := range tt.contains {
				if !strings.Contains(resultXML, substr) {
					t.Errorf("RemoveFromXPath() should contain %q, got: %s", substr, resultXML)
				}
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

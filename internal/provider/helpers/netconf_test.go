package helpers

import (
	"testing"
)

func TestCleanupNamespacePrefixes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Fix _:operation to nc:operation",
			input:    `<element _:operation="delete"></element>`,
			expected: `<element nc:operation="delete" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0"></element>`,
		},
		{
			name:     "Fix xmlns:_xmlns and _xmlns:nc",
			input:    `<element xmlns:_xmlns="xmlns" _xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" _:operation="delete"></element>`,
			expected: `<element xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" nc:operation="delete"></element>`,
		},
		{
			name:     "Fix xmlns:_ to xmlns:nc",
			input:    `<element xmlns:_="urn:ietf:params:xml:ns:netconf:base:1.0" _:operation="delete"></element>`,
			expected: `<element xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" nc:operation="delete"></element>`,
		},
		{
			name: "Complex nested case",
			input: `<interfaces xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-um-interface-cfg">
  <interface>
    <interface-name>Loopback100</interface-name>
    <ipv6>
      <nd xmlns:_xmlns="xmlns" _xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:_="urn:ietf:params:xml:ns:netconf:base:1.0" _:operation="delete">
        <cache-limit _:operation="delete"></cache-limit>
      </nd>
    </ipv6>
  </interface>
</interfaces>`,
			expected: `<interfaces xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-um-interface-cfg" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0">
  <interface>
    <interface-name>Loopback100</interface-name>
    <ipv6>
      <nd nc:operation="delete">
        <cache-limit nc:operation="delete"></cache-limit>
      </nd>
    </ipv6>
  </interface>
</interfaces>`,
		},
	}

}

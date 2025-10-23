package client

import (
	"strings"
	"testing"
)

func TestConvertWithPath_BannerDelete(t *testing.T) {
	path := "Cisco-IOS-XR-um-banner-cfg:/banners/banner[banner-type=login]"
	jsonBody := ""
	operation := "delete"

	result, err := ConvertWithPath(path, jsonBody, operation)
	if err != nil {
		t.Fatalf("ConvertWithPath failed: %v", err)
	}

	// Expected output
	expected := []string{
		`<banners xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-um-banner-cfg">`,
		`<banner xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" nc:operation="delete">`,
		`<banner-type>login</banner-type>`,
		`</banner>`,
		`</banners>`,
	}

	// Check that all expected strings are in the result
	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected XML to contain:\n%s\n\nBut got:\n%s", exp, result)
		}
	}

	t.Logf("Generated XML:\n%s", result)
}

func TestConvertWithPath_InterfaceDelete(t *testing.T) {
	path := "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name=TenGigE0/0/0/26]"
	jsonBody := ""
	operation := "delete"

	result, err := ConvertWithPath(path, jsonBody, operation)
	if err != nil {
		t.Fatalf("ConvertWithPath failed: %v", err)
	}

	// Expected output
	expected := []string{
		`<interfaces xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-um-interface-cfg">`,
		`<interface xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" nc:operation="delete">`,
		`<interface-name>TenGigE0/0/0/26</interface-name>`,
		`</interface>`,
		`</interfaces>`,
	}

	// Check that all expected strings are in the result
	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected XML to contain:\n%s\n\nBut got:\n%s", exp, result)
		}
	}

	t.Logf("Generated XML:\n%s", result)
}

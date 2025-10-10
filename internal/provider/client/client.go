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

package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmic/pkg/api"
	target "github.com/openconfig/gnmic/pkg/api/target"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	DefaultMaxRetries         int     = 2
	DefaultBackoffMinDelay    int     = 4
	DefaultBackoffMaxDelay    int     = 60
	DefaultBackoffDelayFactor float64 = 3
	DefaultConnectTimeout             = 15 * time.Second
	GnmiTimeout                       = 15 * time.Second
)

type SetOperationType string
type ProtocolType string

const (
	Update  SetOperationType = "update"
	Replace SetOperationType = "replace"
	Delete  SetOperationType = "delete"
)

const (
	ProtocolGNMI    ProtocolType = "gnmi"
	ProtocolNETCONF ProtocolType = "netconf"
)

type Client struct {
	Devices map[string]*Device
	// Reuse connection
	ReuseConnection bool
	// Maximum number of retries
	MaxRetries int
	// Minimum delay between two retries
	BackoffMinDelay int
	// Maximum delay between two retries
	BackoffMaxDelay int
	// Backoff delay factor
	BackoffDelayFactor float64
	// Connection timeout - configurable by user
	ConnectTimeout time.Duration
	// gNMI operation timeout - configurable by user
	GnmiTimeout time.Duration
	// NETCONF client
	NetconfClient *NetconfClient
}

type Device struct {
	SetMutex *sync.Mutex
	Target   *target.Target
	Protocol ProtocolType
}

type SetOperation struct {
	Path      string
	Body      string
	Operation SetOperationType
}

func NewClient(reuseConnection bool) Client {
	devices := make(map[string]*Device)
	return Client{
		Devices:            devices,
		ReuseConnection:    reuseConnection,
		MaxRetries:         DefaultMaxRetries,
		BackoffMinDelay:    DefaultBackoffMinDelay,
		BackoffMaxDelay:    DefaultBackoffMaxDelay,
		BackoffDelayFactor: DefaultBackoffDelayFactor,
		ConnectTimeout:     DefaultConnectTimeout,
		GnmiTimeout:        GnmiTimeout,
	}
}

func (c *Client) AddTarget(ctx context.Context, device, host, username, password, certificate, key, caCertificate string, verifyCertificate, Tls bool) error {
	if !strings.Contains(host, ":") {
		host = host + ":57400"
	}

	t, err := api.NewTarget(
		api.Name(device),
		api.Address(host),
		api.Username(username),
		api.Password(password),
		api.TLSCert(certificate),
		api.TLSKey(key),
		api.TLSCA(caCertificate),
		api.SkipVerify(!verifyCertificate),
		api.Insecure(!Tls),
	)
	if err != nil {
		return fmt.Errorf("Unable to create target: %w", err)
	}

	if c.ReuseConnection {
		err = t.CreateGNMIClient(ctx)
		if err != nil {
			return fmt.Errorf("Unable to create gNMI client: %w", err)
		}
	}

	c.Devices[device] = &Device{}
	c.Devices[device].Target = t
	c.Devices[device].SetMutex = &sync.Mutex{}
	c.Devices[device].Protocol = ProtocolGNMI

	return nil
}

// AddNetconfTarget adds a NETCONF target to the client
func (c *Client) AddNetconfTarget(ctx context.Context, device, host, username, password string, port int, useTLS, skipVerify bool) error {
	// Initialize NETCONF client if not exists
	if c.NetconfClient == nil {
		c.NetconfClient = NewNetconfClient(c.ReuseConnection, c.ConnectTimeout, c.GnmiTimeout)
		c.NetconfClient.MaxRetries = c.MaxRetries
		c.NetconfClient.BackoffMinDelay = c.BackoffMinDelay
		c.NetconfClient.BackoffMaxDelay = c.BackoffMaxDelay
		c.NetconfClient.BackoffDelayFactor = c.BackoffDelayFactor
	}

	err := c.NetconfClient.AddTarget(ctx, device, host, username, password, port, useTLS, skipVerify)
	if err != nil {
		return err
	}

	// Add device entry to main client
	c.Devices[device] = &Device{
		SetMutex: &sync.Mutex{},
		Target:   nil, // NETCONF doesn't use gNMI target
		Protocol: ProtocolNETCONF,
	}

	return nil
}

// GenericXMLElement represents a generic XML element
type GenericXMLElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr          `xml:",any,attr"`
	Value    string              `xml:",chardata"`
	Children []GenericXMLElement `xml:",any"`
}

// JSONToXML converts JSON to XML with optional namespace
func JSONToXML(jsonBody string, rootElement string, namespace string) (string, error) {
	// Parse JSON
	var data interface{}
	if err := json.Unmarshal([]byte(jsonBody), &data); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Convert to XML structure
	root := convertToXML(rootElement, data)

	// Add namespace if provided
	if namespace != "" {
		root.Attrs = append(root.Attrs, xml.Attr{
			Name:  xml.Name{Local: "xmlns"},
			Value: namespace,
		})
	}

	// Marshal to XML
	output, err := xml.MarshalIndent(root, "", "        ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return string(output), nil
}

// convertToXML recursively converts JSON data to XML structure
func convertToXML(name string, data interface{}) GenericXMLElement {
	elem := GenericXMLElement{
		XMLName: xml.Name{Local: name},
	}

	switch v := data.(type) {
	case map[string]interface{}:
		// Object: create child elements
		for key, val := range v {
			child := convertToXML(key, val)
			elem.Children = append(elem.Children, child)
		}
	case []interface{}:
		// Array: create multiple elements with same name
		for _, item := range v {
			child := convertToXML(name, item)
			elem.Children = append(elem.Children, child)
		}
	case string:
		elem.Value = v
	case float64:
		elem.Value = fmt.Sprintf("%g", v)
	case bool:
		elem.Value = fmt.Sprintf("%t", v)
	case nil:
		elem.Value = ""
	default:
		elem.Value = fmt.Sprintf("%v", v)
	}

	return elem
}

// extractKeysFromPath extracts key-value pairs from path brackets
func extractKeysFromPath(pathPart string) map[string]interface{} {
	result := make(map[string]interface{})

	// Find bracket notation [key=value]
	start := strings.Index(pathPart, "[")
	if start == -1 {
		return result
	}

	end := strings.Index(pathPart[start:], "]")
	if end == -1 {
		return result
	}

	// Extract content inside brackets
	keyValuePair := pathPart[start+1 : start+end]

	// Parse key=value
	parts := strings.SplitN(keyValuePair, "=", 2)
	if len(parts) == 2 {
		result[parts[0]] = parts[1]
	}

	return result
}

// ConvertWithPath converts using Yang path format
func ConvertWithPath(path string, jsonBody string, operation string) (string, error) {
	parts := strings.Split(path, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid path format")
	}

	yangModule := parts[0]
	pathPart := strings.TrimPrefix(parts[1], "/")
	namespace := fmt.Sprintf("http://cisco.com/ns/yang/%s", yangModule)

	var data interface{}

	// For delete operation with no body, extract keys from path
	if operation == "delete" && jsonBody == "" {
		data = extractKeysFromPath(pathPart)
	} else {
		if err := json.Unmarshal([]byte(jsonBody), &data); err != nil {
			return "", fmt.Errorf("failed to parse JSON: %w", err)
		}
	}

	// Parse path and identify which element has brackets (the list element)
	originalPathParts := strings.Split(pathPart, "/")
	listElemName := ""
	if operation == "delete" {
		for _, p := range originalPathParts {
			if strings.Contains(p, "[") {
				// Extract element name before the bracket
				listElemName = p[:strings.Index(p, "[")]
				break
			}
		}
	}

	// Remove all list key notation before splitting
	cleanPath := pathPart
	for {
		start := strings.Index(cleanPath, "[")
		if start == -1 {
			break
		}
		end := strings.Index(cleanPath[start:], "]")
		if end == -1 {
			break
		}
		cleanPath = cleanPath[:start] + cleanPath[start+end+1:]
	}

	// Split cleaned path into elements
	pathElements := strings.Split(cleanPath, "/")

	// Filter out empty elements
	filteredElements := make([]string, 0, len(pathElements))
	for _, elem := range pathElements {
		if elem != "" {
			filteredElements = append(filteredElements, elem)
		}
	}
	pathElements = filteredElements

	// Build nested structure from innermost to outermost
	current := convertToXML(pathElements[len(pathElements)-1], data)

	// Wrap in parent elements if path has multiple levels
	for i := len(pathElements) - 2; i >= 0; i-- {
		parent := GenericXMLElement{
			XMLName:  xml.Name{Local: pathElements[i]},
			Children: []GenericXMLElement{current},
		}
		current = parent
	}

	// Add namespace to root element
	current.Attrs = append(current.Attrs, xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: namespace,
	})

	output, err := xml.MarshalIndent(current, "", "        ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	// For delete operations, manually inject the nc:operation attribute
	// This is done after marshaling because Go's XML encoder doesn't properly handle namespace prefixes
	if operation == "delete" && listElemName != "" {
		xmlStr := string(output)
		// Find the list element opening tag and add the NETCONF delete attributes
		searchTag := "<" + listElemName + ">"
		replaceTag := "<" + listElemName + ` xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0"` + "\n" +
			strings.Repeat(" ", 19) + `nc:operation="delete">`
		xmlStr = strings.Replace(xmlStr, searchTag, replaceTag, 1)
		return xmlStr, nil
	}

	return string(output), nil
}

func (c *Client) Set(ctx context.Context, device string, operations ...SetOperation) (*gnmi.SetResponse, error) {
	if _, ok := c.Devices[device]; !ok {
		return nil, fmt.Errorf("Device '%s' does not exist in provider configuration.", device)
	}

	// Check protocol and route accordingly
	deviceInfo := c.Devices[device]

	switch deviceInfo.Protocol {
	case ProtocolNETCONF:
		// For NETCONF, we need to convert gNMI operations to NETCONF operations
		// This handles the conversion from gNMI JSON to NETCONF XML format
		for _, op := range operations {
			var config string
			var err error // Declare 'err' before usage

			switch op.Operation {
			case Update, Replace:

				config, err = ConvertWithPath(op.Path, op.Body, "update")
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			case Delete:
				// Create delete operation XML with proper namespace
				config, err = ConvertWithPath(op.Path, "", "delete")
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			}

			tflog.Debug(ctx, fmt.Sprintf("NETCONF config XML: %s", config))

			// Use a simpler NETCONF sequence - edit directly to running config
			// Most IOS-XR systems support direct edit to running
			err = c.NetconfClient.EditConfig(ctx, device, "running", config)
			if err != nil {
				// If direct running edit fails, try candidate approach
				tflog.Debug(ctx, "Direct running config edit failed, trying candidate approach")

				// Lock candidate, edit-config, commit, unlock sequence
				lockErr := c.NetconfClient.Lock(ctx, device, "candidate")
				if lockErr != nil {
					tflog.Warn(ctx, fmt.Sprintf("Failed to lock candidate datastore: %v", lockErr))
				}

				err = c.NetconfClient.EditConfig(ctx, device, "candidate", config)
				if err != nil {
					c.NetconfClient.Unlock(ctx, device, "candidate")
					return nil, fmt.Errorf("NETCONF edit-config failed: %w", err)
				}

				err = c.NetconfClient.Commit(ctx, device)
				if err != nil {
					c.NetconfClient.Unlock(ctx, device, "candidate")
					return nil, fmt.Errorf("NETCONF commit failed: %w", err)
				}

				err = c.NetconfClient.Unlock(ctx, device, "candidate")
				if err != nil {
					tflog.Warn(ctx, fmt.Sprintf("Failed to unlock candidate datastore: %v", err))
				}
			}
		}

		// TODO
		resp := &gnmi.SetResponse{
			Timestamp: time.Now().UnixNano(),
			Response: []*gnmi.UpdateResult{
				{
					Timestamp: time.Now().UnixNano(),
					Path: &gnmi.Path{
						Origin: "netconf",
						Elem: []*gnmi.PathElem{
							{Name: "config"},
						},
					},
					Op: gnmi.UpdateResult_REPLACE,
				},
			},
		}
		return resp, nil

	case ProtocolGNMI:
		// Original gNMI implementation
		target := deviceInfo.Target
		if target == nil {
			return nil, fmt.Errorf("gNMI target is nil for device '%s'", device)
		}

		var ops []api.GNMIOption
		for _, op := range operations {
			if op.Operation == Update {
				ops = append(ops, api.Update(api.Path(op.Path), api.Value(op.Body, "json_ietf")))
				tflog.Debug(ctx, fmt.Sprintf("gNMI operations: %v", op.Body))
			} else if op.Operation == Replace {
				ops = append(ops, api.Replace(api.Path(op.Path), api.Value(op.Body, "json_ietf")))
				tflog.Debug(ctx, fmt.Sprintf("gNMI operations: %v", op.Body))
			} else if op.Operation == Delete {
				ops = append(ops, api.Delete(op.Path))
			}
		}

		setReq, err := api.NewSetRequest(ops...)
		if err != nil {
			return nil, fmt.Errorf("Failed to create set request, got error: %w", err)
		}

		var setResp *gnmi.SetResponse
		for attempts := 0; ; attempts++ {
			deviceInfo.SetMutex.Lock()
			if !c.ReuseConnection {
				err = target.CreateGNMIClient(ctx)
				if err != nil {
					deviceInfo.SetMutex.Unlock()
					if ok := c.Backoff(ctx, attempts); !ok {
						return nil, fmt.Errorf("Unable to create gNMI client: %w", err)
					} else {
						tflog.Error(ctx, fmt.Sprintf("Unable to create gNMI client: %s, retries: %v", err.Error(), attempts))
						continue
					}
				}
			}
			tCtx, cancel := context.WithTimeout(ctx, GnmiTimeout)
			defer cancel()
			tflog.Debug(ctx, fmt.Sprintf("gNMI set request: %s", setReq.String()))
			setResp, err = target.Set(tCtx, setReq)
			tflog.Debug(ctx, fmt.Sprintf("gNMI set response: %s", setResp.String()))
			deviceInfo.SetMutex.Unlock()
			if !c.ReuseConnection {
				target.Close()
			}
			if err != nil {
				if ok := c.Backoff(ctx, attempts); !ok {
					return nil, fmt.Errorf("Set request failed, got error: %w", err)
				} else {
					tflog.Error(ctx, fmt.Sprintf("gNMI set request failed: %s, retries: %v", err, attempts))
					continue
				}
			}
			break
		}
		return setResp, nil

	default:
		return nil, fmt.Errorf("Unsupported protocol for device '%s': %s", device, deviceInfo.Protocol)
	}
}

func (c *Client) Get(ctx context.Context, device, path string) (*gnmi.GetResponse, error) {
	if _, ok := c.Devices[device]; !ok {
		return nil, fmt.Errorf("Device '%s' does not exist in provider configuration.", device)
	}

	// Check protocol and route accordingly
	deviceInfo := c.Devices[device]

	switch deviceInfo.Protocol {
	case ProtocolNETCONF:
		// For NETCONF, we need to convert gNMI path to NETCONF operation
		// This is a simplified conversion - in practice, you might need more sophisticated mapping
		if c.NetconfClient == nil {
			return nil, fmt.Errorf("NETCONF client not initialized for device '%s'", device)
		}

		// Convert gNMI path to NETCONF get-config operation
		// For now, we'll do a basic get-config on running datastore
		configData, err := c.NetconfClient.GetConfig(ctx, device, "running", path)
		if err != nil {
			return nil, fmt.Errorf("NETCONF get-config failed: %w", err)
		}

		// Create a mock gNMI response for compatibility
		// Note: This is a simplified approach - in practice you'd want to parse the XML
		// and convert it to proper gNMI format
		resp := &gnmi.GetResponse{
			Notification: []*gnmi.Notification{
				{
					Timestamp: time.Now().UnixNano(),
					Update: []*gnmi.Update{
						{
							Path: &gnmi.Path{
								Origin: "netconf",
								Elem: []*gnmi.PathElem{
									{Name: "config"},
								},
							},
							Val: &gnmi.TypedValue{
								Value: &gnmi.TypedValue_StringVal{
									StringVal: configData,
								},
							},
						},
					},
				},
			},
		}
		return resp, nil

	case ProtocolGNMI:
		// Original gNMI implementation
		target := deviceInfo.Target
		if target == nil {
			return nil, fmt.Errorf("gNMI target is nil for device '%s'", device)
		}

		getReq, err := api.NewGetRequest(api.Path(path), api.Encoding("json_ietf"))
		if err != nil {
			return nil, fmt.Errorf("Failed to create get request, got error: %w", err)
		}

		var getResp *gnmi.GetResponse
		for attempts := 0; ; attempts++ {
			tflog.Debug(ctx, fmt.Sprintf("gNMI get request: %s", getReq.String()))
			if !c.ReuseConnection {
				err = target.CreateGNMIClient(ctx)
				if err != nil {
					if ok := c.Backoff(ctx, attempts); !ok {
						return nil, fmt.Errorf("Unable to create gNMI client: %w", err)
					} else {
						tflog.Error(ctx, fmt.Sprintf("Unable to create gNMI client: %s, retries: %v", err.Error(), attempts))
						continue
					}
				}
			}
			tCtx, cancel := context.WithTimeout(ctx, GnmiTimeout)
			defer cancel()
			getResp, err = target.Get(tCtx, getReq)
			if !c.ReuseConnection {
				target.Close()
			}
			tflog.Debug(ctx, fmt.Sprintf("gNMI get response: %s", getResp.String()))
			if err != nil {
				if ok := c.Backoff(ctx, attempts); !ok {
					return nil, fmt.Errorf("Get request failed, got error: %w", err)
				} else {
					tflog.Error(ctx, fmt.Sprintf("gNMI get request failed: %s, retries: %v", err, attempts))
					continue
				}
			}
			break
		}
		return getResp, nil

	default:
		return nil, fmt.Errorf("Unsupported protocol for device '%s': %s", device, deviceInfo.Protocol)
	}
}

// Backoff waits following an exponential backoff algorithm
func (c *Client) Backoff(ctx context.Context, attempts int) bool {
	tflog.Debug(ctx, fmt.Sprintf("Begining backoff method: attempts %v on %v", attempts, c.MaxRetries))
	if attempts >= c.MaxRetries {
		tflog.Debug(ctx, "Exit from backoff method with return value false")
		return false
	}

	minDelay := time.Duration(c.BackoffMinDelay) * time.Second
	maxDelay := time.Duration(c.BackoffMaxDelay) * time.Second

	min := float64(minDelay)
	backoff := min * math.Pow(c.BackoffDelayFactor, float64(attempts))
	if backoff > float64(maxDelay) {
		backoff = float64(maxDelay)
	}
	backoff = (rand.Float64()/2+0.5)*(backoff-min) + min
	backoffDuration := time.Duration(backoff)
	tflog.Debug(ctx, fmt.Sprintf("Starting sleeping for %v", backoffDuration.Round(time.Second)))
	time.Sleep(backoffDuration)
	tflog.Debug(ctx, "Exit from backoff method with return value true")
	return true
}

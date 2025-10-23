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
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmic/pkg/api"
	target "github.com/openconfig/gnmic/pkg/api/target"
)

const (
	DefaultGNMIPort    = 57400
	DefaultGNMITimeout = 15 * time.Second
)

// GNMIClient implements the Client interface for gNMI protocol
type GNMIClient struct {
	BaseClient
	Devices map[string]*GNMIDevice
	// gNMI operation timeout
	GnmiTimeout time.Duration
}

type GNMIDevice struct {
	SetMutex *sync.Mutex
	Target   *target.Target
}

func NewGNMIClient(base BaseClient) *GNMIClient {
	devices := make(map[string]*GNMIDevice)

	return &GNMIClient{
		BaseClient:  base,
		Devices:     devices,
		GnmiTimeout: DefaultGNMITimeout,
	}
}

func (gc *GNMIClient) AddTarget(ctx context.Context, device, host, username, password string, port int, certificate, key, caCertificate string, verifyCertificate, useTls bool) error {
	if port == 0 {
		port = DefaultGNMIPort
	}

	if !strings.Contains(host, ":") {
		host = fmt.Sprintf("%s:%d", host, port)
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
		api.Insecure(!useTls),
	)
	if err != nil {
		return fmt.Errorf("Unable to create gNMI target: %w", err)
	}

	if gc.ReuseConnection {
		err = t.CreateGNMIClient(ctx)
		if err != nil {
			return fmt.Errorf("Unable to create gNMI client: %w", err)
		}
	}

	gc.Devices[device] = &GNMIDevice{
		Target:   t,
		SetMutex: &sync.Mutex{},
	}

	return nil
}

func (gc *GNMIClient) Close(ctx context.Context) {
	for device, dev := range gc.Devices {
		if dev.Target != nil {
			tflog.Debug(ctx, fmt.Sprintf("Closing gNMI connection for device: %s", device))
			dev.Target.Close()
		}
	}
}

// SetWithOperations performs gNMI Set operations (Update, Replace, Delete)
func (gc *GNMIClient) SetWithOperations(ctx context.Context, device string, operations ...SetOperation) (*gnmi.SetResponse, error) {
	dev, exists := gc.Devices[device]
	if !exists {
		return nil, fmt.Errorf("Device '%s' does not exist in gNMI client configuration", device)
	}

	var ops []api.GNMIOption
	for _, op := range operations {
		switch op.Operation {
		case Update:
			ops = append(ops, api.Update(api.Path(op.Path), api.Value(op.Body, "json_ietf")))
		case Replace:
			ops = append(ops, api.Replace(api.Path(op.Path), api.Value(op.Body, "json_ietf")))
		case Delete:
			ops = append(ops, api.Delete(op.Path))
		}
	}

	setReq, err := api.NewSetRequest(ops...)
	if err != nil {
		return nil, fmt.Errorf("Failed to create gNMI set request: %w", err)
	}

	var setResp *gnmi.SetResponse
	for attempts := 0; ; attempts++ {
		dev.SetMutex.Lock()

		if !gc.ReuseConnection {
			err = dev.Target.CreateGNMIClient(ctx)
			if err != nil {
				dev.SetMutex.Unlock()
				if ok := gc.Backoff(ctx, attempts); !ok {
					return nil, fmt.Errorf("Unable to create gNMI client: %w", err)
				} else {
					tflog.Error(ctx, fmt.Sprintf("Unable to create gNMI client: %s, retries: %v", err.Error(), attempts))
					continue
				}
			}
		}

		tCtx, cancel := context.WithTimeout(ctx, gc.GnmiTimeout)
		tflog.Debug(ctx, fmt.Sprintf("gNMI set request: %s", setReq.String()))
		setResp, err = dev.Target.Set(tCtx, setReq)
		cancel()

		dev.SetMutex.Unlock()

		if !gc.ReuseConnection {
			dev.Target.Close()
		}

		if err != nil {
			if ok := gc.Backoff(ctx, attempts); !ok {
				return nil, fmt.Errorf("gNMI Set request failed: %w", err)
			} else {
				tflog.Error(ctx, fmt.Sprintf("gNMI set request failed: %s, retries: %v", err, attempts))
				continue
			}
		}

		tflog.Debug(ctx, fmt.Sprintf("gNMI set response: %s", setResp))
		break
	}

	return setResp, nil
}

// Get performs gNMI Get operations
func (gc *GNMIClient) Get(ctx context.Context, device, path string) ([]byte, error) {
	dev, exists := gc.Devices[device]
	if !exists {
		return nil, fmt.Errorf("Device '%s' does not exist in gNMI client configuration", device)
	}

	getReq, err := api.NewGetRequest(api.Path(path), api.Encoding("json_ietf"))
	if err != nil {
		return nil, fmt.Errorf("Failed to create gNMI get request: %w", err)
	}

	var getResp *gnmi.GetResponse
	for attempts := 0; ; attempts++ {
		tflog.Debug(ctx, fmt.Sprintf("gNMI get request: %s", getReq.String()))

		if !gc.ReuseConnection {
			err = dev.Target.CreateGNMIClient(ctx)
			if err != nil {
				if ok := gc.Backoff(ctx, attempts); !ok {
					return nil, fmt.Errorf("Unable to create gNMI client: %w", err)
				} else {
					tflog.Error(ctx, fmt.Sprintf("Unable to create gNMI client: %s, retries: %v", err.Error(), attempts))
					continue
				}
			}
		}

		tCtx, cancel := context.WithTimeout(ctx, gc.GnmiTimeout)
		getResp, err = dev.Target.Get(tCtx, getReq)
		cancel()

		if !gc.ReuseConnection {
			dev.Target.Close()
		}

		if err != nil {
			if ok := gc.Backoff(ctx, attempts); !ok {
				return nil, fmt.Errorf("gNMI Get request failed: %w", err)
			} else {
				tflog.Error(ctx, fmt.Sprintf("gNMI get request failed: %s, retries: %v", err, attempts))
				continue
			}
		}

		tflog.Debug(ctx, fmt.Sprintf("gNMI get response: %s", getResp.Notification[0].Update[0].Val.GetJsonIetfVal()))
		break
	}
	return getResp.Notification[0].Update[0].Val.GetJsonIetfVal(), nil
}

// Set wraps SetWithOperations for the Client interface
func (gc *GNMIClient) Set(ctx context.Context, device string, operations ...SetOperation) (bool, error) {
	_, err := gc.SetWithOperations(ctx, device, operations...)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetProtocol returns the protocol type for this client
func (gc *GNMIClient) GetProtocol() ProtocolType {
	return ProtocolGNMI
}

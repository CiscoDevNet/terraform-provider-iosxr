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
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmic/api"
	"github.com/openconfig/gnmic/target"
)

const (
	DefaultMaxRetries         int     = 2
	DefaultBackoffMinDelay    int     = 4
	DefaultBackoffMaxDelay    int     = 60
	DefaultBackoffDelayFactor float64 = 3
	GnmiTimeout                       = 30 * time.Second
)

type SetOperationType string

const (
	Update  SetOperationType = "update"
	Replace SetOperationType = "replace"
	Delete  SetOperationType = "delete"
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
}

type Device struct {
	SetMutex *sync.Mutex
	Target   *target.Target
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

	return nil
}

func (c *Client) Set(ctx context.Context, device string, operations ...SetOperation) (*gnmi.SetResponse, error) {
	if _, ok := c.Devices[device]; !ok {
		return nil, fmt.Errorf("Device '%s' does not exist in provider configuration.", device)
	}

	target := c.Devices[device].Target

	var ops []api.GNMIOption
	for _, op := range operations {
		if op.Operation == Update {
			ops = append(ops, api.Update(api.Path(op.Path), api.Value(op.Body, "json_ietf")))
		} else if op.Operation == Replace {
			ops = append(ops, api.Replace(api.Path(op.Path), api.Value(op.Body, "json_ietf")))
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
		c.Devices[device].SetMutex.Lock()
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
		tflog.Debug(ctx, fmt.Sprintf("gNMI set request: %s", setReq.String()))
		setResp, err = target.Set(tCtx, setReq)
		tflog.Debug(ctx, fmt.Sprintf("gNMI set response: %s", setResp.String()))
		c.Devices[device].SetMutex.Unlock()
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
}

func (c *Client) Get(ctx context.Context, device, path string) (*gnmi.GetResponse, error) {
	if _, ok := c.Devices[device]; !ok {
		return nil, fmt.Errorf("Device '%s' does not exist in provider configuration.", device)
	}

	target := c.Devices[device].Target

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
				return nil, fmt.Errorf(("Get request failed, got error: %w"), err)
			} else {
				tflog.Error(ctx, fmt.Sprintf("gNMI get request failed: %s, retries: %v", err, attempts))
				continue
			}
		}
		break
	}

	return getResp, nil
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

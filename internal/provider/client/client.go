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
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	DefaultMaxRetries         int     = 2
	DefaultBackoffMinDelay    int     = 4
	DefaultBackoffMaxDelay    int     = 60
	DefaultBackoffDelayFactor float64 = 3
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

// Client is the interface that both NETCONF and gNMI clients implement
type Client interface {
	AddTarget(ctx context.Context, device, host, username, password string, port int, certificate, key, caCertificate string, verifyCertificate, useTls bool) error
	Set(ctx context.Context, device string, operations ...SetOperation) (bool, error)
	Get(ctx context.Context, device, path string) ([]byte, error)
	GetProtocol() ProtocolType
	Close(ctx context.Context)
}

// BaseClient holds common configuration for protocol clients
type BaseClient struct {
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

type SetOperation struct {
	Path      string
	Body      string
	Operation SetOperationType
}

// NewClient creates a new client based on protocol type
func NewClient(protocol ProtocolType, reuseConnection bool, maxRetries int) Client {
	baseConfig := BaseClient{
		ReuseConnection:    reuseConnection,
		MaxRetries:         maxRetries,
		BackoffMinDelay:    DefaultBackoffMinDelay,
		BackoffMaxDelay:    DefaultBackoffMaxDelay,
		BackoffDelayFactor: DefaultBackoffDelayFactor,
	}

	switch protocol {
	case ProtocolNETCONF:
		return NewNetconfClient(baseConfig)
	case ProtocolGNMI:
		return NewGNMIClient(baseConfig)
	default:
		return NewGNMIClient(baseConfig)
	}
}

// Backoff waits following an exponential backoff algorithm
func (b *BaseClient) Backoff(ctx context.Context, attempts int) bool {
	tflog.Debug(ctx, fmt.Sprintf("Begining backoff method: attempts %v on %v", attempts, b.MaxRetries))
	if attempts >= b.MaxRetries {
		tflog.Debug(ctx, "Exit from backoff method with return value false")
		return false
	}

	minDelay := time.Duration(b.BackoffMinDelay) * time.Second
	maxDelay := time.Duration(b.BackoffMaxDelay) * time.Second

	minDelayFloat := float64(minDelay)
	backoff := minDelayFloat * math.Pow(b.BackoffDelayFactor, float64(attempts))
	if backoff > float64(maxDelay) {
		backoff = float64(maxDelay)
	}
	backoff = (rand.Float64()/2+0.5)*(backoff-minDelayFloat) + minDelayFloat
	backoffDuration := time.Duration(backoff)
	tflog.Debug(ctx, fmt.Sprintf("Starting sleeping for %v", backoffDuration.Round(time.Second)))
	time.Sleep(backoffDuration)
	tflog.Debug(ctx, "Exit from backoff method with return value true")
	return true
}

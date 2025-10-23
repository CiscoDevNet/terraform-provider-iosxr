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
	"github.com/nemith/netconf"
	"github.com/nemith/netconf/transport/ssh"
	gossh "golang.org/x/crypto/ssh"
)

const (
	DefaultNetconfPort        = 830
	DefaultNetconfTimeout     = 30 * time.Second
	DefaultNetconfIdleTimeout = 10 * time.Minute
)

type NetconfOperationType string

const (
	NetconfGet        NetconfOperationType = "get"
	NetconfGetConfig  NetconfOperationType = "get-config"
	NetconfEditConfig NetconfOperationType = "edit-config"
	NetconfCommit     NetconfOperationType = "commit"
	NetconfLock       NetconfOperationType = "lock"
	NetconfUnlock     NetconfOperationType = "unlock"
)

type NetconfClient struct {
	Devices map[string]*NetconfDevice
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
	// Connection timeout
	ConnectTimeout time.Duration
	// NETCONF operation timeout
	NetconfTimeout time.Duration
	// Idle timeout for connections
	IdleTimeout time.Duration
}

type NetconfDevice struct {
	SessionMutex *sync.Mutex
	Session      *netconf.Session
	Host         string
	Port         int
	Username     string
	Password     string
	UseTLS       bool
	SkipVerify   bool
}

type NetconfOperation struct {
	Type     NetconfOperationType
	Target   string // datastore target (candidate, running, etc.)
	Filter   string // XML filter for get operations
	Config   string // XML config for edit-config operations
	TestOnly bool   // for validate operations
}

func NewNetconfClient(reuseConnection bool, options ...interface{}) *NetconfClient {
	devices := make(map[string]*NetconfDevice)

	// Set default values
	client := &NetconfClient{
		Devices:            devices,
		ReuseConnection:    reuseConnection,
		MaxRetries:         DefaultMaxRetries,
		BackoffMinDelay:    DefaultBackoffMinDelay,
		BackoffMaxDelay:    DefaultBackoffMaxDelay,
		BackoffDelayFactor: DefaultBackoffDelayFactor,
		ConnectTimeout:     DefaultConnectTimeout,
		NetconfTimeout:     DefaultNetconfTimeout,
		IdleTimeout:        DefaultNetconfIdleTimeout,
	}

	return client
}

func (nc *NetconfClient) AddTarget(ctx context.Context, device, host, username, password string, port int, useTLS, skipVerify bool) error {
	if port == 0 {
		port = DefaultNetconfPort
	}

	// Ensure host includes port if not specified
	if !strings.Contains(host, ":") {
		host = fmt.Sprintf("%s:%d", host, port)
	}

	nc.Devices[device] = &NetconfDevice{
		SessionMutex: &sync.Mutex{},
		Host:         host,
		Port:         port,
		Username:     username,
		Password:     password,
		UseTLS:       useTLS,
		SkipVerify:   skipVerify,
	}

	if nc.ReuseConnection {
		err := nc.createSession(ctx, device)
		if err != nil {
			return fmt.Errorf("Unable to create NETCONF session: %w", err)
		}
	}

	return nil
}

func (nc *NetconfClient) createSession(ctx context.Context, device string) error {
	dev, exists := nc.Devices[device]
	if !exists {
		return fmt.Errorf("Device '%s' does not exist in provider configuration", device)
	}

	// Create connection with timeout
	connectCtx, cancel := context.WithTimeout(ctx, nc.ConnectTimeout)
	defer cancel()

	tflog.Debug(ctx, fmt.Sprintf("Creating NETCONF session to %s with timeout: %v", dev.Host, nc.ConnectTimeout))

	// Create SSH client config
	sshConfig := &gossh.ClientConfig{
		User: dev.Username,
		Auth: []gossh.AuthMethod{
			gossh.Password(dev.Password),
		},
		Timeout: nc.ConnectTimeout,
	}

	if dev.SkipVerify {
		sshConfig.HostKeyCallback = gossh.InsecureIgnoreHostKey()
	} else {
		// For production, you should implement proper host key verification
		sshConfig.HostKeyCallback = gossh.InsecureIgnoreHostKey()
		tflog.Warn(ctx, "Using insecure host key verification for NETCONF connection")
	}

	// Create SSH transport
	transport, err := ssh.Dial(connectCtx, "tcp", dev.Host, sshConfig)
	if err != nil {
		return fmt.Errorf("Failed to create SSH transport to %s: %w", dev.Host, err)
	}

	// Create NETCONF session using the transport
	session, err := netconf.Open(transport)
	if err != nil {
		return fmt.Errorf("Failed to create NETCONF session: %w", err)
	}

	dev.Session = session
	tflog.Debug(ctx, fmt.Sprintf("Successfully created NETCONF session to %s", dev.Host))
	return nil
}

func (nc *NetconfClient) ExecuteRPC(ctx context.Context, device string, operation NetconfOperation) (string, error) {
	dev, exists := nc.Devices[device]
	if !exists {
		return "", fmt.Errorf("Device '%s' does not exist in provider configuration", device)
	}

	var result string
	var err error

	for attempts := 0; ; attempts++ {
		dev.SessionMutex.Lock()

		if !nc.ReuseConnection || dev.Session == nil {
			err = nc.createSession(ctx, device)
			if err != nil {
				dev.SessionMutex.Unlock()
				if ok := nc.Backoff(ctx, attempts); !ok {
					return "", fmt.Errorf("Unable to create NETCONF session: %w", err)
				} else {
					tflog.Error(ctx, fmt.Sprintf("Unable to create NETCONF session: %s, retries: %v", err.Error(), attempts))
					continue
				}
			}
		}

		// Execute the operation with timeout
		opCtx, cancel := context.WithTimeout(ctx, nc.NetconfTimeout)
		result, err = nc.executeOperation(opCtx, dev.Session, operation)
		cancel()

		dev.SessionMutex.Unlock()

		if !nc.ReuseConnection && dev.Session != nil {
			dev.Session.Close(ctx)
			dev.Session = nil
		}

		if err != nil {
			if ok := nc.Backoff(ctx, attempts); !ok {
				return "", fmt.Errorf("NETCONF operation failed: %w", err)
			} else {
				tflog.Error(ctx, fmt.Sprintf("NETCONF operation failed: %s, retries: %v", err, attempts))
				continue
			}
		}
		break
	}

	return result, nil
}

func (nc *NetconfClient) executeOperation(ctx context.Context, session *netconf.Session, operation NetconfOperation) (string, error) {
	switch operation.Type {
	case NetconfGetConfig:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF get-config request on %s", operation.Target))

		// Convert string target to netconf.Datastore
		var datastore netconf.Datastore
		switch operation.Target {
		case "running":
			datastore = netconf.Running
		case "candidate":
			datastore = netconf.Candidate
		case "startup":
			datastore = netconf.Startup
		default:
			return "", fmt.Errorf("Unsupported datastore: %s", operation.Target)
		}

		reply, err := session.GetConfig(ctx, datastore)
		if err != nil {
			return "", err
		}
		tflog.Debug(ctx, fmt.Sprintf("NETCONF get-config response: %s", string(reply)))
		return string(reply), nil

	case NetconfEditConfig:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF edit-config request on %s with config: %s", operation.Target, operation.Config))

		// Convert string target to netconf.Datastore
		var datastore netconf.Datastore
		switch operation.Target {
		case "running":
			datastore = netconf.Running
		case "candidate":
			datastore = netconf.Candidate
		case "startup":
			datastore = netconf.Startup
		default:
			return "", fmt.Errorf("Unsupported datastore: %s", operation.Target)
		}

		err := session.EditConfig(ctx, datastore, operation.Config)
		if err != nil {
			return "", err
		}
		tflog.Debug(ctx, "NETCONF edit-config completed successfully")
		return "OK", nil

	case NetconfCommit:
		tflog.Debug(ctx, "NETCONF commit request")
		err := session.Commit(ctx)
		if err != nil {
			return "", err
		}
		tflog.Debug(ctx, "NETCONF commit completed successfully")
		return "OK", nil

	case NetconfLock:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF lock request on %s", operation.Target))

		// Convert string target to netconf.Datastore
		var datastore netconf.Datastore
		switch operation.Target {
		case "running":
			datastore = netconf.Running
		case "candidate":
			datastore = netconf.Candidate
		case "startup":
			datastore = netconf.Startup
		default:
			return "", fmt.Errorf("Unsupported datastore: %s", operation.Target)
		}

		err := session.Lock(ctx, datastore)
		if err != nil {
			return "", err
		}
		tflog.Debug(ctx, fmt.Sprintf("NETCONF lock on %s completed successfully", operation.Target))
		return "OK", nil

	case NetconfUnlock:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF unlock request on %s", operation.Target))

		// Convert string target to netconf.Datastore
		var datastore netconf.Datastore
		switch operation.Target {
		case "running":
			datastore = netconf.Running
		case "candidate":
			datastore = netconf.Candidate
		case "startup":
			datastore = netconf.Startup
		default:
			return "", fmt.Errorf("Unsupported datastore: %s", operation.Target)
		}

		err := session.Unlock(ctx, datastore)
		if err != nil {
			return "", err
		}
		tflog.Debug(ctx, fmt.Sprintf("NETCONF unlock on %s completed successfully", operation.Target))
		return "OK", nil

	default:
		return "", fmt.Errorf("Unsupported NETCONF operation: %s", operation.Type)
	}
}

// GetConfig retrieves configuration from a datastore
func (nc *NetconfClient) GetConfig(ctx context.Context, device, target, filter string) (string, error) {
	operation := NetconfOperation{
		Type:   NetconfGetConfig,
		Target: target,
		Filter: filter,
	}
	return nc.ExecuteRPC(ctx, device, operation)
}

// EditConfig modifies configuration in a datastore
func (nc *NetconfClient) EditConfig(ctx context.Context, device, target, config string) error {
	operation := NetconfOperation{
		Type:   NetconfEditConfig,
		Target: target,
		Config: config,
	}
	_, err := nc.ExecuteRPC(ctx, device, operation)
	return err
}

// Commit commits the candidate configuration
func (nc *NetconfClient) Commit(ctx context.Context, device string) error {
	operation := NetconfOperation{
		Type: NetconfCommit,
	}
	_, err := nc.ExecuteRPC(ctx, device, operation)
	return err
}

// Lock locks a datastore
func (nc *NetconfClient) Lock(ctx context.Context, device, target string) error {
	operation := NetconfOperation{
		Type:   NetconfLock,
		Target: target,
	}
	_, err := nc.ExecuteRPC(ctx, device, operation)
	return err
}

// Unlock unlocks a datastore
func (nc *NetconfClient) Unlock(ctx context.Context, device, target string) error {
	operation := NetconfOperation{
		Type:   NetconfUnlock,
		Target: target,
	}
	_, err := nc.ExecuteRPC(ctx, device, operation)
	return err
}

// Close closes all NETCONF sessions
func (nc *NetconfClient) Close(ctx context.Context) {
	for device, dev := range nc.Devices {
		dev.SessionMutex.Lock()
		if dev.Session != nil {
			tflog.Debug(ctx, fmt.Sprintf("Closing NETCONF session for device: %s", device))
			dev.Session.Close(ctx)
			dev.Session = nil
		}
		dev.SessionMutex.Unlock()
	}
}

// Backoff waits following an exponential backoff algorithm (reusing from the main client)
func (nc *NetconfClient) Backoff(ctx context.Context, attempts int) bool {
	// Reuse the same backoff logic from the main client
	client := Client{
		MaxRetries:         nc.MaxRetries,
		BackoffMinDelay:    nc.BackoffMinDelay,
		BackoffMaxDelay:    nc.BackoffMaxDelay,
		BackoffDelayFactor: nc.BackoffDelayFactor,
	}
	return client.Backoff(ctx, attempts)
}

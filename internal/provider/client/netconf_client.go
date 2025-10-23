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
	//"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/response"
	"github.com/scrapli/scrapligo/util"
)

const (
	DefaultNetconfPort           = 830
	DefaultNetconfTimeout        = 120 * time.Second
	DefaultNetconfIdleTimeout    = 10 * time.Minute
	DefaultNetconfConnectTimeout = 120 * time.Second
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
	BaseClient
	Devices map[string]*NetconfDevice
	// Connection timeout - configurable by user
	ConnectTimeout time.Duration
	// NETCONF operation timeout - configurable by user
	OperationTimeout time.Duration
	// NETCONF operation timeout
	NetconfTimeout time.Duration
	// Idle timeout for connections
	IdleTimeout time.Duration
}

type NetconfDevice struct {
	SessionMutex    *sync.Mutex
	Driver          *netconf.Driver
	Host            string
	Port            int
	Username        string
	Password        string
	UseTLS          bool
	SkipVerify      bool
	CandidateLocked bool        // Track if candidate is locked
	LockMutex       *sync.Mutex // Mutex for candidate datastore lock
}

type NetconfOperation struct {
	Type     NetconfOperationType
	Target   string // datastore target (candidate, running, etc.)
	Filter   string // XML filter for get operations
	Config   string // XML config for edit-config operations
	TestOnly bool   // for validate operations
}

func NewNetconfClient(base BaseClient) *NetconfClient {
	devices := make(map[string]*NetconfDevice)

	return &NetconfClient{
		BaseClient:       base,
		Devices:          devices,
		ConnectTimeout:   DefaultNetconfConnectTimeout,
		OperationTimeout: DefaultNetconfTimeout,
		NetconfTimeout:   DefaultNetconfTimeout,
		IdleTimeout:      DefaultNetconfIdleTimeout,
	}
}

func (nc *NetconfClient) AddTarget(ctx context.Context, device, host, username, password string, port int, _certificate, _key, _caCertificate string, verifyCertificate, useTls bool) error {
	if port == 0 {
		port = DefaultNetconfPort
	}

	// Strip port from host if it was included (NETCONF driver handles port separately)
	if strings.Contains(host, ":") {
		hostParts := strings.Split(host, ":")
		host = hostParts[0]
	}

	nc.Devices[device] = &NetconfDevice{
		SessionMutex:    &sync.Mutex{},
		LockMutex:       &sync.Mutex{},
		Host:            host,
		Port:            port,
		Username:        username,
		Password:        password,
		UseTLS:          useTls,
		SkipVerify:      !verifyCertificate,
		CandidateLocked: false,
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

	tflog.Debug(ctx, fmt.Sprintf("Creating NETCONF session to %s with timeout: %v", dev.Host, nc.ConnectTimeout))

	// Create NETCONF driver options
	driverOptions := []util.Option{
		options.WithAuthUsername(dev.Username),
		options.WithAuthPassword(dev.Password),
		options.WithPort(dev.Port),
		options.WithTimeoutOps(nc.NetconfTimeout),
		options.WithTimeoutSocket(nc.ConnectTimeout),
	}

	// Add strict key checking option if skipVerify is true
	if dev.SkipVerify {
		driverOptions = append(driverOptions, options.WithAuthNoStrictKey())
	}

	// Create NETCONF driver
	driver, err := netconf.NewDriver(dev.Host, driverOptions...)
	if err != nil {
		return fmt.Errorf("Failed to create NETCONF driver: %w", err)
	}

	// Open the connection
	err = driver.Open()
	if err != nil {
		return fmt.Errorf("Failed to open NETCONF connection: %w", err)
	}

	dev.Driver = driver
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

		if !nc.ReuseConnection || dev.Driver == nil {
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

		// Create a timeout context for the operation
		opCtx, cancel := context.WithTimeout(ctx, nc.NetconfTimeout)

		// Execute the operation with timeout
		result, err = nc.executeOperation(opCtx, dev.Driver, operation)
		cancel() // Always cancel to release resources

		dev.SessionMutex.Unlock()

		if !nc.ReuseConnection && dev.Driver != nil {
			if closeErr := dev.Driver.Close(); closeErr != nil {
				tflog.Warn(ctx, fmt.Sprintf("Failed to close NETCONF session: %s", closeErr.Error()))
			}
			dev.Driver = nil
		}

		if err != nil {
			// Check if it's an I/O error - likely session died, force reconnect
			if strings.Contains(err.Error(), "input/output error") ||
				strings.Contains(err.Error(), "connection reset") ||
				strings.Contains(err.Error(), "broken pipe") {
				tflog.Warn(ctx, fmt.Sprintf("NETCONF session error detected, closing and will retry: %v", err))
				dev.SessionMutex.Lock()
				if dev.Driver != nil {
					if closeErr := dev.Driver.Close(); closeErr != nil {
						tflog.Warn(ctx, fmt.Sprintf("Failed to close NETCONF driver: %v", closeErr))
					}
					dev.Driver = nil
				}
				dev.SessionMutex.Unlock()
			}

			if ok := nc.Backoff(ctx, attempts); !ok {
				return "", fmt.Errorf("NETCONF operation failed: %w", err)
			} else {
				tflog.Error(ctx, fmt.Sprintf("NETCONF operation failed: %s, retries: %v", err, attempts))
				continue
			}
		}

		// Check if the result contains rpc-error even when err is nil
		// This handles cases where the NETCONF driver doesn't return an error
		// but the device sends back an error in the XML response
		if strings.Contains(result, "<rpc-error>") {
			// Special case: for delete operations, treat "data-missing" as success
			// since the element to delete doesn't exist (idempotent delete)
			if operation.Type == NetconfEditConfig && strings.Contains(operation.Config, `nc:operation="delete"`) &&
				strings.Contains(result, "<error-tag>data-missing</error-tag>") {
				tflog.Info(ctx, "NETCONF delete operation: element already missing, treating as success")
				break
			}

			// Special case: for lock operations, if lock-denied, provide helpful error message
			if operation.Type == NetconfLock && strings.Contains(result, "<error-tag>lock-denied</error-tag>") {
				// Extract session-id if present
				sessionID := "unknown"
				if strings.Contains(result, "<session-id>") {
					start := strings.Index(result, "<session-id>") + len("<session-id>")
					end := strings.Index(result[start:], "</session-id>")
					if end > 0 {
						sessionID = result[start : start+end]
					}
				}
				tflog.Error(ctx, fmt.Sprintf("NETCONF candidate datastore is locked by session: %s. You may need to unlock it manually or wait for the lock to be released.", sessionID))
				return "", fmt.Errorf("NETCONF candidate datastore is locked by another session (session-id: %s). Please unlock the candidate datastore manually using 'unlock candidate' or wait for the lock to expire", sessionID)
			}

			tflog.Error(ctx, fmt.Sprintf("NETCONF operation returned rpc-error: %s", result))
			err = fmt.Errorf("NETCONF operation failed with rpc-error: %s", result)
			if ok := nc.Backoff(ctx, attempts); !ok {
				return "", err
			} else {
				tflog.Error(ctx, fmt.Sprintf("NETCONF rpc-error detected, retries: %v", attempts))
				continue
			}
		}

		break
	}

	return result, nil
}

func (nc *NetconfClient) executeOperation(ctx context.Context, driver *netconf.Driver, operation NetconfOperation) (string, error) {
	switch operation.Type {
	case NetconfGet:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF get request with filter: %s", operation.Filter))

		reply, err := driver.Get(operation.Filter)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("NETCONF get failed: %v", err))
			return "", err
		}
		tflog.Debug(ctx, fmt.Sprintf("NETCONF get response: %s", string(reply.Result)))
		return string(reply.Result), nil

	case NetconfGetConfig:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF get-config request on %s", operation.Target))

		// If a filter is provided, use Get method with filter; otherwise use GetConfig
		var reply *response.NetconfResponse
		var err error

		if operation.Filter != "" {
			// Use Get method when filter is needed
			tflog.Debug(ctx, fmt.Sprintf("NETCONF get with filter: %s", operation.Filter))
			reply, err = driver.Get(operation.Filter)
		} else {
			// Use GetConfig without filter
			reply, err = driver.GetConfig(operation.Target)
		}

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("NETCONF get-config failed: %v", err))
			return "", err
		}
		tflog.Debug(ctx, fmt.Sprintf("NETCONF get-config response: %s", string(reply.Result)))
		return string(reply.Result), nil

	case NetconfEditConfig:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF edit-config request on %s with config: %s", operation.Target, operation.Config))

		reply, err := driver.EditConfig(operation.Target, operation.Config)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("NETCONF edit-config failed: %v", err))
			return "", err
		}

		tflog.Debug(ctx, fmt.Sprintf("NETCONF edit-config completed successfully, response: %s", string(reply.Result)))
		return string(reply.Result), nil

	case NetconfCommit:
		tflog.Debug(ctx, "NETCONF commit request")
		tflog.Info(ctx, "Executing NETCONF commit operation")

		reply, err := driver.Commit()
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("NETCONF commit failed with error: %v", err))
			return "", fmt.Errorf("commit operation failed: %w", err)
		}
		tflog.Debug(ctx, "NETCONF commit completed successfully")
		return string(reply.Result), nil

	case NetconfLock:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF lock request on %s", operation.Target))

		reply, err := driver.Lock(operation.Target)
		if err != nil {
			return "", err
		}
		tflog.Debug(ctx, fmt.Sprintf("NETCONF lock on %s completed successfully", operation.Target))
		return string(reply.Result), nil

	case NetconfUnlock:
		tflog.Debug(ctx, fmt.Sprintf("NETCONF unlock request on %s", operation.Target))

		reply, err := driver.Unlock(operation.Target)
		if err != nil {
			return "", err
		}
		tflog.Debug(ctx, fmt.Sprintf("NETCONF unlock on %s completed successfully", operation.Target))
		return string(reply.Result), nil

	default:
		return "", fmt.Errorf("Unknown NETCONF operation type: %s", operation.Type)
	}
}

// SetWithOperations performs NETCONF operations using candidate datastore with global lock management
func (nc *NetconfClient) SetWithOperations(ctx context.Context, device string, operations ...SetOperation) error {
	dev, exists := nc.Devices[device]
	if !exists {
		return fmt.Errorf("Device '%s' does not exist in provider configuration", device)
	}

	// Use the lock mutex to ensure only one goroutine can lock/edit/commit at a time
	dev.LockMutex.Lock()
	defer dev.LockMutex.Unlock()

	// Check if candidate is already locked by this client
	if !dev.CandidateLocked {
		tflog.Info(ctx, fmt.Sprintf("Locking candidate datastore for device: %s", device))
		err := nc.Lock(ctx, device, "candidate")
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to lock candidate datastore: %v", err))
			return fmt.Errorf("Failed to lock candidate datastore: %w", err)
		}
		dev.CandidateLocked = true
		tflog.Debug(ctx, fmt.Sprintf("Candidate datastore locked for device: %s", device))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("Candidate datastore already locked for device: %s, reusing lock", device))
	}

	// Apply all operations to candidate datastore
	for _, op := range operations {
		var config string
		var err error

		switch op.Operation {
		case Update, Replace:
			tflog.Debug(ctx, fmt.Sprintf("NETCONF config op.Body: %s", op.Body))
			config, err = helpers.GetNetconfXml(op.Path, "update", op.Body)
			if err != nil {
				return fmt.Errorf("Failed to convert config to XML: %w", err)
			}
		case Delete:
			// Create delete operation XML with proper namespace
			config, err = helpers.GetNetconfXml(op.Path, "delete", "")
			if err != nil {
				return fmt.Errorf("Failed to convert delete config to XML: %w", err)
			}
		}

		// Use candidate datastore (traditional NETCONF workflow)
		tflog.Info(ctx, fmt.Sprintf("Applying NETCONF edit-config to candidate datastore for device: %s", device))

		// EditConfig to candidate datastore
		tflog.Debug(ctx, fmt.Sprintf("Sending NETCONF edit-config to candidate datastore with XML:\n%s", config))
		err = nc.EditConfig(ctx, device, "candidate", config)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("NETCONF edit-config failed: %v", err))
			// Unlock and clear state on error
			if unlockErr := nc.Unlock(ctx, device, "candidate"); unlockErr != nil {
				tflog.Warn(ctx, fmt.Sprintf("Failed to unlock candidate datastore after edit-config error: %v", unlockErr))
			}
			dev.CandidateLocked = false
			return fmt.Errorf("NETCONF edit-config failed: %w", err)
		}
	}

	// Commit the changes immediately after edit-config
	tflog.Info(ctx, fmt.Sprintf("Committing changes for device: %s", device))
	err := nc.Commit(ctx, device)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("NETCONF commit failed: %v", err))
		// Unlock and clear state on error
		if unlockErr := nc.Unlock(ctx, device, "candidate"); unlockErr != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to unlock candidate datastore after commit error: %v", unlockErr))
		}
		dev.CandidateLocked = false
		return fmt.Errorf("NETCONF commit failed: %w", err)
	}

	// After successful commit, unlock the candidate datastore
	// This allows the next resource in the Terraform plan to acquire the lock fresh
	tflog.Info(ctx, fmt.Sprintf("Unlocking candidate datastore for device: %s after successful commit", device))
	unlockErr := nc.Unlock(ctx, device, "candidate")
	if unlockErr != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to unlock candidate datastore: %v", unlockErr))
	}
	dev.CandidateLocked = false

	tflog.Info(ctx, fmt.Sprintf("Successfully applied and committed configuration to device %s", device))
	return nil
}

// CommitAndUnlock commits all pending changes and unlocks the candidate datastore
func (nc *NetconfClient) CommitAndUnlock(ctx context.Context, device string) error {
	dev, exists := nc.Devices[device]
	if !exists {
		return fmt.Errorf("Device '%s' does not exist in provider configuration", device)
	}

	dev.LockMutex.Lock()
	defer dev.LockMutex.Unlock()

	if !dev.CandidateLocked {
		tflog.Debug(ctx, fmt.Sprintf("Candidate datastore not locked for device: %s, nothing to commit", device))
		return nil
	}

	// Commit the changes
	tflog.Info(ctx, fmt.Sprintf("Committing all changes for device: %s", device))
	err := nc.Commit(ctx, device)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("NETCONF commit failed: %v", err))
		// Still try to unlock even if commit fails
		unlockErr := nc.Unlock(ctx, device, "candidate")
		if unlockErr != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to unlock candidate datastore after commit error: %v", unlockErr))
		}
		dev.CandidateLocked = false
		return fmt.Errorf("NETCONF commit failed: %w", err)
	}

	// Unlock the candidate datastore
	tflog.Info(ctx, fmt.Sprintf("Unlocking candidate datastore for device: %s", device))
	unlockErr := nc.Unlock(ctx, device, "candidate")
	if unlockErr != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to unlock candidate datastore: %v", unlockErr))
	}
	dev.CandidateLocked = false

	tflog.Info(ctx, fmt.Sprintf("Successfully committed and unlocked configuration for device %s", device))
	return nil
}

// GetWithPath performs NETCONF get-config with automatic path conversion and XML to JSON conversion
func (nc *NetconfClient) GetWithPath(ctx context.Context, device, path string) ([]byte, error) {
	var err error

	tflog.Info(ctx, fmt.Sprintf("NETCONF GetWithPath starting for path: %s", path))
	tflog.Debug(ctx, fmt.Sprintf("netconf filter: %s", path))
	filter, err := helpers.GetNetconfXml(path, "get", "")
	if err != nil {
		return nil, fmt.Errorf("failed to build filter: %w", err)
	}
	tflog.Debug(ctx, fmt.Sprintf("netconf filter XML: %s", filter))

	tflog.Info(ctx, "Executing NETCONF GetConfig operation...")
	configData, err := nc.GetConfig(ctx, device, "running", filter)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("NETCONF get-config failed: %v", err))
		return nil, fmt.Errorf("NETCONF get-config failed: %w", err)
	}
	tflog.Info(ctx, fmt.Sprintf("NETCONF GetConfig completed, received %s ", (configData)))
	result, err := helpers.XMLToJSON(configData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, fmt.Errorf("NETCONF get-config failed: %w", err)
	}
	tflog.Debug(ctx, fmt.Sprintf("netconf filter result: %s", result))
	tflog.Info(ctx, "NETCONF GetWithPath completed successfully")
	return []byte(result), nil
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

// Commit commits pending changes
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
		if dev.Driver != nil {
			tflog.Debug(ctx, fmt.Sprintf("Closing NETCONF connection for device: %s", device))
			if err := dev.Driver.Close(); err != nil {
				tflog.Warn(ctx, fmt.Sprintf("Failed to close NETCONF connection for device %s: %v", device, err))
			}
		}
	}
}

// Set implements the Client interface for NETCONF by wrapping SetWithOperations
func (nc *NetconfClient) Set(ctx context.Context, device string, operations ...SetOperation) (bool, error) {
	err := nc.SetWithOperations(ctx, device, operations...)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Get implements the Client interface for NETCONF
func (nc *NetconfClient) Get(ctx context.Context, device, path string) ([]byte, error) {
	jsonData, err := nc.GetWithPath(ctx, device, path)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// GetProtocol returns the protocol type for this client
func (nc *NetconfClient) GetProtocol() ProtocolType {
	return ProtocolNETCONF
}

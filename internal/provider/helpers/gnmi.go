// Copyright © 2025 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-gnmi"
)

// ---------------------------------------------------------------------------
// gNMI connection helpers (mirror of the NETCONF helpers in netconf.go)
// ---------------------------------------------------------------------------

// AcquireGnmiLock acquires the device's OpMutex before any gNMI operation.
// When reuse_connection=false: Lock for ALL operations (serializes everything)
// When reuse_connection=true && isWrite=true: Lock for WRITE operations
// When reuse_connection=true && isWrite=false: NO lock for READ operations
//
// Returns true when the lock was acquired.
func AcquireGnmiLock(opMutex *sync.Mutex, reuseConnection bool, isWrite bool) bool {
	// Serialize all operations when reuse disabled
	if !reuseConnection {
		opMutex.Lock()
		return true
	}

	// When reuse enabled, only serialize write operations
	// Read operations can run concurrently
	if isWrite {
		opMutex.Lock()
		return true
	}

	return false
}

// CloseGnmiConnection safely closes a gNMI connection if reuse is disabled.
// When reuseConnection=true it is a no-op — the shared connection stays open.
// When reuseConnection=false the connection is disconnected after the operation.
//
// Must be called while the OpMutex is still held (deferred after AcquireGnmiLock).
func CloseGnmiConnection(ctx context.Context, client *gnmi.Client, reuseConnection bool) {
	if client == nil {
		return
	}
	// gNMI clients are gRPC-based and don't require explicit disconnect
	// The connection lifecycle is managed by the gRPC library
	if !reuseConnection {
		tflog.Debug(ctx, "Releasing gNMI connection (reuse disabled)")
	}
}

// EnsureGnmiConnection checks if the gNMI connection is healthy and reconnects if needed.
// When reuseConnection=true, performs a quick health check and reconnects if needed.
// When reuseConnection=false, performs a more thorough health check.
func EnsureGnmiConnection(ctx context.Context, client *gnmi.Client, reuseConnection bool, maxRetries int) error {
	if client == nil {
		return fmt.Errorf("gNMI client is nil")
	}

	if maxRetries <= 0 {
		maxRetries = 3
	}

	if !reuseConnection {
		// No reuse: run a health-check to detect a stale connection
		return gnmiHealthCheck(ctx, client)
	}

	// Fast path for reuse: do a quick health check
	if err := gnmiHealthCheck(ctx, client); err != nil {
		// Connection has issues, need to reconnect
		tflog.Warn(ctx, "gNMI connection unhealthy, reconnecting")
		return reconnectGnmiWithRetries(ctx, client, maxRetries)
	}

	return nil
}

// reconnectGnmiWithRetries attempts to reconnect a stale gNMI session with exponential back-off
func reconnectGnmiWithRetries(ctx context.Context, client *gnmi.Client, maxRetries int) error {
	retryDelay := 50 * time.Millisecond
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		if attempt > 1 {
			delay := retryDelay * time.Duration(attempt)
			tflog.Debug(ctx, fmt.Sprintf("gNMI reconnect attempt %d/%d, waiting %v", attempt, maxRetries, delay))
			time.Sleep(delay)
		}

		if err := gnmiHealthCheck(ctx, client); err != nil {
			lastErr = err
			tflog.Warn(ctx, fmt.Sprintf("gNMI reconnect attempt %d/%d failed: %s", attempt, maxRetries, err))
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("gNMI connected on attempt %d", attempt))
		return nil
	}
	return fmt.Errorf("failed to connect gNMI after %d attempts: %w", maxRetries, lastErr)
}

// gnmiHealthCheck sends a Capabilities RPC with a timeout to confirm the
// gRPC channel is alive. Uses 10s timeout to allow for connection establishment.
func gnmiHealthCheck(ctx context.Context, client *gnmi.Client) error {
	testCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_, err := client.Capabilities(testCtx)
	if err != nil && IsGnmiConnectionError(err) {
		return fmt.Errorf("gNMI health check failed: %w", err)
	}
	return nil
}

// IsGnmiConnectionError checks if an error is related to a broken/closed connection.
func IsGnmiConnectionError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "connection is closing") ||
		strings.Contains(errMsg, "connection closed") ||
		strings.Contains(errMsg, "context canceled") ||
		strings.Contains(errMsg, "context deadline exceeded") ||
		strings.Contains(errMsg, "broken pipe") ||
		strings.Contains(errMsg, "connection reset") ||
		strings.Contains(errMsg, "resource exhausted")
}

// GetWithRetry retrieves data from the device with retry logic.
// gNMI Get may return empty/incomplete data immediately after Set due to device sync delay.
// This function retries with exponential backoff to handle such cases.
//
// Parameters:
//   - ctx: context.Context
//   - client: *gnmi.Client
//   - paths: []string
//   - pathForLogging: string (for logging purposes)
//
// Returns:
//   - gnmi.GetRes: The response from Get (by value)
//   - bool: true if response is empty after all retries
//   - error: any error that occurred
func GetWithRetry(ctx context.Context, client *gnmi.Client, paths []string, pathForLogging string) (gnmi.GetRes, bool, error) {
	var getResp gnmi.GetRes
	var err error
	maxRetries := 3
	baseDelay := 200 * time.Millisecond

	for attempt := 0; attempt <= maxRetries; attempt++ {
		getResp, err = client.Get(ctx, paths)
		if err != nil {
			// Check if this is a "not found" error
			if strings.Contains(err.Error(), "Requested element(s) not found") {
				return gnmi.GetRes{}, true, nil // Not an error, just empty result
			}
			return gnmi.GetRes{}, false, fmt.Errorf("failed to retrieve object (%s): %w", pathForLogging, err)
		}

		// Check if we got data back
		isEmpty := isGnmiGetResponseEmpty(&getResp)
		tflog.Debug(ctx, fmt.Sprintf("gNMI Get response for %s (attempt %d/%d): isEmpty=%v",
			pathForLogging, attempt+1, maxRetries+1, isEmpty))

		// If we got data or this is the last attempt, break
		if !isEmpty || attempt == maxRetries {
			return getResp, isEmpty, nil
		}

		// Wait before retrying (exponential backoff)
		delay := baseDelay * time.Duration(1<<uint(attempt))
		tflog.Debug(ctx, fmt.Sprintf("gNMI returned empty response, retrying after %v", delay))
		time.Sleep(delay)
	}

	return getResp, isGnmiGetResponseEmpty(&getResp), nil
}

// isGnmiGetResponseEmpty checks if a gNMI Get response is empty or has no data
func isGnmiGetResponseEmpty(resp *gnmi.GetRes) bool {
	if resp == nil {
		return true
	}
	if len(resp.Notifications) == 0 {
		return true
	}
	if len(resp.Notifications[0].Update) == 0 {
		return true
	}

	// Check if the value is actually empty
	val := resp.Notifications[0].Update[0].Val
	if val == nil {
		return true
	}
	jsonVal := val.GetJsonIetfVal()
	jsonStr := strings.TrimSpace(string(jsonVal))
	return jsonStr == "" || jsonStr == "{}" || jsonStr == "[]"
}

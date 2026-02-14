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
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-gnmi"
)

// CloseGnmiConnection safely closes a gNMI connection if reuse is disabled.
// If reuse is enabled, the connection is kept open for subsequent operations.
//
// Usage:
//
//	defer helpers.CloseGnmiConnection(ctx, device.GnmiClient, device.ReuseConnection)
func CloseGnmiConnection(ctx context.Context, client *gnmi.Client, reuseConnection bool) {
	// Check if client is nil before attempting to disconnect
	if client == nil {
		return
	}

	// Only disconnect if connection reuse is disabled
	// When reuse is enabled, keep the connection open for subsequent operations
	if !reuseConnection {
		if err := client.Disconnect(); err != nil {
			// Log error but don't fail - connection cleanup is best-effort
			tflog.Warn(ctx, fmt.Sprintf("Failed to close gNMI connection: %s", err))
		} else {
			tflog.Debug(ctx, "Successfully closed gNMI connection")
		}
	}
}

// EnsureGnmiConnection checks if the gNMI connection is healthy and reconnects if needed.
// This is important when connection reuse is enabled to handle stale connections.
//
// Returns error if connection is not ready or cannot be established.
//
// Usage (before any gNMI operation):
//
//	if err := helpers.EnsureGnmiConnection(ctx, device.GnmiClient, reuseConnection); err != nil {
//	    resp.Diagnostics.AddError("Connection Error", fmt.Sprintf("Failed to ensure gNMI connection: %s", err))
//	    return
//	}
func EnsureGnmiConnection(ctx context.Context, client *gnmi.Client, reuseConnection bool) error {
	if client == nil {
		return fmt.Errorf("gNMI client is nil")
	}

	// When reuse is disabled, connections are created fresh for each operation
	// No need to check/reconnect
	if !reuseConnection {
		return nil
	}

	// Test connection health with a simple capabilities request with timeout
	// Create a context with timeout to prevent hanging on stale connections
	// Use shorter timeout (3 seconds) to quickly detect stale connections
	testCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := client.Capabilities(testCtx)
	if err != nil {
		// Connection appears stale/broken
		if IsGnmiConnectionError(err) {
			tflog.Warn(ctx, fmt.Sprintf("gNMI connection is stale/broken: %s", err))
			// Force disconnect the stale connection
			_ = client.Disconnect() // Best effort

			return fmt.Errorf("gNMI connection is stale: capabilities request failed: %w", err)
		}
		// Other errors might be transient, log but continue
		tflog.Debug(ctx, fmt.Sprintf("gNMI capabilities check returned error (continuing): %s", err))
	}

	return nil
}

// ForceCloseGnmiConnection forcefully closes a gNMI connection regardless of reuse setting.
// This is useful for cleaning up stale connections.
func ForceCloseGnmiConnection(ctx context.Context, client *gnmi.Client) {
	if client == nil {
		return
	}

	if err := client.Disconnect(); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to force close gNMI connection: %s", err))
	} else {
		tflog.Debug(ctx, "Successfully force closed gNMI connection")
	}
}

// IsGnmiConnectionError checks if an error is related to a broken/closed connection
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

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
	if reuseConnection {
		return // Keep connection open for reuse
	}

	// Check if client is nil before attempting to disconnect
	if client == nil {
		return
	}

	// Close the connection
	if err := client.Disconnect(); err != nil {
		// Log error but don't fail - connection cleanup is best-effort
		tflog.Warn(ctx, fmt.Sprintf("Failed to close gNMI connection: %s", err))
	}
}

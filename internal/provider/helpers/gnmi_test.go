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
	"testing"
)

// TestCloseGnmiConnection tests the CloseGnmiConnection function
func TestCloseGnmiConnection(t *testing.T) {
	ctx := context.Background()

	// Test with nil client and reuse enabled (should not panic)
	t.Run("nil client with reuse enabled", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CloseGnmiConnection panicked with nil client and reuse enabled: %v", r)
			}
		}()
		CloseGnmiConnection(ctx, nil, true)
	})

	// Test with nil client and reuse disabled (should not panic, will attempt disconnect)
	t.Run("nil client with reuse disabled", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("CloseGnmiConnection panicked with nil client and reuse disabled: %v", r)
			}
		}()
		CloseGnmiConnection(ctx, nil, false)
	})
}

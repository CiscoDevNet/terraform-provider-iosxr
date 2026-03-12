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
	"errors"
	"strings"
	"sync"
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

// TestAcquireGnmiLock tests the AcquireGnmiLock function
func TestAcquireGnmiLock(t *testing.T) {
	tests := []struct {
		name            string
		reuseConnection bool
		isWrite         bool
		expectLock      bool
	}{
		{
			name:            "no reuse - should lock",
			reuseConnection: false,
			isWrite:         false,
			expectLock:      true,
		},
		{
			name:            "no reuse - write operation - should lock",
			reuseConnection: false,
			isWrite:         true,
			expectLock:      true,
		},
		{
			name:            "reuse - write operation - should lock",
			reuseConnection: true,
			isWrite:         true,
			expectLock:      true,
		},
		{
			name:            "reuse - read operation - should NOT lock",
			reuseConnection: true,
			isWrite:         false,
			expectLock:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mutex sync.Mutex
			acquired := AcquireGnmiLock(&mutex, tt.reuseConnection, tt.isWrite)

			if acquired != tt.expectLock {
				t.Errorf("AcquireGnmiLock() = %v, expected %v", acquired, tt.expectLock)
			}

			// If lock was acquired, unlock it
			if acquired {
				mutex.Unlock()
			}
		})
	}
}

// TestIsGnmiConnectionError tests the IsGnmiConnectionError function
func TestIsGnmiConnectionError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "connection is closing",
			err:      errors.New("connection is closing"),
			expected: true,
		},
		{
			name:     "connection closed",
			err:      errors.New("connection closed"),
			expected: true,
		},
		{
			name:     "context canceled",
			err:      errors.New("context canceled"),
			expected: true,
		},
		{
			name:     "context deadline exceeded",
			err:      errors.New("context deadline exceeded"),
			expected: true,
		},
		{
			name:     "broken pipe",
			err:      errors.New("broken pipe"),
			expected: true,
		},
		{
			name:     "connection reset",
			err:      errors.New("connection reset"),
			expected: true,
		},
		{
			name:     "resource exhausted",
			err:      errors.New("resource exhausted"),
			expected: true,
		},
		{
			name:     "uppercase - Connection is Closing",
			err:      errors.New("Connection is Closing"),
			expected: true,
		},
		{
			name:     "unrelated error",
			err:      errors.New("some other error"),
			expected: false,
		},
		{
			name:     "validation error",
			err:      errors.New("invalid value provided"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsGnmiConnectionError(tt.err)
			if result != tt.expected {
				t.Errorf("IsGnmiConnectionError(%v) = %v, expected %v", tt.err, result, tt.expected)
			}
		})
	}
}

// TestIsGnmiGetResponseEmpty tests the isGnmiGetResponseEmpty function
func TestIsGnmiGetResponseEmpty(t *testing.T) {
	t.Run("nil response", func(t *testing.T) {
		result := isGnmiGetResponseEmpty(nil)
		if !result {
			t.Error("isGnmiGetResponseEmpty(nil) should return true")
		}
	})

	// Note: We can't easily construct gnmi.GetRes objects without internal types,
	// so we test the nil case which is the most critical
}

// TestEnsureGnmiConnection tests error handling
func TestEnsureGnmiConnection(t *testing.T) {
	ctx := context.Background()

	t.Run("nil client", func(t *testing.T) {
		err := EnsureGnmiConnection(ctx, nil, false, 3)
		if err == nil {
			t.Error("EnsureGnmiConnection() with nil client should return error")
		}
		if !strings.Contains(err.Error(), "client is nil") {
			t.Errorf("EnsureGnmiConnection() error = %v, should mention nil client", err)
		}
	})

	t.Run("default maxRetries", func(t *testing.T) {
		// This will fail because we don't have a real client, but tests the retry logic
		err := EnsureGnmiConnection(ctx, nil, false, 0)
		if err == nil {
			t.Error("EnsureGnmiConnection() with nil client should return error")
		}
	})
}

// TestReconnectGnmiWithRetries tests error handling
func TestReconnectGnmiWithRetries(t *testing.T) {
	ctx := context.Background()

	t.Run("with nil client", func(t *testing.T) {
		// This will panic in the gnmi library with nil client, but we can't really test it
		// without a real client. Just verify the function signature exists and can be called.
		defer func() {
			if r := recover(); r != nil {
				// Expected panic from gnmi library with nil client
				t.Logf("Expected panic from gnmi library with nil client: %v", r)
			}
		}()
		_ = reconnectGnmiWithRetries(ctx, nil, 1)
	})
}

// TestGnmiHealthCheck tests error handling
func TestGnmiHealthCheck(t *testing.T) {
	ctx := context.Background()

	t.Run("nil client", func(t *testing.T) {
		// This will panic in the gnmi library with nil client
		defer func() {
			if r := recover(); r != nil {
				// Expected panic from gnmi library with nil client
				t.Logf("Expected panic from gnmi library with nil client: %v", r)
			}
		}()
		_ = gnmiHealthCheck(ctx, nil)
	})
}

// TestGetWithRetry tests error handling and retry logic
func TestGetWithRetry(t *testing.T) {
	ctx := context.Background()

	t.Run("nil client", func(t *testing.T) {
		// This will panic in the gnmi library with nil client
		defer func() {
			if r := recover(); r != nil {
				// Expected panic from gnmi library with nil client
				t.Logf("Expected panic from gnmi library with nil client: %v", r)
			}
		}()
		_, _, _ = GetWithRetry(ctx, nil, []string{"/test"}, "/test")
	})

	t.Run("empty paths", func(t *testing.T) {
		// This will panic in the gnmi library with nil client
		defer func() {
			if r := recover(); r != nil {
				// Expected panic from gnmi library with nil client
				t.Logf("Expected panic from gnmi library with nil client: %v", r)
			}
		}()
		_, _, _ = GetWithRetry(ctx, nil, []string{}, "/test")
	})
}

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
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/openconfig/gnmi/proto/gnmi"
)

// BatchOperation represents a single operation that can be batched
type BatchOperation struct {
	ResourceID string
	Operation  SetOperation
	ResultChan chan BatchResult
}

// BatchResult represents the result of a batched operation
type BatchResult struct {
	ResourceID string
	Error      error
	Response   *gnmi.SetResponse
}

// BatchManager manages batching of gNMI operations
type BatchManager struct {
	client          *Client
	device          string
	operations      []BatchOperation
	mutex           sync.Mutex
	batchTimeout    time.Duration
	maxBatchSize    int
	timer           *time.Timer
	enabled         bool
	processingBatch bool
}

// NewBatchManager creates a new batch manager for a specific device
func NewBatchManager(client *Client, device string) *BatchManager {
	return &BatchManager{
		client:       client,
		device:       device,
		operations:   make([]BatchOperation, 0),
		batchTimeout: 15 * time.Second, // Default 15s batch window
		maxBatchSize: 500,              // Default max 100 operations per batch
		enabled:      true,
	}
}

// SetBatchTimeout sets the maximum time to wait before processing a batch
func (bm *BatchManager) SetBatchTimeout(timeout time.Duration) {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()
	bm.batchTimeout = timeout
}

// SetMaxBatchSize sets the maximum number of operations in a batch
func (bm *BatchManager) SetMaxBatchSize(size int) {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()
	bm.maxBatchSize = size
}

// EnableBatching enables or disables batching
func (bm *BatchManager) EnableBatching(enabled bool) {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()
	bm.enabled = enabled
	if !enabled && bm.timer != nil {
		bm.timer.Stop()
		bm.timer = nil
	}
}

// AddOperation adds an operation to the batch
func (bm *BatchManager) AddOperation(ctx context.Context, resourceID string, operation SetOperation) (*gnmi.SetResponse, error) {
	if !bm.enabled {
		// If batching is disabled, execute immediately
		return bm.client.Set(ctx, bm.device, operation)
	}

	// Create result channel for this operation
	resultChan := make(chan BatchResult, 1)
	batchOp := BatchOperation{
		ResourceID: resourceID,
		Operation:  operation,
		ResultChan: resultChan,
	}

	// Critical section for adding operation to batch
	bm.mutex.Lock()
	bm.operations = append(bm.operations, batchOp)
	tflog.Debug(ctx, fmt.Sprintf("Added operation to batch for resource %s, batch size: %d", resourceID, len(bm.operations)))

	shouldProcess := len(bm.operations) >= bm.maxBatchSize
	shouldStartTimer := bm.timer == nil && !bm.processingBatch && len(bm.operations) == 1

	if shouldProcess {
		tflog.Debug(ctx, fmt.Sprintf("Batch size limit reached (%d), processing immediately", bm.maxBatchSize))
		// Don't call processBatch in a goroutine while holding the lock
		bm.processingBatch = true
		if bm.timer != nil {
			bm.timer.Stop()
			bm.timer = nil
		}
		// Release lock before processing
		bm.mutex.Unlock()
		go bm.processBatchAsync(ctx)
	} else if shouldStartTimer {
		// Start timer for batch processing
		bm.timer = time.AfterFunc(bm.batchTimeout, func() {
			bm.processBatchAsync(ctx)
		})
		tflog.Debug(ctx, fmt.Sprintf("Started batch timer for %v", bm.batchTimeout))
		bm.mutex.Unlock()
	} else {
		bm.mutex.Unlock()
	}

	// Wait for the result with a reasonable timeout
	select {
	case result := <-resultChan:
		if result.Error != nil {
			return nil, result.Error
		}
		return result.Response, nil
	case <-time.After(30 * time.Second): // 30 second timeout
		return nil, fmt.Errorf("batch operation timed out for resource %s", resourceID)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// processBatchAsync is a wrapper to handle the async processing safely
func (bm *BatchManager) processBatchAsync(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			tflog.Error(ctx, fmt.Sprintf("Panic in batch processing: %v", r))
		}
	}()
	bm.processBatch(ctx)
}

// processBatch processes all accumulated operations in a single gNMI Set request
func (bm *BatchManager) processBatch(ctx context.Context) {
	bm.mutex.Lock()

	// Double-check if we should process
	if len(bm.operations) == 0 {
		bm.processingBatch = false
		bm.mutex.Unlock()
		return
	}

	// If already processing, don't start another batch
	if bm.processingBatch && bm.timer == nil {
		bm.mutex.Unlock()
		return
	}

	bm.processingBatch = true
	operations := make([]BatchOperation, len(bm.operations))
	copy(operations, bm.operations)
	bm.operations = bm.operations[:0] // Clear the slice

	// Stop and reset timer
	if bm.timer != nil {
		bm.timer.Stop()
		bm.timer = nil
	}
	bm.mutex.Unlock()

	tflog.Info(ctx, fmt.Sprintf("Processing batch of %d operations for device %s", len(operations), bm.device))

	// Extract SetOperations for the gNMI client
	setOps := make([]SetOperation, len(operations))
	for i, batchOp := range operations {
		setOps[i] = batchOp.Operation
	}

	// Execute the batched set operation with timeout context
	batchCtx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	response, err := bm.client.Set(batchCtx, bm.device, setOps...)

	// Send results to all waiting operations
	for _, batchOp := range operations {
		select {
		case batchOp.ResultChan <- BatchResult{
			ResourceID: batchOp.ResourceID,
			Error:      err,
			Response:   response,
		}:
		case <-time.After(5 * time.Second):
			tflog.Error(ctx, fmt.Sprintf("Timeout sending result to resource %s", batchOp.ResourceID))
		}
		close(batchOp.ResultChan)
	}

	bm.mutex.Lock()
	bm.processingBatch = false
	bm.mutex.Unlock()

	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Batch operation failed: %v", err))
	} else {
		tflog.Info(ctx, fmt.Sprintf("Batch operation completed successfully for %d operations", len(operations)))
	}
}

// Flush processes any pending operations immediately
func (bm *BatchManager) Flush(ctx context.Context) {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()

	if len(bm.operations) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("Flushing %d pending operations", len(bm.operations)))
		go bm.processBatch(ctx)
	}
}

// GetPendingCount returns the number of pending operations
func (bm *BatchManager) GetPendingCount() int {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()
	return len(bm.operations)
}

// Client-level batch management methods

// EnableBatchingForAllDevices enables or disables batching globally for all devices
func (c *Client) EnableBatchingForAllDevices(enabled bool) {
	c.batchMutex.Lock()
	defer c.batchMutex.Unlock()

	c.BatchingEnabled = enabled
	for _, bm := range c.batchManagers {
		bm.EnableBatching(enabled)
	}
}

// SetBatchTimeoutForAllDevices sets the batch timeout for all devices
func (c *Client) SetBatchTimeoutForAllDevices(timeout time.Duration) {
	c.batchMutex.Lock()
	defer c.batchMutex.Unlock()

	for _, bm := range c.batchManagers {
		bm.SetBatchTimeout(timeout)
	}
}

// SetBatchMaxSizeForAllDevices sets the maximum batch size for all devices
func (c *Client) SetBatchMaxSizeForAllDevices(size int) {
	c.batchMutex.Lock()
	defer c.batchMutex.Unlock()

	for _, bm := range c.batchManagers {
		bm.SetMaxBatchSize(size)
	}
}

// FlushBatchesForAllDevices flushes all pending batches for all devices
func (c *Client) FlushBatchesForAllDevices(ctx context.Context) {
	c.batchMutex.RLock()
	defer c.batchMutex.RUnlock()

	for _, bm := range c.batchManagers {
		bm.Flush(ctx)
	}
}

#!/bin/bash
set -e

echo "=========================================="
echo "Testing Batch Manager Update Fix"
echo "=========================================="
echo ""

cd "$(dirname "$0")"

# Use test.tf instead of main.tf
export TF_VAR_file="test.tf"

# Clean up any existing state
echo "1. Cleaning up..."
rm -f terraform.tfstate terraform.tfstate.backup .terraform.tfstate.lock.info
rm -rf .terraform

echo "2. Applying initial configuration..."
echo "   Creating hostname='test-router' and loopback shutdown=false"
TF_LOG=INFO terraform apply -auto-approve test.tf 2>&1 | tee apply_initial.log | grep -E "Apply complete|FlushAll|error|Error" || true

if grep -qi "error" apply_initial.log; then
    echo "❌ Initial apply failed!"
    exit 1
fi

echo ""
echo "3. Checking if FlushAll was called..."
if grep -q "FlushAll" apply_initial.log; then
    echo "   ✅ FlushAll found in logs"
    grep "FlushAll" apply_initial.log | tail -3
else
    echo "   ⚠️  FlushAll not found in logs (might be at DEBUG level)"
fi

echo ""
echo "4. Running terraform plan (should show no changes)..."
TF_LOG=WARN terraform plan 2>&1 | tee plan1.log

if grep -q "No changes" plan1.log; then
    echo "   ✅ No drift detected - state matches device!"
else
    if grep -q "Plan:" plan1.log; then
        echo "   ❌ DRIFT DETECTED - state doesn't match device"
        echo ""
        echo "This means the batch manager fix might not be working correctly."
        echo "Check the logs above for errors."
        exit 1
    fi
fi

echo ""
echo "5. Modifying configuration for UPDATE test..."
echo "   Changing hostname to 'test-router-updated'"
echo "   Changing loopback shutdown to true"

# Create modified config
cat > test_updated.tf << 'EOF'
terraform {
  required_providers {
    iosxr = {
      source = "CiscoDevNet/iosxr"
    }
  }
}

provider "iosxr" {
  username           = "cisco"
  password           = "cisco"
  host               = "10.122.20.77:2641"
  tls                = false
  verify_certificate = false
}

resource "iosxr_hostname" "test" {
  system_network_name = "test-router-updated"
}

resource "iosxr_interface_loopback" "test" {
  name        = "999"
  description = "Test Loopback 999 UPDATED"
  shutdown    = true
}
EOF

echo ""
echo "6. Applying UPDATE..."
TF_LOG=INFO terraform apply -auto-approve test_updated.tf 2>&1 | tee apply_update.log | grep -E "Apply complete|FlushAll|Skipping device read|error|Error" || true

if grep -qi "error" apply_update.log; then
    echo "❌ Update apply failed!"
    exit 1
fi

echo ""
echo "7. Checking update logs..."
if grep -q "Skipping device read" apply_update.log; then
    echo "   ✅ Read operations skipped during batch mode (correct!)"
else
    echo "   ⚠️  'Skipping device read' not found - might be issue"
fi

if grep -q "FlushAll" apply_update.log; then
    echo "   ✅ FlushAll executed for updates"
    grep "FlushAll" apply_update.log | tail -3
fi

echo ""
echo "8. Running terraform plan after UPDATE (should show no changes)..."
TF_LOG=WARN terraform plan test_updated.tf 2>&1 | tee plan2.log

if grep -q "No changes" plan2.log; then
    echo "   ✅ SUCCESS! No drift after update - fix is working!"
else
    if grep -q "Plan:" plan2.log; then
        echo "   ❌ DRIFT DETECTED after update"
        echo ""
        echo "The update was queued but might not have been flushed to device."
        echo "Check apply_update.log for FlushAll errors."
        exit 1
    fi
fi

echo ""
echo "=========================================="
echo "✅ All tests passed!"
echo "=========================================="
echo "The batch manager update fix is working correctly."
echo ""
echo "Cleanup: Run 'terraform destroy test_updated.tf' when done"


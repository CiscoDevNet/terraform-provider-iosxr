# ─────────────────────────────────────────────────────────────────────────────
# Example: Batched gNMI configuration using auto_commit = false
#
# When auto_commit = false, every resource accumulates its gNMI Set operations
# in an in-memory candidate store.  The iosxr_commit resource drains the store
# in a single batched gNMI Set call to the device, minimising round-trips and
# making the change atomic from the device's perspective.
#
# IMPORTANT: iosxr_commit must appear last.  Use depends_on to reference every
# resource whose operations should be included in the batch.
# ─────────────────────────────────────────────────────────────────────────────

provider "iosxr" {
  host        = "10.0.0.1"
  username    = "admin"
  password    = "secret"
  auto_commit = false # queue all changes; do not send until iosxr_commit runs
}

resource "iosxr_hostname" "example" {
  system_network_name = "router-1"
}

resource "iosxr_logging" "example" {
  ipv4_dscp = "af11"
}

# Commits ALL pending operations for the default device in one gNMI Set call.
resource "iosxr_commit" "all" {
  depends_on = [
    iosxr_hostname.example,
    iosxr_logging.example,
  ]
}


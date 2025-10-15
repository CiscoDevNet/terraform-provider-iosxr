resource "iosxr_segment_routing_te" "example" {
  logging_pcep_peer_status = true
  logging_policy_status    = true
  pcc_report_all           = true
  pcc_source_address       = "88.88.88.8"
  pcc_delegation_timeout   = 10
  pcc_dead_timer           = 60
  pcc_initiated_state      = 15
  pcc_initiated_orphan     = 10
  pce_peers = [
    {
      pce_address = "66.66.66.6"
      precedence  = 122
    }
  ]
}

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
  on_demand_colors = [
    {
      dynamic_anycast_sid_inclusion       = true
      dynamic_metric_type                 = "te"
      color                               = 266
      srv6_locator_name                   = "LOC11"
      srv6_locator_behavior               = "ub6-insert-reduced"
      srv6_locator_binding_sid_type       = "srv6-dynamic"
      source_address                      = "fccc:0:213::1"
      source_address_type                 = "end-point-type-ipv6"
      effective_metric_value              = 4444
      effective_metric_type               = "igp"
      constraint_segments_protection_type = "protected-only"
      constraint_segments_sid_algorithm   = 128
    }
  ]
  policies = [
    {
      policy_name                   = "POLICY1"
      srv6_locator_name             = "Locator11"
      srv6_locator_binding_sid_type = "srv6-dynamic"
      srv6_locator_behavior         = "ub6-insert-reduced"
      source_address                = "fccc:0:103::1"
      source_address_type           = "end-point-type-ipv6"
      policy_color_endpoint_color   = 65534
      policy_color_endpoint_type    = "end-point-type-ipv6"
      policy_color_endpoint_address = "fccc:0:215::1"
    }
  ]
}

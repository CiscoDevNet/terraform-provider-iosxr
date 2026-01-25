resource "iosxr_interface_tunnel_te" "example" {
  name = "100"
  shutdown = false
  logging_events_link_status = true
  bandwidth = 1000000
  description = "My Interface Description"
  load_interval = 30
  ipv4_unnumbered = "Loopback0"
  mpls_mtu = 1400
  affinity_value = "11"
  affinity_mask = "ff"
  autoroute_announce = true
  autoroute_announce_metric_relative = 10
  autoroute_announce_include_ipv6 = true
  autoroute_destinations = [
    {
      address = "192.168.1.5"
    }
  ]
  backup_bw_class_type = "1"
  backup_bw_value = 10000
  signalled_bandwidth = 10000
  signalled_bandwidth_class_type = 1
  fast_reroute = true
  fast_reroute_protect_node = true
  fast_reroute_protect_bandwidth = true
  load_share = 1000
  logging_events_lsp_state = true
  logging_events_lsp_reoptimize = true
  logging_events_lsp_reoptimize_attempts = true
  logging_events_lsp_bw_change = true
  logging_events_lsp_reroute = true
  logging_events_lsp_record_route = true
  logging_events_lsp_switchover = true
  logging_events_lsp_insufficient_bw = true
  logging_events_pcalc_failure = true
  logging_events_bfd_status = true
  signalled_name = "Tunnel-TE-100"
  path_options = [
    {
      preference = 10
      dynamic = true
      isis_instance = "ISIS-1"
      isis_level = 2
      protected_by_index = 20
      protected_by_index_secondary = 30
      lockdown = true
      lockdown_sticky = true
    }
  ]
  priority_setup = 7
  priority_hold = 7
  record_route = true
  binding_sid_mpls_label = 4000
  policy_classes = ["1"]
  auto_bw_limit_min = 1000
  auto_bw_limit_max = 100000
  auto_bw_adjustment_threshold_percent = 80
  auto_bw_adjustment_threshold_min = 100
  auto_bw_overflow_threshold = 80
  auto_bw_overflow_min = 20
  auto_bw_overflow_limit = 3
  auto_bw_underflow_threshold = 10
  auto_bw_underflow_min = 10
  auto_bw_underflow_limit = 5
  auto_bw_resignal_last_bandwidth_timeout = 1000
  path_protection = true
  path_protection_srlg_diverse = true
  path_protection_non_revertive = true
  path_selection_metric_te = true
  path_selection_tiebreaker_min_fill = true
  path_selection_hop_limit = 10
  bidirectional_association_id = 10
  bidirectional_association_source_address = "192.168.1.1"
  bidirectional_association_global_id = 10
  bidirectional_association_corouted = true
  bfd_fast_detect = true
  bfd_minimum_interval = 100
  bfd_multiplier = 3
  bfd_bringup_timeout = 120
  bfd_lsp_ping_interval = 240
  bfd_dampening_initial_wait = 8000
  bfd_dampening_maximum_wait = 30000
  bfd_dampening_secondary_wait = 16000
  destination = "192.168.1.2"
}

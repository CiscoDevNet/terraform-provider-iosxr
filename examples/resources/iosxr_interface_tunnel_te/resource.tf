resource "iosxr_interface_tunnel_te" "example" {
  affinity_mask                           = "ff"
  affinity_value                          = "11"
  auto_bw_adjustment_threshold_min        = 100
  auto_bw_adjustment_threshold_percent    = 80
  auto_bw_limit_max                       = 100000
  auto_bw_limit_min                       = 1000
  auto_bw_overflow_limit                  = 3
  auto_bw_overflow_min                    = 20
  auto_bw_overflow_threshold              = 80
  auto_bw_resignal_last_bandwidth_timeout = 1000
  auto_bw_underflow_limit                 = 5
  auto_bw_underflow_min                   = 10
  auto_bw_underflow_threshold             = 10
  autoroute_announce                      = true
  autoroute_announce_include_ipv6         = true
  autoroute_announce_metric_relative      = 10
  autoroute_destinations = [
    {
      address = "192.168.1.5"
    }
  ]
  backup_bw_class_type                     = "1"
  backup_bw_value                          = 10000
  bandwidth                                = 1000000
  bfd_bringup_timeout                      = 120
  bfd_dampening_initial_wait               = 8000
  bfd_dampening_maximum_wait               = 30000
  bfd_dampening_secondary_wait             = 16000
  bfd_fast_detect                          = true
  bfd_lsp_ping_interval                    = 240
  bfd_minimum_interval                     = 100
  bfd_multiplier                           = 3
  bidirectional_association_corouted       = true
  bidirectional_association_global_id      = 10
  bidirectional_association_id             = 10
  bidirectional_association_source_address = "192.168.1.1"
  binding_sid_mpls_label                   = 4000
  description                              = "My Interface Description"
  destination                              = "192.168.1.2"
  fast_reroute                             = true
  fast_reroute_protect_bandwidth           = true
  fast_reroute_protect_node                = true
  ipv4_unnumbered                          = "Loopback0"
  load_interval                            = 30
  load_share                               = 1000
  logging_events_bfd_status                = true
  logging_events_link_status               = true
  logging_events_lsp_bw_change             = true
  logging_events_lsp_insufficient_bw       = true
  logging_events_lsp_record_route          = true
  logging_events_lsp_reoptimize            = true
  logging_events_lsp_reoptimize_attempts   = true
  logging_events_lsp_reroute               = true
  logging_events_lsp_state                 = true
  logging_events_lsp_switchover            = true
  logging_events_pcalc_failure             = true
  mpls_mtu                                 = 1400
  name                                     = "100"
  path_options = [
    {
      dynamic                      = true
      isis_instance                = "ISIS-1"
      isis_level                   = 2
      lockdown                     = true
      lockdown_sticky              = true
      preference                   = 10
      protected_by_index           = 20
      protected_by_index_secondary = 30
    }
  ]
  path_protection                    = true
  path_protection_non_revertive      = true
  path_protection_srlg_diverse       = true
  path_selection_hop_limit           = 10
  path_selection_metric_te           = true
  path_selection_tiebreaker_min_fill = true
  policy_classes                     = ["1"]
  priority_hold                      = 7
  priority_setup                     = 7
  record_route                       = true
  shutdown                           = false
  signalled_bandwidth                = 10000
  signalled_bandwidth_class_type     = 1
  signalled_name                     = "Tunnel-TE-100"
}

resource "iosxr_router_isis_address_family" "example" {
  process_id                             = "P1"
  af_name                                = "ipv4"
  saf_name                               = "unicast"
  metric_style_wide_transition           = true
  router_id_ip_address                   = "192.168.1.1"
  default_information_originate          = true
  fast_reroute_delay_interval            = 300
  fast_reroute_per_prefix_priority_limit = "critical"
  fast_reroute_per_prefix_priority_limit_levels = [
    {
      level_number   = 1
      priority_limit = "critical"
    }
  ]
  fast_reroute_per_prefix_use_candidate_only              = true
  fast_reroute_per_prefix_srlg_protection_weighted_global = true
  fast_reroute_per_prefix_srlg_protection_weighted_global_levels = [
    {
      level_number = 1
    }
  ]
  fast_reroute_per_prefix_load_sharing_disable = true
  fast_reroute_per_prefix_load_sharing_disable_levels = [
    {
      level_number = 1
    }
  ]
  fast_reroute_per_prefix_tiebreaker_downstream_index           = 10
  fast_reroute_per_prefix_tiebreaker_lc_disjoint_index          = 20
  fast_reroute_per_prefix_tiebreaker_lowest_backup_metric_index = 30
  fast_reroute_per_prefix_tiebreaker_node_protecting_index      = 40
  fast_reroute_per_prefix_tiebreaker_primary_path_index         = 50
  fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index        = 70
  fast_reroute_per_link_priority_limit_levels = [
    {
      level_number   = 1
      priority_limit = "critical"
    }
  ]
  fast_reroute_per_link_use_candidate_only = true
  microloop_avoidance_protected            = true
  microloop_avoidance_rib_update_delay     = 5000
  advertise_passive_only                   = true
  advertise_link_attributes                = true
  mpls_ldp_auto_config                     = false
  mpls_traffic_eng_router_id_ipv4_address  = "1.2.3.4"
  mpls_traffic_eng_level_1_2               = true
  spf_interval_maximum_wait                = 5000
  spf_interval_initial_wait                = 50
  spf_interval_secondary_wait              = 200
  spf_interval_levels = [
    {
      level_number   = 1
      maximum_wait   = 5000
      initial_wait   = 50
      secondary_wait = 200
    }
  ]
  spf_prefix_priority_critical_tag = 100
  spf_prefix_priority_high_tag     = 200
  spf_prefix_priority_medium_tag   = 300
  spf_prefix_priority_critical_levels = [
    {
      level_number = 1
      tag          = 100
    }
  ]
  spf_prefix_priority_high_levels = [
    {
      level_number = 1
      tag          = 200
    }
  ]
  spf_prefix_priority_medium_levels = [
    {
      level_number = 1
      tag          = 300
    }
  ]
  segment_routing_mpls_sr_prefer = true
  maximum_redistributed_prefixes = 100
  maximum_redistributed_prefixes_levels = [
    {
      level_number                   = 1
      maximum_redistributed_prefixes = 1000
    }
  ]
  redistribute_isis = [
    {
      instance_id     = "CORE"
      level           = "level-2"
      metric          = 10
      route_policy    = "ROUTE_POLICY_1"
      metric_type     = "internal"
      down_flag_clear = true
    }
  ]
}

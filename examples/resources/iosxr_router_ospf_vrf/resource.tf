resource "iosxr_router_ospf_vrf" "example" {
  process_name    = "OSPF1"
  vrf_name        = "VRF1"
  domain_id_type  = "0005"
  domain_id_value = "000000001111"
  domain_id_secondaries = [
    {
      type  = "0105"
      value = "001122334455"
    }
  ]
  domain_tag                                = 65001
  disable_dn_bit_check                      = true
  log_adjacency_changes_detail              = true
  router_id                                 = "10.11.12.13"
  redistribute_connected                    = true
  redistribute_connected_tag                = 1
  redistribute_connected_metric_type        = "1"
  redistribute_connected_route_policy       = "ROUTE_POLICY_1"
  redistribute_connected_metric             = 100
  redistribute_connected_lsa_type_summary   = true
  redistribute_connected_nssa_only          = true
  redistribute_static                       = true
  redistribute_static_tag                   = 2
  redistribute_static_metric_type           = "1"
  redistribute_static_route_policy          = "ROUTE_POLICY_1"
  redistribute_static_metric_use_rib_metric = true
  redistribute_static_lsa_type_summary      = true
  redistribute_static_nssa_only             = true
  redistribute_bgp = [
    {
      as_number        = "65001"
      tag              = 3
      metric_type      = "1"
      route_policy     = "ROUTE_POLICY_1"
      preserve_med     = true
      metric           = 100
      lsa_type_summary = true
      nssa_only        = true
    }
  ]
  redistribute_isis = [
    {
      instance_name    = "P1"
      level_1_2        = true
      tag              = 3
      metric_type      = "1"
      route_policy     = "ROUTE_POLICY_1"
      metric           = 100
      lsa_type_summary = true
      nssa_only        = true
    }
  ]
  redistribute_ospf = [
    {
      instance_name           = "OSPF2"
      tag                     = 4
      metric_type             = "1"
      route_policy            = "ROUTE_POLICY_1"
      match_internal          = true
      match_external_two      = true
      match_nssa_external_two = true
      metric                  = 100
      lsa_type_summary        = true
      nssa_only               = true
    }
  ]
  distribute_list_in_acl                     = "ACL_1"
  distribute_list_out_acl                    = "ACL_1"
  distribute_list_out_connected_acl          = "ACL_1"
  distribute_list_out_static_acl             = "ACL_1"
  distribute_list_out_bgp_as                 = "65001"
  distribute_list_out_bgp_acl                = "ACL_1"
  distribute_list_out_ospf_instance_name     = "OSPF2"
  distribute_list_out_ospf_acl               = "ACL_1"
  packet_size                                = 1400
  bfd_fast_detect                            = true
  bfd_fast_detect_strict_mode                = true
  bfd_minimum_interval                       = 300
  bfd_multiplier                             = 3
  security_ttl                               = true
  security_ttl_hops                          = 10
  prefix_suppression                         = true
  prefix_suppression_secondary_address       = true
  default_information_originate_always       = true
  default_information_originate_metric       = 100
  default_information_originate_metric_type  = 1
  default_information_originate_route_policy = "ROUTE_POLICY_1"
  default_metric                             = 1000
  distance_sources = [
    {
      address  = "192.168.1.0"
      wildcard = "0.0.0.255"
      distance = 100
      acl      = "ACL_1"
    }
  ]
  distance_ospf_intra_area                                  = 101
  distance_ospf_inter_area                                  = 102
  distance_ospf_external                                    = 103
  auto_cost_reference_bandwidth                             = 100000
  auto_cost_disable                                         = false
  ignore_lsa_mospf                                          = true
  capability_type7_prefer                                   = true
  max_metric_router_lsa                                     = true
  max_metric_router_lsa_include_stub                        = true
  max_metric_router_lsa_summary_lsa                         = true
  max_metric_router_lsa_summary_lsa_metric                  = 1000
  max_metric_router_lsa_external_lsa                        = true
  max_metric_router_lsa_external_lsa_metric                 = 2000
  max_metric_router_lsa_on_startup_time                     = 300
  max_metric_router_lsa_on_startup_include_stub             = true
  max_metric_router_lsa_on_startup_summary_lsa              = true
  max_metric_router_lsa_on_startup_summary_lsa_metric       = 1000
  max_metric_router_lsa_on_startup_external_lsa             = true
  max_metric_router_lsa_on_startup_external_lsa_metric      = 2000
  max_metric_router_lsa_on_switchover_time                  = 300
  max_metric_router_lsa_on_switchover_include_stub          = true
  max_metric_router_lsa_on_switchover_summary_lsa           = true
  max_metric_router_lsa_on_switchover_summary_lsa_metric    = 1000
  max_metric_router_lsa_on_switchover_external_lsa          = true
  max_metric_router_lsa_on_switchover_external_lsa_metric   = 2000
  max_metric_router_lsa_on_proc_restart_time                = 300
  max_metric_router_lsa_on_proc_restart_include_stub        = true
  max_metric_router_lsa_on_proc_restart_summary_lsa         = true
  max_metric_router_lsa_on_proc_restart_summary_lsa_metric  = 1000
  max_metric_router_lsa_on_proc_restart_external_lsa        = true
  max_metric_router_lsa_on_proc_restart_external_lsa_metric = 2000
  max_lsa                                                   = 20000
  max_lsa_threshold                                         = 75
  max_lsa_ignore_time                                       = 60
  max_lsa_ignore_count                                      = 10
  max_lsa_reset_time                                        = 120
  timers_throttle_spf_initial_delay                         = 500
  timers_throttle_spf_second_delay                          = 1000
  timers_throttle_spf_maximum_delay                         = 2000
  timers_throttle_lsa_all_initial_delay                     = 500
  timers_throttle_lsa_all_minimum_delay                     = 1000
  timers_throttle_lsa_all_maximum_delay                     = 2000
  timers_throttle_fast_reroute                              = 1000
  timers_lsa_group_pacing                                   = 60
  timers_lsa_min_arrival                                    = 60
  timers_lsa_refresh                                        = 2100
  timers_pacing_flood                                       = 10
  nsf_interval                                              = 120
  nsf_lifetime                                              = 300
  nsf_flush_delay_time                                      = 60
  nsf_ietf                                                  = true
  nsf_ietf_strict_lsa_checking                              = true
  nsf_ietf_helper_disable                                   = true
  address_family_ipv4_unicast                               = true
  maximum_interfaces                                        = 500
  maximum_paths                                             = 16
  maximum_redistributed_prefixes                            = 1000
  maximum_redistributed_prefixes_threshold                  = 75
  maximum_redistributed_prefixes_warning_only               = true
  queue_limit_high                                          = 1000
  queue_limit_medium                                        = 1000
  queue_limit_low                                           = 1000
  queue_dispatch_incoming                                   = 60
  queue_dispatch_rate_limited_lsa                           = 60
  queue_dispatch_flush_lsa                                  = 120
  queue_dispatch_spf_lsa_limit                              = 120
  summary_prefixes = [
    {
      address       = "192.168.1.0"
      mask          = "255.255.255.0"
      not_advertise = true
    }
  ]
  spf_prefix_priority_route_policy                              = "ROUTE_POLICY_1"
  fast_reroute_per_prefix                                       = true
  fast_reroute_per_prefix_priority_limit_medium                 = true
  fast_reroute_per_prefix_tiebreaker_downstream_index           = 10
  fast_reroute_per_prefix_tiebreaker_lc_disjoint_index          = 20
  fast_reroute_per_prefix_tiebreaker_lowest_backup_metric_index = 30
  fast_reroute_per_prefix_tiebreaker_node_protecting_index      = 40
  fast_reroute_per_prefix_tiebreaker_primary_path_index         = 50
  fast_reroute_per_prefix_tiebreaker_secondary_path_index       = 60
  fast_reroute_per_prefix_tiebreaker_interface_disjoint_index   = 70
  fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index        = 80
  fast_reroute_per_prefix_load_sharing_disable                  = true
  fast_reroute_per_prefix_exclude_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  fast_reroute_per_prefix_lfa_candidate_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/2"
    }
  ]
  fast_reroute_per_prefix_use_candidate_only_enable = true
  fast_reroute_per_link_priority_limit_medium       = true
  fast_reroute_per_link_exclude_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/3"
    }
  ]
  fast_reroute_per_link_lfa_candidate_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/4"
    }
  ]
  fast_reroute_per_link_use_candidate_only_enable = true
  loopback_stub_network_enable                    = true
  link_down_fast_detect                           = true
  weight                                          = 1000
  microloop_avoidance                             = true
  microloop_avoidance_segment_routing             = true
  microloop_avoidance_rib_update_delay            = 3000
  authentication_key_encrypted                    = "110A1016141D4B"
  message_digest_keys = [
    {
      key_id        = 1
      md5_encrypted = "01100F175804"
    }
  ]
  authentication                               = true
  authentication_message_digest                = true
  authentication_keychain_name                 = "KEY1"
  network_point_to_point                       = true
  mpls_ldp_sync                                = false
  cost                                         = 5000
  cost_fallback_anomaly_delay_igp_metric_value = 500
  cost_fallback_anomaly_delay_te_metric_value  = 600
  hello_interval                               = 10
  dead_interval                                = 40
  priority                                     = 10
  retransmit_interval                          = 1000
  transmit_delay                               = 100
  flood_reduction_enable                       = true
  demand_circuit_enable                        = true
  mtu_ignore_enable                            = true
  database_filter_all_out_enable               = true
  passive_disable                              = true
  external_out_enable                          = true
  summary_in_enable                            = true
  adjacency_stagger_initial_neighbors          = 10
  adjacency_stagger_simultaneous_neighbors     = 20
  snmp_context                                 = "CONTEXT1"
  snmp_trap                                    = true
  ucmp_prefix_list                             = "PREFIX_LIST_1"
  ucmp_exclude_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  ucmp_delay_interval                = 2000
  max_external_lsa                   = 1000
  max_external_lsa_threshold         = 75
  max_external_lsa_suppress_neighbor = true
  exchange_timer                     = 60
  exchange_timer_hold_time           = 120
  exchange_timer_recovery_count      = 10
}

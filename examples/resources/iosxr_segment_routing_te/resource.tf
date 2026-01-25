resource "iosxr_segment_routing_te" "example" {
  te_latency = true
  max_install_standby_paths = 2
  kshortest_paths = 120
  separate_next_hop = true
  maximum_sid_depth = 6
  effective_metric_admin_distance_metric_types = [
    {
      metric_type = "igp"
      admin_distance = 40
    }
  ]
  effective_metric_admin_distance_flex_algo_metric_types = [
    {
      metric_type = 0
      admin_distance = 40
    }
  ]
  steering_labeled_services_disable_on_demand_policies = true
  steering_labeled_services_disable_all_policies = true
  steering_labeled_services_disable_pcep_policies = true
  steering_labeled_services_disable_local_policies = true
  steering_labeled_services_disable_bgp_sr_te_policies = true
  segment_lists_srv6_topology_check = true
  segment_lists_srv6_sid_format = "micro-sid"
  segment_lists_srv6_explicit_segments = [
    {
      path_name = "SEG-1"
        srv6_segments = [
          {
            index = 1
            address = "fcbb:bb00:100::1"
            hop_type = "srv6sid"
          }
        ]
      srv6_topology_check = true
    }
  ]
  segment_lists_sr_mpls_explicit_segments = [
    {
      path_name = "SR-MPLS-1"
        sr_mpls_segments = [
          {
            index = 1
            type = "mpls-label"
            address = "192.168.1.2"
            mpls_label = 16200
            address_type = 2
          }
        ]
    }
  ]
  logging_pcep_peer_status = true
  logging_policy_status = true
  cspf_cache_size = 1000
  timers_candidate_path_cleanup_delay = 60
  timers_initial_verify_restart = 60
  timers_initial_verify_switchover = 120
  timers_initial_verify_startup = 60
  timers_delete_delay = 60
  timers_cleanup_delay = 60
  timers_install_delay = 60
  timers_periodic_reoptimization = 120
  resource_lists = [
    {
      path_name = "RESOURCE-1"
        resources = [
          {
            index = 1
            type = "ipv4-address"
            address = "192.168.1.2"
          }
        ]
    }
  ]
  distribute_link_state = true
  distribute_link_state_throttle = 15
  distribute_link_state_report_candidate_path_inactive = true
  bfd_timers_session_bringup = 120
  binding_sid_rules_explicit = "enforce-srlb"
  binding_sid_rules_dynamic = "disable"
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
        affinities = [
          {
            affinity_name = "AFFINITY-1"
          }
        ]
      metric = 1000
    }
  ]
  pce_peers_ipv4 = [
    {
      pce_address = "192.168.1.2"
      password_encrypted = "14141B180F0B6A"
      pce_group = "PCE-1"
      precedence = 110
    }
  ]
  pce_peers_ipv6 = [
    {
      pce_address = "2001:db8:1::1"
      pce_group = "PCE-1"
      precedence = 120
      tcp_ao_keychain = "KEY-2"
      tcp_ao_include_tcp_options = true
    }
  ]
  pcc_profiles = [
    {
      profile_id = 10
      steering_invalidation_drop = true
      auto_route_include_all_ipv4 = true
      auto_route_include_all_ipv6 = true
      auto_route_force_sr_include = true
      auto_route_forward_class = 1
      auto_route_metric_type = "relative"
      auto_route_metric_relative_value = -10
      auto_route_metric_constant_value = 100
    }
  ]
  pcc_redundancy_pcc_centric = true
  pcc_dead_timer = 60
  pcc_keepalive_timer = 60
  pcc_initiated_state = 120
  pcc_initiated_orphan = 120
  pcc_delegation_timeout = 120
  pcc_report_all = true
  pcc_source_address_ipv4 = "192.168.1.1"
  pcc_source_address_ipv6 = "2001:db8:1::1"
  traces = [
    {
      buffer_name = "bsid"
      trace_count = 1000
    }
  ]
  path_disable_algo_checks_strict_spf_areas = [
    {
      area_id = 1
      protocol = "ospf"
    }
  ]
  path_disable_algo_checks_strict_spf_all_areas = true
  affinity_maps = [
    {
      affinity_name = "AFFINITY-1"
      bit_position = 1
    }
  ]
  candidate_paths = [
    {
      path_type = "candidate-path-type-bgp-srte"
      source_address_selection = true
      source_address_type = "end-point-type-ipv4"
      source_address = "192.168.1.1"
    }
  ]
  srv6_locator = "LOC1"
  srv6_locator_binding_sid_type = "srv6-dynamic"
  srv6_locator_behavior = "ub6-insert-reduced"
  srv6_maximum_sid_depth = 6
}

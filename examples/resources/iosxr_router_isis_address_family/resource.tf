resource "iosxr_router_isis_address_family" "example" {
  process_id = "P1"
  af_name    = "ipv4"
  saf_name   = "unicast"
  distance   = 100
  distance_sources = [
    {
      address      = "192.168.1.1"
      prefix       = 32
      distance     = 101
      route_filter = "ROUTE_POLICY_1"
    }
  ]
  distribute_list_prefix_list_in      = "PREFIX_LIST_1"
  redistribute_connected              = true
  redistribute_connected_level        = "level-2"
  redistribute_connected_metric       = 100
  redistribute_connected_route_policy = "ROUTE_POLICY_1"
  redistribute_connected_metric_type  = "internal"
  redistribute_static                 = true
  redistribute_static_level           = "level-2"
  redistribute_static_metric          = 100
  redistribute_static_route_policy    = "ROUTE_POLICY_1"
  redistribute_static_metric_type     = "internal"
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
  redistribute_bgp = [
    {
      as_number    = "65001"
      level        = "level-2"
      metric       = 100
      route_policy = "ROUTE_POLICY_1"
      metric_type  = "internal"
    }
  ]
  redistribute_ospf = [
    {
      instance_id    = "OSPF1"
      match_external = true
      level          = "level-2"
      metric         = 100
      route_policy   = "ROUTE_POLICY_1"
      metric_type    = "internal"
    }
  ]
  maximum_paths                        = 10
  router_id_ip_address                 = "192.168.1.1"
  advertise_passive_only               = true
  advertise_link_attributes            = true
  microloop_avoidance                  = true
  microloop_avoidance_protected        = true
  microloop_avoidance_rib_update_delay = 5000
  summary_prefixes = [
    {
      address                          = "192.168.2.0"
      prefix                           = 24
      tag                              = 100
      level                            = 1
      algorithm                        = 128
      explicit                         = true
      adv_unreachable                  = true
      unreachable_tag                  = 200
      unreachable_tag_exclude_prefixes = true
      partition_repair                 = true
    }
  ]
  metric = 100
  metric_levels = [
    {
      level_number = 1
      metric       = 100
    }
  ]
  metric_style_wide_transition = true
  metric_style_levels = [
    {
      level_number    = 1
      wide_transition = true
    }
  ]
  spf_interval_maximum_wait   = 5000
  spf_interval_initial_wait   = 50
  spf_interval_secondary_wait = 200
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
  maximum_redistributed_prefixes = 100
  maximum_redistributed_prefixes_levels = [
    {
      level_number                   = 1
      maximum_redistributed_prefixes = 1000
    }
  ]
  propagate_levels = [
    {
      source_level      = 1
      destination_level = 2
      route_policy      = "ROUTE_POLICY_1"
    }
  ]
  adjacency_check_disable                = true
  route_source_first_hop                 = true
  attached_bit_receive_ignore            = true
  attached_bit_send                      = "always-set"
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
  fast_reroute_per_link_use_candidate_only                      = true
  fast_reroute_per_link_use_candidate_only_levels = [
    {
      level_number = 1
    }
  ]
  fast_reroute_per_link_priority_limit = "critical"
  fast_reroute_per_link_priority_limit_levels = [
    {
      level_number   = 1
      priority_limit = "critical"
    }
  ]
  default_information_originate                                   = true
  default_information_originate_route_policy                      = "ROUTE_POLICY_1"
  segment_routing_bundle_member_adj_sid                           = true
  segment_routing_labeled_only                                    = true
  segment_routing_protected_adjacency_sid_delay                   = 60
  segment_routing_mpls_sr_prefer                                  = true
  segment_routing_mpls_unlabeled_protection_route_policy          = "ROUTE_POLICY_1"
  segment_routing_mpls_prefix_sid_map_receive_disable             = true
  segment_routing_mpls_prefix_sid_map_advertise_local_domain_wide = true
  segment_routing_mpls_connected_prefix_sid_map                   = true
  segment_routing_mpls_connected_prefix_sid_map_addresses = [
    {
      ip_address          = "10.1.1.1"
      prefix              = 32
      index_id            = 400
      index_interface     = "GigabitEthernet0/0/0/3"
      index_php_disable   = true
      index_explicit_null = true
    }
  ]
  segment_routing_mpls_connected_prefix_sid_map_flex_algo_addresses = [
    {
      ip_address          = "10.1.1.2"
      prefix              = 32
      flex_algo           = 128
      index_id            = 500
      index_interface     = "GigabitEthernet0/0/0/3"
      index_php_disable   = true
      index_explicit_null = true
    }
  ]
  segment_routing_mpls_connected_prefix_sid_map_strict_spf_addresses = [
    {
      ip_address          = "10.1.1.3"
      prefix              = 32
      index_id            = 600
      index_interface     = "GigabitEthernet0/0/0/3"
      index_php_disable   = true
      index_explicit_null = true
    }
  ]
  partition_detect = true
  partition_detect_tracks = [
    {
      address = "192.168.3.1"
      ipv4    = true
    }
  ]
  partition_detect_external_address_tracks = [
    {
      address          = "192.168.3.2"
      external_address = "10.10.10.1"
    }
  ]
  mpls_ldp_auto_config                            = false
  mpls_traffic_eng_router_id_ipv4_address         = "1.2.3.4"
  mpls_traffic_eng_igp_intact                     = true
  mpls_traffic_eng_multicast_intact               = true
  mpls_traffic_eng_tunnel_restricted              = true
  mpls_traffic_eng_tunnel_preferred               = true
  mpls_traffic_eng_tunnel_metric                  = 100
  mpls_traffic_eng_tunnel_anycast_prefer_igp_cost = true
  mpls_traffic_eng_tunnel_metric_levels = [
    {
      level_number = 1
      metric       = 100
    }
  ]
  mpls_traffic_eng_level_1_2           = true
  prefix_unreachable                   = true
  prefix_unreachable_adv_maximum       = 60
  prefix_unreachable_adv_lifetime      = 120
  prefix_unreachable_adv_metric        = 4261412865
  prefix_unreachable_rx_process_enable = true
}

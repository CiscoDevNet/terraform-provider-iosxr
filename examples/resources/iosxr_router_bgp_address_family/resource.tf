resource "iosxr_router_bgp_address_family" "example" {
  as_number = "65001"
  af_name = "ipv4-unicast"
  distance_bgp_external_route = 200
  distance_bgp_internal_route = 195
  distance_bgp_local_route = 190
  maximum_paths_ebgp_multipath = 10
  maximum_paths_ebgp_selective = true
  maximum_paths_ebgp_route_policy = "ROUTE_POLICY_1"
  maximum_paths_ibgp_multipath = 10
  maximum_paths_ibgp_unequal_cost_deterministic = true
  maximum_paths_ibgp_selective = true
  maximum_paths_ibgp_route_policy = "ROUTE_POLICY_1"
  maximum_paths_unique_nexthop_check_disable = true
  import_from_bridge_domain = true
  additional_paths_send = true
  additional_paths_receive = true
  additional_paths_advertise_limit = 40
  additional_paths_selection_route_policy = "ROUTE_POLICY_1"
  permanent_network_route_policy = "ROUTE_POLICY_1"
  advertise_best_external_labeled_unicast = true
  advertise_local_labeled_route_safi_unicast = "disable"
  advertise_epe_bgp_labeled_unicast = true
  networks = [
    {
      address = "10.1.0.0"
      prefix = 16
      route_policy = "ROUTE_POLICY_1"
    }
  ]
  aggregate_addresses = [
    {
      address = "10.0.0.0"
      prefix = 8
      as_set = false
      as_confed_set = false
      summary_only = true
      route_policy = "ROUTE_POLICY_1"
      description = "Aggregate route description"
      set_tag = 100
    }
  ]
  redistribute_ospf = [
    {
      router_tag = "OSPF1"
      metric = 100
      multipath = true
      route_policy = "ROUTE_POLICY_1"
    }
  ]
  redistribute_eigrp = [
    {
      instance_name = "EIGRP1"
      match_internal_external = true
      metric = 100
      multipath = true
      route_policy = "ROUTE_POLICY_1"
    }
  ]
  redistribute_isis = [
    {
      instance_name = "ISIS1"
      level_1_level_2_level_1_inter_area = true
      metric = 100
      multipath = true
      route_policy = "ROUTE_POLICY_1"
    }
  ]
  redistribute_connected = true
  redistribute_connected_metric = 100
  redistribute_connected_multipath = true
  redistribute_connected_route_policy = "ROUTE_POLICY_1"
  redistribute_static = true
  redistribute_static_metric = 100
  redistribute_static_multipath = true
  redistribute_static_route_policy = "ROUTE_POLICY_1"
  redistribute_rip = true
  redistribute_rip_metric = 100
  redistribute_rip_multipath = true
  redistribute_rip_route_policy = "ROUTE_POLICY_1"
  table_policy = "ROUTE_POLICY_1"
  retain_local_label = 30
  allocate_label_all_unlabeled_path = true
  rnh_install_extcomm_only = true
  prefix_ecmp_delay = 1000
  prefix_ecmp_delay_oor_threshold = 90
  bgp_origin_as_validation_enable = true
  bgp_origin_as_validation_signal_ibgp = true
  bgp_bestpath_origin_as_use_validity = true
  bgp_bestpath_origin_as_allow_invalid = true
  bgp_scan_time = 60
  bgp_attribute_download = true
  bgp_label_delay_seconds = 5
  bgp_label_delay_milliseconds = 500
  bgp_client_to_client_reflection_disable = true
  bgp_client_to_client_reflection_cluster_ids_32bit_format = [
    {
      cluster_as = 65001
      disable = true
    }
  ]
  bgp_client_to_client_reflection_cluster_ids_ip_format = [
    {
      cluster_ip = "192.168.1.1"
      disable = true
    }
  ]
  bgp_dampening_decay_half_life = 30
  bgp_dampening_reuse_threshold = 40
  bgp_dampening_suppress_threshold = 50
  bgp_dampening_max_suppress_time = 30
  event_prefix_route_policy = "ROUTE_POLICY_1"
  dynamic_med_interval = 5
  weight_reset_on_import = true
  nexthop_trigger_delay_critical = 1000
  nexthop_trigger_delay_non_critical = 2000
  nexthop_route_policy = "ROUTE_POLICY_1"
  nexthop_resolution_prefix_length_minimum_ipv4 = 32
  nexthop_resolution_prefix_length_minimum_ipv6 = 128
  update_limit_sub_group_ebgp = 10
  update_limit_sub_group_ibgp = 10
  update_limit_address_family = 10
  update_wait_install = true
  update_wait_install_delay_startup = 300
  as_path_loopcheck_out_disable = true
  epe_backup_enable = true
  default_martian_check_disable = true
  export_to_vrf_allow_backup = true
  export_to_vrf_allow_best_external = true
  segment_routing_prefix_sid_map = true
  segment_routing_srv6_locator = "locator100"
  segment_routing_srv6_usid_allocation_wide_local_id_block = true
  segment_routing_srv6_alloc_mode_per_vrf = true
  peer_set_ids = [
    {
      peer_id = 1
      peer_sid_index = 101
    }
  ]
}

resource "iosxr_router_bgp_vrf_address_family" "example" {
  as_number                                                = "65001"
  vrf_name                                                 = "VRF2"
  af_name                                                  = "ipv4-unicast"
  segment_routing_srv6_locator                             = "locator101"
  segment_routing_srv6_usid_allocation_wide_local_id_block = true
  segment_routing_srv6_alloc_mode_per_vrf                  = true
  distance_bgp_external_route                              = 200
  distance_bgp_internal_route                              = 195
  distance_bgp_local_route                                 = 190
  bgp_attribute_download                                   = true
  allow_vpn_default_originate                              = true
  maximum_paths_ebgp_multipath                             = 10
  maximum_paths_ebgp_selective                             = true
  maximum_paths_ebgp_route_policy                          = "ROUTE_POLICY_1"
  maximum_paths_ibgp_multipath                             = 10
  maximum_paths_ibgp_unequal_cost                          = true
  maximum_paths_ibgp_unequal_cost_deterministic            = true
  maximum_paths_ibgp_selective                             = true
  maximum_paths_ibgp_route_policy                          = "ROUTE_POLICY_1"
  additional_paths_send                                    = true
  additional_paths_receive                                 = true
  additional_paths_advertise_limit                         = 40
  additional_paths_selection_route_policy                  = "ROUTE_POLICY_1"
  advertise_best_external                                  = true
  advertise_local_labeled_route_safi_unicast               = "disable"
  networks = [
    {
      address      = "10.1.0.0"
      prefix       = 16
      route_policy = "ROUTE_POLICY_1"
    }
  ]
  aggregate_addresses = [
    {
      address       = "10.0.0.0"
      prefix        = 8
      as_set        = true
      as_confed_set = false
      summary_only  = true
      route_policy  = "ROUTE_POLICY_1"
      description   = "Aggregate route description"
      set_tag       = 100
    }
  ]
  redistribute_ospf = [
    {
      router_tag                              = "OSPF1"
      match_internal_external_nssa_external_2 = true
      metric                                  = 100
      multipath                               = true
      route_policy                            = "ROUTE_POLICY_1"
    }
  ]
  redistribute_eigrp = [
    {
      instance_name           = "EIGRP1"
      match_internal_external = true
      metric                  = 100
      multipath               = true
      route_policy            = "ROUTE_POLICY_1"
    }
  ]
  redistribute_isis = [
    {
      instance_name                      = "ISIS1"
      level_1_level_2_level_1_inter_area = true
      metric                             = 100
      multipath                          = true
      route_policy                       = "ROUTE_POLICY_1"
    }
  ]
  redistribute_connected               = true
  redistribute_connected_metric        = 100
  redistribute_connected_multipath     = true
  redistribute_connected_route_policy  = "ROUTE_POLICY_1"
  redistribute_static                  = true
  redistribute_static_metric           = 100
  redistribute_static_multipath        = true
  redistribute_static_route_policy     = "ROUTE_POLICY_1"
  redistribute_rip                     = true
  redistribute_rip_metric              = 100
  redistribute_rip_multipath           = true
  redistribute_rip_route_policy        = "ROUTE_POLICY_1"
  table_policy                         = "ROUTE_POLICY_1"
  label_mode_per_prefix                = true
  bgp_origin_as_validation_enable      = true
  bgp_origin_as_validation_signal_ibgp = true
  bgp_bestpath_origin_as_use_validity  = true
  bgp_bestpath_origin_as_allow_invalid = true
  bgp_dampening_decay_half_life        = 30
  bgp_dampening_reuse_threshold        = 40
  bgp_dampening_suppress_threshold     = 50
  bgp_dampening_max_suppress_time      = 30
  dynamic_med_interval                 = 5
  weight_reset_on_import               = true
  nexthop_route_policy                 = "ROUTE_POLICY_1"
  as_path_loopcheck_out_disable        = true
  mvpn_single_forwarder_selection      = "highest-ip-address"
}

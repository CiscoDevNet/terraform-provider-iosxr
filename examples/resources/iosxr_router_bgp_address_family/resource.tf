resource "iosxr_router_bgp_address_family" "example" {
  as_number                                     = "65001"
  af_name                                       = "ipv4-unicast"
  additional_paths_send                         = true
  additional_paths_receive                      = true
  additional_paths_selection_route_policy       = "ADDITIONAL_PATHS_POLICY"
  allocate_label_all                            = true
  allocate_label_all_unlabeled_path             = true
  advertise_best_external                       = true
  maximum_paths_ebgp_multipath                  = 10
  maximum_paths_ebgp_selective                  = true
  maximum_paths_ibgp_multipath                  = 10
  maximum_paths_ibgp_unequal_cost_deterministic = true
  maximum_paths_ibgp_selective                  = true
  maximum_paths_unique_nexthop_check_disable    = true
  nexthop_trigger_delay_critical                = 10
  nexthop_trigger_delay_non_critical            = 20
  aggregate_addresses = [
    {
      address       = "10.0.0.0"
      prefix        = 8
      as_set        = false
      as_confed_set = false
      summary_only  = true
      route_policy  = "ROUTE_POLICY_1"
      description   = "Aggregate route description"
      set_tag       = 100
    }
  ]
  networks = [
    {
      address      = "10.1.0.0"
      prefix       = 16
      route_policy = "ROUTE_POLICY_1"
    }
  ]
  redistribute_ospf = [
    {
      router_tag   = "OSPF1"
      metric       = 100
      multipath    = true
      route_policy = "REDISTRIBUTE_POLICY"
    }
  ]
  redistribute_eigrp = [
    {
      instance_name           = "EIGRP1"
      match_internal          = true
      match_internal_external = true
      metric                  = 100
      multipath               = true
      route_policy            = "REDISTRIBUTE_POLICY"
    }
  ]
  redistribute_isis = [
    {
      instance_name                      = "ISIS1"
      level_1_level_2_level_1_inter_area = true
      metric                             = 100
      multipath                          = true
      route_policy                       = "REDISTRIBUTE_POLICY"
    }
  ]
  redistribute_connected              = true
  redistribute_connected_metric       = 100
  redistribute_connected_multipath    = true
  redistribute_connected_route_policy = "REDISTRIBUTE_POLICY"
  redistribute_static                 = true
  redistribute_static_metric          = 100
  redistribute_static_multipath       = true
  redistribute_static_route_policy    = "REDISTRIBUTE_POLICY"
  redistribute_rip                    = true
  redistribute_rip_metric             = 100
  redistribute_rip_multipath          = true
  redistribute_rip_route_policy       = "REDISTRIBUTE_POLICY"
}

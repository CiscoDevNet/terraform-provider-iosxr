resource "iosxr_router_bgp_vrf_address_family" "example" {
  as_number                               = "65001"
  vrf_name                                = "VRF1"
  af_name                                 = "ipv4-unicast"
  additional_paths_send                   = true
  additional_paths_send_disable           = true
  additional_paths_receive                = true
  additional_paths_receive_disable        = true
  additional_paths_selection_route_policy = "ADDITIONAL_PATHS_SELECTION_POLICY"
  allocate_label_route_policy_name        = "ALLOCATE_LABEL_POLICY"
  advertise_best_external                 = true
  label_mode_per_vrf                      = true
  segment_routing_srv6_locator            = "LocAlgo11"
  aggregate_addresses = [
    {
      address        = "10.0.0.0"
      address_prefix = 8
      as_set         = true
      as_confed_set  = false
      summary_only   = true
      description    = "Aggregate route description"
      set_tag        = 100
    }
  ]
  networks = [
    {
      address        = "10.1.0.0"
      address_prefix = 16
      route_policy   = "ROUTE_POLICY_1"
    }
  ]
  redistribute_ospf = [
    {
      ospf_router_tag                                                       = "1"
      redistribute_ospf_match_internal_external_type_1_nssa_external_type_2 = true
      metric                                                                = 100
      multipath                                                             = true
      route_policy                                                          = "REDISTRIBUTE_POLICY"
    }
  ]
  redistribute_eigrp = [
    {
      eigrp_name                           = "EIGRP1"
      redistribute_eigrp_internal          = true
      redistribute_eigrp_internal_external = true
      metric                               = 100
      multipath                            = true
      route_policy                         = "REDISTRIBUTE_POLICY"
    }
  ]
  redistribute_isis = [
    {
      isis_name                                            = "ISIS1"
      redistribute_isis_level_1_level_2_level_1_inter_area = true
      metric                                               = 100
      multipath                                            = true
      route_policy                                         = "REDISTRIBUTE_POLICY"
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

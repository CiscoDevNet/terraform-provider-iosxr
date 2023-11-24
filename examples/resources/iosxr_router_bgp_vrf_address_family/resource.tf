resource "iosxr_router_bgp_vrf_address_family" "example" {
  as_number                               = "65001"
  vrf_name                                = "VRF1"
  af_name                                 = "ipv4-unicast"
  additional_paths_send                   = true
  additional_paths_receive                = true
  additional_paths_selection_route_policy = "ROUTE_POLICY_1"
  advertise_best_external                 = true
  allocate_label_all                      = true
  maximum_paths_ebgp_multipath            = 10
  maximum_paths_ibgp_multipath            = 10
  label_mode_per_ce                       = false
  label_mode_per_vrf                      = false
  redistribute_connected                  = true
  redistribute_connected_metric           = 10
  redistribute_connected_route_policy     = "ROUTE_POLICY_1"
  redistribute_static                     = true
  redistribute_static_metric              = 10
  redistribute_static_route_policy        = "ROUTE_POLICY_1"
  segment_routing_srv6_locator            = "LocAlgo11"
  segment_routing_srv6_alloc_mode_per_vrf = true
  aggregate_addresses = [
    {
      address       = "10.0.0.0"
      masklength    = 8
      as_set        = false
      as_confed_set = false
      summary_only  = false
    }
  ]
  networks = [
    {
      address      = "10.1.0.0"
      masklength   = 16
      route_policy = "ROUTE_POLICY_1"
    }
  ]
  redistribute_ospf = [
    {
      router_tag                   = "OSPF1"
      match_internal               = true
      match_internal_external      = true
      match_internal_nssa_external = false
      match_external               = false
      match_external_nssa_external = false
      match_nssa_external          = false
      metric                       = 100
      route_policy                 = "ROUTE_POLICY_1"
    }
  ]
}

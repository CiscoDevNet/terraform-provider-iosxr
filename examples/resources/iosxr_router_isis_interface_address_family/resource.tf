resource "iosxr_router_isis_interface_address_family" "example" {
  process_id     = "P1"
  interface_name = "GigabitEthernet0/0/0/1"
  af_name        = "ipv4"
  saf_name       = "unicast"
  metric_default = 500
  metric_levels = [
    {
      level_number   = 1
      metric_default = 600
    }
  ]
  te_metric_flex_algo = 128
  te_metric_flex_algo_levels = [
    {
      level_number = 1
      flex_algo    = 128
    }
  ]
  bandwidth_metric_flex_algo = 129
  bandwidth_metric_flex_algo_levels = [
    {
      level_number = 1
      flex_algo    = 129
    }
  ]
  generic_metric_flex_algos = [
    {
      type   = 130
      metric = 5000
    }
  ]
  generic_metric_flex_algo_levels = [
    {
      level_number = 1
      flex_algos_types = [
        {
          type   = 130
          metric = 5000
        }
      ]
    }
  ]
  mpls_ldp_sync       = true
  mpls_ldp_sync_level = 1
  tag                 = 100
  tag_levels = [
    {
      level_number = 1
      tag          = 100
    }
  ]
  fast_reroute_per_prefix = true
  fast_reroute_levels = [
    {
      level_number = 1
      per_prefix   = true
    }
  ]
  fast_reroute_per_prefix_tiebreaker_node_protecting_index = 10
  fast_reroute_per_prefix_tiebreaker_node_protecting_levels = [
    {
      level_number = 1
      index        = 10
    }
  ]
  fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index = 20
  fast_reroute_per_prefix_tiebreaker_srlg_disjoint_levels = [
    {
      level_number = 1
      index        = 20
    }
  ]
  fast_reroute_per_prefix_tiebreaker_lc_disjoint_index = 30
  fast_reroute_per_prefix_tiebreaker_lc_disjoint_levels = [
    {
      level_number = 1
      index        = 30
    }
  ]
  fast_reroute_per_prefix_remote_lfa_maximum_metric = 100
  fast_reroute_per_prefix_remote_lfa_maximum_metric_levels = [
    {
      level_number   = 1
      maximum_metric = 100
    }
  ]
  fast_reroute_per_prefix_remote_lfa_tunnel_mpls_ldp = true
  fast_reroute_per_prefix_remote_lfa_tunnel_mpls_ldp_levels = [
    {
      level_number = 1
    }
  ]
  fast_reroute_per_prefix_ti_lfa = true
  fast_reroute_per_prefix_ti_lfa_levels = [
    {
      level_number = 1
    }
  ]
  fast_reroute_per_prefix_exclude_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/2"
      level          = 1
    }
  ]
  fast_reroute_per_prefix_lfa_candidate_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/3"
      level          = 1
    }
  ]
  fast_reroute_per_link_exclude_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/2"
      level          = 1
    }
  ]
  fast_reroute_per_link_lfa_candidate_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/3"
      level          = 1
    }
  ]
  link_group_name  = "LINK_GROUP_1"
  link_group_level = 1
  weight           = 500
  weight_levels = [
    {
      level_number = 1
      weight       = 500
    }
  ]
  auto_metric_proactive_protect_metric = 500
  auto_metric_proactive_protect_metric_levels = [
    {
      level_number      = 1
      proactive_protect = 500
    }
  ]
  advertise_prefix_route_policy = "ROUTE_POLICY_1"
  advertise_prefix_route_policy_levels = [
    {
      level_number = 1
      route_policy = "ROUTE_POLICY_2"
    }
  ]
}

resource "iosxr_router_bgp_vrf" "example" {
  as_number                             = "65001"
  bfd_minimum_interval                  = 10
  bfd_multiplier                        = 4
  bgp_auto_policy_soft_reset_disable    = true
  bgp_bestpath_aigp_ignore              = true
  bgp_bestpath_as_path_ignore           = true
  bgp_bestpath_as_path_multipath_relax  = true
  bgp_bestpath_compare_routerid         = true
  bgp_bestpath_cost_community_ignore    = true
  bgp_bestpath_igp_metric_ignore        = true
  bgp_bestpath_med_missing_as_worst     = true
  bgp_bestpath_origin_as_allow_invalid  = true
  bgp_bestpath_origin_as_use_validity   = true
  bgp_bestpath_sr_policy_prefer         = true
  bgp_default_local_preference          = 200
  bgp_enforce_first_as_disable          = true
  bgp_fast_external_fallover_disable    = true
  bgp_log_message_disable               = true
  bgp_log_neighbor_changes_disable      = true
  bgp_multipath_use_cluster_list_length = true
  bgp_origin_as_validation_signal_ibgp  = true
  bgp_redistribute_internal             = true
  bgp_router_id                         = "22.22.22.22"
  bgp_unsafe_ebgp_policy                = true
  default_information_originate         = true
  default_metric                        = 125
  distance_bgp_external                 = 100
  distance_bgp_internal                 = 150
  distance_bgp_local                    = 200
  mpls_activate_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  nexthop_mpls_forwarding_ibgp                     = true
  nexthop_resolution_allow_default                 = true
  rd_auto                                          = true
  socket_receive_buffer_size                       = 1024
  socket_receive_buffer_size_read                  = 1024
  socket_send_buffer_size                          = 4096
  socket_send_buffer_size_write                    = 4096
  timers_bgp_holddown_zero                         = true
  timers_bgp_holddown_zero_minimum_acceptable_zero = true
  timers_bgp_keepalive_interval                    = 0
  vrf_name                                         = "VRF1"
}

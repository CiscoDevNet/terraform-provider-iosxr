resource "iosxr_evpn_interface" "example" {
  interface_name = "Bundle-Ether12"
  core_isolation_group = 11
  timers_peering = 60
  timers_recovery = 120
  timers_carving = 5
  timers_ac_debounce = 2000
  ethernet_segment_esi_zero = "01.01.01.01.01.01.01.01.04"
  ethernet_segment_load_balancing_mode_port_active = true
  ethernet_segment_force_single_homed = true
  ethernet_segment_service_carving_hrw = true
  ethernet_segment_service_carving_multicast_hrw_s_g = true
  ethernet_segment_service_carving_preference_based_weight = 100
  ethernet_segment_service_carving_preference_based_access_driven = true
  ethernet_segment_bgp_rt = "01:01:01:01:01:04"
  ethernet_segment_convergence_reroute = true
  ethernet_segment_convergence_mac_mobility = true
  ethernet_segment_convergence_nexthop_tracking = true
  access_signal_bundle_down = true
}

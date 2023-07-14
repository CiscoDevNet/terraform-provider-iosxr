resource "iosxr_evpn_interface" "example" {
  interface_name                                          = "Bundle-Ether12"
  core_isolation_group                                    = 11
  ethernet_segment_load_balancing_mode_all_active         = false
  ethernet_segment_load_balancing_mode_port_active        = false
  ethernet_segment_load_balancing_mode_single_active      = true
  ethernet_segment_load_balancing_mode_single_flow_active = false
}

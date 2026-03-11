resource "iosxr_evpn_evi" "example" {
  vpn_id = 101
  description = "My Description"
  bgp_rd_four_byte_as_number = 65536
  bgp_rd_four_byte_as_index = 101
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 101
    }
  ]
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 101
    }
  ]
  bgp_implicit_import_disable = true
  bgp_route_policy_import = "EVI_POLICY_1"
  bgp_route_policy_export = "EVI_POLICY_1"
  load_balancing = true
  load_balancing_flow_label_static = true
  preferred_nexthop_modulo = true
  advertise_mac_bvi_mac = true
  unknown_unicast_suppression = true
  control_word_disable = true
  ignore_mtu_mismatch = true
  ignore_mtu_mismatch_disable_deprecated = true
  transmit_mtu_zero = true
  transmit_mtu_zero_disable_deprecated = true
  re_origination_disable = true
  multicast_source_connected = true
  proxy_igmp_snooping = true
  etree = true
  etree_leaf = false
  etree_rt_leaf = true
  vpws_single_active_backup_suppression = true
  bvi_coupled_mode = true
}

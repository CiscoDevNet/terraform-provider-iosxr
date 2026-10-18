resource "iosxr_evpn_evi" "example" {
  advertise_mac_bvi_mac       = true
  bgp_implicit_import_disable = true
  bgp_rd_four_byte_as_index   = 101
  bgp_rd_four_byte_as_number  = 65536
  bgp_route_policy_export     = "EVI_POLICY_1"
  bgp_route_policy_import     = "EVI_POLICY_1"
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 101
    }
  ]
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 101
    }
  ]
  bvi_coupled_mode                       = true
  control_word_disable                   = true
  description                            = "My Description"
  etree                                  = true
  etree_leaf                             = false
  etree_rt_leaf                          = true
  ignore_mtu_mismatch                    = true
  ignore_mtu_mismatch_disable_deprecated = true
  load_balancing                         = true
  load_balancing_flow_label_static       = true
  multicast_source_connected             = true
  preferred_nexthop_modulo               = true
  proxy_igmp_snooping                    = true
  re_origination_disable                 = true
  transmit_mtu_zero                      = true
  transmit_mtu_zero_disable_deprecated   = true
  unknown_unicast_suppression            = true
  vpn_id                                 = 101
  vpws_single_active_backup_suppression  = true
}

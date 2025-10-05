resource "iosxr_evpn_evi" "example" {
  vpn_id                             = 1234
  description                        = "My Description"
  load_balancing                     = true
  load_balancing_flow_label_static   = true
  bgp_rd_two_byte_as_number          = 1
  bgp_rd_two_byte_as_assigned_number = 1
  bgp_route_target_import_two_byte_as_format = [
    {
      as_number       = 1
      assigned_number = 1
    }
  ]
  bgp_route_target_export_ipv4_address_format = [
    {
      ipv4_address    = "1.1.1.1"
      assigned_number = 1
    }
  ]
  bgp_route_policy_import     = "EVI_POLICY_1"
  bgp_route_policy_export     = "EVI_POLICY_1"
  advertise_mac               = true
  unknown_unicast_suppression = true
  control_word_disable        = true
  etree                       = true
  etree_leaf                  = false
  etree_rt_leaf               = true
}

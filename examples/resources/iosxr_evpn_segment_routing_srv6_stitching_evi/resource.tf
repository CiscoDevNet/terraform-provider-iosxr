resource "iosxr_evpn_segment_routing_srv6_stitching_evi" "example" {
  vpn_id = 104
  description = "My Description"
  bgp_rd_four_byte_as_number = 65536
  bgp_rd_four_byte_as_index = 104
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 104
    }
  ]
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 104
    }
  ]
  bgp_route_policy_import = "EVI_POLICY_1"
  bgp_route_policy_export = "EVI_POLICY_1"
  preferred_nexthop_modulo = true
  unknown_unicast_suppression = true
  ignore_mtu_mismatch = true
  ignore_mtu_mismatch_disable = true
  transmit_mtu_zero = true
  transmit_mtu_zero_disable = true
  re_origination_disable = true
}

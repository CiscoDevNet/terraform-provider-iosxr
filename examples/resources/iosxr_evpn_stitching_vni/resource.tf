resource "iosxr_evpn_stitching_vni" "example" {
  bgp_rd_four_byte_as_index  = 106
  bgp_rd_four_byte_as_number = 65536
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 106
    }
  ]
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 106
    }
  ]
  description                 = "My Description"
  preferred_nexthop_lowest_ip = true
  re_origination_disable      = true
  unknown_unicast_suppression = true
  vni_id                      = 106
}

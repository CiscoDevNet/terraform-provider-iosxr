resource "iosxr_evpn_stitching_vni" "example" {
  vni_id = 106
  description = "My Description"
  bgp_rd_four_byte_as_number = 65536
  bgp_rd_four_byte_as_index = 106
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 106
    }
  ]
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 106
    }
  ]
  preferred_nexthop_lowest_ip = true
  unknown_unicast_suppression = true
  re_origination_disable = true
}

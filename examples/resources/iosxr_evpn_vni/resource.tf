resource "iosxr_evpn_vni" "example" {
  bgp_rd_four_byte_as_index  = 105
  bgp_rd_four_byte_as_number = 65536
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 105
    }
  ]
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 105
    }
  ]
  description                 = "My Description"
  preferred_nexthop_lowest_ip = true
  re_origination_disable      = true
  unknown_unicast_suppression = true
  vni_id                      = 105
}

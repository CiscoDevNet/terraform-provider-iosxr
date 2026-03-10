resource "iosxr_evpn_vni" "example" {
  vni_id                     = 105
  description                = "My Description"
  bgp_rd_four_byte_as_number = 65536
  bgp_rd_four_byte_as_index  = 105
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 105
    }
  ]
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number       = 65536
      assigned_number = 105
    }
  ]
  preferred_nexthop_lowest_ip = true
  unknown_unicast_suppression = true
  re_origination_disable      = true
}

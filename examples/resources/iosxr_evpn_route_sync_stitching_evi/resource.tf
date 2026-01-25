resource "iosxr_evpn_route_sync_stitching_evi" "example" {
  vpn_id = 108
  description = "My Description"
  bgp_rd_four_byte_as_number = 65536
  bgp_rd_four_byte_as_index = 108
  bgp_route_target_import_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 108
    }
  ]
  bgp_route_target_export_four_byte_as_format = [
    {
      as_number = 65536
      assigned_number = 108
    }
  ]
}

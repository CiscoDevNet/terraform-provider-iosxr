resource "iosxr_evpn_segment_routing_srv6_evi" "example" {
  vpn_id      = 1235
  description = "My Description"
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
  advertise_mac = true
  locator       = "LOC12"
}

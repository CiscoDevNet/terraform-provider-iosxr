resource "iosxr_evpn_segment_routing_services" "example" {
  vpn_id      = 1234
  description = "My Description"
  evpn_bgp_route_target_import_two_byte_as_format = [
    {
      as_number       = 1
      assigned_number = 1
    }
  ]
  evpn_bgp_route_target_export_ipv4_address_format = [
    {
      ipv4_address    = "1.1.1.1"
      assigned_number = 1
    }
  ]
  advertise_mac_bvi_mac = true
  locator               = "LOC12"
}

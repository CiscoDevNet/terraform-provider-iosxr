resource "iosxr_vrf" "example" {
  vrf_name                                        = "VRF3"
  description                                     = "My VRF Description"
  vpn_id                                          = "1000:1000"
  address_family_ipv4_unicast                     = true
  address_family_ipv4_unicast_import_route_policy = "ROUTE_POLICY_1"
  address_family_ipv4_unicast_export_route_policy = "ROUTE_POLICY_1"
  address_family_ipv4_multicast                   = true
  address_family_ipv4_flowspec                    = true
  address_family_ipv6_unicast                     = true
  address_family_ipv6_unicast_import_route_policy = "ROUTE_POLICY_1"
  address_family_ipv6_unicast_export_route_policy = "ROUTE_POLICY_1"
  address_family_ipv6_multicast                   = true
  address_family_ipv6_flowspec                    = true
  rd_two_byte_as_as_number                        = "123"
  rd_two_byte_as_index                            = 123
  address_family_ipv4_unicast_import_route_target_two_byte_as_format = [
    {
      as_number = 1
      index     = 1
      stitching = true
    }
  ]
  address_family_ipv4_unicast_export_route_target_two_byte_as_format = [
    {
      as_number = 1
      index     = 1
      stitching = true
    }
  ]
}

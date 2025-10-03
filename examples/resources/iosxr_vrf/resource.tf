resource "iosxr_vrf" "example" {
  vrf_name                         = "VRF3"
  description                      = "My VRF Description"
  vpn_id                           = "1000:1000"
  ipv4_unicast                     = true
  ipv4_unicast_import_route_policy = "ROUTE_POLICY_1"
  ipv4_unicast_export_route_policy = "ROUTE_POLICY_1"
  ipv6_unicast                     = true
  ipv6_unicast_import_route_policy = "ROUTE_POLICY_1"
  ipv6_unicast_export_route_policy = "ROUTE_POLICY_1"
  rd_two_byte_as_number            = "65001"
  rd_two_byte_as_index             = 123
  ipv4_unicast_import_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  ipv4_unicast_import_route_target_four_byte_as_format = [
    {
      four_byte_as_number = 100000
      asn4_index          = 1
      stitching           = "enable"
    }
  ]
  ipv4_unicast_import_route_target_ip_address_format = [
    {
      ipv4_address       = "1.1.1.1"
      ipv4_address_index = 1
      stitching          = "enable"
    }
  ]
  ipv4_unicast_export_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  ipv4_unicast_export_route_target_four_byte_as_format = [
    {
      four_byte_as_number = 100000
      asn4_index          = 1
      stitching           = "enable"
    }
  ]
  ipv4_unicast_export_route_target_ip_address_format = [
    {
      ipv4_address       = "1.1.1.1"
      ipv4_address_index = 1
      stitching          = "enable"
    }
  ]
  ipv6_unicast_import_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  ipv6_unicast_import_route_target_four_byte_as_format = [
    {
      four_byte_as_number = 100000
      asn4_index          = 1
      stitching           = "enable"
    }
  ]
  ipv6_unicast_import_route_target_ip_address_format = [
    {
      ipv4_address       = "1.1.1.1"
      ipv4_address_index = 1
      stitching          = "enable"
    }
  ]
  ipv6_unicast_export_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  ipv6_unicast_export_route_target_four_byte_as_format = [
    {
      four_byte_as_number = 100000
      asn4_index          = 1
      stitching           = "enable"
    }
  ]
  ipv6_unicast_export_route_target_ip_address_format = [
    {
      ipv4_address       = "1.1.1.1"
      ipv4_address_index = 1
      stitching          = "enable"
    }
  ]
}

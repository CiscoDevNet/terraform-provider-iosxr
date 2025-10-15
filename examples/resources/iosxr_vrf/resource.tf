resource "iosxr_vrf" "example" {
  vrf_name                         = "VRF4"
  description                      = "My VRF Description"
  vpn_id                           = "1000:1000"
  ipv4_unicast                     = true
  ipv4_unicast_import_route_policy = "VRF_IMPORT_POLICY_1"
  ipv4_unicast_export_route_policy = "VRF_EXPORT_POLICY_1"
  ipv6_unicast                     = true
  ipv6_unicast_import_route_policy = "VRF_IMPORT_POLICY_1"
  ipv6_unicast_export_route_policy = "VRF_EXPORT_POLICY_1"
  rd_two_byte_as_number            = "65001"
  rd_two_byte_as_index             = 123
  ipv4_unicast_import_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
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
  ipv6_unicast_import_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
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
}

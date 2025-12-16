resource "iosxr_vrf" "example" {
  vrf_name                                                  = "VRF4"
  description                                               = "My VRF Description"
  evpn_route_sync                                           = 100
  ipv4_unicast                                              = true
  ipv4_unicast_import_route_policy                          = "VRF_IMPORT_POLICY_1"
  ipv4_unicast_export_route_policy                          = "VRF_EXPORT_POLICY_1"
  ipv4_unicast_import_from_bridge_domain_advertise_as_vpn   = true
  ipv4_unicast_import_from_vrf_advertise_as_vpn             = true
  ipv4_unicast_import_from_vrf_allow_backup                 = true
  ipv4_unicast_import_from_vrf_allow_best_external          = true
  ipv4_unicast_import_from_default_vrf_advertise_as_vpn     = true
  ipv4_unicast_import_from_default_vrf_route_policy         = "VRF_IMPORT_POLICY_1"
  ipv4_unicast_export_to_vrf_allow_imported_vpn             = true
  ipv4_unicast_export_to_vrf_allow_backup                   = true
  ipv4_unicast_export_to_vrf_allow_best_external            = true
  ipv4_unicast_export_to_default_vrf_route_policy           = "VRF_EXPORT_POLICY_1"
  ipv4_unicast_export_to_default_vrf_allow_imported_vpn     = true
  ipv4_unicast_max_prefix_limit                             = 1000
  ipv4_unicast_max_prefix_threshold                         = 75
  ipv4_multicast_import_route_policy                        = "VRF_IMPORT_POLICY_1"
  ipv4_multicast_export_route_policy                        = "VRF_EXPORT_POLICY_1"
  ipv4_multicast_import_from_bridge_domain_advertise_as_vpn = true
  ipv4_multicast_import_from_vrf_advertise_as_vpn           = true
  ipv4_multicast_import_from_vrf_allow_backup               = true
  ipv4_multicast_import_from_vrf_allow_best_external        = true
  ipv4_multicast_import_from_default_vrf_advertise_as_vpn   = true
  ipv4_multicast_import_from_default_vrf_route_policy       = "VRF_IMPORT_POLICY_1"
  ipv4_multicast_export_to_vrf_allow_imported_vpn           = true
  ipv4_multicast_export_to_vrf_allow_backup                 = true
  ipv4_multicast_export_to_vrf_allow_best_external          = true
  ipv4_multicast_export_to_default_vrf_route_policy         = "VRF_EXPORT_POLICY_1"
  ipv4_multicast_export_to_default_vrf_allow_imported_vpn   = true
  ipv4_multicast_max_prefix_limit                           = 1000
  ipv4_multicast_max_prefix_threshold                       = 75
  ipv6_unicast                                              = true
  ipv6_unicast_import_route_policy                          = "VRF_IMPORT_POLICY_1"
  ipv6_unicast_export_route_policy                          = "VRF_EXPORT_POLICY_1"
  ipv6_unicast_import_from_bridge_domain_advertise_as_vpn   = true
  ipv6_unicast_import_from_vrf_advertise_as_vpn             = true
  ipv6_unicast_import_from_vrf_allow_backup                 = true
  ipv6_unicast_import_from_vrf_allow_best_external          = true
  ipv6_unicast_import_from_default_vrf_advertise_as_vpn     = true
  ipv6_unicast_import_from_default_vrf_route_policy         = "VRF_IMPORT_POLICY_1"
  ipv6_unicast_export_to_vrf_allow_imported_vpn             = true
  ipv6_unicast_export_to_vrf_allow_backup                   = true
  ipv6_unicast_export_to_vrf_allow_best_external            = true
  ipv6_unicast_export_to_default_vrf_route_policy           = "VRF_EXPORT_POLICY_1"
  ipv6_unicast_export_to_default_vrf_allow_imported_vpn     = true
  ipv6_unicast_max_prefix_limit                             = 1000
  ipv6_unicast_max_prefix_threshold                         = 75
  ipv6_multicast_import_route_policy                        = "VRF_IMPORT_POLICY_1"
  ipv6_multicast_export_route_policy                        = "VRF_EXPORT_POLICY_1"
  ipv6_multicast_import_from_bridge_domain_advertise_as_vpn = true
  ipv6_multicast_import_from_vrf_advertise_as_vpn           = true
  ipv6_multicast_import_from_vrf_allow_backup               = true
  ipv6_multicast_import_from_vrf_allow_best_external        = true
  ipv6_multicast_import_from_default_vrf_advertise_as_vpn   = true
  ipv6_multicast_import_from_default_vrf_route_policy       = "VRF_IMPORT_POLICY_1"
  ipv6_multicast_export_to_vrf_allow_imported_vpn           = true
  ipv6_multicast_export_to_vrf_allow_backup                 = true
  ipv6_multicast_export_to_vrf_allow_best_external          = true
  ipv6_multicast_export_to_default_vrf_route_policy         = "VRF_EXPORT_POLICY_1"
  ipv6_multicast_export_to_default_vrf_allow_imported_vpn   = true
  ipv6_multicast_max_prefix_limit                           = 1000
  ipv6_multicast_max_prefix_threshold                       = 75
  rd_two_byte_as_number                                     = "65001"
  rd_two_byte_as_index                                      = 123
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
  ipv4_multicast_import_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  ipv4_multicast_export_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  ipv6_multicast_import_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  ipv6_multicast_export_route_target_two_byte_as_format = [
    {
      two_byte_as_number = 65001
      asn2_index         = 1
      stitching          = "enable"
    }
  ]
  vpn_id                         = "1000:1000"
  remote_route_filtering_disable = true
}

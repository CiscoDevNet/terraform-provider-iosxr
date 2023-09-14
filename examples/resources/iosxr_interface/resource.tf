resource "iosxr_interface" "example" {
  interface_name                  = "GigabitEthernet0/0/0/1"
  l2transport                     = false
  point_to_point                  = false
  multipoint                      = false
  dampening_decay_half_life_value = 2
  ipv4_point_to_point             = true
  service_policy_input = [
    {
      name = "PMAP-IN"
    }
  ]
  service_policy_output = [
    {
      name = "PMAP-OUT"
    }
  ]
  shutdown                                                 = true
  mtu                                                      = 9000
  bandwidth                                                = 100000
  description                                              = "My Interface Description"
  load_interval                                            = 30
  vrf                                                      = "VRF1"
  ipv4_address                                             = "1.1.1.1"
  ipv4_netmask                                             = "255.255.255.0"
  ipv4_verify_unicast_source_reachable_via_type            = "any"
  ipv4_verify_unicast_source_reachable_via_allow_self_ping = true
  ipv4_verify_unicast_source_reachable_via_allow_default   = false
  ipv4_access_group_ingress_acl1                           = "ACL1"
  ipv4_access_group_ingress_hardware_count                 = true
  ipv4_access_group_ingress_interface_statistics           = true
  ipv4_access_group_ingress_compress_level                 = 0
  ipv4_access_group_egress_acl                             = "ACL1"
  ipv4_access_group_egress_hardware_count                  = true
  ipv4_access_group_egress_interface_statistics            = true
  ipv4_access_group_egress_compress_level                  = 0
  ipv6_verify_unicast_source_reachable_via_type            = "any"
  ipv6_verify_unicast_source_reachable_via_allow_self_ping = true
  ipv6_verify_unicast_source_reachable_via_allow_default   = false
  ipv6_access_group_ingress_acl1                           = "ACL2"
  ipv6_access_group_ingress_interface_statistics           = true
  ipv6_access_group_ingress_compress_level                 = 0
  ipv6_access_group_egress_acl1                            = "ACL2"
  ipv6_access_group_egress_interface_statistics            = true
  ipv6_access_group_egress_compress_level                  = 0
  ipv6_link_local_address                                  = "fe80::1"
  ipv6_link_local_zone                                     = "0"
  ipv6_autoconfig                                          = false
  ipv6_enable                                              = true
  ipv6_addresses = [
    {
      address       = "2001::1"
      prefix_length = 64
      zone          = "0"
    }
  ]
  bundle_port_priority = 100
}

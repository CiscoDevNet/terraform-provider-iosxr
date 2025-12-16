resource "iosxr_interface_bundle_ether" "example" {
  name                         = "Bundle-Ether10"
  point_to_point               = false
  multipoint                   = false
  dampening                    = true
  dampening_decay_half_life    = 2
  dampening_reuse_threshold    = 10
  dampening_suppress_threshold = 20
  dampening_max_suppress_time  = 30
  shutdown                     = false
  mtu                          = 9000
  logging_events_link_status   = true
  bandwidth                    = 100000
  description                  = "My Interface Description"
  load_interval                = 30
  vrf                          = "VRF1"
  ipv4_address                 = "192.168.1.1"
  ipv4_netmask                 = "255.255.255.0"
  ipv4_route_tag               = 100
  ipv4_algorithm               = 128
  ipv4_secondaries = [
    {
      address   = "192.168.2.1"
      netmask   = "255.255.255.0"
      route_tag = 100
      algorithm = 128
    }
  ]
  ipv4_point_to_point = true
  ipv4_mtu            = 1500
  ipv4_redirects      = true
  ipv4_mask_reply     = true
  ipv4_helper_addresses = [
    {
      address = "192.168.1.1"
      vrf     = "default"
    }
  ]
  ipv4_unreachables_disable      = true
  ipv4_access_group_ingress_acl1 = "ACL1"
  ipv4_access_group_egress_acl   = "ACL1"
  ipv6_access_group_ingress_acl1 = "ACL2"
  ipv6_access_group_egress_acl   = "ACL2"
  ipv6_enable                    = true
  ipv6_addresses = [
    {
      address       = "2001:db8:1:1::1"
      prefix_length = 64
      zone          = "0"
      route_tag     = 100
      algorithm     = 128
    }
  ]
  ipv6_link_local_address   = "fe80::1"
  ipv6_link_local_zone      = "0"
  ipv6_link_local_route_tag = 100
  ipv6_eui64_addresses = [
    {
      address       = "2001:db8:1:2::"
      prefix_length = 64
      zone          = "0"
      route_tag     = 100
      algorithm     = 128
    }
  ]
  ipv6_autoconfig                               = false
  ipv6_dhcp                                     = false
  ipv6_mtu                                      = 1280
  ipv6_unreachables_disable                     = true
  ipv6_nd_reachable_time                        = 1800
  ipv6_nd_cache_limit                           = 1000
  ipv6_nd_dad_attempts                          = 3
  ipv6_nd_unicast_ra                            = true
  ipv6_nd_managed_config_flag                   = true
  ipv6_nd_other_config_flag                     = true
  ipv6_nd_ns_interval                           = 60000
  ipv6_nd_ra_interval_max                       = 10
  ipv6_nd_ra_interval_min                       = 5
  ipv6_nd_ra_lifetime                           = 3600
  ipv6_nd_redirects                             = true
  ipv6_nd_prefix_default_no_adv                 = true
  arp_timeout                                   = 30
  arp_learning_local                            = true
  arp_gratuitous_ignore                         = true
  proxy_arp                                     = true
  bundle_minimum_active_links                   = 1
  bundle_maximum_active_links                   = 8
  bundle_shutdown                               = false
  bundle_lacp_delay                             = 1000
  bundle_lacp_fallback_timeout                  = 60
  lacp_switchover_suppress_flaps                = 150
  lacp_churn_logging                            = "both"
  lacp_cisco_enable                             = true
  lacp_cisco_enable_link_order_signaled         = true
  lacp_non_revertive                            = true
  lacp_mode                                     = "active"
  lacp_system_priority                          = 100
  lacp_system_mac                               = "00:11:22:33:44:55"
  lacp_period                                   = 500
  bfd_mode_ietf                                 = true
  bfd_address_family_ipv4_destination           = "192.168.1.2"
  bfd_address_family_ipv4_minimum_interval      = 1000
  bfd_address_family_ipv4_multiplier            = 3
  bfd_address_family_ipv4_fast_detect           = true
  bfd_address_family_ipv4_echo_minimum_interval = 100
  bfd_address_family_ipv4_timers_start          = 60
  bfd_address_family_ipv4_timers_nbr_unconfig   = 120
  bfd_address_family_ipv6_destination           = "2001:db8::2"
  bfd_address_family_ipv6_minimum_interval      = 1000
  bfd_address_family_ipv6_multiplier            = 3
  bfd_address_family_ipv6_fast_detect           = true
  bfd_address_family_ipv6_timers_start          = 60
  bfd_address_family_ipv6_timers_nbr_unconfig   = 120
  mac_address                                   = "aa:bb:cc:dd:ee:ff"
  mpls_mtu                                      = 1500
  lldp                                          = true
  lldp_transmit_disable                         = true
  lldp_receive_disable                          = true
}

resource "iosxr_interface_bvi" "example" {
  arp_gratuitous_ignore          = true
  arp_learning_local             = true
  arp_timeout                    = 30
  bandwidth                      = 100000
  dampening                      = true
  dampening_decay_half_life      = 2
  dampening_max_suppress_time    = 30
  dampening_reuse_threshold      = 10
  dampening_suppress_threshold   = 20
  description                    = "My Interface Description"
  host_routing                   = true
  ipv4_access_group_egress_acl   = "ACL1"
  ipv4_access_group_ingress_acl1 = "ACL1"
  ipv4_address                   = "192.168.1.1"
  ipv4_algorithm                 = 128
  ipv4_helper_addresses = [
    {
      address = "192.168.1.1"
      vrf     = "default"
    }
  ]
  ipv4_mask_reply     = true
  ipv4_mtu            = 1500
  ipv4_netmask        = "255.255.255.0"
  ipv4_point_to_point = true
  ipv4_redirects      = true
  ipv4_route_tag      = 100
  ipv4_secondaries = [
    {
      address   = "192.168.2.1"
      algorithm = 128
      netmask   = "255.255.255.0"
      route_tag = 100
    }
  ]
  ipv4_unreachables_disable      = true
  ipv6_access_group_egress_acl   = "ACL2"
  ipv6_access_group_ingress_acl1 = "ACL2"
  ipv6_addresses = [
    {
      address       = "2001:db8:1:1::1"
      algorithm     = 128
      prefix_length = 64
      route_tag     = 100
      zone          = "0"
    }
  ]
  ipv6_autoconfig = false
  ipv6_dhcp       = false
  ipv6_enable     = true
  ipv6_eui64_addresses = [
    {
      address       = "2001:db8:1:2::"
      algorithm     = 128
      prefix_length = 64
      route_tag     = 100
      zone          = "0"
    }
  ]
  ipv6_link_local_address                           = "fe80::1"
  ipv6_link_local_route_tag                         = 100
  ipv6_link_local_zone                              = "0"
  ipv6_mtu                                          = 1280
  ipv6_nd_cache_limit                               = 1000
  ipv6_nd_dad_attempts                              = 3
  ipv6_nd_managed_config_flag                       = true
  ipv6_nd_ns_interval                               = 60000
  ipv6_nd_other_config_flag                         = true
  ipv6_nd_prefix_default_no_adv                     = true
  ipv6_nd_ra_interval_max                           = 10
  ipv6_nd_ra_interval_min                           = 5
  ipv6_nd_ra_lifetime                               = 3600
  ipv6_nd_reachable_time                            = 1800
  ipv6_nd_redirects                                 = true
  ipv6_nd_unicast_ra                                = true
  ipv6_unreachables_disable                         = true
  load_interval                                     = 30
  logging_events_link_status                        = true
  mac_address                                       = "aa:bb:cc:dd:ee:ff"
  mtu                                               = 9000
  multipoint                                        = false
  name                                              = "100"
  point_to_point                                    = false
  proxy_arp                                         = true
  ptp                                               = true
  ptp_announce_grant_duration                       = 300
  ptp_announce_interval                             = "2"
  ptp_announce_timeout                              = 5
  ptp_clock_operation_one_step                      = true
  ptp_cos                                           = 6
  ptp_cos_event                                     = 6
  ptp_cos_general                                   = 6
  ptp_delay_asymmetry_unit_microseconds             = true
  ptp_delay_asymmetry_value                         = 1000
  ptp_delay_request_interval                        = "2"
  ptp_delay_response_grant_duration                 = 300
  ptp_delay_response_timeout                        = 3000
  ptp_dscp                                          = 46
  ptp_dscp_event                                    = 46
  ptp_dscp_general                                  = 46
  ptp_interop_domain                                = 24
  ptp_interop_egress_conversion_clock_accuracy      = 33
  ptp_interop_egress_conversion_clock_class_default = 6
  ptp_interop_egress_conversion_clock_class_mappings = [
    {
      clock_class_to_map_from = 6
      clock_class_to_map_to   = 13
    }
  ]
  ptp_interop_egress_conversion_offset_scaled_log_variance = 5
  ptp_interop_egress_conversion_priority1                  = 128
  ptp_interop_egress_conversion_priority2                  = 128
  ptp_interop_ingress_conversion_clock_accuracy            = 33
  ptp_interop_ingress_conversion_clock_class_default       = 6
  ptp_interop_ingress_conversion_clock_class_mappings = [
    {
      clock_class_to_map_from = 13
      clock_class_to_map_to   = 6
    }
  ]
  ptp_interop_ingress_conversion_offset_scaled_log_variance = 5
  ptp_interop_ingress_conversion_priority1                  = 128
  ptp_interop_ingress_conversion_priority2                  = 128
  ptp_interop_profile_g_8275_2                              = true
  ptp_ipv4_ttl                                              = 10
  ptp_ipv6_hop_limit                                        = 10
  ptp_local_priority                                        = 128
  ptp_multicast                                             = true
  ptp_multicast_mixed                                       = true
  ptp_multicast_target_address_mac_forwardable              = true
  ptp_port_state_master_only                                = true
  ptp_profile                                               = "Profile-1"
  ptp_slave_ethernets = [
    {
      address        = "00:11:22:33:44:55"
      non_negotiated = true
    }
  ]
  ptp_slave_ipv4s = [
    {
      address        = "10.2.2.2"
      non_negotiated = true
    }
  ]
  ptp_slave_ipv6s = [
    {
      address        = "2001:db8::2"
      non_negotiated = true
    }
  ]
  ptp_sync_grant_duration                  = 300
  ptp_sync_interval                        = "2"
  ptp_sync_timeout                         = 3000
  ptp_transport_ethernet                   = true
  ptp_unicast_grant_invalid_request_reduce = true
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
  shutdown = false
  vrf      = "VRF1"
}

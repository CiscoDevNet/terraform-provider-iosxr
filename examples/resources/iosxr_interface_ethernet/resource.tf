resource "iosxr_interface_ethernet" "example" {
  type                         = "GigabitEthernet"
  name                         = "0/0/0/1"
  l2transport                  = false
  point_to_point               = false
  multipoint                   = false
  dampening                    = true
  dampening_decay_half_life    = 2
  dampening_reuse_threshold    = 10
  dampening_suppress_threshold = 20
  dampening_max_suppress_time  = 30
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
  shutdown       = true
  mtu            = 9000
  bandwidth      = 100000
  description    = "My Interface Description"
  load_interval  = 30
  vrf            = "VRF1"
  ipv4_address   = "192.168.1.1"
  ipv4_netmask   = "255.255.255.0"
  ipv4_route_tag = 100
  ipv4_algorithm = 128
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
  ipv6_autoconfig               = false
  ipv6_dhcp                     = false
  ipv6_mtu                      = 1280
  ipv6_unreachables_disable     = true
  ipv6_nd_reachable_time        = 1800
  ipv6_nd_cache_limit           = 1000
  ipv6_nd_dad_attempts          = 3
  ipv6_nd_unicast_ra            = true
  ipv6_nd_managed_config_flag   = true
  ipv6_nd_other_config_flag     = true
  ipv6_nd_ns_interval           = 60000
  ipv6_nd_ra_interval_max       = 10
  ipv6_nd_ra_interval_min       = 5
  ipv6_nd_ra_lifetime           = 3600
  ipv6_nd_redirects             = true
  ipv6_nd_prefix_default_no_adv = true
  ethernet_cfm_mep_domains = [
    {
      domain_name                                        = "DOMAIN1"
      service                                            = "SERVICE1"
      mep_id                                             = 1
      propagate_remote_status                            = true
      cos                                                = 5
      loss_measurement_counters_priority_cos_range_start = 1
      loss_measurement_counters_priority_cos_range_end   = 7
      sla_operation_profile_target_mep_ids = [
        {
          profile_name = "SLA-PROFILE-1"
          mep_id       = 2
        }
      ]
      sla_operation_profile_target_mac_addresses = [
        {
          profile_name = "SLA-PROFILE-2"
          mac_address  = "00:11:22:33:44:55"
        }
      ]
    }
  ]
  ethernet_cfm_ais_transmission_up_interval                                          = "1s"
  ethernet_cfm_ais_transmission_up_cos                                               = 5
  ethernet_cfm_bandwidth_notifications_hold_off                                      = 60
  ethernet_cfm_bandwidth_notifications_wait_to_restore                               = 30
  ethernet_cfm_bandwidth_notifications_loss_threshold                                = 5
  ethernet_cfm_bandwidth_notifications_log_changes                                   = true
  frequency_synchronization                                                          = true
  frequency_synchronization_ssm_disable                                              = true
  frequency_synchronization_priority                                                 = 10
  frequency_synchronization_time_of_day_priority                                     = 10
  frequency_synchronization_quality_transmit_lowest_itu_t_option_two_generation_two  = "e-prtc"
  frequency_synchronization_quality_transmit_highest_itu_t_option_two_generation_two = "e-prtc"
  frequency_synchronization_quality_receive_lowest_itu_t_option_two_generation_two   = "e-prtc"
  frequency_synchronization_quality_receive_highest_itu_t_option_two_generation_two  = "e-prtc"
  frequency_synchronization_wait_to_restore                                          = 5
  frequency_synchronization_selection_input                                          = true
  arp_timeout                                                                        = 30
  arp_learning_local                                                                 = true
  arp_gratuitous_ignore                                                              = true
  proxy_arp                                                                          = true
  carrier_delay_up                                                                   = 100
  carrier_delay_down                                                                 = 100
  mac_address                                                                        = "aa:bb:cc:dd:ee:ff"
  mpls_mtu                                                                           = 1500
  lldp                                                                               = true
  lldp_transmit_disable                                                              = true
  lldp_receive_disable                                                               = true
  monitor_sessions = [
    {
      session_name      = "SESSION-1"
      ethernet          = true
      direction_rx_only = true
      port_level        = true
      acl               = true
      acl_ipv4_name     = "ACL1"
      acl_ipv6_name     = "ACL2"
    }
  ]
  ptp                                          = true
  ptp_profile                                  = "Profile-1"
  ptp_transport_ethernet                       = true
  ptp_clock_operation_one_step                 = true
  ptp_announce_interval                        = "2"
  ptp_announce_timeout                         = 5
  ptp_announce_grant_duration                  = 300
  ptp_sync_interval                            = "2"
  ptp_sync_grant_duration                      = 300
  ptp_sync_timeout                             = 3000
  ptp_delay_request_interval                   = "2"
  ptp_cos                                      = 6
  ptp_cos_event                                = 6
  ptp_cos_general                              = 6
  ptp_dscp                                     = 46
  ptp_dscp_event                               = 46
  ptp_dscp_general                             = 46
  ptp_ipv4_ttl                                 = 10
  ptp_ipv6_hop_limit                           = 10
  ptp_delay_asymmetry_value                    = 1000
  ptp_delay_asymmetry_unit_microseconds        = true
  ptp_delay_response_grant_duration            = 300
  ptp_delay_response_timeout                   = 3000
  ptp_unicast_grant_invalid_request_reduce     = true
  ptp_multicast                                = true
  ptp_multicast_mixed                          = true
  ptp_multicast_target_address_mac_forwardable = true
  ptp_port_state_master_only                   = true
  ptp_local_priority                           = 128
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
  ptp_slave_ethernets = [
    {
      address        = "00:11:22:33:44:55"
      non_negotiated = true
    }
  ]
  ptp_master_ipv4s = [
    {
      address         = "10.3.3.3"
      priority        = 100
      clock_class     = 6
      multicast       = true
      multicast_mixed = true
      non_negotiated  = true
      delay_asymmetry = 50
      microseconds    = true
    }
  ]
  ptp_master_ipv6s = [
    {
      address         = "2001:db8::3"
      priority        = 100
      clock_class     = 6
      multicast       = true
      multicast_mixed = true
      non_negotiated  = true
      delay_asymmetry = 50
      microseconds    = true
    }
  ]
  ptp_master_ethernets = [
    {
      address         = "aa:bb:cc:dd:ee:f4"
      priority        = 100
      clock_class     = 6
      multicast       = true
      multicast_mixed = true
      non_negotiated  = true
      delay_asymmetry = 50
      microseconds    = true
    }
  ]
  ptp_interop_profile_g_8275_2                             = true
  ptp_interop_domain                                       = 24
  ptp_interop_egress_conversion_priority1                  = 128
  ptp_interop_egress_conversion_priority2                  = 128
  ptp_interop_egress_conversion_clock_accuracy             = 33
  ptp_interop_egress_conversion_offset_scaled_log_variance = 5
  ptp_interop_egress_conversion_clock_class_default        = 6
  ptp_interop_egress_conversion_clock_class_mappings = [
    {
      clock_class_to_map_from = 6
      clock_class_to_map_to   = 13
    }
  ]
  ptp_interop_ingress_conversion_priority1                  = 128
  ptp_interop_ingress_conversion_priority2                  = 128
  ptp_interop_ingress_conversion_clock_accuracy             = 33
  ptp_interop_ingress_conversion_offset_scaled_log_variance = 5
  ptp_interop_ingress_conversion_clock_class_default        = 6
  ptp_interop_ingress_conversion_clock_class_mappings = [
    {
      clock_class_to_map_from = 13
      clock_class_to_map_to   = 6
    }
  ]
}

resource "iosxr_ptp_profile" "example" {
  announce_grant_duration                       = 300
  announce_timeout                              = 5
  clock_operation_one_step                      = true
  cos                                           = 6
  cos_event                                     = 6
  cos_general                                   = 6
  delay_asymmetry_unit_microseconds             = true
  delay_asymmetry_value                         = 1000
  delay_request_interval                        = "2"
  delay_response_grant_duration                 = 300
  delay_response_timeout                        = 3000
  dscp                                          = 46
  dscp_event                                    = 46
  dscp_general                                  = 46
  interop_domain                                = 24
  interop_egress_conversion_clock_accuracy      = 33
  interop_egress_conversion_clock_class_default = 6
  interop_egress_conversion_clock_class_mappings = [
    {
      clock_class_to_map_from = 6
      clock_class_to_map_to   = 13
    }
  ]
  interop_egress_conversion_offset_scaled_log_variance = 5
  interop_egress_conversion_priority1                  = 128
  interop_egress_conversion_priority2                  = 128
  interop_ingress_conversion_clock_accuracy            = 33
  interop_ingress_conversion_clock_class_default       = 6
  interop_ingress_conversion_clock_class_mappings = [
    {
      clock_class_to_map_from = 13
      clock_class_to_map_to   = 6
    }
  ]
  interop_ingress_conversion_offset_scaled_log_variance = 5
  interop_ingress_conversion_priority1                  = 128
  interop_ingress_conversion_priority2                  = 128
  interop_profile_g_8275_2                              = true
  ipv4_ttl                                              = 10
  ipv6_hop_limit                                        = 10
  master_ethernets = [
    {
      address         = "aa:bb:cc:dd:ee:f4"
      clock_class     = 6
      delay_asymmetry = 50
      microseconds    = true
      multicast       = true
      multicast_mixed = true
      non_negotiated  = true
      priority        = 100
    }
  ]
  master_ipv4s = [
    {
      address         = "10.3.3.3"
      clock_class     = 6
      delay_asymmetry = 50
      microseconds    = true
      multicast       = true
      multicast_mixed = true
      non_negotiated  = true
      priority        = 100
    }
  ]
  master_ipv6s = [
    {
      address         = "2001:db8::3"
      clock_class     = 6
      delay_asymmetry = 50
      microseconds    = true
      multicast       = true
      multicast_mixed = true
      non_negotiated  = true
      priority        = 100
    }
  ]
  multicast                                                       = true
  multicast_mixed                                                 = true
  multicast_target_address_ethernet_mac_address_01_1b_19_00_00_00 = true
  port_state_master_only                                          = true
  profile_name                                                    = "Profile-1"
  slave_ethernets = [
    {
      address        = "00:11:22:33:44:55"
      non_negotiated = true
    }
  ]
  slave_ipv4s = [
    {
      address        = "10.2.2.2"
      non_negotiated = true
    }
  ]
  slave_ipv6s = [
    {
      address        = "2001:db8::2"
      non_negotiated = true
    }
  ]
  source_ipv4_address                  = "10.1.1.1"
  source_ipv6_address                  = "2001:db8::1"
  sync_grant_duration                  = 300
  sync_interval                        = "2"
  sync_timeout                         = 3000
  transport_ethernet                   = true
  unicast_grant_invalid_request_reduce = true
}

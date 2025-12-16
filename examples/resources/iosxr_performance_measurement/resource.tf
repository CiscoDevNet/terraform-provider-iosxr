resource "iosxr_performance_measurement" "example" {
  source_address_ipv4                                                          = "10.1.1.1"
  source_address_ipv6                                                          = "2001:db8::1"
  protocol_twamp_light_measurement_delay_unauthenticated_querier_dst_port      = 862
  protocol_twamp_light_measurement_delay_unauthenticated_querier_src_port      = 10000
  protocol_twamp_light_measurement_delay_unauthenticated_ipv4_timestamp1_label = 1000
  protocol_twamp_light_measurement_delay_unauthenticated_ipv4_timestamp2_label = 2000
  protocol_twamp_light_measurement_delay_unauthenticated_ipv6_timestamp1_label = 1500
  protocol_twamp_light_measurement_delay_unauthenticated_ipv6_timestamp2_label = 2500
  protocol_twamp_light_measurement_delay_responder_allow_querier_ipv4_prefixes = [
    {
      address = "10.1.0.0"
      length  = 24
    }
  ]
  protocol_twamp_light_measurement_delay_responder_allow_querier_ipv4_addresses = [
    {
      address = "10.1.1.100"
    }
  ]
  protocol_twamp_light_measurement_delay_responder_allow_querier_ipv6_prefixes = [
    {
      address = "2001:db8:1:1::"
      length  = 64
    }
  ]
  protocol_twamp_light_measurement_delay_responder_allow_querier_ipv6_addresses = [
    {
      address = "2001:db8::100"
    }
  ]
  protocol_twamp_light_measurement_delay_querier_allow_responder_ipv4_prefixes = [
    {
      address = "10.2.0.0"
      length  = 24
    }
  ]
  protocol_twamp_light_measurement_delay_querier_allow_responder_ipv4_addresses = [
    {
      address = "10.2.1.100"
    }
  ]
  protocol_twamp_light_measurement_delay_querier_allow_responder_ipv6_prefixes = [
    {
      address = "2001:db8:1:2::"
      length  = 64
    }
  ]
  protocol_twamp_light_measurement_delay_querier_allow_responder_ipv6_addresses = [
    {
      address = "2001:db8::100"
    }
  ]
}

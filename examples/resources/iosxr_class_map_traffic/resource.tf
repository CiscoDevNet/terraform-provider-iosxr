resource "iosxr_class_map_traffic" "example" {
  class_map_name = "CM-TRAFFIC"
  match_all      = true
  description    = "Traffic Class Map"
  match_destination_address_ipv4 = [
    {
      address = "10.1.1.0"
      netmask = "255.255.255.0"
    }
  ]
  match_destination_address_ipv6 = [
    {
      address       = "2001:db8:1:1::"
      prefix_length = 64
    }
  ]
  match_destination_port             = ["80"]
  match_dscp                         = ["46"]
  match_dscp_ipv4                    = ["46"]
  match_dscp_ipv6                    = ["46"]
  match_fragment_type_dont_fragment  = true
  match_fragment_type_first_fragment = true
  match_fragment_type_is_fragment    = true
  match_fragment_type_last_fragment  = true
  match_ipv4_icmp_code               = ["1-5"]
  match_ipv4_icmp_type               = ["1-5"]
  match_ipv6_icmp_code               = ["1-5"]
  match_ipv6_icmp_type               = ["1-5"]
  match_packet_length                = ["1000-1200"]
  match_protocol                     = ["udp"]
  match_source_address_ipv4 = [
    {
      address = "10.1.2.0"
      netmask = "255.255.255.0"
    }
  ]
  match_source_address_ipv6 = [
    {
      address       = "2001:db8:1:2::"
      prefix_length = 64
    }
  ]
  match_source_port  = ["80"]
  match_tcp_flag     = 5
  match_tcp_flag_any = true
}

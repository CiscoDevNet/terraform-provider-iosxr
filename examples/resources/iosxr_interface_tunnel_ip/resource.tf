resource "iosxr_interface_tunnel_ip" "example" {
  bandwidth    = 100000
  description  = "My Interface Description"
  ipv4_address = "192.168.1.1"
  ipv4_netmask = "255.255.255.0"
  ipv6_addresses = [
    {
      address       = "2001:db8::1"
      prefix_length = 64
      zone          = "0"
    }
  ]
  ipv6_autoconfig             = false
  ipv6_enable                 = true
  ipv6_link_local_address     = "fe80::1"
  ipv6_link_local_zone        = "0"
  load_interval               = 30
  logging_events_link_status  = true
  mtu                         = 1400
  name                        = "100"
  shutdown                    = false
  tunnel_bfd_destination_ipv4 = "192.168.1.2"
  tunnel_bfd_minimum_interval = 1000
  tunnel_bfd_multiplier       = 3
  tunnel_bfd_period           = 20
  tunnel_bfd_retry            = 3
  tunnel_destination_ipv4     = "192.168.1.2"
  tunnel_df_disable           = true
  tunnel_mode_gre_ipv4        = true
  tunnel_source_ipv4          = "192.168.1.1"
  tunnel_tos                  = 5
  tunnel_ttl_value            = 10
  tunnel_vrf                  = "VRF1"
  vrf                         = "VRF1"
}

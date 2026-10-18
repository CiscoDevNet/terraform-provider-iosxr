resource "iosxr_interface_loopback" "example" {
  dampening                    = true
  dampening_decay_half_life    = 2
  dampening_max_suppress_time  = 30
  dampening_reuse_threshold    = 10
  dampening_suppress_threshold = 20
  description                  = "My Interface Description"
  ipv4_address                 = "192.168.1.1"
  ipv4_algorithm               = 128
  ipv4_helper_addresses = [
    {
      address = "192.168.1.1"
      vrf     = "default"
    }
  ]
  ipv4_mask_reply = true
  ipv4_mtu        = 1500
  ipv4_netmask    = "255.255.255.0"
  ipv4_redirects  = true
  ipv4_route_tag  = 100
  ipv4_secondaries = [
    {
      address   = "192.168.2.1"
      algorithm = 128
      netmask   = "255.255.255.0"
      route_tag = 100
    }
  ]
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
  ipv6_link_local_address       = "fe80::1"
  ipv6_link_local_route_tag     = 100
  ipv6_link_local_zone          = "0"
  ipv6_mtu                      = 1280
  ipv6_nd_cache_limit           = 1000
  ipv6_nd_prefix_default_no_adv = true
  load_interval                 = 30
  name                          = "100"
  shutdown                      = true
  vrf                           = "VRF1"
}

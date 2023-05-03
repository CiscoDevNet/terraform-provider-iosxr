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
  shutdown                = true
  mtu                     = 9000
  bandwidth               = 100000
  description             = "My Interface Description"
  load_interval           = 30
  vrf                     = "VRF1"
  ipv4_address            = "1.1.1.1"
  ipv4_netmask            = "255.255.255.0"
  ipv6_link_local_address = "fe80::1"
  ipv6_link_local_zone    = "0"
  ipv6_autoconfig         = false
  ipv6_enable             = true
  ipv6_addresses = [
    {
      address       = "2001::1"
      prefix_length = 64
      zone          = "0"
    }
  ]
}

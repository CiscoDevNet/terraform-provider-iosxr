resource "iosxr_router_vrrp_interface_ipv6" "example" {
  accept_mode_disable       = true
  address_linklocal         = "fe80::2"
  bfd_fast_detect_peer_ipv6 = "3::3"
  global_addresses = [
    {
      address = "2001:db8::1"
    }
  ]
  interface_name              = "GigabitEthernet0/0/0/2"
  name                        = "TEST2"
  preempt_delay               = 255
  preempt_disable             = false
  priority                    = 250
  timer_advertisement_seconds = 10
  timer_force                 = true
  track_interfaces = [
    {
      interface_name     = "GigabitEthernet0/0/0/4"
      priority_decrement = 12
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT"
      priority_decrement = 22
    }
  ]
  unicast_peer = "fe80::3"
  vrrp_id      = 124
}

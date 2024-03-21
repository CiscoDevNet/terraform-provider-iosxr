resource "iosxr_router_vrrp_interface_ipv6" "example" {
  interface_name               = "GigabitEthernet0/0/0/2"
  vrrp_id                      = 124
  address_linklocal_autoconfig = true
  priority                     = 250
  name                         = "TEST2"
  timer_advertisement_seconds  = 10
  timer_force                  = true
  preempt_disable              = false
  preempt_delay                = 255
  accept_mode_disable          = true
  track_interfaces = [
    {
      interface_name     = "GigabitEthernet0/0/0/5"
      priority_decrement = 12
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT"
      priority_decrement = 22
    }
  ]
  bfd_fast_detect_peer_ipv6 = "3::3"
}

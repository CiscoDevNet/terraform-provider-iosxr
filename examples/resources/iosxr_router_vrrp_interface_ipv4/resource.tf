resource "iosxr_router_vrrp_interface_ipv4" "example" {
  accept_mode_disable       = false
  address                   = "1.1.1.1"
  bfd_fast_detect_peer_ipv4 = "33.33.33.3"
  interface_name            = "GigabitEthernet0/0/0/1"
  name                      = "TEST"
  preempt_delay             = 255
  preempt_disable           = false
  priority                  = 250
  secondary_addresses = [
    {
      address = "2.2.2.2"
    }
  ]
  text_authentication         = "password"
  timer_advertisement_seconds = 123
  timer_force                 = false
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
  unicast_peer = "1.1.1.2"
  version      = 2
  vrrp_id      = 123
}

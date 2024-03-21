resource "iosxr_router_vrrp_interface_ipv4" "example" {
  interface_name              = "GigabitEthernet0/0/0/1"
  vrrp_id                     = 123
  version                     = 2
  address                     = "1.1.1.1"
  priority                    = 250
  name                        = "TEST"
  text_authentication         = "7"
  timer_advertisement_seconds = 123
  timer_force                 = false
  preempt_disable             = false
  preempt_delay               = 255
  accept_mode_disable         = false
  track_interfaces = [
    {
      interface_name     = "GigabitEthernet0/0/0/1"
      priority_decrement = 12
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT"
      priority_decrement = 22
    }
  ]
  bfd_fast_detect_peer_ipv4 = "33.33.33.3"
}

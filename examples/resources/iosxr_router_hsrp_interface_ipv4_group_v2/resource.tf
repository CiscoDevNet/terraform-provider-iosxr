resource "iosxr_router_hsrp_interface_ipv4_group_v2" "example" {
  interface_name            = "GigabitEthernet0/0/0/1"
  group_id                  = 2345
  address                   = "33.33.33.3"
  address_learn             = false
  priority                  = 133
  name                      = "NAME22"
  preempt_delay             = 3100
  bfd_fast_detect_peer_ipv4 = "45.45.45.4"
  track_interfaces = [
    {
      track_name         = "GigabitEthernet0/0/0/7"
      priority_decrement = 66
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT2"
      priority_decrement = 77
    }
  ]
}

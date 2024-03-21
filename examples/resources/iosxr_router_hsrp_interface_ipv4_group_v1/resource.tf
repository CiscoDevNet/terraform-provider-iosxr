resource "iosxr_router_hsrp_interface_ipv4_group_v1" "example" {
  interface_name                 = "GigabitEthernet0/0/0/1"
  group_id                       = 123
  address                        = "22.22.1.1"
  address_learn                  = false
  priority                       = 124
  name                           = "NAME11"
  preempt_delay                  = 3200
  bfd_fast_detect_peer_ipv4      = "44.44.4.4"
  bfd_fast_detect_peer_interface = "GigabitEthernet0/0/0/7"
  track_interfaces = [
    {
      track_name         = "GigabitEthernet0/0/0/1"
      priority_decrement = 166
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT1"
      priority_decrement = 177
    }
  ]
}

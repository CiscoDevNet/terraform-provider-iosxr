resource "iosxr_router_hsrp_interface_ipv4_group_v1" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  group_id = 123
  address = "22.22.1.1"
  address_learn = false
  secondary_ipv4_addresses = [
    {
      address = "2.2.2.2"
    }
  ]
  priority = 124
  preempt_delay = 3200
  track_interfaces = [
    {
      track_name = "GigabitEthernet0/0/0/1"
      priority_decrement = 166
    }
  ]
  track_objects = [
    {
      object_name = "OBJECT1"
      priority_decrement = 177
    }
  ]
  timers_msec = 100
  timers_msec_holdtime = 300
  mac_address = "00:01:00:02:00:02"
  name = "NAME11"
  bfd_fast_detect_peer_ipv4 = "44.44.4.4"
  bfd_fast_detect_peer_interface = "GigabitEthernet0/0/0/1"
}

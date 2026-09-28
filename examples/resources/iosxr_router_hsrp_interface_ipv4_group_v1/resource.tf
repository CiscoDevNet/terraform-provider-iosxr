resource "iosxr_router_hsrp_interface_ipv4_group_v1" "example" {
  address                        = "22.22.1.1"
  address_learn                  = false
  bfd_fast_detect_peer_interface = "GigabitEthernet0/0/0/1"
  bfd_fast_detect_peer_ipv4      = "44.44.4.4"
  group_id                       = 123
  interface_name                 = "GigabitEthernet0/0/0/1"
  mac_address                    = "00:01:00:02:00:02"
  name                           = "NAME11"
  preempt_delay                  = 3200
  priority                       = 124
  secondary_ipv4_addresses = [
    {
      address = "2.2.2.2"
    }
  ]
  timers_msec          = 100
  timers_msec_holdtime = 300
  track_interfaces = [
    {
      priority_decrement = 166
      track_name         = "GigabitEthernet0/0/0/1"
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT1"
      priority_decrement = 177
    }
  ]
}

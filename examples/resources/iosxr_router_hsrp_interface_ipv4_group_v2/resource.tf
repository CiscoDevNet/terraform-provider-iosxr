resource "iosxr_router_hsrp_interface_ipv4_group_v2" "example" {
  address                        = "33.33.33.3"
  address_learn                  = false
  bfd_fast_detect_peer_interface = "GigabitEthernet0/0/0/1"
  bfd_fast_detect_peer_ipv4      = "45.45.45.4"
  group_id                       = 2345
  interface_name                 = "GigabitEthernet0/0/0/1"
  mac_address                    = "00:01:00:02:00:02"
  name                           = "NAME22"
  preempt_delay                  = 3100
  priority                       = 133
  secondary_ipv4_addresses = [
    {
      address = "10.10.1.2"
    }
  ]
  timers_seconds          = 10
  timers_seconds_holdtime = 30
  track_interfaces = [
    {
      priority_decrement = 66
      track_name         = "GigabitEthernet0/0/0/7"
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT2"
      priority_decrement = 77
    }
  ]
}

resource "iosxr_router_hsrp_interface_ipv6_group_v2" "example" {
  interface_name = "GigabitEthernet0/0/0/2"
  group_id       = 4000
  addresses = [
    {
      address = "2001:db8:cafe:2100::bad1:1010"
    }
  ]
  address_link_local_autoconfig                   = true
  address_link_local_autoconfig_legacy_compatible = true
  priority                                        = 244
  preempt_delay                                   = 256
  track_interfaces = [
    {
      track_name         = "GigabitEthernet0/0/0/4"
      priority_decrement = 244
    }
  ]
  track_objects = [
    {
      object_name        = "TOBJ2"
      priority_decrement = 10
    }
  ]
  timers_msec                    = 100
  timers_msec_holdtime           = 300
  mac_address                    = "00:01:00:02:00:02"
  name                           = "gp2"
  bfd_fast_detect_peer_ipv6      = "fe80::240:d0ff:fe48:4672"
  bfd_fast_detect_peer_interface = "GigabitEthernet0/0/0/3"
}

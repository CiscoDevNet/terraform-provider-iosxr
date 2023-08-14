resource "iosxr_router_hsrp_interface_address_family_ipv4" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  group_number_version_1s = [
    {
      group_number_version_1_id      = 123
      address_ipv4_address           = "22.22.1.1"
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
  ]
  group_number_version_2s = [
    {
      group_number_version_2_id = 2345
      address_ipv4_address      = "33.33.33.3"
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
  ]
}

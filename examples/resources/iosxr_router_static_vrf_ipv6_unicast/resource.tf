resource "iosxr_router_static_vrf_ipv6_unicast" "example" {
  vrf_name       = "VRF2"
  prefix_address = "1::"
  prefix_length  = 64
  nexthop_interfaces = [
    {
      interface_name  = "GigabitEthernet0/0/0/1"
      description     = "interface-description"
      tag             = 100
      distance_metric = 122
      permanent       = true
      metric          = 10
    }
  ]
  nexthop_interface_addresses = [
    {
      interface_name                   = "GigabitEthernet0/0/0/2"
      address                          = "2::2"
      description                      = "interface-description"
      tag                              = 103
      distance_metric                  = 144
      permanent                        = true
      metric                           = 10
      bfd_fast_detect_minimum_interval = 100
      bfd_fast_detect_multiplier       = 3
    }
  ]
  nexthop_addresses = [
    {
      address         = "3::3"
      description     = "ip-description"
      tag             = 104
      distance_metric = 155
      track           = "TRACK1"
      metric          = 10
    }
  ]
  vrfs = [
    {
      vrf_name = "VRF1"
      nexthop_interfaces = [
        {
          interface_name  = "GigabitEthernet0/0/0/3"
          description     = "interface-description"
          tag             = 100
          distance_metric = 122
          permanent       = true
          metric          = 10
        }
      ]
      nexthop_interface_addresses = [
        {
          interface_name  = "GigabitEthernet0/0/0/4"
          address         = "2::2"
          description     = "interface-description"
          tag             = 103
          distance_metric = 144
          permanent       = true
          metric          = 10
        }
      ]
      nexthop_addresses = [
        {
          address         = "3::3"
          description     = "ip-description"
          tag             = 104
          distance_metric = 155
          track           = "TRACK1"
          metric          = 10
        }
      ]
    }
  ]
}

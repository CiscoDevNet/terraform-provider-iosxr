resource "iosxr_router_static_ipv6_unicast" "example" {
  nexthop_addresses = [
    {
      address                          = "3::3"
      bfd_fast_detect_minimum_interval = 100
      bfd_fast_detect_multiplier       = 3
      description                      = "ip-description"
      distance_metric                  = 155
      metric                           = 10
      tag                              = 104
      track                            = "TRACK1"
    }
  ]
  nexthop_interface_addresses = [
    {
      address                          = "2::2"
      bfd_fast_detect_minimum_interval = 100
      bfd_fast_detect_multiplier       = 3
      description                      = "interface-description"
      distance_metric                  = 144
      interface_name                   = "GigabitEthernet0/0/0/2"
      metric                           = 10
      tag                              = 103
    }
  ]
  nexthop_interfaces = [
    {
      description     = "interface-description"
      distance_metric = 122
      interface_name  = "GigabitEthernet0/0/0/1"
      metric          = 10
      permanent       = true
      tag             = 100
    }
  ]
  prefix_address = "1::"
  prefix_length  = 64
  sr_policies = [
    {
      description     = "interface-description"
      distance_metric = 144
      metric          = 10
      sr_policy_name  = "sr_te_policy_1"
      tag             = 103
      track           = "TRACK1"
    }
  ]
  vrfs = [
    {
      nexthop_addresses = [
        {
          address                          = "3::3"
          bfd_fast_detect_minimum_interval = 100
          bfd_fast_detect_multiplier       = 3
          description                      = "ip-description"
          distance_metric                  = 155
          metric                           = 10
          tag                              = 104
          track                            = "TRACK1"
        }
      ]
      nexthop_interface_addresses = [
        {
          address                          = "2::2"
          bfd_fast_detect_minimum_interval = 100
          bfd_fast_detect_multiplier       = 3
          description                      = "interface-description"
          distance_metric                  = 144
          interface_name                   = "GigabitEthernet0/0/0/4"
          metric                           = 10
          tag                              = 103
        }
      ]
      nexthop_interfaces = [
        {
          description     = "interface-description"
          distance_metric = 122
          interface_name  = "GigabitEthernet0/0/0/3"
          metric          = 10
          permanent       = true
          tag             = 100
        }
      ]
      sr_policies = [
        {
          description     = "interface-description"
          distance_metric = 144
          metric          = 10
          sr_policy_name  = "sr_te_policy_1"
          tag             = 103
          track           = "TRACK1"
        }
      ]
      vrf_name = "VRF1"
    }
  ]
}

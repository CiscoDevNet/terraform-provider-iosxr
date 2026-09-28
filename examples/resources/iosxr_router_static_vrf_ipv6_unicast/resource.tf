resource "iosxr_router_static_vrf_ipv6_unicast" "example" {
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
  vrf_name = "VRF2"
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

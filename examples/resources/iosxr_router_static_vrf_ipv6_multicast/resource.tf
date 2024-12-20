resource "iosxr_router_static_vrf_ipv6_multicast" "example" {
  vrf_name       = "VRF2"
  prefix_address = "1::"
  prefix_length  = 64
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

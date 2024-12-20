resource "iosxr_router_static_vrf_ipv4_multicast" "example" {
  vrf_name       = "VRF2"
  prefix_address = "100.0.1.0"
  prefix_length  = 24
  nexthop_addresses = [
    {
      address         = "100.0.2.0"
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
          address         = "100.0.2.0"
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

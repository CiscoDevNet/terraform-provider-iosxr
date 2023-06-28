resource "iosxr_router_static_ipv4_multicast" "example" {
  prefix_address = "100.0.1.0"
  prefix_length  = 24
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
      interface_name  = "GigabitEthernet0/0/0/2"
      address         = "11.11.11.1"
      description     = "interface-description"
      tag             = 103
      distance_metric = 144
      permanent       = true
      metric          = 10
    }
  ]
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
          address         = "11.11.11.1"
          description     = "interface-description"
          tag             = 103
          distance_metric = 144
          permanent       = true
          metric          = 10
        }
      ]
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

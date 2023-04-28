resource "iosxr_router_static" "example" {
  prefix_address = "100.0.1.0"
  prefix_length  = 24
  nexthop_interfaces = [
    {
      interface_name  = "GigabitEthernet0/0/0/1"
      description     = "interface-description"
      tag             = 100
      distance_metric = 122
    }
  ]
  nexthop_interface_addresses = [
    {
      interface_name  = "GigabitEthernet0/0/0/2"
      address         = "11.11.11.1"
      description     = "interface-description"
      tag             = 103
      distance_metric = 144
    }
  ]
  nexthop_addresses = [
    {
      address         = "100.0.2.0"
      description     = "ip-description"
      tag             = 104
      distance_metric = 155
    }
  ]
}

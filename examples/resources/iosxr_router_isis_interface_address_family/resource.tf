resource "iosxr_router_isis_interface_address_family" "example" {
  process_id              = "P1"
  interface_name          = "GigabitEthernet0/0/0/1"
  af_name                 = "ipv4"
  saf_name                = "unicast"
  fast_reroute_per_prefix = true
  fast_reroute_levels = [
    {
      level_number = 1
      per_prefix   = true
    }
  ]
  tag                           = 100
  advertise_prefix_route_policy = "ROUTE_POLICY_1"
  advertise_prefix_route_policy_levels = [
    {
      level_number = 1
      route_policy = "ROUTE_POLICY_2"
    }
  ]
  metric_default = 500
}

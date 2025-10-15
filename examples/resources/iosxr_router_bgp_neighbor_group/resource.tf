resource "iosxr_router_bgp_neighbor_group" "example" {
  as_number                                   = "65001"
  name                                        = "GROUP1"
  remote_as                                   = "65001"
  description                                 = "My Neighbor Group Description"
  update_source                               = "Loopback0"
  advertisement_interval_seconds              = 10
  bfd_minimum_interval                        = 3
  bfd_multiplier                              = 4
  bfd_fast_detect                             = true
  bfd_fast_detect_strict_mode                 = false
  password                                    = "12341C2713181F13253920"
  password_inheritance_disable                = false
  timers_keepalive_interval                   = 10
  timers_holdtime                             = 30
  timers_holdtime_minimum_acceptable_holdtime = 30
  address_families = [
    {
      af_name                                    = "ipv4-labeled-unicast"
      soft_reconfiguration_inbound_always        = true
      next_hop_self                              = true
      next_hop_self_inheritance_disable          = true
      route_reflector_client                     = true
      route_reflector_client_inheritance_disable = true
      route_policy_in                            = "ROUTE_POLICY_1"
      route_policy_out                           = "ROUTE_POLICY_1"
    }
  ]
}

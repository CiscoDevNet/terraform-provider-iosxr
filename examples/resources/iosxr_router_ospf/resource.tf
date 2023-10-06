resource "iosxr_router_ospf" "example" {
  process_name                              = "OSPF1"
  mpls_ldp_sync                             = false
  hello_interval                            = 10
  dead_interval                             = 40
  priority                                  = 10
  mtu_ignore_enable                         = true
  mtu_ignore_disable                        = false
  passive_enable                            = false
  passive_disable                           = true
  router_id                                 = "10.11.12.13"
  redistribute_connected                    = true
  redistribute_connected_tag                = 1
  redistribute_connected_metric_type        = "1"
  redistribute_static                       = true
  redistribute_static_tag                   = 2
  redistribute_static_metric_type           = "1"
  bfd_fast_detect                           = true
  bfd_minimum_interval                      = 300
  bfd_multiplier                            = 3
  default_information_originate             = true
  default_information_originate_always      = true
  default_information_originate_metric_type = 1
  auto_cost_reference_bandwidth             = 100000
  auto_cost_disable                         = false
  segment_routing_mpls                      = true
  segment_routing_sr_prefer                 = true
  areas = [
    {
      area_id = "0"
    }
  ]
  redistribute_bgp = [
    {
      as_number   = "65001"
      tag         = 3
      metric_type = "1"
    }
  ]
  redistribute_isis = [
    {
      instance_name = "P1"
      level_1       = true
      level_2       = false
      level_1_2     = false
      tag           = 3
      metric_type   = "1"
    }
  ]
  redistribute_ospf = [
    {
      instance_name       = "OSPF2"
      match_internal      = true
      match_external      = false
      match_nssa_external = false
      tag                 = 4
      metric_type         = "1"
    }
  ]
}

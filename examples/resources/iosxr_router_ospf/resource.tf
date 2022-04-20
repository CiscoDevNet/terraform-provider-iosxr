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
}

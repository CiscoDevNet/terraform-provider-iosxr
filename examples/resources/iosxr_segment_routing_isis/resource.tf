resource "iosxr_segment_routing_isis" "example" {
  process_id                          = "P1"
  af_name                             = "ipv6"
  saf_name                            = "unicast"
  metric_style_wide                   = true
  microloop_avoidance_segment_routing = true
  router_id_interface_name            = "Loopback0"
  locators = [
    {
      locator_name = "AlgoLocator"
      level        = 1
    }
  ]
}

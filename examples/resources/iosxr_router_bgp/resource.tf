resource "iosxr_router_bgp" "example" {
  as_number                     = "65001"
  default_information_originate = true
  default_metric                = 125
  timers_bgp_keepalive_interval = 5
  timers_bgp_holdtime           = "20"
  bfd_minimum_interval          = 10
  bfd_multiplier                = 4
}

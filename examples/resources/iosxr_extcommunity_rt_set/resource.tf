resource "iosxr_extcommunity_rt_set" "example" {
  set_name = "ROUTE1"
  rpl      = "extcommunity-set rt ROUTE1\nend-set\n"
}

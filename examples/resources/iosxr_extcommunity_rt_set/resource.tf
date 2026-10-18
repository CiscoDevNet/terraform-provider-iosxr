resource "iosxr_extcommunity_rt_set" "example" {
  rpl      = "extcommunity-set rt ROUTE1\nend-set\n"
  set_name = "ROUTE1"
}

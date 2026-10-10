resource "iosxr_extcommunity_cost_set" "example" {
  rpl      = "extcommunity-set cost COST2\nend-set\n"
  set_name = "COST2"
}

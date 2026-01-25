resource "iosxr_extcommunity_cost_set" "example" {
  set_name = "COST2"
  rpl      = "extcommunity-set cost COST2\nend-set\n"
}

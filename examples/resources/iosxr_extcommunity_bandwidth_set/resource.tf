resource "iosxr_extcommunity_bandwidth_set" "example" {
  set_name = "BANDWIDTH1"
  rpl = "extcommunity-set bandwidth BANDWIDTH1\n  1234:5678,\n  1234:8765\nend-set\n"
}

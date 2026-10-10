resource "iosxr_extcommunity_opaque_set" "example" {
  rpl      = "extcommunity-set opaque BLUE\n  100\nend-set\n"
  set_name = "BLUE"
}

resource "iosxr_extcommunity_opaque_set" "example" {
  set_name = "BLUE"
  rpl      = "extcommunity-set opaque BLUE\n  100\nend-set\n"
}

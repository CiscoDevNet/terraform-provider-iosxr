resource "iosxr_extcommunity_set" "example" {
  set_name                          = "BLUE"
  rpl_extended_community_opaque_set = "extcommunity-set opaque BLUE\n  100\nend-set\n"
}

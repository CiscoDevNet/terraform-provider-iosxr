resource "iosxr_extcommunity_soo_set" "example" {
  rpl      = "extcommunity-set soo SITE1\nend-set\n"
  set_name = "SITE1"
}

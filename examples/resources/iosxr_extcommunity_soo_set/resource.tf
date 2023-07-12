resource "iosxr_extcommunity_soo_set" "example" {
  set_name = "SITE1"
  rpl      = "extcommunity-set soo SITE1\nend-set\n"
}

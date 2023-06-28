resource "iosxr_community_set" "example" {
  set_name = "TEST11"
  rpl      = "community-set TEST11\nend-set\n"
}

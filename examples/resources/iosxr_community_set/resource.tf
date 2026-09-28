resource "iosxr_community_set" "example" {
  rpl      = "community-set TEST11\nend-set\n"
  set_name = "TEST11"
}

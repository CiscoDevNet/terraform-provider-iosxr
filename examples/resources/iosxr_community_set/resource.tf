resource "iosxr_community_set" "example" {
  set_name = "WORD"
  rpl      = "community-set WORD\nend-set\n"
}

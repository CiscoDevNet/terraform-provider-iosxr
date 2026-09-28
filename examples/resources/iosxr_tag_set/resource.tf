resource "iosxr_tag_set" "example" {
  rpl      = "tag-set TAG_SET_1\n  4297\nend-set\n"
  set_name = "TAG_SET_1"
}

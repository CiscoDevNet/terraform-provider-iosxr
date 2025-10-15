resource "iosxr_tag_set" "example" {
  set_name    = "TAG_SET_1"
  rpl_tag_set = "tag-set TAG_SET_1\n  4297\nend-set\n"
}

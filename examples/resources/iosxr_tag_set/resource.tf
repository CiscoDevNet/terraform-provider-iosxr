resource "iosxr_tag_set" "example" {
  set_name    = "TEST"
  rpl_tag_set = "tag-set TEST\n  4297\nend-set\n"
}

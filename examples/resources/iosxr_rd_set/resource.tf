resource "iosxr_rd_set" "example" {
  set_name = "set1"
  rpl      = "rd-set set1\nend-set\n"
}

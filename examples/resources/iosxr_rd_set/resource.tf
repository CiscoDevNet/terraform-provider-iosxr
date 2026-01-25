resource "iosxr_rd_set" "example" {
  set_name = "set1"
  rpl = "rd-set set1\n  65001:1,\n  123456:1,\n  192.0.2.1:1,\n  match any\nend-set\n"
}

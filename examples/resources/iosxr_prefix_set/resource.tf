resource "iosxr_prefix_set" "example" {
  set_name = "PREFIX_SET_1"
  rpl      = "prefix-set PREFIX_SET_1\n  10.1.1.0/26 ge 26,\n  10.1.2.0/26 ge 26\nend-set\n"
}

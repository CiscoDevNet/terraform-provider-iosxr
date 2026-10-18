resource "iosxr_mac_set" "example" {
  rpl      = "mac-set MAC1\n  0011.2233.4455,\n  2233.4455.5566\nend-set\n"
  set_name = "MAC1"
}

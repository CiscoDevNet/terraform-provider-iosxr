resource "iosxr_extcommunity_seg_nh_set" "example" {
  set_name = "SEG1"
  rpl      = "extcommunity-set seg-nh SEG1\n  10.1.1.1,\n  192.168.1.1\nend-set\n"
}

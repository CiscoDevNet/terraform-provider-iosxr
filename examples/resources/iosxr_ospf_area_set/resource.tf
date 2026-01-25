resource "iosxr_ospf_area_set" "example" {
  set_name = "OSPF1"
  rpl      = "ospf-area-set OSPF1\n  65536,\n  192.168.1.1\nend-set\n"
}

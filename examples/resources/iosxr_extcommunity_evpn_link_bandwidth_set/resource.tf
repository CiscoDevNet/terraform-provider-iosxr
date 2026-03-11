resource "iosxr_extcommunity_evpn_link_bandwidth_set" "example" {
  set_name = "EVPN1"
  rpl = "extcommunity-set evpn-link-bandwidth EVPN1\n  0:65001,\n  1:65001\nend-set\n"
}

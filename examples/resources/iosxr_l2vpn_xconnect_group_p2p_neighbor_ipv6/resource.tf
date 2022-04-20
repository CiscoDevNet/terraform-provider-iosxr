resource "iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6" "example" {
  group_name        = "P2P"
  p2p_xconnect_name = "XC"
  address           = "2001::2"
  pw_id             = 2
  pw_class          = "PW_CLASS_1"
}

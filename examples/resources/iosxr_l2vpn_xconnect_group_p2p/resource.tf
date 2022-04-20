resource "iosxr_l2vpn_xconnect_group_p2p" "example" {
  group_name        = "P2P"
  p2p_xconnect_name = "XC"
  description       = "My P2P Description"
}

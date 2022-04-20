data "iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv4" "example" {
  group_name        = "P2P"
  p2p_xconnect_name = "XC"
  address           = "2.3.4.5"
  pw_id             = 1
}

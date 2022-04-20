resource "iosxr_l2vpn_xconnect_group_p2p_interface" "example" {
  group_name        = "P2P"
  p2p_xconnect_name = "XC"
  interface_name    = "GigabitEthernet0/0/0/2"
}

resource "iosxr_l2vpn_xconnect_group_p2p" "example" {
  group_name        = "P2P"
  p2p_xconnect_name = "XC"
  description       = "My P2P Description"
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/2"
    }
  ]
  ipv4_neighbors = [
    {
      address  = "2.3.4.5"
      pw_id    = 1
      pw_class = "PW_CLASS_1"
    }
  ]
}

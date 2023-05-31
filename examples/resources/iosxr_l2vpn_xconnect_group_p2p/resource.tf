resource "iosxr_l2vpn_xconnect_group_p2p" "example" {
  group_name        = "P2P"
  p2p_xconnect_name = "XC"
  description       = "My P2P Description"
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/2"
    }
  ]
  neighbor_evpn_evi_segment_routing_services = [
    {
      vpn_id                       = 4600
      service_id                   = 600
      segment_routing_srv6_locator = "LOC11"
    }
  ]
}

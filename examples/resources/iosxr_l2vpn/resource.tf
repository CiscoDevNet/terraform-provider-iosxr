resource "iosxr_l2vpn" "example" {
  description                     = "My L2VPN Description"
  router_id                       = "1.2.3.4"
  load_balancing_flow_src_dst_mac = false
  load_balancing_flow_src_dst_ip  = true
  xconnect_groups = [
    {
      group_name = "P2P"
    }
  ]
}

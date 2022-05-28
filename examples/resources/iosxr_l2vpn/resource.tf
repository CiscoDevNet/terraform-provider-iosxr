resource "iosxr_l2vpn" "example" {
  description = "My L2VPN Description"
  router_id   = "1.2.3.4"
  xconnect_groups = [
    {
      group_name = "P2P"
    }
  ]
}

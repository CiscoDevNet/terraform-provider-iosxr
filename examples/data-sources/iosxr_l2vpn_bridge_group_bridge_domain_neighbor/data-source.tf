data "iosxr_l2vpn_bridge_group_bridge_domain_neighbor" "example" {
  address            = "10.1.1.3"
  bridge_domain_name = "BD123"
  bridge_group_name  = "BG123"
  pw_id              = 1000
}

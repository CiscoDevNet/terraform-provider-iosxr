data "iosxr_l2vpn_bridge_group_bridge_domain_neighbor" "example" {
  bridge_group_name = "BG123"
  bridge_domain_name = "BD123"
  address = "10.1.1.3"
  pw_id = 1000
}

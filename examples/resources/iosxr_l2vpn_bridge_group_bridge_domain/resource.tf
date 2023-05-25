resource "iosxr_l2vpn_bridge_group_bridge_domain" "example" {
  bridge_group_name  = "BG123"
  bridge_domain_name = "BD123"
  evis = [
    {
      vpn_id = 1234
    }
  ]
  vnis = [
    {
      vni_id = 1234
    }
  ]
}

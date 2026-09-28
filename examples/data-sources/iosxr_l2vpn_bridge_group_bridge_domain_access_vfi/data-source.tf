data "iosxr_l2vpn_bridge_group_bridge_domain_access_vfi" "example" {
  access_vfi_name    = "ACCESS_VFI1"
  bridge_domain_name = "BD123"
  bridge_group_name  = "BG123"
}

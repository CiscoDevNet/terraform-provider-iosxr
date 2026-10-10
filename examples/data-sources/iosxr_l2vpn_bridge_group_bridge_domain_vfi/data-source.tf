data "iosxr_l2vpn_bridge_group_bridge_domain_vfi" "example" {
  bridge_domain_name = "BD123"
  bridge_group_name  = "BG123"
  vfi_name           = "VFI1"
}

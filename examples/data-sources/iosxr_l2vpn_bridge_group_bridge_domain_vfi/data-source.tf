data "iosxr_l2vpn_bridge_group_bridge_domain_vfi" "example" {
  bridge_group_name  = "BG123"
  bridge_domain_name = "BD123"
  vfi_name           = "VFI1"
}

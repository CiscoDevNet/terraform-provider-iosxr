resource "iosxr_l2vpn_bridge_group_bridge_domain_access_vfi" "example" {
  bridge_group_name  = "BG123"
  bridge_domain_name = "BD123"
  access_vfi_name    = "ACCESS_VFI1"
  shutdown           = true
  neighbors = [
    {
      address = "10.1.1.1"
      pw_id   = 1000
      static_mac_addresses = [
        {
          mac_address = "aa:bb:cc:dd:ee:01"
        }
      ]
      pw_class = "PW_CLASS_1"
    }
  ]
}

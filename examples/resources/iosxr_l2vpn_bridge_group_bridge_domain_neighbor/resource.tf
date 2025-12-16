resource "iosxr_l2vpn_bridge_group_bridge_domain_neighbor" "example" {
  bridge_group_name  = "BG123"
  bridge_domain_name = "BD123"
  address            = "10.1.1.3"
  pw_id              = 1000
  static_mac_addresses = [
    {
      mac_address = "aa:bb:cc:dd:ee:05"
    }
  ]
  pw_class                    = "PW_CLASS_1"
  split_horizon_group         = true
  mac_aging_time              = 300
  mac_aging_type_inactivity   = true
  mac_learning                = true
  mac_port_down_flush_disable = true
  backup_neighbors = [
    {
      address  = "10.1.1.5"
      pw_id    = 1005
      pw_class = "PW_CLASS_2"
    }
  ]
}

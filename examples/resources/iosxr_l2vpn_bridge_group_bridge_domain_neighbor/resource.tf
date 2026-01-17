resource "iosxr_l2vpn_bridge_group_bridge_domain_neighbor" "example" {
  bridge_group_name  = "BG123"
  bridge_domain_name = "BD123"
  address            = "10.1.1.3"
  pw_id              = 1000
  flooding_disable   = true
  static_mac_addresses = [
    {
      mac_address = "aa:bb:cc:dd:ee:05"
    }
  ]
  pw_class                             = "PW_CLASS_1"
  split_horizon_group                  = true
  mac_aging_time                       = 300
  mac_aging_type_inactivity            = true
  mac_learning                         = true
  mac_limit_maximum                    = 1000
  mac_limit_action_no_flood            = true
  mac_port_down_flush_disable          = true
  mac_secure                           = true
  mac_secure_logging_disable           = true
  mac_secure_action_shutdown           = true
  mac_secure_shutdown_recovery_timeout = 300
  dhcp_ipv4_snooping_profile           = "DHCP_PROFILE"
  igmp_snooping_profile                = "IGMP_PROFILE"
  mld_snooping_profile                 = "MLD_PROFILE"
  backup_neighbors = [
    {
      address  = "10.1.1.5"
      pw_id    = 1005
      pw_class = "PW_CLASS_2"
    }
  ]
}

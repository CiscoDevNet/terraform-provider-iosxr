resource "iosxr_l2vpn_bridge_group_bridge_domain" "example" {
  bridge_domain_name                                = "BD123"
  bridge_group_name                                 = "BG123"
  coupled_mode                                      = true
  description                                       = "Bridge domain description"
  dynamic_arp_inspection                            = true
  dynamic_arp_inspection_address_validation_dst_mac = true
  dynamic_arp_inspection_address_validation_ipv4    = true
  dynamic_arp_inspection_address_validation_src_mac = true
  dynamic_arp_inspection_logging                    = true
  efp_visibility                                    = true
  etree                                             = true
  etree_leaf                                        = true
  evis = [
    {
      vpn_id = 1234
    }
  ]
  flooding_disable      = true
  igmp_snooping_profile = "PROFILE_1"
  interfaces = [
    {
      dynamic_arp_inspection_address_validation_dst_mac_disable = true
      dynamic_arp_inspection_address_validation_ipv4_disable    = true
      dynamic_arp_inspection_address_validation_src_mac_disable = true
      dynamic_arp_inspection_logging_disable                    = true
      flooding_disable                                          = true
      igmp_snooping_profile                                     = "PROFILE_1"
      interface_name                                            = "Bundle-Ether11.1234"
      ip_source_guard                                           = true
      ip_source_guard_logging                                   = true
      mac_aging_time                                            = 300
      mac_aging_type_inactivity                                 = true
      mac_learning                                              = true
      mac_limit_action_no_flood                                 = true
      mac_limit_maximum                                         = 1000
      mac_limit_notification_both                               = true
      mac_port_down_flush_disable                               = true
      mac_secure                                                = true
      mac_secure_action_shutdown                                = true
      mac_secure_logging_disable                                = true
      mac_secure_shutdown_recovery_timeout                      = 300
      mld_snooping_profile                                      = "PROFILE_2"
      split_horizon_group                                       = true
      static_mac_addresses = [
        {
          mac_address = "aa:bb:cc:dd:ee:ff"
        }
      ]
    }
  ]
  ip_source_guard                      = true
  ip_source_guard_logging              = true
  mac_aging_time                       = 300
  mac_aging_type_absolute              = true
  mac_learning_disable                 = true
  mac_limit_action_shutdown            = true
  mac_limit_maximum                    = 1000
  mac_limit_notification_both          = true
  mac_port_down_flush_disable          = true
  mac_secure                           = true
  mac_secure_action_shutdown           = true
  mac_secure_logging                   = true
  mac_secure_shutdown_recovery_timeout = 60
  mac_secure_threshold                 = true
  mac_withdraw_access_pw_disable       = true
  mac_withdraw_disable                 = true
  mac_withdraw_relay                   = true
  mac_withdraw_state_down              = true
  member_vnis_vni = [
    {
      static_mac_addresses = [
        {
          mac_address = "aa:bb:cc:dd:ee:04"
          next_hop    = "10.1.1.3"
        }
      ]
      vni_id = 1234
    }
  ]
  mld_snooping_profile       = "PROFILE_2"
  mtu                        = 1500
  multicast_source_ipv4_ipv6 = true
  neighbors_evpn_evi = [
    {
      target = 400
      vpn_id = 300
    }
  ]
  shutdown                        = false
  transport_mode_vlan_passthrough = true
}

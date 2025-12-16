resource "iosxr_l2vpn_bridge_group_bridge_domain" "example" {
  bridge_group_name  = "BG123"
  bridge_domain_name = "BD123"
  mtu                = 1500
  description        = "Bridge domain description"
  evis = [
    {
      vpn_id = 1234
    }
  ]
  coupled_mode                    = true
  transport_mode_vlan_passthrough = true
  flooding_disable                = true
  mld_snooping_profile            = "PROFILE_2"
  interfaces = [
    {
      interface_name              = "Bundle-Ether11.1234"
      mac_aging_time              = 300
      mac_aging_type_inactivity   = true
      mac_learning                = true
      mac_port_down_flush_disable = true
      mld_snooping_profile        = "PROFILE_2"
      split_horizon_group         = true
      static_mac_addresses = [
        {
          mac_address = "aa:bb:cc:dd:ee:ff"
        }
      ]
    }
  ]
  shutdown                       = false
  mac_aging_time                 = 300
  mac_aging_type_absolute        = true
  mac_learning_disable           = true
  mac_withdraw_disable           = true
  mac_withdraw_access_pw_disable = true
  mac_withdraw_relay             = true
  mac_withdraw_state_down        = true
  mac_limit_action_shutdown      = true
  mac_limit_notification_both    = true
  mac_port_down_flush_disable    = true
  neighbors_evpn_evi = [
    {
      vpn_id = 300
      target = 400
    }
  ]
  etree      = true
  etree_leaf = true
  member_vnis_vni = [
    {
      vni_id = 1234
      static_mac_addresses = [
        {
          mac_address = "aa:bb:cc:dd:ee:04"
          next_hop    = "10.1.1.3"
        }
      ]
    }
  ]
}

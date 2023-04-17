resource "iosxr_snmp_server" "example" {
  traps_rf                          = true
  traps_bfd                         = true
  traps_ntp                         = true
  traps_ethernet_oam_events         = true
  traps_copy_complete               = true
  traps_snmp_linkup                 = true
  traps_snmp_linkdown               = true
  traps_power                       = true
  traps_config                      = true
  traps_entity                      = true
  traps_system                      = true
  traps_bridgemib                   = true
  traps_entity_state_operstatus     = true
  traps_entity_redundancy_all       = true
  trap_source_both                  = "Loopback10"
  traps_l2vpn_all                   = true
  traps_l2vpn_vc_up                 = true
  traps_l2vpn_vc_down               = true
  traps_sensor                      = true
  traps_fru_ctrl                    = true
  traps_isis_authentication_failure = "enable"
  traps_bgp_cbgp2_updown            = true
  traps_bgp_bgp4_mib_updown         = true
  snmp_server_user = [
    {
      user_name                  = "USER1"
      group_name                 = "GROUP1"
      v3_auth_md5_encryption_aes = "073C05626E2A4841141D"
    }
  ]
  snmp_server_group = [
    {
      group_name = "GROUP12"
      v3_priv    = true
      v3_read    = "READ-STRING"
      v3_write   = "WRITE-STRING"
      v3_context = "CONTEXT-STRING"
      v3_notify  = "NOTIFY-STRING"
      v3_ipv4    = "IPV4-STRING"
      v3_ipv6    = "IPV6-STRING"
    }
  ]
}

resource "iosxr_snmp_server" "example" {
  rf                          = true
  bfd                         = true
  ntp                         = true
  ethernet_oam_events         = true
  copy_complete               = true
  traps_snmp_linkup           = true
  traps_snmp_linkdown         = true
  power                       = true
  config                      = true
  entity                      = true
  system                      = true
  bridgemib                   = true
  entity_state_operstatus     = true
  entity_redundancy_all       = true
  trap_source_both            = "Loopback10"
  l2vpn_all                   = true
  l2vpn_vc_up                 = true
  l2vpn_vc_down               = true
  sensor                      = true
  fru_ctrl                    = true
  isis_authentication_failure = "enable"
  bgp_cbgp2_updown            = true
  bgp_bgp4_mib_updown         = true
  users = [
    {
      user_name                  = "USER1"
      group_name                 = "GROUP1"
      v3_auth_md5_encryption_aes = "073C05626E2A4841141D"
    }
  ]
  groups = [
    {
      group_name = "GROUP12"
      v3_priv    = true
      v3_read    = "VIEW1"
      v3_write   = "VIEW2"
      v3_context = "CONTEXT1"
      v3_notify  = "VIEW3"
      v3_ipv4    = "ACL1"
      v3_ipv6    = "ACL2"
    }
  ]
}

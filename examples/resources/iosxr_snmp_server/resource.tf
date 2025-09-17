resource "iosxr_snmp_server" "example" {
  location                                  = "My location"
  contact                                   = "My contact"
  traps_rf                                  = true
  traps_bfd                                 = true
  traps_ntp                                 = true
  traps_ethernet_oam_events                 = true
  traps_copy_complete                       = true
  traps_snmp_linkup                         = true
  traps_snmp_linkdown                       = true
  traps_power                               = true
  traps_config                              = true
  traps_entity                              = true
  traps_system                              = true
  traps_bridgemib                           = true
  traps_entity_state_operstatus             = true
  traps_entity_redundancy_all               = true
  trap_source_both                          = "Loopback10"
  traps_l2vpn_all                           = true
  traps_l2vpn_vc_up                         = true
  traps_l2vpn_vc_down                       = true
  traps_sensor                              = true
  traps_fru_ctrl                            = true
  traps_isis_database_overload              = true
  traps_isis_manual_address_drops           = true
  traps_isis_corrupted_lsp_detected         = true
  traps_isis_attempt_to_exceed_max_sequence = true
  traps_isis_id_len_mismatch                = true
  traps_isis_max_area_addresses_mismatch    = true
  traps_isis_own_lsp_purge                  = true
  traps_isis_sequence_number_skip           = true
  traps_isis_authentication_type_failure    = true
  traps_isis_authentication_failure         = true
  traps_isis_version_skew                   = true
  traps_isis_area_mismatch                  = true
  traps_isis_rejected_adjacency             = true
  traps_isis_lsp_too_large_to_propagate     = true
  traps_isis_orig_lsp_buff_size_mismatch    = true
  traps_isis_protocols_supported_mismatch   = true
  traps_isis_adjacency_change               = true
  traps_isis_lsp_error_detected             = true
  traps_bgp_cbgp_two_enable                 = true
  traps_bgp_enable_cisco_bgp4_mib           = true
  users = [
    {
      user_name                  = "USER1"
      group_name                 = "GROUP1"
      v3_auth_md5_encryption_aes = "073C05626E2A4841141D"
      v3_ipv4                    = "ACL1"
      v3_systemowner             = true
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
  communities = [
    {
      community   = "COMMUNITY1"
      view        = "VIEW1"
      ro          = true
      rw          = false
      sdrowner    = false
      systemowner = true
      ipv4        = "ACL1"
      ipv6        = "ACL2"
    }
  ]
}

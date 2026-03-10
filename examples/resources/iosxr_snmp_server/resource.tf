resource "iosxr_snmp_server" "example" {
  location               = "My location"
  contact                = "My contact"
  chassis_id             = "Chassis1"
  packetsize             = 1024
  trap_timeout           = 10
  queue_length           = 100
  throttle_time          = 100
  overload_control       = 10
  overload_throttle_rate = 20
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
  traps_snmp_authentication                              = true
  traps_snmp_coldstart                                   = true
  traps_snmp_warmstart                                   = true
  traps_snmp_linkup                                      = true
  traps_snmp_linkdown                                    = true
  traps_snmp_all                                         = true
  traps_l2vpn_all                                        = true
  traps_l2vpn_vc_up                                      = true
  traps_l2vpn_vc_down                                    = true
  traps_l2vpn_cisco                                      = true
  traps_vpls_all                                         = true
  traps_vpls_status                                      = true
  traps_vpls_full_raise                                  = true
  traps_vpls_full_clear                                  = true
  traps_bfd                                              = true
  traps_config                                           = true
  traps_cfm                                              = true
  traps_ethernet_oam_events                              = true
  traps_rf                                               = true
  traps_sensor                                           = true
  traps_mpls_l3vpn_all                                   = true
  traps_mpls_l3vpn_vrf_up                                = true
  traps_mpls_l3vpn_vrf_down                              = true
  traps_mpls_l3vpn_mid_threshold_exceeded                = true
  traps_mpls_l3vpn_max_threshold_exceeded                = true
  traps_mpls_l3vpn_max_threshold_cleared                 = true
  traps_mpls_l3vpn_max_threshold_reissue_notif_time      = 100
  traps_mpls_traffic_eng_cisco                           = true
  traps_mpls_traffic_eng_cisco_ext_bringup_fail          = true
  traps_mpls_traffic_eng_cisco_ext_insuff_bw             = true
  traps_mpls_traffic_eng_cisco_ext_preempt               = true
  traps_mpls_traffic_eng_cisco_ext_reroute_pending       = true
  traps_mpls_traffic_eng_cisco_ext_reroute_pending_clear = true
  traps_mpls_traffic_eng_down                            = true
  traps_mpls_traffic_eng_p2mp_down                       = true
  traps_mpls_traffic_eng_p2mp_up                         = true
  traps_mpls_traffic_eng_reoptimize                      = true
  traps_mpls_traffic_eng_reroute                         = true
  traps_mpls_traffic_eng_up                              = true
  traps_ntp                                              = true
  traps_bgp_cbgp_two_enable                              = true
  traps_bgp_enable_cisco_bgp4_mib                        = true
  traps_hsrp                                             = true
  traps_isis_database_overload                           = true
  traps_isis_manual_address_drops                        = true
  traps_isis_corrupted_lsp_detected                      = true
  traps_isis_attempt_to_exceed_max_sequence              = true
  traps_isis_id_len_mismatch                             = true
  traps_isis_max_area_addresses_mismatch                 = true
  traps_isis_own_lsp_purge                               = true
  traps_isis_sequence_number_skip                        = true
  traps_isis_authentication_type_failure                 = true
  traps_isis_authentication_failure                      = true
  traps_isis_version_skew                                = true
  traps_isis_area_mismatch                               = true
  traps_isis_rejected_adjacency                          = true
  traps_isis_lsp_too_large_to_propagate                  = true
  traps_isis_orig_lsp_buff_size_mismatch                 = true
  traps_isis_protocols_supported_mismatch                = true
  traps_isis_adjacency_change                            = true
  traps_isis_lsp_error_detected                          = true
  traps_vrrp_events                                      = true
  traps_alarm                                            = true
  traps_bridgemib                                        = true
  traps_copy_complete                                    = true
  traps_entity                                           = true
  traps_cisco_entity_ext                                 = true
  traps_entity_redundancy_all                            = true
  traps_entity_redundancy_switchover                     = true
  traps_entity_redundancy_status                         = true
  traps_entity_state_switchover                          = true
  traps_entity_state_operstatus                          = true
  traps_flash_insertion                                  = true
  traps_flash_removal                                    = true
  traps_fru_ctrl                                         = true
  traps_ipsla                                            = true
  traps_mpls_ldp_down                                    = true
  traps_mpls_ldp_up                                      = true
  traps_mpls_ldp_threshold                               = true
  traps_pim_neighbor_change                              = true
  traps_pim_interface_state_change                       = true
  traps_pim_invalid_message_received                     = true
  traps_pim_rp_mapping_change                            = true
  traps_power                                            = true
  traps_syslog                                           = true
  traps_system                                           = true
  hosts = [
    {
      address = "11.11.11.11"
      traps_unencrypted_strings = [
        {
          community_string          = "COMMUNITY1"
          udp_port                  = 1100
          version_v3_security_level = "auth"
        }
      ]
      informs_unencrypted_strings = [
        {
          community_string          = "COMMUNITY2"
          udp_port                  = 1100
          version_v3_security_level = "auth"
        }
      ]
    }
  ]
  views = [
    {
      view_name = "VIEW1"
      mib_view_families = [
        {
          name     = "1.3.6.1.2.1.1"
          included = true
        }
      ]
    }
  ]
  trap_source                     = "Loopback10"
  trap_source_ipv4                = "Loopback0"
  trap_source_ipv6                = "Loopback1"
  trap_source_port                = 1200
  trap_throttle_time              = 100
  trap_authentication_vrf_disable = true
  trap_delay_timer                = 30
  ipv4_dscp                       = "ef"
  ipv6_dscp                       = "ef"
  drop_unknown_user               = true
  drop_report_acl_ipv4            = "ACL1"
  drop_report_acl_ipv6            = "ACL1"
  groups = [
    {
      group_name = "GROUP12"
      v3_priv    = true
      v3_read    = "VIEW1"
      v3_write   = "VIEW2"
      v3_context = "CONTEXT1"
      v3_notify  = "VIEW3"
      v3_ipv4    = "ACL1"
      v3_ipv6    = "ACL1"
    }
  ]
  engine_id_local = "80000009030000C0A80101"
  engine_id_remotes = [
    {
      address   = "11.11.11.11"
      engine_id = "80000009030000C0A80101"
      udp_port  = 1100
    }
  ]
  users = [
    {
      user_name                  = "USER1"
      group_name                 = "GROUP1"
      v3                         = true
      v3_auth_md5_encryption_aes = "073C05626E2A4841141D"
      v3_ipv4                    = "ACL1"
      v3_ipv6                    = "ACL1"
      v3_systemowner             = true
    }
  ]
  oid_poll_stats                   = true
  timeouts_subagent                = 20
  timeouts_duplicate               = 10
  timeouts_in_qdrop                = 20
  timeouts_threshold               = 10
  timeouts_pdu_stats               = 10
  logging_threshold_oid_processing = 10
  logging_threshold_pdu_processing = 10
  inform_retries                   = 10
  inform_timeout                   = 10
  inform_pending                   = 10
}

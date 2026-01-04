resource "iosxr_evpn" "example" {
  bgp_rd_ipv4_address                               = "192.168.1.1"
  bgp_rd_ipv4_address_index                         = 100
  timers_recovery                                   = 120
  timers_peering                                    = 60
  timers_carving                                    = 5
  timers_ac_debounce                                = 2000
  timers_backup_replacement_delay                   = 3000
  timers_mac_postpone                               = 240
  load_balancing_flow_label_static                  = true
  source_interface                                  = "Loopback0"
  cost_out                                          = true
  startup_cost_in                                   = 60
  staggered_bringup_timer                           = 3000
  logging_df_election                               = true
  ethernet_segment_type_one_auto_generation_disable = true
  groups = [
    {
      group_name = 10
      core_interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/2"
        }
      ]
    }
  ]
  srv6 = true
  srv6_locators = [
    {
      locator_name                        = "LOC1"
      usid_allocation_wide_local_id_block = true
    }
  ]
  srv6_usid_allocation_wide_local_id_block                  = true
  ignore_mtu_mismatch                                       = true
  transmit_mtu_zero                                         = true
  host_ipv4_duplicate_detection_move_count                  = 10
  host_ipv4_duplicate_detection_move_interval               = 360
  host_ipv4_duplicate_detection_freeze_time                 = 120
  host_ipv4_duplicate_detection_retry_count                 = "5"
  host_ipv4_duplicate_detection_reset_freeze_count_interval = 48
  host_ipv6_duplicate_detection_move_count                  = 10
  host_ipv6_duplicate_detection_move_interval               = 360
  host_ipv6_duplicate_detection_freeze_time                 = 120
  host_ipv6_duplicate_detection_retry_count                 = "5"
  host_ipv6_duplicate_detection_reset_freeze_count_interval = 48
  virtual_neighbors = [
    {
      address                                            = "192.168.1.1"
      pw_id                                              = 100
      timers_peering                                     = 60
      timers_recovery                                    = 120
      timers_carving                                     = 5
      timers_ac_debounce                                 = 2000
      ethernet_segment_esi_zero                          = "01.01.01.01.01.01.01.01.01"
      ethernet_segment_service_carving_manual_primary    = "100-101,103"
      ethernet_segment_service_carving_manual_secondary  = "200-201,203"
      ethernet_segment_service_carving_multicast_hrw_s_g = true
      ethernet_segment_bgp_rt                            = "01:01:01:01:01:01"
    }
  ]
  virtual_vfis = [
    {
      vfi_name                                          = "VFI1"
      timers_peering                                    = 60
      timers_recovery                                   = 120
      timers_carving                                    = 5
      timers_ac_debounce                                = 2000
      ethernet_segment_esi_zero                         = "01.01.01.01.02.02.02.02.02"
      ethernet_segment_service_carving_manual_primary   = "100-101,103"
      ethernet_segment_service_carving_manual_secondary = "200-201,203"
      ethernet_segment_bgp_rt                           = "01:01:01:01:01:02"
    }
  ]
  virtual_access_evi_ethernet_segment_esi_zero = "01.01.01.01.01.01.01.01.03"
  virtual_access_evi_ethernet_segment_bgp_rt   = "01:01:01:01:01:03"
}

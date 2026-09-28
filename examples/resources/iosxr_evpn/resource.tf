resource "iosxr_evpn" "example" {
  bgp_rd_ipv4_address                               = "192.168.1.1"
  bgp_rd_ipv4_address_index                         = 100
  cost_out                                          = true
  ethernet_segment_type_one_auto_generation_disable = true
  groups = [
    {
      core_interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/2"
        }
      ]
      group_id = 10
    }
  ]
  host_ipv4_duplicate_detection_freeze_time                 = 120
  host_ipv4_duplicate_detection_move_count                  = 10
  host_ipv4_duplicate_detection_move_interval               = 360
  host_ipv4_duplicate_detection_reset_freeze_count_interval = 48
  host_ipv4_duplicate_detection_retry_count                 = "5"
  host_ipv6_duplicate_detection_freeze_time                 = 120
  host_ipv6_duplicate_detection_move_count                  = 10
  host_ipv6_duplicate_detection_move_interval               = 360
  host_ipv6_duplicate_detection_reset_freeze_count_interval = 48
  host_ipv6_duplicate_detection_retry_count                 = "5"
  ignore_mtu_mismatch                                       = true
  load_balancing_flow_label_static                          = true
  logging_df_election                                       = true
  source_interface                                          = "Loopback0"
  srv6                                                      = true
  srv6_locators = [
    {
      locator_name                        = "LOC1"
      usid_allocation_wide_local_id_block = true
    }
  ]
  srv6_usid_allocation_wide_local_id_block     = true
  staggered_bringup_timer                      = 3000
  startup_cost_in                              = 60
  timers_ac_debounce                           = 2000
  timers_backup_replacement_delay              = 3000
  timers_carving                               = 5
  timers_mac_postpone                          = 240
  timers_peering                               = 60
  timers_recovery                              = 120
  transmit_mtu_zero                            = true
  virtual_access_evi_ethernet_segment_bgp_rt   = "01:01:01:01:01:03"
  virtual_access_evi_ethernet_segment_esi_zero = "01.01.01.01.01.01.01.01.03"
  virtual_neighbors = [
    {
      address                                            = "192.168.1.1"
      ethernet_segment_bgp_rt                            = "01:01:01:01:01:01"
      ethernet_segment_esi_zero                          = "01.01.01.01.01.01.01.01.01"
      ethernet_segment_service_carving_manual_primary    = "100-101,103"
      ethernet_segment_service_carving_manual_secondary  = "200-201,203"
      ethernet_segment_service_carving_multicast_hrw_s_g = true
      pw_id                                              = 100
      timers_ac_debounce                                 = 2000
      timers_carving                                     = 5
      timers_peering                                     = 60
      timers_recovery                                    = 120
    }
  ]
  virtual_vfis = [
    {
      ethernet_segment_bgp_rt                           = "01:01:01:01:01:02"
      ethernet_segment_esi_zero                         = "01.01.01.01.02.02.02.02.02"
      ethernet_segment_service_carving_manual_primary   = "100-101,103"
      ethernet_segment_service_carving_manual_secondary = "200-201,203"
      timers_ac_debounce                                = 2000
      timers_carving                                    = 5
      timers_peering                                    = 60
      timers_recovery                                   = 120
      vfi_name                                          = "VFI1"
    }
  ]
}

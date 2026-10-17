resource "iosxr_rsvp" "example" {
  authentication_key_chain                          = "KEY1"
  authentication_life_time                          = 60
  authentication_retransmit                         = 120
  authentication_window_size                        = 32
  bandwidth_mam_percentage_max_reservable_bandwidth = 10000
  bandwidth_mam_percentage_max_reservable_bc0       = 2000
  bandwidth_mam_percentage_max_reservable_bc1       = 2000
  bandwidth_rdm_percentage_max_reservable_bc0       = 10000
  bandwidth_rdm_percentage_max_reservable_bc1       = 2000
  interfaces = [
    {
      authentication_key_chain                                     = "KEY1"
      authentication_life_time                                     = 60
      authentication_window_size                                   = 32
      bandwidth_mam_percentage_bc0_total                           = 80
      bandwidth_mam_percentage_bc1_total                           = 20
      bandwidth_mam_percentage_flow                                = 20
      bandwidth_mam_percentage_total                               = 80
      interface_name                                               = "GigabitEthernet0/0/0/1"
      signalling_dscp                                              = 48
      signalling_hello_graceful_restart_interface_based            = true
      signalling_rate_limit_enable                                 = true
      signalling_rate_limit_interval                               = 1000
      signalling_rate_limit_rate                                   = 200
      signalling_refresh_interval                                  = 120
      signalling_refresh_missed                                    = 2
      signalling_refresh_oob_interval                              = 200
      signalling_refresh_oob_missed                                = 2
      signalling_refresh_reduction_bundle_max_size                 = 1000
      signalling_refresh_reduction_reliable_ack_hold_time          = 1000
      signalling_refresh_reduction_reliable_ack_max_size           = 2000
      signalling_refresh_reduction_reliable_retransmit_queue_depth = 100
      signalling_refresh_reduction_reliable_retransmit_time        = 1000
      signalling_refresh_reduction_reliable_summary_refresh        = true
      signalling_refresh_reduction_summary_max_size                = 100
    }
  ]
  latency_threshold                = 120
  logging_events_issu              = true
  logging_events_nsr               = true
  ltrace_buffer_multiplier         = 3
  ltrace_buffer_multiplier_common  = true
  ltrace_buffer_multiplier_dbg_err = true
  ltrace_buffer_multiplier_intf    = true
  ltrace_buffer_multiplier_rare    = true
  ltrace_buffer_multiplier_sig     = true
  ltrace_buffer_multiplier_sig_err = true
  ltrace_buffer_multiplier_sync    = true
  neighbors = [
    {
      address                    = "192.168.1.1"
      authentication_key_chain   = "KEY1"
      authentication_life_time   = 60
      authentication_window_size = 32
    }
  ]
  signalling_checksum_disable                          = true
  signalling_event_per_pulse                           = 60
  signalling_graceful_restart                          = true
  signalling_graceful_restart_lsp_ctype                = true
  signalling_graceful_restart_recovery_time            = 300
  signalling_graceful_restart_restart_time             = 220
  signalling_hello_graceful_restart_refresh_interval   = 3000
  signalling_hello_graceful_restart_refresh_misses     = 5
  signalling_message_bundle_disable                    = true
  signalling_nodeid_subobject_disable                  = true
  signalling_oob_vrf                                   = "VRF1"
  signalling_patherr_state_removal_disable             = true
  signalling_prefix_filtering_acl                      = "ACL1"
  signalling_prefix_filtering_default_deny_action_drop = true
}

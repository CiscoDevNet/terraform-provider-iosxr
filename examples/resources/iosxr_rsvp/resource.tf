resource "iosxr_rsvp" "example" {
  signalling_graceful_restart                          = true
  signalling_graceful_restart_lsp_ctype                = true
  signalling_graceful_restart_recovery_time            = 300
  signalling_graceful_restart_restart_time             = 220
  signalling_hello_graceful_restart_refresh_interval   = 3000
  signalling_hello_graceful_restart_refresh_misses     = 5
  signalling_event_per_pulse                           = 60
  signalling_prefix_filtering_acl                      = "ACL1"
  signalling_prefix_filtering_default_deny_action_drop = true
  signalling_message_bundle_disable                    = true
  signalling_nodeid_subobject_disable                  = true
  signalling_patherr_state_removal_disable             = true
  signalling_checksum_disable                          = true
  signalling_oob_vrf                                   = "VRF1"
  authentication_key_chain                             = "KEY1"
  authentication_window_size                           = 32
  authentication_life_time                             = 60
  authentication_retransmit                            = 120
  neighbors = [
    {
      address                    = "192.168.1.1"
      authentication_key_chain   = "KEY1"
      authentication_window_size = 32
      authentication_life_time   = 60
    }
  ]
  bandwidth_mam_percentage_max_reservable_bandwidth = 10000
  bandwidth_mam_percentage_max_reservable_bc0       = 2000
  bandwidth_mam_percentage_max_reservable_bc1       = 2000
  bandwidth_rdm_percentage_max_reservable_bc0       = 10000
  bandwidth_rdm_percentage_max_reservable_bc1       = 2000
  latency_threshold                                 = 120
  logging_events_nsr                                = true
  logging_events_issu                               = true
  ltrace_buffer_multiplier                          = 3
  ltrace_buffer_multiplier_rare                     = true
  ltrace_buffer_multiplier_common                   = true
  ltrace_buffer_multiplier_sig                      = true
  ltrace_buffer_multiplier_sig_err                  = true
  ltrace_buffer_multiplier_intf                     = true
  ltrace_buffer_multiplier_dbg_err                  = true
  ltrace_buffer_multiplier_sync                     = true
}

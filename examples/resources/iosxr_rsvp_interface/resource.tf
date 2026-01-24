resource "iosxr_rsvp_interface" "example" {
  interface_name                                               = "GigabitEthernet0/0/0/1"
  bandwidth_mam_percentage_total                               = 80
  bandwidth_mam_percentage_flow                                = 20
  bandwidth_mam_percentage_bc0_total                           = 80
  bandwidth_mam_percentage_bc1_total                           = 20
  signalling_dscp                                              = 48
  signalling_rate_limit_enable                                 = true
  signalling_rate_limit_rate                                   = 200
  signalling_rate_limit_interval                               = 1000
  signalling_refresh_interval                                  = 120
  signalling_refresh_missed                                    = 2
  signalling_refresh_oob_interval                              = 200
  signalling_refresh_oob_missed                                = 2
  signalling_refresh_reduction_reliable_ack_hold_time          = 1000
  signalling_refresh_reduction_reliable_ack_max_size           = 2000
  signalling_refresh_reduction_reliable_retransmit_time        = 1000
  signalling_refresh_reduction_reliable_retransmit_queue_depth = 100
  signalling_refresh_reduction_reliable_summary_refresh        = true
  signalling_refresh_reduction_summary_max_size                = 100
  signalling_refresh_reduction_bundle_max_size                 = 1000
  signalling_hello_graceful_restart_interface_based            = true
  authentication_key_chain                                     = "KEY1"
  authentication_window_size                                   = 32
  authentication_life_time                                     = 60
}

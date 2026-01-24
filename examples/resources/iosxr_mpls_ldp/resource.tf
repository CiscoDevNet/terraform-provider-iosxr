resource "iosxr_mpls_ldp" "example" {
  router_id                                                = "1.2.3.4"
  graceful_restart                                         = true
  graceful_restart_reconnect_timeout                       = 120
  graceful_restart_forwarding_state_holdtime               = 180
  graceful_restart_helper_peer_maintain_on_local_reset_for = "ACL1"
  ltrace_buffer_multiplier                                 = 2
  default_vrf_implicit_ipv4_disable                        = true
  session_backoff_time_initial                             = 10
  session_backoff_time_maximum                             = 100
  session_holdtime                                         = 180
  session_downstream_on_demand_with                        = "ACL1"
  session_protection                                       = true
  session_protection_for_acl                               = "ACL1"
  session_protection_for_acl_duration                      = 120
  nsr                                                      = true
  entropy_label                                            = true
  entropy_label_add_el                                     = true
  signalling_dscp                                          = 48
  igp_sync_delay_on_session_up                             = 10
  igp_sync_delay_on_proc_restart                           = 100
  capabilities_sac                                         = true
  log_hello_adjacency                                      = true
  log_neighbor                                             = true
  log_nsr                                                  = true
  log_graceful_restart                                     = true
  log_session_protection                                   = true
  discovery_hello_holdtime                                 = 10
  discovery_hello_interval                                 = 3
  discovery_targeted_hello_holdtime                        = 10
  discovery_targeted_hello_interval                        = 3
  discovery_instance_tlv_disable                           = true
  discovery_ds_tlv_disable                                 = true
  discovery_rtr_id_arb_tlv_disable                         = true
  discovery_quick_start_disable                            = true
  neighbor_dual_stack_transport_connection_prefer_ipv4     = true
  neighbor_dual_stack_transport_connection_max_wait        = 30
  neighbor_dual_stack_tlv_compliance                       = true
  neighbors = [
    {
      neighbor_address   = "192.168.2.1"
      label_space_id     = 0
      password_encrypted = "060506324F41"
    }
  ]
}

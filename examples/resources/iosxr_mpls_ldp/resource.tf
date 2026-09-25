resource "iosxr_mpls_ldp" "example" {
  capabilities_sac = true
  default_vrf_implicit_ipv4_disable = true
  discovery_ds_tlv_disable = true
  discovery_hello_holdtime = 10
  discovery_hello_interval = 3
  discovery_instance_tlv_disable = true
  discovery_quick_start_disable = true
  discovery_rtr_id_arb_tlv_disable = true
  discovery_targeted_hello_holdtime = 10
  discovery_targeted_hello_interval = 3
  entropy_label = true
  entropy_label_add_el = true
  graceful_restart = true
  graceful_restart_forwarding_state_holdtime = 180
  graceful_restart_helper_peer_maintain_on_local_reset_for = "ACL1"
  graceful_restart_reconnect_timeout = 120
  igp_sync_delay_on_proc_restart = 100
  igp_sync_delay_on_session_up = 10
  log_graceful_restart = true
  log_hello_adjacency = true
  log_neighbor = true
  log_nsr = true
  log_session_protection = true
  ltrace_buffer_multiplier = 2
  neighbor_dual_stack_tlv_compliance = true
  neighbor_dual_stack_transport_connection_max_wait = 30
  neighbor_dual_stack_transport_connection_prefer_ipv4 = true
  neighbor_password_encrypted = "060506324F41"
  neighbors = [
    {
      label_space_id = 0
      neighbor_address = "192.168.2.1"
      password_encrypted = "060506324F41"
    }
  ]
  nsr = true
  router_id = "1.2.3.4"
  session_backoff_time_initial = 10
  session_backoff_time_maximum = 100
  session_downstream_on_demand_with = "ACL1"
  session_holdtime = 180
  session_protection = true
  session_protection_for_acl = "ACL1"
  session_protection_for_acl_duration = 120
  signalling_dscp = 48
}

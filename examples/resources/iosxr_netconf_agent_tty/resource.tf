resource "iosxr_netconf_agent_tty" "example" {
  session_timeout         = 30
  throttle_memory         = 300
  throttle_offload_memory = 0
  throttle_process_rate   = 5000
}

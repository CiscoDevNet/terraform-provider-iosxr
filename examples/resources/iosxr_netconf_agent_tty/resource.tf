resource "iosxr_netconf_agent_tty" "example" {
  throttle_process_rate = 5000
  throttle_memory = 300
  throttle_offload_memory = 0
  session_timeout = 30
}

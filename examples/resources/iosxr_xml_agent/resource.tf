resource "iosxr_xml_agent" "example" {
  enable                    = true
  tty_enable                = true
  tty_streaming_size        = 1000
  tty_iteration_size        = "off"
  tty_throttle_process_rate = 5000
  tty_throttle_memory       = 300
  tty_session_timeout       = 30
  ssl_enable                = true
  ssl_streaming_size        = 1000
  ssl_iteration_size        = "off"
  ssl_throttle_process_rate = 5000
  ssl_throttle_memory       = 300
  ssl_session_timeout       = 30
  ssl_vrfs = [
    {
      vrf_name         = "default"
      shutdown         = true
      ipv4_access_list = "ACL_IPV4"
    }
  ]
  ipv6_enable           = true
  ipv4_disable          = true
  streaming_size        = 1000
  iteration_size        = "off"
  throttle_process_rate = 5000
  throttle_memory       = 300
  session_timeout       = 30
  vrfs = [
    {
      vrf_name         = "VRF1"
      shutdown         = false
      ipv6_access_list = "ACL1"
      ipv4_access_list = "ACL2"
    }
  ]
}

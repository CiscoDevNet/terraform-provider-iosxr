resource "iosxr_xml_agent" "example" {
  enable                    = true
  ipv4_disable              = true
  ipv6_enable               = true
  iteration_size            = "off"
  session_timeout           = 30
  ssl_enable                = true
  ssl_iteration_size        = "off"
  ssl_session_timeout       = 30
  ssl_streaming_size        = 1000
  ssl_throttle_memory       = 300
  ssl_throttle_process_rate = 5000
  ssl_vrfs = [
    {
      ipv4_access_list = "ACL_IPV4"
      shutdown         = true
      vrf_name         = "default"
    }
  ]
  streaming_size            = 1000
  throttle_memory           = 300
  throttle_process_rate     = 5000
  tty_enable                = true
  tty_iteration_size        = "off"
  tty_session_timeout       = 30
  tty_streaming_size        = 1000
  tty_throttle_memory       = 300
  tty_throttle_process_rate = 5000
  vrfs = [
    {
      ipv4_access_list = "ACL2"
      ipv6_access_list = "ACL1"
      shutdown         = false
      vrf_name         = "VRF1"
    }
  ]
}

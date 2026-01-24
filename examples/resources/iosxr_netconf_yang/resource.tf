resource "iosxr_netconf_yang" "example" {
  agent                         = true
  ssh_enable                    = true
  with_defaults_support         = true
  rate_limit                    = 4096
  session_limit                 = 50
  session_idle_timeout          = 30
  session_absolute_timeout      = 1440
  netconf_v1_0                  = "1.0-only"
  netconf_v1_streaming_disabled = true
}

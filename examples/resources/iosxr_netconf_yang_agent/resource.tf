resource "iosxr_netconf_yang_agent" "example" {
  session_absolute_timeout = 1440
  session_idle_timeout     = 30
  session_limit            = 50
  with_defaults_support    = true
}

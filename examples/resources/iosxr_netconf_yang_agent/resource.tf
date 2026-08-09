resource "iosxr_netconf_yang_agent" "example" {
  with_defaults_support    = true
  session_limit            = 50
  session_idle_timeout     = 30
  session_absolute_timeout = 1440
}

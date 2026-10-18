resource "iosxr_line_console" "example" {
  absolute_timeout            = 3600
  access_class_egress         = "CONSOLE_ACL"
  access_class_ingress        = "CONSOLE_ACL"
  cli_whitespace_completion   = true
  disconnect_character        = "0x0a"
  escape_character            = "0x0a"
  exec_timeout_minutes        = 30
  exec_timeout_seconds        = 0
  length                      = 25
  pager                       = "none"
  secret_encrypted            = "$1$UgkY$I2SEocww.URG7gvDI7oz01"
  session_limit               = 15
  session_timeout             = 1440
  session_timeout_output      = true
  telnet_transparent          = true
  timeout_login_response      = 60
  timestamp_disable           = true
  transport_input_ssh         = true
  transport_output_ssh_telnet = true
  transport_preferred_ssh     = true
  users_group = [
    {
      group_name = "cisco-support"
    }
  ]
  width = 81
}

resource "iosxr_line_template" "example" {
  template_name               = "Template-1"
  autocommand                 = "show version"
  access_class_ingress        = "CONSOLE_ACL"
  access_class_egress         = "CONSOLE_ACL"
  disconnect_character        = "0x0a"
  escape_character            = "0x0a"
  session_timeout             = 1440
  session_timeout_output      = true
  transport_input_ssh         = true
  transport_output_ssh_telnet = true
  transport_preferred_ssh     = true
  session_limit               = 15
  cli_whitespace_completion   = true
  secret_encrypted            = "$1$UgkY$I2SEocww.URG7gvDI7oz01"
  timeout_login_response      = 60
  users_group = [
    {
      group_name = "cisco-support"
    }
  ]
  exec_timeout_minutes = 30
  exec_timeout_seconds = 0
  absolute_timeout     = 3600
  width                = 81
  length               = 25
  timestamp_disable    = true
  pager                = "none"
  telnet_transparent   = true
}

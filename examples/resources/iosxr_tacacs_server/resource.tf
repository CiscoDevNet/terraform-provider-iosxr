resource "iosxr_tacacs_server" "example" {
  hosts = [
    {
      address                        = "9.0.1.68"
      port                           = 49
      timeout                        = 10
      holddown_time                  = 300
      key_type_7                     = "0235347225301B204F4F0A0A"
      single_connection              = true
      single_connection_idle_timeout = 1000
    }
  ]
  key_type_7    = "0235347225301B204F4F0A0A"
  timeout       = 5
  holddown_time = 600
  ipv4_dscp     = "cs6"
  ipv6_dscp     = "cs6"
}

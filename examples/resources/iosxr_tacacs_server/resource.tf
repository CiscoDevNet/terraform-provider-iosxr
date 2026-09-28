resource "iosxr_tacacs_server" "example" {
  holddown_time = 600
  hosts = [
    {
      address                        = "9.0.1.68"
      holddown_time                  = 300
      key_type_7                     = "0235347225301B204F4F0A0A"
      port                           = 49
      single_connection              = true
      single_connection_idle_timeout = 1000
      timeout                        = 10
    }
  ]
  ipv4_dscp  = "cs6"
  ipv6_dscp  = "cs6"
  key_type_7 = "0235347225301B204F4F0A0A"
  timeout    = 5
}

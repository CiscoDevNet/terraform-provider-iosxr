resource "iosxr_http_client" "example" {
  connection_retry           = 5
  connection_timeout         = 30
  response_timeout           = 60
  secure_verify_host_disable = true
  secure_verify_peer_disable = true
  source_interface_ipv4      = "MgmtEth0/RP0/CPU0/0"
  ssl_version_tls13          = true
  tcp_window_scale           = 14
  version_default            = true
  vrf                        = "MGMT"
}

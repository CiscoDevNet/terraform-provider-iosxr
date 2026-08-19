resource "iosxr_http_client" "example" {
  vrf                        = "MGMT"
  secure_verify_peer_disable = true
  secure_verify_host_disable = true
  response_timeout           = 60
  connection_timeout         = 30
  connection_retry           = 5
  source_interface_ipv4      = "MgmtEth0/RP0/CPU0/0"
  version_default            = true
  tcp_window_scale           = 14
  ssl_version_tls13          = true
}

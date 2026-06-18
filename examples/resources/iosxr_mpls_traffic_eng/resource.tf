resource "iosxr_mpls_traffic_eng" "example" {
  traffic_eng = true
  # Supported from version 25.1
  disable = true
  # Supported from version 25.1
  reoptimize_reoptimization_period_in = 3600
  # Supported from version 25.1
  server_ipv4 = "192.0.2.1"
}

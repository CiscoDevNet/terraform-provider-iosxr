resource "iosxr_ssh" "example" {
  server_dscp          = 48
  server_logging       = true
  server_rate_limit    = 60
  server_session_limit = 10
  server_v2            = true
  server_vrfs = [
    {
      vrf_name = "VRF1"
    }
  ]
}

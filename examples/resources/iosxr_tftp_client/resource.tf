resource "iosxr_tftp_client" "example" {
  client_vrfs = [
    {
      dscp             = "default"
      retries          = 10
      source_interface = "Loopback0"
      timeout          = 30
      vrf_name         = "VRF1"
    }
  ]
}

resource "iosxr_tftp_client" "example" {
  client_vrfs = [
    {
      vrf_name = "VRF1"
      source_interface = "Loopback0"
      retries = 10
      timeout = 30
      dscp = "default"
    }
  ]
}

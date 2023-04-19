resource "iosxr_logging_source_interface" "example" {
  name = "Loopback0"
  vrfs = [
    {
      name = "VRF1"
    }
  ]
}

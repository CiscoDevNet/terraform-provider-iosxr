resource "iosxr_radius_source_interface" "example" {
  source_interface = "Loopback0"
  vrf              = "VRF1"
}

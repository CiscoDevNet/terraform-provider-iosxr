resource "iosxr_radius_source_interface" "example" {
  vrf = "VRF1"
  source_interface = "Loopback0"
}

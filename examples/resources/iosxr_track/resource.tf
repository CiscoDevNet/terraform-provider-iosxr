resource "iosxr_track" "example" {
  track_name = "TRACK1"
  delay_up = 10
  delay_down = 5
  route_address_prefix = "2001:db8::"
  route_address_prefix_length = 64
  route_vrf = "VRF1"
}

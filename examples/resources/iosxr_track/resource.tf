resource "iosxr_track" "example" {
  delay_down                  = 5
  delay_up                    = 10
  route_address_prefix        = "2001:db8::"
  route_address_prefix_length = 64
  route_vrf                   = "VRF1"
  track_name                  = "TRACK1"
}

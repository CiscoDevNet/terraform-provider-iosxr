resource "iosxr_track" "example" {
  track = [
    {
      track_name              = "TRACK1"
      delay_up                = 10
      delay_down              = 5
      line_protocol_interface = "GigabitEthernet0/0/0/1"
    }
  ]
}

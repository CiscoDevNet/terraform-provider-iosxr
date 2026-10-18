resource "iosxr_monitor_session" "example" {
  default_capture_disable = true
  monitor_sessions = [
    {
      destination_interface = "GigabitEthernet0/0/0/1"
      discard_class         = 1
      mirror_first          = 256
      session_name          = "SPAN1"
      traffic_class         = 5
      traffic_type          = "ethernet"
    }
  ]
  router_id = 1
}

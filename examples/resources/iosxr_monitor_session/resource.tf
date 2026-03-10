resource "iosxr_monitor_session" "example" {
  monitor_sessions = [
    {
      session_name          = "SPAN1"
      traffic_type          = "ethernet"
      destination_interface = "GigabitEthernet0/0/0/1"
      discard_class         = 1
      traffic_class         = 5
      mirror_first          = 256
    }
  ]
  router_id               = 1
  default_capture_disable = true
}

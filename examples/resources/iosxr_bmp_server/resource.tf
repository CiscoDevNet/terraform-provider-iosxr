resource "iosxr_bmp_server" "example" {
  all_route_monitorings = [
    {
      route_mon = "inbound-pre-policy"
      advertisement_interval = 60
      scan_time = 5
    }
  ]
  all_max_buffer_size = 15
  servers = [
    {
      number = 1
      shutdown = false
      host = "192.168.1.100"
      port = 5000
      initial_delay = 60
      flapping_delay = 300
      initial_refresh_delay = 30
      initial_refresh_spread = 60
      stats_reporting_period = 60
      description = "BMP Server 1"
      dscp_value = "ef"
      update_source = "Loopback0"
      vrf = "OOB"
      tcp_mss = 1460
      tcp_keep_alive = 60
    }
  ]
}

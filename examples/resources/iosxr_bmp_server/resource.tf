resource "iosxr_bmp_server" "example" {
  all_max_buffer_size = 15
  all_route_monitorings = [
    {
      advertisement_interval = 60
      route_mon              = "inbound-pre-policy"
      scan_time              = 5
    }
  ]
  servers = [
    {
      description            = "BMP Server 1"
      dscp_value             = "ef"
      flapping_delay         = 300
      host                   = "192.168.1.100"
      initial_delay          = 60
      initial_refresh_delay  = 30
      initial_refresh_spread = 60
      number                 = 1
      port                   = 5000
      shutdown               = false
      stats_reporting_period = 60
      tcp_keep_alive         = 60
      tcp_mss                = 1460
      update_source          = "Loopback0"
      vrf                    = "OOB"
    }
  ]
}

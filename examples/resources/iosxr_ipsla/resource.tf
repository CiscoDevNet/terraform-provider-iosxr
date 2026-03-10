resource "iosxr_ipsla" "example" {
  low_memory = 100000
  key_chain = "KEY_CHAIN"
  hw_timestamp_disable = true
  operations = [
    {
      operation_number = 1
      udp_echo = true
      udp_echo_tag = "UDP_ECHO_TAG"
      udp_echo_frequency = 60
      udp_echo_datasize_request = 64
      udp_echo_timeout = 5000
      udp_echo_source_ipv4 = "10.1.1.1"
      udp_echo_source_port = 1024
      udp_echo_destination_ipv4 = "10.1.1.2"
      udp_echo_destination_port = 7
      udp_echo_control_disable = false
      udp_echo_verify_data = false
      udp_echo_tos = 0
      udp_echo_vrf = "VRF1"
      udp_echo_statistics_hourly_buckets = 2
      udp_echo_statistics_hourly_distribution_count = 1
      udp_echo_statistics_hourly_distribution_interval = 20
        udp_echo_statistics_intervals = [
          {
            interval = 300
            buckets = 10
          }
        ]
      udp_echo_history_buckets = 15
      udp_echo_history_filter_all = true
      udp_echo_history_lives = 2
    }
  ]
  schedules = [
    {
      operation_number = 1
      life_time = 86200
      start_hour = 12
      start_minute = 0
      start_second = 0
      start_month = "january"
      start_day_of_month = 15
      start_year = 2032
      start_pending = false
      recurring = true
      ageout = 300
    }
  ]
  server_twamp = true
  server_twamp_port = 862
  server_twamp_timer_inactivity = 600
}

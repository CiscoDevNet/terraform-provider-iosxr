resource "iosxr_tcp" "example" {
  window_size = 32768
  synwait_time = 15
  path_mtu_discovery = true
  path_mtu_discovery_age_timer = "20"
  receive_queue = 200
  timestamp = true
  throttle = 40
  throttle_high_water_mark = 70
  selective_ack = true
  mss = 1460
  accept_rate = 500
  ao = true
  ao_keychains = [
    {
      keychain_name = "TCP_KEYCHAIN"
        keys = [
          {
            key_name = "200"
            send_id = 10
            receive_id = 20
          }
        ]
    }
  ]
}

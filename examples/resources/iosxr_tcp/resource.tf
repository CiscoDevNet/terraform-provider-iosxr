resource "iosxr_tcp" "example" {
  accept_rate = 500
  ao          = true
  ao_keychains = [
    {
      keychain_name = "TCP_KEYCHAIN"
      keys = [
        {
          key_name   = "200"
          receive_id = 20
          send_id    = 10
        }
      ]
    }
  ]
  mss                          = 1460
  path_mtu_discovery           = true
  path_mtu_discovery_age_timer = "20"
  receive_queue                = 200
  selective_ack                = true
  synwait_time                 = 15
  throttle                     = 40
  throttle_high_water_mark     = 70
  timestamp                    = true
  window_size                  = 32768
}

resource "iosxr_router_isis_interface" "example" {
  process_id     = "P1"
  interface_name = "GigabitEthernet0/0/0/1"
  circuit_type   = "level-1"
  hello_padding  = "disable"
  hello_padding_levels = [
    {
      level_number  = 1
      hello_padding = "always"
    }
  ]
  priority = 10
  priority_levels = [
    {
      level_number = 1
      priority     = 64
    }
  ]
  point_to_point                    = false
  state                             = "passive"
  hello_password_hmac_md5_encrypted = "060506324F41584B564B0F49584B"
  hello_password_hmac_md5_send_only = true
  hello_password_levels = [
    {
      level_number                        = 1
      level_hello_password_text_encrypted = "060506324F41584B564B0F49584B"
      level_hello_password_text_send_only = true
    }
  ]
  bfd_fast_detect_ipv4 = true
  bfd_fast_detect_ipv6 = true
  bfd_minimum_interval = 50
  bfd_multiplier       = 3
}

resource "iosxr_key_chain" "example" {
  name = "KEY11"
  keys = [
    {
      key_name                                = "1"
      key_string_password                     = "00071A150754"
      cryptographic_algorithm                 = "hmac-md5"
      accept_lifetime_start_time_hour         = 11
      accept_lifetime_start_time_minute       = 52
      accept_lifetime_start_time_second       = 55
      accept_lifetime_start_time_day_of_month = 21
      accept_lifetime_start_time_month        = "january"
      accept_lifetime_start_time_year         = 2023
      accept_lifetime_infinite                = true
      send_lifetime_start_time_hour           = 8
      send_lifetime_start_time_minute         = 36
      send_lifetime_start_time_second         = 22
      send_lifetime_start_time_day_of_month   = 15
      send_lifetime_start_time_month          = "january"
      send_lifetime_start_time_year           = 2023
      send_lifetime_infinite                  = true
    }
  ]
}

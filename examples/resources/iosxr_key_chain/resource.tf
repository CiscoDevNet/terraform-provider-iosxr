resource "iosxr_key_chain" "example" {
  name                   = "KEY11"
  accept_tolerance_value = 1000
  macsec_keys = [
    {
      ckn                                = "01"
      key_string_password                = "115949554642595C577A7F747D63637244574E535806580851040D531C465C0851020204055A0957540903554657520A5C575770151F08480746115A08552F7A22"
      key_string_cryptographic_algorithm = "aes-256-cmac"
      lifetime_start_time_hour           = 11
      lifetime_start_time_minute         = 52
      lifetime_start_time_second         = 55
      lifetime_start_time_month          = "january"
      lifetime_start_time_day_of_month   = 21
      lifetime_start_time_year           = 2023
      lifetime_end_time_hour             = 11
      lifetime_end_time_minute           = 52
      lifetime_end_time_second           = 55
      lifetime_end_time_month            = "january"
      lifetime_end_time_day_of_month     = 21
      lifetime_end_time_year             = 2026
    }
  ]
  keys = [
    {
      key_name                                = "1"
      key_string_password6                    = "00071A150754"
      cryptographic_algorithm                 = "hmac-md5"
      accept_lifetime_start_time_hour         = 11
      accept_lifetime_start_time_minute       = 52
      accept_lifetime_start_time_second       = 55
      accept_lifetime_start_time_month        = "january"
      accept_lifetime_start_time_day_of_month = 21
      accept_lifetime_start_time_year         = 2023
      accept_lifetime_end_time_hour           = 11
      accept_lifetime_end_time_minute         = 52
      accept_lifetime_end_time_second         = 55
      accept_lifetime_end_time_month          = "january"
      accept_lifetime_end_time_day_of_month   = 21
      accept_lifetime_end_time_year           = 2026
      send_lifetime_start_time_hour           = 8
      send_lifetime_start_time_minute         = 36
      send_lifetime_start_time_second         = 22
      send_lifetime_start_time_month          = "january"
      send_lifetime_start_time_day_of_month   = 15
      send_lifetime_start_time_year           = 2023
      send_lifetime_end_time_hour             = 8
      send_lifetime_end_time_minute           = 36
      send_lifetime_end_time_second           = 22
      send_lifetime_end_time_month            = "january"
      send_lifetime_end_time_day_of_month     = 15
      send_lifetime_end_time_year             = 2026
    }
  ]
  timezone_local = true
}

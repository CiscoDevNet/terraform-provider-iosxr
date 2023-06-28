resource "iosxr_service_timestamps" "example" {
  debug_datetime_localtime     = true
  debug_datetime_msec          = true
  debug_datetime_show_timezone = true
  debug_datetime_year          = true
  debug_uptime                 = true
  debug_disable                = true
  log_datetime_localtime       = true
  log_datetime_msec            = true
  log_datetime_show_timezone   = true
  log_datetime_year            = true
  log_uptime                   = true
  log_disable                  = true
}

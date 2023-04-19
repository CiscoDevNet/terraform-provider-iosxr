resource "iosxr_logging" "example" {
  ipv4_dscp                    = "cs6"
  trap                         = "informational"
  events_display_location      = true
  events_level                 = "informational"
  console                      = "disable"
  monitor                      = "disable"
  buffered_logging_buffer_size = 4000000
  buffered_level               = "debugging"
  facility_level               = "local7"
  hostnameprefix               = "HOSTNAME01"
  suppress_duplicates          = true
}

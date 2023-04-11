resource "iosxr_syslog" "example" {
  ipv4_dscp                    = "cs6"
  trap                         = "informational"
  events_display_location      = true
  events_level                 = "informational"
  console                      = "disable"
  monitor                      = "disable"
  buffered_logging_buffer_size = 4000000
  buffered_level               = "debugging"
  facility_level               = "local7"
  hostnameprefix               = "ALMTX1P01"
  suppress_duplicates          = true
  source_interfaces = [
    {
      source_interface_name = "Loopback10"
    }
  ]
}

resource "iosxr_logging" "example" {
  console                         = "disable"
  trap                            = "informational"
  monitor                         = "disable"
  console_facility                = "all"
  archive_disk0                   = true
  archive_frequency_daily         = true
  archive_filesize                = 100
  archive_size                    = 500
  archive_length                  = 4
  archive_severity                = "informational"
  archive_threshold               = 80
  ipv4_dscp                       = "cs6"
  ipv6_dscp                       = "ef"
  facility_level                  = "local7"
  buffered_entries_count          = 10000
  buffered_size                   = 4000000
  buffered_level                  = "debugging"
  buffered_discriminator_match1   = "BUFFERED1"
  buffered_discriminator_match2   = "BUFFERED2"
  buffered_discriminator_match3   = "BUFFERED3"
  buffered_discriminator_nomatch1 = "BUFFERED_NOMATCH1"
  buffered_discriminator_nomatch2 = "BUFFERED_NOMATCH2"
  buffered_discriminator_nomatch3 = "BUFFERED_NOMATCH3"
  container_all                   = true
  container_fetch_timestamp       = true
  file = [
    {
      file_name                                      = "logfile1"
      path                                           = "/disk0:"
      maxfilesize                                    = 1024
      severity                                       = "info"
      local_accounting_send_to_remote_facility_level = "local0"
      discriminator_match1                           = "MATCH1"
      discriminator_match2                           = "MATCH2"
      discriminator_match3                           = "MATCH3"
      discriminator_nomatch1                         = "NOMATCH1"
      discriminator_nomatch2                         = "NOMATCH2"
      discriminator_nomatch3                         = "NOMATCH3"
    }
  ]
  history        = "emergencies"
  history_size   = 500
  hostnameprefix = "HOSTNAME01"
  localfilesize  = 1000
  source_interfaces = [
    {
      name = "Loopback0"
      vrfs = [
        {
          name = "VRF1"
        }
      ]
    }
  ]
  suppress_duplicates = true
  format_rfc5424      = true
  yang                = "emergencies"
  suppress_rules = [
    {
      rule_name = "RULE1"
      alarms = [
        {
          message_category = "SECURITY"
          group_name       = "SSHD"
          message_code     = "INFO"
        }
      ]
      apply_all_of_router = true
    }
  ]
  events_buffer_size = 10000
  filter_matches = [
    {
      match = "MATCH1"
    }
  ]
  events_display_location = true
  events_level            = "informational"
  events_threshold        = 80
}

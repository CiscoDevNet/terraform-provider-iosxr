resource "iosxr_logging" "example" {
  console = "disable"
  trap    = "informational"
  monitor = "disable"
  # Not supported from version 25.1 and above
  console_facility = "all"
  # Not supported from version 25.1 and above
  archive_disk0 = true
  # Not supported from version 25.1 and above
  archive_frequency_daily = true
  archive_filesize        = 100
  archive_size            = 500
  archive_length          = 4
  archive_severity        = "informational"
  archive_threshold       = 80
  # Not supported from version 25.1 and above
  ipv4_dscp = "cs6"
  # Not supported from version 25.1 and above
  ipv6_dscp = "ef"
  # Not supported from version 25.1 and above
  facility_level         = "local7"
  buffered_entries_count = 100
  # Not supported from version 25.1 and above
  buffered_size = 4000000
  # Not supported from version 25.1 and above
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
      file_name = "logfile1"
      # Not supported from version 25.1 and above
      path = "/disk0:"
      # Not supported from version 25.1 and above
      maxfilesize = 1024
      # Not supported from version 25.1 and above
      severity = "info"
      # Not supported from version 25.1 and above
      local_accounting_send_to_remote_facility_level = "local0"
      discriminator_match1                           = "MATCH1"
      discriminator_match2                           = "MATCH2"
      discriminator_match3                           = "MATCH3"
      discriminator_nomatch1                         = "NOMATCH1"
      discriminator_nomatch2                         = "NOMATCH2"
      discriminator_nomatch3                         = "NOMATCH3"
      # Supported from version 25.1
      local_accounting = true
      # Supported from version 25.1
      send_to_remote = true
      # Supported from version 25.1
      send_to_remote_facility = "auth"
      # Supported from version 25.1
      path_maxfilesize = 100
      # Supported from version 25.1
      path_path_name = "TODO"
      # Supported from version 25.1
      path_severity = "alerts"
    }
  ]
  history = "emergencies"
  # Not supported from version 25.1 and above
  history_size   = 100
  hostnameprefix = "HOSTNAME01"
  localfilesize  = 1000
  # Not supported from version 25.1 and above
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
  # Not supported from version 25.1 and above
  format_rfc5424 = true
  yang           = "emergencies"
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
  # Supported from version 25.1
  buffered_buffered_level = "alerts"
  # Supported from version 25.1
  buffered_log_buffer_size = 100
  # Supported from version 25.1
  console_console_level = "alerts"
  # Supported from version 25.1
  console_discriminator_match1 = "TODO"
  # Supported from version 25.1
  console_discriminator_match2 = "TODO"
  # Supported from version 25.1
  console_discriminator_match3 = "TODO"
  # Supported from version 25.1
  console_discriminator_nomatch1 = "TODO"
  # Supported from version 25.1
  console_discriminator_nomatch2 = "TODO"
  # Supported from version 25.1
  console_discriminator_nomatch3 = "TODO"
  # Supported from version 25.1
  facility_all = "all"
  # Supported from version 25.1
  history_level = "alerts"
  # Supported from version 25.1
  monitor_monitor_level = "alerts"
}

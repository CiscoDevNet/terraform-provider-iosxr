resource "iosxr_logging" "example" {
  archive_disk0                   = true
  archive_filesize                = 100
  archive_frequency               = "daily"
  archive_length                  = 4
  archive_severity                = "informational"
  archive_size                    = 500
  archive_threshold               = 80
  buffered_discriminator_match1   = "BUFFERED1"
  buffered_discriminator_match2   = "BUFFERED2"
  buffered_discriminator_match3   = "BUFFERED3"
  buffered_discriminator_nomatch1 = "BUFFERED_NOMATCH1"
  buffered_discriminator_nomatch2 = "BUFFERED_NOMATCH2"
  buffered_discriminator_nomatch3 = "BUFFERED_NOMATCH3"
  buffered_entries_count          = 10000
  buffered_level                  = "debugging"
  buffered_size                   = 4000000
  console                         = "disable"
  console_discriminator_match1    = "CONSOLE1"
  console_discriminator_match2    = "CONSOLE2"
  console_discriminator_match3    = "CONSOLE3"
  console_discriminator_nomatch1  = "CONSOLE_NOMATCH1"
  console_discriminator_nomatch2  = "CONSOLE_NOMATCH2"
  console_discriminator_nomatch3  = "CONSOLE_NOMATCH3"
  console_facility                = "all"
  container_fetch_timestamp       = true
  events_buffer_size              = 10000
  events_display_location         = true
  events_level                    = "informational"
  events_threshold                = 80
  facility_level                  = "local7"
  file = [
    {
      discriminator_match1                           = "MATCH1"
      discriminator_match2                           = "MATCH2"
      discriminator_match3                           = "MATCH3"
      discriminator_nomatch1                         = "NOMATCH1"
      discriminator_nomatch2                         = "NOMATCH2"
      discriminator_nomatch3                         = "NOMATCH3"
      file_name                                      = "logfile1"
      local_accounting                               = true
      local_accounting_send_to_remote                = true
      local_accounting_send_to_remote_facility_level = "local0"
      maxfilesize                                    = 1024
      path                                           = "/disk0:"
      severity                                       = "informational"
    }
  ]
  filter_matches = [
    {
      match = "MATCH1"
    }
  ]
  format         = "rfc5424"
  history        = "emergencies"
  history_size   = 500
  hostnameprefix = "HOSTNAME01"
  ipv4_dscp      = "cs6"
  ipv6_dscp      = "ef"
  localfilesize  = 1000
  monitor        = "disable"
  source_interfaces = [
    {
      name = "Loopback0"
      vrf  = "default"
    }
  ]
  suppress_duplicates = true
  suppress_rules = [
    {
      alarms = [
        {
          group_name       = "SSHD"
          message_category = "SECURITY"
          message_code     = "INFO"
        }
      ]
      apply_all_of_router = true
      rule_name           = "RULE1"
    }
  ]
  tls_servers = [
    {
      address_ipv4     = "1.1.1.1"
      name             = "TLS-SERVER1"
      severity         = "informational"
      source_interface = "Loopback0"
      trustpoint       = "TRUSTPOINT1"
      vrf              = "VRF1"
    }
  ]
  trap = "informational"
  yang = "debugging"
}

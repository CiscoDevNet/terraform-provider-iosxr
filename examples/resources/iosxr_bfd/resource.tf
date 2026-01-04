resource "iosxr_bfd" "example" {
  echo_latency_detect                          = true
  echo_latency_detect_percentage               = 200
  echo_latency_detect_count                    = 10
  echo_startup_validate_force                  = true
  echo_ipv4_source                             = "10.1.1.1"
  echo_ipv4_bundle_per_member_minimum_interval = 200
  trap_singlehop_pre_mapped                    = true
  multipath_locations = [
    {
      location_id = "0/0/CPU0"
    }
  ]
  multihop_ttl_drop_threshold            = 200
  dampening_initial_wait                 = 3600
  dampening_secondary_wait               = 3200
  dampening_maximum_wait                 = 3100
  dampening_threshold                    = 60000
  dampening_extensions_down_monitoring   = true
  dampening_bundle_member_l3_only_mode   = true
  dampening_bundle_member_initial_wait   = 5184
  dampening_bundle_member_secondary_wait = 6184
  dampening_bundle_member_maximum_wait   = 7184
  bundle_coexistence_bob_blb             = "inherit"
  ipv6_checksum_disable                  = true
  interfaces = [
    {
      interface_name        = "GigabitEthernet0/0/0/0"
      echo_ipv4_source      = "12.1.1.1"
      ipv6_checksum_disable = true
      local_address         = "12.1.1.1"
      tx_interval           = 10000
      rx_interval           = 30000
      multiplier            = 40
    }
  ]
}

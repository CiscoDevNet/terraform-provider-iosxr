resource "iosxr_tpa" "example" {
  statistics_update_frequency = 60
  statistics_max_lpts_events = 10000000
  statistics_max_intf_events = 10000000
  logging_file_max_size_kb = 1024
  logging_rotation_max_files = 10
  vrfs = [
    {
      vrf_name = "default"
      ipv4_update_source_dataports = "GigabitEthernet0/0/0/0"
      ipv6_update_source_dataports = "GigabitEthernet0/0/0/0"
        ipv4_update_source_destinations = [
          {
            destination_interface = "MgmtEth0/RP0/CPU0/0"
            source_interface = "GigabitEthernet0/0/0/0"
          }
        ]
        ipv6_update_source_destinations = [
          {
            destination_interface = "MgmtEth0/RP0/CPU0/0"
            source_interface = "GigabitEthernet0/0/0/0"
          }
        ]
        east_west_interfaces = [
          {
            interface_name = "GigabitEthernet0/0/0/1"
            referenced_vrf = "default"
            referenced_interface = "GigabitEthernet0/0/0/1"
          }
        ]
    }
  ]
}

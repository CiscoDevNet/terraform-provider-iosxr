resource "iosxr_linux_networking" "example" {
  statistics_synchronization_sixty_seconds = true
  exposed_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/0"
      linux_managed = "disable"
      statistics_synchronization_sixty_seconds = true
    }
  ]
  vrfs = [
    {
      vrf_name = "default"
      ipv4_source_interface_default_route = "GigabitEthernet0/0/0/0"
      ipv6_source_interface_default_route = "GigabitEthernet0/0/0/0"
    }
  ]
}

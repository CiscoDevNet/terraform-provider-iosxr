resource "iosxr_linux_networking" "example" {
  exposed_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/0"
      linux_managed  = "disable"
    }
  ]
  vrfs = [
    {
      vrf_name                            = "default"
      ipv4_source_interface_default_route = "GigabitEthernet0/0/0/0"
      ipv6_source_interface_default_route = "GigabitEthernet0/0/0/0"
    }
  ]
}

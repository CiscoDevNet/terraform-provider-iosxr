resource "iosxr_interface_ipv6" "example" {
  interface_name     = "GigabitEthernet0/0/0/1"
  link_local_address = "fe80::1"
  link_local_zone    = "0"
  autoconfig         = false
  enable             = true
}

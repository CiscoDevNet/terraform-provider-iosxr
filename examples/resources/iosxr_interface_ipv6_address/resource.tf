resource "iosxr_interface_ipv6_address" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  address        = "2001::1"
  prefix_length  = 64
  zone           = "0"
}

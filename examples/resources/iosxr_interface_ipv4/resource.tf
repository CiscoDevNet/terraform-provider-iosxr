resource "iosxr_interface_ipv4" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  address        = "1.1.1.1"
  netmask        = "255.255.255.0"
}

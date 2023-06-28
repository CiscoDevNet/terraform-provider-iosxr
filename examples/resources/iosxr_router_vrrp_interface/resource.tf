resource "iosxr_router_vrrp_interface" "example" {
  interface_name       = "GigabitEthernet0/0/0/1"
  mac_refresh          = 14
  delay_minimum        = 1234
  delay_reload         = 4321
  bfd_minimum_interval = 255
  bfd_multiplier       = 33
}

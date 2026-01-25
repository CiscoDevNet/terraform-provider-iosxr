data "iosxr_router_vrrp_interface_ipv6" "example" {
  interface_name = "GigabitEthernet0/0/0/2"
  vrrp_id        = 124
}

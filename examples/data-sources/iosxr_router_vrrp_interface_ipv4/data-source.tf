data "iosxr_router_vrrp_interface_ipv4" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  vrrp_id        = 123
  version        = 2
}

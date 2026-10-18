data "iosxr_router_vrrp_interface_ipv4" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  version        = 2
  vrrp_id        = 123
}

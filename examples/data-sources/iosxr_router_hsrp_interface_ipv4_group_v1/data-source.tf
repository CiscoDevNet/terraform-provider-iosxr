data "iosxr_router_hsrp_interface_ipv4_group_v1" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  group_id       = 123
}

data "iosxr_router_isis_interface_address_family" "example" {
  af_name        = "ipv4"
  interface_name = "GigabitEthernet0/0/0/1"
  process_id     = "P1"
  saf_name       = "unicast"
}

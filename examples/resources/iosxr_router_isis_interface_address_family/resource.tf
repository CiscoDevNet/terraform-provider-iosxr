resource "iosxr_router_isis_interface_address_family" "example" {
  process_id     = "P1"
  interface_name = "GigabitEthernet0/0/0/1"
  af_name        = "ipv4"
  saf_name       = "unicast"
}

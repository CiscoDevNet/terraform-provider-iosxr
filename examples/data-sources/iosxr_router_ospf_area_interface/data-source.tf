data "iosxr_router_ospf_area_interface" "example" {
  process_name   = "OSPF1"
  area_id        = "0"
  interface_name = "Loopback1"
}

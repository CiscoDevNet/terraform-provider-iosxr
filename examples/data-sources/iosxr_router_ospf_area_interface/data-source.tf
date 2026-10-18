data "iosxr_router_ospf_area_interface" "example" {
  area_id        = "0"
  interface_name = "Loopback1"
  process_name   = "OSPF1"
}

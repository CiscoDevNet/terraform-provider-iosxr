data "iosxr_router_ospf_vrf_area_interface" "example" {
  process_name   = "OSPF1"
  vrf_name       = "VRF1"
  area_id        = "0"
  interface_name = "Loopback2"
}

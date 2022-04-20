resource "iosxr_router_ospf_vrf_area" "example" {
  process_name = "OSPF1"
  vrf_name     = "VRF1"
  area_id      = "0"
}

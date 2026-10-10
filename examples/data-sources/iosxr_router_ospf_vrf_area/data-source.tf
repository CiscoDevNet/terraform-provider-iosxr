data "iosxr_router_ospf_vrf_area" "example" {
  area_id      = "1"
  process_name = "OSPF1"
  vrf_name     = "VRF1"
}

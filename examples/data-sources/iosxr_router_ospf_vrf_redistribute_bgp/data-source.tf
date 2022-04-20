data "iosxr_router_ospf_vrf_redistribute_bgp" "example" {
  process_name = "OSPF1"
  vrf_name     = "VRF1"
  as_number    = "65001"
}

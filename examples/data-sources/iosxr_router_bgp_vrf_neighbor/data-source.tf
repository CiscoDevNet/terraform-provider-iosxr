data "iosxr_router_bgp_vrf_neighbor" "example" {
  address = "10.1.1.2"
  as_number = "65001"
  vrf_name = "VRF1"
}

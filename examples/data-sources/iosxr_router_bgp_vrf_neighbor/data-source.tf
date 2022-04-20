data "iosxr_router_bgp_vrf_neighbor" "example" {
  as_number        = "65001"
  vrf_name         = "VRF1"
  neighbor_address = "10.1.1.2"
}

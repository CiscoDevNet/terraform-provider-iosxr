data "iosxr_router_bgp_vrf_neighbor_address_family" "example" {
  as_number        = "65001"
  vrf_name         = "VRF1"
  neighbor_address = "10.1.1.2"
  af_name          = "ipv4-unicast"
}

data "iosxr_router_bgp_vrf_neighbor_address_family" "example" {
  as_number        = "123456"
  vrf_name         = "VRF33"
  neighbor_address = "44.44.44.44"
  af_name          = "ipv4-labeled-unicast"
}

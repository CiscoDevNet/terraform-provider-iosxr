data "iosxr_router_bgp_vrf_address_family" "example" {
  af_name   = "ipv4-unicast"
  as_number = "65001"
  vrf_name  = "VRF2"
}

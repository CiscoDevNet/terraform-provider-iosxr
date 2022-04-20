resource "iosxr_router_bgp_vrf_address_family_network" "example" {
  as_number  = "65001"
  vrf_name   = "VRF1"
  af_name    = "ipv4-unicast"
  address    = "10.1.0.0"
  masklength = 16
}

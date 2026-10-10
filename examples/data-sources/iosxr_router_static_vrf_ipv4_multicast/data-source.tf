data "iosxr_router_static_vrf_ipv4_multicast" "example" {
  prefix_address = "100.0.1.0"
  prefix_length  = 24
  vrf_name       = "VRF2"
}

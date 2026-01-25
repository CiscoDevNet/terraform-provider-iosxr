data "iosxr_router_static_vrf_ipv4_unicast" "example" {
  vrf_name       = "VRF2"
  prefix_address = "100.0.1.0"
  prefix_length  = 24
}

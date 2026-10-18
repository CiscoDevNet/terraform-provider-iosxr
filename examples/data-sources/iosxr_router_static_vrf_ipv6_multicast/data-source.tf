data "iosxr_router_static_vrf_ipv6_multicast" "example" {
  prefix_address = "1::"
  prefix_length  = 64
  vrf_name       = "VRF2"
}

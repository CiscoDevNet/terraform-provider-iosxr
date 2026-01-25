data "iosxr_router_static_ipv6_unicast" "example" {
  prefix_address = "1::"
  prefix_length  = 64
}

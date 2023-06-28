data "iosxr_router_static_ipv4_unicast" "example" {
  prefix_address = "100.0.1.0"
  prefix_length  = 24
}

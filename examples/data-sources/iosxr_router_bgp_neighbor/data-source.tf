data "iosxr_router_bgp_neighbor" "example" {
  as_number        = "65001"
  neighbor_address = "10.1.1.2"
}

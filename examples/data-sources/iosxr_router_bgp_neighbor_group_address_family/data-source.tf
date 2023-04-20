data "iosxr_router_bgp_neighbor_group_address_family" "example" {
  as_number           = "65001"
  neighbor_group_name = "NEIGHBOR1"
  af_name             = "ipv4 labled-unicast"
}

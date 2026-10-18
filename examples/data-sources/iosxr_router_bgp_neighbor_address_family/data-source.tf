data "iosxr_router_bgp_neighbor_address_family" "example" {
  address   = "10.1.1.2"
  af_name   = "vpnv4-unicast"
  as_number = "65001"
}

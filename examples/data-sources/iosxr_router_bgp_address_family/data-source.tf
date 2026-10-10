data "iosxr_router_bgp_address_family" "example" {
  af_name   = "ipv4-unicast"
  as_number = "65001"
}

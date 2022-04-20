data "iosxr_router_bgp_address_family_redistribute_isis" "example" {
  as_number     = "65001"
  af_name       = "ipv4-unicast"
  instance_name = "P1"
}

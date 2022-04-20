data "iosxr_router_bgp_address_family_redistribute_ospf" "example" {
  as_number  = "65001"
  af_name    = "ipv4-unicast"
  router_tag = "OSPF1"
}

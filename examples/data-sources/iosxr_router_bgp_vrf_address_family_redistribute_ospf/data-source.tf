data "iosxr_router_bgp_vrf_address_family_redistribute_ospf" "example" {
  as_number  = "65001"
  vrf_name   = "VRF1"
  af_name    = "ipv4-unicast"
  router_tag = "OSPF1"
}

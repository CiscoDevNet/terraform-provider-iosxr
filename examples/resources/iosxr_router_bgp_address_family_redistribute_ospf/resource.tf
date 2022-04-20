resource "iosxr_router_bgp_address_family_redistribute_ospf" "example" {
  as_number                    = "65001"
  af_name                      = "ipv4-unicast"
  router_tag                   = "OSPF1"
  match_internal               = true
  match_internal_external      = true
  match_internal_nssa_external = false
  match_external               = false
  match_external_nssa_external = false
  match_nssa_external          = false
  metric                       = 100
}

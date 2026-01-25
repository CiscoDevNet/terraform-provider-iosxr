data "iosxr_router_bgp_af_group" "example" {
  as_number     = "65001"
  af_group_name = "AFGROUP1"
  af_name       = "vpnv4-unicast"
}

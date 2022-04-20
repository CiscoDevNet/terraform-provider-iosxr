resource "iosxr_router_bgp_address_family_aggregate_address" "example" {
  as_number     = "65001"
  af_name       = "ipv4-unicast"
  address       = "10.0.0.0"
  masklength    = 8
  as_set        = false
  as_confed_set = false
  summary_only  = false
}

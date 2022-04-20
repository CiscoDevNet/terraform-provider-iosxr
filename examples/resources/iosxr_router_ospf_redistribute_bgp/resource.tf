resource "iosxr_router_ospf_redistribute_bgp" "example" {
  process_name = "OSPF1"
  as_number    = "65001"
  tag          = 3
  metric_type  = "1"
}

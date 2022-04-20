resource "iosxr_router_ospf_vrf_redistribute_isis" "example" {
  process_name  = "OSPF1"
  vrf_name      = "VRF1"
  instance_name = "P1"
  level_1       = true
  level_2       = false
  level_1_2     = false
  tag           = 3
  metric_type   = "1"
}

data "iosxr_router_ospf_vrf_redistribute_isis" "example" {
  process_name  = "OSPF1"
  vrf_name      = "VRF1"
  instance_name = "P1"
}

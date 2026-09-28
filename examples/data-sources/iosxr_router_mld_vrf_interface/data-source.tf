data "iosxr_router_mld_vrf_interface" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  vrf_name       = "VRF1"
}

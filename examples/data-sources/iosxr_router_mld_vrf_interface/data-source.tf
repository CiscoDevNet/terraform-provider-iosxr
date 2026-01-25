data "iosxr_router_mld_vrf_interface" "example" {
  vrf_name       = "VRF1"
  interface_name = "GigabitEthernet0/0/0/1"
}

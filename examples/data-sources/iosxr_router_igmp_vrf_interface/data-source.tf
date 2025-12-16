data "iosxr_router_igmp_vrf_interface" "example" {
  vrf_name       = "VRF1"
  interface_name = "GigabitEthernet0/0/0/1"
}

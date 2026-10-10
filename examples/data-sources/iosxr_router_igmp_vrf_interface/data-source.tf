data "iosxr_router_igmp_vrf_interface" "example" {
  interface_name = "GigabitEthernet0/0/0/2"
  vrf_name       = "VRF1"
}

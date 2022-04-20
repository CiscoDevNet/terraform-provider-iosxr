resource "iosxr_vrf_route_target_ip_address_format" "example" {
  vrf_name           = "VRF1"
  address_family     = "ipv4"
  sub_address_family = "unicast"
  direction          = "import"
  ip_address         = "1.1.1.1"
  index              = 1
  stitching          = true
}

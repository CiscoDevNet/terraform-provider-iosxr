resource "iosxr_interface" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  l2transport    = false
  point_to_point = false
  multipoint     = false
  shutdown       = true
  mtu            = 9000
  bandwidth      = 100000
  description    = "My Interface Description"
  vrf            = "VRF1"
}

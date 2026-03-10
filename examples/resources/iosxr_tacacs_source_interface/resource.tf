resource "iosxr_tacacs_source_interface" "example" {
  source_interface = "MgmtEth0/RP0/CPU0/0"
  source_interfaces = [
    {
      vrf       = "VRF1"
      interface = "MgmtEth0/RP0/CPU0/0"
    }
  ]
}

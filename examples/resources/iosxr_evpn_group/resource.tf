resource "iosxr_evpn_group" "example" {
  group_id = 1
  core_interfaces = [
    {
      interface_name = "Bundle-Ether111"
    }
  ]
}

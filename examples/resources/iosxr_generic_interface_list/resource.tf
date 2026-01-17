resource "iosxr_generic_interface_list" "example" {
  list_name = "INTF-LIST1"
  interfaces = [
    {
      interface_name = "Bundle-Ether101"
    }
  ]
}

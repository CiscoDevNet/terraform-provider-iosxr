resource "iosxr_generic_interface_list" "example" {
  interfaces = [
    {
      interface_name = "Bundle-Ether101"
    }
  ]
  list_name = "INTF-LIST1"
}

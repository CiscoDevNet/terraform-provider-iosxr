resource "iosxr_snmp_server_view" "example" {
  view_name = "VIEW12"
  name = [
    {
      name     = "iso"
      included = true
    }
  ]
}

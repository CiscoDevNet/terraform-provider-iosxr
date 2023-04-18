resource "iosxr_snmp_server_view" "example" {
  view_name = "VIEW12"
  mib_view_families = [
    {
      name     = "iso"
      included = true
    }
  ]
}

resource "iosxr_snmp_server_view" "example" {
  view_name = "VIEW12"
  mib_view_families = [
    {
      mib_view_family_name = "iso"
      included             = true
    }
  ]
}

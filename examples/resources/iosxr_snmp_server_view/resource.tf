resource "iosxr_snmp_server_view" "example" {
  view_name = "VIEW12"
  name = [
    {
      mib_view_family_name = "iso"
      included             = true
    }
  ]
}

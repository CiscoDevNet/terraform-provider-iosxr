resource "iosxr_snmp_vrf" "example" {
  vrf_name                                                       = "11.11.11.11"
  traps_unencrypted_unencrypted_string_version_v3_security_level = "true"
}
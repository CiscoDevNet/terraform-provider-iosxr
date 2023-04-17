resource "iosxr_snmp_server_vrf_host" "example" {
  vrf_name = "11.11.11.11"
  unencrypted_strings = [
    {
      community_string          = "COMMUNITY1"
      version_v3_security_level = "auth"
    }
  ]
}

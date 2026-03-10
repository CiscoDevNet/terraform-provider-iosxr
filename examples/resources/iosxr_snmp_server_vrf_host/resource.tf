resource "iosxr_snmp_server_vrf_host" "example" {
  vrf_name = "VRF1"
  address  = "11.11.11.11"
  traps_unencrypted_strings = [
    {
      community_string          = "COMMUNITY1"
      version_v3_security_level = "auth"
    }
  ]
  informs_unencrypted_strings = [
    {
      community_string          = "COMMUNITY2"
      version_v3_security_level = "auth"
    }
  ]
}

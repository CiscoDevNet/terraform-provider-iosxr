resource "iosxr_snmp_server_vrf" "example" {
  contexts = [
    {
      name = "CONTEXT1"
    }
  ]
  hosts = [
    {
      address = "11.11.11.11"
      informs_encrypted_default = [
        {
          community_string = "15021E0E082328"
          udp_port         = "1100"
          version_v2c      = true
        }
      ]
      informs_unencrypted_strings = [
        {
          community_string          = "COMMUNITY2"
          version_v3_security_level = "auth"
        }
      ]
      traps_encrypted_default = [
        {
          community_string = "15021E0E082328"
          udp_port         = "1100"
          version_v2c      = true
        }
      ]
      traps_unencrypted_strings = [
        {
          community_string = "COMMUNITY1"
          version_v2c      = true
        }
      ]
    }
  ]
  vrf_name = "VRF1"
}

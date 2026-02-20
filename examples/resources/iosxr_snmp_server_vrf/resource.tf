resource "iosxr_snmp_server_vrf" "example" {
  vrf_name = "VRF1"
  hosts = [
    {
      address = "11.11.11.11"
      traps_unencrypted_strings = [
        {
          community_string = "COMMUNITY1"
          version_v2c      = true
        }
      ]
      traps_encrypted_default = [
        {
          community_string = "15021E0E082328"
          udp_port         = "1100"
          version_v2c      = true
        }
      ]
      traps_encrypted_aes = [
        {
          community_string = "06253E2C5A471E1C5E"
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
      informs_encrypted_default = [
        {
          community_string = "15021E0E082328"
          udp_port         = "1100"
          version_v2c      = true
        }
      ]
      informs_encrypted_aes = [
        {
          community_string = "06253E2C5A471E1C5E"
          udp_port         = "1100"
          version_v2c      = true
        }
      ]
    }
  ]
  contexts = [
    {
      name = "CONTEXT1"
    }
  ]
}

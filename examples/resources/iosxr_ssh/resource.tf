resource "iosxr_ssh" "example" {
  client_algorithms_ciphers                 = ["aes128-ctr"]
  client_algorithms_key_exchanges           = ["ecdh-sha2-nistp521"]
  client_disable_hmac_sha1                  = true
  client_disable_hmac_sha2_256              = false
  client_disable_hmac_sha2_512              = false
  client_dscp                               = 16
  client_enable_cipher_3des_cbc             = true
  client_enable_cipher_aes_cbc              = true
  client_rekey_time                         = 60
  client_rekey_volume                       = 2048
  client_source_interface                   = "Loopback100"
  client_tcp_window_scale                   = 7
  client_v2                                 = true
  client_vrf                                = "MGMT"
  server_algorithms_ciphers                 = ["aes128-ctr"]
  server_algorithms_host_key_dsa            = true
  server_algorithms_host_key_ecdsa_nistp256 = true
  server_algorithms_host_key_ecdsa_nistp384 = true
  server_algorithms_host_key_ecdsa_nistp521 = true
  server_algorithms_host_key_ed25519        = true
  server_algorithms_host_key_rsa            = true
  server_algorithms_host_key_rsa_sha256     = true
  server_algorithms_host_key_rsa_sha512     = true
  server_algorithms_host_key_ssh_rsa        = true
  server_algorithms_host_key_x509v3_ssh_rsa = true
  server_algorithms_key_exchanges           = ["ecdh-sha2-nistp521"]
  server_disable_hmac_sha1                  = true
  server_disable_hmac_sha2_256              = false
  server_disable_hmac_sha2_512              = false
  server_dscp                               = 48
  server_enable_cipher_3des_cbc             = true
  server_enable_cipher_aes_cbc              = true
  server_logging                            = true
  server_max_auth_limit                     = 10
  server_netconf_port                       = 830
  server_netconf_vrfs = [
    {
      ipv4_access_list = "ACL1"
      ipv6_access_list = "ACL2"
      vrf_name         = "VRF2"
    }
  ]
  server_netconf_xml           = true
  server_port                  = 5522
  server_port_forwarding_local = true
  server_rate_limit            = 60
  server_rekey_time            = 60
  server_rekey_volume          = 2048
  server_session_limit         = 10
  server_tcp_window_scale      = 7
  server_usernames = [
    {
      keystring = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCv60WjxoM39LgPDbiW7ne3gu18q0NIVv0RE6rDLNal1quXZ6k5I9nV0WbPSqJLRm4Q2aHEGQ3NG2dJ5ZZ3xYDOm5X9JtMSjLFCJhSHVnGz6w+s8zPKiLmBjBD4VmxBKGMj0C/4LlZJ1F3yJfPTCzDwIMAMF8fJBJ8PqFKfvMTMqLkBfjB7xhXIx5N3jAZJdmxPkzdPPLnqLOKUjGKHRgmLWbynKZwRkjqvNJPQd3pf9Yb/HGqhWLvXc0z2xGlqODBhC3vLg0tlSKFpSdcJqj6eZLmKQ5BLHhZkJHDVdKzKNw5r0dBbLqFzF7nHiJ3uD+fUgPNzKOc7vF/TzLmNDlWr"
      username  = "cisco"
    }
  ]
  server_v2 = true
  server_vrfs = [
    {
      ipv4_access_list = "ACL1"
      ipv6_access_list = "ACL2"
      vrf_name         = "VRF1"
    }
  ]
  timeout = 60
}

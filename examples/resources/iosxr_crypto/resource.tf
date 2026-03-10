resource "iosxr_crypto" "example" {
  ca_trustpoint_system_description = "System trustpoint description"
  ca_trustpoint_system_enrollment_retry_count = 10
  ca_trustpoint_system_enrollment_retry_period = 5
  ca_trustpoint_system_enrollment_url = "http://ca.example.com"
  ca_trustpoint_system_sftp_username = "sftpuser"
  ca_trustpoint_system_sftp_password = "1511021F0725"
  ca_trustpoint_system_auto_enroll = 80
  ca_trustpoint_system_renewal_message_type_pkcsreq = true
  ca_trustpoint_system_skip_challenge_password = true
  ca_trustpoint_system_rsa_keypair = "KEY1"
  ca_trustpoint_system_ca_keypair_ecdsanistp521 = "KEY4"
  ca_trustpoint_system_crl_optional = true
  ca_trustpoint_system_query_url = "ldap://ca.example.com/certsrv"
  ca_trustpoint_system_ip_address = "10.1.1.1"
  ca_trustpoint_system_subject_name = "CN=Router1,OU=Network,O=Example,C=US"
  ca_trustpoint_system_subject_name_ca_certificate = "CN=CA,OU=Security,O=Example,C=US"
  ca_trustpoint_system_subject_alternative_name = "DNS:router1.example.com,IP:192.168.1.1"
  ca_trustpoint_system_serial_number = true
  ca_trustpoint_system_vrf = "VRF1"
  ca_trustpoint_system_lifetime_ca_certificate = 90
  ca_trustpoint_system_lifetime_certificate = 90
  ca_trustpoint_system_message_digest = "sha256"
  ca_trustpoints = [
    {
      trustpoint_name = "TP1"
      description = "Custom trustpoint"
      enrollment_retry_count = 10
      enrollment_retry_period = 5
      enrollment_url = "http://ca.example.com"
      sftp_username = "sftpuser"
      sftp_password = "1511021F0725"
      auto_enroll = 80
      renewal_message_type_renewalreq = true
      skip_challenge_password = true
      rsakeypair = "KEY1"
      crl_optional = true
      query_url = "ldap://ca.example.com/certsrv"
      ip_address = "10.1.1.2"
      subject_name = "CN=Router2,OU=Network,O=Example,C=US"
      subject_alternative_name = "DNS:router2.example.com,IP:192.168.1.2"
      serial_number = true
      vrf = "VRF1"
      message_digest = "sha256"
      method_est_credential_certificate = "EST-BOOTSTRAP"
    }
  ]
  ca_openssh_trustpoints = [
    {
      trustpoint_name = "OPENSSH-TP1"
      rsakeypair = "KEY1"
    }
  ]
  ca_http_proxy = "proxy.example.com"
  ca_http_proxy_port = 8080
  ca_source_interface_ipv4 = "Loopback0"
  ca_rsa_1024_disable = true
  ca_fqdn_check_ip_address_allow = true
  ca_crl_curl_timeout = 10
  fips_mode = true
}

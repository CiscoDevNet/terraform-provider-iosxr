resource "iosxr_tftp_server" "example" {
  vrfs = [
    {
      vrf_name                = "VRF1"
      ipv4_server_access_list = "ACL1"
      ipv4_server_max_servers = "no-limit"
      ipv4_server_homedir     = "disk0:"
      ipv4_server_dscp        = "default"
      ipv6_server_access_list = "ACL1"
      ipv6_server_max_servers = "no-limit"
      ipv6_server_homedir     = "disk0:"
      ipv6_server_dscp        = "default"
    }
  ]
}

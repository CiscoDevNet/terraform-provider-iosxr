resource "iosxr_telnet" "example" {
  ipv4_client_source_interface = "GigabitEthernet0/0/0/1"
  vrfs = [
    {
      vrf_name                = "ROI"
      ipv4_server_max_servers = 32
      ipv4_server_access_list = "ACCESS1"
      ipv6_server_max_servers = 34
      ipv6_server_access_list = "ACCESS11"
    }
  ]
  vrfs_dscp = [
    {
      vrf_name  = "TOI"
      ipv4_dscp = 55
    }
  ]
}

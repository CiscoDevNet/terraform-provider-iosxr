resource "iosxr_telnet" "example" {
  ipv4_client_source_interface = "GigabitEthernet0/0/0/1"
  ipv6_client_source_interface = "GigabitEthernet0/0/0/1"
  vrfs = [
    {
      ipv4_server_access_list = "ACCESS1"
      ipv4_server_max_servers = 32
      ipv6_server_access_list = "ACCESS11"
      ipv6_server_max_servers = 34
      vrf_name                = "ROI"
    }
  ]
  vrfs_dscp = [
    {
      ipv4_dscp = 55
      vrf_name  = "TOI"
    }
  ]
}

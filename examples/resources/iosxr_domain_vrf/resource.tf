resource "iosxr_domain_vrf" "example" {
  domains = [
    {
      domain_name = "example.com"
      order       = 12345
    }
  ]
  ipv4_hosts = [
    {
      host_name  = "HOST_NAME_IPV4"
      ip_address = ["10.0.0.10"]
    }
  ]
  ipv6_hosts = [
    {
      host_name    = "HOST_NAME_IPV6"
      ipv6_address = ["10::10"]
    }
  ]
  lookup_disable          = true
  lookup_source_interface = "Loopback214"
  multicast               = "multicast.cisco.com"
  name                    = "cisco.com"
  name_servers = [
    {
      address = "10.0.0.1"
      order   = 0
    }
  ]
  vrf_name = "TEST-VRF"
}

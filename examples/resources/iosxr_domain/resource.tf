resource "iosxr_domain" "example" {
  domains = [
    {
      domain_name = "example.com"
      order = 0
    }
  ]
  lookup_disable = true
  lookup_source_interface = "Loopback214"
  name = "cisco.com"
  ipv4_hosts = [
    {
      host_name = "HOST_NAME_IPV4"
      ip_address = ["10.0.0.10"]
    }
  ]
  name_servers = [
    {
      address = "10.0.0.1"
      order = 345
    }
  ]
  ipv6_hosts = [
    {
      host_name = "HOST_NAME_IPV6"
      ipv6_address = ["10::10"]
    }
  ]
  multicast = "multicast.cisco.com"
  default_flows_disable = true
}

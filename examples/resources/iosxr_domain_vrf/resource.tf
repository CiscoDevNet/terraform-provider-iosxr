resource "iosxr_domain_vrf" "example" {
  vrf_name = "TEST-VRF"
  domains = [
    {
      domain_name = "DOMAIN11"
      order       = 12345
    }
  ]
  lookup_disable          = true
  lookup_source_interface = "Loopback2147483647"
  name                    = "DNAME"
  ipv4_hosts = [
    {
      host_name  = "HOST-AGC"
      ip_address = ["10.0.0.0"]
    }
  ]
  name_servers = [
    {
      address = "10.0.0.1"
      order   = 0
    }
  ]
  ipv6_hosts = [
    {
      host_name    = "HOST-ACC"
      ipv6_address = ["10::10"]
    }
  ]
  multicast = "TESTACC"
}

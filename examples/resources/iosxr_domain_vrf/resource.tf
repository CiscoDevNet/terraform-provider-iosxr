resource "iosxr_domain_vrf" "example" {
  vrf_name = "WORD"
  list_domain = [
    {
      domain_name = "WORD"
      order       = 0
    }
  ]
  lookup_disable          = true
  lookup_source_interface = "Loopback2147483647"
  name                    = "WORD"
  ipv4_host = [
    {
      host_name  = "WORD"
      ip_address = ["10.0.0.0"]
    }
  ]
  name_server = [
    {
      address = "10.0.0.1"
      order   = 0
    }
  ]
  ipv6_host = [
    {
      host_name    = "WORD"
      ipv6_address = ["10::10"]
    }
  ]
  multicast = "WORD"
}

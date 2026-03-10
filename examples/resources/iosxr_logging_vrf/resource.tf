resource "iosxr_logging_vrf" "example" {
  vrf_name = "default"
  hostnames = [
    {
      name                    = "server.cisco.com"
      severity                = "info"
      port                    = 514
      operator                = "equals"
      facility                = "local0"
      hostname_source_address = "1.1.1.2"
    }
  ]
  host_ipv4_addresses = [
    {
      ipv4_address        = "1.1.1.1"
      severity            = "info"
      port                = 514
      operator            = "equals"
      facility            = "local0"
      ipv4_source_address = "1.1.1.2"
    }
  ]
  host_ipv6_addresses = [
    {
      ipv6_address        = "2001:db8::1"
      severity            = "info"
      port                = 514
      operator            = "equals-or-higher"
      facility            = "local0"
      ipv6_source_address = "2001:db8::2"
    }
  ]
}

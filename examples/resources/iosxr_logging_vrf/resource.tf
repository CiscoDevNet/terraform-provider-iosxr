resource "iosxr_logging_vrf" "example" {
  vrf_name = "default"
  host_ipv4_addresses = [
    {
      ipv4_address = "1.1.1.1"
      severity     = "info"
      port         = 514
      operator     = "equals"
    }
  ]
  host_ipv6_addresses = [
    {
      ipv6_address = "2001::1"
      severity     = "info"
      port         = 514
      operator     = "equals-or-higher"
    }
  ]
}

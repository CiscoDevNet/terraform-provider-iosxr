resource "iosxr_logging_vrf" "example" {
  vrf_name = "VRF1"
  host_ipv4_addresses = [
    {
      ipv4_address = "1.1.1.1"
      severity     = "info"
    }
  ]
  host_ipv6_addresses = [
    {
      ipv6_address = "2001::1"
      severity     = "info"
    }
  ]
}

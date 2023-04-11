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
      ipv6_address = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
      severity     = "info"
    }
  ]
}

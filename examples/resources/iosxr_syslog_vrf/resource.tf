resource "iosxr_syslog_vrf" "example" {
  vrf_name = "CORE-Mgmt"
  host_ipv4_addresses = [
    {
      ipv4_address = "10.5.110.120"
      severity     = "info"
    }
  ]
}

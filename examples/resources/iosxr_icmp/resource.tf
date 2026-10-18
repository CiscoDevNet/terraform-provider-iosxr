resource "iosxr_icmp" "example" {
  ipv4_mpls_extended_diagnostics      = true
  ipv4_rate_limit_unreachable_df_rate = 1000
  ipv4_rate_limit_unreachable_rate    = 1000
  ipv4_source_vrf                     = true
  ipv4_vrfs = [
    {
      extended_diagnostics_permitted_remote_addresses = [
        {
          address = "10.0.0.0"
          length  = 24
        }
      ]
      vrf_name = "VRF1"
    }
  ]
  ipv6_mpls_extended_diagnostics = true
  ipv6_source_vrf                = true
  ipv6_vrfs = [
    {
      extended_diagnostics_permitted_remote_addresses = [
        {
          address = "2001:db8::"
          length  = 64
        }
      ]
      vrf_name = "VRF1"
    }
  ]
}

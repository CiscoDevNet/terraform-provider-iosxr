resource "iosxr_icmp" "example" {
  ipv4_source_vrf = true
  ipv4_rate_limit_unreachable_rate = 1000
  ipv4_rate_limit_unreachable_df_rate = 1000
  ipv6_source_vrf = true
}

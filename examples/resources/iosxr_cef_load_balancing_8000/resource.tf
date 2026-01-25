resource "iosxr_cef_load_balancing_8000" "example" {
  platform_load_balance_hash_rotate = 1
  platform_load_balance_fields_userdata_ipv6_udp = [
    {
      location_string = "0_RP0_CPU0"
      ipv6_udp_hash_offset = 5
      ipv6_udp_hash_size = 3
    }
  ]
  platform_load_balance_fields_userdata_ipv4_udp = [
    {
      location_string = "0_RP0_CPU0"
      ipv4_udp_hash_offset = 5
      ipv4_udp_hash_size = 3
    }
  ]
  platform_load_balance_mpls_hashing_inner_non_ip_label_only = true
}

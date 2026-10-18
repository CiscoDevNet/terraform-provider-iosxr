resource "iosxr_ipv6" "example" {
  assembler_frag_hdr_incomplete_enable = true
  assembler_max_packets                = 40
  assembler_overlap_frag_drop_enable   = true
  assembler_reassembler_drop_enable    = true
  assembler_timeout                    = 50
  hop_limit                            = 123
  icmp_error_interval                  = 2111
  icmp_error_interval_bucket_size      = 123
  path_mtu_enable                      = true
  path_mtu_timeout                     = 10
  source_route                         = true
}

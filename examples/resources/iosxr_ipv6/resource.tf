resource "iosxr_ipv6" "example" {
  hop_limit                            = 123
  icmp_error_interval                  = 2111
  icmp_error_interval_bucket_size      = 123
  source_route                         = true
  assembler_timeout                    = 50
  assembler_max_packets                = 40
  assembler_reassembler_drop_enable    = true
  assembler_frag_hdr_incomplete_enable = true
  assembler_overlap_frag_drop_enable   = true
  path_mtu_enable                      = true
  path_mtu_timeout                     = 10
}

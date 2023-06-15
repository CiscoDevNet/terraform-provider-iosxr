resource "iosxr_ipv6_prefix_list" "example" {
  prefix_list_name = "LIST1"
  sequences = [
    {
      sequence_number        = 4096
      remark                 = "REMARK"
      permission             = "permit"
      prefix                 = "2001:db8:3333:4444:5555:6666:7777:8888"
      mask                   = 64
      match_prefix_length_eq = 10
      match_prefix_length_ge = 20
      match_prefix_length_le = 20
    }
  ]
}

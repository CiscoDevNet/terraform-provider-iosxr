resource "iosxr_ipv6_prefix_list" "example" {
  prefix_list_name = "LIST1"
  sequences = [
    {
      sequence_number = 4096
      permission = "permit"
      prefix = "2001:db8::"
      zone = "1"
      mask = 32
      match_prefix_length_ge = 64
      match_prefix_length_le = 128
    }
  ]
}

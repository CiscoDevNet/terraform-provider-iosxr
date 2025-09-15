resource "iosxr_ipv4_prefix_list" "example" {
  prefix_list_name = "LIST1"
  sequences = [
    {
      sequence_number        = 4096
      permission             = "deny"
      prefix                 = "10.1.0.0"
      mask                   = "255.255.0.0"
      match_prefix_length_eq = 12
      match_prefix_length_ge = 22
      match_prefix_length_le = 32
    }
  ]
}

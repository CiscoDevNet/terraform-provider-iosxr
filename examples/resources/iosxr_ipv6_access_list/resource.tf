resource "iosxr_ipv6_access_list" "example" {
  access_list_name = "TEST1"
  sequences = [
    {
      sequence_number                = 22
      permit_protocol                = "tcp"
      permit_source_address          = "1::1"
      permit_source_prefix_length    = 64
      permit_source_port_range_start = "100"
      permit_source_port_range_end   = "200"
      permit_destination_host        = "2::1"
      permit_destination_port_eq     = "10"
      permit_nexthop1_ipv6           = "3::3"
      permit_nexthop2_ipv6           = "4::4"
      permit_log                     = true
    }
  ]
}

resource "iosxr_ipv4_access_list" "example" {
  access_list_name = "ACCESS1"
  sequences = [
    {
      sequence_number                = 11
      permit_protocol                = "tcp"
      permit_source_address          = "18.0.0.0"
      permit_source_wildcard_mask    = "0.255.255.255"
      permit_source_port_range_start = "100"
      permit_source_port_range_end   = "200"
      permit_destination_host        = "11.1.1.1"
      permit_destination_port_eq     = "300"
      permit_dscp                    = "cs1"
      permit_ttl_eq                  = 10
      permit_nexthop1_ipv4           = "1.2.3.4"
      permit_nexthop2_ipv4           = "3.4.5.6"
      permit_log                     = true
    }
  ]
}

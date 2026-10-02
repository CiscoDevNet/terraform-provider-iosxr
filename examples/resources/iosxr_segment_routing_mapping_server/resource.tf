resource "iosxr_segment_routing_mapping_server" "example" {
  mapping_prefix_sid_address_family = [
    {
      addresses = [
        {
          attached              = true
          ip_address            = "10.1.1.0"
          prefix                = 24
          range                 = 10
          start_sid_index_range = 500
        }
      ]
      af_name = "ipv4"
    }
  ]
}

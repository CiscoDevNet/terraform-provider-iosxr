resource "iosxr_segment_routing_mapping_server" "example" {
  mapping_prefix_sid_address_family = [
    {
      af_name = "ipv4"
      prefix_addresses = [
        {
          address   = "10.1.1.0"
          attached  = true
          length    = "24"
          range     = 10
          sid_index = 500
        }
      ]
    }
  ]
}

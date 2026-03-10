resource "iosxr_segment_routing_v6" "example" {
  enable                 = true
  sid_holdtime           = 10
  logging_locator_status = true
  formats = [
    {
      name          = "usid-f3216"
      format_enable = true
    }
  ]
  locators = [
    {
      locator_enable         = true
      name                   = "Locator1"
      micro_segment_behavior = "unode-psp-usd"
      prefix                 = "fccc:0:214::"
      prefix_length          = 48
      anycast                = true
      algorithm              = 128
    }
  ]
  encapsulation_traffic_class_option = "value"
  encapsulation_traffic_class_value  = 1
  encapsulation_hop_limit_option     = "count"
  encapsulation_hop_limit_value      = 1
  encapsulation_source_address       = "fccc:0:214::1"
}

resource "iosxr_segment_routing_v6" "example" {
  enable                             = true
  encapsulation_hop_limit_option     = "count"
  encapsulation_hop_limit_value      = 1
  encapsulation_source_address       = "fccc:0:214::1"
  encapsulation_traffic_class_option = "value"
  encapsulation_traffic_class_value  = 1
  formats = [
    {
      format_enable = true
      name          = "usid-f3216"
    }
  ]
  locators = [
    {
      algorithm              = 128
      anycast                = true
      locator_enable         = true
      micro_segment_behavior = "unode-psp-usd"
      name                   = "Locator1"
      prefix                 = "fccc:0:214::"
      prefix_length          = 48
    }
  ]
  logging_locator_status = true
  sid_holdtime           = 10
}

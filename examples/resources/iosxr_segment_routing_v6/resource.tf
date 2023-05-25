resource "iosxr_segment_routing_v6" "example" {
  enable                       = true
  encapsulation_source_address = "fccc:0:214::1"
  locators = [
    {
      locator_enable         = true
      name                   = "Locator1"
      micro_segment_behavior = "unode-psp-usd"
      prefix                 = "fccc:0:214::"
      prefix_length          = 48
    }
  ]
}

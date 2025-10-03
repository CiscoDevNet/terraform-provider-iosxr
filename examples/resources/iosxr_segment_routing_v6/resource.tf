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
  formats = [
    {
      name                                         = "usid-f3216"
      format_enable                                = true
      usid_local_id_block_ranges_lib_start         = 57344
      usid_local_id_block_ranges_explict_lib_start = 65024
      usid_wide_local_id_block_explicit_range      = 65527
    }
  ]
}

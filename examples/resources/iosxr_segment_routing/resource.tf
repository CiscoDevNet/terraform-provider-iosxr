resource "iosxr_segment_routing" "example" {
  local_block_lower_bound = 15000
  local_block_upper_bound = 15999
  global_block_lower_bound = 16000
  global_block_upper_bound = 29999
  enable = true
}

data "iosxr_segment_routing_isis" "example" {
  process_id = "P1"
  af_name    = "ipv6"
  saf_name   = "unicast"
}

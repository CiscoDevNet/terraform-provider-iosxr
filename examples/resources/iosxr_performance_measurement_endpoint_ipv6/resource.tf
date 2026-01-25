resource "iosxr_performance_measurement_endpoint_ipv6" "example" {
  address = "2001:db8::1"
  vrf_name = "VRF1"
  source_address_ipv6 = "2001:1:1:1::100"
  description = "PM Endpoint for testing"
  delay_measurement = true
  delay_measurement_profile_name = "DELAY_PROFILE_1"
  segment_list_names = [
    {
      list_name = "SEG_LIST_1"
    }
  ]
  segment_routing = true
  segment_routing_te_explicit_segment_lists = [
    {
      list_name = "SEG_LIST_SR_1"
      reverse_path_segment_list = "SEG_LIST_REVERSE_1"
      insert_srh_sl_zero = true
    }
  ]
  segment_routing_te_explicit_reverse_path_list = "SEG_LIST_GLOBAL_REVERSE"
}

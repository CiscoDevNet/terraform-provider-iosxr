resource "iosxr_performance_measurement_endpoint_ipv4" "example" {
  address                        = "10.1.1.1"
  delay_measurement              = true
  delay_measurement_profile_name = "DELAY_PROFILE_1"
  description                    = "PM Endpoint for testing"
  segment_list_names = [
    {
      list_name = "SEG_LIST_1"
    }
  ]
  segment_routing                               = true
  segment_routing_te_explicit_reverse_path_list = "SEG_LIST_GLOBAL_REVERSE"
  segment_routing_te_explicit_segment_lists = [
    {
      insert_srh_sl_zero        = true
      list_name                 = "SEG_LIST_SR_1"
      reverse_path_segment_list = "SEG_LIST_REVERSE_1"
    }
  ]
  source_address_ipv4 = "10.1.1.100"
  vrf_name            = "VRF1"
}

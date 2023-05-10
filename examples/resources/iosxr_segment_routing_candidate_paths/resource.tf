resource "iosxr_segment_routing_candidate_paths" "example" {
  policy_name = "POLICY1"
  path_index  = 100
  candidate_paths_type = [
    {
      type               = "dynamic"
      pcep               = false
      metric_metric_type = "igp"
      hop_type           = "srv6"
      segment_list_name  = "LIST1"
    }
  ]
}

resource "iosxr_segment_routing_te_candidate_path" "example" {
  policy_name = "POLICY1"
  path_index  = 100
  candidate_paths_type = [
    {
      type               = "dynamic"
      pcep               = false
      metric_metric_type = "igp"
      hop_type           = "mpls"
      segment_list_name  = "dynamic"
    }
  ]
}

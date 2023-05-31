resource "iosxr_segment_routing_te_policy_candidate_path" "example" {
  policy_name = "POLICY1"
  path_index  = 100
  path_infos = [
    {
      type              = "dynamic"
      pcep              = true
      metric_type       = "igp"
      hop_type          = "mpls"
      segment_list_name = "dynamic"
    }
  ]
}

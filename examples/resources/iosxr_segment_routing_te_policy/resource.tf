resource "iosxr_segment_routing_te_policy" "example" {
  auto_route_force_sr_include = true
  auto_route_forward_class    = 1
  auto_route_include_all_ipv6 = true
  auto_route_include_prefixes = [
    {
      address = "192.168.2.0"
      af_type = "af-type-ipv4"
      length  = 24
    }
  ]
  auto_route_metric_constant_value = 100
  auto_route_metric_relative_value = -10
  auto_route_metric_type           = "relative"
  bandwidth                        = 100000
  binding_sid_mpls_label           = 16501
  binding_sid_type                 = "mpls-label-specified"
  candidate_paths_preferences = [
    {
      backup_ineligible            = true
      bidirectional_association_id = 10
      bidirectional_corouted       = true
      constraints_affinity_rules = [
        {
          affinities = [
            {
              affinity_name = "AFFINITY-1"
            }
          ]
          affinity_type = "affinity-exclude-any"
        }
      ]
      constraints_bounds = [
        {
          metric_type = "latency"
          type        = "bound-scope-cumulative"
          value       = 100000
        }
      ]
      constraints_disjoint_path_fallback_disable   = true
      constraints_disjoint_path_group_id           = 10
      constraints_disjoint_path_shortest_path      = true
      constraints_disjoint_path_sub_id             = 1
      constraints_disjoint_path_type               = "link"
      constraints_segment_rules_adjacency_sid_only = true
      constraints_segment_rules_protection_type    = "protected-preferred"
      constraints_segment_rules_sid_algorithm      = 128
      effective_metric_type                        = "igp"
      effective_metric_value                       = 1000
      path_index                                   = 1
      paths = [
        {
          anycast                = true
          hop_type               = "mpls"
          metric_margin_relative = 100
          metric_margin_type     = "relative"
          metric_sid_limit       = 6
          metric_type            = "latency"
          pcep                   = true
          segment_list_name      = "dynamic"
          sticky                 = true
          type                   = "dynamic"
        }
      ]
      pce_group = "PCE-1"
    }
  ]
  effective_metric_type                                         = "igp"
  effective_metric_value                                        = 1000
  ipv6_disable                                                  = true
  pce_group                                                     = "PCE-1"
  performance_measurement_liveness_backup_profile               = "LIVENESS-PROFILE-2"
  performance_measurement_liveness_invalidation_action          = "invalid-ation-action-down"
  performance_measurement_liveness_logging_session_state_change = true
  performance_measurement_liveness_profile                      = "LIVENESS-PROFILE-1"
  performance_measurement_reverse_path_segment_list             = "SEG-1"
  policy_color                                                  = 20
  policy_color_endpoint_address                                 = "192.168.1.1"
  policy_color_endpoint_type                                    = "end-point-type-ipv4"
  policy_name                                                   = "POLICY1"
  shutdown                                                      = false
  source_address                                                = "192.168.1.1"
  source_address_type                                           = "end-point-type-ipv4"
  srv6_locator_behavior                                         = "ub6-insert-reduced"
  srv6_locator_binding_sid_type                                 = "srv6-dynamic"
  srv6_locator_name                                             = "LOC1"
  steering_invalidation_drop                                    = true
  steering_labeled_services_disable                             = true
  transit_eligible                                              = true
}

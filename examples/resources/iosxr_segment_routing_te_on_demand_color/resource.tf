resource "iosxr_segment_routing_te_on_demand_color" "example" {
  color = 100
  dynamic_anycast_sid_inclusion = true
  dynamic_metric_type = "latency"
  dynamic_metric_margin_type = "relative"
  dynamic_metric_margin_relative = 100
  dynamic_pcep = true
  dynamic_disjoint_path_group_id = 10
  dynamic_disjoint_path_type = "link"
  dynamic_disjoint_path_sub_id = 1
  dynamic_disjoint_path_shortest_path = true
  dynamic_disjoint_path_fallback_disable = true
  dynamic_affinity_rules = [
    {
      affinity_type = "affinity-exclude-any"
        affinities = [
          {
            affinity_name = "AFFINITY-1"
          }
        ]
    }
  ]
  dynamic_bounds = [
    {
      type = "bound-scope-cumulative"
      metric_type = "latency"
      value = 100000
    }
  ]
  constraint_segments_protection_type = "protected-preferred"
  constraint_segments_sid_algorithm = 128
  steering_labeled_services_disable = true
  steering_invalidation_drop = true
  bandwidth = 100000
  performance_measurement_liveness_profile = "LIVENESS-PROFILE-1"
  performance_measurement_liveness_backup_profile = "LIVENESS-PROFILE-2"
  performance_measurement_liveness_logging_session_state_change = true
  performance_measurement_liveness_invalidation_action = "invalid-ation-action-down"
  performance_measurement_reverse_path_segment_list = "SEG-1"
  maximum_sid_depth = 6
  pce_group = "GROUP-1"
  source_address_type = "end-point-type-ipv4"
  source_address = "192.168.1.1"
  effective_metric_value = 1000
  effective_metric_type = "igp"
  srv6_locator_name = "LOC1"
  srv6_locator_binding_sid_type = "srv6-dynamic"
  srv6_locator_behavior = "ub6-insert-reduced"
}

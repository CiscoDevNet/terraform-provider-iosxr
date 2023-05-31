resource "iosxr_segment_routing_traffic_engineering" "example" {
  on_demand_colors = [
    {
      color                                           = 266
      srv6_locator_locator_name                       = "LOC11"
      srv6_locator_behavior                           = "ub6-insert-reduced"
      srv6_locator_binding_sid_type                   = "srv6-dynamic"
      source_address_source_address                   = "fccc:0:213::1"
      source_address_ip_address_type                  = "end-point-type-ipv6"
      effective_metric_metric_value_type_metric_value = 4444
      effective_metric_metric_value_type_metric_type  = "igp"
      constraint_segments_protection_type             = "protected-only"
      constraint_segments_sid_algorithm               = 128
    }
  ]
  policies = [
    {
      policy_name                             = "POLICY1"
      srv6_locator_locator_name               = "Locator11"
      srv6_locator_binding_sid_type           = "srv6-dynamic"
      srv6_locator_behavior                   = "ub6-insert-reduced"
      source_address_source_address           = "fccc:0:103::1"
      source_address_ip_address_type          = "end-point-type-ipv6"
      policy_color_endpoint_color             = 65534
      policy_color_endpoint_end_point_type    = "end-point-type-ipv6"
      policy_color_endpoint_end_point_address = "fccc:0:215::1"
    }
  ]
}

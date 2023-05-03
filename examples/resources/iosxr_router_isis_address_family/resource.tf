resource "iosxr_router_isis_address_family" "example" {
  process_id              = "P1"
  af_name                 = "ipv4"
  saf_name                = "unicast"
  metric_style_narrow     = false
  metric_style_wide       = true
  metric_style_transition = false
  metric_style_levels = [
    {
      level_id   = 1
      narrow     = false
      wide       = true
      transition = false
    }
  ]
  router_id_ip_address                            = "1.2.3.4"
  default_information_originate                   = true
  fast_reroute_delay_interval                     = 300
  fast_reroute_per_link_priority_limit_critical   = true
  fast_reroute_per_link_priority_limit_high       = false
  fast_reroute_per_link_priority_limit_medium     = false
  fast_reroute_per_prefix_priority_limit_critical = true
  fast_reroute_per_prefix_priority_limit_high     = false
  fast_reroute_per_prefix_priority_limit_medium   = false
  microloop_avoidance_protected                   = false
  microloop_avoidance_segment_routing             = true
  advertise_passive_only                          = true
  advertise_link_attributes                       = true
  mpls_ldp_auto_config                            = false
  mpls_traffic_eng_router_id_ip_address           = "1.2.3.4"
  mpls_traffic_eng_level_1_2                      = false
  mpls_traffic_eng_level_1                        = false
  mpls_traffic_eng_level_2_only                   = true
  spf_interval_maximum_wait                       = 5000
  spf_interval_initial_wait                       = 50
  spf_interval_secondary_wait                     = 200
  spf_prefix_priorities = [
    {
      priority = "critical"
      tag      = 100
    }
  ]
  segment_routing_mpls_sr_prefer = true
  maximum_redistributed_prefixes = 100
  maximum_redistributed_prefixes_levels = [
    {
      level_id         = 1
      maximum_prefixes = 1000
    }
  ]
}

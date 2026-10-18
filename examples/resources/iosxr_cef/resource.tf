resource "iosxr_cef" "example" {
  adjacency_route_override_rib                                     = true
  consistent_hashing_auto_recovery                                 = true
  encap_sharing_disable                                            = true
  load_balancing_mode_hierarchical_ecmp_min_paths                  = 16
  load_balancing_recursive_oor_mode_dampening_and_dlb              = true
  load_balancing_recursive_oor_mode_dampening_and_dlb_max_duration = 200
  load_balancing_recursive_oor_mode_dampening_resource_threshold   = 80
  load_balancing_recursive_oor_mode_dlb_resource_threshold         = 90
  ltrace_multiplier                                                = 8
  platform_lsm_frr_holdtime                                        = 60
  proactive_arp_nd_enable                                          = true
  retry_service_time                                               = 50
  retry_syslog_timer                                               = 120
  retry_timeout                                                    = 15
}

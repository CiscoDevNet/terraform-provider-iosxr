resource "iosxr_hw_module_profile" "example" {
  profile_load_balance_algorithm_hash_polynomial_index = 5
  profile_qos_max_classmap_size = "8"
  profile_qos_max_classmap_size_locations = [
    {
      location_name = "0/0/CPU0"
      max_classmap_size = "8"
    }
  ]
  profile_qos_qosg_dscp_mark_enable_first = 0
  profile_qos_qosg_dscp_mark_enable_second = 63
  profile_qos_free_buffer_int_threshold_set = 50
  profile_qos_free_buffer_int_threshold_clear = 75
  profile_qos_hqos_enable = true
  profile_qos_stats_collection = true
  profile_qos_ecn_marking_stats = true
  profile_qos_shared_policer_per_class_stats = true
  profile_qos_wred_stats_enable = true
  profile_qos_lag_scheduler = true
  profile_qos_conform_aware_policer = true
  profile_qos_arp_isis_priority_enable = true
  profile_qos_gre_exp_classification_enable = true
  profile_qos_egress_compensation_setting_force = true
  profile_qos_policer_scale = "64000"
  profile_qos_nif_hp_fifo_reserve_percent = 10
  profile_qos_nif_hp_fifo_reserve_locations = [
    {
      location_name = "0/0/CPU0"
      percent = 10
    }
  ]
  netflow_ipfix315_enable_locations = [
    {
      location_name = "0/0/CPU0"
      location_name2 = "0/0/CPU0"
    }
  ]
  netflow_sflow_enable_locations = [
    {
      location_name = "0/0/CPU0"
      location_name2 = "0/0/CPU0"
    }
  ]
  stats_tx_scale_enhanced_ingress_sr = true
  srv6_mode_micro_segment_format_f3216 = true
  srv6_encapsulation_l3_traffic_class_with_hoplimit_propagate = true
  sr_policy_v6_null_label_autopush = true
  oam_four8byte_cfm_maid_enable = true
  fib_bgp_pic_multipath_core_enable = true
  bgp_mp_pic_auto_protect_enable = true
}

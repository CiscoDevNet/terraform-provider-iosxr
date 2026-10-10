resource "iosxr_router_isis" "example" {
  adjacency_stagger                   = true
  adjacency_stagger_initial_neighbors = 5
  adjacency_stagger_max_neighbors     = 20
  affinity_maps = [
    {
      affinity_name = "22"
      bit_position  = 4
    }
  ]
  authentication_check_disable            = true
  distribute_link_state                   = true
  distribute_link_state_exclude_external  = true
  distribute_link_state_exclude_interarea = true
  distribute_link_state_instance_id       = 32
  distribute_link_state_level             = 2
  distribute_link_state_route_policy      = "ROUTE_POLICY_1"
  distribute_link_state_throttle          = 1
  extended_admin_group                    = "both"
  flex_algos = [
    {
      address_family = [
        {
          af_name       = "ipv4"
          maximum_paths = 10
          saf_name      = "unicast"
        }
      ]
      advertise_definition                      = true
      affinity_exclude_any                      = ["AFFINITY-2"]
      affinity_include_all                      = ["AFFINITY-1"]
      affinity_include_any                      = ["AFFINITY-1"]
      affinity_reverse_exclude_any              = ["AFFINITY-2"]
      affinity_reverse_include_all              = ["AFFINITY-1"]
      affinity_reverse_include_any              = ["AFFINITY-1"]
      auto_cost_reference_bandwidth             = 1000000000
      auto_cost_reference_bandwidth_granularity = 1000
      auto_cost_reference_group_mode            = true
      data_plane_ip                             = false
      data_plane_segment_routing                = true
      fast_reroute_disable                      = true
      maximum_delay                             = 1000000
      metric_type                               = "delay"
      microloop_avoidance_disable               = true
      minimum_bandwidth                         = 1000000000
      number                                    = 128
      prefix_metric                             = true
      priority                                  = 10
      srlg_exclude_any                          = ["SRLG-EXCLUDE-1"]
      ucmp_disable                              = true
    }
  ]
  hello_padding             = "adaptive"
  hostname_dynamic_disable  = true
  ignore_lsp_errors_disable = true
  iid_disable               = true
  instance_id               = 1
  is_type                   = "level-1"
  log_adjacency_changes     = true
  log_format_brief          = true
  log_pdu_drops             = true
  log_sizes = [
    {
      log_type    = "adjacency"
      size_number = 30
    }
  ]
  lsp_check_interval = 20
  lsp_check_interval_levels = [
    {
      level_number       = 1
      lsp_check_interval = 20
    }
  ]
  lsp_fast_flooding                   = true
  lsp_fast_flooding_max_lsp_tx        = 500
  lsp_fast_flooding_remote_psnp_delay = 1000
  lsp_gen_interval_initial_wait       = 50
  lsp_gen_interval_levels = [
    {
      initial_wait   = 50
      level_number   = 1
      maximum_wait   = 5000
      secondary_wait = 200
    }
  ]
  lsp_gen_interval_maximum_wait   = 5000
  lsp_gen_interval_secondary_wait = 200
  lsp_mtu                         = 1400
  lsp_mtu_levels = [
    {
      level_number = 1
      lsp_mtu      = 1400
    }
  ]
  lsp_password_hmac_md5_enable_poi    = true
  lsp_password_hmac_md5_encrypted     = "060506324F41584B564B0F49584B"
  lsp_password_hmac_md5_send_only     = true
  lsp_password_hmac_md5_snp_send_only = true
  lsp_password_levels = [
    {
      hmac_md5_enable_poi    = true
      hmac_md5_encrypted     = "060506324F41584B564B0F49584B"
      hmac_md5_send_only     = true
      hmac_md5_snp_send_only = true
      level_number           = 1
    }
  ]
  lsp_refresh_interval = 16000
  lsp_refresh_interval_levels = [
    {
      level_number         = 1
      lsp_refresh_interval = 16000
    }
  ]
  max_lsp_lifetime = 1200
  max_lsp_lifetime_levels = [
    {
      level_number     = 1
      max_lsp_lifetime = 1200
    }
  ]
  max_metric               = true
  max_metric_default_route = true
  max_metric_delay         = true
  max_metric_external      = true
  max_metric_interlevel    = true
  max_metric_levels = [
    {
      default_route        = true
      delay                = true
      external             = true
      interlevel           = true
      level_number         = 1
      on_startup_advertise = 300
      srv6_locator         = true
      te                   = true
    }
  ]
  max_metric_on_startup_advertise = 300
  max_metric_srv6_locator         = true
  max_metric_te                   = true
  min_lsp_arrival_initial_wait    = 40
  min_lsp_arrival_levels = [
    {
      initial_wait   = 40
      level_number   = 1
      maximum_wait   = 2000
      secondary_wait = 100
    }
  ]
  min_lsp_arrival_maximum_wait   = 2000
  min_lsp_arrival_secondary_wait = 100
  mpls_ldp_sync                  = true
  mpls_ldp_sync_level            = 1
  multi_part_tlv_disable         = true
  multi_part_tlv_disable_levels = [
    {
      level_number      = 1
      neighbor          = true
      prefix_tlvs       = true
      router_capability = true
    }
  ]
  multi_part_tlv_disable_neighbor          = true
  multi_part_tlv_disable_prefix_tlvs       = true
  multi_part_tlv_disable_router_capability = true
  nets = [
    {
      net_id = "49.0001.2222.2222.2222.00"
    }
  ]
  nsf_ietf                                     = true
  nsf_interface_expires                        = 2
  nsf_interface_timer                          = 5
  nsf_lifetime                                 = 10
  nsr                                          = true
  nsr_restart_time                             = 240
  oor_set_overload_bit_disable                 = true
  process_id                                   = "P1"
  protocol_shutdown                            = true
  psnp_interval                                = 100
  purge_transmit_strict                        = true
  purge_transmit_strict_value                  = "level-1"
  receive_application_flex_algo_delay_app_only = true
  segment_routing_global_block_lower_bound     = 16000
  segment_routing_global_block_upper_bound     = 29999
  set_overload_bit                             = true
  set_overload_bit_advertise_external          = true
  set_overload_bit_advertise_interlevel        = true
  set_overload_bit_levels = [
    {
      advertise_external           = true
      advertise_interlevel         = true
      level_number                 = 1
      on_startup_time_to_advertise = 300
    }
  ]
  set_overload_bit_on_startup_time_to_advertise = 300
  srlg_admin_weight                             = 500
  srlg_names = [
    {
      admin_weight = 500
      srlg_name    = "SRLG-1"
      static_ipv4_addresses = [
        {
          local_end_point  = "10.0.0.1"
          remote_end_point = "10.0.0.2"
        }
      ]
    }
  ]
}

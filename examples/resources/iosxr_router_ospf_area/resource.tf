resource "iosxr_router_ospf_area" "example" {
  process_name = "OSPF1"
  area_id      = "1"
  ranges = [
    {
      address       = "192.168.1.0"
      mask          = "255.255.255.0"
      advertise     = true
      not_advertise = false
    }
  ]
  default_cost = 100
  multi_area_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
      neighbors = [
        {
          address                 = "192.168.2.1"
          database_filter_all_out = true
          poll_interval           = 10
          cost                    = 100
        }
      ]
      authentication_key_encrypted = "110A1016141D4B"
      message_digest_keys = [
        {
          key_id        = 1
          md5_encrypted = "01100F175804"
        }
      ]
      authentication                 = true
      authentication_message_digest  = true
      authentication_keychain_name   = "KEY1"
      cost                           = 500
      cost_fallback                  = 600
      cost_fallback_threshold        = 10000
      hello_interval                 = 10
      dead_interval                  = 40
      retransmit_interval            = 1000
      transmit_delay                 = 100
      mtu_ignore_enable              = true
      database_filter_all_out_enable = true
      distribute_list_acl            = "ACL_1"
      packet_size                    = 1400
      fast_reroute_per_link_exclude_interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/1"
        }
      ]
      fast_reroute_per_link_lfa_candidate_interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/2"
        }
      ]
      fast_reroute_per_link_use_candidate_only_enable = true
      fast_reroute_per_prefix                         = true
      fast_reroute_per_prefix_exclude_interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/3"
        }
      ]
      fast_reroute_per_prefix_lfa_candidate_interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/4"
        }
      ]
      fast_reroute_per_prefix_use_candidate_only_enable             = true
      fast_reroute_per_prefix_remote_lfa_tunnel_mpls_ldp            = true
      fast_reroute_per_prefix_remote_lfa_maximum_cost               = 500
      fast_reroute_per_prefix_ti_lfa_enable                         = true
      fast_reroute_per_prefix_tiebreaker_downstream_index           = 10
      fast_reroute_per_prefix_tiebreaker_lc_disjoint_index          = 20
      fast_reroute_per_prefix_tiebreaker_lowest_backup_metric_index = 30
      fast_reroute_per_prefix_tiebreaker_node_protecting_index      = 40
      fast_reroute_per_prefix_tiebreaker_primary_path_index         = 50
      fast_reroute_per_prefix_tiebreaker_secondary_path_index       = 60
      fast_reroute_per_prefix_tiebreaker_interface_disjoint_index   = 70
      fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index        = 80
      passive_disable                                               = true
      delay_normalize_interval                                      = 2000
      delay_normalize_offset                                        = 0
    }
  ]
  mpls_traffic_eng                   = true
  mpls_ldp_sync                      = true
  mpls_ldp_sync_igp_shortcuts        = true
  mpls_ldp_auto_config               = true
  route_policy_in                    = "ROUTE_POLICY_1"
  route_policy_out                   = "ROUTE_POLICY_1"
  external_out_enable                = true
  summary_in_enable                  = true
  segment_routing_mpls               = true
  segment_routing_forwarding_mpls    = true
  distribute_link_state_excl_nssa    = true
  distribute_link_state_excl_summary = true
  authentication_key_encrypted       = "110A1016141D4B"
  message_digest_keys = [
    {
      key_id        = 1
      md5_encrypted = "01100F175804"
    }
  ]
  authentication                               = true
  authentication_message_digest                = true
  authentication_keychain_name                 = "KEY1"
  network_point_to_point                       = true
  cost                                         = 500
  cost_fallback_anomaly_delay_igp_metric_value = 500
  cost_fallback_anomaly_delay_te_metric_value  = 600
  hello_interval                               = 10
  dead_interval                                = 40
  priority                                     = 10
  retransmit_interval                          = 1000
  transmit_delay                               = 100
  flood_reduction_enable                       = true
  demand_circuit_enable                        = true
  mtu_ignore_enable                            = true
  database_filter_all_out_enable               = true
  passive_disable                              = true
  distribute_list_acl                          = "ACL_1"
  bfd_fast_detect                              = true
  bfd_fast_detect_strict_mode                  = true
  bfd_minimum_interval                         = 300
  bfd_multiplier                               = 3
  security_ttl                                 = true
  security_ttl_hops                            = 10
  prefix_suppression                           = true
  prefix_suppression_secondary_address         = true
  fast_reroute_per_link_exclude_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  fast_reroute_per_link_lfa_candidate_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/2"
    }
  ]
  fast_reroute_per_link_use_candidate_only_enable = true
  fast_reroute_per_prefix                         = true
  fast_reroute_per_prefix_exclude_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/3"
    }
  ]
  fast_reroute_per_prefix_lfa_candidate_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/4"
    }
  ]
  fast_reroute_per_prefix_use_candidate_only_enable             = true
  fast_reroute_per_prefix_remote_lfa_tunnel_mpls_ldp            = true
  fast_reroute_per_prefix_remote_lfa_maximum_cost               = 500
  fast_reroute_per_prefix_ti_lfa_enable                         = true
  fast_reroute_per_prefix_tiebreaker_downstream_index           = 10
  fast_reroute_per_prefix_tiebreaker_lc_disjoint_index          = 20
  fast_reroute_per_prefix_tiebreaker_lowest_backup_metric_index = 30
  fast_reroute_per_prefix_tiebreaker_node_protecting_index      = 40
  fast_reroute_per_prefix_tiebreaker_primary_path_index         = 50
  fast_reroute_per_prefix_tiebreaker_secondary_path_index       = 60
  fast_reroute_per_prefix_tiebreaker_interface_disjoint_index   = 70
  fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index        = 80
  loopback_stub_network_enable                                  = true
  link_down_fast_detect                                         = true
  weight                                                        = 1000
  delay_normalize_interval                                      = 2000
  delay_normalize_offset                                        = 0
  virtual_links = [
    {
      address                      = "192.168.1.4"
      hello_interval               = 10
      dead_interval                = 40
      retransmit_interval          = 1000
      transmit_delay               = 100
      authentication_key_encrypted = "110A1016141D4B"
      message_digest_keys = [
        {
          key_id        = 1
          md5_encrypted = "01100F175804"
        }
      ]
      authentication                = true
      authentication_message_digest = true
      authentication_keychain_name  = "KEY1"
    }
  ]
}

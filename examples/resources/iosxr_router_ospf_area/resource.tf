resource "iosxr_router_ospf_area" "example" {
  area_id                                      = "1"
  authentication                               = true
  authentication_key_encrypted                 = "110A1016141D4B"
  authentication_keychain_name                 = "KEY1"
  authentication_message_digest                = true
  bfd_fast_detect                              = true
  bfd_fast_detect_strict_mode                  = true
  bfd_minimum_interval                         = 300
  bfd_multiplier                               = 3
  cost                                         = 500
  cost_fallback_anomaly_delay_igp_metric_value = 500
  cost_fallback_anomaly_delay_te_metric_value  = 600
  database_filter_all_out_enable               = true
  dead_interval                                = 40
  default_cost                                 = 100
  demand_circuit_enable                        = true
  distribute_link_state_excl_nssa              = true
  distribute_link_state_excl_summary           = true
  distribute_list_in_acl                       = "ACL_1"
  external_out_enable                          = true
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
  fast_reroute_per_prefix_remote_lfa_maximum_cost               = 500
  fast_reroute_per_prefix_remote_lfa_tunnel_mpls_ldp            = true
  fast_reroute_per_prefix_ti_lfa_enable                         = true
  fast_reroute_per_prefix_tiebreaker_downstream_index           = 10
  fast_reroute_per_prefix_tiebreaker_interface_disjoint_index   = 70
  fast_reroute_per_prefix_tiebreaker_lc_disjoint_index          = 20
  fast_reroute_per_prefix_tiebreaker_lowest_backup_metric_index = 30
  fast_reroute_per_prefix_tiebreaker_node_protecting_index      = 40
  fast_reroute_per_prefix_tiebreaker_primary_path_index         = 50
  fast_reroute_per_prefix_tiebreaker_secondary_path_index       = 60
  fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index        = 80
  fast_reroute_per_prefix_use_candidate_only_enable             = true
  flood_reduction_enable                                        = true
  hello_interval                                                = 10
  link_down_fast_detect                                         = true
  loopback_stub_network_enable                                  = true
  message_digest_keys = [
    {
      key_id        = 1
      md5_encrypted = "01100F175804"
    }
  ]
  mpls_ldp_auto_config        = true
  mpls_ldp_sync               = true
  mpls_ldp_sync_igp_shortcuts = true
  mpls_traffic_eng            = true
  mtu_ignore_enable           = true
  multi_area_interfaces = [
    {
      authentication                 = true
      authentication_key_encrypted   = "110A1016141D4B"
      authentication_keychain_name   = "KEY1"
      authentication_message_digest  = true
      cost                           = 500
      cost_fallback                  = 600
      cost_fallback_threshold        = 10000
      database_filter_all_out_enable = true
      dead_interval                  = 40
      distribute_list_in_acl         = "ACL_1"
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
      fast_reroute_per_prefix_remote_lfa_maximum_cost               = 500
      fast_reroute_per_prefix_remote_lfa_tunnel_mpls_ldp            = true
      fast_reroute_per_prefix_ti_lfa_enable                         = true
      fast_reroute_per_prefix_tiebreaker_downstream_index           = 10
      fast_reroute_per_prefix_tiebreaker_interface_disjoint_index   = 70
      fast_reroute_per_prefix_tiebreaker_lc_disjoint_index          = 20
      fast_reroute_per_prefix_tiebreaker_lowest_backup_metric_index = 30
      fast_reroute_per_prefix_tiebreaker_node_protecting_index      = 40
      fast_reroute_per_prefix_tiebreaker_primary_path_index         = 50
      fast_reroute_per_prefix_tiebreaker_secondary_path_index       = 60
      fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index        = 80
      fast_reroute_per_prefix_use_candidate_only_enable             = true
      hello_interval                                                = 10
      interface_name                                                = "GigabitEthernet0/0/0/1"
      message_digest_keys = [
        {
          key_id        = 1
          md5_encrypted = "01100F175804"
        }
      ]
      mtu_ignore_enable = true
      neighbors = [
        {
          address                 = "192.168.2.1"
          cost                    = 100
          database_filter_all_out = true
          poll_interval           = 10
        }
      ]
      packet_size         = 1400
      passive_disable     = true
      retransmit_interval = 1000
      transmit_delay      = 100
    }
  ]
  network_point_to_point               = true
  passive_disable                      = true
  prefix_suppression                   = true
  prefix_suppression_secondary_address = true
  priority                             = 10
  process_name                         = "OSPF1"
  ranges = [
    {
      address       = "192.168.1.0"
      advertise     = true
      mask          = "255.255.255.0"
      not_advertise = false
    }
  ]
  retransmit_interval             = 1000
  route_policy_in                 = "ROUTE_POLICY_1"
  route_policy_out                = "ROUTE_POLICY_1"
  security_ttl                    = true
  security_ttl_hops               = 10
  segment_routing_forwarding_mpls = true
  segment_routing_mpls            = true
  summary_in_enable               = true
  transmit_delay                  = 100
  virtual_links = [
    {
      address                       = "192.168.1.4"
      authentication                = true
      authentication_key_encrypted  = "110A1016141D4B"
      authentication_keychain_name  = "KEY1"
      authentication_message_digest = true
      dead_interval                 = 40
      hello_interval                = 10
      message_digest_keys = [
        {
          key_id        = 1
          md5_encrypted = "01100F175804"
        }
      ]
      retransmit_interval = 1000
      transmit_delay      = 100
    }
  ]
  weight = 1000
}

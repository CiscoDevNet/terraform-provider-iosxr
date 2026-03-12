resource "iosxr_router_ospf_vrf_area" "example" {
  process_name = "OSPF1"
  vrf_name     = "VRF1"
  area_id      = "1"
  ranges = [
    {
      address       = "192.168.1.0"
      mask          = "255.255.255.0"
      advertise     = true
      not_advertise = false
    }
  ]
  default_cost                 = 100
  route_policy_in              = "ROUTE_POLICY_1"
  route_policy_out             = "ROUTE_POLICY_1"
  external_out_enable          = true
  summary_in_enable            = true
  authentication_key_encrypted = "110A1016141D4B"
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
  distribute_list_in_acl                       = "ACL_1"
  packet_size                                  = 1400
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
  sham_links = [
    {
      local_address                = "192.168.1.1"
      remote_address               = "192.168.1.2"
      cost                         = 100
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

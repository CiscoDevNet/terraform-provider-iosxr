resource "iosxr_router_pim_ipv6" "example" {
  accept_register                 = "REGISTER_ACL_V6"
  allow_rp                        = true
  allow_rp_group_list             = "ALLOW_GROUP_ACL_V6"
  allow_rp_list                   = "ALLOW_RP_ACL_V6"
  bsr_candidate_bsr_address       = "2001:db8::12"
  bsr_candidate_bsr_hash_mask_len = 126
  bsr_candidate_bsr_priority      = 100
  bsr_candidate_rps = [
    {
      address          = "2001:db8::13"
      bidir_group_list = "BSR_RP_ACL_V6"
      bidir_interval   = 60
      bidir_priority   = 192
      group_list       = "BSR_RP_ACL_V6"
      interval         = 60
      priority         = 192
    }
  ]
  bsr_relay_vrfs = [
    {
      listen   = true
      vrf_name = "VRF1"
    }
  ]
  convergence_link_down_prune_delay           = 10
  convergence_rpf_conflict_join_delay         = 5
  dr_priority                                 = 100
  global_maximum_bsr_crp_cache_maximum        = 1000
  global_maximum_bsr_crp_cache_threshold      = 800
  global_maximum_group_mappings_bsr           = 5000
  global_maximum_group_mappings_bsr_threshold = 4000
  global_maximum_packet_queue_high_priority   = 1048576
  global_maximum_packet_queue_low_priority    = 524288
  global_maximum_register_states              = 30000
  global_maximum_register_states_threshold    = 25000
  global_maximum_route_interfaces             = 500000
  global_maximum_route_interfaces_threshold   = 400000
  global_maximum_routes                       = 50000
  global_maximum_routes_threshold             = 40000
  hello_interval                              = 30
  interfaces = [
    {
      bfd_fast_detect                      = true
      bfd_minimum_interval                 = 150
      bfd_multiplier                       = 3
      bsr_border                           = true
      dr_priority                          = 100
      enable                               = true
      hello_interval                       = 30
      interface_name                       = "GigabitEthernet0/0/0/1"
      join_prune_interval                  = 60
      join_prune_mtu                       = 1500
      maximum_route_interfaces             = 10000
      maximum_route_interfaces_access_list = "INTF_MAX_ROUTES_ACL_V6"
      maximum_route_interfaces_threshold   = 8000
      neighbor_filter                      = "INTF_NEIGHBOR_ACL_V6"
      override_interval                    = 2500
      propagation_delay                    = 500
    }
  ]
  join_prune_interval                         = 60
  join_prune_mtu                              = 1500
  log_neighbor_changes                        = true
  maximum_bsr_crp_cache_maximum               = 1000
  maximum_bsr_crp_cache_threshold             = 800
  maximum_group_mappings_bsr                  = 5000
  maximum_group_mappings_bsr_threshold        = 4000
  maximum_register_states                     = 30000
  maximum_register_states_threshold           = 25000
  maximum_route_interfaces                    = 500000
  maximum_route_interfaces_threshold          = 400000
  maximum_routes                              = 50000
  maximum_routes_threshold                    = 40000
  mdt_c_multicast_announce_pim_join_tlv       = true
  mdt_c_multicast_migration_route_policy      = "PIM_POLICY"
  mdt_c_multicast_shared_tree_prune           = true
  mdt_c_multicast_shared_tree_prune_delay     = 3
  mdt_c_multicast_source_tree_prune_delay     = 3
  mdt_c_multicast_suppress_pim_data_signaling = true
  mdt_c_multicast_suppress_shared_tree_join   = true
  mdt_c_multicast_type                        = "pim"
  mdt_data_announce_interval                  = 60
  mdt_data_switchover_interval                = 30
  mdt_hello_interval                          = 60
  mdt_neighbor_filter                         = "MDT_NEIGHBOR_ACL_V6"
  neighbor_check_on_recv                      = true
  neighbor_check_on_send                      = true
  neighbor_filter                             = "NEIGHBOR_ACL_V6"
  old_register_checksum                       = true
  override_interval                           = 2500
  propagation_delay                           = 500
  register_source                             = "Loopback0"
  rp_addresses = [
    {
      access_list = "RP_ACL_V6"
      address     = "2001:db8::1"
      override    = true
    }
  ]
  rp_addresses_bidir = [
    {
      access_list = "BIDIR_ACL_V6"
      address     = "2001:db8::2"
      override    = true
    }
  ]
  rp_static_deny                    = "DENY_ACL_V6"
  rpf_topology_route_policy         = "PIM_POLICY"
  rpf_vector_allow_ebgp             = true
  rpf_vector_disable_ibgp           = true
  rpf_vector_standard_encoding      = true
  sg_expiry_timer                   = 180
  sg_list                           = "SG_ACL_V6"
  spt_threshold_infinity            = true
  spt_threshold_infinity_group_list = "SPT_GROUPS_V6"
  ssm_allow_override                = true
  ssm_range                         = "SSM_ACL_V6"
  suppress_data_registers           = true
  suppress_rpf_change_prunes        = true
}

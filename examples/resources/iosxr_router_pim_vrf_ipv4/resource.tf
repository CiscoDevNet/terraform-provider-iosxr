resource "iosxr_router_pim_vrf_ipv4" "example" {
  vrf_name = "VRF1"
  rp_addresses = [
    {
      address     = "10.1.1.1"
      access_list = "RP_ACL"
      override    = true
    }
  ]
  rp_static_deny                          = "DENY_ACL"
  accept_register                         = "REGISTER_ACL"
  suppress_data_registers                 = true
  register_source                         = "Loopback0"
  suppress_rpf_change_prunes              = true
  neighbor_filter                         = "NEIGHBOR_ACL"
  convergence_rpf_conflict_join_delay     = 5
  convergence_link_down_prune_delay       = 10
  spt_threshold_infinity                  = true
  spt_threshold_infinity_group_list       = "SPT_GROUPS"
  old_register_checksum                   = true
  neighbor_check_on_send                  = true
  neighbor_check_on_recv                  = true
  hello_interval                          = 30
  dr_priority                             = 100
  join_prune_interval                     = 60
  join_prune_mtu                          = 1500
  propagation_delay                       = 500
  override_interval                       = 2500
  maximum_routes                          = 10000
  maximum_routes_threshold                = 8000
  maximum_route_interfaces                = 100000
  maximum_route_interfaces_threshold      = 80000
  maximum_register_states                 = 5000
  maximum_register_states_threshold       = 4000
  maximum_group_mappings_bsr              = 1000
  maximum_group_mappings_bsr_threshold    = 800
  maximum_group_mappings_autorp           = 1000
  maximum_group_mappings_autorp_threshold = 800
  maximum_bsr_crp_cache                   = 500
  maximum_bsr_crp_cache_threshold         = 400
  log_neighbor_changes                    = true
  rpf_vector_allow_ebgp                   = true
  rpf_vector_disable_ibgp                 = true
  rpf_vector_standard_encoding            = true
  rpf_vector_injects = [
    {
      source_address = "10.1.1.100"
      source_mask    = 24
      rpf_vectors    = ["10.1.1.1"]
    }
  ]
  explicit_rpf_vector_injects = [
    {
      source_address = "10.2.1.100"
      source_mask    = 24
      rpf_vectors    = ["10.2.1.1"]
    }
  ]
  rpf_topology_route_policy                   = "PIM_POLICY"
  mdt_neighbor_filter                         = "MDT_NEIGHBOR_ACL"
  mdt_data_switchover_interval                = 30
  mdt_data_announce_interval                  = 60
  mdt_c_multicast_type                        = "pim"
  mdt_c_multicast_announce_pim_join_tlv       = true
  mdt_c_multicast_shared_tree_prune           = true
  mdt_c_multicast_suppress_shared_tree_join   = true
  mdt_c_multicast_suppress_pim_data_signaling = true
  mdt_c_multicast_shared_tree_prune_delay     = 3
  mdt_c_multicast_source_tree_prune_delay     = 3
  mdt_c_multicast_migration_route_policy      = "PIM_POLICY"
  allow_rp                                    = true
  allow_rp_list                               = "ALLOW_RP_ACL"
  allow_rp_group_list                         = "ALLOW_GROUP_ACL"
  sg_expiry_timer                             = 180
  sg_list                                     = "SG_ACL"
  ssm_range                                   = "SSM_ACL"
  ssm_allow_override                          = true
  auto_rp_listen_disable                      = true
  auto_rp_relay_vrfs = [
    {
      vrf_name = "VRF2"
      listen   = true
    }
  ]
  bsr_candidate_bsr_address       = "10.1.1.12"
  bsr_candidate_bsr_hash_mask_len = 30
  bsr_candidate_bsr_priority      = 100
  bsr_candidate_rps = [
    {
      address    = "10.1.1.13"
      group_list = "BSR_RP_ACL"
      priority   = 192
      interval   = 60
    }
  ]
  bsr_relay_vrfs = [
    {
      vrf_name = "VRF2"
      listen   = true
    }
  ]
  interfaces = [
    {
      interface_name                       = "GigabitEthernet0/0/0/1"
      enable                               = true
      dr_priority                          = 100
      hello_interval                       = 30
      join_prune_interval                  = 60
      join_prune_mtu                       = 1500
      propagation_delay                    = 500
      override_interval                    = 2500
      neighbor_filter                      = "INTF_NEIGHBOR_ACL"
      maximum_route_interfaces             = 10000
      maximum_route_interfaces_threshold   = 8000
      maximum_route_interfaces_access_list = "INTF_MAX_ROUTES_ACL"
      bfd_multiplier                       = 3
      bfd_minimum_interval                 = 150
      bfd_fast_detect                      = true
      bsr_border                           = true
    }
  ]
}

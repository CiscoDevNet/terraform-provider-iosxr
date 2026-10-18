resource "iosxr_multicast_routing_vrf" "example" {
  ipv4                                            = true
  ipv4_accounting_per_prefix                      = true
  ipv4_accounting_per_prefix_forward_only         = true
  ipv4_bgp_auto_discovery_core_tree               = "pim"
  ipv4_bgp_auto_discovery_dynamic_ssm_translation = true
  ipv4_bgp_auto_discovery_inherit_vrf_ri_vrf      = "VRF2"
  ipv4_bgp_auto_discovery_inter_as                = true
  ipv4_bgp_auto_discovery_leaf_info_required      = true
  ipv4_bgp_auto_discovery_receiver_site           = true
  ipv4_core_tree_protocol_rsvp_te                 = true
  ipv4_core_tree_protocol_rsvp_te_group_list      = "RSVP_ACL"
  ipv4_export_rt_four_byte_as = [
    {
      four_byte_as_number = 65537
      index_number        = 2
    }
  ]
  ipv4_export_rt_ip_address = [
    {
      index_number = 3
      ip_address   = "10.1.1.1"
    }
  ]
  ipv4_export_rt_two_byte_as = [
    {
      index_number       = 1
      two_byte_as_number = 65001
    }
  ]
  ipv4_import_rt_four_byte_as = [
    {
      four_byte_as_number = 65537
      index_number        = 2
    }
  ]
  ipv4_import_rt_ip_address = [
    {
      index_number = 3
      ip_address   = "10.1.1.1"
    }
  ]
  ipv4_import_rt_two_byte_as = [
    {
      index_number       = 1
      two_byte_as_number = 65001
    }
  ]
  ipv4_interfaces = [
    {
      boundary       = "BOUNDARY_ACL"
      enable         = true
      interface_name = "GigabitEthernet0/0/0/2"
      ttl_threshold  = 100
    }
  ]
  ipv4_interfaces_all_enable                  = true
  ipv4_log_traps                              = true
  ipv4_mdt_data_group_access_list             = "DATA_MDT_GROUP_ACL"
  ipv4_mdt_data_group_address                 = "232.111.1.2"
  ipv4_mdt_data_group_immediate_switch        = true
  ipv4_mdt_data_group_prefix_length           = 32
  ipv4_mdt_data_group_route_policy            = "PIM_POLICY"
  ipv4_mdt_data_group_threshold               = 10
  ipv4_mdt_data_maximum_access_list           = "DATA_MDT_ACL"
  ipv4_mdt_data_maximum_immediate_switch      = true
  ipv4_mdt_data_maximum_number                = 100
  ipv4_mdt_data_maximum_route_policy          = "PIM_POLICY"
  ipv4_mdt_data_maximum_threshold             = 5
  ipv4_mdt_default_ipv4_address               = "232.111.1.1"
  ipv4_mdt_default_segment_routing_mpls_color = "100"
  ipv4_mdt_mtu                                = 1500
  ipv4_mdt_source                             = "Loopback0"
  ipv4_multipath                              = true
  ipv4_multipath_hash_source                  = true
  ipv4_rate_per_route                         = true
  ipv4_ssm_allow_override                     = true
  ipv4_ssm_range                              = "SSM_ACL"
  ipv4_static_rpf_address_ranges = [
    {
      interface        = "GigabitEthernet0/0/0/2"
      next_hop_address = "10.1.1.1"
      prefix_address   = "10.10.10.10"
      prefix_length    = 32
    }
  ]
  ipv6                                            = true
  ipv6_accounting_per_prefix                      = true
  ipv6_accounting_per_prefix_forward_only         = true
  ipv6_bgp_auto_discovery_core_tree               = "pim"
  ipv6_bgp_auto_discovery_dynamic_ssm_translation = true
  ipv6_bgp_auto_discovery_inherit_vrf_ri_vrf      = "VRF2"
  ipv6_bgp_auto_discovery_inter_as                = true
  ipv6_bgp_auto_discovery_leaf_info_required      = true
  ipv6_bgp_auto_discovery_receiver_site           = true
  ipv6_core_tree_protocol_rsvp_te                 = true
  ipv6_core_tree_protocol_rsvp_te_group_list      = "RSVP_ACL"
  ipv6_export_rt_four_byte_as = [
    {
      four_byte_as_number = 65537
      index_number        = 2
    }
  ]
  ipv6_export_rt_ip_address = [
    {
      index_number = 3
      ip_address   = "10.1.1.1"
    }
  ]
  ipv6_export_rt_two_byte_as = [
    {
      index_number       = 1
      two_byte_as_number = 65001
    }
  ]
  ipv6_import_rt_four_byte_as = [
    {
      four_byte_as_number = 65537
      index_number        = 2
    }
  ]
  ipv6_import_rt_ip_address = [
    {
      index_number = 3
      ip_address   = "10.1.1.1"
    }
  ]
  ipv6_import_rt_two_byte_as = [
    {
      index_number       = 1
      two_byte_as_number = 65001
    }
  ]
  ipv6_interfaces = [
    {
      boundary       = "BOUNDARY_ACL"
      enable         = true
      interface_name = "GigabitEthernet0/0/0/2"
      ttl_threshold  = 100
    }
  ]
  ipv6_interfaces_all_enable             = true
  ipv6_log_traps                         = true
  ipv6_mdt_data_group_access_list        = "DATA_MDT_GROUP_ACL"
  ipv6_mdt_data_group_address            = "232.111.1.2"
  ipv6_mdt_data_group_immediate_switch   = true
  ipv6_mdt_data_group_prefix_length      = 32
  ipv6_mdt_data_group_route_policy       = "PIM_POLICY"
  ipv6_mdt_data_group_threshold          = 10
  ipv6_mdt_data_maximum_access_list      = "DATA_MDT_ACL"
  ipv6_mdt_data_maximum_immediate_switch = true
  ipv6_mdt_data_maximum_number           = 100
  ipv6_mdt_data_maximum_route_policy     = "PIM_POLICY"
  ipv6_mdt_data_maximum_threshold        = 5
  ipv6_mdt_default_ipv4_address          = "232.111.1.1"
  ipv6_mdt_mtu                           = 1500
  ipv6_mdt_source                        = "Loopback0"
  ipv6_multipath                         = true
  ipv6_multipath_hash_source_nexthop     = true
  ipv6_rate_per_route                    = true
  ipv6_ssm_allow_override                = true
  ipv6_ssm_range                         = "SSM_ACL"
  ipv6_static_rpf_address_ranges = [
    {
      interface        = "GigabitEthernet0/0/0/2"
      next_hop_address = "2001:db8:1::1"
      prefix_address   = "2001:db8:10::10"
      prefix_length    = 128
    }
  ]
  vrf_name = "VRF1"
}

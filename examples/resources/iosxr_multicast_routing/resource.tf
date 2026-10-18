resource "iosxr_multicast_routing" "example" {
  ipv4                                            = true
  ipv4_accounting_per_prefix                      = true
  ipv4_accounting_per_prefix_forward_only         = true
  ipv4_bgp_auto_discovery_core_tree               = "pim"
  ipv4_bgp_auto_discovery_dynamic_ssm_translation = true
  ipv4_bgp_auto_discovery_inherit_vrf_ri_vrf      = "VRF1"
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
  ipv4_forwarding_latency_delay = 100
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
  ipv4_interface_inheritance_disable = true
  ipv4_interfaces = [
    {
      boundary       = "BOUNDARY_ACL"
      enable         = true
      interface_name = "GigabitEthernet0/0/0/1"
      ttl_threshold  = 100
    }
  ]
  ipv4_interfaces_all_enable         = true
  ipv4_log_traps                     = true
  ipv4_maximum_disable               = true
  ipv4_mdt_mtu                       = 1500
  ipv4_mdt_source                    = "Loopback0"
  ipv4_mofrr_lockout_timer           = 30
  ipv4_mofrr_loss_detection_timer    = 30
  ipv4_multipath                     = true
  ipv4_multipath_hash_source         = true
  ipv4_oom_handling                  = true
  ipv4_rate_per_route                = true
  ipv4_route_rate_freshwalk_interval = 10
  ipv4_route_rate_interval           = 1000
  ipv4_route_rate_routes             = 1000
  ipv4_ssm_allow_override            = true
  ipv4_ssm_range                     = "SSM_ACL"
  ipv4_static_rpf_address_ranges = [
    {
      interface        = "GigabitEthernet0/0/0/1"
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
  ipv6_bgp_auto_discovery_inherit_vrf_ri_vrf      = "VRF1"
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
  ipv6_forwarding_latency_delay = 100
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
  ipv6_interface_inheritance_disable = true
  ipv6_interfaces = [
    {
      boundary       = "BOUNDARY_ACL"
      enable         = true
      interface_name = "GigabitEthernet0/0/0/1"
      ttl_threshold  = 100
    }
  ]
  ipv6_interfaces_all_enable         = true
  ipv6_log_traps                     = true
  ipv6_maximum_disable               = true
  ipv6_mdt_mtu                       = 1500
  ipv6_mdt_source                    = "Loopback0"
  ipv6_mofrr_lockout_timer           = 30
  ipv6_mofrr_loss_detection_timer    = 30
  ipv6_multipath                     = true
  ipv6_multipath_hash_source_nexthop = true
  ipv6_rate_per_route                = true
  ipv6_ssm_allow_override            = true
  ipv6_ssm_range                     = "SSM_ACL"
  ipv6_static_rpf_address_ranges = [
    {
      interface        = "GigabitEthernet0/0/0/1"
      next_hop_address = "2001:db8:1::1"
      prefix_address   = "2001:db8:10::10"
      prefix_length    = 128
    }
  ]
}

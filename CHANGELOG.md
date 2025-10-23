## 0.6.0 (IOS-XR Version 24.4.2 Compatibility)

- BREAKING CHANGE: Rename 'address_family_ipv4_unicast' to 'ipv4_unicast' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_import_route_policy' to 'ipv4_unicast_import_route_policy' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_export_route_policy' to 'ipv4_unicast_export_route_policy' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_multicast' to 'ipv4_multicast' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_flowspec' to 'ipv4_flowspec' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast' to 'ipv6_unicast' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_import_route_policy' to 'ipv6_unicast_import_route_policy' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_export_route_policy' to 'ipv6_unicast_export_route_policy' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_multicast' to 'ipv6_multicast' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_flowspec' to 'ipv6_flowspec' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_import_route_target_two_byte_as_format' to 'ipv4_unicast_import_route_target_two_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_import_route_target_four_byte_as_format' to 'ipv4_unicast_import_route_target_four_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_import_route_target_ip_address_format' to 'ipv4_unicast_import_route_target_ip_address_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_export_route_target_two_byte_as_format' to 'ipv4_unicast_export_route_target_two_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_export_route_target_four_byte_as_format' to 'ipv4_unicast_export_route_target_four_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv4_unicast_export_route_target_ip_address_format' to 'ipv4_unicast_export_route_target_ip_address_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_import_route_target_two_byte_as_format' to 'ipv6_unicast_import_route_target_two_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_import_route_target_four_byte_as_format' to 'ipv6_unicast_import_route_target_four_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_import_route_target_ip_address_format' to 'ipv6_unicast_import_route_target_ip_address_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_export_route_target_two_byte_as_format' to 'ipv6_unicast_export_route_target_two_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_export_route_target_four_byte_as_format' to 'ipv6_unicast_export_route_target_four_byte_as_format' in 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Rename 'address_family_ipv6_unicast_export_route_target_ip_address_format' to 'ipv6_unicast_export_route_target_ip_address_format' in 'iosxr_vrf' resource and data source
- Add 'bfd_fast_detect_ipv4' to 'iosxr_router_isis_interface' resource and data source
- Add 'bfd_minimum_interval' to 'iosxr_router_isis_interface' resource and data source
- Add 'bfd_multiplier' to 'iosxr_router_isis_interface' resource and data source
- Add 'interfaces' (type: List) to 'iosxr_evpn' resource and data source
- Add 'ethernet_segment_enable' to 'iosxr_evpn' resource and data source
- Add 'ethernet_segment_esi_zero' to 'iosxr_evpn' resource and data source
- Add 'vpn_id' to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'description' to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'bgp_route_target_import_two_byte_as_format' (type: List) to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'bgp_route_target_import_four_byte_as_format' (type: List) to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'bgp_route_target_import_ipv4_address_format' (type: List) to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'bgp_route_target_export_two_byte_as_format' (type: List) to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'bgp_route_target_export_four_byte_as_format' (type: List) to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'bgp_route_target_export_ipv4_address_format' (type: List) to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'advertise_mac' to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'locators' (type: List) to 'iosxr_evpn_segment_routing_srv6_evi' resource and data source
- Add 'segment_routing_srv6' to 'iosxr_evpn' resource and data source
- Add 'segment_routing_srv6_locators' (type: List) to 'iosxr_evpn' resource and data source
- Add 'segment_routing_srv6_usid_allocation_wide_local_id_block' to 'iosxr_evpn' resource and data source
- Add 'password' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'password_inheritance_disable' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'segment_routing_mpls_enable' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'adjacency_sid_indices' (type: List) to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'adjacency_sid_absolutes' (type: List) to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'formats' (type: List) to 'iosxr_segment_routing_v6' resource and data source
- BREAKING CHANGE: Update YANG path to 'Cisco-IOS-XR-segment-routing-ms-cfg:/sr/srv6' in 'iosxr_segment_routing_v6' resource and data source
- BREAKING CHANGE: Rename 'name' to 'locator_name' in 'iosxr_segment_routing_v6' resource and data source
- BREAKING CHANGE: Remove 'locator_enable' in 'iosxr_segment_routing_v6' resource and data source
- Add 'micro_segment_behavior_unode_psp_usd' to 'iosxr_segment_routing_v6' resource and data source
- BREAKING CHANGE: Update YANG path to 'Cisco-IOS-XR-um-segment-routing-cfg:/segment-routing' in 'iosxr_segment_routing' resource and data source
- BREAKING CHANGE: Remove 'rd_two_byte_as_number', 'rd_two_byte_as_index', 'rd_four_byte_as_number', 'rd_four_byte_as_index', 'rd_ipv4_address', 'rd_ipv4_address_index' in 'iosxr_vrf' resource and data source
- Add 'rd_two_byte_as_number' to 'iosxr_vrf' resource and data source
- Add 'rd_two_byte_as_index' to 'iosxr_vrf' resource and data source
- Add 'rd_four_byte_as_number' to 'iosxr_vrf' resource and data source
- Add 'rd_four_byte_as_index' to 'iosxr_vrf' resource and data source
- Add 'rd_ipv4_address' to 'iosxr_vrf' resource and data source
- Add 'rd_ipv4_address_index' to 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Remove 'address_family_ipv4_unicast_import_route_target_two_byte_as_format', 'address_family_ipv4_unicast_import_route_target_four_byte_as_format', 'address_family_ipv4_unicast_import_route_target_ip_address_format', 'address_family_ipv4_unicast_export_route_target_two_byte_as_format', 'address_family_ipv4_unicast_export_route_target_four_byte_as_format', 'address_family_ipv4_unicast_export_route_target_ip_address_format', 'address_family_ipv6_unicast_import_route_target_two_byte_as_format', 'address_family_ipv6_unicast_import_route_target_four_byte_as_format', 'address_family_ipv6_unicast_import_route_target_ip_address_format', 'address_family_ipv6_unicast_export_route_target_two_byte_as_format', 'address_family_ipv6_unicast_export_route_target_four_byte_as_format', 'address_family_ipv6_unicast_export_route_target_ip_address_format' (type: List) in 'iosxr_vrf' resource and data source
- Add 'address_family_ipv4_unicast_import_route_target_two_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv4_unicast_import_route_target_four_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv4_unicast_import_route_target_ip_address_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv4_unicast_export_route_target_two_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv4_unicast_export_route_target_four_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv4_unicast_export_route_target_ip_address_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv6_unicast_import_route_target_two_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv6_unicast_import_route_target_four_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv6_unicast_import_route_target_ip_address_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv6_unicast_export_route_target_two_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv6_unicast_export_route_target_four_byte_as_format' (type: List) to 'iosxr_vrf' resource and data source
- Add 'address_family_ipv6_unicast_export_route_target_ip_address_format' (type: List) to 'iosxr_vrf' resource and data source
- BREAKING CHANGE: Remove 'traps_bgp_cbgp2_updown', 'traps_bgp_bgp4_mib_updown' in 'iosxr_snmp_server' resource and data source
- Add 'traps_bgp_cbgp_two_updown' to 'iosxr_snmp_server' resource and data source
- Add 'traps_bgp_cbgp_two_enable' to 'iosxr_snmp_server' resource and data source
- Add 'traps_bgp_enable_updown' to 'iosxr_snmp_server' resource and data source
- Add 'traps_bgp_enable_cisco_bgp4_mib' to 'iosxr_snmp_server' resource and data source
- BREAKING CHANGE: Remove 'traps_isis_all', 'traps_isis_database_overload', 'traps_isis_manual_address_drops', 'traps_isis_corrupted_lsp_detected', 'traps_isis_attempt_to_exceed_max_sequence', 'traps_isis_id_len_mismatch', 'traps_isis_max_area_addresses_mismatch', 'traps_isis_own_lsp_purge', 'traps_isis_sequence_number_skip', 'traps_isis_authentication_type_failure', 'traps_isis_authentication_failure', 'traps_isis_version_skew', 'traps_isis_area_mismatch', 'traps_isis_rejected_adjacency', 'traps_isis_lsp_too_large_to_propagate', 'traps_isis_orig_lsp_buff_size_mismatch', 'traps_isis_protocols_supported_mismatch', 'traps_isis_adjacency_change', 'traps_isis_lsp_error_detected' in 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_all' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_database_overload' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_manual_address_drops' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_corrupted_lsp_detected' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_attempt_to_exceed_max_sequence' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_id_len_mismatch' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_max_area_addresses_mismatch' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_own_lsp_purge' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_sequence_number_skip' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_authentication_type_failure' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_authentication_failure' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_version_skew' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_area_mismatch' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_rejected_adjacency' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_lsp_too_large_to_propagate' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_orig_lsp_buff_size_mismatch' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_protocols_supported_mismatch' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_adjacency_change' to 'iosxr_snmp_server' resource and data source
- Add 'traps_isis_lsp_error_detected' to 'iosxr_snmp_server' resource and data source
- Add 'metric_levels' (type: List) to 'iosxr_router_isis_interface_address_family' resource and data source
- BREAKING CHANGE: Remove 'metric', 'maximum' in 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'metric_default_metric' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'metric_maximum' to 'iosxr_router_isis_interface_address_family' resource and data source
- BREAKING CHANGE: Remove 'prefix_sid_absolute', 'prefix_sid_index', 'prefix_sid_n_flag_clear' in 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_index_id' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_index_php_disable' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_index_explicit_null' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_index_n_flag_clear' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_absolute_id' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_absolute_php_disable' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_absolute_explicit_null' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_absolute_n_flag_clear' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_index_id' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_index_php_disable' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_index_explicit_null' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_index_n_flag_clear' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_absolute_id' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_absolute_php_disable' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_absolute_explicit_null' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_strict_spf_absolute_n_flag_clear' to 'iosxr_router_isis_interface_address_family' resource and data source
- Add 'prefix_sid_algorithms' (type: List) to 'iosxr_router_isis_interface_address_family' resource and data source
- BREAKING CHANGE: Remove 'hello_password_hmac_md5', 'hello_password_keychain', 'hello_password_text' in 'iosxr_router_isis_interface' resource and data source
- BREAKING CHANGE: Remove 'passive', 'shutdown', 'suppressed' in 'iosxr_router_isis_interface' resource and data source
- BREAKING CHANGE: Remove 'hello_padding_disable', 'hello_padding_sometimes' in 'iosxr_router_isis_interface' resource and data source
- Add 'hello_padding' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_padding_levels' (type: List) to 'iosxr_router_isis_interface' resource and data source
- Add 'state' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_accepts_encrypted' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_accepts_levels' (type: List) to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_text_encrypted' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_text_send_only' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_hmac_md5_encrypted' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_hmac_md5_send_only' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_keychain_name' to 'iosxr_router_isis_interface' resource and data source
- Add 'hello_password_levels' (type: List) to 'iosxr_router_isis_interface' resource and data source
- Add 'priority_levels' (type: List) to 'iosxr_router_isis_interface' resource and data source
- Add 'metric_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'tag' to 'iosxr_router_isis_address_family' resource and data source
- Add 'tag_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Rename 'redistribute_isis' to 'redistribute_isis_processes' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Rename 'instance_id' to 'isis_string' in 'iosxr_router_isis_address_family' resource and data source
- Add 'redistribute_route_level' to 'iosxr_router_isis_address_family' resource and data source
- Add 'metric' to 'iosxr_router_isis_address_family' resource and data source
- Add 'metric_type' to 'iosxr_router_isis_address_family' resource and data source
- Add 'down_flag_clear' to 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'spf_prefix_priorities' (type: List) in 'router_isis_address_family' resource and data source
- Add 'spf_interval_ietf' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_interval_ietf_initial_wait' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_interval_ietf_short_wait' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_interval_ietf_long_wait' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_interval_ietf_learn_interval' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_interval_ietf_holddown_interval' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_interval_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_critical_tag' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_critical_prefixlist_name' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_high_tag' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_high_prefixlist_name' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_medium_tag' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_medium_prefixlist_name' to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_critical_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_high_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'spf_prefix_priority_medium_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Rename 'mpls_traffic_eng_router_id_ip_address' to 'mpls_traffic_eng_router_id_ipv4_address' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Rename 'mpls_traffic_eng_router_id_interface' to 'mpls_traffic_eng_router_id_interface_name' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'microloop_avoidance_protected' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'microloop_avoidance_segment_routing' in 'iosxr_router_isis_address_family' resource and data source
- Add 'microloop_avoidance_enable' to 'iosxr_router_isis_address_family' resource and data source
- Add 'microloop_avoidance_enable_protected' to 'iosxr_router_isis_address_family' resource and data source
- Add 'microloop_avoidance_enable_segment_routing_route_policy' to 'iosxr_router_isis_address_family' resource and data source
- Add 'microloop_avoidance_rib_update_delay' to 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'fast_reroute_per_link_priority_limit_critical' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'fast_reroute_per_link_priority_limit_high' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'fast_reroute_per_link_priority_limit_medium' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'fast_reroute_per_prefix_priority_limit_critical' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'fast_reroute_per_prefix_priority_limit_high' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'fast_reroute_per_prefix_priority_limit_medium' in 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_priority_limit' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_priority_limit_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_use_candidate_only' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_srlg_protection_weighted_global' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_srlg_protection_weighted_global_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_remote_lfa_prefix_list' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_remote_lfa_prefix_list_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_load_sharing_disable' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_load_sharing_disable_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_tiebreaker_downstream_index' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_tiebreaker_lc_disjoint_index' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_tiebreaker_lowest_backup_metric_index' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_tiebreaker_node_protecting_index' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_tiebreaker_primary_path_index' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_tiebreaker_secondary_path_index' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_link_priority_limit' to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_link_priority_limit_levels' (type: List) to 'iosxr_router_isis_address_family' resource and data source
- Add 'fast_reroute_per_link_use_candidate_only' to 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Rename 'interface_name' to 'router_id_interface_name' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Rename 'ip_address' to 'router_id_ip_address' in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Rename 'level_id' to 'level_number' in 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'narrow' from 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'wide' from 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'transition' from 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- Add 'metric_style_narrow' to 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- Add 'metric_style_narrow_transition' to 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- Add 'metric_style_wide' to 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- Add 'metric_style_wide_transition' to 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- Add 'metric_style_transition' to 'metric_style_levels' list in 'iosxr_router_isis_address_family' resource and data source
- BREAKING CHANGE: Remove 'passive' from 'interfaces' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Remove 'suppressed' from 'interfaces' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Remove 'shutdown' from 'interfaces' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'state' (enumeration) to 'interfaces' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Remove 'priority' from 'interfaces' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'priority_levels' (type: List) to 'interfaces' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Change 'hello_padding' from Bool to String (enumeration) in 'interfaces' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Rename 'algorithm_number' to 'flex_algo_number' in 'flex_algos' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Rename 'metric_type_delay' to 'metric_type' in 'flex_algos' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Rename 'affinity_map_name' to 'name' in 'affinity_maps' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Remove 'lsp_password_keychain' from 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_accept_encrypted' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_accept_levels' (type: List) to 'router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_text_encrypted' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_text_send_only' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_text_snp_send_only' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_text_enable_poi' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_hmac_md5_encrypted' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_hmac_md5_send_only' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_hmac_md5_snp_send_only' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_hmac_md5_enable_poi' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_keychain_name' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_keychain_send_only' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_keychain_snp_send_only' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'lsp_password_keychain_enable_poi' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Rename 'lsp_refresh_interval_lsp_refresh_interval_time' to 'lsp_refresh_interval' in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Rename 'max_lsp_lifetime_max_lsp_lifetime' to 'max_lsp_lifetime' in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Rename 'level_id' to 'level_number' in 'set_overload_bit_levels' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'nsf_cisco' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'nsf_ietf' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'set_overload_bit_on_startup_time_to_advertise_seconds' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'set_overload_bit_on_startup_wait_for_bgp' to 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'on_startup_time_to_advertise_seconds' to 'set_overload_bit_levels' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'on_startup_wait_for_bgp' to 'set_overload_bit_levels' list in 'iosxr_router_isis' resource and data source
- BREAKING CHANGE: Add 'segment_routing_srv6_usid_allocation_wide_local_id_block' in 'iosxr_router_bgp' resource and data source
- BREAKING CHANGE: Add 'segment_routing_srv6_alloc_mode_per_ce' in 'iosxr_router_bgp_vrf_address_family' resource and data source
- BREAKING CHANGE: Add 'label_mode_per_vrf' to 'segment_routing_srv6_alloc_mode_per_vrf' in 'iosxr_router_bgp_vrf_address_family' resource and data source
- BREAKING CHANGE: Add 'label_mode_per_vrf_46' to 'segment_routing_srv6_alloc_mode_per_vrf_46' in 'iosxr_router_bgp_vrf_address_family' resource and data source
- BREAKING CHANGE: Add 'label_mode_route_policy' to 'segment_routing_srv6_alloc_mode_route_policy' in 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'timers_zero' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'timers_keepalive_zero' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'timers_zero_minimum_acceptable_holdtime' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'timers_holdtime' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'timers_holdtime_number' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'timers_holdtime_zero' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'timers_holdtime_minimum_acceptable_holdtime' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'timers_zero' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'timers_bgp_keepalive_zero' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'timers_zero_minimum_acceptable_holdtime' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'timers_holdtime' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'timers_holdtime_number' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'timers_holdtime_zero' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'timers_holdtime_minimum_acceptable_holdtime' to 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'local_as_replace_as' to 'local_as_no_prepend_replace_as' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'local_as_dual_as' to 'local_as_no_prepend_replace_as_dual_as' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'neighbor' to 'address' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'rd_auto' to 'rd_auto' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'rd_two_byte_as_as_number' to 'rd_two_byte_as_number' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'rd_two_byte_as_index' to 'rd_two_byte_as_index' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'rd_four_byte_as_as_number' to 'rd_four_byte_as_number' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'rd_four_byte_as_index' to 'rd_four_byte_as_index' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'rd_ip_address_ipv4_address' to 'rd_ipv4_address_address' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'rd_ip_address_index' to 'rd_ipv4_address_index' in 'iosxr_router_bgp_vrf' resource and data source
- Add 'rd_two_byte_as' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'rd_four_byte_as' to 'iosxr_router_bgp_vrf' resource and data source
- Add 'rd_ipv4_address' to 'iosxr_router_bgp_vrf' resource and data source

- BREAKING CHANGE: Remove 'ao_key_chain_name' in 'iosxr_router_bgp_neighbor_group' resource and data source
- BREAKING CHANGE: Remove 'ao_include_tcp_options_enable' in 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'ao_key_chain_name' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'ao_key_chain_name_name' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'ao_key_chain_include_tcp_options' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'ao_key_chain_accept_ao_mismatch_connection' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'ao_inheritance_disable' to 'iosxr_router_bgp_neighbor_group' resource and data source

- BREAKING CHANGE: Rename 'default_originate_route_policy' in 'iosxr_router_bgp_neighbor_address_family' resource and data source
- BREAKING CHANGE: Rename 'default_originate_inheritance_disable' in 'iosxr_router_bgp_neighbor_address_family' resource and data source
- BREAKING CHANGE: Rename 'default_originate_route_policy' in 'iosxr_router_bgp_vrf_neighbor_address_family' resource and data source
- BREAKING CHANGE: Rename 'default_originate_inheritance_disable' in 'iosxr_router_bgp_vrf_neighbor_address_family' resource and data source
- Add 'redistribute_ospf' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_ospfv3' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_eigrp' in 'iosxr_router_bgp_address_family' resource and data source
- BREAKING CHANGE: Rename 'redistribute_isis' (yang_name: 'redistribute/isis' to 'redistribute/isis-processes/isis-process') in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_connected_multipath' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_rip' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_rip_metric' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_rip_multipath' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_rip_route_policy' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'redistribute_static_multipath' in 'iosxr_router_bgp_address_family' resource and data source

- BREAKING CHANGE: Remove 'import' in 'iosxr_router_bgp_neighbor_address_family' resource and data source
- Add 'import_stitching_rt' to 'iosxr_router_bgp_neighbor_address_family' resource and data source
- Add 'import_stitching_rt_re_originate' to 'iosxr_router_bgp_neighbor_address_family' resource and data source
- Add 'import_stitching_rt_re_originate_stitching_rt' to 'iosxr_router_bgp_neighbor_address_family' resource and data source
- Add 'import_re_originate' to 'iosxr_router_bgp_neighbor_address_family' resource and data source
- BREAKING CHANGE: Rename 'encapsulation-type/srv6' to 'encapsulation_type' in 'iosxr_router_bgp_neighbor_address_family' resource and data source

- BREAKING CHANGE: Rename 'masklength' to 'address_prefix' (yang_name: 'masklength' to 'address-prefix') in 'iosxr_router_bgp_address_family' resource and data source
- Add 'route_policy' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'description' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'set_tag' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'backdoor' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'multipath' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_per_prefix' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_per_ce' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_per_vrf' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_per_vrf_46' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_route_policy' to 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_route_policy_name' to 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_per_nexthop_received_label' to 'iosxr_router_bgp_address_family' resource and data source
- Add 'label_mode_per_nexthop_received_label_allocate_secondary_label' to 'iosxr_router_bgp_address_family' resource and data source

- BREAKING CHANGE: Rename 'masklength' to 'address_prefix' in 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'route_policy', 'description', 'set_tag' to 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'route_policy', 'backdoor', 'multipath' to 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'redistribute_isis_processes' (type: List) to 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'level_one', 'level_one_level_two', 'level_one_level_two_level_one_inter_area', 'level_one_level_one_inter_area', 'level_two', 'level_two_level_one_inter_area', 'level_one_inter_area' to 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'metric', 'multipath', 'route_policy' to 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'label_mode_per_prefix', 'label_mode_per_ce', 'label_mode_per_vrf', 'label_mode_per_vrf_46' to 'iosxr_router_bgp_address_family' resource and data source

- BREAKING CHANGE: Remove 'echo_ipv4_bundle_per_member_preferred_minimum_interval' in 'iosxr_bfd' resource and data source
- Add 'echo_ipv4_bundle_per_member_minimum_interval' to 'iosxr_bfd' resource and data source
- BREAKING CHANGE: Rename 'location-name' to 'location-id' in 'iosxr_bfd' resource and data source

- BREAKING CHANGE: Remove 'ethernet_segment_identifier_type_zero_bytes_1', 'ethernet_segment_identifier_type_zero_bytes_23', 'ethernet_segment_identifier_type_zero_bytes_45', 'ethernet_segment_identifier_type_zero_bytes_67', 'ethernet_segment_identifier_type_zero_bytes_89' in 'iosxr_evpn_interface' resource and data source

- BREAKING CHANGE: Rename 'compress-level' to 'compress' in 'iosxr_interface' resource and data source
- BREAKING CHANGE: Rename 'ipv6_access_group_egress_acl1' to 'ipv6_access_group_egress_acl' in 'iosxr_interface' resource and data source
- Add 'timers_bgp_keepalive_zero', 'timers_bgp_holdtime_zero', 'timers_bgp_holdtime_minimum_acceptable_holdtime', 'timers_bgp_zero_zero', 'timers_bgp_zero_minimum_acceptable_holdtime' to 'iosxr_router_bgp' resource and data source
- BREAKING CHANGE: Rename 'timers_bgp_holdtime_holdtime' to 'timers_bgp_holdtime_number' in 'iosxr_router_bgp' resource and data source
- BREAKING CHANGE: Rename 'neighbor-address' to 'address' in 'iosxr_router_bgp' resource and data source
- Add 'remote_as' to 'iosxr_router_bgp' resource and data source
- BREAKING CHANGE: Rename 'advertisement_interval_time_in_seconds' to 'advertisement_interval_advertisement_interval_time_in_seconds' in 'iosxr_router_bgp' resource and data source
- BREAKING CHANGE: Rename 'advertisement_interval_time_in_milliseconds' to 'advertisement_interval_advertisement_interval_time_in_milliseconds' in 'iosxr_router_bgp' resource and data source
- BREAKING CHANGE: Remove 'bfd_fast_detect_inheritance_disable' in 'iosxr_router_bgp' resource and data source
- Add 'bfd_fast_detect_disable' to 'iosxr_router_bgp' resource and data source
- Add 'local_as_inheritance_disable', 'local_as_no_prepend', 'local_as_replace_as', 'local_as_dual_as' to 'iosxr_router_bgp' resource and data source
- Add 'bgp_bestpath_sr_policy_prefer', 'bgp_bestpath_sr_policy_force' to 'iosxr_router_bgp' resource and data source

- Add 'timers_bgp_keepalive_zero', 'timers_bgp_holdtime_zero', 'timers_bgp_holdtime_minimum_acceptable_holdtime', 'timers_bgp_zero_zero', 'timers_bgp_zero_minimum_acceptable_holdtime' to 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'neighbor-address' to 'address' in 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Remove 'bgp_bestpath_med_confed' in 'iosxr_router_bgp_vrf' resource and data source
- Add 'bgp_bestpath_sr_policy_prefer', 'bgp_bestpath_sr_policy_force' to 'iosxr_router_bgp_vrf' resource and data source
- BREAKING CHANGE: Rename 'advertisement_interval_time_in_seconds' to 'advertisement_interval_advertisement_interval_time_in_seconds' in 'iosxr_router_bgp_neighbor_group' resource and data source
- BREAKING CHANGE: Rename 'advertisement_interval_time_in_milliseconds' to 'advertisement_interval_advertisement_interval_time_in_milliseconds' in 'iosxr_router_bgp_neighbor_group' resource and data source
- BREAKING CHANGE: Remove 'bfd_fast_detect_inheritance_disable' in 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'bfd_fast_detect_disable' to 'iosxr_router_bgp_neighbor_group' resource and data source
- Add 'additional_paths_selection_route_policy', 'additional_paths_selection_disable' to 'iosxr_router_bgp_address_family' resource and data source
- BREAKING CHANGE: Rename 'neighbor-address' to 'address' in 'iosxr_router_bgp_vrf_neighbor_address_family' resource and data source
- Add 'additional_paths_selection_route_policy', 'additional_paths_selection_disable' to 'iosxr_router_bgp_vrf_address_family' resource and data source
- BREAKING CHANGE: Remove 'allocate_label_all_unlabeled_path' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'allocate_label_all', 'allocate_label_all_unlabeled_path', 'allocate_label_route_policy', 'allocate_label_route_policy_name', 'allocate_label_route_policy_unlabeled_path' to 'iosxr_router_bgp_address_family' resource and data source
- BREAKING CHANGE: Remove 'allocate_label_all_unlabeled_path' in 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'allocate_label_all', 'allocate_label_all_unlabeled_path', 'allocate_label_route_policy', 'allocate_label_route_policy_name', 'allocate_label_route_policy_unlabeled_path' to 'iosxr_router_bgp_vrf_address_family' resource and data source
- BREAKING CHANGE: Remove 'maximum_paths_ebgp_multipath', 'maximum_paths_ibgp_multipath', 'maximum_paths_eibgp_multipath' in 'iosxr_router_bgp_address_family' resource and data source
- Add 'maximum_paths_ebgp', 'maximum_paths_ebgp_number', 'maximum_paths_ebgp_selective', 'maximum_paths_ebgp_route_policy', 'maximum_paths_ibgp', 'maximum_paths_ibgp_number', 'maximum_paths_ibgp_unequal_cost', 'maximum_paths_ibgp_unequal_cost_deterministic', 'maximum_paths_ibgp_selective', 'maximum_paths_ibgp_route_policy', 'maximum_paths_eibgp', 'maximum_paths_eibgp_number', 'maximum_paths_eibgp_equal_cost', 'maximum_paths_eibgp_selective', 'maximum_paths_eibgp_route_policy', 'maximum_paths_unique_nexthop_check_disable' to 'iosxr_router_bgp_address_family' resource and data source
- BREAKING CHANGE: Remove 'maximum_paths_ebgp_multipath', 'maximum_paths_ibgp_multipath', 'maximum_paths_eibgp_multipath' in 'iosxr_router_bgp_vrf_address_family' resource and data source
- Add 'maximum_paths_ebgp', 'maximum_paths_ebgp_number', 'maximum_paths_ebgp_selective', 'maximum_paths_ebgp_route_policy', 'maximum_paths_ibgp', 'maximum_paths_ibgp_number', 'maximum_paths_ibgp_unequal_cost', 'maximum_paths_ibgp_unequal_cost_deterministic', 'maximum_paths_ibgp_selective', 'maximum_paths_ibgp_route_policy', 'maximum_paths_eibgp', 'maximum_paths_eibgp_number', 'maximum_paths_eibgp_equal_cost', 'maximum_paths_eibgp_selective', 'maximum_paths_eibgp_route_policy', 'maximum_paths_unique_nexthop_check_disable' to 'iosxr_router_bgp_vrf_address_family' resource and data source

## 0.5.3

- Revert workaround for issue related to interpreting dotted decimal AS number notation, as a fix has been implemented in recent XR versions, e.g. 24.2.2
- Handle state refresh when objects have been removed, [link](https://github.com/CiscoDevNet/terraform-provider-iosxr/issues/273)

## 0.5.2

- Implement workaround for issue related to interpreting dotted decimal AS number notation, [link](https://github.com/CiscoDevNet/terraform-provider-iosxr/issues/263)
- Add `use_af_group` attribute to `iosxr_router_bgp_neighbor_group` resource and data source

## 0.5.1

- Add `description`, `local_as`, `local_as_dual_as`, `local_as_no_prepend` and `local_as_replace_as` attributes to `iosxr_router_bgp_neighbor_group` resource and data source
- Add `bgp_router_id` and `use_neighbor_group` attributes to `iosxr_router_bgp_vrf` resource and data source

## 0.5.0

- Add `next_hop_self`, `soft_reconfiguration_inbound_always`, `send_community_ebgp`, `send_community_ebgp_inheritance_disable`, `maximum_prefix_limit`, `maximum_prefix_threshold`, `maximum_prefix_restart`, `maximum_prefix_discard_extra_paths`, `maximum_prefix_warning_only`, `default_originate_route_policy` and `default_originate_inheritance_disable` attributes to `iosxr_router_bgp_neighbor_address_family` resource and data source
- Add `timers_keepalive_interval`, `timers_holdtime` and `timers_minimum_acceptable_holdtime` attributes to `iosxr_router_bgp_neighbor_group` resource and data source
- Make `index_sid_index` and `absolute_sid_label` attribute of `iosxr_router_ospf_area_interface` resource optional
- Add `iosxr_router_static_vrf_ipv4_unicast` resource and data source
- Add `iosxr_router_static_vrf_ipv4_multicast` resource and data source
- Add `iosxr_router_static_vrf_ipv6_unicast` resource and data source
- Add `iosxr_router_static_vrf_ipv6_multicast` resource and data source
- Add `next_hop_self` to `iosxr_router_bgp_neighbor_group` resource and data source
- Add `next_hop_self` to `iosxr_router_bgp_vrf_neighbor_address_family` resource and data source
- Add `address_link_local_autoconfig` attribute to `iosxr_hsrp_interface_address_family_ipv6_group_v2` resource and data source
- BREAKING CHANGE: Rename `iosxr_router_hsrp_interface_address_family_ipv4_group_v1` resource and data source to `iosxr_router_hsrp_interface_ipv4_group_v1`
- BREAKING CHANGE: Rename `iosxr_router_hsrp_interface_address_family_ipv4_group_v2` resource and data source to `iosxr_router_hsrp_interface_ipv4_group_v2`
- BREAKING CHANGE: Rename `iosxr_router_hsrp_interface_address_family_ipv6_group_v2` resource and data source to `iosxr_router_hsrp_interface_ipv6_group_v2`
- BREAKING CHANGE: Rename `iosxr_router_vrrp_interface_address_family_ipv4` resource and data source to `iosxr_router_vrrp_interface_ipv4`
- BREAKING CHANGE: Rename `iosxr_router_vrrp_interface_address_family_ipv6` resource and data source to `iosxr_router_vrrp_interface_ipv6`

## 0.4.0

- BREAKING CHANGE: Refactor resource import functionality to use a comma separated list of key attribute values instead of a gNMI path
- Add support for empty YANG containers to the `iosxr_gnmi` resource using the `<EMPTY>` keyword

## 0.3.2

- Add `route_policy_in` and `route_policy_out` attributes to `iosxr_router_bgp_neighbor_group` resource and data source
- Add `fast_reroute_per_prefix_ti_lfa` and `fast_reroute_node_protecting_srlg_disjoint` attributes to `iosxr_router_ospf_area_interface` resource and data source
- Add `prefix_sid_strict_spf` and `prefix_sid_algorithm` attributes to `iosxr_router_ospf_area_interface` resource and data source
- Add `segment_routing_mpls` and `segment_routing_sr_prefer` attributes to `iosxr_router_ospf` resource and data source
- Add `v3_sha_encryption` and `v3_aes_encryption` attributes to `iosxr_snmp_server` resource and data source
- Add `fast_reroute_per_prefix` and `fast_reroute_per_prefix_ti_lfa` attributes to `iosxr_router_isis_interface_address_family` resource and data source
- Add `reuse_connection` provider attribute
- Add `redistribute_connected_route_policy` and `redistribute_static_route_policy` attributes to `iosxr_router_bgp_address_family` resource and data source
- Add `networks.route_policy` and `redistribute_isis.route_policy` attributes to `iosxr_router_bgp_address_family` resource and data source
- Add `additional_paths`, `allocate_label` and `advertise_best_external` attributes to `iosxr_router_bgp_vrf_address_family` resource and data source
- Add `redistribute_connected_route_policy` and `redistribute_static_route_policy` attributes to `iosxr_router_bgp_vrf_address_family` resource and data source
- Add `networks.route_policy` and `redistribute_isis.route_policy` attributes to `iosxr_router_bgp_vrf_address_family` resource and data source
- Add `route_policy_in` and `route_policy_out` attributes to `iosxr_router_bgp_neighbor_address_family` resource and data source

## 0.3.1

- Make various BGP neighbor attributes optional
- Make `set_overload_bit_on_startup_advertise_as_overloaded_time_to_advertise` and `on_startup_advertise_as_overloaded_time_to_advertise` attributes of `iosxr_router_isis` resource optional
- Make `make_before_break_delay` attribute of `iosxr_mpls_ldp` resource optional
- Make various `iosxr_key_chain` resource attributes optional

## 0.3.0

- Add `ipv4_verify_unicast_source_reachable_via` and `ipv4_access_group` attributes to `iosxr_interface` resource and data source
- Add `ipv6_verify_unicast_source_reachable_via` and `ipv6_access_group` attributes to `iosxr_interface` resource and data source
- BREAKING CHANGE: Rename traps related attributes of `iosxr_snmp_server` resource and data source to `traps_*`

## 0.2.6

- Add `auto_cost` attributes to `iosxr_router_ospf` and `iosxr_router_ospf_vrf` resources and data sources
- Add support for leaf-lists to `iosxr_gnmi` resource
- Add `port` and `operator` attributes to `iosxr_logging_vrf` resource and data source
- Add `iosxr_router_hsrp_interface_address_family_ipv6_group_v2` resource and data source
- Add `route_reflector_client` attribute to `iosxr_router_bgp_neighbor_address_family` and `iosxr_router_bgp_neighbor_group` resource and data source
- Add `communities` attributes to `iosxr_snmp_server` resource and data source
- When removing attributes from a resource (or setting them to `null`) which were previously set, the corresponding configuration will be removed from the device

## 0.2.5

- Add `ipv4_access_list` and `ipv6_access_list` attributes to `iosxr_ssh` resource and data source
- Add `iosxr_router_hsrp_interface` resource and data source
- Add `iosxr_router_hsrp_interface_address_family_ipv4_group_v1` resource and data source
- Add `iosxr_router_hsrp_interface_address_family_ipv4_group_v2` resource and data source

## 0.2.4

- Add `record_ipv4`, `record_ipv6`, `record_mpls` and `sflow_options` attributes to `iosxr_flow_monitor_map` resource and data source
- Add `set_overload_bit` attributes to `iosxr_router_isis` resource and data source
- Add `metric` attributes to `iosxr_router_isis_interface_address_family` resource and data source
- Add `nexthop_trigger_delay_critical` and `nexthop_trigger_delay_non_critical` attributes to `iosxr_router_bgp_address_family` resource and data source
- Add `advertisement_interval` attributes to `iosxr_router_bgp`, `iosxr_router_bgp_vrf` and `iosxr_router_bgp_neighbor_group` resources and data sources
- Add `load_balancing_flow_src_dst_mac` and `load_balancing_flow_src_dst_ip` attributes to `iosxr_l2vpn` resource and data source
- Add `iosxr_l2vpn_pw_class` resource and data source
- Add `igp_sync_delay` and `label_local_allocate` attributes to `iosxr_mpls_ldp` resource and data source

## 0.2.3

- Add `timers_bgp_minimum_acceptable_holdtime` attribute to `iosxr_router_bgp` resource and data source
- Add `iosxr_flow_sampler_map` resource and data source
- Add `iosxr_flow_monitor_map` resource and data source
- Add `iosxr_ntp` resource and data source
- Add `iosxr_bfd` resource and data source
- Add `iosxr_flow_exporter_map` resource and data source
- Add `bgp_bestpath` attributes to `iosxr_router_bgp` resource and data source
- Add `flow_ipv4` and `flow_ipv6` attributes to `iosxr_interface` resource and data source
- BREAKING CHANGE: Remove `neighbor_groups` attributes from `iosxr_router_bgp` resource and data source
- Add `bfd_fast_detect` attributes to `iosxr_router_bgp` resource and data source
- Add `bfd_multiplier` and `bfd_fast_detect` attributes to `iosxr_router_bgp_neighbor_group` resource and data source
- Add `bfd_fast_detect` attributes to `iosxr_router_bgp_vrf` resource and data source

## 0.2.2

- Make `icmp_error_interval_interval_time` attribute of `iosxr_ipv6` resource optional
- Make `timers_bgp_keepalive_interval` attribute of `iosxr_router_bgp_vrf` resource optional
- Make `timers_bgp_holdtime` attribute of `iosxr_router_bgp_vrf` resource optional
- Make `timers_bgp_keepalive_interval` attribute of `iosxr_router_bgp` resource optional
- Make `timers_bgp_holdtime` attribute of `iosxr_router_bgp` resource optional
- Make `lsp_password_keychain` attribute of `iosxr_router_isis` resource optional
- BREAKING CHANGE: Rename `icmp_error_interval_interval_time` attribute of `iosxr_ipv6` resource and data source to `icmp_error_interval`
- Add `iosxr_telnet` resource and data source
- Add `iosxr_tag_set` resource and data source
- Add `iosxr_error_disable_recovery` resource and data source
- Add `iosxr_extcommunity_rt_set` resource and data source
- Add `iosxr_extcommunity_soo_set` resource and data source
- Add `iosxr_fpd` resource and data source
- Add `iosxr_extcommunity_cost_set` resource and data source
- Add `iosxr_rd_set` resource and data source
- Add `contact` and `location` attributes to `iosxr_snmp_server` resource and data source
- Add `police_conform_action_transmit`, `police_conform_action_drop`, `police_exceed_action_transmit`, `police_exceed_action_drop`, `police_violate_action_transmit`, `police_violate_action_drop` attributes to `iosxr_policy_map_qos` resource and data source

## 0.2.1

- Add `bundle` attributes to `iosxr_interface` resource and data source
- Make `dampening_decay_half_life_value` attribute of `iosxr_interface` resource optional
- Do not configure `ipv6_link_local_zone` attribute of `iosxr_interface` resource by default

## 0.2.0

- Introduce more granular controls around resource delete operations (`delete_mode`)
- BREAKING CHANGE: Rename `iosxr_router_static` resource and data source to `iosxr_router_static_ipv4_unicast`
- BREAKING CHANGE: Remove `iosxr_oc_system_config` resource and data source
- Validate if referenced `device` exists in provider configuration
- Add `vrfs`, `permanent`, `track` and `metric` attributes to `iosxr_router_static_ipv4_unicast` resource and data source
- Add `iosxr_router_static_ipv4_multicast` resource and data source
- Add `iosxr_router_static_ipv6_unicast` resource and data source
- Add `iosxr_router_static_ipv6_multicast` resource and data source
- Add `iosxr_lldp` resource and data source
- Add `iosxr_community_set` resource and data source
- Add `iosxr_domain` resource and data source
- Add `iosxr_domain_vrf` resource and data source
- Add `iosxr_service_timestamps` resource and data source
- Add `iosxr_lacp` resource and data source
- Add `iosxr_as_path_set` resource and data source
- Add `iosxr_esi_set` resource and data source
- Add `iosxr_ipv6` resource and data source
- Add `iosxr_router_vrrp_interface` resource and data source
- Add `iosxr_router_vrrp_interface_address_family_ipv4` resource and data source
- Add `iosxr_router_vrrp_interface_address_family_ipv6` resource and data source

## 0.1.11

- Migrate provider to `CiscoDevNet` namespace

## 0.1.10

- Add `iosxr_banner` resource and data source
- Add `iosxr_cdp` resource and data source
- Add `iosxr_extcommunity_opaque_set` resource and data source
- Add `iosxr_segment_routing_v6` resource and data source
- Add `iosxr_segment_routing_te` resource and data source
- Add `iosxr_segment_routing_te_policy_candidate_path` resource and data source
- Add `iosxr_router_bgp_neighbor_address_family` resource and data source
- Add `iosxr_evpn_segment_routing_srv6_evi` resource and data source
- Add `segment_routing_srv6_evis` attribute to `iosxr_l2vpn_bridge_group_bridge_domain` resource and data source
- Add `neighbor_evpn_evi_segment_routing_services` attribute to `iosxr_l2vpn_xconnect_group_p2p` resource and data source
- Add `nexthop_validation_color_extcomm_sr_policy`, `nexthop_validation_color_extcomm_disable`, `bfd_minimum_interval`, `nsr_disable`, `bgp_redistribute_internal` and `segment_routing_srv6_locator` attributes to `iosxr_router_bgp` resource and data source
- Add `segment_routing_srv6_locator` and `segment_routing_srv6_alloc_mode_per_vrf` attributes to `iosxr_router_bgp_vrf_address_family` resource and data source
- Add `redistribute_isis` and `segment_routing_srv6_locators` attributes to `iosxr_router_isis_address_family` resource and data source
- Add `bfd_fast_detect_ipv6` attribute to `iosxr_router_isis_interface` resource and data source
- Add `bfd_minimum_interval`, `next_hop_self_inheritance_disable`, `route_reflector_client_inheritance_disable`, `bfd_fast_detect` attribute to `iosxr_router_bgp_neighbor_group` resource and data source
- Add `iosxr_segment_routing` resource and data source
- Add `iosxr_pce` resource and data source
- Add `iosxr_class_map_qos` resource and data source
- Add `iosxr_policy_map_qos` resource and data source
- Add `iosxr_ipv4_access_list` resource and data source
- Add `iosxr_ipv4_access_list_options` resource and data source
- Add `iosxr_ipv4_prefix_list` resource and data source
- Add `iosxr_ipv6_access_list` resource and data source
- Add `iosxr_ipv6_access_list_options` resource and data source
- Add `iosxr_ipv6_prefix_list` resource and data source
- Add `capabilities_sac_ipv4_disable`, `capabilities_sac_ipv6_disable`, `capabilities_sac_fec128_disable`, `capabilities_sac_fec129_disable`, `mldp_logging_notifications`, `mldp_address_families` and `session_protection` attributes to `iosxr_mpls_ldp` resource and data source
- Add `logging_pcep_peer_status`, `logging_policy_status`, `pcc_report_all`, `pcc_source_address`, `pcc_delegation_timeout`, `pcc_dead_timer`, `pcc_initiated_state`, `pcc_initiated_orphan`, `pce_peers` `dynamic_anycast_sid_inclusion` and `dynamic_metric_type` attributes to `iosxr_segment_routing_te` resource and data source

## 0.1.9

- Add `iosxr_router_bgp_vrf_neighbor_address_family` resource and data source
- Add `iosxr_router_bgp_neighbor_group` resource and data source
- Add `iosxr_router_static` resource and data source
- Add `iosxr_key_chain` resource and data source
- Add attributes to `iosxr_interface` resource and data source
- Add attributes to `iosxr_router_bgp` resource and data source
- Add attributes to `iosxr_router_bgp_address_family` resource and data source
- Add attributes to `iosxr_router_bgp_vrf` resource and data source
- Add attributes to `iosxr_router_isis` resource and data source
- Add attributes to `iosxr_router_isis_address_family` resource and data source
- Add attributes to `iosxr_router_isis_interface_address_family` resource and data source

## 0.1.8

- Fix incompatibility with gNMI and IOS-XR 7.6
- Remove `iosxr_segment_routing` resource and data source due to unified model being deprecated

## 0.1.7

- Add `iosxr_mpls_traffic_eng` resource and data source
- Add `iosxr_mpls_oam` resource and data source
- Add `iosxr_segment_routing` resource and data source
- Add `iosxr_logging` resource and data source
- Add `iosxr_logging_source_interface` resource and data source
- Add `iosxr_logging_vrf` resource and data source
- Add `iosxr_snmp_server` resource and data source
- Add `iosxr_snmp_server_mib` resource and data source
- Add `iosxr_snmp_server_view` resource and data source
- Add `iosxr_snmp_server_vrf_host` resource and data source
- Add `verify_certificate` provider attribute
- Add `tls` provider attribute
- Add `certificate` provider attribute
- Add `key` provider attribute
- Add `ca_certificate` provider attribute
- BREAKING CHANGE: Use TLS by default

## 0.1.6

- BREAKING CHANGE: Remove `address_families` from `iosxr_router_isis` resource and data source
- Add `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `interfaces` from `iosxr_router_isis` resource and data source
- Add `iosxr_router_isis_interface` resource and data source

## 0.1.5

- Add `iosxr_prefix_set` resource and data source
- Add `iosxr_route_policy` resource and data source
- Add `address_family_ipv4_unicast_import_route_policy` attribute to `iosxr_vrf` resource
- Add `address_family_ipv4_unicast_export_route_policy` attribute to `iosxr_vrf` resource
- Add `address_family_ipv6_unicast_import_route_policy` attribute to `iosxr_vrf` resource
- Add `address_family_ipv6_unicast_export_route_policy` attribute to `iosxr_vrf` resource
- Add `l2transport_encapsulation_dot1q_vlan_id` attribute to `iosxr_interface` resource
- Add `l2transport_encapsulation_dot1q_second_dot1q` attribute to `iosxr_interface` resource
- Add `rewrite_ingress_tag_one` attribute to `iosxr_interface` resource
- Add `rewrite_ingress_tag_two` attribute to `iosxr_interface` resource
- Add `encapsulation_dot1q_vlan_id` attribute to `iosxr_interface` resource
- Add `load_interval` attribute to `iosxr_interface` resource
- Add `iosxr_evpn_evi` resource and data source
- Add `iosxr_evpn_group` resource and data source
- Add `iosxr_evpn_interface` resource and data source
- Add `iosxr_evpn` resource and data source
- Add `iosxr_l2vpn_bridge_group` resource and data source
- Add `iosxr_l2vpn_bridge_group_bridge_domain` resource and data source
- Add `evpn_target_neighbors` attribute to `iosxr_l2vpn_xconnect_group_p2p` resource
- Add `evpn_service_neighbors` attribute to `iosxr_l2vpn_xconnect_group_p2p` resource

## 0.1.4

- Add `iosxr_ssh` resource and data source
- Allow concurrent changes across different devices

## 0.1.3

- Add support for IOS-XR 7.8.1+

## 0.1.2

- Update dependencies and go version

## 0.1.1

- Add delete attribute to `iosxr_gnmi` resource
- Allow provider config without host attribute in case `devices` attribute is being used
- Enhance `iosxr_gnmi` resource to support nested attributes (within YANG containers) using "/" as separator
- Enhance `iosxr_gnmi` resource to support nested YANG lists
- BREAKING CHANGE: remove `iosxr_vrf_route_target_two_byte_as_format` resource and data source
- BREAKING CHANGE: remove `iosxr_vrf_route_target_four_byte_as_format` resource and data source
- BREAKING CHANGE: remove `iosxr_vrf_route_target_ip_address_format` resource and data source
- Add route target attributes to `iosxr_vrf` resource and data source
- BREAKING CHANGE: remove `iosxr_interface_ipv4` resource and data source
- BREAKING CHANGE: remove `iosxr_interface_ipv6` resource and data source
- BREAKING CHANGE: remove `iosxr_interface_ipv6_address` resource and data source
- Add ipv4 and ipv6 attributes to `iosxr_interface` resource and data source
- BREAKING CHANGE: remove `iosxr_mpls_ldp_address_family` resource and data source
- BREAKING CHANGE: remove `iosxr_mpls_ldp_interface` resource and data source
- Add `address_family` and interface attributes to `iosxr_mpls_ldp` resource and data source
- BREAKING CHANGE: remove `iosxr_l2vpn_xconnect_group` resource and data source
- Add `xconnect_group` attributes to `iosxr_l2vpn resource` and data source
- BREAKING CHANGE: remove `iosxr_l2vpn_xconnect_group_p2p_interface` resource and data source
- BREAKING CHANGE: remove `iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv4` resource and data source
- BREAKING CHANGE: remove `iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6` resource and data source
- Add interface and neighbor attributes to `iosxr_l2vpn_xconnect_group_p2p` resource and data source
- BREAKING CHANGE: remove `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: remove `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: remove `iosxr_router_isis_net` resource and data source
- Add address family, interface and net attributes to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_area` resource and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_redistribute_bgp` resource and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_redistribute_isis` resource and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_redistribute_ospf` resource and data source
- Add area and redistribute attributes to `iosxr_router_ospf resource` and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_vrf_area` resource and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_vrf_redistribute_bgp` resource and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_vrf_redistribute_isis` resource and data source
- BREAKING CHANGE: remove `iosxr_router_ospf_vrf_redistribute_ospf` resource and data source
- Add area and redistribute attributes to `iosxr_router_ospf_vrf` resource and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_neighbor` resource and data source
- Add neighbor attributes to `iosxr_router_bgp resource` and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_vrf_neighbor` resource and data source
- Add neighbor attributes to `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_address_family_aggregate_address` resource and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_address_family_network` resource and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_address_family_redistribute_isis` resource and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_address_family_redistribute_ospf` resource and data source
- Add aggregate address, network and redistribute attributes to `iosxr_router_bgp_address_family` resource and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_vrf_address_family_aggregate_address` resource and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_vrf_address_family_network resource` and data source
- BREAKING CHANGE: remove `iosxr_router_bgp_vrf_address_family_redistribute_ospf` resource and data source
- Add aggregate address, network and redistribute attributes to `iosxr_router_bgp_vrf_address_family` resource and data source

## 0.1.0

- Initial release

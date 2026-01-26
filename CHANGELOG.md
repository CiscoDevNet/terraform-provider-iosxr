## 0.7.0

- BREAKING CHANGE: Rename `purge_transmit_strict_strict_value` to `purge_transmit_strict_value` in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Rename `distribute_list_acl` to `distribute_list_in_acl` in `iosxr_router_ospf_area_interface`, `iosxr_router_ospf_area`, `iosxr_router_ospf_vrf_area_interface`, `iosxr_router_ospf_vrf_area` resource and data source
- BREAKING CHANGE: Rename `distribute_list_route_policy` to `distribute_list_in_route_policy` in `iosxr_router_ospf_area_interface`, `iosxr_router_ospf_area`, `iosxr_router_ospf_vrf_area_interface`, `iosxr_router_ospf_vrf_area` resource and data source
- BREAKING CHANGE: Rename `algorithm-number` to `number` in `router_ospf_area_interface` resource and data source
- BREAKING CHANGE: Rename `timers_bgp_keepalive_zero` to `timers_bgp_holddown_zero` in `iosxr_router_bgp` and `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `timers_bgp_keepalive_zero_holdtime_zero` to `timers_bgp_holddown_zero_minimum_acceptable_zero` in `iosxr_router_bgp` and `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `timers_bgp_keepalive_zero_minimum_acceptable_holdtime` to `timers_bgp_holddown_zero_minimum_acceptable_holdtime` in `iosxr_router_bgp` and `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `timers_keepalive_zero` to `timers_holddown_zero` in `iosxr_router_bgp`, `iosxr_router_bgp_vrf`, `iosxr_router_bgp_neighbor_group` and `iosxr_router_bgp_session_group` resource and data source
- BREAKING CHANGE: Rename `timers_keepalive_zero_holdtime_zero` to `timers_holddown_zero_minimum_acceptable_zero` in `iosxr_router_bgp`, `iosxr_router_bgp_vrf`, `iosxr_router_bgp_neighbor_group` and `iosxr_router_bgp_session_group` resource and data source
- BREAKING CHANGE: Rename `timers_keepalive_zero_minimum_acceptable_holdtime` to `timers_holddown_zero_minimum_acceptable_holdtime` in `iosxr_router_bgp`, `iosxr_router_bgp_vrf`, `iosxr_router_bgp_neighbor_group` and `iosxr_router_bgp_session_group` resource and data source
- BREAKING CHANGE: Decompose `iosxr_interface` into individual interface type resource and data source
  - Interface types:
      `iosxr_interface_ethernet` `iosxr_interface_ethernet_subinterface` `iosxr_interface_bundle_ether` `iosxr_interface_bundle_ether_subinterface` `iosxr_interface_bvi` `iosxr_interface_loopback` `iosxr_interface_tunnel_ip` `iosxr_interface_tunnel_te`
- BREAKING CHANGE: Consolidate `iosxr_logging_source_interface` into `iosxr_logging` resource and data source
- BREAKING CHANGE: Consolidate `iosxr_evpn_group` into `iosxr_evpn` resource and data source
- BREAKING CHANGE: Consolidate `iosxr_segment_routing_te_policy_candidate_path` into `iosxr_segment_routing_te_policy` resource and data source
- BREAKING CHANGE: Consolidate `iosxr_snmp_server_view` into `iosxr_snmp_server` resource and data source
- BREAKING CHANGE: Consolidate `iosxr_l2vpn_xconnect_group_p2p` into `iosxr_l2vpn_xconnect_group` resource and data source
- BREAKING CHANGE: Consolidate `iosxr_l2vpn_bridge_group` into `iosxr_l2vpn_bridge_group_bridge_domain` resource and data source
- BREAKING CHANGE: Decompose `iosxr_segment_routing_te` on-demand-colors into `iosxr_segment_routing_te_on_demand_color` resource and data source
- BREAKING CHANGE: Decompose `iosxr_mpls_ldp` into `iosxr_mpls_ldp`, `iosxr_mpls_ldp_address_family`, `iosxr_mpls_ldp_interface`, `iosxr_mpls_ldp_mldp`, `iosxr_mpls_ldp_vrf` resource and data source
- BREAKING CHANGE: Rename `fast_reroute_per_prefix_tiebreaker_node_protecting` to `fast_reroute_per_prefix_tiebreaker_node_protecting_index` in `iosxr_router_ospf_area_interface` resource and data source
- BREAKING CHANGE: Rename `fast_reroute_per_prefix_tiebreaker_srlg_disjoint` to `fast_reroute_per_prefix_tiebreaker_srlg_disjoint_index` in `iosxr_router_ospf_area_interface` resource and data source
- BREAKING CHANGE: Rename `fast_reroute_per_prefix_ti_lfa` to `fast_reroute_per_prefix_ti_lfa_enable` in `iosxr_router_ospf_area_interface` resource and data source
- BREAKING CHANGE: Rename `dampening_decay_half_life_value` to `dampening_decay_half_life` in `iosxr_interface` resource and data source
- BREAKING CHANGE: Rename `local_as_replace_as` to `local_as_no_prepend_replace_as` in `iosxr_router_bgp_neighbor_group` resource and data source
- BREAKING CHANGE: Rename `local_as_dual_as` to `local_as_no_prepend_replace_as_dual_as` in `iosxr_router_bgp_neighbor_group` resource and data source
- BREAKING CHANGE: Rename `buffered_logging_buffer_size` to `buffered_size` in `iosxr_logging` resource and data source
- BREAKING CHANGE: Rename `encapsulation_mpls_load_balancing_flow_label_code_one7` to `encapsulation_mpls_load_balancing_flow_label_code_17` in `iosxr_l2vpn_pw_class` resource and data source
- BREAKING CHANGE: Rename `encapsulation_mpls_load_balancing_flow_label_code_one7_disable` to `encapsulation_mpls_load_balancing_flow_label_code_17_disable` in `iosxr_l2vpn_pw_class` resource and data source
- BREAKING CHANGE: Rename `encapsulation_mpls_transport_mode_passthrough` to `encapsulation_mpls_transport_mode_vlan_passthrough` in `iosxr_l2vpn_pw_class` resource and data source
- BREAKING CHANGE: Rename `bgp_rd` attributes for consistency with other definitions in `iosxr_evpn_evi` resource and data source
- BREAKING CHANGE: Rename `pcc_source_address` to `pcc_source_address_ipv4` in `iosxr_segment_routing_te` resource and data source
- BREAKING CHANGE: Rename `pce_peers` to `pce_peers_ipv4` in `iosxr_segment_routing_te` resource and data source
- BREAKING CHANGE: Rename `unencrypted_strings` to `traps_unencrypted_strings` in `iosxr_snmp_server_vrf_host` resource and data source
- BREAKING CHANGE: Rename `hello_keychain_send_only` to `keychain_send_only` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hello_keychain_name` to `keychain_name` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hello_password_hmac_md5_send_only` to `hmac_md5_send_only` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hello_password_hmac_md5_encrypted` to `hmac_md5_encrypted` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hello_password_text_send_only` to `text_send_only` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hello_password_text_encrypted` to `text_encrypted` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `rpl-tag-set` to `rpl` in `iosxr_tag_set` resource and data source
- BREAKING CHANGE: Rename `lsp_password_levels` to `lsp_password_accept_levels` in `iosxr_router_isis` resource and data source

- Add many new attributes across all resources and data sources
- Add `iosxr_ethernet_sla` resource and data source
- Add `iosxr_ethernet_cfm` resource and data source
- Add `iosxr_cef_load_balancing_8000` resource and data source
- Add `iosxr_generic_interface_list` resource and data source
- Add `iosxr_srlg` resource and data source
- Add `iosxr_evpn_route_sync_stitching_evi` resource and data source
- Add `iosxr_evpn_route_sync_evi` resource and data source
- Add `iosxr_evpn_stitching_vni` resource and data source
- Add `iosxr_evpn_vni` resource and data source
- Add `iosxr_evpn_segment_routing_srv6_stitching_evi` resource and data source
- Add `iosxr_evpn_stitching_evi` resource and data source
- Add `iosxr_router_bgp_session_group` resource and data source
- Add `iosxr_router_bgp_af_group` resource and data source
- Add `iosxr_bmp_server` resource and data source
- Add `iosxr_hw_module_profile_8000` resource and data source
- Add `iosxr_rsvp` resource and data source
- Add `iosxr_rsvp_interface` resource and data source
- Add `iosxr_class_map_traffic` resource and data source
- Add `iosxr_policy_map_pbr` resource and data source
- Add `iosxr_aaa` resource and data source
- Add `iosxr_aaa_authentication` resource and data source
- Add `iosxr_aaa_authorization` resource and data source
- Add `iosxr_aaa_accounting` resource and data source
- Add `iosxr_radius_source_interface` resource and data source
- Add `iosxr_radius_server` resource and data source
- Add `iosxr_tacacs_source_interface` resource and data source
- Add `iosxr_tacacs_server` resource and data source
- Add `iosxr_segment_routing_te_on_demand_color` resource and data source
- Add `iosxr_router_ospf_vrf_area` resource and data source
- Add `iosxr_router_ospf_area` resource and data source
- Add `iosxr_controller_optics` resource and data source
- Add `iosxr_etag_set` resource and data source
- Add `iosxr_policy_global_set` resource and data source
- Add `iosxr_ospf_area_set` resource and data source
- Add `iosxr_mac_set` resource and data source
- Add `iosxr_large_community_set` resource and data source
- Add `iosxr_extcommunity_seg_nh_set` resource and data source
- Add `iosxr_extcommunity_evpn_link_bandwidth_set` resource and data source
- Add `iosxr_extcommunity_bandwidth_set` resource and data source
- Add `iosxr_frequency_synchronization` resource and data source
- Add `iosxr_ptp` resource and data source
- Add `iosxr_ptp_profile` resource and data source
- Add `iosxr_macsec` resource and data source
- Add `iosxr_macsec_policy` resource and data source
- Add `iosxr_crypto` resource and data source
- Add `iosxr_hw_module_profile` resource and data source
- Add `iosxr_lpts_punt_police` resource and data source
- Add `iosxr_monitor_session` resource and data source
- Add `iosxr_icmp` resource and data source
- Add `iosxr_ipsa_responder` resource and data source
- Add `iosxr_vty_pool` resource and data source
- Add `iosxr_telemetry_model_driven` resource and data source
- Add `iosxr_l2vpn_bridge_group_bridge_domain_neighbor` resource and data source
- Add `iosxr_l2vpn_bridge_group_bridge_access_vfi` resource and data source
- Add `iosxr_l2vpn_bridge_group_bridge_domain_vfi` resource and data source
- Add `iosxr_control_plane` resource and data source
- Add `iosxr_segment_routing_mapping` resource and data source
- Add `iosxr_call_home` resource and data source
- Add `iosxr_tcp` resource and data source
- Add `iosxr_line_console` resource and data source
- Add `iosxr_line_default` resource and data source
- Add `iosxr_line_template` resource and data source
- Add `iosxr_track` resource and data source
- Add `iosxr_ipsla_responder` resource and data source
- Add `iosxr_ipsla` resource and data source
- Add `iosxr_router_mld` resource and data source
- Add `iosxr_router_mld_interface` resource and data source
- Add `iosxr_router_mld_vrf` resource and data source
- Add `iosxr_router_mld_vrf_inteface` resource and data source
- Add `iosxr_router_igmp` resource and data source
- Add `iosxr_router_igmp_interface` resource and data source
- Add `iosxr_router_igmp_vrf` resource and data source
- Add `iosxr_router_igmp_vrf_inteface` resource and data source
- Add `iosxr_router_pim_ipv4` resource and data source
- Add `iosxr_router_pim_ipv6` resource and data source
- Add `iosxr_router_pim_vrf_ipv4` resource and data source
- Add `iosxr_router_pim_vrf_ipv6` resource and data source
- Add `iosxr_ftp` resource and data source
- Add `iosxr_tftp_server` resource and data source
- Add `iosxr_tftp_client` resource and data source
- Add `iosxr_cli_alias` resource and data source
- Add `iosxr_linux_networking` resource and data source
- Add `iosxr_cef` resource and data source
- Add `iosxr_cef_pbts_forward_class` resource and data source
- Add `iosxr_hw_module_shutdown` resource and data source
- Add `iosxr_performance_measurement` resource and data source
- Add `iosxr_performance_measurement_interface` resource and data source
- Add `iosxr_performance_measurement_delay_profile` resource and data source
- Add `iosxr_performance_measurement_liveness_profile` resource and data source
- Add `iosxr_performance_measurement_endpoint_ipv6` resource and data source
- Add `iosxr_performance_measurement_endpoint_ipv4` resource and data source
- Add `iosxr_xml_agent` resource and data source
- Add `iosxr_netconf_agent_tty` resource and data source
- Add `iosxr_netconf_yang_agent` resource and data source
- Add `iosxr_tpa` resource and data source
- Add `iosxr_cli` resource
- Make `maximum_prefix_limit` attribute of `iosxr_bgp_neighbor_address_family` resource optional
- Make `maximum_prefix_threshold` attribute of `iosxr_bgp_neighbor_address_family` resource optional
- Change `iosxr_ssh` default `delete_mode` to attributes
- Fix xpath for `prefix_list_name` attribute in `spf_prefix_priority_medium_levels` list for `iosxr_router_isis_address_family` resource and data source
- Fix xpath for `prefix_list_name` attribute in `spf_prefix_priority_high_levels` list for `iosxr_router_isis_address_family` resource and data source
- Fix xpath for `metric_default` attribute in `metric_levels` list for `iosxr_router_isis_interface_address_family` resource and data source
- Fix xpath for `metric_maximum` attribute in `metric_levels` list for `iosxr_router_isis_interface_address_family` resource and data source

## 0.6.0

- Major changes to existing datasources/resources for compatability with IOS-XR Version 24.4.2
- BREAKING CHANGE: Rename `keychain_send_only` to `hello_keychain_send_only` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `keychain_name` to `hello_keychain_name` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hmac_md5_send_only` to `hello_password_hmac_md5_send_only` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hmac_md5_encrypted` to `hello_password_hmac_md5_encrypted` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `text_send_only` to `hello_password_text_send_only` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `text_encrypted` to `hello_password_text_encrypted` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `hello_password_keychain` to `hello_password_keychain_name` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `local_as_dual_as` to `local_as_no_prepend_replace_as_dual_as` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `local_as_replace_as` to `local_as_no_prepend_replace_as` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `timers_bgp_keepalive_zero` to `timers_bgp_keepalive_zero_holdtime_zero` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `timers_msec2` to `timers_msec_holdtime`, `timers_hold_time` to `timers_seconds`, `timers_hold_time2` to `timers_seconds_holdtime` in `iosxr_router_hsrp_interface_ipv4_group_v1` resource and data source
- BREAKING CHANGE: Rename `timers_msec2` to `timers_msec_holdtime`, `timers_hold_time` to `timers_seconds`, `timers_hold_time2` to `timers_seconds_holdtime` in `iosxr_router_hsrp_interface_ipv4_group_v2` resource and data source
- BREAKING CHANGE: Rename `timers_msec2` to `timers_msec_holdtime`, `timers_hold_time` to `timers_seconds`, `timers_hold_time2` to `timers_seconds_holdtime` in `iosxr_router_hsrp_interface_ipv6_group_v2` resource and data source
- BREAKING CHANGE: Rename `trap_source_both` to `trap_source` in `iosxr_snmp_server` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast` to `ipv4_unicast` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_import_route_policy` to `ipv4_unicast_import_route_policy` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_export_route_policy` to `ipv4_unicast_export_route_policy` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_multicast` to `ipv4_multicast` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_flowspec` to `ipv4_flowspec` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast` to `ipv6_unicast` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_import_route_policy` to `ipv6_unicast_import_route_policy` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_export_route_policy` to `ipv6_unicast_export_route_policy` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_multicast` to `ipv6_multicast` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_flowspec` to `ipv6_flowspec` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_import_route_target_two_byte_as_format` to `ipv4_unicast_import_route_target_two_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_import_route_target_four_byte_as_format` to `ipv4_unicast_import_route_target_four_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_import_route_target_ip_address_format` to `ipv4_unicast_import_route_target_ip_address_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_export_route_target_two_byte_as_format` to `ipv4_unicast_export_route_target_two_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_export_route_target_four_byte_as_format` to `ipv4_unicast_export_route_target_four_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv4_unicast_export_route_target_ip_address_format` to `ipv4_unicast_export_route_target_ip_address_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_import_route_target_two_byte_as_format` to `ipv6_unicast_import_route_target_two_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_import_route_target_four_byte_as_format` to `ipv6_unicast_import_route_target_four_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_import_route_target_ip_address_format` to `ipv6_unicast_import_route_target_ip_address_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_export_route_target_two_byte_as_format` to `ipv6_unicast_export_route_target_two_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_export_route_target_four_byte_as_format` to `ipv6_unicast_export_route_target_four_byte_as_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Rename `address_family_ipv6_unicast_export_route_target_ip_address_format` to `ipv6_unicast_export_route_target_ip_address_format` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Update YANG path to `Cisco-IOS-XR-segment-routing-ms-cfg:/sr/srv6` in `iosxr_segment_routing_v6` resource and data source
- BREAKING CHANGE: Rename `name` to `locator_name` in `iosxr_segment_routing_v6` resource and data source
- BREAKING CHANGE: Remove `locator_enable` in `iosxr_segment_routing_v6` resource and data source
- BREAKING CHANGE: Update YANG path to `Cisco-IOS-XR-um-segment-routing-cfg:/segment-routing` in `iosxr_segment_routing` resource and data source
- BREAKING CHANGE: Remove `rd_two_byte_as_number`, `rd_two_byte_as_index`, `rd_four_byte_as_number`, `rd_four_byte_as_index`, `rd_ipv4_address`, `rd_ipv4_address_index` in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Remove `address_family_ipv4_unicast_import_route_target_two_byte_as_format`, `address_family_ipv4_unicast_import_route_target_four_byte_as_format`, `address_family_ipv4_unicast_import_route_target_ip_address_format`, `address_family_ipv4_unicast_export_route_target_two_byte_as_format`, `address_family_ipv4_unicast_export_route_target_four_byte_as_format`, `address_family_ipv4_unicast_export_route_target_ip_address_format`, `address_family_ipv6_unicast_import_route_target_two_byte_as_format`, `address_family_ipv6_unicast_import_route_target_four_byte_as_format`, `address_family_ipv6_unicast_import_route_target_ip_address_format`, `address_family_ipv6_unicast_export_route_target_two_byte_as_format`, `address_family_ipv6_unicast_export_route_target_four_byte_as_format`, `address_family_ipv6_unicast_export_route_target_ip_address_format` (type: List) in `iosxr_vrf` resource and data source
- BREAKING CHANGE: Remove `traps_bgp_cbgp2_updown`, `traps_bgp_bgp4_mib_updown` in `iosxr_snmp_server` resource and data source
- BREAKING CHANGE: Remove `traps_isis_all`, `traps_isis_database_overload`, `traps_isis_manual_address_drops`, `traps_isis_corrupted_lsp_detected`, `traps_isis_attempt_to_exceed_max_sequence`, `traps_isis_id_len_mismatch`, `traps_isis_max_area_addresses_mismatch`, `traps_isis_own_lsp_purge`, `traps_isis_sequence_number_skip`, `traps_isis_authentication_type_failure`, `traps_isis_authentication_failure`, `traps_isis_version_skew`, `traps_isis_area_mismatch`, `traps_isis_rejected_adjacency`, `traps_isis_lsp_too_large_to_propagate`, `traps_isis_orig_lsp_buff_size_mismatch`, `traps_isis_protocols_supported_mismatch`, `traps_isis_adjacency_change`, `traps_isis_lsp_error_detected` in `iosxr_snmp_server` resource and data source
- BREAKING CHANGE: Remove `metric`, `maximum` in `iosxr_router_isis_interface_address_family` resource and data source
- BREAKING CHANGE: Remove `prefix_sid_absolute`, `prefix_sid_index`, `prefix_sid_n_flag_clear` in `iosxr_router_isis_interface_address_family` resource and data source
- BREAKING CHANGE: Remove `hello_password_hmac_md5`, `hello_password_keychain`, `hello_password_text` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Remove `passive`, `shutdown`, `suppressed` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Remove `hello_padding_disable`, `hello_padding_sometimes` in `iosxr_router_isis_interface` resource and data source
- BREAKING CHANGE: Rename `redistribute_isis` to `redistribute_isis_processes` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Rename `instance_id` to `isis_string` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `spf_prefix_priorities` (type: List) in `router_isis_address_family` resource and data source
- BREAKING CHANGE: Rename `mpls_traffic_eng_router_id_ip_address` to `mpls_traffic_eng_router_id_ipv4_address` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Rename `mpls_traffic_eng_router_id_interface` to `mpls_traffic_eng_router_id_interface_name` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `microloop_avoidance_protected` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `microloop_avoidance_segment_routing` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `fast_reroute_per_link_priority_limit_critical` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `fast_reroute_per_link_priority_limit_high` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `fast_reroute_per_link_priority_limit_medium` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `fast_reroute_per_prefix_priority_limit_critical` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `fast_reroute_per_prefix_priority_limit_high` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `fast_reroute_per_prefix_priority_limit_medium` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Rename `interface_name` to `router_id_interface_name` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Rename `ip_address` to `router_id_ip_address` in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Rename `level_id` to `level_number` in `metric_style_levels` list in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `narrow` from `metric_style_levels` list in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `wide` from `metric_style_levels` list in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `transition` from `metric_style_levels` list in `iosxr_router_isis_address_family` resource and data source
- BREAKING CHANGE: Remove `passive` from `interfaces` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Remove `suppressed` from `interfaces` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Remove `shutdown` from `interfaces` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `state` (enumeration) to `interfaces` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Remove `priority` from `interfaces` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `priority_levels` (type: List) to `interfaces` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Change `hello_padding` from Bool to String (enumeration) in `interfaces` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Rename `algorithm_number` to `flex_algo_number` in `flex_algos` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Rename `metric_type_delay` to `metric_type` in `flex_algos` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Rename `affinity_map_name` to `name` in `affinity_maps` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Remove `lsp_password_keychain` from `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_accept_encrypted` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_accept_levels` (type: List) to `router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_text_encrypted` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_text_send_only` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_text_snp_send_only` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_text_enable_poi` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_hmac_md5_encrypted` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_hmac_md5_send_only` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_hmac_md5_snp_send_only` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_hmac_md5_enable_poi` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_keychain_name` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_keychain_send_only` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_keychain_snp_send_only` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `lsp_password_keychain_enable_poi` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Rename `lsp_refresh_interval_lsp_refresh_interval_time` to `lsp_refresh_interval` in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Rename `max_lsp_lifetime_max_lsp_lifetime` to `max_lsp_lifetime` in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Rename `level_id` to `level_number` in `set_overload_bit_levels` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `nsf_cisco` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `nsf_ietf` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `set_overload_bit_on_startup_time_to_advertise_seconds` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `set_overload_bit_on_startup_wait_for_bgp` to `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `on_startup_time_to_advertise_seconds` to `set_overload_bit_levels` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `on_startup_wait_for_bgp` to `set_overload_bit_levels` list in `iosxr_router_isis` resource and data source
- BREAKING CHANGE: Add `segment_routing_srv6_usid_allocation_wide_local_id_block` in `iosxr_router_bgp` resource and data source
- BREAKING CHANGE: Add `segment_routing_srv6_alloc_mode_per_ce` in `iosxr_router_bgp_vrf_address_family` resource and data source
- BREAKING CHANGE: Add `label_mode_per_vrf` to `segment_routing_srv6_alloc_mode_per_vrf` in `iosxr_router_bgp_vrf_address_family` resource and data source
- BREAKING CHANGE: Add `label_mode_per_vrf_46` to `segment_routing_srv6_alloc_mode_per_vrf_46` in `iosxr_router_bgp_vrf_address_family` resource and data source
- BREAKING CHANGE: Add `label_mode_route_policy` to `segment_routing_srv6_alloc_mode_route_policy` in `iosxr_router_bgp_vrf_address_family` resource and data source
- BREAKING CHANGE: Rename `local_as_replace_as` to `local_as_no_prepend_replace_as` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `local_as_dual_as` to `local_as_no_prepend_replace_as_dual_as` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `neighbor` to `address` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `rd_auto` to `rd_auto` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `rd_two_byte_as_as_number` to `rd_two_byte_as_number` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `rd_two_byte_as_index` to `rd_two_byte_as_index` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `rd_four_byte_as_as_number` to `rd_four_byte_as_number` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `rd_four_byte_as_index` to `rd_four_byte_as_index` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `rd_ip_address_ipv4_address` to `rd_ipv4_address_address` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `rd_ip_address_index` to `rd_ipv4_address_index` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Remove `ao_key_chain_name` in `iosxr_router_bgp_neighbor_group` resource and data source
- BREAKING CHANGE: Remove `ao_include_tcp_options_enable` in `iosxr_router_bgp_neighbor_group` resource and data source
- BREAKING CHANGE: Rename `default_originate_route_policy` in `iosxr_router_bgp_neighbor_address_family` resource and data source
- BREAKING CHANGE: Rename `default_originate_inheritance_disable` in `iosxr_router_bgp_neighbor_address_family` resource and data source
- BREAKING CHANGE: Rename `default_originate_route_policy` in `iosxr_router_bgp_vrf_neighbor_address_family` resource and data source
- BREAKING CHANGE: Rename `default_originate_inheritance_disable` in `iosxr_router_bgp_vrf_neighbor_address_family` resource and data source
- BREAKING CHANGE: Rename `redistribute_isis` (yang_name: `redistribute/isis` to `redistribute/isis-processes/isis-process`) in `iosxr_router_bgp_address_family` resource and data source
- BREAKING CHANGE: Remove `import` in `iosxr_router_bgp_neighbor_address_family` resource and data source
- BREAKING CHANGE: Rename `encapsulation-type/srv6` to `encapsulation_type` in `iosxr_router_bgp_neighbor_address_family` resource and data source
- BREAKING CHANGE: Rename `masklength` to `address_prefix` (yang_name: `masklength` to `address-prefix`) in `iosxr_router_bgp_address_family` resource and data source
- BREAKING CHANGE: Rename `masklength` to `address_prefix` in `iosxr_router_bgp_vrf_address_family` resource and data source
- BREAKING CHANGE: Remove `echo_ipv4_bundle_per_member_preferred_minimum_interval` in `iosxr_bfd` resource and data source
- BREAKING CHANGE: Rename `location-name` to `location-id` in `iosxr_bfd` resource and data source
- BREAKING CHANGE: Remove `ethernet_segment_identifier_type_zero_bytes_1`, `ethernet_segment_identifier_type_zero_bytes_23`, `ethernet_segment_identifier_type_zero_bytes_45`, `ethernet_segment_identifier_type_zero_bytes_67`, `ethernet_segment_identifier_type_zero_bytes_89` in `iosxr_evpn_interface` resource and data source
- BREAKING CHANGE: Rename `compress-level` to `compress` in `iosxr_interface` resource and data source
- BREAKING CHANGE: Rename `ipv6_access_group_egress_acl1` to `ipv6_access_group_egress_acl` in `iosxr_interface` resource and data source
- BREAKING CHANGE: Rename `timers_bgp_holdtime_holdtime` to `timers_bgp_holdtime_number` in `iosxr_router_bgp` resource and data source
- BREAKING CHANGE: Rename `neighbor-address` to `address` in `iosxr_router_bgp` resource and data source
- BREAKING CHANGE: Rename `advertisement_interval_time_in_seconds` to `advertisement_interval_advertisement_interval_time_in_seconds` in `iosxr_router_bgp` resource and data source
- BREAKING CHANGE: Rename `advertisement_interval_time_in_milliseconds` to `advertisement_interval_advertisement_interval_time_in_milliseconds` in `iosxr_router_bgp` resource and data source
- BREAKING CHANGE: Remove `bfd_fast_detect_inheritance_disable` in `iosxr_router_bgp` resource and data source
- BREAKING CHANGE: Rename `neighbor-address` to `address` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Remove `bgp_bestpath_med_confed` in `iosxr_router_bgp_vrf` resource and data source
- BREAKING CHANGE: Rename `advertisement_interval_time_in_seconds` to `advertisement_interval_advertisement_interval_time_in_seconds` in `iosxr_router_bgp_neighbor_group` resource and data source
- BREAKING CHANGE: Rename `advertisement_interval_time_in_milliseconds` to `advertisement_interval_advertisement_interval_time_in_milliseconds` in `iosxr_router_bgp_neighbor_group` resource and data source
- BREAKING CHANGE: Remove `bfd_fast_detect_inheritance_disable` in `iosxr_router_bgp_neighbor_group` resource and data source
- BREAKING CHANGE: Rename `neighbor-address` to `address` in `iosxr_router_bgp_vrf_neighbor_address_family` resource and data source
- BREAKING CHANGE: Remove `allocate_label_all_unlabeled_path` in `iosxr_router_bgp_address_family` resource and data source
- BREAKING CHANGE: Remove `allocate_label_all_unlabeled_path` in `iosxr_router_bgp_vrf_address_family` resource and data source
- BREAKING CHANGE: Remove `maximum_paths_ebgp_multipath`, `maximum_paths_ibgp_multipath`, `maximum_paths_eibgp_multipath` in `iosxr_router_bgp_address_family` resource and data source
- BREAKING CHANGE: Remove `maximum_paths_ebgp_multipath`, `maximum_paths_ibgp_multipath`, `maximum_paths_eibgp_multipath` in `iosxr_router_bgp_vrf_address_family` resource and data source

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

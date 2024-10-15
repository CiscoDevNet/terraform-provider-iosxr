---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

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


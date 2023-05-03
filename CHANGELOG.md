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

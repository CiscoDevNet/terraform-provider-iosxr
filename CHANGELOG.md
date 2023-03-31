## 0.1.4 (unreleased)

- Add `iosxr_ssh` resource and data source
- Allow concurrent changes across different devices

## 0.1.3

- Add support for IOS-XR 7.8.1+

## 0.1.2

- Update dependencies and go version

## 0.1.1

- Add delete attribute to iosxr_gnmi resource
- Allow provider config without host attribute in case devices attribute is being used
- Enhance iosxr_gnmi resource to support nested attributes (within YANG containers) using "/" as separator
- Enhance iosxr_gnmi resource to support nested YANG lists
- BREAKING CHANGE: remove iosxr_vrf_route_target_two_byte_as_format resource and data source
- BREAKING CHANGE: remove iosxr_vrf_route_target_four_byte_as_format resource and data source
- BREAKING CHANGE: remove iosxr_vrf_route_target_ip_address_format resource and data source
- Add route target attributes to iosxr_vrf resource and data source
- BREAKING CHANGE: remove iosxr_interface_ipv4 resource and data source
- BREAKING CHANGE: remove iosxr_interface_ipv6 resource and data source
- BREAKING CHANGE: remove iosxr_interface_ipv6_address resource and data source
- Add ipv4 and ipv6 attributes to iosxr_interface resource and data source
- BREAKING CHANGE: remove iosxr_mpls_ldp_address_family resource and data source
- BREAKING CHANGE: remove iosxr_mpls_ldp_interface resource and data source
- Add address_family and interface attributes to iosxr_mpls_ldp resource and data source
- BREAKING CHANGE: remove iosxr_l2vpn_xconnect_group resource and data source
- Add xconnect_group attributes to iosxr_l2vpn resource and data source
- BREAKING CHANGE: remove iosxr_l2vpn_xconnect_group_p2p_interface resource and data source
- BREAKING CHANGE: remove iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv4 resource and data source
- BREAKING CHANGE: remove iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6 resource and data source
- Add interface and neighbor attributes to iosxr_l2vpn_xconnect_group_p2p resource and data source
- BREAKING CHANGE: remove iosxr_router_isis_address_family resource and data source
- BREAKING CHANGE: remove iosxr_router_isis_interface resource and data source
- BREAKING CHANGE: remove iosxr_router_isis_net resource and data source
- Add address_family, interface and net attributes to iosxr_router_isis resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_area resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_redistribute_bgp resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_redistribute_isis resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_redistribute_ospf resource and data source
- Add area and redistribute attributes to iosxr_router_ospf resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_vrf_area resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_vrf_redistribute_bgp resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_vrf_redistribute_isis resource and data source
- BREAKING CHANGE: remove iosxr_router_ospf_vrf_redistribute_ospf resource and data source
- Add area and redistribute attributes to iosxr_router_ospf_vrf resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_neighbor resource and data source
- Add neighbor attributes to iosxr_router_bgp resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_vrf_neighbor resource and data source
- Add neighbor attributes to iosxr_router_bgp_vrf resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_address_family_aggregate_address resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_address_family_network resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_address_family_redistribute_isis resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_address_family_redistribute_ospf resource and data source
- Add aggregate_address, network and redistribute attributes to iosxr_router_bgp_address_family resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_vrf_address_family_aggregate_address resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_vrf_address_family_network resource and data source
- BREAKING CHANGE: remove iosxr_router_bgp_vrf_address_family_redistribute_ospf resource and data source
- Add aggregate_address, network and redistribute attributes to iosxr_router_bgp_vrf_address_family resource and data source

## 0.1.0

- Initial release

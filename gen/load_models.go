// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var models = []string{
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hostname-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-ip-address-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-vrf-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-interface-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-ipv4-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-statistics-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-static-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-service-policy-qos-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-bundle-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-ethernet-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-l2vpn-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-key-chain-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-location-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mpls-ldp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mpls-te-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-policymap-classmap-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-pce-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-cfg-mibs-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-hsrp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-entity-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-error-disable-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-line-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-line-exec-timeout-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-line-general-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-line-timestamp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-telnet-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-system-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-bridgemib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-entity-state-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-entity-redundancy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-l2vpn-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-flow-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mibs-ifmib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-mpls-ldp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ipv4-access-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ipv6-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-vrrp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ipv6-access-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-access-list-datatypes.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ipv6-prefix-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ipv4-prefix-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mpls-l3vpn-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mibs-sensormib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-fru-ctrl-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-isis-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-bgp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ntp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-segment-routing-srv6-datatypes.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-segment-routing-srv6-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-segment-routing-ms-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-infra-xtc-agent-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-config-copy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-power-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ethernet-oam-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-bfd-sbfd-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mibs-rfmib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mpls-oam-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-segment-routing-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-bgp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-logging-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-logging-events-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-isis-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-ospf-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-snmp-server-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-vrf-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ssh-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-route-policy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-l2-ethernet-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-l2transport-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-statistics-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/cisco-semver.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/ietf-inet-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/ietf-yang-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/tailf-cli-extensions.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/tailf-common.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/tailf-meta-extensions.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-banner-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-cdp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-lldp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-lacp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-domain-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-service-timestamps-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-fpd-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-flow-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-access-group-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-ipv6-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-aaa-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-aaa-tacacs-server-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-aaa-task-user-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-cli-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-linux-networking-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-netconf-yang-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-xml-agent-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-tpa-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-performance-measurement-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-performance-mgmt-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ipsla-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-track-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-license-smart-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-smart-license-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-call-home-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-cef-accounting-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-cef-load-balancing-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-cef-pd-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-profile-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-acl-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-l3-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-port-range-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-profile-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-quad-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-service-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-shut-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-subslot-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-hw-module-vrf-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-monitor-session-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-lawful-intercept-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-lpts-profiling-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-lpts-punt-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-lpts-punt-flow-trap-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-cli-alias-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ftp-tftp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-crypto-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-dhcp-ipv4-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-dhcp-ipv6-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-icmp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-flowspec-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-subscriber-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-pbr-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-pbr-policy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-dynamic-template-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-dyn-tmpl-service-policy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-gnss-receiver-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ptp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ptp-log-servo-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-igmp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-mld-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-igmp-snooping-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mld-snooping-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-pim-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-tcp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-tunnel-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-mac-address-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-ifmgr-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ipv6-nd-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-arp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-if-mpls-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-control-plane-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-router-rib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ethernet-cfm-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-flash-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-syslog-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-alarm-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-ipsla-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-traps-pim-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-mibs-cbqosmib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-snmp-server-mroutemib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-snmp-server-notification-log-mib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-logging-correlator-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-telemetry-model-driven-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-vty-pool-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-icmp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-npu-hw-profile-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-macsec-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-ptp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-frequency-synchronization-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-optics-speed-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-optics-driver-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-controller-optics-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-wanphy-ui-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-aaa-radius-server-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-rsvp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-8000-hw-module-profile-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-um-evpn-host-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/2442/Cisco-IOS-XR-8000-fib-platform-cfg.yang",
}

const (
	modelsPath = "./gen/models/"
)

func main() {
	for _, model := range models {
		f := strings.Split(model, "/")
		path := filepath.Join(modelsPath, f[len(f)-1])
		if _, err := os.Stat(path); err != nil {
			err := downloadModel(path, model)
			if err != nil {
				panic(err)
			}
			fmt.Println("Downloaded model: " + path)
		}
	}
}

func downloadModel(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

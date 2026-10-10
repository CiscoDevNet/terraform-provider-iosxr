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
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ModelEntry represents a YANG model with its URL, release version, and filename
type ModelEntry struct {
	URL      string
	Version  string
	Filename string
}

// release describes one IOS-XR release to download YANG models for.
type release struct {
	Version     string   // dotted major.minor; names gen/models/<Version> and gen/definitions/<Version>
	UpstreamDir string   // literal YangModels directory under vendor/cisco/xr (for example "2442")
	Files       []string // YANG file names in UpstreamDir, each saved to gen/models/<Version>/
}

// yangBaseURL is the YangModels directory that holds one sub-directory per IOS-XR release.
const yangBaseURL = "https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/"

// releases lists every supported IOS-XR release with its full file list.
// To add a release, append an entry. Version is the dotted major.minor used locally (for example
// "25.10"). UpstreamDir is the YangModels directory name, written out by hand: it concatenates the
// release digits (24.4.2 is "2442") and is never derived from Version.
var releases = []release{
	{
		Version:     "24.4",
		UpstreamDir: "2442", // IOS-XR 24.4.2 (base version)
		Files: []string{
			"Cisco-IOS-XR-types.yang",
			"Cisco-IOS-XR-um-hostname-cfg.yang",
			"Cisco-IOS-XR-um-if-ip-address-cfg.yang",
			"Cisco-IOS-XR-um-if-vrf-cfg.yang",
			"Cisco-IOS-XR-um-interface-cfg.yang",
			"Cisco-IOS-XR-um-if-ipv4-cfg.yang",
			"Cisco-IOS-XR-um-statistics-cfg.yang",
			"Cisco-IOS-XR-um-router-static-cfg.yang",
			"Cisco-IOS-XR-um-if-service-policy-qos-cfg.yang",
			"Cisco-IOS-XR-um-if-bundle-cfg.yang",
			"Cisco-IOS-XR-um-if-ethernet-cfg.yang",
			"Cisco-IOS-XR-um-l2vpn-cfg.yang",
			"Cisco-IOS-XR-um-key-chain-cfg.yang",
			"Cisco-IOS-XR-um-location-cfg.yang",
			"Cisco-IOS-XR-um-mpls-ldp-cfg.yang",
			"Cisco-IOS-XR-um-mpls-te-cfg.yang",
			"Cisco-IOS-XR-um-policymap-classmap-cfg.yang",
			"Cisco-IOS-XR-um-pce-cfg.yang",
			"Cisco-IOS-XR-um-cfg-mibs-cfg.yang",
			"Cisco-IOS-XR-um-router-hsrp-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-cfg.yang",
			"Cisco-IOS-XR-um-error-disable-cfg.yang",
			"Cisco-IOS-XR-um-line-cfg.yang",
			"Cisco-IOS-XR-um-line-exec-timeout-cfg.yang",
			"Cisco-IOS-XR-um-line-general-cfg.yang",
			"Cisco-IOS-XR-um-line-timestamp-cfg.yang",
			"Cisco-IOS-XR-um-telnet-cfg.yang",
			"Cisco-IOS-XR-um-traps-system-cfg.yang",
			"Cisco-IOS-XR-um-traps-bridgemib-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-state-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-redundancy-cfg.yang",
			"Cisco-IOS-XR-um-flow-cfg.yang",
			"Cisco-IOS-XR-um-mibs-ifmib-cfg.yang",
			"Cisco-IOS-XR-um-traps-mpls-ldp-cfg.yang",
			"Cisco-IOS-XR-um-ipv4-access-list-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-router-vrrp-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-access-list-cfg.yang",
			"Cisco-IOS-XR-um-access-list-datatypes.yang",
			"Cisco-IOS-XR-um-ipv6-prefix-list-cfg.yang",
			"Cisco-IOS-XR-um-ipv4-prefix-list-cfg.yang",
			"Cisco-IOS-XR-um-mpls-l3vpn-cfg.yang",
			"Cisco-IOS-XR-um-mibs-sensormib-cfg.yang",
			"Cisco-IOS-XR-um-traps-fru-ctrl-cfg.yang",
			"Cisco-IOS-XR-um-router-isis-cfg.yang",
			"Cisco-IOS-XR-um-router-bgp-cfg.yang",
			"Cisco-IOS-XR-um-ntp-cfg.yang",
			"Cisco-IOS-XR-segment-routing-srv6-datatypes.yang",
			"Cisco-IOS-XR-segment-routing-srv6-cfg.yang",
			"Cisco-IOS-XR-segment-routing-ms-cfg.yang",
			"Cisco-IOS-XR-infra-xtc-agent-cfg.yang",
			"Cisco-IOS-XR-um-traps-config-copy-cfg.yang",
			"Cisco-IOS-XR-um-traps-power-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-oam-cfg.yang",
			"Cisco-IOS-XR-um-bfd-sbfd-cfg.yang",
			"Cisco-IOS-XR-um-mibs-rfmib-cfg.yang",
			"Cisco-IOS-XR-um-mpls-oam-cfg.yang",
			"Cisco-IOS-XR-um-segment-routing-cfg.yang",
			"Cisco-IOS-XR-um-logging-cfg.yang",
			"Cisco-IOS-XR-um-logging-events-cfg.yang",
			"Cisco-IOS-XR-um-clock-exr-cfg.yang",
			"Cisco-IOS-XR-um-router-ospf-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-cfg.yang",
			"Cisco-IOS-XR-um-vrf-cfg.yang",
			"Cisco-IOS-XR-um-ssh-cfg.yang",
			"Cisco-IOS-XR-um-route-policy-cfg.yang",
			"Cisco-IOS-XR-um-l2-ethernet-cfg.yang",
			"Cisco-IOS-XR-um-if-l2transport-cfg.yang",
			"cisco-semver.yang",
			"ietf-inet-types.yang",
			"ietf-yang-types.yang",
			"tailf-cli-extensions.yang",
			"tailf-common.yang",
			"tailf-meta-extensions.yang",
			"Cisco-IOS-XR-um-banner-cfg.yang",
			"Cisco-IOS-XR-um-cdp-cfg.yang",
			"Cisco-IOS-XR-um-lldp-cfg.yang",
			"Cisco-IOS-XR-um-lacp-cfg.yang",
			"Cisco-IOS-XR-um-domain-cfg.yang",
			"Cisco-IOS-XR-um-service-timestamps-cfg.yang",
			"Cisco-IOS-XR-um-fpd-cfg.yang",
			"Cisco-IOS-XR-um-if-access-group-cfg.yang",
			"Cisco-IOS-XR-um-if-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-aaa-cfg.yang",
			"Cisco-IOS-XR-um-aaa-tacacs-server-cfg.yang",
			"Cisco-IOS-XR-um-aaa-task-user-cfg.yang",
			"Cisco-IOS-XR-cli-cfg.yang",
			"Cisco-IOS-XR-um-linux-networking-cfg.yang",
			"Cisco-IOS-XR-um-netconf-yang-cfg.yang",
			"Cisco-IOS-XR-um-xml-agent-cfg.yang",
			"Cisco-IOS-XR-um-tpa-cfg.yang",
			"Cisco-IOS-XR-um-performance-measurement-cfg.yang",
			"Cisco-IOS-XR-um-performance-mgmt-cfg.yang",
			"Cisco-IOS-XR-um-ipsla-cfg.yang",
			"Cisco-IOS-XR-um-track-cfg.yang",
			"Cisco-IOS-XR-um-license-smart-cfg.yang",
			"Cisco-IOS-XR-smart-license-cfg.yang",
			"Cisco-IOS-XR-um-call-home-cfg.yang",
			"Cisco-IOS-XR-um-cef-accounting-cfg.yang",
			"Cisco-IOS-XR-um-cef-load-balancing-cfg.yang",
			"Cisco-IOS-XR-um-cef-pd-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-profile-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-acl-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-l3-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-port-range-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-quad-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-service-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-shut-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-subslot-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-vrf-cfg.yang",
			"Cisco-IOS-XR-um-monitor-session-cfg.yang",
			"Cisco-IOS-XR-um-lawful-intercept-cfg.yang",
			"Cisco-IOS-XR-um-lpts-profiling-cfg.yang",
			"Cisco-IOS-XR-um-lpts-punt-cfg.yang",
			"Cisco-IOS-XR-um-lpts-punt-flow-trap-cfg.yang",
			"Cisco-IOS-XR-um-cli-alias-cfg.yang",
			"Cisco-IOS-XR-um-ftp-tftp-cfg.yang",
			"Cisco-IOS-XR-um-http-client-cfg.yang",
			"Cisco-IOS-XR-um-crypto-cfg.yang",
			"Cisco-IOS-XR-um-dhcp-ipv4-cfg.yang",
			"Cisco-IOS-XR-um-dhcp-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-icmp-cfg.yang",
			"Cisco-IOS-XR-um-flowspec-cfg.yang",
			"Cisco-IOS-XR-um-subscriber-cfg.yang",
			"Cisco-IOS-XR-um-pbr-cfg.yang",
			"Cisco-IOS-XR-um-pbr-policy-cfg.yang",
			"Cisco-IOS-XR-um-dynamic-template-cfg.yang",
			"Cisco-IOS-XR-um-dyn-tmpl-service-policy-cfg.yang",
			"Cisco-IOS-XR-um-gnss-receiver-cfg.yang",
			"Cisco-IOS-XR-um-ptp-cfg.yang",
			"Cisco-IOS-XR-um-ptp-log-servo-cfg.yang",
			"Cisco-IOS-XR-um-router-igmp-cfg.yang",
			"Cisco-IOS-XR-um-router-mld-cfg.yang",
			"Cisco-IOS-XR-um-igmp-snooping-cfg.yang",
			"Cisco-IOS-XR-um-mld-snooping-cfg.yang",
			"Cisco-IOS-XR-um-router-pim-cfg.yang",
			"Cisco-IOS-XR-um-tcp-cfg.yang",
			"Cisco-IOS-XR-um-if-tunnel-cfg.yang",
			"Cisco-IOS-XR-um-if-mac-address-cfg.yang",
			"Cisco-IOS-XR-ifmgr-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-nd-cfg.yang",
			"Cisco-IOS-XR-um-if-arp-cfg.yang",
			"Cisco-IOS-XR-um-if-mpls-cfg.yang",
			"Cisco-IOS-XR-um-control-plane-cfg.yang",
			"Cisco-IOS-XR-um-router-rib-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-cfm-cfg.yang",
			"Cisco-IOS-XR-um-traps-flash-cfg.yang",
			"Cisco-IOS-XR-um-traps-syslog-cfg.yang",
			"Cisco-IOS-XR-um-traps-alarm-cfg.yang",
			"Cisco-IOS-XR-um-traps-ipsla-cfg.yang",
			"Cisco-IOS-XR-um-traps-pim-cfg.yang",
			"Cisco-IOS-XR-um-mibs-cbqosmib-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-mroutemib-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-notification-log-mib-cfg.yang",
			"Cisco-IOS-XR-um-logging-correlator-cfg.yang",
			"Cisco-IOS-XR-um-telemetry-model-driven-cfg.yang",
			"Cisco-IOS-XR-um-vty-pool-cfg.yang",
			"Cisco-IOS-XR-npu-hw-profile-cfg.yang",
			"Cisco-IOS-XR-um-macsec-cfg.yang",
			"Cisco-IOS-XR-um-frequency-synchronization-cfg.yang",
			"Cisco-IOS-XR-optics-speed-cfg.yang",
			"Cisco-IOS-XR-optics-driver-cfg.yang",
			"Cisco-IOS-XR-controller-optics-cfg.yang",
			"Cisco-IOS-XR-wanphy-ui-cfg.yang",
			"Cisco-IOS-XR-um-aaa-radius-server-cfg.yang",
			"Cisco-IOS-XR-um-rsvp-cfg.yang",
			"Cisco-IOS-XR-um-8000-hw-module-profile-cfg.yang",
			"Cisco-IOS-XR-um-evpn-host-cfg.yang",
			"Cisco-IOS-XR-8000-fib-platform-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-sla-cfg.yang",
		},
	},
	{
		Version:     "25.4",
		UpstreamDir: "2542", // IOS-XR 25.4.2
		Files: []string{
			"Cisco-IOS-XR-types.yang",
			"Cisco-IOS-XR-um-hostname-cfg.yang",
			"Cisco-IOS-XR-um-if-ip-address-cfg.yang",
			"Cisco-IOS-XR-um-if-vrf-cfg.yang",
			"Cisco-IOS-XR-um-interface-cfg.yang",
			"Cisco-IOS-XR-um-if-ipv4-cfg.yang",
			"Cisco-IOS-XR-um-statistics-cfg.yang",
			"Cisco-IOS-XR-um-router-static-cfg.yang",
			"Cisco-IOS-XR-um-if-service-policy-qos-cfg.yang",
			"Cisco-IOS-XR-um-if-bundle-cfg.yang",
			"Cisco-IOS-XR-um-if-ethernet-cfg.yang",
			"Cisco-IOS-XR-um-l2vpn-cfg.yang",
			"Cisco-IOS-XR-um-key-chain-cfg.yang",
			"Cisco-IOS-XR-um-location-cfg.yang",
			"Cisco-IOS-XR-um-mpls-ldp-cfg.yang",
			"Cisco-IOS-XR-um-mpls-te-cfg.yang",
			"Cisco-IOS-XR-um-policymap-classmap-cfg.yang",
			"Cisco-IOS-XR-um-pce-cfg.yang",
			"Cisco-IOS-XR-um-cfg-mibs-cfg.yang",
			"Cisco-IOS-XR-um-router-hsrp-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-cfg.yang",
			"Cisco-IOS-XR-um-error-disable-cfg.yang",
			"Cisco-IOS-XR-um-line-cfg.yang",
			"Cisco-IOS-XR-um-line-exec-timeout-cfg.yang",
			"Cisco-IOS-XR-um-line-general-cfg.yang",
			"Cisco-IOS-XR-um-line-timestamp-cfg.yang",
			"Cisco-IOS-XR-um-telnet-cfg.yang",
			"Cisco-IOS-XR-um-traps-system-cfg.yang",
			"Cisco-IOS-XR-um-traps-bridgemib-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-state-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-redundancy-cfg.yang",
			"Cisco-IOS-XR-um-flow-cfg.yang",
			"Cisco-IOS-XR-um-mibs-ifmib-cfg.yang",
			"Cisco-IOS-XR-um-traps-mpls-ldp-cfg.yang",
			"Cisco-IOS-XR-um-ipv4-access-list-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-router-vrrp-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-access-list-cfg.yang",
			"Cisco-IOS-XR-um-access-list-datatypes.yang",
			"Cisco-IOS-XR-um-ipv6-prefix-list-cfg.yang",
			"Cisco-IOS-XR-um-ipv4-prefix-list-cfg.yang",
			"Cisco-IOS-XR-um-mpls-l3vpn-cfg.yang",
			"Cisco-IOS-XR-um-mibs-sensormib-cfg.yang",
			"Cisco-IOS-XR-um-traps-fru-ctrl-cfg.yang",
			"Cisco-IOS-XR-um-router-isis-cfg.yang",
			"Cisco-IOS-XR-um-router-bgp-cfg.yang",
			"Cisco-IOS-XR-um-ntp-cfg.yang",
			"Cisco-IOS-XR-segment-routing-srv6-datatypes.yang",
			"Cisco-IOS-XR-segment-routing-srv6-cfg.yang",
			"Cisco-IOS-XR-segment-routing-ms-cfg.yang",
			"Cisco-IOS-XR-infra-xtc-agent-cfg.yang",
			"Cisco-IOS-XR-um-traps-config-copy-cfg.yang",
			"Cisco-IOS-XR-um-traps-power-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-oam-cfg.yang",
			"Cisco-IOS-XR-um-bfd-sbfd-cfg.yang",
			"Cisco-IOS-XR-um-mibs-rfmib-cfg.yang",
			"Cisco-IOS-XR-um-mpls-oam-cfg.yang",
			"Cisco-IOS-XR-um-segment-routing-cfg.yang",
			"Cisco-IOS-XR-um-logging-cfg.yang",
			"Cisco-IOS-XR-um-logging-events-cfg.yang",
			"Cisco-IOS-XR-um-router-ospf-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-cfg.yang",
			"Cisco-IOS-XR-um-vrf-cfg.yang",
			"Cisco-IOS-XR-um-ssh-cfg.yang",
			"Cisco-IOS-XR-um-route-policy-cfg.yang",
			"Cisco-IOS-XR-um-l2-ethernet-cfg.yang",
			"Cisco-IOS-XR-um-if-l2transport-cfg.yang",
			"cisco-semver.yang",
			"ietf-inet-types.yang",
			"ietf-yang-types.yang",
			"tailf-cli-extensions.yang",
			"tailf-common.yang",
			"tailf-meta-extensions.yang",
			"Cisco-IOS-XR-um-banner-cfg.yang",
			"Cisco-IOS-XR-um-cdp-cfg.yang",
			"Cisco-IOS-XR-um-lldp-cfg.yang",
			"Cisco-IOS-XR-um-lacp-cfg.yang",
			"Cisco-IOS-XR-um-domain-cfg.yang",
			"Cisco-IOS-XR-um-service-cfg.yang",
			"Cisco-IOS-XR-um-fpd-cfg.yang",
			"Cisco-IOS-XR-um-if-access-group-cfg.yang",
			"Cisco-IOS-XR-um-if-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-aaa-cfg.yang",
			"Cisco-IOS-XR-um-aaa-tacacs-server-cfg.yang",
			"Cisco-IOS-XR-um-aaa-task-user-cfg.yang",
			"Cisco-IOS-XR-cli-cfg.yang",
			"Cisco-IOS-XR-um-linux-networking-cfg.yang",
			"Cisco-IOS-XR-um-netconf-yang-cfg.yang",
			"Cisco-IOS-XR-um-xml-agent-cfg.yang",
			"Cisco-IOS-XR-um-performance-measurement-cfg.yang",
			"Cisco-IOS-XR-um-performance-mgmt-cfg.yang",
			"Cisco-IOS-XR-um-ipsla-cfg.yang",
			"Cisco-IOS-XR-um-track-cfg.yang",
			"Cisco-IOS-XR-um-license-smart-cfg.yang",
			"Cisco-IOS-XR-smart-license-cfg.yang",
			"Cisco-IOS-XR-um-call-home-cfg.yang",
			"Cisco-IOS-XR-um-cef-accounting-cfg.yang",
			"Cisco-IOS-XR-um-cef-load-balancing-cfg.yang",
			"Cisco-IOS-XR-um-cef-pd-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-profile-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-acl-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-l3-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-port-range-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-quad-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-service-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-shut-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-subslot-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-vrf-cfg.yang",
			"Cisco-IOS-XR-um-monitor-session-cfg.yang",
			"Cisco-IOS-XR-um-lawful-intercept-cfg.yang",
			"Cisco-IOS-XR-um-lpts-profiling-cfg.yang",
			"Cisco-IOS-XR-um-lpts-punt-cfg.yang",
			"Cisco-IOS-XR-um-lpts-punt-flow-trap-cfg.yang",
			"Cisco-IOS-XR-um-cli-alias-cfg.yang",
			"Cisco-IOS-XR-um-ftp-tftp-cfg.yang",
			"Cisco-IOS-XR-um-crypto-cfg.yang",
			"Cisco-IOS-XR-um-dhcp-ipv4-cfg.yang",
			"Cisco-IOS-XR-um-dhcp-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-icmp-cfg.yang",
			"Cisco-IOS-XR-um-flowspec-cfg.yang",
			"Cisco-IOS-XR-um-subscriber-cfg.yang",
			"Cisco-IOS-XR-um-pbr-cfg.yang",
			"Cisco-IOS-XR-um-pbr-policy-cfg.yang",
			"Cisco-IOS-XR-um-dynamic-template-cfg.yang",
			"Cisco-IOS-XR-um-dyn-tmpl-service-policy-cfg.yang",
			"Cisco-IOS-XR-um-gnss-receiver-cfg.yang",
			"Cisco-IOS-XR-um-ptp-cfg.yang",
			"Cisco-IOS-XR-um-ptp-log-servo-cfg.yang",
			"Cisco-IOS-XR-um-router-igmp-cfg.yang",
			"Cisco-IOS-XR-um-router-mld-cfg.yang",
			"Cisco-IOS-XR-um-igmp-snooping-cfg.yang",
			"Cisco-IOS-XR-um-mld-snooping-cfg.yang",
			"Cisco-IOS-XR-um-router-pim-cfg.yang",
			"Cisco-IOS-XR-um-tcp-cfg.yang",
			"Cisco-IOS-XR-um-if-tunnel-cfg.yang",
			"Cisco-IOS-XR-um-if-mac-address-cfg.yang",
			"Cisco-IOS-XR-ifmgr-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-nd-cfg.yang",
			"Cisco-IOS-XR-um-if-arp-cfg.yang",
			"Cisco-IOS-XR-um-if-mpls-cfg.yang",
			"Cisco-IOS-XR-um-control-plane-cfg.yang",
			"Cisco-IOS-XR-um-router-rib-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-cfm-cfg.yang",
			"Cisco-IOS-XR-um-traps-flash-cfg.yang",
			"Cisco-IOS-XR-um-traps-syslog-cfg.yang",
			"Cisco-IOS-XR-um-traps-alarm-cfg.yang",
			"Cisco-IOS-XR-um-traps-ipsla-cfg.yang",
			"Cisco-IOS-XR-um-traps-pim-cfg.yang",
			"Cisco-IOS-XR-um-mibs-cbqosmib-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-mroutemib-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-notification-log-mib-cfg.yang",
			"Cisco-IOS-XR-um-logging-correlator-cfg.yang",
			"Cisco-IOS-XR-um-telemetry-model-driven-cfg.yang",
			"Cisco-IOS-XR-um-vty-pool-cfg.yang",
			"Cisco-IOS-XR-npu-hw-profile-cfg.yang",
			"Cisco-IOS-XR-um-macsec-cfg.yang",
			"Cisco-IOS-XR-um-frequency-synchronization-cfg.yang",
			"Cisco-IOS-XR-optics-speed-cfg.yang",
			"Cisco-IOS-XR-optics-driver-cfg.yang",
			"Cisco-IOS-XR-controller-optics-cfg.yang",
			"Cisco-IOS-XR-wanphy-ui-cfg.yang",
			"Cisco-IOS-XR-um-aaa-radius-server-cfg.yang",
			"Cisco-IOS-XR-um-rsvp-cfg.yang",
			"Cisco-IOS-XR-um-8000-hw-module-profile-cfg.yang",
			"Cisco-IOS-XR-um-evpn-host-cfg.yang",
			"Cisco-IOS-XR-8000-fib-platform-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-sla-cfg.yang",
			"Cisco-IOS-XR-um-http-client-cfg.yang",
		},
	},
	{
		Version:     "26.2",
		UpstreamDir: "2621", // IOS-XR 26.2.1
		Files: []string{
			"Cisco-IOS-XR-types.yang",
			"Cisco-IOS-XR-um-hostname-cfg.yang",
			"Cisco-IOS-XR-um-if-ip-address-cfg.yang",
			"Cisco-IOS-XR-um-if-vrf-cfg.yang",
			"Cisco-IOS-XR-um-interface-cfg.yang",
			"Cisco-IOS-XR-um-if-ipv4-cfg.yang",
			"Cisco-IOS-XR-um-statistics-cfg.yang",
			"Cisco-IOS-XR-um-router-static-cfg.yang",
			"Cisco-IOS-XR-um-if-service-policy-qos-cfg.yang",
			"Cisco-IOS-XR-um-if-bundle-cfg.yang",
			"Cisco-IOS-XR-um-if-ethernet-cfg.yang",
			"Cisco-IOS-XR-um-l2vpn-cfg.yang",
			"Cisco-IOS-XR-um-key-chain-cfg.yang",
			"Cisco-IOS-XR-um-location-cfg.yang",
			"Cisco-IOS-XR-um-mpls-ldp-cfg.yang",
			"Cisco-IOS-XR-um-mpls-te-cfg.yang",
			"Cisco-IOS-XR-um-policymap-classmap-cfg.yang",
			"Cisco-IOS-XR-um-pce-cfg.yang",
			"Cisco-IOS-XR-um-cfg-mibs-cfg.yang",
			"Cisco-IOS-XR-um-router-hsrp-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-cfg.yang",
			"Cisco-IOS-XR-um-error-disable-cfg.yang",
			"Cisco-IOS-XR-um-line-cfg.yang",
			"Cisco-IOS-XR-um-line-exec-timeout-cfg.yang",
			"Cisco-IOS-XR-um-line-general-cfg.yang",
			"Cisco-IOS-XR-um-line-timestamp-cfg.yang",
			"Cisco-IOS-XR-um-telnet-cfg.yang",
			"Cisco-IOS-XR-um-traps-system-cfg.yang",
			"Cisco-IOS-XR-um-traps-bridgemib-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-state-cfg.yang",
			"Cisco-IOS-XR-um-traps-entity-redundancy-cfg.yang",
			"Cisco-IOS-XR-um-flow-cfg.yang",
			"Cisco-IOS-XR-um-mibs-ifmib-cfg.yang",
			"Cisco-IOS-XR-um-traps-mpls-ldp-cfg.yang",
			"Cisco-IOS-XR-um-ipv4-access-list-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-router-vrrp-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-access-list-cfg.yang",
			"Cisco-IOS-XR-um-access-list-datatypes.yang",
			"Cisco-IOS-XR-um-ipv6-prefix-list-cfg.yang",
			"Cisco-IOS-XR-um-ipv4-prefix-list-cfg.yang",
			"Cisco-IOS-XR-um-mpls-l3vpn-cfg.yang",
			"Cisco-IOS-XR-um-mibs-sensormib-cfg.yang",
			"Cisco-IOS-XR-um-traps-fru-ctrl-cfg.yang",
			"Cisco-IOS-XR-um-router-isis-cfg.yang",
			"Cisco-IOS-XR-um-router-bgp-cfg.yang",
			"Cisco-IOS-XR-um-ntp-cfg.yang",
			"Cisco-IOS-XR-segment-routing-srv6-datatypes.yang",
			"Cisco-IOS-XR-segment-routing-srv6-cfg.yang",
			"Cisco-IOS-XR-segment-routing-ms-cfg.yang",
			"Cisco-IOS-XR-infra-xtc-agent-cfg.yang",
			"Cisco-IOS-XR-um-traps-config-copy-cfg.yang",
			"Cisco-IOS-XR-um-traps-power-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-oam-cfg.yang",
			"Cisco-IOS-XR-um-bfd-sbfd-cfg.yang",
			"Cisco-IOS-XR-um-mibs-rfmib-cfg.yang",
			"Cisco-IOS-XR-um-mpls-oam-cfg.yang",
			"Cisco-IOS-XR-um-segment-routing-cfg.yang",
			"Cisco-IOS-XR-um-logging-cfg.yang",
			"Cisco-IOS-XR-um-logging-events-cfg.yang",
			"Cisco-IOS-XR-um-router-ospf-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-cfg.yang",
			"Cisco-IOS-XR-um-vrf-cfg.yang",
			"Cisco-IOS-XR-um-ssh-cfg.yang",
			"Cisco-IOS-XR-um-route-policy-cfg.yang",
			"Cisco-IOS-XR-um-l2-ethernet-cfg.yang",
			"Cisco-IOS-XR-um-if-l2transport-cfg.yang",
			"cisco-semver.yang",
			"ietf-inet-types.yang",
			"ietf-yang-types.yang",
			"tailf-cli-extensions.yang",
			"tailf-common.yang",
			"tailf-meta-extensions.yang",
			"Cisco-IOS-XR-um-banner-cfg.yang",
			"Cisco-IOS-XR-um-cdp-cfg.yang",
			"Cisco-IOS-XR-um-lldp-cfg.yang",
			"Cisco-IOS-XR-um-lacp-cfg.yang",
			"Cisco-IOS-XR-um-domain-cfg.yang",
			"Cisco-IOS-XR-um-service-cfg.yang",
			"Cisco-IOS-XR-um-fpd-cfg.yang",
			"Cisco-IOS-XR-um-if-access-group-cfg.yang",
			"Cisco-IOS-XR-um-if-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-aaa-cfg.yang",
			"Cisco-IOS-XR-um-aaa-tacacs-server-cfg.yang",
			"Cisco-IOS-XR-um-aaa-task-user-cfg.yang",
			"Cisco-IOS-XR-cli-cfg.yang",
			"Cisco-IOS-XR-um-linux-networking-cfg.yang",
			"Cisco-IOS-XR-um-netconf-yang-cfg.yang",
			"Cisco-IOS-XR-um-xml-agent-cfg.yang",
			"Cisco-IOS-XR-um-performance-measurement-cfg.yang",
			"Cisco-IOS-XR-um-performance-mgmt-cfg.yang",
			"Cisco-IOS-XR-um-ipsla-cfg.yang",
			"Cisco-IOS-XR-um-track-cfg.yang",
			"Cisco-IOS-XR-um-license-smart-cfg.yang",
			"Cisco-IOS-XR-smart-license-cfg.yang",
			"Cisco-IOS-XR-um-call-home-cfg.yang",
			"Cisco-IOS-XR-um-cef-accounting-cfg.yang",
			"Cisco-IOS-XR-um-cef-load-balancing-cfg.yang",
			"Cisco-IOS-XR-um-cef-pd-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-profile-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-acl-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-l3-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-port-range-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-quad-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-service-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-shut-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-subslot-cfg.yang",
			"Cisco-IOS-XR-um-hw-module-vrf-cfg.yang",
			"Cisco-IOS-XR-um-monitor-session-cfg.yang",
			"Cisco-IOS-XR-um-lawful-intercept-cfg.yang",
			"Cisco-IOS-XR-um-lpts-profiling-cfg.yang",
			"Cisco-IOS-XR-um-lpts-punt-cfg.yang",
			"Cisco-IOS-XR-um-lpts-punt-flow-trap-cfg.yang",
			"Cisco-IOS-XR-um-cli-alias-cfg.yang",
			"Cisco-IOS-XR-um-ftp-tftp-cfg.yang",
			"Cisco-IOS-XR-um-crypto-cfg.yang",
			"Cisco-IOS-XR-um-dhcp-ipv4-cfg.yang",
			"Cisco-IOS-XR-um-dhcp-ipv6-cfg.yang",
			"Cisco-IOS-XR-um-icmp-cfg.yang",
			"Cisco-IOS-XR-um-flowspec-cfg.yang",
			"Cisco-IOS-XR-um-subscriber-cfg.yang",
			"Cisco-IOS-XR-um-pbr-cfg.yang",
			"Cisco-IOS-XR-um-pbr-policy-cfg.yang",
			"Cisco-IOS-XR-um-dynamic-template-cfg.yang",
			"Cisco-IOS-XR-um-dyn-tmpl-service-policy-cfg.yang",
			"Cisco-IOS-XR-um-gnss-receiver-cfg.yang",
			"Cisco-IOS-XR-um-ptp-cfg.yang",
			"Cisco-IOS-XR-um-ptp-log-servo-cfg.yang",
			"Cisco-IOS-XR-um-router-igmp-cfg.yang",
			"Cisco-IOS-XR-um-router-mld-cfg.yang",
			"Cisco-IOS-XR-um-igmp-snooping-cfg.yang",
			"Cisco-IOS-XR-um-mld-snooping-cfg.yang",
			"Cisco-IOS-XR-um-router-pim-cfg.yang",
			"Cisco-IOS-XR-um-tcp-cfg.yang",
			"Cisco-IOS-XR-um-if-tunnel-cfg.yang",
			"Cisco-IOS-XR-um-if-mac-address-cfg.yang",
			"Cisco-IOS-XR-ifmgr-cfg.yang",
			"Cisco-IOS-XR-um-ipv6-nd-cfg.yang",
			"Cisco-IOS-XR-um-if-arp-cfg.yang",
			"Cisco-IOS-XR-um-if-mpls-cfg.yang",
			"Cisco-IOS-XR-um-control-plane-cfg.yang",
			"Cisco-IOS-XR-um-router-rib-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-cfm-cfg.yang",
			"Cisco-IOS-XR-um-traps-flash-cfg.yang",
			"Cisco-IOS-XR-um-traps-syslog-cfg.yang",
			"Cisco-IOS-XR-um-traps-alarm-cfg.yang",
			"Cisco-IOS-XR-um-traps-ipsla-cfg.yang",
			"Cisco-IOS-XR-um-traps-pim-cfg.yang",
			"Cisco-IOS-XR-um-mibs-cbqosmib-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-mroutemib-cfg.yang",
			"Cisco-IOS-XR-um-snmp-server-notification-log-mib-cfg.yang",
			"Cisco-IOS-XR-um-logging-correlator-cfg.yang",
			"Cisco-IOS-XR-um-telemetry-model-driven-cfg.yang",
			"Cisco-IOS-XR-um-vty-pool-cfg.yang",
			"Cisco-IOS-XR-npu-hw-profile-cfg.yang",
			"Cisco-IOS-XR-um-macsec-cfg.yang",
			"Cisco-IOS-XR-um-frequency-synchronization-cfg.yang",
			"Cisco-IOS-XR-optics-speed-cfg.yang",
			"Cisco-IOS-XR-optics-driver-cfg.yang",
			"Cisco-IOS-XR-controller-optics-cfg.yang",
			"Cisco-IOS-XR-wanphy-ui-cfg.yang",
			"Cisco-IOS-XR-um-aaa-radius-server-cfg.yang",
			"Cisco-IOS-XR-um-rsvp-cfg.yang",
			"Cisco-IOS-XR-um-8000-hw-module-profile-cfg.yang",
			"Cisco-IOS-XR-um-evpn-host-cfg.yang",
			"Cisco-IOS-XR-8000-fib-platform-cfg.yang",
			"Cisco-IOS-XR-um-ethernet-sla-cfg.yang",
		},
	},
}

// modelEntries expands releases into one ModelEntry per file.
func modelEntries(rs []release) []ModelEntry {
	var out []ModelEntry
	for _, r := range rs {
		for _, f := range r.Files {
			out = append(out, ModelEntry{URL: yangBaseURL + r.UpstreamDir + "/" + f, Version: r.Version, Filename: f})
		}
	}
	return out
}

// models is the combined list of all YANG models across all releases.
var models = modelEntries(releases)

const (
	modelsBasePath      = "./gen/models/"
	definitionsBasePath = "./gen/definitions/"
)

// majorMinor returns the major.minor part of a dotted version, dropping any patch.
//
//	"24.4.2" → "24.4"
//	"24.4"   → "24.4"
func majorMinor(v string) string {
	parts := strings.Split(strings.TrimSpace(v), ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return v
}

func main() {
	versionFlag := flag.String("version", "", `Download models only for a specific version (e.g. --version=24.4 or --version=24.4.2)`)
	flag.Parse()

	// If --version is given, filter the models list down to the release whose Version
	// matches the requested major.minor. The patch is ignored.
	// Examples:
	//   --version=24.4   → matches release "24.4"
	//   --version=24.4.2 → matches release "24.4"
	activeModels := models
	if *versionFlag != "" {
		want := majorMinor(*versionFlag)
		var filtered []ModelEntry
		for _, entry := range models {
			if entry.Version == want {
				filtered = append(filtered, entry)
			}
		}
		if len(filtered) == 0 {
			fmt.Printf("ERROR: no models found for version %q (major.minor: %q)\n", *versionFlag, want)
			fmt.Printf("Available versions in model list:\n")
			for _, r := range releases {
				fmt.Printf("  %s\n", r.Version)
			}
			os.Exit(1)
		}
		fmt.Printf("Version filter %q active — downloading %d/%d models\n",
			*versionFlag, len(filtered), len(models))
		activeModels = filtered
	}

	fmt.Println("Starting IOS-XR YANG model download...")

	// Track versions we've seen (keyed by dotted dir name)
	versionsCreated := make(map[string]bool)
	downloadedCount := 0
	skippedCount := 0

	for _, entry := range activeModels {
		dirVersion := entry.Version // e.g. "24.4"

		// Create version-specific directories if they don't exist
		if !versionsCreated[dirVersion] {
			modelsDir := filepath.Join(modelsBasePath, dirVersion)
			definitionsDir := filepath.Join(definitionsBasePath, dirVersion)

			if err := os.MkdirAll(modelsDir, 0755); err != nil {
				fmt.Printf("Failed to create models directory %s: %v\n", modelsDir, err)
				panic(err)
			}

			if err := os.MkdirAll(definitionsDir, 0755); err != nil {
				fmt.Printf("Failed to create definitions directory %s: %v\n", definitionsDir, err)
				panic(err)
			}

			fmt.Printf("Created directories for version %s\n", dirVersion)
			versionsCreated[dirVersion] = true
		}

		// Build the destination path with dotted version subdirectory
		destPath := filepath.Join(modelsBasePath, dirVersion, entry.Filename)

		// Skip if file already exists
		if _, err := os.Stat(destPath); err == nil {
			fmt.Printf("  ✓ Already exists: %s/%s\n", dirVersion, entry.Filename)
			skippedCount++
			continue
		}

		// Download the model
		err := downloadModel(destPath, entry.URL)
		if err != nil {
			fmt.Printf("  ✗ Failed to download %s/%s: %v\n", dirVersion, entry.Filename, err)
			panic(err)
		}

		fmt.Printf("  ✓ Downloaded: %s/%s\n", dirVersion, entry.Filename)
		downloadedCount++
	}

	fmt.Printf("\n=== Download Summary ===\n")
	fmt.Printf("Total models: %d\n", len(activeModels))
	fmt.Printf("Downloaded: %d\n", downloadedCount)
	fmt.Printf("Skipped (already exists): %d\n", skippedCount)
	fmt.Printf("Versions: %d\n", len(versionsCreated))

	fmt.Printf("\nDirectory structure:\n")
	fmt.Printf("  models/\n")
	for version := range versionsCreated {
		fmt.Printf("    %s/\n", version)
	}
	fmt.Printf("  definitions/\n")
	for version := range versionsCreated {
		fmt.Printf("    %s/\n", version)
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

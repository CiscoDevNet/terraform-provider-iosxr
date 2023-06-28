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
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-hostname-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-if-ip-address-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-if-vrf-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-interface-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-if-ipv4-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-statistics-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-router-static-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-if-service-policy-qos-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-if-bundle-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-l2vpn-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-key-chain-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-location-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-mpls-ldp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-mpls-te-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-policymap-classmap-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-pce-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-cfg-mibs-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-entity-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-system-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-bridgemib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-entity-state-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-entity-redundancy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-l2vpn-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-mibs-ifmib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-mpls-ldp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-ipv4-access-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-ipv6-access-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-access-list-datatypes.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-ipv6-prefix-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-ipv4-prefix-list-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-mpls-l3vpn-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-mibs-sensormib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-fru-ctrl-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-router-isis-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-router-bgp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-ntp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-segment-routing-srv6-datatypes.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-segment-routing-srv6-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-segment-routing-ms-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-infra-xtc-agent-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-config-copy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-traps-power-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-ethernet-oam-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-bfd-sbfd-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-mibs-rfmib-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-mpls-oam-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-segment-routing-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-router-bgp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-logging-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-logging-events-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-router-isis-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-router-ospf-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-snmp-server-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-vrf-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-ssh-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-route-policy-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-l2-ethernet-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-if-l2transport-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-statistics-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/cisco-semver.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/ietf-inet-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/ietf-yang-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/tailf-cli-extensions.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/tailf-common.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/tailf-meta-extensions.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-banner-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-cdp-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xr/761/Cisco-IOS-XR-um-lldp-cfg.yang",
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

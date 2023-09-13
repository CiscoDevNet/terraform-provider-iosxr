resource "iosxr_gnmi" "hostname" {
  path = "openconfig-system:/system/config"
  attributes = {
    hostname = "ROUTER-1"
  }
}

resource "iosxr_gnmi" "vrf" {
  path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]"
  attributes = {
    vrf-name    = "VRF1"
    description = "My Desc"
    "vpn/id"    = "1:1"
  }
  lists = [
    {
      name = "address-family/ipv4/unicast/Cisco-IOS-XR-um-router-bgp-cfg:import/route-target/ip-addresse-rts/ip-address-rt"
      key  = "ip-address,index,stitching"
      items = [
        {
          ip-address = "1.1.1.1"
          index      = "1"
          stitching  = "true"
        }
      ]
    }
  ]
}

resource "iosxr_gnmi" "domain_host" {
  path = "Cisco-IOS-XR-um-domain-cfg:/domain/ipv4/hosts/host[host-name=abc.cisco.com]"
  attributes = {
    "host-name" = "abc.cisco.com"
  }
  lists = [
    {
      name   = "ip-address"
      values = ["1.2.3.4"]
    }
  ]
}

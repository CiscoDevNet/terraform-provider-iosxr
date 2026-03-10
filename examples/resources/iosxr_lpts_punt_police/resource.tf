resource "iosxr_lpts_punt_police" "example" {
  mcast_rate = 1000
  bcast_rate = 1000
  protocol_arp_rate = 2000
  protocol_cdp_rate = 2000
  protocol_lacp_rate = 2000
  protocol_lldp_rate = 2000
  protocol_ssfp_rate = 2000
  protocol_ipv6_nd_proxy_rate = 2000
  domains = [
    {
      domain_name = "DOMAIN1"
      mcast_rate = 1500
      bcast_rate = 1500
      protocol_arp_rate = 2500
      protocol_cdp_rate = 2500
      protocol_lacp_rate = 2500
      protocol_lldp_rate = 2500
      protocol_ssfp_rate = 2500
      protocol_ipv6_nd_proxy_rate = 2500
    }
  ]
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
      mcast_rate = 800
      bcast_rate = 800
    }
  ]
}

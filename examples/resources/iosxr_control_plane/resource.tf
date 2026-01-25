resource "iosxr_control_plane" "example" {
  mgmt_inband_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/0"
      ssh = true
        ssh_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        ssh_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        ssh_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        ssh_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      telnet = true
        telnet_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        telnet_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        telnet_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        telnet_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      snmp = true
        snmp_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        snmp_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        snmp_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        snmp_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      tftp = true
        tftp_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        tftp_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        tftp_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        tftp_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      http = true
        http_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        http_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
      xml = true
        xml_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        xml_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        xml_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        xml_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      netconf = true
        netconf_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        netconf_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        netconf_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        netconf_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      allow_all = true
        allow_all_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        allow_all_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        allow_all_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        allow_all_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
    }
  ]
  mgmt_oob_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
      ssh = true
        ssh_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        ssh_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        ssh_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        ssh_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      telnet = true
        telnet_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        telnet_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        telnet_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        telnet_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      snmp = true
        snmp_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        snmp_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        snmp_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        snmp_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      tftp = true
        tftp_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        tftp_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        tftp_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        tftp_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      http = true
        http_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        http_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
      xml = true
        xml_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        xml_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        xml_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        xml_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      netconf = true
        netconf_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        netconf_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        netconf_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        netconf_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
      allow_all = true
        allow_all_ipv4_prefixes = [
          {
            address = "192.168.1.0"
            length = 24
          }
        ]
        allow_all_ipv4_hosts = [
          {
            address = "192.168.1.1"
          }
        ]
        allow_all_ipv6_prefixes = [
          {
            address = "2001:db8:1:1::"
            length = 64
          }
        ]
        allow_all_ipv6_hosts = [
          {
            address = "2001:db8:1:1::1"
          }
        ]
    }
  ]
  mgmt_oob_vrf = "VRF1"
}

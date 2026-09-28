resource "iosxr_l2vpn" "example" {
  autodiscovery_bgp_signaling_protocol_bgp_mtu_mismatch_ignore = true
  capability_high_mode                                         = true
  description                                                  = "My L2VPN Description"
  flexible_xconnect_service_vlan_aware_evis = [
    {
      interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/2.200"
        }
      ]
      vpn_id = 200
    }
  ]
  flexible_xconnect_service_vlan_unaware = [
    {
      interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/1.100"
        }
      ]
      neighbor_evpn_evis = [
        {
          remote_ac_id = 1000
          vpn_id       = 100
        }
      ]
      service_name = "XC-1"
    }
  ]
  ignore_mtu_mismatch                            = true
  ignore_mtu_mismatch_ad                         = true
  load_balancing_flow_src_dst_ip                 = true
  load_balancing_flow_src_dst_mac                = false
  logging_bridge_domain                          = true
  logging_nsr                                    = true
  logging_pseudowire                             = true
  logging_pwhe_replication_disable               = true
  logging_vfi                                    = true
  mac_limit_threshold                            = 50
  neighbors_all_ldp_flap                         = true
  pw_grouping                                    = true
  pw_oam_refresh_transmit                        = 20
  pw_routing_bgp_rd_four_byte_as_assigned_number = 1
  pw_routing_bgp_rd_four_byte_as_number          = 65536
  pw_routing_global_id                           = 100
  pw_status_disable                              = true
  redundancy_iccp_groups = [
    {
      group_number = 100
      interfaces = [
        {
          interface_name    = "Bundle-Ether20"
          mac_flush_stp_tcn = true
          primary_vlan      = "10-15"
          recovery_delay    = 60
          secondary_vlan    = "20-25"
        }
      ]
      multi_homing_node_id = 1
    }
  ]
  router_id                          = "1.2.3.4"
  snmp_mib_interface_format_external = true
  snmp_mib_pseudowire_statistics     = true
  tcn_propagation                    = true
}

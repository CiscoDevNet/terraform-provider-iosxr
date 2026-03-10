resource "iosxr_l2vpn" "example" {
  description = "My L2VPN Description"
  router_id   = "1.2.3.4"
  redundancy_iccp_groups = [
    {
      group_number = 100
      interfaces = [
        {
          interface_name    = "Bundle-Ether20"
          primary_vlan      = "10-15"
          secondary_vlan    = "20-25"
          mac_flush_stp_tcn = true
          recovery_delay    = 60
        }
      ]
      multi_homing_node_id = 1
    }
  ]
  flexible_xconnect_service_vlan_unaware = [
    {
      service_name = "XC-1"
      interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/1.100"
        }
      ]
      neighbor_evpn_evis = [
        {
          vpn_id       = 100
          remote_ac_id = 1000
        }
      ]
    }
  ]
  flexible_xconnect_service_vlan_aware_evis = [
    {
      vpn_id = 200
      interfaces = [
        {
          interface_name = "GigabitEthernet0/0/0/2.200"
        }
      ]
    }
  ]
  ignore_mtu_mismatch                                          = true
  ignore_mtu_mismatch_ad                                       = true
  pw_status_disable                                            = true
  load_balancing_flow_src_dst_mac                              = false
  load_balancing_flow_src_dst_ip                               = true
  capability_high_mode                                         = true
  pw_oam_refresh_transmit                                      = 20
  tcn_propagation                                              = true
  pw_grouping                                                  = true
  neighbors_all_ldp_flap                                       = true
  mac_limit_threshold                                          = 50
  logging_pseudowire                                           = true
  logging_bridge_domain                                        = true
  logging_vfi                                                  = true
  logging_nsr                                                  = true
  logging_pwhe_replication_disable                             = true
  autodiscovery_bgp_signaling_protocol_bgp_mtu_mismatch_ignore = true
  pw_routing_global_id                                         = 100
  pw_routing_bgp_rd_four_byte_as_number                        = 65536
  pw_routing_bgp_rd_four_byte_as_assigned_number               = 1
  snmp_mib_interface_format_external                           = true
  snmp_mib_pseudowire_statistics                               = true
}

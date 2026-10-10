resource "iosxr_l2vpn_xconnect_group" "example" {
  group_name = "P2P"
  mp2mps = [
    {
      autodiscovery_bgp                        = true
      autodiscovery_bgp_rd_four_byte_as_index  = 100
      autodiscovery_bgp_rd_four_byte_as_number = 65536
      autodiscovery_bgp_route_policy_export    = "EXPORT_POLICY"
      autodiscovery_bgp_route_target_export_four_byte_as_format = [
        {
          assigned_number     = 200
          four_byte_as_number = 65536
        }
      ]
      autodiscovery_bgp_route_target_import_four_byte_as_format = [
        {
          assigned_number     = 100
          four_byte_as_number = 65536
        }
      ]
      autodiscovery_bgp_signaling_protocol_bgp_ce_ids = [
        {
          interfaces = [
            {
              interface_name = "GigabitEthernet0/0/0/1"
              remote_ce_ids = [
                {
                  remote_ce_id_value = 20
                }
              ]
            }
          ]
          local_ce_id_value         = 10
          vpws_seamless_integration = true
        }
      ]
      autodiscovery_bgp_signaling_protocol_bgp_ce_range                       = 11
      autodiscovery_bgp_signaling_protocol_bgp_load_balancing_flow_label_both = true
      control_word_disable                                                    = true
      instance_name                                                           = "MP2MP1"
      l2_encapsulation                                                        = "ethernet"
      mtu                                                                     = 1500
      shutdown                                                                = false
      vpn_id                                                                  = 100
    }
  ]
  p2ps = [
    {
      description = "My P2P Description"
      evpn_target_neighbors_segment_routing = [
        {
          remote_ac_id                 = 8000
          segment_routing_srv6_locator = "LOC12"
          source                       = 7001
          vpn_id                       = 7000
        }
      ]
      interfaces = [
        {
          interface_name = "Bundle-Ether11"
        }
      ]
      p2p_xconnect_name = "XC"
    }
  ]
}

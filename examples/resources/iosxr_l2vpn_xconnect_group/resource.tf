resource "iosxr_l2vpn_xconnect_group" "example" {
  group_name = "P2P"
  p2ps = [
    {
      p2p_xconnect_name = "XC"
      description       = "My P2P Description"
      interfaces = [
        {
          interface_name = "Bundle-Ether11"
        }
      ]
      evpn_target_neighbors_segment_routing = [
        {
          vpn_id                       = 7000
          remote_ac_id                 = 8000
          source                       = 7001
          segment_routing_srv6_locator = "LOC12"
        }
      ]
    }
  ]
  mp2mps = [
    {
      instance_name                            = "MP2MP1"
      vpn_id                                   = 100
      mtu                                      = 1500
      shutdown                                 = false
      l2_encapsulation                         = "ethernet"
      control_word_disable                     = true
      autodiscovery_bgp                        = true
      autodiscovery_bgp_rd_four_byte_as_number = 65536
      autodiscovery_bgp_rd_four_byte_as_index  = 100
      autodiscovery_bgp_route_target_import_four_byte_as_format = [
        {
          four_byte_as_number = 65536
          assigned_number     = 100
        }
      ]
      autodiscovery_bgp_route_target_export_four_byte_as_format = [
        {
          four_byte_as_number = 65536
          assigned_number     = 200
        }
      ]
      autodiscovery_bgp_signaling_protocol_bgp_ce_ids = [
        {
          local_ce_id_value = 10
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
          vpws_seamless_integration = true
        }
      ]
      autodiscovery_bgp_signaling_protocol_bgp_ce_range                       = 11
      autodiscovery_bgp_signaling_protocol_bgp_load_balancing_flow_label_both = true
      autodiscovery_bgp_route_policy_export                                   = "EXPORT_POLICY"
    }
  ]
}

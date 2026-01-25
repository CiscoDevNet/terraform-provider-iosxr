resource "iosxr_l2vpn_bridge_group_bridge_domain_vfi" "example" {
  bridge_group_name                        = "BG123"
  bridge_domain_name                       = "BD123"
  vfi_name                                 = "VFI1"
  vpn_id                                   = 100
  shutdown                                 = true
  autodiscovery_bgp                        = true
  autodiscovery_bgp_rd_four_byte_as_number = 65536
  autodiscovery_bgp_rd_four_byte_as_index  = 1
  autodiscovery_bgp_route_target_import_ipv4_address_format = [
    {
      ipv4_address    = "10.0.0.1"
      assigned_number = 100
    }
  ]
  autodiscovery_bgp_route_target_export_ipv4_address_format = [
    {
      ipv4_address    = "10.0.0.1"
      assigned_number = 100
    }
  ]
  autodiscovery_bgp_control_word                                                    = true
  autodiscovery_bgp_signaling_protocol_bgp                                          = true
  autodiscovery_bgp_signaling_protocol_bgp_ve_id                                    = 100
  autodiscovery_bgp_signaling_protocol_bgp_ve_range                                 = 50
  autodiscovery_bgp_signaling_protocol_bgp_load_balancing_flow_label_both           = true
  autodiscovery_bgp_signaling_protocol_bgp_load_balancing_flow_label_static         = true
  autodiscovery_bgp_signaling_protocol_ldp                                          = true
  autodiscovery_bgp_signaling_protocol_ldp_vpls_id_ipv4_address                     = "10.0.0.1"
  autodiscovery_bgp_signaling_protocol_ldp_vpls_id_ipv4_address_index               = 1
  autodiscovery_bgp_signaling_protocol_ldp_vpls_id_load_balancing_flow_label_both   = true
  autodiscovery_bgp_signaling_protocol_ldp_vpls_id_load_balancing_flow_label_static = true
  multicast_p2mp                                                                    = true
  multicast_p2mp_transport_rsvp_te                                                  = true
  multicast_p2mp_transport_rsvp_te_attribute_set_p2mp_te                            = "TE_ATTRIBUTE_SET_1"
  multicast_p2mp_signaling_protocol_bgp                                             = true
  neighbors = [
    {
      address = "10.1.1.1"
      pw_id   = 1000
      static_mac_addresses = [
        {
          mac_address = "aa:bb:cc:dd:ee:01"
        }
      ]
      mpls_static_label_local    = 16001
      mpls_static_label_remote   = 17001
      pw_class                   = "PW_CLASS_1"
      dhcp_ipv4_snooping_profile = "DHCP_PROFILE"
      igmp_snooping_profile      = "IGMP_PROFILE"
      mld_snooping_profile       = "MLD_PROFILE"
    }
  ]
}

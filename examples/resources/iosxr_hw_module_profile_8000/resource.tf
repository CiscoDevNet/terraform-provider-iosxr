resource "iosxr_hw_module_profile_8000" "example" {
  profile_tcam_fib_ipv4_unicast_percent = 50
  profile_tcam_fib_ipv6_unicast_percent = 50
  profile_tcam_format_access_list_ipv4_src_addr = true
  profile_tcam_format_access_list_ipv4_dst_addr = true
  profile_tcam_format_access_list_ipv4_src_port = true
  profile_tcam_format_access_list_ipv4_dst_port = true
  profile_tcam_format_access_list_ipv4_proto = true
  profile_tcam_format_access_list_ipv4_precedence = true
  profile_tcam_format_access_list_ipv4_ttl_match = true
  profile_tcam_format_access_list_ipv4_tcp_flags = true
  profile_tcam_format_access_list_ipv4_frag_bit = true
  profile_tcam_format_access_list_ipv4_packet_len = true
  profile_tcam_format_access_list_ipv4_fragment_offset = true
  profile_tcam_format_access_list_ipv6_src_addr = true
  profile_tcam_format_access_list_ipv6_dst_addr = true
  profile_tcam_format_access_list_ipv6_dst_port = true
  profile_tcam_format_access_list_ipv6_next_hdr = true
  profile_tcam_format_access_list_ipv6_traffic_class = true
  profile_tcam_format_access_list_ipv6_frag_bit = true
  profile_tcam_format_access_list_ipv6_tcp_flags = true
  profile_tcam_format_access_list_ipv6_packet_len = true
  profile_qos_voq_mode_fair_eight = true
  profile_qos_l2_mode = "L3"
  profile_qos_low_latency_mode = "1"
  profile_qos_intra_npu_over_fabric = "disable"
  profile_qos_qos_stats_push_collection = true
  profile_qos_high_water_marks = true
  profile_cef_dark_bw = "enable"
  profile_cef_sropt = "enable"
  profile_cef_bgplu = "enable"
  profile_cef_cbf = "enable"
  profile_cef_cbf_forward_class_list = [0]
  profile_cef_ipv6_hop_limit = "punt"
  profile_cef_lpts_acl = true
  profile_cef_lpts_pifib_entry_counters = 256
  profile_cef_vxlan_ipv6_tnl_scale = true
  profile_cef_mplsoudp_scale = true
  profile_cef_stats_label_app_default = "dynamic"
  profile_cef_ttl_tunnel_ip_decrement = "disable"
  profile_cef_te_tunnel_highscale_no_ldp_over_te = true
  profile_cef_te_tunnel_highscale_ldp_over_te_no_sr_over_srte = true
  profile_cef_te_tunnel_label_over_te_counters = true
  profile_cef_ip_redirect = "enable"
  profile_cef_unipath_surpf_enable = true
  profile_cef_source_rtbh_enable = true
  profile_encap_exact_interfaces = [
    {
      interface_name = "FourHundredGigE0/0/0/0"
    }
  ]
  profile_encap_exact_locations = [
    {
      location_name = "0/RP0/CPU0"
    }
  ]
  profile_encap_exact_locations_all = true
  profile_encap_exact_locations_all_virtual = true
  profile_stats_voqs_sharing_counters = "1"
  profile_stats_no_bvi_ingress = true
  profile_stats_acl_permit = true
  profile_bw_threshold = "80"
  profile_gue_udp_dest_port_ipv4 = 7000
  profile_gue_udp_dest_port_ipv6 = 8000
  profile_gue_udp_dest_port_mpls = 9000
  profile_npu_buffer_extended_locations = [
    {
      location_name = "0/RP0/CPU0"
      bandwidth_congestion_detection_enable = true
      bandwidth_congestion_protect_enable = true
    }
  ]
  profile_l2fib_pw_stats = true
  profile_l2fib_bridge_flush_convergence = true
  profile_l2fib_higher_scale = true
  profile_route_scale_ipv6_unicast_connected_prefix_high = true
  profile_flowspec_ipv6_packet_len_enable = true
}

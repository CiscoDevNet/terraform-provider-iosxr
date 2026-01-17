resource "iosxr_pce" "example" {
  address_ipv4 = "77.77.77.1"
  address_ipv6 = "2001:db8::1"
  state_sync_ipv4s = [
    {
      address = "100.100.100.11"
    }
  ]
  state_sync_ipv6s = [
    {
      address = "2001:db8::2"
    }
  ]
  tcp_buffer_size                      = 256000
  tcp_ao_keychain_name                 = "KEY_CHAIN_1"
  tcp_ao_include_tcp_options           = true
  tcp_ao_accept_ao_mismatch_connection = true
  disjoint_path_maximum_attempts       = 5
  disjoint_path_group_ids = [
    {
      group_id                                = 10
      link_disjoint                           = true
      link_disjoint_strict                    = true
      link_disjoint_lsp_one_pcc_address_type  = "ipv4"
      link_disjoint_lsp_one_pcc_ip_address    = "100.100.100.10"
      link_disjoint_lsp_one_pcc_lsp_name      = "LSP_10"
      link_disjoint_lsp_one_pcc_shortest_path = true
      link_disjoint_lsp_one_pcc_exclude_srlg  = 10
      link_disjoint_lsp_two_pcc_address_type  = "ipv4"
      link_disjoint_lsp_two_pcc_ip_address    = "100.100.100.11"
      link_disjoint_lsp_two_pcc_lsp_name      = "LSP_11"
      link_disjoint_lsp_two_pcc_exclude_srlg  = 11
      link_disjoint_sub_ids = [
        {
          sub_id                    = 12
          strict                    = true
          lsp_one_pcc_address_type  = "ipv4"
          lsp_one_pcc_ip_address    = "100.100.100.12"
          lsp_one_pcc_lsp_name      = "LSP_12"
          lsp_one_pcc_shortest_path = true
          lsp_one_pcc_exclude_srlg  = 12
          lsp_two_pcc_address_type  = "ipv4"
          lsp_two_pcc_ip_address    = "200.200.200.13"
          lsp_two_pcc_lsp_name      = "LSP_13"
          lsp_two_pcc_exclude_srlg  = 13
        }
      ]
      node_disjoint                           = true
      node_disjoint_strict                    = true
      node_disjoint_lsp_one_pcc_address_type  = "ipv4"
      node_disjoint_lsp_one_pcc_ip_address    = "100.100.100.14"
      node_disjoint_lsp_one_pcc_lsp_name      = "LSP_14"
      node_disjoint_lsp_one_pcc_shortest_path = true
      node_disjoint_lsp_one_pcc_exclude_srlg  = 14
      node_disjoint_lsp_two_pcc_address_type  = "ipv4"
      node_disjoint_lsp_two_pcc_ip_address    = "100.100.100.15"
      node_disjoint_lsp_two_pcc_lsp_name      = "LSP_15"
      node_disjoint_lsp_two_pcc_exclude_srlg  = 15
      node_disjoint_sub_ids = [
        {
          sub_id                    = 16
          strict                    = true
          lsp_one_pcc_address_type  = "ipv4"
          lsp_one_pcc_ip_address    = "100.100.100.16"
          lsp_one_pcc_lsp_name      = "LSP_16"
          lsp_one_pcc_shortest_path = true
          lsp_one_pcc_exclude_srlg  = 16
          lsp_two_pcc_address_type  = "ipv4"
          lsp_two_pcc_ip_address    = "100.100.100.17"
          lsp_two_pcc_lsp_name      = "LSP_17"
          lsp_two_pcc_exclude_srlg  = 17
        }
      ]
      srlg_disjoint                           = true
      srlg_disjoint_strict                    = true
      srlg_disjoint_lsp_one_pcc_address_type  = "ipv4"
      srlg_disjoint_lsp_one_pcc_ip_address    = "100.100.100.18"
      srlg_disjoint_lsp_one_pcc_lsp_name      = "LSP_18"
      srlg_disjoint_lsp_one_pcc_shortest_path = true
      srlg_disjoint_lsp_one_pcc_exclude_srlg  = 18
      srlg_disjoint_lsp_two_pcc_address_type  = "ipv4"
      srlg_disjoint_lsp_two_pcc_ip_address    = "100.100.100.19"
      srlg_disjoint_lsp_two_pcc_lsp_name      = "LSP_19"
      srlg_disjoint_lsp_two_pcc_exclude_srlg  = 19
      srlg_disjoint_sub_ids = [
        {
          sub_id                    = 20
          strict                    = true
          lsp_one_pcc_address_type  = "ipv4"
          lsp_one_pcc_ip_address    = "100.100.100.20"
          lsp_one_pcc_lsp_name      = "LSP_20"
          lsp_one_pcc_shortest_path = true
          lsp_one_pcc_exclude_srlg  = 20
          lsp_two_pcc_address_type  = "ipv4"
          lsp_two_pcc_ip_address    = "100.100.100.21"
          lsp_two_pcc_lsp_name      = "LSP_21"
          lsp_two_pcc_exclude_srlg  = 21
        }
      ]
      srlg_node_disjoint                           = true
      srlg_node_disjoint_strict                    = true
      srlg_node_disjoint_lsp_one_pcc_address_type  = "ipv4"
      srlg_node_disjoint_lsp_one_pcc_ip_address    = "100.100.100.22"
      srlg_node_disjoint_lsp_one_pcc_lsp_name      = "LSP_22"
      srlg_node_disjoint_lsp_one_pcc_shortest_path = true
      srlg_node_disjoint_lsp_one_pcc_exclude_srlg  = 22
      srlg_node_disjoint_lsp_two_pcc_address_type  = "ipv4"
      srlg_node_disjoint_lsp_two_pcc_ip_address    = "100.100.100.23"
      srlg_node_disjoint_lsp_two_pcc_lsp_name      = "LSP_23"
      srlg_node_disjoint_lsp_two_pcc_exclude_srlg  = 23
      srlg_node_disjoint_sub_ids = [
        {
          sub_id                    = 24
          strict                    = true
          lsp_one_pcc_address_type  = "ipv4"
          lsp_one_pcc_ip_address    = "100.100.100.24"
          lsp_one_pcc_lsp_name      = "LSP_24"
          lsp_one_pcc_shortest_path = true
          lsp_one_pcc_exclude_srlg  = 24
          lsp_two_pcc_address_type  = "ipv4"
          lsp_two_pcc_ip_address    = "100.100.100.25"
          lsp_two_pcc_lsp_name      = "LSP_25"
          lsp_two_pcc_exclude_srlg  = 25
        }
      ]
    }
  ]
  peer_ipv4s = [
    {
      address                              = "200.200.200.10"
      tcp_ao_keychain_name                 = "KEY_CHAIN_1"
      tcp_ao_include_tcp_options           = true
      tcp_ao_accept_ao_mismatch_connection = true
    }
  ]
  peer_ipv6s = [
    {
      address                              = "2001:db8::10"
      tcp_ao_keychain_name                 = "KEY_CHAIN_1"
      tcp_ao_include_tcp_options           = true
      tcp_ao_accept_ao_mismatch_connection = true
    }
  ]
  netconf_ssh_user               = "netconf-user"
  netconf_ssh_password_encrypted = "00071A150754"
  api_authentication_digest      = true
  api_sibling_ipv4               = "100.100.100.2"
  api_vrf                        = "VRF1"
  api_users = [
    {
      user_name          = "rest-user"
      password_encrypted = "00141215174C04140B"
    }
  ]
  api_ipv4_address                               = "100.100.100.3"
  api_ipv6_address                               = "2001:db8::1"
  timers_reoptimization                          = 600
  timers_keepalive                               = 60
  timers_minimum_peer_keepalive                  = 40
  timers_peer_zombie                             = 120
  timers_init_verify_restart                     = 80
  timers_init_verify_switchover                  = 120
  timers_init_verify_startup                     = 360
  backoff_ratio                                  = 10
  backoff_difference                             = 10
  backoff_threshold                              = 10
  logging_no_path                                = true
  logging_fallback                               = true
  logging_pcep_pcerr_received                    = true
  logging_pcep_api_send_queue_congestion_disable = true
  logging_pcep_disjointness_status               = true
  segment_routing_strict_sid_only                = true
  srte_affinity_bitmaps = [
    {
      affinity_color_name   = "COLOR_1"
      affinity_bit_position = 1
    }
  ]
  srte_segment_lists = [
    {
      segment_list_name = "SEGMENT_LIST_1"
      indexes = [
        {
          index_number = 1
          mpls_label   = 16100
        }
      ]
    }
  ]
  srte_ipv4_peers = [
    {
      address = "100.100.100.26"
      policies = [
        {
          policy_name                     = "PEER_POLICY1"
          candidate_paths_append_sid_mpls = 16100
          candidate_paths_preferences = [
            {
              preference_id               = 100
              dynamic_mpls                = true
              dynamic_metric_type_latency = true
              dynamic_metric_sid_limit    = 5
              explicit_segment_list_names = [
                {
                  segment_list_name = "SEGMENT_LIST_1"
                }
              ]
              constraints_segments_sid_algorithm                  = 128
              constraints_segments_protection_protected_preferred = true
            }
          ]
          candidate_paths_affinity_include_any_colors = [
            {
              affinity_color_name = "COLOR_1"
            }
          ]
          candidate_paths_affinity_include_all_colors = [
            {
              affinity_color_name = "COLOR_2"
            }
          ]
          candidate_paths_affinity_exclude_colors = [
            {
              affinity_color_name = "COLOR_3"
            }
          ]
          color                    = 100
          end_point_ipv4           = "100.100.100.1"
          binding_sid_mpls         = 24200
          shutdown                 = false
          profile_id               = 5
          path_selection_protected = true
        }
      ]
    }
  ]
  srte_cspf_anycast_sid_inclusion = true
  srte_cspf_sr_native             = true
  srte_cspf_sr_native_force       = true
  srte_p2mp_endpoint_sets = [
    {
      endpoint_set_name = "ENDPOINT_SET_1"
      ipv4s = [
        {
          address = "100.100.100.30"
        }
      ]
    }
  ]
  srte_p2mp_policies = [
    {
      policy_name      = "P2MP_POLICY_1"
      color            = 200
      endpoint_set     = "ENDPOINT_SET_1"
      source_ipv4      = "100.100.100.1"
      shutdown         = false
      fast_reroute_lfa = true
      treesid_mpls     = 24200
      candidate_paths_constraints_affinity_include_any_colors = [
        {
          affinity_color_name = "COLOR_1"
        }
      ]
      candidate_paths_constraints_affinity_include_all_colors = [
        {
          affinity_color_name = "COLOR_2"
        }
      ]
      candidate_paths_constraints_affinity_exclude_colors = [
        {
          affinity_color_name = "COLOR_3"
        }
      ]
      candidate_paths_preferences = [
        {
          preference_id               = 100
          dynamic                     = true
          dynamic_metric_type_latency = true
        }
      ]
    }
  ]
  srte_p2mp_timers_reoptimization = 300
  srte_p2mp_timers_cleanup        = 60
  srte_p2mp_label_range_min       = 16000
  srte_p2mp_label_range_max       = 17000
  srte_p2mp_multipath_disable     = true
  srte_p2mp_fast_reroute_lfa      = true
  srte_p2mp_frr_node_set_from_ipv4s = [
    {
      address = "100.100.100.31"
    }
  ]
  srte_p2mp_frr_node_set_to_ipv4s = [
    {
      address = "100.100.100.32"
    }
  ]
  peer_filter_ipv4_access_list     = "ACL_1"
  hierarchical_underlay_enable_all = true
}

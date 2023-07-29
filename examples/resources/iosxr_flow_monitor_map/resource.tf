resource "iosxr_flow_monitor_map" "example" {
  name = "monitor_map1"
  exporters = [
    {
      name = "exporter1"
    }
  ]
  option_outphysint                          = true
  option_filtered                            = true
  option_bgpattr                             = true
  option_outbundlemember                     = true
  record_ipv4_destination                    = true
  record_ipv4_destination_tos                = true
  record_ipv4_as                             = true
  record_ipv4_protocol_port                  = true
  record_ipv4_prefix                         = true
  record_ipv4_source_prefix                  = true
  record_ipv4_destination_prefix             = true
  record_ipv4_as_tos                         = true
  record_ipv4_protocol_port_tos              = true
  record_ipv4_prefix_tos                     = true
  record_ipv4_source_prefix_tos              = true
  record_ipv4_destination_prefix_tos         = true
  record_ipv4_prefix_port                    = true
  record_ipv4_bgp_nexthop_tos                = true
  record_ipv4_peer_as                        = true
  record_ipv4_gtp                            = true
  record_ipv6_destination                    = true
  record_ipv6_peer_as                        = true
  record_ipv6_gtp                            = true
  record_mpls_ipv4_fields                    = true
  record_mpls_ipv6_fields                    = true
  record_mpls_ipv4_ipv6_fields               = true
  record_mpls_labels                         = 2
  record_map_t                               = true
  record_sflow                               = true
  record_datalink_record                     = true
  record_default_rtp                         = true
  record_default_mdi                         = true
  cache_entries                              = 5000
  cache_timeout_active                       = 1
  cache_timeout_inactive                     = 0
  cache_timeout_update                       = 1
  cache_timeout_rate_limit                   = 5000
  cache_permanent                            = true
  cache_immediate                            = true
  hw_cache_timeout_inactive                  = 50
  sflow_options_extended_router              = true
  sflow_options_extended_gateway             = true
  sflow_options_extended_ipv4_tunnel_egress  = true
  sflow_options_extended_ipv6_tunnel_egress  = true
  sflow_options_if_counters_polling_interval = 5
  sflow_options_sample_header_size           = 128
  sflow_options_input_ifindex                = "physical"
  sflow_options_output_ifindex               = "physical"
}

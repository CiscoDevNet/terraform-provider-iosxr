resource "iosxr_flow_monitor_map" "example" {
  cache_entries            = 5000
  cache_immediate          = true
  cache_timeout_active     = 1
  cache_timeout_inactive   = 0
  cache_timeout_rate_limit = 5000
  cache_timeout_update     = 1
  exporters = [
    {
      name = "exporter_map1"
    }
  ]
  hw_cache_timeout_inactive                  = 50
  name                                       = "monitor_map1"
  option_bgpattr                             = true
  option_filtered                            = true
  option_outbundlemember                     = true
  option_outphysint                          = true
  record_mpls_labels                         = 2
  sflow_options                              = true
  sflow_options_extended_gateway             = true
  sflow_options_extended_ipv4_tunnel_egress  = true
  sflow_options_extended_ipv6_tunnel_egress  = true
  sflow_options_extended_router              = true
  sflow_options_if_counters_polling_interval = 5
  sflow_options_input_ifindex                = "physical"
  sflow_options_output_ifindex               = "physical"
  sflow_options_sample_header_size           = 128
}

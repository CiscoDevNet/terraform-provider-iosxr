resource "iosxr_router_bgp_session_group" "example" {
  advertisement_interval_seconds                 = 10
  allowas_in                                     = 3
  as_number                                      = "65001"
  as_override                                    = "enable"
  as_path_loopcheck_out                          = "enable"
  bfd_fast_detect                                = true
  bfd_fast_detect_strict_mode_negotiate_override = true
  bfd_minimum_interval                           = 10
  bfd_multiplier                                 = 4
  bmp_activate_servers = [
    {
      server_number = 1
    }
  ]
  capability_additional_paths_receive           = true
  capability_additional_paths_send              = true
  capability_suppress_all                       = true
  capability_suppress_extended_nexthop_encoding = true
  capability_suppress_four_byte_as              = true
  cluster_id_32bit_format                       = 100010
  dampening                                     = "enable"
  default_policy_action_in                      = "reject"
  default_policy_action_out                     = "reject"
  description                                   = "Session Group Description"
  dscp                                          = "ef"
  egress_engineering                            = true
  fast_fallover                                 = true
  idle_watch_time                               = 240
  internal_vpn_client                           = true
  local_address                                 = "192.168.1.1"
  log_message_in_size                           = 256
  log_message_out_size                          = 256
  log_neighbor_changes_detail                   = true
  maximum_peers                                 = 1000
  name                                          = "SGROUP1"
  password                                      = "030752180500"
  receive_buffer_size                           = 1024
  receive_buffer_size_read                      = 1024
  remote_as                                     = "65001"
  send_buffer_size                              = 4096
  send_buffer_size_write                        = 4096
  session_open_mode                             = "active-only"
  shutdown                                      = false
  tcp_mss_value                                 = 1460
  tcp_mtu_discovery                             = true
  timers_holdtime                               = 30
  timers_holdtime_minimum_acceptable_holdtime   = 30
  timers_keepalive_interval                     = 10
  update_in_error_handling_treat_as_withdraw    = "enable"
  update_in_labeled_unicast_equivalent          = true
  update_source                                 = "Loopback0"
}

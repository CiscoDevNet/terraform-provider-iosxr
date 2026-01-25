resource "iosxr_router_bgp_session_group" "example" {
  as_number                      = "65001"
  name                           = "SGROUP1"
  remote_as                      = "65001"
  maximum_peers                  = 1000
  as_path_loopcheck_out          = "enable"
  dampening                      = "enable"
  as_override                    = "enable"
  advertisement_interval_seconds = 10
  description                    = "Session Group Description"
  internal_vpn_client            = true
  tcp_mss_value                  = 1460
  tcp_mtu_discovery              = true
  bmp_activate_servers = [
    {
      server_number = 1
    }
  ]
  bfd_minimum_interval                           = 10
  bfd_multiplier                                 = 4
  bfd_fast_detect                                = true
  bfd_fast_detect_strict_mode_negotiate_override = true
  password                                       = "030752180500"
  receive_buffer_size                            = 1024
  receive_buffer_size_read                       = 1024
  send_buffer_size                               = 4096
  send_buffer_size_write                         = 4096
  shutdown                                       = false
  timers_keepalive_interval                      = 10
  timers_holdtime                                = 30
  timers_holdtime_minimum_acceptable_holdtime    = 30
  local_address                                  = "192.168.1.1"
  log_neighbor_changes_detail                    = true
  log_message_in_size                            = 256
  log_message_out_size                           = 256
  update_source                                  = "Loopback0"
  session_open_mode                              = "active-only"
  dscp                                           = "ef"
  capability_additional_paths_send               = true
  capability_additional_paths_receive            = true
  capability_suppress_all                        = true
  capability_suppress_extended_nexthop_encoding  = true
  capability_suppress_four_byte_as               = true
  cluster_id_32bit_format                        = 100010
  idle_watch_time                                = 240
  allowas_in                                     = 3
  egress_engineering                             = true
  default_policy_action_in                       = "reject"
  default_policy_action_out                      = "reject"
  fast_fallover                                  = true
  update_in_labeled_unicast_equivalent           = true
  update_in_error_handling_treat_as_withdraw     = "enable"
}

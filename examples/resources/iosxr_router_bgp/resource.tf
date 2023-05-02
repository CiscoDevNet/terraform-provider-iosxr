resource "iosxr_router_bgp" "example" {
  as_number                             = "65001"
  default_information_originate         = true
  default_metric                        = 125
  timers_bgp_keepalive_interval         = 5
  timers_bgp_holdtime                   = "20"
  bgp_router_id                         = "22.22.22.22"
  bgp_graceful_restart_graceful_reset   = true
  ibgp_policy_out_enforce_modifications = true
  bgp_log_neighbor_changes_detail       = true
  bfd_minimum_interval                  = 10
  bfd_multiplier                        = 4
  neighbors = [
    {
      neighbor_address                = "10.1.1.2"
      remote_as                       = "65002"
      description                     = "My Neighbor Description"
      use_neighbor_group              = "GROUP1"
      ignore_connected_check          = true
      ebgp_multihop_maximum_hop_count = 10
      bfd_minimum_interval            = 10
      bfd_multiplier                  = 4
      local_as                        = "65003"
      local_as_no_prepend             = true
      local_as_replace_as             = true
      local_as_dual_as                = true
      password                        = "12341C2713181F13253920"
      shutdown                        = false
      timers_keepalive_interval       = 5
      timers_holdtime                 = "20"
      update_source                   = "GigabitEthernet0/0/0/1"
      ttl_security                    = false
    }
  ]
  neighbor_groups = [
    {
      name          = "GROUP1"
      remote_as     = "65001"
      update_source = "Loopback0"
    }
  ]
}

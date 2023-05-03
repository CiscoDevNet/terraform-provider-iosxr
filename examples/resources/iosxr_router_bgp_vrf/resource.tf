resource "iosxr_router_bgp_vrf" "example" {
  as_number                     = "65001"
  vrf_name                      = "VRF2"
  rd_auto                       = false
  rd_ip_address_ipv4_address    = "14.14.14.14"
  rd_ip_address_index           = 3
  default_information_originate = true
  default_metric                = 125
  timers_bgp_keepalive_interval = 5
  timers_bgp_holdtime           = "20"
  bfd_minimum_interval          = 10
  bfd_multiplier                = 4
  neighbors = [
    {
      neighbor_address                = "10.1.1.2"
      remote_as                       = "65002"
      description                     = "My Neighbor Description"
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
}

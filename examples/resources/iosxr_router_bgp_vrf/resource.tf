resource "iosxr_router_bgp_vrf" "example" {
  as_number                     = "65001"
  vrf_name                      = "VRF2"
  default_information_originate = true
  default_metric                = 125
  rd_auto                       = true
  timers_bgp_keepalive_interval = 5
  timers_bgp_holdtime           = 20
  bgp_router_id                 = "22.22.22.22"
  bfd_minimum_interval          = 10
  bfd_multiplier                = 4
  neighbors = [
    {
      address                         = "10.1.1.2"
      remote_as                       = "65002"
      description                     = "My Neighbor Description"
      use_neighbor_group              = "GROUP1"
      advertisement_interval_seconds  = 10
      ignore_connected_check          = true
      ebgp_multihop_maximum_hop_count = 10
      bfd_minimum_interval            = 10
      bfd_multiplier                  = 4
      bfd_fast_detect                 = true
      bfd_fast_detect_strict_mode     = false
      bfd_fast_detect_disable         = false
      password                        = "12341C2713181F13253920"
      password_inheritance_disable    = false
      shutdown                        = false
      timers_keepalive_interval       = 10
      timers_holdtime                 = 20
      update_source                   = "GigabitEthernet0/0/0/1"
      ttl_security                    = false
    }
  ]
}

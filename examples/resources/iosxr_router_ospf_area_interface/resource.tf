resource "iosxr_router_ospf_area_interface" "example" {
  process_name                                       = "OSPF1"
  area_id                                            = "0"
  interface_name                                     = "GigabitEthernet0/0/0/1"
  network_broadcast                                  = false
  network_non_broadcast                              = false
  network_point_to_point                             = true
  network_point_to_multipoint                        = false
  cost                                               = 20
  priority                                           = 100
  passive_enable                                     = false
  passive_disable                                    = true
  fast_reroute_per_prefix_ti_lfa                     = true
  fast_reroute_per_prefix_tiebreaker_srlg_disjoint   = 22
  fast_reroute_per_prefix_tiebreaker_node_protecting = 33
}

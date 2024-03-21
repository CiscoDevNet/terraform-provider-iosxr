resource "iosxr_router_bgp_neighbor_address_family" "example" {
  as_number                                                 = "65001"
  neighbor_address                                          = "10.1.1.2"
  af_name                                                   = "vpnv4-unicast"
  import_stitching_rt_re_originate_stitching_rt             = true
  route_reflector_client                                    = true
  route_reflector_client_inheritance_disable                = true
  advertise_vpnv4_unicast_enable_re_originated_stitching_rt = true
  next_hop_self                                             = true
  next_hop_self_inheritance_disable                         = true
  encapsulation_type_srv6                                   = true
  route_policy_in                                           = "ROUTE_POLICY_1"
  route_policy_out                                          = "ROUTE_POLICY_1"
  soft_reconfiguration_inbound_always                       = true
  send_community_ebgp                                       = true
  send_community_ebgp_inheritance_disable                   = true
  maximum_prefix_limit                                      = 1248576
  maximum_prefix_threshold                                  = 80
  maximum_prefix_warning_only                               = true
  default_originate_route_policy                            = "ROUTE_POLICY_1"
  default_originate_inheritance_disable                     = true
}

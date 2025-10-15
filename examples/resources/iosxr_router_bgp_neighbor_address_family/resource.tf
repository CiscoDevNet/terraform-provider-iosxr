resource "iosxr_router_bgp_neighbor_address_family" "example" {
  as_number                                          = "65001"
  address                                            = "10.1.1.2"
  af_name                                            = "vpnv4-unicast"
  route_reflector_client                             = true
  route_reflector_client_inheritance_disable         = true
  advertise_vpnv4_unicast                            = true
  advertise_vpnv4_unicast_re_originated              = true
  advertise_vpnv4_unicast_re_originated_stitching_rt = true
  next_hop_self                                      = true
  next_hop_self_inheritance_disable                  = true
  encapsulation_type                                 = "srv6"
  route_policy_in                                    = "ROUTE_POLICY_1"
  route_policy_out                                   = "ROUTE_POLICY_1"
  soft_reconfiguration_inbound                       = true
  soft_reconfiguration_inbound_always                = true
  maximum_prefix_limit                               = 1248576
  maximum_prefix_threshold                           = 80
  maximum_prefix_warning_only                        = true
  default_originate                                  = true
  default_originate_route_policy                     = "ROUTE_POLICY_1"
}

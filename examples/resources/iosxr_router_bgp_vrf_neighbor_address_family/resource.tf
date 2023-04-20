resource "iosxr_router_bgp_vrf_neighbor_address_family" "example" {
  as_number                                     = "33333"
  vrf_name                                      = "VRF33"
  neighbor_address                              = "44.44.44.44"
  af_name                                       = "ipv4-unicast1"
  route_policy_retention_route_policy_name      = "ROUTE-POLICY1"
  route_policy_in                               = "true"
  route_policy_out                              = "true"
  default_originate_route_policy                = "POLICY11"
  next_hop_self_inheritance_disable             = true
  soft_reconfiguration_inbound_always           = true
  send_community_ebgp_inheritance_disable       = true
  remove_private_as_inheritance_disable         = true
  remove_private_as_inbound_inheritance_disable = true
}

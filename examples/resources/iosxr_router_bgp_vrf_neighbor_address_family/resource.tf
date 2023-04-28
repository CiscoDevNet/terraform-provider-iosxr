resource "iosxr_router_bgp_vrf_neighbor_address_family" "example" {
  as_number                               = "65001"
  vrf_name                                = "VRF1"
  neighbor_address                        = "10.1.1.2"
  af_name                                 = "ipv4-unicast"
  route_policy_in                         = "ROUTE_POLICY_1"
  route_policy_out                        = "ROUTE_POLICY_1"
  default_originate_route_policy          = "ROUTE_POLICY_1"
  next_hop_self_inheritance_disable       = true
  soft_reconfiguration_inbound_always     = true
  send_community_ebgp_inheritance_disable = true
  remove_private_as_inheritance_disable   = true
}

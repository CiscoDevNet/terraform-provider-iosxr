resource "iosxr_mpls_ldp_mldp" "example" {
  address_family = [
    {
      carrier_supporting_carrier        = true
      forwarding_recursive              = true
      forwarding_recursive_route_policy = "LDP_POLICY_1"
      make_before_break_delay           = 60
      make_before_break_delete_delay    = 40
      make_before_break_route_policy    = "LDP_POLICY_1"
      mofrr_enable                      = true
      mofrr_route_policy                = "LDP_POLICY_1"
      name                              = "ipv4"
      neighbors = [
        {
          neighbor_address          = "192.168.2.1"
          neighbor_route_policy_in  = "LDP_POLICY_1"
          neighbor_route_policy_out = "LDP_POLICY_1"
        }
      ]
      neighbors_route_policy_in  = "LDP_POLICY_1"
      neighbors_route_policy_out = "LDP_POLICY_1"
      recursive_fec_enable       = true
      recursive_fec_route_policy = "LDP_POLICY_1"
      rib_unicast_always         = true
      statics = [
        {
          lsp_address = "192.168.2.1"
          p2mp        = 5
        }
      ]
    }
  ]
  logging_internal      = true
  logging_notifications = true
}

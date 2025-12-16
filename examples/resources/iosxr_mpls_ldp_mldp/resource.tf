resource "iosxr_mpls_ldp_mldp" "example" {
  logging_notifications = true
  logging_internal      = true
  address_families = [
    {
      name = "ipv4"
      statics = [
        {
          lsp_address = "192.168.2.1"
          p2mp        = 5
        }
      ]
      make_before_break_delay        = 60
      make_before_break_delete_delay = 40
      make_before_break_route_policy = "LDP_POLICY_1"
      carrier_supporting_carrier     = true
      mofrr_enable                   = true
      mofrr_route_policy             = "LDP_POLICY_1"
      recursive_fec_enable           = true
      recursive_fec_route_policy     = "LDP_POLICY_1"
      neighbors_route_policy_in      = "LDP_POLICY_1"
      neighbors_route_policy_out     = "LDP_POLICY_1"
      neighbors = [
        {
          neighbor_address          = "192.168.2.1"
          neighbor_route_policy_in  = "LDP_POLICY_1"
          neighbor_route_policy_out = "LDP_POLICY_1"
        }
      ]
      forwarding_recursive              = true
      forwarding_recursive_route_policy = "LDP_POLICY_1"
      rib_unicast_always                = true
    }
  ]
  vrfs = [
    {
      vrf_name = "VRF1"
      address_families = [
        {
          name                              = "ipv4"
          make_before_break_delay           = 60
          make_before_break_delete_delay    = 40
          make_before_break_route_policy    = "LDP_POLICY_1"
          mofrr_enable                      = true
          mofrr_route_policy                = "LDP_POLICY_1"
          forwarding_recursive              = true
          forwarding_recursive_route_policy = "LDP_POLICY_1"
          rib_unicast_always                = true
        }
      ]
    }
  ]
}

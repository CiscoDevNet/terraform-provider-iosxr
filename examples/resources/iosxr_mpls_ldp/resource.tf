resource "iosxr_mpls_ldp" "example" {
  router_id = "1.2.3.4"
  address_families = [
    {
      af_name = "ipv4"
    }
  ]
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  capabilities_sac_ipv4_disable   = true
  capabilities_sac_ipv6_disable   = true
  capabilities_sac_fec128_disable = true
  capabilities_sac_fec129_disable = true
  mldp_logging_notifications      = true
  mldp_address_families = [
    {
      name                              = "ipv4"
      make_before_break_delay           = 30
      forwarding_recursive              = true
      forwarding_recursive_route_policy = "ROUTE_POLICY_1"
      recursive_fec                     = true
    }
  ]
  session_protection = true
}

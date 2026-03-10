resource "iosxr_mpls_ldp_vrf" "example" {
  vrf_name                                                 = "VRF1"
  router_id                                                = "1.2.3.4"
  session_downstream_on_demand_with                        = "ACL1"
  graceful_restart_helper_peer_maintain_on_local_reset_for = "ACL2"
  neighbors_password_encrypted                             = "060506324F41"
  neighbors = [
    {
      neighbor_address   = "192.168.2.1"
      label_space_id     = 0
      password_encrypted = "060506324F41"
    }
  ]
  address_family_ipv4                                                    = true
  address_family_ipv4_discovery_transport_address_ip                     = "192.168.1.1"
  address_family_ipv4_label_local_allocate_for_acl                       = "ACL1"
  address_family_ipv4_label_local_default_route                          = true
  address_family_ipv4_label_local_implicit_null_override_for             = "ACL1"
  address_family_ipv4_label_local_advertise_explicit_null                = true
  address_family_ipv4_label_local_advertise_explicit_null_for_acl        = "ACL1"
  address_family_ipv4_label_local_advertise_explicit_null_for_acl_to_acl = "ACL2"
  address_family_ipv4_label_local_advertise_to_neighbors = [
    {
      neighbor_address = "192.168.1.1"
      label_space_id   = 0
      for              = "ACL1"
    }
  ]
  address_family_ipv4_label_local_advertise_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  address_family_ipv4_label_local_advertise_disable = true
  address_family_ipv4_label_remote_accept_from_neighbors = [
    {
      neighbor_address = "192.168.1.2"
      label_space_id   = 0
      for              = "ACL1"
    }
  ]
  address_family_ipv6                                                    = true
  address_family_ipv6_discovery_transport_address_ip                     = "2001:db8::1"
  address_family_ipv6_label_local_allocate_for_acl                       = "ACL1"
  address_family_ipv6_label_local_default_route                          = true
  address_family_ipv6_label_local_implicit_null_override_for             = "ACL1"
  address_family_ipv6_label_local_advertise_explicit_null                = true
  address_family_ipv6_label_local_advertise_explicit_null_for_acl        = "ACL1"
  address_family_ipv6_label_local_advertise_explicit_null_for_acl_to_acl = "ACL2"
  address_family_ipv6_label_local_advertise_to_neighbors = [
    {
      neighbor_address = "192.168.1.2"
      label_space_id   = 0
      for              = "ACL1"
    }
  ]
  address_family_ipv6_label_local_advertise_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  address_family_ipv6_label_local_advertise_disable = true
  address_family_ipv6_label_remote_accept_from_neighbors = [
    {
      neighbor_address = "192.168.1.2"
      label_space_id   = 0
      for              = "ACL1"
    }
  ]
  interfaces = [
    {
      interface_name                                            = "GigabitEthernet0/0/0/1"
      address_family_ipv4                                       = true
      address_family_ipv4_discovery_transport_address_interface = true
      address_family_ipv6                                       = true
      address_family_ipv6_discovery_transport_address_ip        = "2001:db8::1"
    }
  ]
}

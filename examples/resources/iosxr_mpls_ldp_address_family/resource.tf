resource "iosxr_mpls_ldp_address_family" "example" {
  af_name                                            = "ipv4"
  discovery_targeted_hello_accept                    = true
  discovery_targeted_hello_accept_from               = "ACL1"
  discovery_transport_address_ipv4                   = "192.168.1.1"
  label_local_advertise_disable                      = true
  label_local_advertise_explicit_null                = true
  label_local_advertise_explicit_null_for_acl        = "ACL1"
  label_local_advertise_explicit_null_for_acl_to_acl = "ACL2"
  label_local_advertise_interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
    }
  ]
  label_local_advertise_to_neighbors = [
    {
      label_space_id   = 0
      neighbor_address = "192.168.1.1"
      for              = "ACL1"
    }
  ]
  label_local_allocate_for_host_routes   = true
  label_local_default_route              = true
  label_local_implicit_null_override_for = "ACL1"
  label_remote_accept_from_neighbors = [
    {
      label_space_id   = 0
      neighbor_address = "192.168.1.2"
      for              = "ACL1"
    }
  ]
  neighbor_ipv4_targeted = [
    {
      neighbor_address = "192.168.1.2"
    }
  ]
  neighbor_sr_policies = [
    {
      policy_name = "sr_te_policy_1"
      targeted    = true
    }
  ]
  redistribute_bgp_advertise_to = "ACL1"
  redistribute_bgp_as           = "65001"
  traffic_eng_auto_tunnel_mesh_groups = [
    {
      group_id = 1
    }
  ]
}

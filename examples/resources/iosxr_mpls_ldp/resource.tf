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
}

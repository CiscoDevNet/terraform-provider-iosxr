resource "iosxr_router_bgp_neighbor_group" "example" {
  as_number     = "65001"
  name          = "GROUP1"
  remote_as     = "65001"
  update_source = "Loopback0"
  address_families = [
    {
      af_name                             = "ipv4-labeled-unicast"
      soft_reconfiguration_inbound_always = true
    }
  ]
}

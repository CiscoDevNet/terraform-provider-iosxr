resource "iosxr_evpn" "example" {
  source_interface = "Loopback0"
  interfaces = [
    {
      interface_name            = "GigabitEthernet0/0/0/1"
      ethernet_segment_enable   = true
      ethernet_segment_esi_zero = "01.02.03.04.05.06.07.08.09"
    }
  ]
  segment_routing_srv6 = true
  segment_routing_srv6_locators = [
    {
      locator_name                        = "LOC1"
      usid_allocation_wide_local_id_block = true
    }
  ]
  segment_routing_srv6_usid_allocation_wide_local_id_block = true
}

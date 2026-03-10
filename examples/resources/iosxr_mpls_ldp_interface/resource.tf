resource "iosxr_mpls_ldp_interface" "example" {
  interface_name                 = "GigabitEthernet0/0/0/1"
  discovery_hello_holdtime       = 30
  discovery_hello_interval       = 3
  discovery_hello_dual_stack_tlv = "ipv4"
  discovery_quick_start_disable  = true
  igp_sync_delay_on_session_up   = 20
  address_family = [
    {
      af_name                        = "ipv4"
      discovery_transport_address_ip = "192.168.1.1"
      igp_auto_config_disable        = true
      mldp_disable                   = true
    }
  ]
}

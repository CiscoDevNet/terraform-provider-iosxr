resource "iosxr_router_isis" "example" {
  process_id = "P1"
  is_type    = "level-1"
  nets = [
    {
      net_id = "49.0001.2222.2222.2222.00"
    }
  ]
  address_families = [
    {
      af_name                       = "ipv4"
      saf_name                      = "unicast"
      mpls_ldp_auto_config          = false
      metric_style_narrow           = false
      metric_style_wide             = true
      metric_style_transition       = false
      router_id_ip_address          = "1.2.3.4"
      default_information_originate = true
    }
  ]
  interfaces = [
    {
      interface_name          = "GigabitEthernet0/0/0/1"
      circuit_type            = "level-1"
      hello_padding_disable   = true
      hello_padding_sometimes = false
      priority                = 10
      point_to_point          = false
      passive                 = false
      suppressed              = false
      shutdown                = false
    }
  ]
}

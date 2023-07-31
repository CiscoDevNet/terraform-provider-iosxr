resource "iosxr_router_isis" "example" {
  process_id                                                            = "P1"
  is_type                                                               = "level-1"
  set_overload_bit_on_startup_advertise_as_overloaded                   = true
  set_overload_bit_on_startup_advertise_as_overloaded_time_to_advertise = 10
  set_overload_bit_on_startup_wait_for_bgp                              = false
  set_overload_bit_advertise_external                                   = true
  set_overload_bit_advertise_interlevel                                 = true
  set_overload_bit_levels = [
    {
      level_id                                             = 1
      on_startup_advertise_as_overloaded                   = true
      on_startup_advertise_as_overloaded_time_to_advertise = 10
      on_startup_wait_for_bgp                              = false
      advertise_external                                   = true
      advertise_interlevel                                 = true
    }
  ]
  nsr                               = true
  nsf_cisco                         = true
  nsf_ietf                          = false
  nsf_lifetime                      = 10
  nsf_interface_timer               = 5
  nsf_interface_expires             = 2
  log_adjacency_changes             = true
  lsp_gen_interval_maximum_wait     = 5000
  lsp_gen_interval_initial_wait     = 50
  lsp_gen_interval_secondary_wait   = 200
  lsp_refresh_interval              = 65000
  max_lsp_lifetime                  = 65535
  lsp_password_keychain             = "ISIS-KEY"
  distribute_link_state_instance_id = 32
  affinity_maps = [
    {
      name         = "22"
      bit_position = 4
    }
  ]
  flex_algos = [
    {
      algorithm_number     = 128
      advertise_definition = true
      metric_type_delay    = true
    }
  ]
  nets = [
    {
      net_id = "49.0001.2222.2222.2222.00"
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

resource "iosxr_router_isis" "example" {
  process_id                            = "P1"
  is_type                               = "level-1"
  set_overload_bit                      = true
  set_overload_bit_on_startup_seconds   = 300
  set_overload_bit_advertise_external   = true
  set_overload_bit_advertise_interlevel = true
  set_overload_bit_levels = [
    {
      level_number            = 1
      on_startup_time_seconds = 300
      advertise_external      = true
      advertise_interlevel    = true
    }
  ]
  nsr                               = true
  nsf_ietf                          = true
  nsf_lifetime                      = 10
  nsf_interface_timer               = 5
  nsf_interface_expires             = 2
  log_adjacency_changes             = true
  lsp_gen_interval_maximum_wait     = 5000
  lsp_gen_interval_initial_wait     = 50
  lsp_gen_interval_secondary_wait   = 200
  lsp_refresh_interval              = 65000
  max_lsp_lifetime                  = 65535
  lsp_password_text_encrypted       = "060506324F41584B564B0F49584B"
  distribute_link_state_instance_id = 32
  affinity_maps = [
    {
      name         = "22"
      bit_position = 4
    }
  ]
  flex_algos = [
    {
      flex_algo_number     = 128
      advertise_definition = true
      metric_type          = "delay"
    }
  ]
  nets = [
    {
      net_id = "49.0001.2222.2222.2222.00"
    }
  ]
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
      circuit_type   = "level-1"
      hello_padding  = "always"
      priority_levels = [
        {
          level_number = 1
          priority     = 10
        }
      ]
      point_to_point = false
      state          = "passive"
    }
  ]
}

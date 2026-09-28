resource "iosxr_performance_measurement_liveness_profile" "example" {
  endpoint_default                                         = true
  endpoint_default_liveness_detection_logging_state_change = true
  endpoint_default_liveness_detection_multiplier           = 5
  endpoint_default_probe_flow_label_from                   = 100
  endpoint_default_probe_flow_label_increment              = 50
  endpoint_default_probe_flow_label_to                     = 500
  endpoint_default_probe_sweep_destination_ipv4            = "127.0.0.1"
  endpoint_default_probe_sweep_destination_range           = 10
  endpoint_default_probe_tos_dscp                          = 48
  endpoint_default_probe_tx_interval                       = 30000
  profiles = [
    {
      liveness_detection_logging_state_change = true
      liveness_detection_multiplier           = 5
      probe_flow_label_from                   = 100
      probe_flow_label_increment              = 50
      probe_flow_label_to                     = 500
      probe_sweep_destination_ipv4            = "127.0.0.1"
      probe_sweep_destination_range           = 10
      probe_tos_dscp                          = 48
      probe_tx_interval                       = 30000
      profile_name                            = "LIVENESS_PROFILE_1"
    }
  ]
  sr_policy_default                               = true
  sr_policy_default_liveness_detection_multiplier = 5
  sr_policy_default_probe_flow_label_from         = 100
  sr_policy_default_probe_flow_label_increment    = 50
  sr_policy_default_probe_flow_label_to           = 500
  sr_policy_default_probe_sweep_destination_ipv4  = "127.0.0.1"
  sr_policy_default_probe_sweep_destination_range = 10
  sr_policy_default_probe_tos_dscp                = 48
  sr_policy_default_probe_tx_interval             = 30000
}

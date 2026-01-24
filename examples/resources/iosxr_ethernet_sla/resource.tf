resource "iosxr_ethernet_sla" "example" {
  profile_name                             = "SLA-PROFILE-1"
  type                                     = "cfm-delay-measurement"
  probe_send_burst_every_interval          = 30
  probe_send_burst_every_seconds           = true
  probe_send_burst_packet_count            = 5
  probe_send_burst_packet_interval_seconds = 1
  probe_packet_size                        = 1500
  probe_packet_test_pattern_hex            = 4369
  probe_priority                           = 7
  statistics_measure = [
    {
      type                                         = "round-trip-delay"
      aggregate_bins                               = 5
      aggregate_width                              = 1000
      aggregate_usec                               = true
      buckets_size                                 = 5
      buckets_probes                               = true
      buckets_archive                              = 10
      thresholds_stateful_log_on_max_value         = 1000
      thresholds_stateful_log_on_mean_value        = 1000
      thresholds_stateful_log_on_sample_count      = 10
      thresholds_stateful_log_on_in_and_above_bin  = 5
      thresholds_stateful_efd_on_max_value         = 2000
      thresholds_stateful_efd_on_mean_value        = 2000
      thresholds_stateful_efd_on_sample_count      = 10
      thresholds_stateful_efd_on_in_and_above_bin  = 5
      thresholds_stateless_log_on_max_value        = 2000
      thresholds_stateless_log_on_mean_value       = 2000
      thresholds_stateless_log_on_sample_count     = 10
      thresholds_stateless_log_on_in_and_above_bin = 5
    }
  ]
  schedule_every_minutes  = 1
  schedule_every_for_time = 1
  schedule_every_for_unit = "minutes"
}

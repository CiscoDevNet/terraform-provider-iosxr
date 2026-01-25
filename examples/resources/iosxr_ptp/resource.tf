resource "iosxr_ptp" "example" {
  frequency_priority = 10
  time_of_day_priority = 10
  ipv6_verify_checksum = true
  min_clock_class = 128
  utc_offset_baseline = 37
  utc_offsets = [
    {
      date = "2025-12-01"
      offset_value = 37
    }
  ]
  uncalibrated_clock_class_clock_class = 248
  uncalibrated_clock_class_unless_from_holdover = true
  uncalibrated_traceable_override = true
  startup_clock_class = 165
  freerun_clock_class = 249
  double_failure_clock_class = 250
  physical_layer_frequency = true
  network_type_high_pdv = true
  servo_slow_tracking = 16
  holdover_spec_clock_class = 164
  holdover_spec_duration = 600
  holdover_spec_traceable_override = true
  apts = true
  performance_monitoring = true
  log_best_primary_clock_changes = true
  log_servo_events = true
  virtual_port = true
  virtual_port_priority1 = 128
  virtual_port_priority2 = 129
  virtual_port_clock_class = 6
  virtual_port_clock_accuracy = 33
  virtual_port_offset_scaled_log_variance = 4100
  virtual_port_local_priority = 128
  virtual_port_gm_threshold_breach = 1000000
  clock_identity_mac_address_custom = "00:11:22:33:44:55"
  clock_domain = 24
  clock_priority1 = 128
  clock_priority2 = 128
  clock_clock_class = 6
  clock_timescale_ptp = true
  clock_time_source_gps = true
  clock_profile_g_8275_1_clock_type_t_bc = true
}

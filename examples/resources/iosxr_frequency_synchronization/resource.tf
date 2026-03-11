resource "iosxr_frequency_synchronization" "example" {
  quality_itu_t_option_one = true
  clock_interface_timing_mode_system = true
  system_timing_mode_clock_only = true
  log_selection_changes = true
  clock_identity_mac_address = "00:11:22:33:44:55"
}

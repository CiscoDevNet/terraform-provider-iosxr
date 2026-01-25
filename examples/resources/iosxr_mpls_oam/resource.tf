resource "iosxr_mpls_oam" "example" {
  oam = true
  oam_echo_disable_vendor_extension = true
  oam_echo_reply_mode_control_channel_allow_reverse_lsp = true
  oam_echo_revision_four = true
  oam_dpm_pps = 10
  oam_dpm_interval = 60
  oam_dpm_downstream_ecmp_faults = true
}

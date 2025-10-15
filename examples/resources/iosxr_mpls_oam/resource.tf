resource "iosxr_mpls_oam" "example" {
  oam                                                   = true
  oam_echo_disable_vendor_extension                     = true
  oam_echo_reply_mode_control_channel_allow_reverse_lsp = true
  oam_dpm_pps                                           = 10
  oam_dpm_interval                                      = 60
}

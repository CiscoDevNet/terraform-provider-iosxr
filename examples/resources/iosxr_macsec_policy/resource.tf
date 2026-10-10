resource "iosxr_macsec_policy" "example" {
  allow_lacp_in_clear                = true
  allow_lldp_in_clear                = true
  allow_pause_frame_in_clear         = true
  cipher_suite                       = "GCM-AES-256"
  conf_offset                        = "CONF-OFFSET-50"
  delay_protection                   = true
  enable_legacy_fallback             = true
  enable_legacy_sak_write            = true
  impose_overhead_on_bundle          = true
  include_icv_indicator              = true
  key_server_priority                = 100
  logging_sak_rekey_disable          = true
  logging_sak_rekey_summary_interval = 60
  max_an                             = "1"
  policy_exception                   = "lacp-in-clear"
  policy_name                        = "POLICY1"
  ppk                                = true
  ppk_sks_profile                    = "SKS-PROFILE1"
  sak_rekey_interval_seconds         = 1800
  security_policy                    = "must-secure"
  suspend_for_disable                = true
  suspend_on_request_disable         = true
  use_eapol_pae_in_icv               = true
  vlan_tags_in_clear                 = 1
  window_size                        = 512
}

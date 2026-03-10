resource "iosxr_macsec_policy" "example" {
  policy_name                        = "POLICY1"
  key_server_priority                = 100
  cipher_suite                       = "GCM-AES-256"
  window_size                        = 512
  conf_offset                        = "CONF-OFFSET-50"
  security_policy                    = "must-secure"
  vlan_tags_in_clear                 = 1
  policy_exception                   = "lacp-in-clear"
  sak_rekey_interval_seconds         = 1800
  include_icv_indicator              = true
  delay_protection                   = true
  use_eapol_pae_in_icv               = true
  suspend_on_request_disable         = true
  suspend_for_disable                = true
  enable_legacy_fallback             = true
  enable_legacy_sak_write            = true
  impose_overhead_on_bundle          = true
  max_an                             = "1"
  allow_lacp_in_clear                = true
  allow_pause_frame_in_clear         = true
  allow_lldp_in_clear                = true
  ppk                                = true
  ppk_sks_profile                    = "SKS-PROFILE1"
  logging_sak_rekey_disable          = true
  logging_sak_rekey_summary_interval = 60
}

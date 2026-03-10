resource "iosxr_error_disable_recovery" "example" {
  link_oam_session_down_interval          = 30
  link_oam_discovery_timeout_interval     = 40
  link_oam_capabilities_conflict_interval = 50
  link_oam_miswired_interval              = 60
  link_oam_link_fault_interval            = 70
  link_oam_dying_gasp_interval            = 80
  link_oam_critical_event_interval        = 90
  link_oam_threshold_breached_interval    = 92
  stp_bpdu_guard_interval                 = 94
  stp_legacy_bpdu_interval                = 96
  cluster_udld_interval                   = 98
  cluster_minlinks_interval               = 100
  udld_unidirectional_interval            = 110
  udld_neighbor_mismatch_interval         = 120
  udld_timeout_interval                   = 140
  udld_loopback_interval                  = 150
  pvrst_pvid_mismatch_interval            = 160
  l2vpn_bport_mac_move_interval           = 170
  ot_track_state_change_interval          = 180
  link_oam_dampening_interval             = 190
}

resource "iosxr_router_isis_interface" "example" {
  process_id     = "P1"
  interface_name = "GigabitEthernet0/0/0/1"
  mesh_group     = 1
  state          = "passive"
  circuit_type   = "level-1"
  csnp_interval  = 10
  csnp_interval_levels = [
    {
      level_number  = 1
      csnp_interval = 10
    }
  ]
  hello_padding = "disable"
  hello_padding_levels = [
    {
      level_number  = 1
      hello_padding = "always"
    }
  ]
  hello_interval = 5
  hello_interval_levels = [
    {
      level_number   = 1
      hello_interval = 5
    }
  ]
  hello_multiplier = 5
  hello_multiplier_levels = [
    {
      level_number     = 1
      hello_multiplier = 5
    }
  ]
  lsp_interval = 10
  lsp_interval_levels = [
    {
      level_number = 1
      lsp_interval = 10
    }
  ]
  hello_password_hmac_md5_encrypted = "060506324F41584B564B0F49584B"
  hello_password_hmac_md5_send_only = true
  hello_password_levels = [
    {
      level_number   = 1
      text_encrypted = "060506324F41584B564B0F49584B"
      text_send_only = true
    }
  ]
  remote_psnp_delay = 1000
  priority          = 10
  priority_levels = [
    {
      level_number = 1
      priority     = 10
    }
  ]
  point_to_point      = true
  retransmit_interval = 50
  retransmit_interval_levels = [
    {
      level_number        = 1
      retransmit_interval = 50
    }
  ]
  retransmit_throttle_interval = 10000
  retransmit_throttle_interval_levels = [
    {
      level_number                 = 1
      retransmit_throttle_interval = 10000
    }
  ]
  link_down_fast_detect         = true
  affinity_flex_algos           = ["AFFINITY-1"]
  affinity_flex_algos_anomalies = ["AFFINITY-2"]
  override_metrics              = "high"
  delay_normalize_interval      = 10000
  delay_normalize_offset        = 1000
  mpls_ldp_sync                 = true
  mpls_ldp_sync_level           = 1
  bfd_fast_detect_ipv4          = true
  bfd_fast_detect_ipv6          = true
  bfd_minimum_interval          = 50
  bfd_multiplier                = 3
}

resource "iosxr_cdp" "example" {
  enable                = true
  holdtime              = 12
  timer                 = 34
  advertise_v1          = true
  log_adjacency_changes = true
}

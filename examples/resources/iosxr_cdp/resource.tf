resource "iosxr_cdp" "example" {
  advertise_v1          = true
  enable                = true
  holdtime              = 12
  log_adjacency_changes = true
  timer                 = 34
}

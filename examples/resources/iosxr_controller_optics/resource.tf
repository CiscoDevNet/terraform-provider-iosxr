resource "iosxr_controller_optics" "example" {
  active   = "act"
  breakout = "4x25"
  name     = "0/0/0/1"
  shutdown = true
  speed    = "10g"
  type     = "Optics"
}

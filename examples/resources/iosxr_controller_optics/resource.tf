resource "iosxr_controller_optics" "example" {
  type = "Optics"
  name = "0/0/0/1"
  active = "act"
  shutdown = true
  speed = "10g"
  breakout = "4x25"
}

resource "iosxr_banner" "example" {
  banner_type = "login"
  line        = " Hello user !"
}

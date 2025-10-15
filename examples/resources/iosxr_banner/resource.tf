resource "iosxr_banner" "example" {
  banner_type = "login"
  line        = "^C Hello World! ^C"
}

resource "iosxr_banner" "example" {
  banner_type = "prompt-timeout"
  line        = ", banner-text ,"
}

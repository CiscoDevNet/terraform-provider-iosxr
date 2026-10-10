resource "iosxr_cef_pbts_forward_class" "example" {
  fallback_to_drop = true
  forward_class    = "1"
}

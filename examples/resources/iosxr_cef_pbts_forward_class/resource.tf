resource "iosxr_cef_pbts_forward_class" "example" {
  forward_class = "1"
  fallback_to_drop = true
}

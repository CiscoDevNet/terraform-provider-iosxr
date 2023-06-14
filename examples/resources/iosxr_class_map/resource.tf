resource "iosxr_class_map" "example" {
  class_map_name                        = "TEST"
  match_any                             = true
  description                           = "description1"
  match_dscp_value                      = "46"
  match_mpls_experimental_topmost_label = 5
}

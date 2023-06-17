resource "iosxr_class_map_qos" "example" {
  class_map_name                  = "TEST"
  match_any                       = true
  description                     = "description1"
  match_dscp                      = ["46"]
  match_mpls_experimental_topmost = [5]
}

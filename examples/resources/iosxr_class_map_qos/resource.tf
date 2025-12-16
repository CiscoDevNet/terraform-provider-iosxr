resource "iosxr_class_map_qos" "example" {
  class_map_name                  = "TEST"
  match_any                       = true
  description                     = "description1"
  match_dscp                      = ["46"]
  match_dscp_ipv4                 = ["46"]
  match_dscp_ipv6                 = ["46"]
  match_mpls_experimental_topmost = [5]
  match_precedence                = ["5"]
  match_precedence_ipv4           = ["5"]
  match_precedence_ipv6           = ["5"]
  match_qos_group                 = ["1"]
}

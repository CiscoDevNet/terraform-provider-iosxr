resource "iosxr_class_map_qos" "example" {
  class_map_name                  = "CM-QOS"
  match_any                       = true
  description                     = "QoS Class Map"
  match_cos_inner                 = [4]
  match_discard_class             = [1]
  match_dscp                      = ["46"]
  match_dscp_ipv4                 = ["46"]
  match_dscp_ipv6                 = ["46"]
  match_mpls_experimental_topmost = [5]
  match_precedence                = ["5"]
  match_qos_group                 = ["1"]
}

resource "iosxr_qos_policy_map" "example" {
  policy_map_name                     = "core-ingress-classifier"
  class_name                          = "class-default"
  class_type                          = "qos"
  class_set_mpls_experimental_topmost = 0
  class_set_dscp                      = "0"
}

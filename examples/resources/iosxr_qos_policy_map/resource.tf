resource "iosxr_qos_policy_map" "example" {
  policy_map_name               = "core-ingress-classifier"
  name                          = "class-default"
  type                          = "qos"
  set_mpls_experimental_topmost = 0
  set_dscp                      = "0"
}

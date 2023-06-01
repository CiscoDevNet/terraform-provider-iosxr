resource "iosxr_qos_policy_map" "example" {
  policy_map_name                     = "core-ingress-classifier"
  class_name                          = "class-default"
  class_type                          = "qos"
  class_set_mpls_experimental_topmost = 0
  class_set_dscp                      = "0"
  class_queue_limits_queue_limit = [
    {
      value = "100"
      unit  = "us"
    }
  ]
  class_service_policy_name      = "SERVICEPOLICY"
  class_police_rate_value        = "5"
  class_police_rate_unit         = "gbps"
  class_shape_average_rate_value = "100"
  class_shape_average_rate_unit  = "gbps"
}

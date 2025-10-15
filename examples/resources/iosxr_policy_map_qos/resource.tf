resource "iosxr_policy_map_qos" "example" {
  policy_map_name = "PM1"
  description     = "My description"
  classes = [
    {
      name                          = "class-default"
      type                          = "qos"
      set_mpls_experimental_topmost = 0
      set_dscp                      = "0"
      priority_level                = 1
      queue_limits = [
        {
          value = "100"
          unit  = "us"
        }
      ]
      service_policy_name = "SERVICEPOLICY"
      police_rate_value   = "5"
      police_rate_unit    = "gbps"
    }
  ]
}

resource "iosxr_policy_map_qos" "example" {
  policy_map_name = "PM1"
  description     = "My description"
  classes = [
    {
      name                   = "class-default"
      type                   = "qos"
      police_rate_value      = "5"
      police_rate_unit       = "gbps"
      police_peak_rate_value = "6"
      police_peak_rate_unit  = "gbps"
      priority_level         = 1
      queue_limits = [
        {
          value = "100"
          unit  = "ms"
        }
      ]
      service_policy_name           = "CHILD_POLICY"
      set_mpls_experimental_topmost = 5
    }
  ]
}

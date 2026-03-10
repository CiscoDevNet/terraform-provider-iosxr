resource "iosxr_policy_map_qos" "example" {
  policy_map_name = "PM-QOS"
  description     = "My description"
  classes = [
    {
      name                    = "class-default"
      type                    = "qos"
      police_rate_value       = "5"
      police_rate_unit        = "gbps"
      police_burst_value      = 500
      police_burst_unit       = "bytes"
      police_peak_rate_value  = "6"
      police_peak_rate_unit   = "gbps"
      police_peak_burst_value = 1000
      police_peak_burst_unit  = "bytes"
      priority_level          = 1
      queue_limits = [
        {
          value = "100"
          unit  = "ms"
        }
      ]
      random_detect_default = true
      random_detect = [
        {
          minimum_threshold_value = 100
          minimum_threshold_unit  = "ms"
          maximum_threshold_value = 200
          maximum_threshold_unit  = "ms"
        }
      ]
      service_policy_name           = "CHILD_POLICY"
      set_traffic_class             = 1
      set_discard_class             = 1
      set_mpls_experimental_topmost = 5
    }
  ]
}

resource "iosxr_policy_map_qos" "example" {
  classes = [
    {
      name                    = "class-default"
      police_burst_unit       = "bytes"
      police_burst_value      = 500
      police_peak_burst_unit  = "bytes"
      police_peak_burst_value = 1000
      police_peak_rate_unit   = "gbps"
      police_peak_rate_value  = "6"
      police_rate_unit        = "gbps"
      police_rate_value       = "5"
      priority_level          = 1
      queue_limits = [
        {
          unit  = "ms"
          value = "100"
        }
      ]
      random_detect = [
        {
          maximum_threshold_unit  = "ms"
          maximum_threshold_value = 200
          minimum_threshold_unit  = "ms"
          minimum_threshold_value = 100
        }
      ]
      random_detect_ecn             = true
      service_policy_name           = "CHILD_POLICY"
      set_discard_class             = 1
      set_mpls_experimental_topmost = 5
      set_traffic_class             = 1
      type                          = "qos"
    }
  ]
  description     = "My description"
  policy_map_name = "PM-QOS"
}

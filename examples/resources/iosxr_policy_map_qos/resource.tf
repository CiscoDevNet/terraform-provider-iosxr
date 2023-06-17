resource "iosxr_policy_map_qos" "example" {
  policy_map_name = "PM1"
  description     = "My description"
  classes = [
    {
      name                          = "class-default"
      type                          = "qos"
      set_mpls_experimental_topmost = 0
      set_dscp                      = "0"
      queue_limits = [
        {
          value = "100"
          unit  = "us"
        }
      ]
    }
  ]
}

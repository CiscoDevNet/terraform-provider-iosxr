resource "iosxr_router_igmp_interface" "example" {
  interface_name                         = "GigabitEthernet0/0/0/1"
  version                                = 3
  router_enable                          = true
  dvmrp_enable                           = true
  query_interval                         = 125
  query_timeout                          = 255
  query_max_response_time                = 10
  explicit_tracking_enable               = true
  explicit_tracking_acl                  = "IGMP_ACL"
  access_group                           = "IGMP_ACL"
  maximum_groups_per_interface           = 25000
  maximum_groups_per_interface_threshold = 20000
  maximum_groups_per_interface_acl       = "IGMP_ACL"
  static_groups = [
    {
      group_address      = "239.1.1.1"
      group_address_only = true
      suppress_reports   = true
    }
  ]
  join_groups = [
    {
      group_address = "239.1.1.100"
      source_addresses = [
        {
          source_ip = "10.1.1.1"
          include   = true
        }
      ]
    }
  ]
}

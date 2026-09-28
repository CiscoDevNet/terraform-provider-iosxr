resource "iosxr_router_igmp_interface" "example" {
  access_group             = "IGMP_ACL"
  dvmrp_enable             = true
  explicit_tracking_acl    = "IGMP_ACL"
  explicit_tracking_enable = true
  interface_name           = "GigabitEthernet0/0/0/1"
  join_groups = [
    {
      group_address = "239.1.1.100"
      source_addresses = [
        {
          include   = true
          source_ip = "10.1.1.1"
        }
      ]
    }
  ]
  maximum_groups_per_interface           = 25000
  maximum_groups_per_interface_acl       = "IGMP_ACL"
  maximum_groups_per_interface_threshold = 20000
  query_interval                         = 125
  query_max_response_time                = 10
  query_timeout                          = 255
  router_enable                          = true
  static_groups = [
    {
      group_address      = "239.1.1.1"
      group_address_only = true
      suppress_reports   = true
    }
  ]
  version = 3
}

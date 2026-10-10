resource "iosxr_router_mld_vrf_interface" "example" {
  access_group             = "MLD_ACL"
  dvmrp_enable             = true
  explicit_tracking_acl    = "MLD_ACL"
  explicit_tracking_enable = true
  interface_name           = "GigabitEthernet0/0/0/1"
  join_groups = [
    {
      group_address = "ff3e::100"
      source_addresses = [
        {
          include   = true
          source_ip = "2001:db8::1"
        }
      ]
    }
  ]
  maximum_groups_per_interface           = 25000
  maximum_groups_per_interface_acl       = "MLD_ACL"
  maximum_groups_per_interface_threshold = 20000
  query_interval                         = 125
  query_max_response_time                = 10
  query_timeout                          = 255
  router_enable                          = true
  static_groups = [
    {
      group_address      = "ff3e::1"
      group_address_only = true
      suppress_reports   = true
    }
  ]
  version  = 2
  vrf_name = "VRF1"
}

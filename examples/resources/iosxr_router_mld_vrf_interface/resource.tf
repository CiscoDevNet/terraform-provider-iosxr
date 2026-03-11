resource "iosxr_router_mld_vrf_interface" "example" {
  vrf_name = "VRF1"
  interface_name = "GigabitEthernet0/0/0/1"
  version = 2
  router_enable = true
  query_interval = 125
  query_timeout = 255
  query_max_response_time = 10
  explicit_tracking_enable = true
  explicit_tracking_acl = "MLD_ACL"
  access_group = "MLD_ACL"
  maximum_groups_per_interface = 25000
  maximum_groups_per_interface_threshold = 20000
  maximum_groups_per_interface_acl = "MLD_ACL"
  static_groups = [
    {
      group_address = "ff3e::1"
      group_address_only = true
      suppress_reports = true
    }
  ]
  join_groups = [
    {
      group_address = "ff3e::100"
        source_addresses = [
          {
            source_ip = "2001:db8::1"
            include = true
          }
        ]
    }
  ]
  dvmrp_enable = true
}

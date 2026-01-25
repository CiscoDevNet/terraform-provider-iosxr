resource "iosxr_router_igmp" "example" {
  accounting_max_history                 = 30
  nsf_lifetime                           = 300
  dvmrp_enable                           = true
  robustness_count                       = 5
  maximum_groups                         = 50000
  maximum_groups_per_interface           = 25000
  maximum_groups_per_interface_threshold = 20000
  maximum_groups_per_interface_acl       = "IGMP_ACL"
  version                                = 3
  query_interval                         = 125
  query_timeout                          = 255
  query_max_response_time                = 10
  explicit_tracking                      = true
  explicit_tracking_acl                  = "IGMP_ACL"
  access_group                           = "IGMP_ACL"
  ssm_map_statics = [
    {
      address     = "10.1.1.1"
      access_list = "SSM_ACL"
    }
  ]
  ssm_map_query_dns             = true
  missed_packets_gen_query      = 5000
  missed_packets_grp_spec_query = 5000
  missed_packets_ssm_query      = 5000
  missed_packets_member_report  = 5000
}

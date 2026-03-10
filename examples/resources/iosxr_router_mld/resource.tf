resource "iosxr_router_mld" "example" {
  nsf_lifetime = 300
  robustness_variable = 5
  maximum_groups = 50000
  maximum_groups_per_interface = 25000
  maximum_groups_per_interface_threshold = 20000
  maximum_groups_per_interface_acl = "MLD_ACL"
  version = 2
  query_interval = 125
  query_timeout = 255
  query_max_response_time = 10
  explicit_tracking = true
  explicit_tracking_acl = "MLD_ACL"
  access_group = "MLD_ACL"
  ssm_map_statics = [
    {
      address = "2001:db8::1"
      access_list = "SSM_ACL"
    }
  ]
  ssm_map_query_dns = true
  missed_packets_gen_query = 5000
  missed_packets_grp_spec_query = 5000
  missed_packets_ssm_query = 5000
  missed_packets_member_report = 5000
}

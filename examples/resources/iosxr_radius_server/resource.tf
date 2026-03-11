resource "iosxr_radius_server" "example" {
  hosts = [
    {
      address = "10.1.1.1"
      auth_port = 1812
      acct_port = 1813
      timeout = 120
      retransmit = 5
      key_type_7 = "060506324F41584B"
      test_username = "cisco"
      idle_time = 30
      ignore_auth_port = true
      ignore_acct_port = true
    }
  ]
  key_type_7 = "060506324F41584B"
  timeout = 120
  retransmit_retries = 5
  load_balance_method_least_outstanding_batch_size = 25
  load_balance_method_least_outstanding_ignore_preferred_server = true
  throttle_access = 100
  throttle_access_timeout = 5
  throttle_accounting = 50
  deadtime = 10
  dead_criteria_time = 10
  dead_criteria_tries = 5
  source_port_extended = true
  ipv4_dscp = "cs6"
  ipv6_dscp = "cs6"
  vsa_attribute_ignore_unknown = true
  disallow_null_username = true
  attribute_lists = [
    {
      name = "ATTR-LIST-1"
      radius_attributes = "1,2,3,4,5"
        attribute_vendor_ids = [
          {
            id = 9
              vendor_types = [
                {
                  vendor_type_id = 1
                }
              ]
          }
        ]
    }
  ]
  attribute_acct_session_id_prepend_nas_port_id = true
  attribute_acct_multi_session_id_include_parent_session_id = true
  attribute_filter_id_11_default_direction = "inbound"
}

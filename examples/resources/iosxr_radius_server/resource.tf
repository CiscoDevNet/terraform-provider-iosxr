resource "iosxr_radius_server" "example" {
  attribute_acct_multi_session_id_include_parent_session_id = true
  attribute_acct_session_id_prepend_nas_port_id             = true
  attribute_filter_id_11_default_direction                  = "inbound"
  attribute_lists = [
    {
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
      name              = "ATTR-LIST-1"
      radius_attributes = "1,2,3,4,5"
    }
  ]
  dead_criteria_time     = 10
  dead_criteria_tries    = 5
  deadtime               = 10
  disallow_null_username = true
  hosts = [
    {
      acct_port        = 1813
      address          = "10.1.1.1"
      auth_port        = 1812
      idle_time        = 30
      ignore_acct_port = true
      ignore_auth_port = true
      key_type_7       = "060506324F41584B"
      retransmit       = 5
      test_username    = "cisco"
      timeout          = 120
    }
  ]
  ipv4_dscp                                                     = "cs6"
  ipv6_dscp                                                     = "cs6"
  key_type_7                                                    = "060506324F41584B"
  load_balance_method_least_outstanding_batch_size              = 25
  load_balance_method_least_outstanding_ignore_preferred_server = true
  retransmit_retries                                            = 5
  source_port_extended                                          = true
  throttle_access                                               = 100
  throttle_access_timeout                                       = 5
  throttle_accounting                                           = 50
  timeout                                                       = 120
  vsa_attribute_ignore_unknown                                  = true
}

resource "iosxr_aaa" "example" {
  tacacs_server_hosts = [
    {
      address                        = "9.0.1.68"
      port                           = 49
      timeout                        = 10
      holddown_time                  = 300
      key_type_7                     = "0235347225301B204F4F0A0A"
      single_connection              = true
      single_connection_idle_timeout = 1000
    }
  ]
  tacacs_server_key_type_7    = "0235347225301B204F4F0A0A"
  tacacs_server_timeout       = 5
  tacacs_server_holddown_time = 600
  tacacs_server_ipv4_dscp     = "cs6"
  tacacs_server_ipv6_dscp     = "cs6"
  tacacs_server_groups = [
    {
      group_name = "AAA"
      servers = [
        {
          address = "9.0.1.68"
        }
      ]
      vrf           = "VRF1"
      holddown_time = 300
      server_privates = [
        {
          address                        = "9.0.1.68"
          port                           = 49
          key_type_7                     = "0235347225301B204F4F0A0A"
          single_connection              = true
          single_connection_idle_timeout = 1000
          timeout                        = 10
          holddown_time                  = 300
        }
      ]
    }
  ]
  tacacs_source_interface = "MgmtEth0/RP0/CPU0/0"
  tacacs_vrfs = [
    {
      vrf_name         = "VRF1"
      source_interface = "MgmtEth0/RP0/CPU0/0"
    }
  ]
  usernames = [
    {
      name                 = "terraform-user"
      secret_type_10       = "$6$JX.ze16kjuwfCe1.$aAaA/Zc0/lI2C7VrT4PmeBh.tsQj1fKAR3q.Oqq214vrnCLjhezD9/asda5WB8tbO4wqKp3MsA5cv3B2uHJOj/"
      login_history_enable = true
      group_root_lr        = true
    }
  ]
}

resource "iosxr_call_home" "example" {
  service_active = true
  mail_servers = [
    {
      mail_server_name = "smtp.example.com"
      priority = 1
    }
  ]
  sender_from = "router@example.com"
  sender_reply_to = "admin@example.com"
  contact_email = "admin@example.com"
  phone_number = "+14085551234"
  street_address = "170 West Tasman Drive, San Jose, CA 95134"
  customer_id = "CUST12345"
  contract_id = "CONTRACT67890"
  site_id = "SITE001"
  rate_limit = 3
  data_privacy_level_high = true
  http_proxy_name = "proxy.example.com"
  http_proxy_port = 8080
  source_interface = "Loopback0"
  syslog_throttling = true
  vrf = "OOB"
  aaa_authorization = true
  aaa_authorization_username = "callhome"
  profiles = [
    {
      profile_name = "cisco-sl"
      active = true
        destination_addresses = [
          {
            address_type = "http"
            destination_address = "https://cisco-license.customer.com/Transportgateway/services/DeviceRequestHandler"
          }
        ]
      destination_message_size_limit = 100000
      destination_msg_format_long = true
      destination_transport_method_email_disable = true
      destination_transport_method_http = true
      reporting_smart_call_home_data = true
      reporting_smart_licensing_data = true
      anonymous_reporting_only = true
    }
  ]
}

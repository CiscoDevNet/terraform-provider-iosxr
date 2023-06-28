resource "iosxr_domain" "example" {
  list_domain = [
    {
      domain_name = "DOMAIN1"
      order       = 0
    }
  ]
  lookup_disable          = true
  lookup_source_interface = "Loopback2147483647"
  name                    = "DOMAIN"
  ipv4_host = [
    {
      host_name  = "HOST_NAME"
      ip_address = ["10.0.0.0"]
    }
  ]
  name_server = [
    {
      address = "10.0.0.1"
      order   = 345
    }
  ]
  ipv6_host = [
    {
      host_name    = "HOST_NAME_IPV6"
      ipv6_address = ["10::10"]
    }
  ]
  multicast             = "DOMAIN1_ACC"
  default_flows_disable = true
}

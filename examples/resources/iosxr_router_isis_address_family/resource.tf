resource "iosxr_router_isis_address_family" "example" {
  process_id                    = "P1"
  af_name                       = "ipv4"
  saf_name                      = "unicast"
  mpls_ldp_auto_config          = false
  metric_style_narrow           = false
  metric_style_wide             = true
  metric_style_transition       = false
  router_id_ip_address          = "1.2.3.4"
  default_information_originate = true
}

data "iosxr_router_isis_address_family" "example" {
  process_id = "P1"
  af_name    = "ipv4"
  saf_name   = "unicast"
}

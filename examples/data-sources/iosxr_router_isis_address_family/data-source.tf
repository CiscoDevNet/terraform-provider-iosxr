data "iosxr_router_isis_address_family" "example" {
  af_name    = "ipv4"
  process_id = "P1"
  saf_name   = "unicast"
}

resource "iosxr_router_bgp_address_family_redistribute_isis" "example" {
  as_number                    = "65001"
  af_name                      = "ipv4-unicast"
  instance_name                = "P1"
  level_one                    = true
  level_one_two                = true
  level_one_two_one_inter_area = false
  level_one_one_inter_area     = false
  level_two                    = false
  level_two_one_inter_area     = false
  level_one_inter_area         = false
  metric                       = 100
}

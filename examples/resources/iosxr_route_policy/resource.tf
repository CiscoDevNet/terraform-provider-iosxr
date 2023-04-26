resource "iosxr_route_policy" "example" {
  route_policy_name = "BGP_POLICY_NAME"
  rpl               = "route-policy BGP_POLICY_NAME\n  pass\nend-policy\n"
}

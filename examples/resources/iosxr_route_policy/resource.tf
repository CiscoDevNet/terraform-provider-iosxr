resource "iosxr_route_policy" "example" {
  route_policy_name = "ROUTE_POLICY_1"
  rpl = "route-policy ROUTE_POLICY_1\n  if destination in PREFIX_SET_1 then\n    set extcommunity rt (12345:1) additive\n  endif\n  pass\nend-policy\n"
}

resource "iosxr_cef_accounting" "example" {
  interfaces_mpls_ipv4_rsvp_te                            = true
  interfaces_segment_routing_mpls_ipv4                    = true
  interfaces_segment_routing_mpls_ipv6                    = true
  prefixes_ipv6_mode_per_prefix_per_nexthop_srv6_locators = true
  segment_routing_policies_srv6_disable                   = true
}

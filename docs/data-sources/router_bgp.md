---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_bgp Data Source - terraform-provider-iosxr"
subcategory: "BGP"
description: |-
  This data source can read the Router BGP configuration.
---

# iosxr_router_bgp (Data Source)

This data source can read the Router BGP configuration.

## Example Usage

```terraform
data "iosxr_router_bgp" "example" {
  as_number = "65001"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_number` (String) bgp as-number

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `bfd_minimum_interval` (Number) Hello interval
- `bfd_multiplier` (Number) Detect multiplier
- `bgp_bestpath_aigp_ignore` (Boolean) Ignore AIGP attribute
- `bgp_bestpath_as_path_ignore` (Boolean) Ignore as-path length
- `bgp_bestpath_as_path_multipath_relax` (Boolean) Relax as-path check for multipath selection
- `bgp_bestpath_compare_routerid` (Boolean) Compare router-id for identical EBGP paths
- `bgp_bestpath_cost_community_ignore` (Boolean) Ignore cost-community comparison
- `bgp_bestpath_igp_metric_ignore` (Boolean) Ignore IGP metric during path comparison
- `bgp_bestpath_igp_metric_sr_policy` (Boolean) Use next-hop admin/metric from SR policy at Next Hop metric comparison stage
- `bgp_bestpath_med_always` (Boolean) Allow comparing MED from different neighbors
- `bgp_bestpath_med_confed` (Boolean) Compare MED among confederation paths
- `bgp_bestpath_med_missing_as_worst` (Boolean) Treat missing MED as the least preferred one
- `bgp_bestpath_origin_as_allow_invalid` (Boolean) BGP bestpath selection will allow 'invalid' origin-AS
- `bgp_bestpath_origin_as_use_validity` (Boolean) BGP bestpath selection will use origin-AS validity
- `bgp_bestpath_sr_policy_force` (Boolean) Consider only paths over SR Policy for bestpath selection, eBGP no-color ineligible
- `bgp_bestpath_sr_policy_prefer` (Boolean) Consider only paths over SR Policy for bestpath selection, eBGP no-color eligible
- `bgp_graceful_restart_graceful_reset` (Boolean) Reset gracefully if configuration change forces a peer reset
- `bgp_log_neighbor_changes_detail` (Boolean) Include extra detail in change messages
- `bgp_redistribute_internal` (Boolean) Allow redistribution of iBGP into IGPs (dangerous)
- `bgp_router_id` (String) Configure Router-id
- `default_information_originate` (Boolean) Distribute a default route
- `default_metric` (Number) default redistributed metric
- `ibgp_policy_out_enforce_modifications` (Boolean) Allow policy to modify all attributes
- `id` (String) The path of the retrieved object.
- `neighbors` (Attributes List) Neighbor address (see [below for nested schema](#nestedatt--neighbors))
- `nexthop_validation_color_extcomm_disable` (Boolean) Disable next-hop reachability validation for color-extcomm path
- `nexthop_validation_color_extcomm_sr_policy` (Boolean) Enable BGP next-hop reachability validation by SR Policy for color-extcomm paths
- `nsr_disable` (Boolean) Disable non-stop-routing support for all neighbors
- `segment_routing_srv6_locator` (String) Configure locator name
- `timers_bgp_holdtime` (String) Holdtime. Set 0 to disable keepalives/hold time.
- `timers_bgp_keepalive_interval` (Number) BGP timers
- `timers_bgp_minimum_acceptable_holdtime` (String) Minimum acceptable holdtime from neighbor. Set 0 to disable keepalives/hold time.

<a id="nestedatt--neighbors"></a>
### Nested Schema for `neighbors`

Read-Only:

- `advertisement_interval_milliseconds` (Number) time in milliseconds
- `advertisement_interval_seconds` (Number) Minimum interval between sending BGP routing updates
- `bfd_fast_detect` (Boolean) Enable Fast detection
- `bfd_fast_detect_inheritance_disable` (Boolean) Prevent bfd settings from being inherited from the parent
- `bfd_fast_detect_strict_mode` (Boolean) Hold down neighbor session until BFD session is up
- `bfd_minimum_interval` (Number) Hello interval
- `bfd_multiplier` (Number) Detect multiplier
- `description` (String) Neighbor specific description
- `ebgp_multihop_maximum_hop_count` (Number) maximum hop count
- `ignore_connected_check` (Boolean) Bypass the directly connected nexthop check for single-hop eBGP peering
- `local_as` (String) bgp as-number
- `local_as_dual_as` (Boolean) Dual-AS mode
- `local_as_no_prepend` (Boolean) Do not prepend local AS to announcements from this neighbor
- `local_as_replace_as` (Boolean) Prepend only local AS to announcements to this neighbor
- `neighbor_address` (String) Neighbor address
- `password` (String) Specifies an ENCRYPTED password will follow
- `remote_as` (String) bgp as-number
- `shutdown` (Boolean) Administratively shut down this neighbor
- `timers_holdtime` (String) Holdtime. Set 0 to disable keepalives/hold time.
- `timers_keepalive_interval` (Number) BGP timers
- `timers_minimum_acceptable_holdtime` (String) Minimum acceptable holdtime from neighbor. Set 0 to disable keepalives/hold time.
- `ttl_security` (Boolean) Enable EBGP TTL security
- `update_source` (String) Source of routing updates
- `use_neighbor_group` (String) Inherit configuration from a neighbor-group

resource "iosxr_router_bgp_vrf_address_family" "example" {
  as_number                     = "65001"
  vrf_name                      = "VRF1"
  af_name                       = "ipv4-unicast"
  maximum_paths_ebgp_multipath  = 10
  maximum_paths_ibgp_multipath  = 10
  label_mode_per_ce             = false
  label_mode_per_vrf            = false
  redistribute_connected        = true
  redistribute_connected_metric = 10
  redistribute_static           = true
  redistribute_static_metric    = 10
}

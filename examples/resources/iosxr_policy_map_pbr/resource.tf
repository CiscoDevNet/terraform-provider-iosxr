resource "iosxr_policy_map_pbr" "example" {
  classes = [
    {
      decapsulate_gre                         = true
      name                                    = "class-default"
      police_rate_unit                        = "gbps"
      police_rate_value                       = 5
      redirect_ipv4_nexthop1_address          = "192.168.253.1"
      redirect_ipv4_nexthop1_vrf              = "VRF1"
      redirect_ipv4_nexthop2_address          = "192.168.253.3"
      redirect_ipv4_nexthop2_vrf              = "VRF2"
      redirect_ipv4_nexthop3_address          = "192.168.253.5"
      redirect_ipv4_nexthop3_vrf              = "VRF3"
      redirect_ipv6_nexthop1_address          = "2001:db8:1::1"
      redirect_ipv6_nexthop1_vrf              = "VRF1"
      redirect_ipv6_nexthop2_address          = "2001:db8:2::1"
      redirect_ipv6_nexthop2_vrf              = "VRF2"
      redirect_ipv6_nexthop3_address          = "2001:db8:3::1"
      redirect_ipv6_nexthop3_vrf              = "VRF3"
      redirect_nexthop_route_target_as_format = "65001:100"
      set_dscp                                = "ef"
      set_forward_class                       = 1
      type                                    = "traffic"
    }
  ]
  description     = "PBR Policy-map"
  policy_map_name = "PM-PBR"
}

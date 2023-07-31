resource "iosxr_l2vpn_pw_class" "example" {
  name                                                           = "PWC1"
  encapsulation_mpls                                             = true
  encapsulation_mpls_transport_mode_ethernet                     = true
  encapsulation_mpls_transport_mode_vlan                         = false
  encapsulation_mpls_transport_mode_passthrough                  = false
  encapsulation_mpls_load_balancing_pw_label                     = true
  encapsulation_mpls_load_balancing_flow_label_both              = true
  encapsulation_mpls_load_balancing_flow_label_both_static       = true
  encapsulation_mpls_load_balancing_flow_label_code_one7         = true
  encapsulation_mpls_load_balancing_flow_label_code_one7_disable = true
}

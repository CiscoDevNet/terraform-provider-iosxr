resource "iosxr_l2vpn_pw_class" "example" {
  pw_class_name                                                = "PW-CLASS1"
  encapsulation_mpls                                           = true
  encapsulation_mpls_protocol_ldp                              = true
  encapsulation_mpls_control_word                              = true
  encapsulation_mpls_transport_mode_vlan                       = true
  encapsulation_mpls_vccv_verification_type_none               = true
  encapsulation_mpls_preferred_path_sr_te_policy               = "sr-policy-1"
  encapsulation_mpls_preferred_path_fallback_disable           = true
  encapsulation_mpls_tag_rewrite_ingress_vlan                  = 100
  encapsulation_mpls_redundancy_one_way                        = true
  encapsulation_mpls_redundancy_initial_delay                  = 60
  encapsulation_mpls_load_balancing_flow_label_both            = true
  encapsulation_mpls_load_balancing_flow_label_both_static     = true
  encapsulation_mpls_load_balancing_flow_label_code_17         = true
  encapsulation_mpls_load_balancing_flow_label_code_17_disable = true
  encapsulation_mpls_ipv4_source                               = "1.2.3.4"
  backup_disable_delay                                         = 10
  mac_withdraw                                                 = true
}

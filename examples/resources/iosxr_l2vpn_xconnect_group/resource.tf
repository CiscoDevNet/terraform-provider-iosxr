resource "iosxr_l2vpn_xconnect_group" "example" {
  group_name = "P2P"
  p2ps = [
    {
      p2p_xconnect_name = "XC"
      description       = "My P2P Description"
      interfaces = [
        {
          interface_name = "Bundle-Ether11"
        }
      ]
      ipv4_neighbors = [
        {
          address                  = "1.1.1.1"
          pw_id                    = 1
          pw_class                 = "PW_CLASS_1"
          bandwidth                = 1000000000
          mpls_static_label_local  = 1002
          mpls_static_label_remote = 1003
        }
      ]
    }
  ]
}

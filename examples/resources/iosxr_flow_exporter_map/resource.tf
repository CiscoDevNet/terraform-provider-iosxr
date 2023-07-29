resource "iosxr_flow_exporter_map" "example" {
  name                                    = "TEST"
  destination_ipv4_address                = "10.1.1.1"
  destination_ipv6_address                = "1::1"
  destination_vrf                         = "VRF1"
  source                                  = "GigabitEthernet0/0/0/1"
  dscp                                    = 62
  packet_length                           = 512
  transport_udp                           = 1033
  dfbit_set                               = true
  version_export_format                   = "v9"
  version_template_data_timeout           = 1024
  version_template_options_timeout        = 3033
  version_template_timeout                = 2222
  version_options_interface_table_timeout = 6048
  version_options_sampler_table_timeout   = 4096
  version_options_class_table_timeout     = 255
  version_options_vrf_table_timeout       = 122
}

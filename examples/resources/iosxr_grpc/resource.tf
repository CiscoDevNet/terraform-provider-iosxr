resource "iosxr_grpc" "example" {
  aaa_accounting_history_memory              = 100
  aaa_accounting_queue_size                  = 100
  dscp                                       = "af11"
  gnmi                                       = true
  gnmi_get_limit_multiple                    = 10
  gnmi_port                                  = 57401
  gnpsi                                      = true
  gnpsi_stream_queue_size                    = 4000
  gribi                                      = true
  gribi_port                                 = 50051
  keepalive_time                             = 300
  keepalive_timeout                          = 30
  local_connection                           = true
  max_concurrent_streams                     = 64
  max_request_per_user                       = 10
  max_request_total                          = 128
  max_streams                                = 32
  max_streams_per_user                       = 32
  memory_limit                               = 1000
  min_keepalive_interval                     = 300
  name                                       = "grpc-1"
  segment_routing                            = true
  segment_routing_traffic_eng                = true
  segment_routing_traffic_eng_policy_service = true
  segment_routing_traffic_eng_service_instances = [
    {
      ipv4_addresses = [
        {
          ipv4_address          = "192.0.2.10"
          no_tls                = true
          port                  = 50051
          source_interface      = "Loopback0"
          tls_hostname_override = "service.example.com"
          vrf                   = "default"
        }
      ]
      ipv6_addresses = [
        {
          ipv6_address          = "2001:db8::10"
          no_tls                = true
          port                  = 50051
          source_interface      = "Loopback0"
          tls_hostname_override = "service.example.com"
          vrf                   = "default"
        }
      ]
      service_instance_name = "SERVICE-1"
    }
  ]
  service_layer               = true
  service_layer_auto_register = true
  tls_cipher_default          = "enable"
  tls_max_version             = "1.3"
  tls_min_version             = "1.2"
  ttl                         = 64
}

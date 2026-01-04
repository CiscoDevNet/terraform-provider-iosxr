resource "iosxr_telemetry_model_driven" "example" {
  max_containers_per_path             = 16
  max_sensor_paths                    = 1000
  tcp_send_timeout                    = 30
  strict_timer                        = true
  gnmi_target_defined_cadence_factor  = 5
  gnmi_target_defined_minimum_cadence = 60
  include_select_leaves_on_events     = true
  include_empty_values                = true
  gnmi_heartbeat_always               = true
  gnmi_bundling                       = true
  gnmi_bundling_size                  = 1024
  destination_groups = [
    {
      name = "DEST-GROUP-1"
      vrf  = "VRF1"
      address_families = [
        {
          af_name              = "ipv4"
          address              = "10.1.1.1"
          port                 = 57500
          encoding             = "json"
          protocol_grpc        = true
          protocol_grpc_no_tls = true
          protocol_grpc_gzip   = true
        }
      ]
      destinations = [
        {
          address                 = "collector.example.com"
          port                    = 57500
          address_family          = "ipv4"
          encoding                = "self-describing-gpb"
          protocol_udp            = true
          protocol_udp_packetsize = 1024
        }
      ]
    }
  ]
  subscriptions = [
    {
      name               = "SUB-1"
      source_qos_marking = "ef"
      source_interface   = "Loopback0"
      sensor_group_ids = [
        {
          name               = "SENSOR-GROUP-1"
          heartbeat_always   = true
          heartbeat_interval = 30000
          strict_timer       = true
          sample_interval    = 0
        }
      ]
      destination_ids = [
        {
          name = "DEST-GROUP-1"
        }
      ]
      send_retry          = 5
      send_retry_duration = 10000
    }
  ]
  sensor_groups = [
    {
      name = "SENSOR-GROUP-1"
      sensor_paths = [
        {
          name = "Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters"
        }
      ]
    }
  ]
}

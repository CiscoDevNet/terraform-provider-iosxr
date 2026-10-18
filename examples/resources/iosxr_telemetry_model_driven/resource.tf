resource "iosxr_telemetry_model_driven" "example" {
  destination_groups = [
    {
      address_family = [
        {
          address              = "10.1.1.1"
          af_name              = "ipv4"
          encoding             = "json"
          port                 = 57500
          protocol_grpc        = true
          protocol_grpc_gzip   = true
          protocol_grpc_no_tls = true
        }
      ]
      destinations = [
        {
          address                 = "collector.example.com"
          address_family          = "ipv4"
          encoding                = "self-describing-gpb"
          port                    = 57500
          protocol_udp            = true
          protocol_udp_packetsize = 1024
        }
      ]
      name = "DEST-GROUP-1"
      vrf  = "VRF1"
    }
  ]
  gnmi_bundling                       = true
  gnmi_bundling_size                  = 1024
  gnmi_heartbeat_always               = true
  gnmi_target_defined_cadence_factor  = 5
  gnmi_target_defined_minimum_cadence = 60
  include_empty_values                = true
  include_select_leaves_on_events     = true
  max_containers_per_path             = 16
  max_sensor_paths                    = 1000
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
  strict_timer = true
  subscriptions = [
    {
      destination_ids = [
        {
          name = "DEST-GROUP-1"
        }
      ]
      name                = "SUB-1"
      send_retry          = 5
      send_retry_duration = 10000
      sensor_group_ids = [
        {
          heartbeat_always   = true
          heartbeat_interval = 30000
          name               = "SENSOR-GROUP-1"
          sample_interval    = 0
          strict_timer       = true
        }
      ]
      source_interface   = "Loopback0"
      source_qos_marking = "ef"
    }
  ]
  tcp_send_timeout = 30
}

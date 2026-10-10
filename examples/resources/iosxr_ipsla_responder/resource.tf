resource "iosxr_ipsla_responder" "example" {
  twamp = true
  twamp_light_sessions = [
    {
      authentication = true
      encryption     = true
      local_ipv4_addresses = [
        {
          address    = "10.1.1.1"
          local_port = 862
          remote_ipv4_addresses = [
            {
              address     = "10.1.1.2"
              remote_port = "862"
              vrf         = "default"
            }
          ]
        }
      ]
      session_id = 1
      timeout    = 3600
    }
  ]
  twamp_timeout = 600
  type_udp_ipv4 = [
    {
      address = "10.1.1.1"
      ports = [
        {
          port_number = 888
        }
      ]
    }
  ]
}

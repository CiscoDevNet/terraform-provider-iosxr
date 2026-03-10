resource "iosxr_ipsla_responder" "example" {
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
  twamp = true
  twamp_timeout = 600
  twamp_light_sessions = [
    {
      session_id = 1
        local_ipv4_addresses = [
          {
            address = "10.1.1.1"
            local_port = 862
              remote_ipv4_addresses = [
                {
                  address = "10.1.1.2"
                  remote_port = "862"
                  vrf = "default"
                }
              ]
          }
        ]
      authentication = true
      encryption = true
      timeout = 3600
    }
  ]
}

resource "iosxr_ntp" "example" {
  access_group_ipv4_peer       = "peer1"
  access_group_ipv4_query_only = "query1"
  access_group_ipv4_serve      = "serve1"
  access_group_ipv4_serve_only = "serve-only123"
  access_group_ipv6_peer       = "peer1"
  access_group_ipv6_query_only = "query1"
  access_group_ipv6_serve      = "serve1"
  access_group_ipv6_serve_only = "serve-only123"
  access_group_vrfs = [
    {
      ipv4_peer       = "peer1"
      ipv4_query_only = "query1"
      ipv4_serve      = "serve1"
      ipv4_serve_only = "serve-only123"
      ipv6_peer       = "peer1"
      ipv6_query_only = "query1"
      ipv6_serve      = "serve1"
      ipv6_serve_only = "serve-only123"
      vrf_name        = "ntp_vrf"
    }
  ]
  authenticate = true
  authentication_keys = [
    {
      key_number    = 10
      md5_encrypted = "1212000E43"
    }
  ]
  broadcastdelay = 10
  cmac_authentication_keys = [
    {
      cmac_encrypted = "135445415F59527D737D78626771475240"
      key_number     = 2
    }
  ]
  drift_aging_time = 10
  drift_file_disk0 = true
  drift_filename   = "drift.txt"
  hmac_sha1_authentication_keys = [
    {
      hmac_sha1_encrypted = "101F5B4A5142445C545D7A7A767B676074"
      key_number          = 3
    }
  ]
  hmac_sha2_authentication_keys = [
    {
      hmac_sha2_encrypted = "091D1C5A4D5041455355547B79777C6663"
      key_number          = 4
    }
  ]
  hostname_peers_servers = [
    {
      burst         = true
      fqdn_hostname = "ntp.cisco.com"
      iburst        = true
      key           = 1
      maxpoll       = 5
      minpoll       = 4
      prefer        = true
      source        = "GigabitEthernet0/0/0/1"
      type          = "peer"
      version       = 2
    }
  ]
  interfaces = [
    {
      broadcast_client      = true
      broadcast_destination = "1.2.3.4"
      broadcast_key         = 1
      broadcast_version     = 2
      interface_name        = "Bundle-Ether1"
    }
  ]
  ipv4_peers_servers = [
    {
      address = "1.2.3.4"
      burst   = true
      iburst  = true
      key     = 1
      maxpoll = 5
      minpoll = 4
      prefer  = true
      source  = "GigabitEthernet0/0/0/1"
      type    = "server"
      version = 2
    }
  ]
  ipv4_precedence = "network"
  ipv6_dscp       = "af11"
  ipv6_peers_servers = [
    {
      address      = "2001::1"
      burst        = true
      iburst       = true
      ipv6_address = "2001::1"
      key          = 1
      maxpoll      = 5
      minpoll      = 4
      prefer       = true
      source       = "GigabitEthernet0/0/0/1"
      type         = "peer"
      version      = 2
    }
  ]
  log_internal_sync = true
  max_associations  = 10
  passive           = true
  peers_servers_vrfs = [
    {
      hostname_peers_servers = [
        {
          burst         = true
          fqdn_hostname = "ntp.cisco.com"
          iburst        = true
          key           = 1
          maxpoll       = 5
          minpoll       = 4
          prefer        = true
          source        = "GigabitEthernet0/0/0/1"
          type          = "peer"
          version       = 2
        }
      ]
      ipv4_peers_servers = [
        {
          address = "1.2.3.4"
          burst   = true
          iburst  = true
          key     = 1
          maxpoll = 5
          minpoll = 4
          prefer  = true
          source  = "GigabitEthernet0/0/0/1"
          type    = "server"
          version = 2
        }
      ]
      ipv6_peers_servers = [
        {
          address      = "2001::1"
          burst        = true
          iburst       = true
          ipv6_address = "2001::1"
          key          = 1
          maxpoll      = 5
          minpoll      = 4
          prefer       = true
          source       = "GigabitEthernet0/0/0/1"
          type         = "peer"
          version      = 2
        }
      ]
      vrf_name = "vrf1"
    }
  ]
  source_interface_name = "BVI1"
  source_vrfs = [
    {
      interface_name = "BVI1"
      vrf_name       = "source_vrf"
    }
  ]
  trusted_keys = [
    {
      key_number = 8
    }
  ]
  update_calendar = true
}

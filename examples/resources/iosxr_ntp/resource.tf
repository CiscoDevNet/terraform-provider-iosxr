resource "iosxr_ntp" "example" {
  ipv4_precedence              = "network"
  ipv6_dscp                    = "af11"
  access_group_ipv6_peer       = "peer1"
  access_group_ipv6_query_only = "query1"
  access_group_ipv6_serve      = "serve1"
  access_group_ipv6_serve_only = "serve-only123"
  access_group_ipv4_peer       = "peer1"
  access_group_ipv4_query_only = "query1"
  access_group_ipv4_serve      = "serve1"
  access_group_ipv4_serve_only = "serve-only123"
  access_group_vrfs = [
    {
      vrf_name        = "ntp_vrf"
      ipv6_peer       = "peer1"
      ipv6_query_only = "query1"
      ipv6_serve      = "serve1"
      ipv6_serve_only = "serve-only123"
      ipv4_peer       = "peer1"
      ipv4_query_only = "query1"
      ipv4_serve      = "serve1"
      ipv4_serve_only = "serve-only123"
    }
  ]
  authenticate = true
  authentication_keys = [
    {
      key_number    = 10
      md5_encrypted = "1212000E43"
    }
  ]
  broadcastdelay   = 10
  max_associations = 1
  trusted_keys = [
    {
      key_number = 8
    }
  ]
  update_calendar       = true
  log_internal_sync     = true
  source_interface_name = "BVI1"
  source_vrfs = [
    {
      vrf_name       = "source_vrf"
      interface_name = "BVI1"
    }
  ]
  passive = true
  cmac_authentication_keys = [
    {
      key_number     = 2
      cmac_encrypted = "135445415F59527D737D78626771475240"
    }
  ]
  hmac_sha1_authentication_keys = [
    {
      key_number          = 3
      hmac_sha1_encrypted = "101F5B4A5142445C545D7A7A767B676074"
    }
  ]
  hmac_sha2_authentication_keys = [
    {
      key_number          = 4
      hmac_sha2_encrypted = "091D1C5A4D5041455355547B79777C6663"
    }
  ]
  ipv4_peers_servers = [
    {
      address = "1.2.3.4"
      type    = "server"
      version = 2
      key     = 1
      minpoll = 4
      maxpoll = 5
      prefer  = true
      burst   = true
      iburst  = true
      source  = "GigabitEthernet0/0/0/1"
    }
  ]
}

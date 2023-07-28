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
  vrfs = [
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
  auth_keys = [
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
}

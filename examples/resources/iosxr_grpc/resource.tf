resource "iosxr_grpc" "example" {
  aaa_accounting_history_memory = 100
  aaa_accounting_queue_size     = 100
  address_family_dual           = true
  dscp                          = "af11"
  local_connection              = true
  max_concurrent_streams        = 64
  max_request_per_user          = 10
  max_request_total             = 128
  max_streams                   = 32
  max_streams_per_user          = 32
  memory_limit                  = 1000
  port                          = 57400
}

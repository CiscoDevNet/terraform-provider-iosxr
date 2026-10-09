resource "iosxr_grpc" "example" {
  port                 = 57400
  address_family_dual  = true
  local_connection     = true
  max_request_total    = 128
  max_request_per_user = 10
  max_streams          = 32
  max_streams_per_user = 32
}

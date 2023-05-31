resource "iosxr_router_isis_interface" "example" {
  process_id              = "P1"
  interface_name          = "GigabitEthernet0/0/0/1"
  circuit_type            = "level-1"
  hello_padding_disable   = true
  hello_padding_sometimes = false
  priority                = 10
  point_to_point          = false
  passive                 = false
  suppressed              = false
  shutdown                = false
  hello_password_keychain = "KEY_CHAIN_1"
  bfd_fast_detect_ipv6    = true
}
